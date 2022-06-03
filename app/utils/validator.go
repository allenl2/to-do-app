package utils

import (
	"fmt"
	"log"
	"strings"
	"to-do-app/app/models"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(input models.User) *fiber.Error {
	var errors []string
	err := validate.Struct(input)

	if err != nil {
		//check if the validation itself doesn't work
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			errors = append(errors, err.Error())
		} else {
			for _, err := range err.(validator.ValidationErrors) {
				currError := fmt.Sprintf("Field: `%v`, Contraint: `%v`, Contraint Value: `%v`", err.Field(), err.Tag(), err.Param())
				errors = append(errors, currError)
				log.Println("Validation Error -", err.Field(), ":", err.Tag(), "->", err.Param())
			}
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}
	return nil
}

func ValidateUserPassword(input models.User) *fiber.Error {
	var errors []string

	validate.RegisterStructValidation(UserStructValidation, models.User{})

	err := validate.Struct(input)

	if err != nil {
		//check if the validation itself doesn't work
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			errors = append(errors, err.Error())
		} else {
			for _, err := range err.(validator.ValidationErrors) {
				currError := fmt.Sprintf("Field: `%v`, Contraint: `%v`, Contraint Value: `%v`", err.Field(), err.Tag(), err.Param())
				errors = append(errors, currError)
				log.Println("Validation Error -", err.Field(), ":", err.Tag(), "->", err.Param())
			}
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}
	return nil
}

//validates fields on user struct
func UserStructValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(models.User)

	//validates that username exists
	if len(user.Username) == 0 {
		sl.ReportError(user.Username, "username", "Username", "must exist", "")
	}

	//validate password requirements
	hasUpper := false
	hasLower := false
	hasNum := false

	for _, char := range user.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNum = true
		}
	}

	if !hasUpper {
		sl.ReportError(user.Password, "password", "Password", "uppercase letter", "at least one")
	}
	if !hasLower {
		sl.ReportError(user.Password, "password", "Password", "lowercase letter", "at least one")
	}
	if !hasNum {
		sl.ReportError(user.Password, "password", "Password", "number", "at least one")
	}

}
