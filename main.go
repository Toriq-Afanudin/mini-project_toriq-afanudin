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
	r.PUT("editJadwal/:nomer", controls.EditJadwal)

	r.POST("presensi/:nim", controls.Presensi)

	r.Run()
}
