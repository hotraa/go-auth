package main

import (
	"log"
	"net/http"

	authcontroller "github.com/hotraa/pustaka-pinjam/controllers"
	pagecontroller "github.com/hotraa/pustaka-pinjam/controllers"
	bidangilmucontroller "github.com/hotraa/pustaka-pinjam/controllers/bidangilmucontroller"
	bukucontroller "github.com/hotraa/pustaka-pinjam/controllers/bukucontroller"
	bukumodel "github.com/hotraa/pustaka-pinjam/models/bukumodel"
)

func main() {
	// Menyajikan file statis dari direktori "public"
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))
	
	// User
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/kategori", pagecontroller.Kategori)
	http.HandleFunc("/tentang", pagecontroller.Tentang)
	http.HandleFunc("/faq", pagecontroller.Faq)
	http.HandleFunc("/kontak", pagecontroller.Kontak)

	// --- Admin ---
	http.Handle("/admin", pagecontroller.AdminMiddleware(http.HandlerFunc(pagecontroller.Admin)))
	// http.HandleFunc("/admin", pagecontroller.Admin)
	// http.HandleFunc("/bidangilmu", pagecontroller.BidangIlmu)
	// http.HandleFunc("/buku", pagecontroller.Buku)
	http.Handle("/peminjaman", pagecontroller.AdminMiddleware(http.HandlerFunc(pagecontroller.Peminjaman)))
	
	// --- Bidang Ilmu ---
	// http.HandleFunc("/bidangilmu", bidangilmucontroller.Index)
	// http.HandleFunc("/bidangilmu/add", bidangilmucontroller.Add)
	// http.HandleFunc("/bidangilmu/edit", bidangilmucontroller.Edit)
	http.Handle("/bidangilmu", pagecontroller.AdminMiddleware(http.HandlerFunc(bidangilmucontroller.Index)))
    http.Handle("/bidangilmu/add", pagecontroller.AdminMiddleware(http.HandlerFunc(bidangilmucontroller.Add)))
    http.Handle("/bidangilmu/edit", pagecontroller.AdminMiddleware(http.HandlerFunc(bidangilmucontroller.Edit)))
	http.HandleFunc("/bidangilmu/delete", bidangilmucontroller.Delete)
	
	// --- Buku ---
	// http.HandleFunc("/buku", bukucontroller.Index)
	http.HandleFunc("/image", bukumodel.GetImage)
	// http.HandleFunc("/buku/add", bukucontroller.Add)
	// http.HandleFunc("/buku/detail", bukucontroller.Detail)
	// http.HandleFunc("/buku/edit", bukucontroller.Edit)
	http.Handle("/buku", pagecontroller.AdminMiddleware(http.HandlerFunc(bukucontroller.Index)))
    http.Handle("/buku/add", pagecontroller.AdminMiddleware(http.HandlerFunc(bukucontroller.Add)))
    http.Handle("/buku/detail", pagecontroller.AdminMiddleware(http.HandlerFunc(bukucontroller.Detail)))
    http.Handle("/buku/edit", pagecontroller.AdminMiddleware(http.HandlerFunc(bukucontroller.Edit)))
	http.HandleFunc("/buku/delete", bukucontroller.Delete)

	// --- Peminjaman User ---
	http.HandleFunc("/detailpeminjaman", pagecontroller.DetailPeminjaman)
	http.HandleFunc("/peminjamanbuku", bukucontroller.PeminjamanBuku)

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
