package dosens

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func LihatJadwal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type jadwal struct {
		Matakuliah string
		Kelas      string
		Tanggal    string
		Jam        string
	}
	var column tabels.Dosen
	db.Where("nip = ?", c.Param("nip")).Find(&column)
	var tabel []jadwal
	db.Raw("SELECT matakuliah, kelas, tanggal, jam FROM sistem_presensi.jadwals WHERE dosen = ?;", column.Id).Scan(&tabel)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   tabel,
	})
}
