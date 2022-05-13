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
	var akses []int
	for i := 0; i < len(Penjadwalan); i++ {
		jadwal = Penjadwalan[i]
		matakuliah = append(matakuliah, jadwal.Matakuliah)
		tanggal = append(tanggal, jadwal.Tanggal_perkuliahan)
		akses = append(akses, jadwal.Akses)
	}

	//MEMASTIKAN MATAKULIAH DAN TANGGAL PENGAMPU YANG DI INPUT ADA DALAM TABEL PENJADWALAN
	var validasi1 int
	var validasi int
	for i := 0; i < len(matakuliah); i++ {
		if (Input.Matakuliah == matakuliah[i]) && (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal[i]) {
			validasi1++
		}
		if akses[i] == 1 {
			validasi++
		}
	}
	if validasi1 == 0 {
		message := "MATAKULIAH: '" + Input.Matakuliah + "' ATAU TANGGAL: '" + Input.Tanggal_perkuliahan + "' YANG ANDA INPUTKAN TIDAK DI TEMUKAN"
		c.JSON(http.StatusBadRequest, gin.H{"PRESENSI GAGAL": message})
	}
	if validasi != 1 {
		message := "PRESENSI TIDAK DI IZINKAN ATAU SEDANG DI TUTUP OLEH DOSEN PENGAMPU"
		c.JSON(http.StatusBadRequest, gin.H{"PRESENSI GAGAL": message})
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
	var validasi2 int
	for i := 0; i < len(Nama_mahasiswa); i++ {
		if Input.Nama_mahasiswa == Nama_mahasiswa[i] {
			validasi2++
		}
	}
	if validasi2 == 0 {
		message := "NAMA MAHASISWA: '" + Input.Nama_mahasiswa + "' YANG ANDA INPUTKAN TIDAK DITEMUKAN"
		c.JSON(http.StatusBadRequest, gin.H{"PRESENSI GAGAL": message})

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
	var validasi3 int
	if matkul != nil {
		for i := len(matkul) - 1; i >= 0; i-- {
			if (Input.Matakuliah == matkul[i]) && (Input.Nama_mahasiswa == mahasiswa[i]) && (Input.Tanggal_perkuliahan+"T00:00:00+07:00" == tanggal2[i]) {
				validasi3++
			}
		}
		if validasi3 != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": "ANDA SUDAH MELAKUKAN PRESENSI UNTUK PERTEMUAN INI"})
		}
	}

	//MEMASTIKAN INPUT KEHADIRAN = 1
	var validasi4 int
	if Input.Kehadiran != 1 {
		validasi4 = 1
		c.JSON(http.StatusBadRequest, gin.H{"PRESENSI GAGAL": "SILAKAN INPUT KEHADIRAN DENGAN ANGKA '1'"})
	}

	//JIKA SEMUA VALIDASI LOLOS MAKA DATA AKAN DI INPUTKAN
	if (validasi == 1) && (validasi1 != 0) && (validasi2 != 0) && (validasi3 == 0) && (validasi4 != 1) {
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
		c.JSON(http.StatusOK, gin.H{"PRESENSI BERHASIL DILAKUKAN": t})
	}
}
