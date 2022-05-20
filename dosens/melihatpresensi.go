package dosens

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

type presensi struct {
	Nama       string
	Matakuliah string
	Tanggal    string
}

type jadwal struct {
	Matakuliah string
	Kelas      string
	Tanggal    string
	Jam        string
}

type akumulasi struct {
	Nama       string
	Matakuliah string
	Hadir      int
}

func MelihatPresensi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var dosen tabels.Dosen
	db.Where("nama = ?", claims["id"]).Find(&dosen)
	var presensi []presensi
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
	var Jadwal []jadwal
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
