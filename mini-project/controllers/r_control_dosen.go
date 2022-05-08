package controllers

import (
	"mini_project/models"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TAMPIL DATA
func Dosen_pengampu_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Daftar_dosen []models.Dosen_pengampu
	db.Find(&Daftar_dosen)
	for i := 0; i < 15; i++ {
		fmt.Println(Daftar_dosen[i])
	}
	c.JSON(http.StatusOK, gin.H{"data": Daftar_dosen})
}
