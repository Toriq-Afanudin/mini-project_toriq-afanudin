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

	r.GET("/mahasiswa", controllers.Daftar_mahasiswa_tampil)

	r.GET("/dosen", controllers.Dosen_pengampu_tampil)

	r.GET("/kelas", controllers.Kelas_tampil)

	r.GET("/penjadwalan", controllers.Penjadwalan_tampil)
	r.POST("/penjadwalan", controllers.Penjadwalan_tambah)
	r.PUT("/penjadwalan/:id_penjadwalan", controllers.Penjadwalan_ubah)
	r.DELETE("/penjadwalan/:id_penjadwalan", controllers.Penjadwalan_hapus)

	r.GET("/kehadiran", controllers.Kehadiran_tampil)
	r.POST("/kehadiran", controllers.Kehadiran_tambah)
	r.PUT("/kehadiran/:id_kehadiran", controllers.Kehadiran_ubah)
	r.DELETE("/kehadiran/:id_kehadiran", controllers.Kehadiran_hapus)

	r.GET("/akumulasi", controllers.Akumulasi_tampil)

	r.Run()
}
