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
	var v1 bool //validasi
	if (k.Matakuliah != Input.Matakuliah) && (k.Dosen_pengampu_tanpa_gelar != Input.Dosen_pengampu_tanpa_gelar) {
		v1 = true
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
	var v2 bool
	if j.Jam_perkuliahan != Input.Jam_perkuliahan {
		v2 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jam perkuliahan tidak ditemukan",
		})
		return
	}

	//VALIDASI TANGGAL DAN JAM
	var a models.Penjadwalan
	db.Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Where("jam_perkuliahan = ?", Input.Jam_perkuliahan).Find(&a)
	var v3 bool
	if (a.Tanggal_perkuliahan == Input.Tanggal_perkuliahan) && (a.Jam_perkuliahan == Input.Jam_perkuliahan) {
		v3 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal dan jam perkuliahan sudah digunakan",
		})
		return
	}

	//MENGHITUNG JUMLAH PERTEMUAN
	var p []models.Penjadwalan
	db.Where("matakuliah = ?", Input.Matakuliah).Find(&p)
	var v4 bool
	if len(p) >= 7 {
		v4 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jumlah pertemuan sudah 7 kali",
		})
		return
	}

	//VALIDASI: MEMASTIKAN MATAKULIAH BELUM DIJADWALNKAN DALAM TANGGAL TERTENTU
	var m models.Penjadwalan
	db.Where("tanggal_perkuliahan = ?", Input.Tanggal_perkuliahan).Where("matakuliah = ?", Input.Matakuliah).Find(&m)
	var v5 bool
	if (m.Tanggal_perkuliahan == Input.Tanggal_perkuliahan) && (m.Matakuliah == Input.Matakuliah) {
		v5 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "matakuliah sudah dijadwalkan pada tanggal tersebut",
		})
		return
	}

	//VALIDASI: MEMASTIKAN BUKAN HARI LIBUR/TANGGAL MERAH
	var d models.Libur
	db.Where("tanggal = ?", Input.Tanggal_perkuliahan).Find(&d)
	var v6 bool
	fmt.Println(c)
	if d.Tanggal == Input.Tanggal_perkuliahan {
		v6 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tidak bisa melakukan perkuliahan karena hari libur " + d.Keterangan,
		})
		return
	}

	//VALIDASI: MEMASTIKAN TANGGAL MERUPAKAN MASA PERKULIAHAN
	var b models.Tanggal
	db.Where("tanggal = ?", Input.Tanggal_perkuliahan).Find(&b)
	var v7 bool
	fmt.Println(b)
	if b.Tanggal != Input.Tanggal_perkuliahan {
		v7 = true
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal diluar masa perkuliahan atau format salah",
		})
		return
	}

	//JIKA SUDAH LOLOS VALIDASI MAKA DATA AKAN DI INPUTKAN
	var trigger bool
	if !v1 && !v2 && !v3 && !v4 && !v5 && !v6 && !v7 {
		trigger = true
		db.Create(&input)
		type berhasil struct {
			Matakuliah string
			Dosen      string
			Tanggal    string
			Waktu      string
		}
		var a berhasil
		a.Matakuliah = Input.Matakuliah
		a.Dosen = k.Dosen_pengampu
		a.Tanggal = Input.Tanggal_perkuliahan
		a.Waktu = Input.Jam_perkuliahan
		c.JSON(200, gin.H{
			"status": "data berhasil di tambahkan",
			"data":   a,
		})
	}

	//TRIGGER JUMLAH PERTEMUAN
	if trigger {
		var a models.Akumulasi
		var Akumulasi []models.Akumulasi
		db.Where("matakuliah = ?", Input.Matakuliah).Find(&a)
		a.Pertemuan++
		db.Model(&Akumulasi).Where("matakuliah = ?", Input.Matakuliah).Update("pertemuan", a.Pertemuan)
		return
	}
}
