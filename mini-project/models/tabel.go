package models

type Penjadwalan struct {
	Id_penjadwalan             int    `json:"id_penjadwalan"`
	Matakuliah                 string `json:"matakuliah"`
	Dosen_pengampu_tanpa_gelar string `json:"dosen_pengampu_tanpa_gelar"`
	Tanggal_perkuliahan        string `json:"tanggal_perkuliahan"`
	Jam_perkuliahan            string `json:"jam_perkuliahan"`
	Akses                      int    `json:"akses"`
}

type Kehadiran struct {
	Id_kehadiran        int    `json:"id_kehadiran"`
	Matakuliah          string `json:"matakuliah"`
	Nama_mahasiswa      string `json:"nama_mahasiswa"`
	Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
}

type Dosen_pengampu struct {
	Id_dosen     string `json:"id_dosen"`
	Nama         string `json:"nama"`
	Niy_nidn_nip string `json:"niy_nidn_nip"`
}

type Kelas struct {
	Id_kelas                   int    `json:"id_kelas"`
	Kode                       string `json:"kode"`
	Matakuliah                 string `json:"matakuliah"`
	Kelas                      string `json:"kelas"`
	Dosen_pengampu             string `json:"dosen_pengampu"`
	Dosen_pengampu_tanpa_gelar string `json:"dosen_pengampu_tanpa_gelar"`
}

type Daftar_mahasiswa struct {
	Id_mahasiswa int    `json:"id_mahasiswa"`
	Nama         string `json:"nama"`
	Nim          string `json:"nim"`
}

type Jam_perkuliahan struct {
	Id_jam          int    `json:"Id_jam"`
	Jam_perkuliahan string `json:"jam_perkuliahan"`
}

type Akumulasi struct {
	Id_akumulasi int    `json:"id_akumulasi"`
	Matakuliah   string `json:"matakuliah"`
	Nama         string `json:"nama"`
	Pertemuan    int    `json:"pertemuan"`
	Hadir        int    `json:"hadir"`
	Tidak        int    `json:"tidak"`
}
