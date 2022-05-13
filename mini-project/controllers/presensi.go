package controllers

import (
	"mini_project/models"
	"net/http"

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
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "data tidak dalam bentuk json",
		})
		return
	}

	//PROSES INPUT
	input := models.Kehadiran{
		Matakuliah:          Input.Matakuliah,
		Nama_mahasiswa:      Input.Nama_mahasiswa,
		Tanggal_perkuliahan: Input.Tanggal_perkuliahan,
		Kehadiran:           Input.Kehadiran,
	}

	//VALIDASI: MEMASTIKAN DATA 'MATAKULIAH' DAN 'TANGGAL PERKULIAHAN' ADA DI DATABASE
	var m models.Penjadwalan
	db.Where("matakuliah = ?", Input.Matakuliah).Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Find(&m)
	var v1 int
	if (Input.Matakuliah != m.Matakuliah) && (Input.Tanggal_perkuliahan+"T00:00:00+07:00" != m.Tanggal_perkuliahan) {
		v1 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "MATAKULIAH DAN TANGGAL TIDAK DITEMUKAN",
		})
		return
	}

	//VALIDASI: MEMASTIKAN NAMA MAHASISWA ADA DI DATABASE
	var n models.Daftar_mahasiswa
	db.Where("nama = ?", Input.Nama_mahasiswa).Find(&n)
	var v2 int
	if Input.Nama_mahasiswa != n.Nama {
		v2 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "NAMA MAHASISWA TIDAK DITEMUKAN",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA BELUM MELAKUKAN PRESENSI
	var s models.Kehadiran
	db.Where("matakuliah = ?", Input.Matakuliah).Where("nama_mahasiswa = ?", Input.Nama_mahasiswa).Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Find(&s)
	var v3 int
	if (s.Matakuliah == Input.Matakuliah) && (s.Nama_mahasiswa == Input.Nama_mahasiswa) && (s.Tanggal_perkuliahan == Input.Tanggal_perkuliahan+"T00:00:00+07:00") {
		v3 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "ANDA SUDAH MELAKUKAN PRESENSI",
		})
		return
	}

	//VALIDASI: MEMASTIKAN INPUT KEHADIRAN = 1
	var v4 int
	if Input.Kehadiran != 1 {
		v4 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "INPUT KEHADIRAN DENGAN ANGKA 1",
		})
		return
	}

	//JIKA SEMUA VALIDASI LOLOS MAKA DATA AKAN DI INPUTKAN
	if (v1 != 1) && (v2 != 1) && (v3 != 1) && (v4 != 1) {
		db.Create(&input)
		type Tampilkan struct {
			Matakuliah string
			Nama       string
			Tanggal    string
			Kehadiran  string
		}
		var t Tampilkan
		t.Matakuliah = Input.Matakuliah
		t.Nama = Input.Nama_mahasiswa
		t.Tanggal = Input.Tanggal_perkuliahan
		t.Kehadiran = "HADIR"
		c.JSON(200, gin.H{
			"status": "data berhasil di tambahkan",
			"data":   t,
		})
		return
	}
}
