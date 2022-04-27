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

	r.GET("/daftar_mahasiswa", controllers.Daftar_mahasiswa_tampil)

	r.GET("/daftar_dosen", controllers.Dosen_pengampu_tampil)

	r.GET("/daftar_kelas", controllers.Kelas_tampil)

	r.Run()
}
