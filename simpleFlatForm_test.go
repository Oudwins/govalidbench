package main

import (
	"regexp"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/asaskevich/govalidator"
	validate "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-playground/validator/v10"
)

type SimpleFlatForm struct {
	Id        string `json:"id" validate:"required,uuid4" valid:"required,uuidv4"`
	Name      string `json:"name" validate:"required" valid:"required"`
	Email     string `json:"email" validate:"required,email" valid:"required,email"`
	Age       int    `json:"age" validate:"required,min=18,max=100" valid:"required,min=18,max=100"`
	Isbn      string `json:"isbn" validate:"required,isbn10" valid:"required,isbn10"`
	Address   string `json:"address" validate:"required" valid:"required"`
	City      string `json:"city" validate:"required,oneof=New York Los Angeles Chicago Houston Phoenix" valid:"required,oneof=New York Los Angeles Chicago Houston Phoenix"`
	State     string `json:"state" validate:"required" valid:"required"`
	ZipCode   string `json:"zip_code" validate:"required" valid:"required"`
	Country   string `json:"country" validate:"required" valid:"required"`
	Website   string `json:"website" validate:"required,url" valid:"required,url"`
	CreatedAt string `json:"created_at" validate:"required" valid:"required"`
}

var cities = []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}

var isbnRegex = regexp.MustCompile(`^(?:[0-9]{9}X|[0-9]{10})$`)

// Zog

var SimpleFlatFormZogSchema = z.Struct(
	z.Schema{
		"id":        z.String().Required().UUID(),
		"name":      z.String().Required(),
		"email":     z.String().Required().Email(),
		"age":       z.Int().Required().GTE(18).LTE(100),
		"isbn":      z.String().Required().Match(isbnRegex),
		"address":   z.String().Required(),
		"city":      z.String().Required().OneOf(cities),
		"state":     z.String().Required(),
		"zipCode":   z.String().Required(),
		"country":   z.String().Required(),
		"website":   z.String().Required().URL(),
		"createdAt": z.String().Required(),
	},
)

func OzzoFlatForm(form *SimpleFlatForm) error {
	return validate.ValidateStruct(form,
		validate.Field(&form.Id, validate.Required, is.UUIDv4),
		validate.Field(&form.Name, validate.Required),
		validate.Field(&form.Email, validate.Required, is.Email),
		validate.Field(&form.Age, validate.Required, validate.Min(18), validate.Max(100)),
		validate.Field(&form.Isbn, validate.Required, is.ISBN10),
		validate.Field(&form.Address, validate.Required),
		validate.Field(&form.City, validate.Required, validate.In(cities)),
		validate.Field(&form.State, validate.Required),
		validate.Field(&form.ZipCode, validate.Required),
		validate.Field(&form.Country, validate.Required),
		validate.Field(&form.Website, validate.Required, is.URL),
		validate.Field(&form.CreatedAt, validate.Required),
	)

}

// Package Validator

var SimpleFlatFormPkgValidator = validator.New(validator.WithRequiredStructEnabled())

func BenchmarkFlatFormSchemaFilled(b *testing.B) {
	filledForm := SimpleFlatForm{
		Id:        "123e4567-e89b-12d3-a456-426614174000",
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Age:       25,
		Isbn:      "0-19-853453-1",
		Address:   "123 Main St",
		City:      "New York",
		State:     "NY",
		ZipCode:   "10001",
		Country:   "USA",
		Website:   "https://www.example.com",
		CreatedAt: "2021-01-01",
	}
	data := map[string]any{
		"id":        "123e4567-e89b-12d3-a456-426614174000",
		"name":      "John Doe",
		"email":     "john.doe@example.com",
		"age":       25,
		"isbn":      "0-19-853453-1",
		"address":   "123 Main St",
		"city":      "New York",
		"state":     "NY",
		"zipCode":   "10001",
		"country":   "USA",
		"website":   "https://www.example.com",
		"createdAt": "2021-01-01",
	}

	b.Run("Zog Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SimpleFlatFormZogSchema.Validate(&filledForm)
		}
	})

	b.Run("Zog Parse", func(b *testing.B) {
		var form SimpleFlatForm
		for i := 0; i < b.N; i++ {
			SimpleFlatFormZogSchema.Parse(data, &form)
		}
	})

	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			OzzoFlatForm(&filledForm)
		}
	})

	b.Run("Package Validator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SimpleFlatFormPkgValidator.Struct(&filledForm)
		}
	})

	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(&filledForm)
		}
	})

}

func BenchmarkFlatFormSchemaEmpty(b *testing.B) {
	emptyForm := SimpleFlatForm{}
	emptyData := map[string]any{}

	b.Run("Zog Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SimpleFlatFormZogSchema.Validate(&emptyForm)
		}
	})

	b.Run("ZogParse", func(b *testing.B) {
		var form SimpleFlatForm
		for i := 0; i < b.N; i++ {
			SimpleFlatFormZogSchema.Parse(emptyData, &form)
		}
	})

	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			OzzoFlatForm(&emptyForm)
		}
	})

	b.Run("Package Validator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SimpleFlatFormPkgValidator.Struct(&emptyForm)
		}
	})

	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(&emptyForm)
		}
	})

}
