package mahasiswas

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	jwt "github.com/appleboy/gin-jwt/v2"
)

type presensi struct {
	Matakuliah string `json:"matakuliah"`
	Tanggal    string `json:"tanggal"`
}

type akumulasi struct {
	Matakuliah string `json:"matakuliah"`
	Hadir      string `json:"hadir"`
}

func HistoriPresensi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var student []presensi
	db.Raw("SELECT matakuliah, tanggal FROM sistem_presensi.presensis WHERE nama = ?", claims["id"]).Scan(&student)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   student,
	})
}

func AkumulasiPresensi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var akumulation []akumulasi
	db.Raw("SELECT matakuliah, hadir FROM sistem_presensi.akumulasis WHERE nama = ?", claims["id"]).Scan(&akumulation)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   akumulation,
	})
}
