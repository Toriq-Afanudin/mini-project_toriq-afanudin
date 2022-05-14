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
	r.GET("/penjadwalan", controllers.Get_penjadwalan)
	r.POST("/penjadwalan", controllers.Post_penjadwalan)
	r.PUT("/akses", controllers.Update_akses)

	//REST API UNTUK MAHASISWA
	r.GET("/presensi", controllers.Get_presensi)
	r.POST("/presensi", controllers.Post_presensi)

	//REST API UNTUK PROGRAMMER/ MELIHAT TABEL REFERENSI
	r.GET("/dosen", controllers.Dosen)
	r.GET("/kelas", controllers.Kelas)
	r.GET("/mahasiswa", controllers.Mahasiswa)
	r.GET("/akumulasi", controllers.Get_akumulasi)
	r.GET("/jam", controllers.Jam_perkuliahan)

	r.Run()
}
