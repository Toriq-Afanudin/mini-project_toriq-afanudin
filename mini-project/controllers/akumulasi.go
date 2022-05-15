package controllers

import (
	"mini_project/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Akumulasi_matakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var h models.Dosen_pengampu
	db.Where("nip = ?", c.Param("nip")).Find(&h)

	var g models.Kelas
	db.Where("dosen_pengampu_tanpa_gelar = ?", h.Tanpa_gelar).Find(&g)

	var Akumulasi []models.Akumulasi
	db.Where("matakuliah = ?", g.Matakuliah).Find(&Akumulasi)
	c.JSON(200, gin.H{
		"status": "data berhasil di peroleh",
		"data":   Akumulasi,
	})
}

func Akumulasi_mahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Akumulasi []models.Akumulasi
	db.Where("nama = ?", c.Param("nama")).Find(&Akumulasi)
	c.JSON(200, gin.H{
		"status": "data berhasil di peroleh",
		"data":   Akumulasi,
	})
}
