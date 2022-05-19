package tabels

type Akumulasi struct {
	Matakuliah string `json:"matakuliah"`
	Nama       string `json:"nama"`
	Hadir      int    `json:"hadir"`
}

type Dosen struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Gelar    string `json:"gelar"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

type Jadwal struct {
	Id         int    `json:"id"`
	Matakuliah string `json:"matakuliah"`
	Kelas      string `json:"kelas"`
	Pertemuan  int    `json:"pertemuan"`
	Tanggal    string `json:"tanggal"`
	Jam        string `json:"jam"`
	Akses      int    `json:"akses"`
	Presensi   int    `json:"presensi"`
	Dosen      int    `json:"dosen"`
}

type Jam struct {
	Id  int    `json:"id"`
	Jam string `json:"jam"`
}

type Krs struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Matakuliah string `json:"matakuliah"`
	Kelas      string `json:"kelas"`
	Hari       string `json:"hari"`
	Jam        string `json:"jam"`
	Dosen      string `json:"dosen"`
}

type Login struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Nomer string `json:"nomer"`
}

type Mahasiswa struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Nim      string `json:"nim"`
	Password string `json:"password"`
}

type Presensi struct {
	Nama       string `json:"nama"`
	Matakuliah string `json:"matakuliah"`
	Kelas      string `json:"kelas"`
	Tanggal    string `json:"tanggal"`
	Dosen      int    `json:"dosen"`
}

type Tanggal struct {
	Id      int    `json:"id"`
	Tanggal string `json:"tanggal"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
