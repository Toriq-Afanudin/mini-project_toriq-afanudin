package controls

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func EditJadwal(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	//TYPE INPUTAN
	type edit struct {
		Matakuliah string `json:"matakuliah"`
		Kelas      string `json:"kelas"`
		Pertemuan  int    `json:"pertemuan"`
		Tanggal    string `json:"tanggal"`
		Jam        string `json:"jam"`
	}

	//VALIDASI JSON
	var ed edit
	if err := c.ShouldBindJSON(&ed); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	//VALIDASI TANGGAL PERKULIAHAN
	var column2 tabels.Tanggal
	db.Where("tanggal = ?", ed.Tanggal).Find(&column2)
	if ed.Tanggal != column2.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal yang di input bukan masa perkuliahan",
		})
		return
	}

	//VALIDASI:	MEMASTIKAN JAM PERKULIAHAN SESUAI DI DATABASE
	var column3 tabels.Jam
	db.Where("jam = ?", ed.Jam).Find(&column3)
	if ed.Jam != column3.Jam {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jam bukan waktu perkuliahan",
		})
		return
	}

	//VALIDASI: MEMASTIKAN DOSEN MENGAJAR KELAS SESUAI INPUTAN
	var column4 tabels.Dosen
	db.Where("nip = ?", c.Param("nip")).Find(&column4)
	var column5 tabels.Krs
	db.Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("dosen = ?", column4.Gelar).Find(&column5)
	if ed.Matakuliah != column5.Matakuliah {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak mengajar di kelas ini",
		})
		return
	}

	//VALIDASI: MEMASTIKAN BELUM ADA MAHASISWA YANG PRESENSI
	var column tabels.Jadwal
	db.Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("pertemuan = ?", ed.Pertemuan).Find(&column)
	if column.Presensi != 0 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tidak dapat mengubah jadwal, sudah ada mahasiswa yang presensi",
		})
		return
	}

	//JIKA LOLOS VALIDASI MAKA DATA AKAN DI UPDATE
	var tabel []tabels.Jadwal
	db.Model(&tabel).Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("pertemuan = ?", ed.Pertemuan).Update("tanggal", ed.Tanggal).Update("jam", ed.Jam)
	c.JSON(200, gin.H{
		"status":     "berhasil merubah jadwal",
		"nama":       column4.Gelar,
		"matakuliah": ed.Matakuliah,
		"kelas":      ed.Kelas,
		"tanggal":    ed.Tanggal,
		"jam":        ed.Jam,
	})
}
