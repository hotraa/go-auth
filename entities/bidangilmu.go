package entities

import "time"

type BidangIlmu struct {
	Id            		 uint
	Nama          		 string
	WaktuPembuatan 		 time.Time
	WaktuPembaruan		 time.Time
	JumlahBidangIlmu     int
}