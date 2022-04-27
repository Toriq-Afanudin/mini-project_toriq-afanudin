package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	Id_mahasiswa int    `json:"id_mahasiswa"`
	Nama         string `json:"nama"`
	Nim          string `json:"nim"`
}

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

//TAMBAH DATA
