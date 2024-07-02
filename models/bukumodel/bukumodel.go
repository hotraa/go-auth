package models

import (
	"net/http"

	"github.com/hotraa/pustaka-pinjam/config"
	"github.com/hotraa/pustaka-pinjam/entities"
)

func GetAll() []entities.Buku {
	rows, err := config.DB.Query(`
		SELECT
			buku.id,
			buku.sampul,
			buku.judul,
			buku.penulis,
			buku.tahun_terbit,
			bidang_ilmu.nama as nama_bidang_ilmu,
			buku.stok,
			buku.deskripsi,
			buku.waktu_pembuatan,
			buku.waktu_pembaruan
		FROM buku
		JOIN bidang_ilmu ON buku.id_bidang_ilmu = bidang_ilmu.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var buku []entities.Buku

	for rows.Next() {
		var bukuTemp entities.Buku
		err := rows.Scan(
			&bukuTemp.Id,
			&bukuTemp.Sampul,
			&bukuTemp.Judul,
			&bukuTemp.Penulis,
			&bukuTemp.TahunTerbit,
			&bukuTemp.BidangIlmu.Nama,
			&bukuTemp.Stok,
			&bukuTemp.Deskripsi,
			&bukuTemp.WaktuPembuatan,
			&bukuTemp.WaktuPembaruan,
		)

		if err != nil {
			panic(err)
		}

		buku = append(buku, bukuTemp)
	}

	return buku
}

func GetImage(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    var imageData []byte
    err := config.DB.QueryRow(`SELECT sampul FROM buku WHERE id = ?`, id).Scan(&imageData)
    if err != nil {
        http.Error(w, "Gambar tidak ditemukan", http.StatusNotFound)
        return
    }

    // Set the content type
    w.Header().Set("Content-Type", "image/png")
    w.Write(imageData)
}

func Create(buku entities.Buku) bool {
	result, err := config.DB.Exec(`
		INSERT INTO buku (
			sampul, judul, penulis, tahun_terbit, id_bidang_ilmu, stok, deskripsi, waktu_pembuatan, waktu_pembaruan
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		buku.Sampul,
		buku.Judul,
		buku.Penulis,
		buku.TahunTerbit,
		buku.BidangIlmu.Id,
		buku.Stok,
		buku.Deskripsi,
		buku.WaktuPembuatan,
		buku.WaktuPembaruan,
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

func Detail(id int) entities.Buku {
	row := config.DB.QueryRow(`
		SELECT
			buku.id,
			buku.sampul,
			buku.judul,
			buku.penulis,
			buku.tahun_terbit,
			bidang_ilmu.nama as nama_bidang_ilmu,
			buku.stok,
			buku.deskripsi,
			buku.waktu_pembuatan,
			buku.waktu_pembaruan
		FROM buku
		JOIN bidang_ilmu ON buku.id_bidang_ilmu = bidang_ilmu.id
		WHERE buku.id = ?
	`, id)

	var bukuTemp entities.Buku

	err := row.Scan(
		&bukuTemp.Id,
		&bukuTemp.Sampul,
		&bukuTemp.Judul,
		&bukuTemp.Penulis,
		&bukuTemp.TahunTerbit,
		&bukuTemp.BidangIlmu.Nama,
		&bukuTemp.Stok,
		&bukuTemp.Deskripsi,
		&bukuTemp.WaktuPembuatan,
		&bukuTemp.WaktuPembaruan,
	)

	if err != nil {
		panic(err)
	}

	return bukuTemp
}

func Update(id int, buku entities.Buku) bool {
	query, err := config.DB.Exec(`
		UPDATE buku SET
			sampul = ?,
			judul = ?,
			penulis = ?,
			tahun_terbit = ?,
			id_bidang_ilmu = ?,
			stok = ?,
			deskripsi = ?,
			waktu_pembaruan = ?
		WHERE id = ?`,
		buku.Sampul,
		buku.Judul,
		buku.Penulis,
		buku.TahunTerbit,
		buku.BidangIlmu.Id,
		buku.Stok,
		buku.Deskripsi,
		buku.WaktuPembaruan,
		id,
	)

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
	_, err := config.DB.Exec(`DELETE FROM buku WHERE id = ?`, id)
	return err
}

func JumlahBuku() (int) {
	var jumlah int
	row := config.DB.QueryRow("SELECT COUNT(*) FROM buku")

	err := row.Scan(&jumlah)

	if err != nil {
		panic(err)
	}

	return jumlah
}

// func DetailStokBuku(id int) entities.Buku {
// 	row := config.DB.QueryRow(`
// 		SELECT
// 			buku.id,
// 			buku.stok,
// 			buku.waktu_pembaruan
// 		FROM buku
// 		WHERE buku.id = ?
// 	`, id)

// 	var bukuTemp entities.Buku

// 	err := row.Scan(
// 		&bukuTemp.Id,
// 		&bukuTemp.Stok,
// 		&bukuTemp.WaktuPembaruan,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return bukuTemp
// }

func UpdateStokBuku(id int, buku entities.Buku) bool {
	query, err := config.DB.Exec(`
		UPDATE buku SET
			stok = stok - 1,
			waktu_pembaruan = ?
		WHERE id = ?`,
		buku.WaktuPembaruan,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0 
}
