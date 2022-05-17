package main

import (
	"mini.com/controllers"
	"mini.com/dosens"
	"mini.com/mahasiswas"
	"mini.com/tabels"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := tabels.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	//REST API USER DOSEN
	r.PUT("editJadwal/:nip", dosens.EditJadwal)
	r.GET("akumulasi/:nip", dosens.DosenAkumulasi)
	r.GET("melihatpresensi/:nip", dosens.MelihatPresensi)

	//REST API USER MAHASISWA
	r.POST("presensi/:nim", mahasiswas.Presensi)
	r.GET("akumulasiMahasiswa/:nim", mahasiswas.MahasiswaAkumulasi)
	r.GET("presensi/:nim", mahasiswas.LihatPresensi)

	//MAHASISWA DAN DOSEN
	r.POST("login", controllers.Login)
	r.GET("melihatjadwal", controllers.LihatJadwal)

	r.Run()
}
