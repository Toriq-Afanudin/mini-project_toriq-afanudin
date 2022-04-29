package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Penjadwalan []models.Penjadwalan
	db.Find(&Penjadwalan)
	c.JSON(http.StatusOK, gin.H{"data": Penjadwalan})
}

//TAMBAH DATA
func Penjadwalan_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var Penjadwalan models.Penjadwalan
	if err := c.ShouldBindJSON(&Penjadwalan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	setting := models.Penjadwalan{
		Id_penjadwalan:      Penjadwalan.Id_penjadwalan,
		Id_kelas:            Penjadwalan.Id_kelas,
		Tanggal_perkuliahan: Penjadwalan.Tanggal_perkuliahan,
		Jam_perkuliahan:     Penjadwalan.Jam_perkuliahan,
	}

	db.Create(&setting)

	c.JSON(http.StatusOK, gin.H{"data": setting})
}
