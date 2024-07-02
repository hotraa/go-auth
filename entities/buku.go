package entities

import "time"

type Buku struct {
	Id             uint
	Sampul         []byte
	Judul          string
	Penulis        string
	TahunTerbit    int
	BidangIlmu     BidangIlmu
	Stok           int64
	Deskripsi 	   string
	WaktuPembuatan time.Time
	WaktuPembaruan time.Time
	JumlahBuku     int
}