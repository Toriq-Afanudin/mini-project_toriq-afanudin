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
	var penjadwalan []models.Penjadwalan
	db.Find(&penjadwalan)
	var penjadwalan_struct models.Penjadwalan
	var pertemuan_metode_numerik int
	var pertemuan_analisis_vektor int
	var pertemuan_matematika_diskret int
	for i := 0; i < len(penjadwalan); i++ {
		penjadwalan_struct = penjadwalan[i]
		if penjadwalan_struct.Matakuliah == "Metode Numerik" {
			pertemuan_metode_numerik++
		}
		if penjadwalan_struct.Matakuliah == "Analisis Vektor" {
			pertemuan_analisis_vektor++
		}
		if penjadwalan_struct.Matakuliah == "Matematika Diskret" {
			pertemuan_matematika_diskret++
		}
	}

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

	//AMBIL DATA NAMA MAHASISWA
	var Daftar_mahasiswa []models.Daftar_mahasiswa
	db.Find(&Daftar_mahasiswa)
	var daftar_mahasiswa []string
	var daftar_mahasiswa_struct models.Daftar_mahasiswa
	for i := 0; i < len(Daftar_mahasiswa); i++ {
		daftar_mahasiswa_struct = Daftar_mahasiswa[i]
		daftar_mahasiswa = append(daftar_mahasiswa, daftar_mahasiswa_struct.Nama)
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
		Nama_mahasiswa   string
		Jumlah_pertemuan int
		Jumlah_hadir     int
	}
	var metode_numerik Struct_metode_numerik
	var Akumulasi_metode_numerik []interface{}
	for i := 0; i < len(daftar_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Metode_numerik); j++ {
			if daftar_mahasiswa[i] == Metode_numerik[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_metode_numerik[j]
			}
		}
		metode_numerik.Nama_mahasiswa = daftar_mahasiswa[i]
		metode_numerik.Jumlah_hadir = Jumlah_hadir
		metode_numerik.Jumlah_pertemuan = pertemuan_metode_numerik
		Akumulasi_metode_numerik = append(Akumulasi_metode_numerik, metode_numerik)
	}

	//AKUMULASI KELAS ANALISIS VEKTOR
	type Struct_analisis_vektor struct {
		Nama_mahasiswa   string
		Jumlah_pertemuan int
		Jumlah_hadir     int
	}
	var analisis_vektor Struct_analisis_vektor
	var Akumulasi_analisis_vektor []interface{}
	for i := 0; i < len(daftar_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Analisis_vektor); j++ {
			if daftar_mahasiswa[i] == Analisis_vektor[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_analisis_vektor[j]
			}
		}
		analisis_vektor.Nama_mahasiswa = daftar_mahasiswa[i]
		analisis_vektor.Jumlah_hadir = Jumlah_hadir
		analisis_vektor.Jumlah_pertemuan = pertemuan_analisis_vektor
		Akumulasi_analisis_vektor = append(Akumulasi_analisis_vektor, analisis_vektor)
	}

	//AKUMULASI KELAS MATEMATIKA DISKRET
	type Struct_matematika_diskret struct {
		Nama_mahasiswa   string
		Jumlah_pertemuan int
		Jumlah_hadir     int
	}
	var matematika_diskret Struct_matematika_diskret
	var Akumulasi_matematika_diskret []interface{}
	for i := 0; i < len(daftar_mahasiswa); i++ {
		var Jumlah_hadir int
		for j := 0; j < len(Matematika_diskret); j++ {
			if daftar_mahasiswa[i] == Matematika_diskret[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_matematika_diskret[j]
			}
		}
		matematika_diskret.Nama_mahasiswa = daftar_mahasiswa[i]
		matematika_diskret.Jumlah_hadir = Jumlah_hadir
		matematika_diskret.Jumlah_pertemuan = pertemuan_matematika_diskret
		Akumulasi_matematika_diskret = append(Akumulasi_matematika_diskret, matematika_diskret)
	}

	//MENAMPILKAN DI LOCALHOST8080
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Metode Numerik": Akumulasi_metode_numerik})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Analisis Vektor": Akumulasi_analisis_vektor})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Matematika Diskret": Akumulasi_matematika_diskret})
}
