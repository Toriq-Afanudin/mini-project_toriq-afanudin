package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"

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

type jadwal struct {
	Matakuliah string `json:"matakuliah"`
	Tanggal    string `json:"tanggal"`
	Jam        string `json:"jam"`
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

func CreatePresensi(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var input presensi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input is not in json form",
		})
		return
	}
	var krs tabels.Krs
	db.Where("nama = ?", claims["id"]).Where("matakuliah = ?", input.Matakuliah).Find(&krs)
	var jadwal tabels.Jadwal
	db.Where("tanggal = ?", input.Tanggal).Where("matakuliah = ?", input.Matakuliah).Where("kelas = ?", krs.Kelas).Find(&jadwal)
	if input.Tanggal != jadwal.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal tidak ditemukan di jadwal",
		})
		return
	}
	var presensi tabels.Presensi
	db.Where("nama = ?", claims["id"]).Where("tanggal = ?", input.Tanggal).Where("matakuliah = ?", input.Matakuliah).Where("kelas = ?", krs.Kelas).Find(&presensi)
	if input.Tanggal == presensi.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda telah melakukan presensi untuk tanggal ini",
		})
		return
	}
	var dosen tabels.Dosen
	db.Where("gelar = ?", krs.Dosen).Find(&dosen)
	created := tabels.Presensi{
		Nama:       krs.Nama,
		Matakuliah: input.Matakuliah,
		Kelas:      krs.Kelas,
		Tanggal:    input.Tanggal,
		Dosen:      dosen.Id,
	}
	db.Create(&created)
	c.JSON(200, gin.H{
		"status":     "berhasil presensi",
		"matakuliah": input.Matakuliah,
		"tanggal":    input.Tanggal,
	})
	jadwal.Presensi++
	var Jadwal []tabels.Jadwal
	db.Model(&Jadwal).Where("tanggal = ?", input.Tanggal).Where("matakuliah = ?", input.Matakuliah).Where("kelas = ?", krs.Kelas).Update("presensi", jadwal.Presensi)
	var akumulasi tabels.Akumulasi
	db.Where("nama = ?", claims["id"]).Where("matakuliah = ?", input.Matakuliah).Find(&akumulasi)
	akumulasi.Hadir++
	var Akumulasi []tabels.Akumulasi
	db.Model(&Akumulasi).Where("nama = ?", claims["id"]).Where("matakuliah = ?", input.Matakuliah).Update("hadir", akumulasi.Hadir)
}

func GetJadwal(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	var Krs []tabels.Krs
	db.Raw("SELECT matakuliah, dosen FROM sistem_presensi.krs WHERE nama = ?", claims["id"]).Scan(&Krs)
	var Jadwal []interface{}
	for i := 0; i < len(Krs); i++ {
		for j := 1; j <= 5; j++ {
			var dosen tabels.Dosen
			db.Where("gelar = ?", Krs[i].Dosen).Find(&dosen)
			var jadwal jadwal
			db.Where("dosen = ?", dosen.Id).Where("matakuliah = ?", Krs[i].Matakuliah).Where("pertemuan = ?", j).Find(&jadwal)
			Jadwal = append(Jadwal, jadwal)
		}

	}
	c.JSON(200, gin.H{
		"status": "berhasil",
		"jadwal": Jadwal,
	})
}
