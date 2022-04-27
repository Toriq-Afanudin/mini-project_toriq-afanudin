package models

type Daftar_mahasiswa struct {
	Id_mahasiswa int    `json:"id_mahasiswa"`
	Nama         string `json:"nama"`
	Nim          string `json:"nim"`
}

type Dosen_pengampu struct {
	Id_dosen int    `json:"id_dosen"`
	Nama     string `json:"nama"`
}

type Kelas struct {
	Id_kelas       int    `json:"id_kelas"`
	Kode           string `json:"kode"`
	Matakuliah     string `json:"matakuliah"`
	Kelas          string `json:"kelas"`
	Dosen_pengampu string `json:"dosen_pengampu"`
	Id_dosen       string `json:"id_dosen"`
}
