package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mini_project/models"

	"mini_project/controllers"
)

func main() {
	r := gin.Default()

	//MODEL
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Sistem Presensi Mahasiswa"})
	})

	//REST API UNTUK DOSEN
	r.GET("/jadwal/:nip", controllers.Get_jadwal)
	r.POST("/jadwal/:nip", controllers.Post_jadwal)
	r.PUT("/akses/:nip", controllers.Update_akses)
	r.GET("/akumulasi/:nip", controllers.Akumulasi_matakuliah)
	r.GET("/kehadiran/:nip/:tanggal_perkuliahan", controllers.Presensi_matakuliah)

	//REST API UNTUK MAHASISWA
	r.GET("/penjadwalan", controllers.Get_penjadwalan)
	r.GET("/presensi_mahasiswa/:nama_mahasiswa", controllers.Presensi_mahasiswa)
	r.POST("/presensi", controllers.Post_presensi)
	r.GET("/akumulasi_nama/:nama", controllers.Akumulasi_mahasiswa)

	//REST API UNTUK PROGRAMMER/ MELIHAT TABEL REFERENSI
	r.GET("/dosen", controllers.Dosen)
	r.GET("/kelas", controllers.Kelas)
	r.GET("/mahasiswa", controllers.Mahasiswa)
	r.GET("/jam", controllers.Jam_perkuliahan)

	r.Run()
}
