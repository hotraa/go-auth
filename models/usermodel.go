package models

import (
	"database/sql"

	"github.com/hotraa/pustaka-pinjam/config"
	"github.com/hotraa/pustaka-pinjam/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
    row, err := u.db.Query(`SELECT id, nama_lengkap, email, password, role FROM users WHERE ` + fieldName + ` = ? LIMIT 1`, fieldValue)

    if err != nil {
        return err
    }

    defer row.Close()

    for row.Next() {
        row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Password, &user.Role)
    }

    return nil
}


func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.db.Exec(`INSERT INTO users (nama_lengkap, email, password) VALUES(? , ?, ?)`,
		user.NamaLengkap, 
		user.Email,
		user.Password,
	)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}

func JumlahUser() (int) {
	var jumlah int
	row := config.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE role = 'biasa'`)

	err := row.Scan(&jumlah)

	if err != nil {
		panic(err)
	}

	return jumlah
}

func GetBukuById(id int) entities.Buku  {
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
