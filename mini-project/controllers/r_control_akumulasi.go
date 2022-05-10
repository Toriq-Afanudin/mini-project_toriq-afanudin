package controllers

import (
	"mini_project/models"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Data_input_akumulasi struct {
	Id_akumulasi       int `json:"id_akumulasi"`
	Id_mahasiswa       int `json:"id_mahasiswa"`
	Id_kelas           int `json:"id_kelas"`
	Jumlah_kelas       int `json:"jumlah_kelas"`
	Jumlah_hadir       int `json:"jumlah_hadir"`
	Jumlah_tidak_hadir int `json:"jumlah_tidak_hadir"`
}

//TAMPIL DATA (GET)
func Akumulasi_kelas_teori_ring_B(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Hadir []models.Kehadiran
	var Mahasiswa []models.Daftar_mahasiswa
	db.Find(&Hadir)
	db.Find(&Mahasiswa)

	var j models.Kehadiran
	var d models.Daftar_mahasiswa
	var Id_mahasiswa []int
	var Kehadiran []int
	var Nama_mahasiswa []string

	for i := 0; i < 21; i++ {
		j = Hadir[i]
		Id_mahasiswa = append(Id_mahasiswa, j.Id_mahasiswa)
		Kehadiran = append(Kehadiran, j.Kehadiran)
	}

	for i := 0; i < 7; i++ {
		d = Mahasiswa[i]
		Nama_mahasiswa = append(Nama_mahasiswa, d.Nama)
	}

	fmt.Println(Id_mahasiswa)
	fmt.Println(Kehadiran)
	fmt.Println(Nama_mahasiswa)

	var Unix_id_mahasiswa []int
	for i := 1; 1 <= 7; i++ {
		Unix_id_mahasiswa = append(Unix_id_mahasiswa, i)
	}

	var Jumlah int
	for i := 0; i < 21; i++ {
		if Unix_id_mahasiswa[0] == Id_mahasiswa[i] {
			Jumlah++
		}
	}
	fmt.Println(Jumlah)

	c.JSON(http.StatusOK, gin.H{"data": "Akumulasi Berhasil Dilakukan"})
}
