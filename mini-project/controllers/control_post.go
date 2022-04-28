package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func Kehadiran_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var setting_data_input models.Kehadiran
	if err := c.ShouldBindJSON(&setting_data_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	setting := models.Kehadiran{
		Id_kehadiran: setting_data_input.Id_kehadiran,
		Id_mahasiswa: setting_data_input.Id_mahasiswa,
		Id_setting:   setting_data_input.Id_setting,
		Kehadiran:    setting_data_input.Kehadiran,
	}

	db.Create(&setting)

	c.JSON(http.StatusOK, gin.H{"data": setting})
}
