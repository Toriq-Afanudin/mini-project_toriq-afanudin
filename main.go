package main

import (
	"mini.com/controls"
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

	r.POST("login", controls.Login)
	r.PUT("editJadwal/:nip", controls.EditJadwal)
	r.GET("akumulasi/:nip", controls.DosenAkumulasi)

	r.POST("presensi/:nim", controls.Presensi)
	r.GET("akumulasiMahasiswa/:nim", controls.MahasiswaAkumulasi)
	r.GET("presensi/:nim", controls.LihatPresensi)

	r.Run()
}
