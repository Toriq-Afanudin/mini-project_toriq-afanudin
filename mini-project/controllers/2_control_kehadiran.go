package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Data_input_kehadiran struct {
	Id_penjadwalan int `json:"id_setting"`
	Id_mahasiswa   int `json:"id_mahasiswa"`
	Kehadiran      int `json:"kehadiran"`
}

//TAMPIL DATA
func Kehadiran_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_kehadiran []models.Kehadiran
	db.Find(&Daftar_kehadiran)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kehadiran})
}

func Kehadiran_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var setting_data_input Data_input_kehadiran
	if err := c.ShouldBindJSON(&setting_data_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	setting := models.Kehadiran{
		Id_penjadwalan: setting_data_input.Id_penjadwalan,
		Id_mahasiswa:   setting_data_input.Id_mahasiswa,
		Kehadiran:      setting_data_input.Kehadiran,
	}

	db.Create(&setting)

	c.JSON(http.StatusOK, gin.H{"data": setting})
}

//UBAH DATA (PUT)
func Kehadiran_ubah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Presensi models.Kehadiran
	if err := db.Where("id_kehadiran = ?", c.Param("id_kehadiran")).First(&Presensi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi input/masukan
	var dataInput Data_input_kehadiran
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&Presensi).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": Presensi})
}

//HAPUS DATA (DELETE)
func Kehadiran_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Presensi models.Kehadiran
	if err := db.Where("id_kehadiran = ?", c.Param("id_kehadiran")).First(&Presensi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data mahasiswa tidak di temukan"})
		return
	}

	//proses hapus data
	db.Where("id_kehadiran = ?", c.Param("id_kehadiran")).Delete(&Presensi)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
