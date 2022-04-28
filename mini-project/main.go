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

	r.GET("/daftar/mahasiswa", controllers.Daftar_mahasiswa_tampil)

	r.GET("/daftar/dosen", controllers.Dosen_pengampu_tampil)

	r.GET("/daftar/kelas", controllers.Kelas_tampil)

	r.GET("/daftar/setting", controllers.Setting_tampil)
	r.POST("/daftar/setting/tambah", controllers.Setting_tambah)

	r.GET("/daftar/kehadiran", controllers.Kehadiran_tampil)
	r.POST("/daftar/kehadiran/tambah", controllers.Kehadiran_tambah)

	r.Run()
}
