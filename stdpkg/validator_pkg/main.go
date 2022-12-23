package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// User1 contains user information
type User1 struct {
	FirstName   string `json:"fname" validate:"alpha"`
	LastName    string `json:"lname" validate:"alpha"`
	Age         uint8  `validate:"gte=20,lte=65"`
	Email       string `json:"e-mail" validate:"required,email"`
	JoiningDate string `validate:"datetime"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func normal() {

	validate = validator.New()

	user := &User1{
		FirstName:   "Test25",
		LastName:    "Test",
		Age:         75,
		Email:       "Badger.Smith@",
		JoiningDate: "005-25-10",
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("%v's type is [%v], tag:%v value_got:%v param:%v", err.Field(), err.Kind(), err.Tag(), err.Value(), err.Param())
			fmt.Println(msg)
		}
		return
	}
}

// ----------------------------------------------------------------
func main() {
	normal()
	custom_validaton()
}

// ----------------------------------------------------------------

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"passwd"` // <-- a custom validation rule
}

func custom_validaton() {
	v := validator.New()
	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	a := User{
		Email:    "something@gmail.com",
		Name:     "A girl has no name",
		Password: "1234",
	}
	err := v.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}
