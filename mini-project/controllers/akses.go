package controllers

import (
	"mini_project/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Update_akses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type Input_akses struct {
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

	var h models.Dosen_pengampu
	db.Where("nip = ?", c.Param("nip")).Find(&h)

	var a models.Penjadwalan
	var Penjadwalan []models.Penjadwalan
	db.Where("dosen_pengampu_tanpa_gelar = ?", h.Tanpa_gelar).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Find(&a)
	var b int
	if a.Akses == 1 {
		b = 0
		db.Model(&Penjadwalan).Where("dosen_pengampu_tanpa_gelar = ?", h.Tanpa_gelar).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Update("akses", b)
		c.JSON(200, gin.H{
			"status":     "presensi tidak di izinkan",
			"matakuliah": a.Matakuliah,
			"tanggal":    input.Tanggal_perkuliahan,
		})
		return
	} else {
		b = 1
		db.Model(&Penjadwalan).Where("dosen_pengampu_tanpa_gelar = ?", h.Tanpa_gelar).Where("tanggal_perkuliahan = ?", input.Tanggal_perkuliahan).Update("akses", b)
		c.JSON(200, gin.H{
			"status":     "presensi di izinkan",
			"matakuliah": a.Matakuliah,
			"tanggal":    input.Tanggal_perkuliahan,
		})
		return
	}
}
