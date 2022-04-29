package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Kelas_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_kelas []models.Kelas
	db.Find(&Daftar_kelas)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_kelas})
}
