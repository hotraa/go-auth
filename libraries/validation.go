package libraries

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/hotraa/pustaka-pinjam/config"

	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type Validation struct {
	conn *sql.DB
}

func NewValidation() *Validation {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &Validation{
		conn: conn,
	}
}

func (v *Validation) Init() (*validator.Validate, ut.Translator) {
	// memanggil package translator
	translator := id.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("id")

	validate := validator.New()

	// register default translation (id)
	id_translations.RegisterDefaultTranslations(validate, trans)

	// mengubah label default
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		labelName := field.Tag.Get("label")
		return labelName
	})

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} tidak boleh kosong", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} harus berupa alamat email yang valid", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		split_params := strings.Split(params, "-")

		tableName := split_params[0]
		fieldName := split_params[1]
		fieldValue := fl.Field().String()

		return v.checkIsUnique(tableName, fieldName, fieldValue)
	})

	validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		return ut.Add("isunique", "{0} sudah digunakan", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isunique", fe.Field())
		return t
	})

	validate.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} harus memiliki panjang minimal {1} karakter", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field(), fe.Param())
		return t
	})

	validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
		return ut.Add("eqfield", "{0} harus sama dengan {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field(), fe.Param())
		return t
	})

	return validate, trans
}

func (v *Validation) Struct(s interface{}) interface{} {

	validate, trans := v.Init()

	vErrors := make(map[string]interface{})

	err := validate.Struct(s)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vErrors[e.StructField()] = e.Translate(trans)
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil

}

func (v *Validation) checkIsUnique(tableName, fieldName, fieldValue string) bool {

	row, _ := v.conn.Query(`SELECT ` + fieldName + ` from ` + tableName + ` where ` + fieldName + ` = ?`, fieldValue)

	defer row.Close()

	var result string
	for row.Next() {
		row.Scan(&result)
	}

	return result != fieldValue
}
