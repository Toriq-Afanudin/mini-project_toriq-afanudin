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
func Penjadwalan_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Jadwal []models.Penjadwalan
	db.Find(&Jadwal)
	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

//TAMBAH DATA (POST)
func Penjadwalan_tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukan
	var Input Data_input
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	setting := models.Penjadwalan{
		Id_penjadwalan:      Input.Id_penjadwalan,
		Id_kelas:            Input.Id_kelas,
		Tanggal_perkuliahan: Input.Tanggal_perkuliahan,
		Jam_perkuliahan:     Input.Jam_perkuliahan,
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
	db.Where("id_penjadwalan = ?", c.Param("id_penjadwalan")).Model(&Jadwal).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

//HAPUS DATA (DELETE)
func Penjadwalan_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Jadwal models.Penjadwalan
	if err := db.Where("id_penjadwalan = ?", c.Param("id_penjadwalan")).First(&Jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data mahasiswa tidak di temukan"})
		return
	}

	//proses hapus data
	db.Where("id_penjadwalan = ?", c.Param("id_penjadwalan")).Delete(&Jadwal)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
