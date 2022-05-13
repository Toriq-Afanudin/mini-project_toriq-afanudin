package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Input_akumulasi struct {
	Nama_mahasiswa   string `json:"nama_mahasiswa"`
	Jumlah_pertemuan int    `json:"jumlah_pertemuan"`
	Jumlah_hadir     int    `json:"jumlah_hadir"`
}

func Get_akumulasi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	//MENGHITUNG JUMLAH PERTEMUAN
	var matkul []string
	matkul = append(matkul, "Metode Numerik", "Analisis Vektor", "Matematika Diskret")
	var mahasiswa []string
	mahasiswa = append(mahasiswa, "Brilian Cahya Puspitaningrum", "Najla Zaul Fadaukas", "Laela Durrotun Nafisah", "Anindya Putri", "Azizah Nurul Wahidah", "Anggita Nur Fadhilah", "Ika Winda Kusumasari")
	type Record struct {
		Nama             string
		Matakuliah       string
		Jumlah_pertemuan int
		Jumlah_hadir     int
	}
	var r Record
	var record []interface{}
	for i := 0; i < len(matkul); i++ {
		for j := 0; j < len(mahasiswa); j++ {
			var p []models.Penjadwalan
			db.Where("matakuliah = ?", matkul[i]).Find(&p)
			var q []models.Kehadiran
			db.Where("matakuliah = ?", matkul[i]).Where("nama_mahasiswa = ?", mahasiswa[j]).Find(&q)
			r.Nama = mahasiswa[j]
			r.Matakuliah = matkul[i]
			r.Jumlah_pertemuan = len(p)
			r.Jumlah_hadir = len(q)
			record = append(record, r)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}
