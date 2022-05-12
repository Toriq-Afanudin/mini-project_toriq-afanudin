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
	c.JSON(http.StatusOK, gin.H{"JADWAL": Jadwal})
}

func Post_penjadwalan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//MEMASTIKAN INPUTAN DALAM BENTUK JSON
	var Input Input_penjadwalan
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROE": "INPU TIDAK DALAM BENTUK JSON"})
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
	var validasi1 int
	for i := 0; i < len(matakuliah); i++ {
		if (Input.Matakuliah == matakuliah[i]) && (Input.Dosen_pengampu_tanpa_gelar == dosen[i]) {
			validasi1++
		}
	}
	if validasi1 == 0 {
		message := "MATAKULIAH: '" + Input.Matakuliah + "' TIDAK DIAMPU OLEH DOSEN '" + Input.Dosen_pengampu_tanpa_gelar + "'"
		c.JSON(http.StatusBadRequest, gin.H{"GAGAL MENAMBAHKAN JADWAL": message})

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
	var validasi2 int
	for i := 0; i < len(jam_kuliah); i++ {
		if Input.Jam_perkuliahan == jam_kuliah[i] {
			validasi2++
		}
	}
	if validasi2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"GAGAL MENAMBAHKAN JADWAL": "JAM PERKULIAHAN: '" + Input.Jam_perkuliahan + "' YANG ANDA INPUT TIDAK DI TEMUKAN"})

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
	var validasi3 int
	if tanggal != nil {
		for i := len(tanggal) - 1; i >= 0; i-- {
			if (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal[i]) && (Input.Jam_perkuliahan == jam[i]) {
				validasi3++
			}
		}
		if validasi3 != 0 {
			message := "TANGGAL DAN JADWAL TELAH DIGUNAKAN"
			c.JSON(http.StatusBadRequest, gin.H{"GAGAL MENAMBAHKAN JADWAL": message})
		}
	}

	//MEMASTIKAN JUMLAH PERTEMUAN TIDAK LEBIH DARI 7

	//JIKA MATAKULIAH DAN DOSEN DAN JAM PERKULIAHAN ADA MAKA DATA AKAN DI INPUTKAN
	if (validasi1 != 0) && (validasi2 != 0) && (validasi3 == 0) {
		db.Create(&input)
		type berhasil struct {
			Matakuliah          string
			Dosen_pengampu      string
			Tanggal_perkuliahan string
			Jam_perkuliahan     string
			Keterangan          string
		}
		var a berhasil
		a.Matakuliah = Input.Matakuliah
		a.Dosen_pengampu = Input.Dosen_pengampu_tanpa_gelar
		a.Tanggal_perkuliahan = Input.Tanggal_perkuliahan
		a.Jam_perkuliahan = Input.Jam_perkuliahan
		a.Keterangan = "BERHASIL"
		c.JSON(http.StatusOK, gin.H{"PENJADWALAN BERHASIL DITAMBAHKAN": a})
	}
}
