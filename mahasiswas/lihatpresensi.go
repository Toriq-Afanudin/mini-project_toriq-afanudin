package mahasiswas

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func LihatPresensi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)
	var column tabels.Mahasiswa
	db.Where("nim = ?", c.Param("nim")).Find(&column)
	type presensi struct {
		Matakuliah string
		Tanggal    string
	}
	var tabel []presensi
	db.Raw("SELECT matakuliah, tanggal FROM sistem_presensi.presensis WHERE nama = ?;", column.Nama).Scan(&tabel)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
