package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Daftar []interface{}

//TAMPIL DATA
func Dosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Daftar_dosen []models.Dosen_pengampu
	db.Find(&Daftar_dosen)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_dosen})
}

//TAMPIL DATA
func Kelas(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Daftar_kelas []models.Kelas
	db.Find(&Daftar_kelas)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kelas})
}

//TAMPIL DATA
func Mahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_mahasiswa []models.Daftar_mahasiswa
	db.Find(&Daftar_mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_mahasiswa})
}

//TAMPIL DATA
func Jam_perkuliahan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Jam []models.Jam_perkuliahan
	db.Find(&Jam)
	c.JSON(http.StatusOK, gin.H{"data": Jam})
}
