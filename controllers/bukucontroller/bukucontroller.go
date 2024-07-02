package controllers

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/hotraa/pustaka-pinjam/entities"
	modelsrelasi "github.com/hotraa/pustaka-pinjam/models/bidangilmumodel"
	models "github.com/hotraa/pustaka-pinjam/models/bukumodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	buku := models.GetAll()

	data := map[string]any {
		"buku" : buku,
	}

	var filepath = path.Join("views/buku/buku.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filepath = path.Join("views/buku/bukucreate.html")
		var temp, err = template.ParseFiles(filepath)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		bidangIlmu := modelsrelasi.GetAll()
		data := map[string]any {
			"bidangIlmu" : bidangIlmu,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var buku entities.Buku

		file, _, err := r.FormFile("sampul")
		if err != nil {
			http.Error(w, "Tidak dapat membaca file gambar/sampul", http.StatusBadRequest)
			return
		}
		defer file.Close()

		var buf bytes.Buffer
		_, err = io.Copy(&buf, file)
		if err != nil {
			http.Error(w, "Tidak dapat membaca konten file gambar/sampul", http.StatusInternalServerError)
			return
		}
		sampul := buf.Bytes()

		tahunTerbit, err := strconv.Atoi(r.FormValue("tahun_terbit"))
		if err != nil {
			panic(err)
		}

		stok, err := strconv.Atoi(r.FormValue("stok"))
		if err != nil {
			panic(err)
		}

		bidangIlmuId, err := strconv.Atoi(r.FormValue("id_bidang_ilmu"))
		if err != nil {
			panic(err)
		}

		buku.Sampul = []byte(sampul)
		buku.Judul = r.FormValue("judul")
		buku.Penulis = r.FormValue("penulis")
		buku.TahunTerbit = int(tahunTerbit)
		buku.BidangIlmu.Id = uint(bidangIlmuId)
		buku.Stok = int64(stok)
		buku.Deskripsi = r.FormValue("deskripsi")
		buku.WaktuPembuatan = time.Now()
		buku.WaktuPembaruan = time.Now()

		if ok := models.Create(buku); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/buku", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	buku := models.Detail(id)
	data := map[string]any {
		"buku" : buku,
	}

	var filepath = path.Join("views/buku/bukudetail.html")
	temp, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filepath = path.Join("views/buku/bukuedit.html")
		var temp, err = template.ParseFiles(filepath)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		buku := models.Detail(id)

		bidangIlmu := modelsrelasi.GetAll()
		data := map[string]any {
			"bidangIlmu" : bidangIlmu,
			"buku" : buku,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var buku entities.Buku

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		file, _, err := r.FormFile("sampul")
		if err != nil {
			http.Error(w, "Tidak dapat membaca file gambar/sampul", http.StatusBadRequest)
			return
		}
		defer file.Close()

		var buf bytes.Buffer
		_, err = io.Copy(&buf, file)
		if err != nil {
			http.Error(w, "Tidak dapat membaca konten file gambar/sampul", http.StatusInternalServerError)
			return
		}
		sampul := buf.Bytes()

		tahunTerbit, err := strconv.Atoi(r.FormValue("tahun_terbit"))
		if err != nil {
			panic(err)
		}

		stok, err := strconv.Atoi(r.FormValue("stok"))
		if err != nil {
			panic(err)
		}

		bidangIlmuId, err := strconv.Atoi(r.FormValue("id_bidang_ilmu"))
		if err != nil {
			panic(err)
		}

		buku.Sampul = []byte(sampul)
		buku.Judul = r.FormValue("judul")
		buku.Penulis = r.FormValue("penulis")
		buku.TahunTerbit = int(tahunTerbit)
		buku.BidangIlmu.Id = uint(bidangIlmuId)
		buku.Stok = int64(stok)
		buku.Deskripsi = r.FormValue("deskripsi")
		buku.WaktuPembaruan = time.Now()

		if ok := models.Update(id, buku); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/buku", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/buku", http.StatusSeeOther)
}

func PeminjamanBuku(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	var filepath = path.Join("views/detailpeminjaman.html")
	// 	var temp, err = template.ParseFiles(filepath)
	
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	
	// 	idString := r.URL.Query().Get("id")
	// 	id, err := strconv.Atoi(idString)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	buku := models.DetailStokBuku(id)

	// 	// bidangIlmu := modelsrelasi.GetAll()
	// 	data := map[string]any {
	// 		// "bidangIlmu" : bidangIlmu,
	// 		"buku" : buku,
	// 	}

	// 	temp.Execute(w, data)
	// }

	if r.Method == "POST" {
		var buku entities.Buku

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		stok, err := strconv.Atoi(r.FormValue("stok"))
		if err != nil {
			panic(err)
		}

		buku.Stok = int64(stok)
		buku.WaktuPembaruan = time.Now()

		if ok := models.UpdateStokBuku(id, buku); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/peminjamanbuku", http.StatusSeeOther)
	}
}

// var modelBuku = models.NewModelBuku()

// func PeminjamanBuku(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		r.ParseForm()
// 		tanggalPeminjaman := r.FormValue("tanggalPeminjaman")

// 		// Kurangi stok buku di database
// 		idBuku := r.FormValue("idBuku") // Dapatkan idBuku dari form
// 		err := modelBuku.KurangiStok(idBuku)
// 		if err != nil {
// 			http.Error(w, "Gagal meminjam buku", http.StatusInternalServerError)
// 			return
// 		}

// 		// Generate kode peminjaman
// 		kodePeminjaman := generateKodePeminjaman(6)

// 		// Set tanggal kembali (7 hari dari tanggal peminjaman)
// 		tanggalKembali, _ := time.Parse("2006-01-02", tanggalPeminjaman)
// 		tanggalKembali = tanggalKembali.AddDate(0, 0, 7)
// 		// tanggalKembaliStr := tanggalKembali.Format("2006-01-02")

// 		// Redirect ke detail peminjaman dengan ID buku dan kode peminjaman
// 		http.Redirect(w, r, "/detailpeminjaman?id=" + idBuku + "&kode=" + kodePeminjaman, http.StatusSeeOther)
// 	}
// }

// func generateKodePeminjaman(n int) string {
// 	const huruf = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	rand.Seed(time.Now().UnixNano())
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = huruf[rand.Intn(len(huruf))]
// 	}
// 	return string(b)
// }