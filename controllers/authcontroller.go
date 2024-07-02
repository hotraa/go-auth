package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/hotraa/pustaka-pinjam/config"
	"github.com/hotraa/pustaka-pinjam/entities"
	"github.com/hotraa/pustaka-pinjam/libraries"
	"github.com/hotraa/pustaka-pinjam/models"
	"golang.org/x/crypto/bcrypt"

	modelsbuku "github.com/hotraa/pustaka-pinjam/models/bukumodel"
)

type UserInput struct {
	Email string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			buku := modelsbuku.GetAll()

			// bukuMaks := 8
			// if len(buku) > bukuMaks {
			// 	buku = buku[:bukuMaks]
			// }

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
				"buku" : buku,
			}

			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, data)
		}

	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// proses login
		r.ParseForm()
		UserInput := &UserInput{
			Email: r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)

		} else {

			var user entities.User
			userModel.Where(&user, "email", UserInput.Email)

			var message error
			if user.Email == "" {
				message = errors.New("Email atau Password salah!")
			} else {
				// pengecekan password
				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Email atau Password salah!")
				}
			}

			if message != nil {
				data := map[string]interface{}{
					"error": message,
				}
				temp, _ := template.ParseFiles("views/login.html")
				temp.Execute(w, data)
			} else {
				session, _ := config.Store.Get(r, config.SESSION_ID)
				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["nama_lengkap"] = user.NamaLengkap
				session.Values["role"] = user.Role
				session.Save(r, w)
			
				if user.Role == "admin" {
					http.Redirect(w, r, "/admin", http.StatusSeeOther)
				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi
		// mengambil inputan form
		r.ParseForm()

		user := entities.User{
			NamaLengkap: r.Form.Get("nama_lengkap"),
			Email:       r.Form.Get("email"),
			Password:    r.Form.Get("password"),
			Cpassword:   r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			// insert ke database
			userModel.Create(user)

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, nil)
		}
	}
}