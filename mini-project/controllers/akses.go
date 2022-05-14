package controllers

import (
	"mini_project/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Update_akses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type Input_akses struct {
		Matakuliah          string `json:"matakuliah"`
		Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
	}

	var input Input_akses

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	var a models.Penjadwalan
	var Penjadwalan []models.Penjadwalan
	db.Where("matakuliah = ?", input.Matakuliah).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Find(&a)
	var b int
	if a.Akses == 1 {
		b = 0
		db.Model(&Penjadwalan).Where("matakuliah = ?", input.Matakuliah).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Update("akses", b)
		c.JSON(200, gin.H{
			"status":     "presensi tidak di izinkan",
			"matakuliah": input.Matakuliah,
			"tanggal":    input.Tanggal_perkuliahan,
		})
		return
	} else {
		b = 1
		db.Model(&Penjadwalan).Where("matakuliah = ?", input.Matakuliah).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Update("akses", b)
		c.JSON(200, gin.H{
			"status":     "presensi di izinkan",
			"matakuliah": input.Matakuliah,
			"tanggal":    input.Tanggal_perkuliahan,
		})
		return
	}
}
