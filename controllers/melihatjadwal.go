package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LihatJadwal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type jadwal struct {
		Matakuliah string
		Kelas      string
		Tanggal    string
		Jam        string
	}
	var tabel []jadwal
	db.Raw("SELECT matakuliah, kelas, tanggal, jam FROM sistem_presensi.jadwals;").Scan(&tabel)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
