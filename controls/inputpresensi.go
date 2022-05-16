package controls

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func Presensi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	//TYPE INPUTAN
	type presensi struct {
		Matakuliah string `json:"matakuliah"`
		Kelas      string `json:"kelas"`
		Tanggal    string `json:"tanggal"`
	}

	//VALIDASI JSON
	var p presensi
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	//VALIDASI: MEMASTIKAN INPUTAN ADA DI JADWAL
	var column2 tabels.Jadwal
	db.Where("tanggal = ?", p.Tanggal).Where("matakuliah = ?", p.Matakuliah).Where("kelas = ?", p.Kelas).Where("akses = ?", 1).Find(&column2)
	if p.Tanggal != column2.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "kelas ini tidak dijadwalkan atau di tutup oleh dosen",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA BELUM PRESENSI
	var column tabels.Mahasiswa
	db.Where("nim = ?", c.Param("nim")).Find(&column)
	var column3 tabels.Presensi
	db.Where("nama = ?", column.Nama).Where("matakuliah = ?", p.Matakuliah).Where("tanggal = ?", p.Tanggal).Find(&column3)
	if column.Nama == column3.Nama {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": column.Nama + " sudah melakukan presensi",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA MELAKUKAN KRS SESUAI INPUTAN
	var column4 tabels.Krs
	db.Where("nama = ?", column.Nama).Where("matakuliah = ?", p.Matakuliah).Where("kelas = ?", p.Kelas).Find(&column4)
	if column.Nama != column4.Nama {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak melakukan krs untuk matakuliah ini",
		})
		return
	}

	//JIKA SUDAH LOLOS VALIDASI, MAKA DATA AKAN DI INPUTKAN
	pr := tabels.Presensi{
		Nama:       column.Nama,
		Matakuliah: p.Matakuliah,
		Kelas:      p.Kelas,
		Tanggal:    p.Tanggal,
	}
	db.Create(&pr)
	c.JSON(200, gin.H{
		"status":     "presensi berhasil",
		"nama":       pr.Nama,
		"matakuliah": pr.Matakuliah,
		"tanggal":    pr.Tanggal,
	})

	//TRIGGER PRESENSI
	column2.Presensi++
	var tabel []tabels.Jadwal
	db.Model(&tabel).Where("matakuliah = ?", pr.Matakuliah).Where("kelas = ?", pr.Kelas).Where("tanggal = ?", pr.Tanggal).Update("presensi", column2.Presensi)
	var column6 tabels.Akumulasi
	db.Where("nama = ?", pr.Nama).Where("matakuliah = ?", pr.Matakuliah).Find(&column6)
	column6.Hadir++
	var tabel2 []tabels.Akumulasi
	db.Model(&tabel2).Where("nama = ?", pr.Nama).Where("matakuliah = ?", pr.Matakuliah).Update("hadir", column6.Hadir)
}
