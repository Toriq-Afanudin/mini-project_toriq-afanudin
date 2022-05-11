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
func Akumulasi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Hadir []models.Kehadiran
	var Mahasiswa []models.Daftar_mahasiswa
	db.Find(&Hadir)
	db.Find(&Mahasiswa)

	var j models.Kehadiran
	var d models.Daftar_mahasiswa
	var Id_mahasiswa []int
	var Kehadiran []int
	var Kelas []int

	//MENGAMBIL DATA DARI TABEL PRESENSI KEHADIRAN
	for i := 0; i < len(Hadir); i++ {
		j = Hadir[i]
		Id_mahasiswa = append(Id_mahasiswa, j.Id_mahasiswa)
		Kehadiran = append(Kehadiran, j.Kehadiran)
		Kelas = append(Kelas, j.Id_kelas)
	}

	//PRESENSI KELAS ILMU DAKWAH
	var Ilmu_dakwah []int
	var Kehadiran_ilmu_dakwah []int
	for i := 0; i < len(Kelas); i++ {
		if Kelas[i] == 1 {
			Ilmu_dakwah = append(Ilmu_dakwah, Id_mahasiswa[i])
			Kehadiran_ilmu_dakwah = append(Kehadiran_ilmu_dakwah, Kehadiran[i])
		}
	}

	//PRESENSI KELAS ANALISIS VEKTOR
	var Analisis_vektor []int
	var Kehadiran_analisis_vektor []int
	for i := 0; i < len(Kelas); i++ {
		if Kelas[i] == 2 {
			Analisis_vektor = append(Analisis_vektor, Id_mahasiswa[i])
			Kehadiran_analisis_vektor = append(Kehadiran_analisis_vektor, Kehadiran[i])
		}
	}

	//PRESENSI KELAS MATEMATIKA DISKRET
	var Matematika_diskret []int
	var Kehadiran_matematika_diskret []int
	for i := 0; i < len(Kelas); i++ {
		if Kelas[i] == 3 {
			Matematika_diskret = append(Matematika_diskret, Id_mahasiswa[i])
			Kehadiran_matematika_diskret = append(Kehadiran_matematika_diskret, Kehadiran[i])
		}
	}

	//MENGAMBIL DATA NAMA MAHASISWA
	var Nama_mahasiswa []string
	for i := 0; i < len(Mahasiswa[0:7]); i++ {
		d = Mahasiswa[i]
		Nama_mahasiswa = append(Nama_mahasiswa, d.Nama)
	}

	//SORTIR UNIX ID MAHASISWA
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

	//AKUMULASI KELAS ILMU DAKWAH
	type Struct_ilmu_dakwah struct {
		Nama_mahasiswa     string
		Jumlah_pertemuan   int
		Jumlah_hadir       int
		Jumlah_tidak_hadir int
	}
	var ilmu_dakwah Struct_ilmu_dakwah
	var Akumulasi_ilmu_dakwah []interface{}
	for i := 0; i < len(Unix_id_mahasiswa); i++ {
		var Jumlah_hadir int
		var Jumlah_pertemuan int
		var Jumlah_tidak_hadir int
		Jumlah_hadir = 0
		for j := 0; j < len(Ilmu_dakwah); j++ {
			if Unix_id_mahasiswa[i] == Ilmu_dakwah[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_ilmu_dakwah[j]
				Jumlah_pertemuan++
			}
			Jumlah_tidak_hadir = Jumlah_pertemuan - Jumlah_hadir
		}
		ilmu_dakwah.Nama_mahasiswa = Nama_mahasiswa[i]
		ilmu_dakwah.Jumlah_hadir = Jumlah_hadir
		ilmu_dakwah.Jumlah_pertemuan = Jumlah_pertemuan
		ilmu_dakwah.Jumlah_tidak_hadir = Jumlah_tidak_hadir
		Akumulasi_ilmu_dakwah = append(Akumulasi_ilmu_dakwah, ilmu_dakwah)
	}

	//AKUMULASI KELAS ANALISIS VEKTOR
	type Struct_analisis_vektor struct {
		Nama_mahasiswa     string
		Jumlah_pertemuan   int
		Jumlah_hadir       int
		Jumlah_tidak_hadir int
	}
	var analisis_vektor Struct_analisis_vektor
	var Akumulasi_analisis_vektor []interface{}
	for i := 0; i < len(Unix_id_mahasiswa); i++ {
		var Jumlah_hadir int
		var Jumlah_pertemuan int
		var Jumlah_tidak_hadir int
		Jumlah_hadir = 0
		for j := 0; j < len(Analisis_vektor); j++ {
			if Unix_id_mahasiswa[i] == Analisis_vektor[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_analisis_vektor[j]
				Jumlah_pertemuan++
			}
			Jumlah_tidak_hadir = Jumlah_pertemuan - Jumlah_hadir
		}
		analisis_vektor.Nama_mahasiswa = Nama_mahasiswa[i]
		analisis_vektor.Jumlah_hadir = Jumlah_hadir
		analisis_vektor.Jumlah_pertemuan = Jumlah_pertemuan
		analisis_vektor.Jumlah_tidak_hadir = Jumlah_tidak_hadir
		Akumulasi_analisis_vektor = append(Akumulasi_analisis_vektor, analisis_vektor)
	}

	//AKUMULASI KELAS MATEMATIKA DISKRET
	type Struct_matematika_diskret struct {
		Nama_mahasiswa     string
		Jumlah_pertemuan   int
		Jumlah_hadir       int
		Jumlah_tidak_hadir int
	}
	var matematika_diskret Struct_matematika_diskret
	var Akumulasi_matematika_diskret []interface{}
	for i := 0; i < len(Unix_id_mahasiswa); i++ {
		var Jumlah_hadir int
		var Jumlah_pertemuan int
		var Jumlah_tidak_hadir int
		Jumlah_hadir = 0
		for j := 0; j < len(Matematika_diskret); j++ {
			if Unix_id_mahasiswa[i] == Matematika_diskret[j] {
				Jumlah_hadir = Jumlah_hadir + Kehadiran_matematika_diskret[j]
				Jumlah_pertemuan++
			}
			Jumlah_tidak_hadir = Jumlah_pertemuan - Jumlah_hadir
		}
		matematika_diskret.Nama_mahasiswa = Nama_mahasiswa[i]
		matematika_diskret.Jumlah_hadir = Jumlah_hadir
		matematika_diskret.Jumlah_pertemuan = Jumlah_pertemuan
		matematika_diskret.Jumlah_tidak_hadir = Jumlah_tidak_hadir
		Akumulasi_matematika_diskret = append(Akumulasi_matematika_diskret, matematika_diskret)
	}

	//MENAMPILKAN DI LOKAL HOST8080
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Ilmu Dakwah": Akumulasi_ilmu_dakwah})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Analisis Vektor": Akumulasi_analisis_vektor})
	c.JSON(http.StatusOK, gin.H{"Akumulasi Kelas Matematika Diskret": Akumulasi_matematika_diskret})
}
