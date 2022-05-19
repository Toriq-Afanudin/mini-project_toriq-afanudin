package mahasiswas

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func GetPresence(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	//TYPE INPUTAN
	type presence struct {
		Matakuliah string `json:"matakuliah"`
		Kelas      string `json:"kelas"`
		Tanggal    string `json:"tanggal"`
	}

	//VALIDASI JSON
	var studentPresence presence
	if err := c.ShouldBindJSON(&studentPresence); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input is not in json form",
		})
		return
	}

	//VALIDASI: MEMASTIKAN INPUTAN ADA DI JADWAL
	var timetable tabels.Jadwal
	db.Where("tanggal = ?", studentPresence.Tanggal).Where("matakuliah = ?", studentPresence.Matakuliah).Where("kelas = ?", studentPresence.Kelas).Where("akses = ?", 1).Find(&timetable)
	if studentPresence.Tanggal != timetable.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "this class is not scheduled or closed by the lecturer",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA BELUM PRESENSI
	var column tabels.Mahasiswa
	db.Where("nim = ?", c.Param("nim")).Find(&column)
	var column3 tabels.Presensi
	db.Where("nama = ?", column.Nama).Where("matakuliah = ?", studentPresence.Matakuliah).Where("tanggal = ?", studentPresence.Tanggal).Find(&column3)
	if column.Nama == column3.Nama {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": column.Nama + " sudah melakukan presensi",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA MELAKUKAN KRS SESUAI INPUTAN
	var column4 tabels.Krs
	db.Where("nama = ?", column.Nama).Where("matakuliah = ?", studentPresence.Matakuliah).Where("kelas = ?", studentPresence.Kelas).Find(&column4)
	if column.Nama != column4.Nama {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak melakukan krs untuk matakuliah ini",
		})
		return
	}

	var column5 tabels.Krs
	db.Where("matakuliah = ?", studentPresence.Matakuliah).Where("kelas = ?", studentPresence.Kelas).Find(&column5)
	var column7 tabels.Dosen
	db.Where("gelar = ?", column5.Dosen).Find(&column7)

	//JIKA SUDAH LOLOS VALIDASI, MAKA DATA AKAN DI INPUTKAN
	pr := tabels.Presensi{
		Nama:       column.Nama,
		Matakuliah: studentPresence.Matakuliah,
		Kelas:      studentPresence.Kelas,
		Tanggal:    studentPresence.Tanggal,
		Dosen:      column7.Id,
	}
	db.Create(&pr)
	c.JSON(200, gin.H{
		"status":     "presensi berhasil",
		"nama":       pr.Nama,
		"matakuliah": pr.Matakuliah,
		"tanggal":    pr.Tanggal,
	})

	//TRIGGER PRESENSI
	timetable.Presensi++
	var tabel []tabels.Jadwal
	db.Model(&tabel).Where("matakuliah = ?", pr.Matakuliah).Where("kelas = ?", pr.Kelas).Where("tanggal = ?", pr.Tanggal).Update("presensi", timetable.Presensi)
	var column6 tabels.Akumulasi
	db.Where("nama = ?", pr.Nama).Where("matakuliah = ?", pr.Matakuliah).Find(&column6)
	column6.Hadir++
	var tabel2 []tabels.Akumulasi
	db.Model(&tabel2).Where("nama = ?", pr.Nama).Where("matakuliah = ?", pr.Matakuliah).Update("hadir", column6.Hadir)
}
