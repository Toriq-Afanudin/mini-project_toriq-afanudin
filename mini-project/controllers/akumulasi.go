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

	//MENGAMBIL DATA DARI TABEL PRESENSI KEHADIRAN
	var Hadir []models.Kehadiran
	db.Find(&Hadir)
	var j models.Kehadiran
	var Matkul []string
	var Nama_mahasiswa []string
	var Kehadiran []int
	for i := 0; i < len(Hadir); i++ {
		j = Hadir[i]
		Matkul = append(Matkul, j.Matakuliah)
		Nama_mahasiswa = append(Nama_mahasiswa, j.Nama_mahasiswa)
		Kehadiran = append(Kehadiran, j.Kehadiran)
	}

	//SORTIR UNIX NAMA MAHASISWA
	var Unix_nama_mahasiswa []string
	Unix_nama_mahasiswa = append(Unix_nama_mahasiswa, Nama_mahasiswa[0])
	if Nama_mahasiswa[1] != Nama_mahasiswa[0] {
		Unix_nama_mahasiswa = append(Unix_nama_mahasiswa, Nama_mahasiswa[1])
	}
	for i := 2; i < len(Nama_mahasiswa); i++ {
		var Jumlah int
		for m := i - 1; m >= 0; m-- {
			if Nama_mahasiswa[i] == Nama_mahasiswa[m] {
				Jumlah++
			}
		}
		if Jumlah == 0 {
			Unix_nama_mahasiswa = append(Unix_nama_mahasiswa, Nama_mahasiswa[i])
		}
	}

	//PRESENSI KELAS METODE NUMERIK
	var Metode_numerik []string
	var Kehadiran_metode_numerik []int
	for i := 0; i < len(Nama_mahasiswa); i++ {
		if Matkul[i] == "Metode Numerik" {
			Metode_numerik = append(Metode_numerik, Nama_mahasiswa[i])
			Kehadiran_metode_numerik = append(Kehadiran_metode_numerik, Kehadiran[i])
		}
	}

	//PRESENSI KELAS ANALISIS VEKTOR
	var Analisis_vektor []string
	var Kehadiran_analisis_vektor []int
	for i := 0; i < len(Nama_mahasiswa); i++ {
		if Matkul[i] == "Analisis Vektor" {
			Analisis_vektor = append(Analisis_vektor, Nama_mahasiswa[i])
			Kehadiran_analisis_vektor = append(Kehadiran_analisis_vektor, Kehadiran[i])
		}
	}

	//PRESENSI KELAS MATEMATIKA DISKRET
	var Matematika_diskret []string
	var Kehadiran_matematika_diskret []int
	for i := 0; i < len(Nama_mahasiswa); i++ {
		if Matkul[i] == "Matematika Diskret" {
			Matematika_diskret = append(Matematika_diskret, Nama_mahasiswa[i])
			Kehadiran_matematika_diskret = append(Kehadiran_matematika_diskret, Kehadiran[i])
		}
	}

	//AKUMULASI KELAS METODE NUMERIK
	type Struct_metode_numerik struct {
		Nama_mahasiswa string
		// Jumlah_pertemuan   int
		Jumlah_hadir int
		// Jumlah_tidak_hadir int
	}
	var metode_numerik Struct_metode_numerik
	var Akumulasi_metode_numerik []interface{}
	for i := 0; i < len(Unix_nama_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Metode_numerik); j++ {
			if Unix_nama_mahasiswa[i] == Metode_numerik[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_metode_numerik[j]
			}
		}
		// var Jumlah_pertemuan int
		metode_numerik.Nama_mahasiswa = Nama_mahasiswa[i]
		metode_numerik.Jumlah_hadir = Jumlah_hadir
		// analisis_vektor.Jumlah_pertemuan = Jumlah_pertemuan
		Akumulasi_metode_numerik = append(Akumulasi_metode_numerik, metode_numerik)
	}

	//AKUMULASI KELAS ANALISIS VEKTOR
	type Struct_analisis_vektor struct {
		Nama_mahasiswa string
		// Jumlah_pertemuan   int
		Jumlah_hadir int
		// Jumlah_tidak_hadir int
	}
	var analisis_vektor Struct_analisis_vektor
	var Akumulasi_analisis_vektor []interface{}
	for i := 0; i < len(Unix_nama_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Analisis_vektor); j++ {
			if Unix_nama_mahasiswa[i] == Analisis_vektor[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_analisis_vektor[j]
			}
		}
		// var Jumlah_pertemuan int
		analisis_vektor.Nama_mahasiswa = Nama_mahasiswa[i]
		analisis_vektor.Jumlah_hadir = Jumlah_hadir
		// analisis_vektor.Jumlah_pertemuan = Jumlah_pertemuan
		Akumulasi_analisis_vektor = append(Akumulasi_analisis_vektor, analisis_vektor)
	}

	//AKUMULASI KELAS MATEMATIKA DISKRET
	type Struct_matematika_diskret struct {
		Nama_mahasiswa string
		// Jumlah_pertemuan   int
		Jumlah_hadir int
		// Jumlah_tidak_hadir int
	}
	var matematika_diskret Struct_matematika_diskret
	var Akumulasi_matematika_diskret []interface{}
	for i := 0; i < len(Unix_nama_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Matematika_diskret); j++ {
			if Unix_nama_mahasiswa[i] == Matematika_diskret[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_matematika_diskret[j]
			}
		}
		// var Jumlah_pertemuan int
		matematika_diskret.Nama_mahasiswa = Nama_mahasiswa[i]
		matematika_diskret.Jumlah_hadir = Jumlah_hadir
		// analisis_vektor.Jumlah_pertemuan = Jumlah_pertemuan
		Akumulasi_matematika_diskret = append(Akumulasi_matematika_diskret, matematika_diskret)
	}

	//MENAMPILKAN DI LOCALHOST8080
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Metode Numerik": Akumulasi_metode_numerik})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Analisis Vektor": Akumulasi_analisis_vektor})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Matematika Diskret": Akumulasi_matematika_diskret})
}
