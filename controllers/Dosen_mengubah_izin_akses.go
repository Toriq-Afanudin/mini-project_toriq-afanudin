package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

type tanggal struct {
	Tanggal    string `json:"tanggal"`
	Matakuliah string `json:"matakuliah"`
}

func UpdateAkses(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var tanggal tanggal
	if err := c.ShouldBindJSON(&tanggal); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var jadwal tabels.Jadwal
	var Jadwal []tabels.Jadwal
	db.Where("dosen = ?", dosen.Id).Where("tanggal = ?", tanggal.Tanggal).Where("matakuliah = ?", tanggal.Matakuliah).Find(&jadwal)
	if tanggal.Tanggal == jadwal.Tanggal {
		if jadwal.Akses == 0 {
			var a = 1
			db.Model(&Jadwal).Where("dosen = ?", dosen.Id).Where("tanggal = ?", tanggal.Tanggal).Update("akses", a)
			c.JSON(200, gin.H{
				"status":     "berhasil",
				"matakuliah": jadwal.Matakuliah,
				"tanggal":    jadwal.Tanggal,
				"message":    "diperbolehkan melakukan presensi",
			})
			return
		}
		if jadwal.Akses == 1 {
			var a = 0
			db.Model(&Jadwal).Where("dosen = ?", dosen.Id).Where("tanggal = ?", tanggal.Tanggal).Update("akses", a)
			c.JSON(200, gin.H{
				"status":     "berhasil",
				"matakuliah": jadwal.Matakuliah,
				"tanggal":    jadwal.Tanggal,
				"message":    "TIDAK diperbolehkan melakukan presensi",
			})
			return
		}
	} else {
		c.JSON(400, gin.H{
			"status":  "gagal",
			"message": "tanggal atau matakuliah yang dimasukan salah",
		})
		return
	}
}
