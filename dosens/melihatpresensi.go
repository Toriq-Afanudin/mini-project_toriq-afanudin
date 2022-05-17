package dosens

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func MelihatPresensi(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)
	var column tabels.Dosen
	db.Where("nip = ?", c.Param("nip")).Find(&column)
	type presensi struct {
		Nama       string
		Matakuliah string
		Tanggal    string
	}
	var tabel []presensi
	db.Raw("SELECT nama,matakuliah,tanggal FROM sistem_presensi.presensis;").Scan(&tabel)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
