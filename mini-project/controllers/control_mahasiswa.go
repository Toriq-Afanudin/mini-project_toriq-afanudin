package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Daftar_mahasiswa_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_mahasiswa []models.Daftar_mahasiswa
	db.Find(&Daftar_mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_mahasiswa})
}
