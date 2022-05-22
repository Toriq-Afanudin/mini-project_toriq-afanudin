package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

type presensiDosen struct {
	Nama       string
	Matakuliah string
	Tanggal    string
}

type jadwalDosen struct {
	Matakuliah string
	Kelas      string
	Tanggal    string
	Jam        string
}

type akumulasiDosen struct {
	Nama       string
	Matakuliah string
	Hadir      int
}

func MelihatPresensi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var presensi []presensiDosen
	db.Raw("SELECT nama, matakuliah, tanggal FROM sistem_presensi.presensis WHERE dosen = ?", dosen.Id).Scan(&presensi)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   presensi,
	})
}

func GetJadwalDosen(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var Jadwal []jadwalDosen
	db.Raw("SELECT matakuliah, kelas, tanggal, jam FROM sistem_presensi.jadwals WHERE dosen = ?", dosen.Id).Scan(&Jadwal)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   Jadwal,
	})
}

func GetAkumulasi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var Akumulasi []akumulasi
	db.Raw("SELECT nama, matakuliah, hadir FROM sistem_presensi.akumulasis WHERE dosen = ?", dosen.Id).Scan(&Akumulasi)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   Akumulasi,
	})
}
