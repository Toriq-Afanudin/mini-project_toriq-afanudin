package controls

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini.com/tabels"
)

func EditJadwal(c *gin.Context) {
	//KONEKSI KE DATABASE
	db := c.MustGet("db").(*gorm.DB)

	//TYPE INPUTAN
	type edit struct {
		Matakuliah string `json:"matakuliah"`
		Kelas      string `json:"kelas"`
		Pertemuan  int    `json:"pertemuan"`
		Tanggal    string `json:"tanggal"`
		Jam        string `json:"jam"`
	}

	//VALIDASI JSON
	var ed edit
	if err := c.ShouldBindJSON(&ed); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}

	//VALIDASI TANGGAL PERKULIAHAN
	var column2 tabels.Tanggal
	db.Where("tanggal = ?", ed.Tanggal).Find(&column2)
	if ed.Tanggal != column2.Tanggal {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "tanggal yang di input bukan masa perkuliahan",
		})
		return
	}

	//VALIDASI JAM PERKULIAHAN
	var column3 tabels.Jam
	db.Where("jam = ?", ed.Jam).Find(&column3)
	if ed.Jam != column3.Jam {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "jam bukan waktu perkuliahan",
		})
		return
	}

	//VALIDASI DOSEN MENGAJAR
	var column4 tabels.Dosen
	db.Where("nip = ?", c.Param("nomer")).Find(&column4)
	var column5 tabels.Krs
	db.Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("dosen = ?", column4.Gelar).Find(&column5)
	if ed.Matakuliah != column5.Matakuliah {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "anda tidak mengajar di kelas ini",
		})
		return
	}

	//JIKA LOLOS VALIDASI MAKA DATA AKAN DI UPDATE
	var column tabels.Jadwal
	var tabel []tabels.Jadwal
	db.Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("pertemuan = ?", ed.Pertemuan).Find(&column)
	if ed.Matakuliah == column.Matakuliah {
		db.Model(&tabel).Where("matakuliah = ?", ed.Matakuliah).Where("kelas = ?", ed.Kelas).Where("pertemuan = ?", ed.Pertemuan).Update("tanggal", ed.Tanggal).Update("jam", ed.Jam)
		c.JSON(200, gin.H{
			"status":     "berhasil merubah jadwal",
			"nama":       column4.Gelar,
			"matakuliah": ed.Matakuliah,
			"kelas":      ed.Kelas,
			"tanggal":    ed.Tanggal,
			"jam":        ed.Jam,
		})
		return
	}
}
