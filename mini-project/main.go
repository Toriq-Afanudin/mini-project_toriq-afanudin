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

	r.GET("/dosen", controllers.Dosen)

	r.GET("/kelas", controllers.Kelas)

	r.GET("/mahasiswa", controllers.Mahasiswa)

	r.GET("/penjadwalan", controllers.Get_penjadwalan)
	r.POST("/penjadwalan", controllers.Post_penjadwalan)
	// r.PUT("/penjadwalan/:id_penjadwalan", controllers.Penjadwalan_ubah)
	// r.DELETE("/penjadwalan/:id_penjadwalan", controllers.Penjadwalan_hapus)

	r.GET("/presensi", controllers.Get_presensi)
	r.POST("/presensi", controllers.Post_presensi)
	// r.PUT("/kehadiran/:id_kehadiran", controllers.Kehadiran_ubah)
	// r.DELETE("/kehadiran/:id_kehadiran", controllers.Kehadiran_hapus)

	r.GET("/akumulasi", controllers.Get_akumulasi)

	r.GET("/jam", controllers.Jam_perkuliahan)

	r.Run()
}
