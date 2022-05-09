package models

type Penjadwalan struct {
	Id_penjadwalan      int    `json:"id_penjadwalan"`
	Id_kelas            int    `json:"id_kelas"`
	Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
	Jam_perkuliahan     string `json:"jam_perkuliahan"`
}

type Kehadiran struct {
	Id_kehadiran   int `json:"id_kehadiran"`
	Id_penjadwalan int `json:"id_penjadwalan"`
	Id_mahasiswa   int `json:"id_mahasiswa"`
	Kehadiran      int `json:"kehadiran"`
}

type Akumulasi_per_kelas struct {
	Id_akumulasi       int `json:"id_akumulasi"`
	Id_mahasiswa       int `json:"id_mahasiswa"`
	Id_kelas           int `json:"id_kelas"`
	Jumlah_kelas       int `json:"jumlah_kelas"`
	Jumlah_hadir       int `json:"jumlah_hadir"`
	Jumlah_tidak_hadir int `json:"jumlah_tidak_hadir"`
}

type Dosen_pengampu struct {
	Id_dosen     string `gorm:"column:id_dosen"`
	Nama         string `gorm:"column:nama"`
	Niy_nidn_nip string `gorm:"column:niy_nidn_nip"`
}

type Kelas struct {
	Id_kelas       int    `json:"id_kelas"`
	Kode           string `json:"kode"`
	Matakuliah     string `json:"matakuliah"`
	Kelas          string `json:"kelas"`
	Dosen_pengampu string `json:"dosen_pengampu"`
}

type Daftar_mahasiswa struct {
	Id_mahasiswa int    `json:"id_mahasiswa"`
	Nama         string `json:"nama"`
	Nim          string `json:"nim"`
}
