package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Dosen_pengampu_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Daftar_dosen []models.Dosen_pengampu
	db.Find(&Daftar_dosen)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_dosen})
}
