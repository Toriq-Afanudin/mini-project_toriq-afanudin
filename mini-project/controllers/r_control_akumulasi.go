package controllers

import (
	"mini_project/models"
	"net/http"

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

	for i := 0; i < len(Hadir); i++ {
		j = Hadir[i]
		Id_mahasiswa = append(Id_mahasiswa, j.Id_mahasiswa)
		Kehadiran = append(Kehadiran, j.Kehadiran)
	}

	var Nama_mahasiswa []string
	for i := 0; i < len(Mahasiswa[0:7]); i++ {
		d = Mahasiswa[i]
		Nama_mahasiswa = append(Nama_mahasiswa, d.Nama)
	}

	var Unix_id_mahasiswa []int
	Unix_id_mahasiswa = append(Unix_id_mahasiswa, Id_mahasiswa[0])
	if Id_mahasiswa[1] != Id_mahasiswa[0] {
		Unix_id_mahasiswa = append(Unix_id_mahasiswa, Id_mahasiswa[1])
	}

	for i := 2; i < len(Hadir); i++ {
		var Jumlah int
		for m := i - 1; m >= 0; m-- {
			if Id_mahasiswa[i] == Id_mahasiswa[m] {
				Jumlah++
			}
		}
		if Jumlah == 0 {
			Unix_id_mahasiswa = append(Unix_id_mahasiswa, Id_mahasiswa[i])
		}
	}

	type Akumulasi struct {
		Nama_mahasiswa     string
		Jumlah_pertemuan   int
		Jumlah_hadir       int
		Jumlah_tidak_hadir int
	}

	var akumulasi Akumulasi
	var Akm []interface{}

	for i := 0; i < len(Unix_id_mahasiswa); i++ {
		var Jumlah_hadir int
		var Jumlah_pertemuan int
		var Jumlah_tidak_hadir int
		Jumlah_hadir = 0
		for j := 0; j < len(Id_mahasiswa); j++ {
			if Unix_id_mahasiswa[i] == Id_mahasiswa[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran[j]
				Jumlah_pertemuan++
			}
			Jumlah_tidak_hadir = Jumlah_pertemuan - Jumlah_hadir
		}
		akumulasi.Nama_mahasiswa = Nama_mahasiswa[i]
		akumulasi.Jumlah_hadir = Jumlah_hadir
		akumulasi.Jumlah_pertemuan = Jumlah_pertemuan
		akumulasi.Jumlah_tidak_hadir = Jumlah_tidak_hadir
		Akm = append(Akm, akumulasi)
	}

	// fmt.Println(Akm)
	// fmt.Println(Unix_id_mahasiswa)
	// fmt.Println(Id_mahasiswa)
	// fmt.Println(Kehadiran)
	// fmt.Println(Nama_mahasiswa)

	c.JSON(http.StatusOK, gin.H{"Akumulasi Teori Ring Kelas B": Akm})
}
