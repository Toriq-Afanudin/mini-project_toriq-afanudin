package main

import (
	"mini.com/dosens"
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
	r.POST("login", dosens.Login)
	r.PUT("editJadwal/:nip", dosens.EditJadwal)
	r.GET("akumulasi/:nip", dosens.DosenAkumulasi)

	//REST API USER MAHASISWA
	r.POST("presensi/:nim", mahasiswas.Presensi)
	r.GET("akumulasiMahasiswa/:nim", mahasiswas.MahasiswaAkumulasi)
	r.GET("presensi/:nim", mahasiswas.LihatPresensi)

	r.Run()
}
