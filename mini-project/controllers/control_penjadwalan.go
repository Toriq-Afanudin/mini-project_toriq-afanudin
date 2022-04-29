package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Data_input struct {
	Id_penjadwalan      int    `json:"id_penjadwalan"`
	Id_kelas            int    `json:"id_kelas"`
	Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
	Jam_perkuliahan     string `json:"jam_perkuliahan"`
}

//TAMPIL DATA (GET)
func Penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Penjadwalan []models.Penjadwalan
	db.Find(&Penjadwalan)
	c.JSON(http.StatusOK, gin.H{"data": Penjadwalan})
}

//TAMBAH DATA (POST)
func Penjadwalan_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var Penjadwalan Data_input
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

//UBAH DATA (PUT)
func Penjadwalan_ubah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Jadwal models.Penjadwalan
	if err := db.Where("id_penjadwalan = ?", c.Param("id_penjadwalan")).First(&Jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi input/masukan
	var dataInput Data_input
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&Jadwal).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

//DATA HAPUS (DELETE)
func Penjadwalan_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Jadwal Data_input
	if err := db.Where("id_penjadwalan = ?", c.Param("id_penjadwalan")).First(&Jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data mahasiswa tidak di temukan"})
		return
	}

	//proses ubah data
	db.Delete(&Jadwal)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
