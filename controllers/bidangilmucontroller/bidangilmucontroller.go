package controllers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/hotraa/pustaka-pinjam/entities"
	models "github.com/hotraa/pustaka-pinjam/models/bidangilmumodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	bidangIlmu := models.GetAll()

	data := map[string]any {
		"bidangIlmu" : bidangIlmu,
	}

	var filepath = path.Join("views/bidangilmu/bidangilmu.html")
	var temp, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filepath = path.Join("views/bidangilmu/bidangilmucreate.html")
		var temp, err = template.ParseFiles(filepath)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		temp.Execute(w, nil)
	}
	
	if r.Method == "POST" {
		var bidangIlmu entities.BidangIlmu

		bidangIlmu.Nama = r.FormValue("nama")
		bidangIlmu.WaktuPembuatan = time.Now()
		bidangIlmu.WaktuPembaruan = time.Now()

		if ok := models.Create(bidangIlmu); !ok {
			var filepath = path.Join("views/bidangilmu/bidangilmucreate.html")
			var temp, _ = template.ParseFiles(filepath)
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/bidangilmu", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filepath = path.Join("views/bidangilmu/bidangilmuedit.html")
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

		bidangIlmu := models.Detail(id)
		data := map[string]any {
			"bidangIlmu" : bidangIlmu,
		}
	
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var bidangIlmu entities.BidangIlmu

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		bidangIlmu.Nama = r.FormValue("nama")
		bidangIlmu.WaktuPembaruan = time.Now()

		if ok := models.Update(id, bidangIlmu); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/bidangilmu", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/bidangilmu", http.StatusSeeOther)
}
