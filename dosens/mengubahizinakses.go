package dosens

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
	var column tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&column)
	var column2 tanggal
	if err := c.ShouldBindJSON(&column2); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	var column3 tabels.Jadwal
	var column4 []tabels.Jadwal
	db.Where("dosen = ?", column.Id).Where("tanggal = ?", column2.Tanggal).Where("matakuliah = ?", column2.Matakuliah).Find(&column3)
	if column2.Tanggal == column3.Tanggal {
		if column3.Akses == 0 {
			var a = 1
			db.Model(&column4).Where("dosen = ?", column.Id).Where("tanggal = ?", column2.Tanggal).Update("akses", a)
			c.JSON(200, gin.H{
				"status":     "berhasil",
				"matakuliah": column3.Matakuliah,
				"tanggal":    column3.Tanggal,
				"message":    "diperbolehkan melakukan presensi",
			})
			return
		}
		if column3.Akses == 1 {
			var a = 0
			db.Model(&column4).Where("dosen = ?", column.Id).Where("tanggal = ?", column2.Tanggal).Update("akses", a)
			c.JSON(200, gin.H{
				"status":     "berhasil",
				"matakuliah": column3.Matakuliah,
				"tanggal":    column3.Tanggal,
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
