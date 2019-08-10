package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name string `json:"name" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	msg := ""
	for _, e := range err.(validator.ValidationErrors).Translate(cv.trans) {
		if msg != "" {
			msg += ","
		}
		msg += e
	}
	return errors.New(msg)
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	e := echo.New()

	v, trans := setupValidator()
	e.Validator = &CustomValidator{validator: v, trans: trans}
	e.POST("/", test)
	e.Logger.Fatal(e.Start(":1113"))
}

func setupValidator() (*validator.Validate, ut.Translator) {
	uni := ut.New(ja_JP.New(), ja_JP.New())

	// 変換処理生成
	jaTrans, found := uni.GetTranslator("ja_JP")
	if !found {
		log.Fatal("translator not found")
		return nil, nil
	}

	v := validator.New()
	// フィールド名を登録
	_ = jaTrans.Add("Name", "名前", false)

	// エラーメッセージを登録
	v.RegisterTranslation("required", jaTrans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}は必須です。", false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		// フィールド名取得
		fld, _ := ut.T(fe.Field())
		// エラーメッセージ取得
		t, _ := ut.T(fe.Tag(), fld)
		return t
	})

	return v, jaTrans
}

func test(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, &Error{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
