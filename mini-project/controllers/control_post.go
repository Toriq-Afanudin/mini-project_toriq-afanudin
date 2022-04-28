package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMBAH DATA
func Setting_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var setting_data_input models.Setting_presensi_oleh_dosen
	if err := c.ShouldBindJSON(&setting_data_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	setting := models.Setting_presensi_oleh_dosen{
		Id_setting:          setting_data_input.Id_setting,
		Id_kelas:            setting_data_input.Id_kelas,
		Tanggal_perkuliahan: setting_data_input.Tanggal_perkuliahan,
		Jam_perkuliahan:     setting_data_input.Jam_perkuliahan,
	}

	db.Create(&setting)

	c.JSON(http.StatusOK, gin.H{"data": setting})
}
