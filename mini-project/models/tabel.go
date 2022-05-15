package models

type Penjadwalan struct {
	Matakuliah                 string `json:"matakuliah"`
	Dosen_pengampu_tanpa_gelar string `json:"dosen_pengampu_tanpa_gelar"`
	Tanggal_perkuliahan        string `json:"tanggal_perkuliahan"`
	Jam_perkuliahan            string `json:"jam_perkuliahan"`
	Akses                      int    `json:"akses"`
}

type Kehadiran struct {
	Matakuliah          string `json:"matakuliah"`
	Nama_mahasiswa      string `json:"nama_mahasiswa"`
	Tanggal_perkuliahan string `json:"tanggal_perkuliahan"`
}

type Dosen_pengampu struct {
	Id_dosen    string `json:"id_dosen"`
	Nama        string `json:"nama"`
	Nip         string `json:"nip"`
	Tanpa_gelar string `json:"tanpa_gelar"`
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
	Matakuliah string `json:"matakuliah"`
	Nama       string `json:"nama"`
	Pertemuan  int    `json:"pertemuan"`
	Hadir      int    `json:"hadir"`
	Tidak      int    `json:"tidak"`
}

type Tanggal struct {
	Tanggal string `json:"tanggal"`
}

type Libur struct {
	Tanggal    string `json:"tanggal"`
	Keterangan string `json:"keterangan"`
}
