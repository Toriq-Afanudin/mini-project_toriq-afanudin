package controllers

import (
	"mini_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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
	c.JSON(http.StatusOK, gin.H{"data": Jadwal})
}

func Post_penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//MEMASTIKAN INPUTAN DALAM BENTUK JSON
	var Input Input_penjadwalan
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//PROSES INPUT
	input := models.Penjadwalan{
		Matakuliah:                 Input.Matakuliah,
		Dosen_pengampu_tanpa_gelar: Input.Dosen_pengampu_tanpa_gelar,
		Tanggal_perkuliahan:        Input.Tanggal_perkuliahan,
		Jam_perkuliahan:            Input.Jam_perkuliahan,
	}

	//MENGAMBIL DATA MATAKULIAH DAN DOSEN DARI TABEL KELAS
	var Kelas []models.Kelas
	db.Find(&Kelas)
	var kelas models.Kelas
	var matakuliah []string
	var dosen []string
	for i := 0; i < len(Kelas); i++ {
		kelas = Kelas[i]
		matakuliah = append(matakuliah, kelas.Matakuliah)
		dosen = append(dosen, kelas.Dosen_pengampu_tanpa_gelar)
	}

	//MEMASTIKAN MATAKULIAH DAN DOSEN PENGAMPU YANG DI INPUT ADA DALAM TABEL KELAS
	var hitung1 int
	for i := 0; i < len(matakuliah); i++ {
		if Input.Matakuliah == matakuliah[i] {
			hitung1++
		}
	}
	if hitung1 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Matakuliah yang Anda input tidak ditemukan"})

	}
	var hitung2 int
	for i := 0; i < len(dosen); i++ {
		if Input.Dosen_pengampu_tanpa_gelar == dosen[i] {
			hitung2++
		}
	}
	if hitung2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Nama Dosen yang Anda input tidak ditemukan"})

	}

	//MENGAMBIL DATA JAM_PERKULIAHAN
	var Jam []models.Jam_perkuliahan
	db.Find(&Jam)
	var j models.Jam_perkuliahan
	var jam_kuliah []string
	for i := 0; i < len(Jam); i++ {
		j = Jam[i]
		jam_kuliah = append(jam_kuliah, j.Jam_perkuliahan)
	}

	//MEMASTIKAN JAM PERKULIAHAN ADA
	var hitung3 int
	for i := 0; i < len(jam_kuliah); i++ {
		if Input.Jam_perkuliahan == jam_kuliah[i] {
			hitung3++
		}
	}
	if hitung3 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Jam perkuliahan yang Anda input tidak ditemukan"})

	}

	//MENGAMBIL DATA TANGGAL DAN JAM DARI TABEL PENJADWALAN
	var Jadwal []models.Penjadwalan
	db.Find(&Jadwal)
	var jad models.Penjadwalan
	var tanggal []string
	var jam []string
	for i := 0; i < len(Jadwal); i++ {
		jad = Jadwal[i]
		tanggal = append(tanggal, jad.Tanggal_perkuliahan)
		jam = append(jam, jad.Jam_perkuliahan)
	}

	//MEMASTIKAN TANGGAL DAN JAM BELUM DIGUNAKAN
	var hitung4 int
	if tanggal != nil {
		for i := len(tanggal) - 1; i >= 0; i-- {
			if (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal[i]) && (Input.Jam_perkuliahan == jam[i]) {
				hitung4++
			}
		}
		if hitung4 != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Tanggal dan jam perkuliahan sudah digunakan"})
		}
	}

	//JIKA MATAKULIAH DAN DOSEN DAN JAM PERKULIAHAN ADA MAKA DATA AKAN DI INPUTKAN
	if (hitung1 != 0) && (hitung2 != 0) && (hitung3 != 0) && (hitung4 == 0) {
		db.Create(&input)
		c.JSON(http.StatusOK, gin.H{"Data yang telah di tambahkans": input})
	}
}
