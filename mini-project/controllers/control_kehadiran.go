package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Kehadiran_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_kehadiran []models.Penjadwalan
	db.Find(&Daftar_kehadiran)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kehadiran})
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
