package controllers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/hotraa/pustaka-pinjam/config"
	modelspeminjaman "github.com/hotraa/pustaka-pinjam/models"
	modelsrelasi "github.com/hotraa/pustaka-pinjam/models/bidangilmumodel"
	models "github.com/hotraa/pustaka-pinjam/models/bukumodel"
)

func Kategori(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/kategori.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, nil)
}

func Tentang(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/tentang.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, nil)
}

func Faq(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/faq.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, nil)
}

func Kontak(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/kontak.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, nil)
}

// --- Admin ---
func Admin(w http.ResponseWriter, r *http.Request) {
	buku := models.GetAll()
	jumlahBuku := models.JumlahBuku()
	jumlahBidangIlmu := modelsrelasi.JumlahBidangIlmu()
	jumlahUser := modelspeminjaman.JumlahUser()

	data := map[string]any {
		"buku" : buku,
		"jumlahBuku" : jumlahBuku,
		"jumlahBidangIlmu" : jumlahBidangIlmu,
		"jumlahUser" : jumlahUser,
	}

	var filepath = path.Join("views/dashboard/dashboard.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, data)
}

func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := config.Store.Get(r, config.SESSION_ID)

        if session.Values["role"] != "admin" {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func Peminjaman(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views/dashboard/peminjaman.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	temp.Execute(w, nil)
}

func DetailPeminjaman(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	kodePeminjaman := r.URL.Query().Get("kode")
	tanggalPeminjaman := r.URL.Query().Get("tanggalPeminjaman")

    buku := modelspeminjaman.GetBukuById(id)

    var filepath = path.Join("views/detailpeminjaman.html")
    var temp, errTemplate = template.ParseFiles(filepath)

    if errTemplate != nil {
        http.Error(w, errTemplate.Error(), http.StatusInternalServerError)
        return
    }

	data := map[string]interface{}{
		"buku":             buku,
		"kodePeminjaman":   kodePeminjaman,
		"tanggalPeminjaman": tanggalPeminjaman,
	}

    temp.Execute(w, data)
}