package controllers

import (
	"mini_project/models"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Input_presensi struct {
	Matakuliah          string `json:"matakuliah"`
	Nama_mahasiswa      string `json:"nama_mahasiswa"`
	Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
	Kehadiran           int    `json:"kehadiran"`
}

func Get_presensi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Jadwal []models.Kehadiran
	db.Find(&Jadwal)
	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

func Post_presensi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//MEMASTIKAN INPUTAN DALAM BENTUK JSON
	var Input Input_presensi
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//PROSES INPUT
	input := models.Kehadiran{
		Matakuliah:          Input.Matakuliah,
		Nama_mahasiswa:      Input.Nama_mahasiswa,
		Tanggal_perkuliahan: Input.Tanggal_perkuliahan,
		Kehadiran:           Input.Kehadiran,
	}

	//MENGAMBIL DATA MATAKULIAH DAN TANGGAL PERKULIAHAN DARI TABEL PENJADWALAN
	var Penjadwalan []models.Penjadwalan
	db.Find(&Penjadwalan)
	var jadwal models.Penjadwalan
	var matakuliah []string
	var tanggal []string
	for i := 0; i < len(Penjadwalan); i++ {
		jadwal = Penjadwalan[i]
		matakuliah = append(matakuliah, jadwal.Matakuliah)
		tanggal = append(tanggal, jadwal.Tanggal_perkuliahan)
	}
	fmt.Println(matakuliah)
	fmt.Println(tanggal)
	fmt.Println(input)

	//MEMASTIKAN MATAKULIAH DAN TANGGAL PENGAMPU YANG DI INPUT ADA DALAM TABEL PENJADWALAN
	var hitung1 int
	for i := 0; i < len(matakuliah); i++ {
		if (Input.Matakuliah == matakuliah[i]) && (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal[i]) {
			hitung1++
		}
	}
	if hitung1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Matakuliah dan tanggal yang Anda input tidak ditemukan"})

	}

	//MENGAMBIL DATA NAMA MAHASISWA
	var Nama []models.Daftar_mahasiswa
	db.Find(&Nama)
	var nama models.Daftar_mahasiswa
	var Nama_mahasiswa []string
	for i := 0; i < len(Nama); i++ {
		nama = Nama[i]
		Nama_mahasiswa = append(Nama_mahasiswa, nama.Nama)
	}

	//MEMASTIKAN NAMA MAHASISWA ADA
	var hitung2 int
	for i := 0; i < len(Nama_mahasiswa); i++ {
		if Input.Nama_mahasiswa == Nama_mahasiswa[i] {
			hitung2++
		}
	}
	if hitung2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Nama mahasiswa yang Anda input tidak ditemukan"})

	}

	//MENGAMBIL DATA NAMA MAHASISWA
	var Kehadiran []models.Kehadiran
	db.Find(&Kehadiran)
	var kehadiran models.Kehadiran
	var matkul []string
	var mahasiswa []string
	var tanggal2 []string
	for i := 0; i < len(Kehadiran); i++ {
		kehadiran = Kehadiran[i]
		matkul = append(matkul, kehadiran.Matakuliah)
		mahasiswa = append(mahasiswa, kehadiran.Nama_mahasiswa)
		tanggal2 = append(tanggal2, kehadiran.Tanggal_perkuliahan)
	}

	//MEMASTIKAN TANGGAL DAN JAM BELUM DIGUNAKAN
	var hitung3 int
	if matkul != nil {
		for i := len(matkul) - 1; i >= 0; i-- {
			if (Input.Matakuliah == matkul[i]) && (Input.Nama_mahasiswa == mahasiswa[i]) && (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal2[i]) {
				hitung3++
			}
		}
		if hitung3 != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Anda sudah melakukan presensi untuk pertemuan ini"})
		}
	}

	//JIKA SYARAT TERPENUHI MAKA DATA AKAN DI INPUTKAN
	if (hitung1 != 0) && (hitung2 != 0) && (hitung3 == 0) {
		db.Create(&input)
		c.JSON(http.StatusOK, gin.H{"Data yang telah di tambahkan": input})
	}
}
