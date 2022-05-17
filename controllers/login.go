package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type login struct {
		Nama     string `json:"nama"`
		Nip      string `json:"nip"`
		Password string `json:"password"`
	}

	var log login
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	l := tabels.Dosen{
		Nama:     log.Nama,
		Nip:      log.Nip,
		Password: log.Password,
	}

	var column tabels.Dosen
	db.Where("nama = ?", log.Nama).Where("nip = ?", log.Nip).Where("password = ?", log.Password).Find(&column)
	var column2 tabels.Login
	db.Where("nama = ?", log.Nama).Where("nomer = ?", log.Nip).Find(&column2)

	if (l.Nama == column2.Nama) && (l.Nip == column2.Nomer) {
		c.JSON(200, gin.H{
			"data":   l.Nama,
			"status": "anda sudah login",
		})
		return
	}

	if (l.Nama == column.Nama) && (l.Nip == column.Nip) && (l.Password == column.Password) {
		dl := tabels.Login{
			Nama:  l.Nama,
			Nomer: l.Nip,
		}
		db.Create(&dl)
		c.JSON(200, gin.H{
			"data":   dl.Nama,
			"status": "berhasil login",
		})
		return
	} else {
		c.JSON(400, gin.H{
			"status": "login gagal",
		})
		return
	}
}
