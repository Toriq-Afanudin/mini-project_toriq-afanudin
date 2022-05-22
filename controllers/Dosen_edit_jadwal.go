package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

type edit struct {
	Matakuliah string `json:"matakuliah"`
	Kelas      string `json:"kelas"`
	Pertemuan  int    `json:"pertemuan"`
	Tanggal    string `json:"tanggal"`
	Jam        string `json:"jam"`
}

func EditJadwal(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	//VALIDASI JSON
	var edit edit
	if err := c.ShouldBindJSON(&edit); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	//VALIDASI TANGGAL PERKULIAHAN
	var tanggal tabels.Tanggal
	db.Where("tanggal = ?", edit.Tanggal).Find(&tanggal)
	if edit.Tanggal != tanggal.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal yang di input bukan masa perkuliahan",
		})
		return
	}
	//VALIDASI:	MEMASTIKAN JAM PERKULIAHAN SESUAI DI DATABASE
	var jam tabels.Jam
	db.Where("jam = ?", edit.Jam).Find(&jam)
	if edit.Jam != jam.Jam {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jam bukan waktu perkuliahan",
		})
		return
	}
	//VALIDASI: MEMASTIKAN DOSEN MENGAJAR KELAS SESUAI INPUTAN
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var krs tabels.Krs
	db.Where("matakuliah = ?", edit.Matakuliah).Where("kelas = ?", edit.Kelas).Where("dosen = ?", dosen.Gelar).Find(&krs)
	if edit.Matakuliah != krs.Matakuliah {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak mengajar di kelas ini",
		})
		return
	}
	//VALIDASI: MEMASTIKAN BELUM ADA MAHASISWA YANG PRESENSI
	var column tabels.Jadwal
	db.Where("matakuliah = ?", edit.Matakuliah).Where("kelas = ?", edit.Kelas).Where("pertemuan = ?", edit.Pertemuan).Find(&column)
	if column.Presensi != 0 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tidak dapat mengubah jadwal, sudah ada mahasiswa yang presensi",
		})
		return
	}
	//JIKA LOLOS VALIDASI MAKA DATA AKAN DI UPDATE
	var tabel []tabels.Jadwal
	db.Model(&tabel).Where("matakuliah = ?", edit.Matakuliah).Where("kelas = ?", edit.Kelas).Where("pertemuan = ?", edit.Pertemuan).Update("tanggal", edit.Tanggal).Update("jam", edit.Jam)
	c.JSON(200, gin.H{
		"status":     "berhasil merubah jadwal",
		"nama":       dosen.Nama,
		"matakuliah": edit.Matakuliah,
		"kelas":      edit.Kelas,
		"tanggal":    edit.Tanggal,
		"jam":        edit.Jam,
	})
}
