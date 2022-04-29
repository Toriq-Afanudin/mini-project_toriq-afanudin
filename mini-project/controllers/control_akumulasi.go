package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Data_input_akumulasi struct {
	Jumlah_kelas       int `json:"jumlah_kelas"`
	Jumlah_hadir       int `json:"jumlah_hadir"`
	Jumlah_tidak_hadir int `json:"jumlah_tidak_hadir"`
}

//TAMPIL DATA (GET)
func Akumulasi_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Akumulasi []models.Akumulasi_per_kelas
	db.Find(&Akumulasi)
	c.JSON(http.StatusOK, gin.H{"data": Akumulasi})
}

//UBAH DATA (PUT)
func Akumulasi_ubah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Akumulasi models.Akumulasi_per_kelas
	if err := db.Where("id_akumulasi = ?", c.Param("id_akumulasi")).First(&Akumulasi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validasi input/masukan
	var dataInput Data_input_akumulasi
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&Akumulasi).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": Akumulasi})
}
