package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Data_input_kehadiran struct {
	Id_kehadiran int `json:"id_kehadiran"`
	Id_mahasiswa int `json:"id_mahasiswa"`
	Id_setting   int `json:"id_setting"`
	Kehadiran    int `json:"kehadiran"`
}

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
	var setting_data_input Data_input_kehadiran
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

//UBAH DATA (PUT)
func Kehadiran_ubah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Presensi models.Penjadwalan
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

	var Presensi models.Penjadwalan
	if err := db.Where("id_kehadiran = ?", c.Param("id_kehadiran")).First(&Presensi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data mahasiswa tidak di temukan"})
		return
	}

	//proses hapus data
	db.Where("id_kehadiran = ?", c.Param("id_kehadiran")).Delete(&Presensi)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
