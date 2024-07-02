package models

import (
	"github.com/hotraa/pustaka-pinjam/config"
	"github.com/hotraa/pustaka-pinjam/entities"
)

func GetAll() []entities.BidangIlmu {
	rows, err := config.DB.Query(`SELECT * FROM bidang_ilmu`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var bidangIlmu []entities.BidangIlmu

	for rows.Next() {
		var bidangIlmuTemp entities.BidangIlmu
		if err := rows.Scan(&bidangIlmuTemp.Id, &bidangIlmuTemp.Nama, &bidangIlmuTemp.WaktuPembuatan, &bidangIlmuTemp.WaktuPembaruan); err != nil {
			panic(err)
		}

		bidangIlmu = append(bidangIlmu, bidangIlmuTemp)
	}

	return bidangIlmu
}

func Create(bidangIlmu entities.BidangIlmu) bool {
	result, err := config.DB.Exec(`
		INSERT INTO bidang_ilmu (nama, waktu_pembuatan, waktu_pembaruan)
		VALUE (?, ?, ?)`,
		bidangIlmu.Nama,
		bidangIlmu.WaktuPembuatan,
		bidangIlmu.WaktuPembaruan,
	)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.BidangIlmu {
	row := config.DB.QueryRow(`SELECT id, nama FROM bidang_ilmu WHERE id = ?`, id)

	var bidangIlmu entities.BidangIlmu
	if err := row.Scan(&bidangIlmu.Id, &bidangIlmu.Nama); err != nil {
		panic(err.Error())
	}

	return bidangIlmu
}

func Update(id int, bidangIlmu entities.BidangIlmu) bool {
	query, err := config.DB.Exec(`UPDATE bidang_ilmu SET nama = ?, waktu_pembaruan = ? WHERE id = ?`, bidangIlmu.Nama, bidangIlmu.WaktuPembaruan, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE from bidang_ilmu WHERE id = ?`, id)
	return err
}

func JumlahBidangIlmu() (int) {
	var jumlah int
	row := config.DB.QueryRow("SELECT COUNT(*) FROM bidang_ilmu")

	err := row.Scan(&jumlah)

	if err != nil {
		panic(err)
	}

	return jumlah
}

