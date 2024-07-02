package config

import "github.com/gorilla/sessions"

const SESSION_ID = "pustakapinjam_auth_sess"

var Store = sessions.NewCookieStore([]byte("qwe3423as987654321"))
