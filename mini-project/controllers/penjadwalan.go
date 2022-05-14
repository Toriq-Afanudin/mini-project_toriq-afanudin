package controllers

import (
	"fmt"
	"mini_project/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//STRUCT INI DIGUNAKAN UNTUK POST DAN PUT
type Input_penjadwalan struct {
	Matakuliah                 string `json:"matakuliah"`
	Dosen_pengampu_tanpa_gelar string `json:"dosen_pengampu_tanpa_gelar"`
	Tanggal_perkuliahan        string `json:"tanggal_perkuliahan"`
	Jam_perkuliahan            string `json:"jam_perkuliahan"`
}

func Get_penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Jadwal []models.Penjadwalan
	db.Find(&Jadwal)
	c.JSON(200, gin.H{
		"status": "data berhasil di peroleh",
		"data":   Jadwal,
	})
}

func Post_penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//MEMASTIKAN INPUTAN DALAM BENTUK JSON
	var Input Input_penjadwalan
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	//PROSES INPUT
	input := models.Penjadwalan{
		Matakuliah:                 Input.Matakuliah,
		Dosen_pengampu_tanpa_gelar: Input.Dosen_pengampu_tanpa_gelar,
		Tanggal_perkuliahan:        Input.Tanggal_perkuliahan,
		Jam_perkuliahan:            Input.Jam_perkuliahan,
	}

	//VALIDASI KELAS (MATAKULIAH DAN DOSEN)
	var k models.Kelas
	db.Where("matakuliah = ?", Input.Matakuliah).Where("dosen_pengampu_tanpa_gelar = ?", Input.Dosen_pengampu_tanpa_gelar).Find(&k)
	var v1 int //validasi
	if (k.Matakuliah != Input.Matakuliah) && (k.Dosen_pengampu_tanpa_gelar != Input.Dosen_pengampu_tanpa_gelar) {
		v1 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "kelas tidak ditemukan",
			"saran":   "cek kembali matakuliah atau nama dosen",
		})
		return
	}

	//VALIDASI JAM
	var j models.Jam_perkuliahan
	db.Where("jam_perkuliahan = ?", Input.Jam_perkuliahan).Find(&j)
	var v2 int
	if j.Jam_perkuliahan != Input.Jam_perkuliahan {
		v2 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jam perkuliahan tidak ditemukan",
		})
		return
	}

	//VALIDASI TANGGAL DAN JAM
	var a models.Penjadwalan
	db.Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Where("jam_perkuliahan = ?", Input.Jam_perkuliahan).Find(&a)
	var v3 int
	if (a.Tanggal_perkuliahan == Input.Tanggal_perkuliahan+"T00:00:00+07:00") && (a.Jam_perkuliahan == Input.Jam_perkuliahan) {
		v3 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal dan jam perkuliahan sudah digunakan",
		})
		return
	}

	//MENGHITUNG JUMLAH PERTEMUAN
	var p []models.Penjadwalan
	db.Where("matakuliah = ?", Input.Matakuliah).Find(&p)
	var v4 int
	if len(p) >= 7 {
		v4 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jumlah pertemuan sudah 7 kali",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MATAKULIAH BELUM DIJADWALNKAN DALAM TANGGAL TERTENTU
	var m models.Penjadwalan
	db.Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Where("matakuliah = ?", Input.Matakuliah).Find(&m)
	var v5 int
	if (m.Tanggal_perkuliahan == Input.Tanggal_perkuliahan+"T00:00:00+07:00") && (m.Matakuliah == Input.Matakuliah) {
		v5 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "matakuliah sudah dijadwalkan pada tanggal tersebut",
		})
		return
	}

	//VALIDASI PENULISAN TANGGAL 1
	var d models.Libur
	db.Where("tanggal = ?", Input.Tanggal_perkuliahan).Find(&d)
	var v6 int
	fmt.Println(c)
	if d.Tanggal == Input.Tanggal_perkuliahan {
		v6 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tidak bisa melakukan perkuliahan karena hari libur " + d.Keterangan,
		})
		return
	}

	//VALIDASI PENULISAN TANGGAL 2
	var b models.Tanggal
	db.Where("tanggal = ?", Input.Tanggal_perkuliahan).Find(&b)
	var v7 int
	fmt.Println(b)
	if b.Tanggal != Input.Tanggal_perkuliahan {
		v7 = 1
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal diluar masa perkuliahan",
		})
		return
	}

	//JIKA SUDAH LOLOS VALIDASI MAKA DATA AKAN DI INPUTKAN
	var trigger int
	if (v1 != 1) && (v2 != 1) && (v3 != 1) && (v4 != 1) && (v5 != 1) && (v6 != 1) && (v7 != 1) {
		trigger = 1
		db.Create(&input)
		type berhasil struct {
			Matakuliah string
			Dosen      string
			Tanggal    string
			Waktu      string
		}
		var a berhasil
		a.Matakuliah = Input.Matakuliah
		a.Dosen = Input.Dosen_pengampu_tanpa_gelar
		a.Tanggal = Input.Tanggal_perkuliahan[0:10]
		a.Waktu = Input.Jam_perkuliahan
		c.JSON(200, gin.H{
			"status": "data berhasil di tambahkan",
			"data":   a,
		})
	}

	//TRIGGER JUMLAH PERTEMUAN
	if trigger == 1 {
		var a models.Akumulasi
		var Akumulasi []models.Akumulasi
		db.Where("matakuliah = ?", Input.Matakuliah).Find(&a)
		a.Pertemuan++
		db.Model(&Akumulasi).Where("matakuliah = ?", Input.Matakuliah).Update("pertemuan", a.Pertemuan)
		return
	}
}
