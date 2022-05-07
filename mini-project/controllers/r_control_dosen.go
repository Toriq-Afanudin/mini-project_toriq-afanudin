package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Dosen_pengampu struct {
	Id_dosen     int    `json:"id_dosen"`
	Nama         string `json:"nama"`
	Niy_nidn_nip string `json:"niy_nidn_nip"`
}

//TAMPIL DATA
func Dosen_pengampu_tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Daftar_dosen []Dosen_pengampu

	db.Find(&Daftar_dosen)
	c.JSON(http.StatusOK, gin.H{"data": Daftar_dosen})
}
