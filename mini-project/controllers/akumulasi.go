package controllers

import (
	"mini_project/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Get_akumulasi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Akumulasi []models.Akumulasi
	db.Find(&Akumulasi)
	c.JSON(200, gin.H{
		"status": "data berhasil di peroleh",
		"data":   Akumulasi,
	})
}
