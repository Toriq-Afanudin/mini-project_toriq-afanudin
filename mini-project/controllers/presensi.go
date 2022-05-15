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
}

func Presensi_mahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Jadwal []models.Kehadiran
	db.Where("nama_mahasiswa = ?", c.Param("nama_mahasiswa")).Find(&Jadwal)
	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

func Presensi_matakuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var h models.Dosen_pengampu
	db.Where("nip = ?", c.Param("nip")).Find(&h)

	var g models.Kelas
	db.Where("dosen_pengampu_tanpa_gelar = ?", h.Tanpa_gelar).Find(&g)

	var Kehadiran []models.Kehadiran
	db.Where("matakuliah = ?", g.Matakuliah).Where("tanggal_perkuliahan = ?", c.Param("tanggal_perkuliahan")).Find(&Kehadiran)
	c.JSON(http.StatusOK, gin.H{"data": Kehadiran})
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
	}

	//VALIDASI: MEMASTIKAN DATA 'MATAKULIAH' DAN 'TANGGAL PERKULIAHAN' ADA DI DATABASE
	var m models.Penjadwalan
	db.Where("matakuliah = ?", Input.Matakuliah).Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Find(&m)
	var v1 bool
	if (Input.Matakuliah != m.Matakuliah) && (Input.Tanggal_perkuliahan != m.Tanggal_perkuliahan) {
		v1 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "matakuliah `" + Input.Matakuliah + "` belum dijadwalkan pada tanggal `" + input.Tanggal_perkuliahan + "`",
		})
		return
	}

	//VALIDASI: MEMASTIKAN NAMA MAHASISWA ADA DI DATABASE
	var n models.Daftar_mahasiswa
	db.Where("nama = ?", Input.Nama_mahasiswa).Find(&n)
	var v2 bool
	if Input.Nama_mahasiswa != n.Nama {
		v2 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "nama mahasiswa `" + Input.Nama_mahasiswa + "` tidak ditemukan",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MAHASISWA BELUM MELAKUKAN PRESENSI
	var s models.Kehadiran
	db.Where("matakuliah = ?", Input.Matakuliah).Where("nama_mahasiswa = ?", Input.Nama_mahasiswa).Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Find(&s)
	var v3 bool
	if (s.Matakuliah == Input.Matakuliah) && (s.Nama_mahasiswa == Input.Nama_mahasiswa) && (s.Tanggal_perkuliahan == Input.Tanggal_perkuliahan) {
		v3 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda sudah melakukan presensi",
		})
		return
	}

	//MEMASTIKAN PRESENSI DI IZINKAN OLEH DOSEN
	var e models.Penjadwalan
	var v4 bool
	db.Where("matakuliah = ?", Input.Matakuliah).Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Find(&e)
	if e.Akses != 1 {
		v4 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "presensi tidak di izinkan oleh dosen pengampu",
		})
		return
	}

	//JIKA SEMUA VALIDASI LOLOS MAKA DATA AKAN DI INPUTKAN
	var trigger bool
	if !v1 && !v2 && !v3 && !v4 {
		trigger = true
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
		t.Kehadiran = "Hadir"
		c.JSON(200, gin.H{
			"status": "data berhasil di tambahkan",
			"data":   t,
		})
	}

	//TRIGGER JUMLAH HADIR
	if trigger {
		var a models.Akumulasi
		var Akumulasi []models.Akumulasi
		db.Where("matakuliah = ?", Input.Matakuliah).Where("nama = ?", Input.Nama_mahasiswa).Find(&a)
		a.Hadir++
		a.Tidak = a.Pertemuan - a.Hadir
		db.Model(&Akumulasi).Where("matakuliah = ?", Input.Matakuliah).Where("nama = ?", Input.Nama_mahasiswa).Update("hadir", a.Hadir).Update("tidak", a.Tidak)
		return
	}
}
