package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Daftar_mahasiswa_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_mahasiswa []models.Daftar_mahasiswa
	db.Find(&Daftar_mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_mahasiswa})
}

func Dosen_pengampu_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_dosen []models.Dosen_pengampu
	db.Find(&Daftar_dosen)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_dosen})
}

func Kelas_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_kelas []models.Kelas
	db.Find(&Daftar_kelas)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kelas})
}

func Penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Penjadwalan []models.Penjadwalan
	db.Find(&Penjadwalan)
	c.JSON(http.StatusOK, gin.H{"data": Penjadwalan})
}

func Kehadiran_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_kehadiran []models.Penjadwalan
	db.Find(&Daftar_kehadiran)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kehadiran})
}
