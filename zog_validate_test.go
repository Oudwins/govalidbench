package main

import (
	"regexp"
	"testing"

	z "github.com/Oudwins/zog"
)

type FlatForm struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
	Website   string `json:"website"`
	CreatedAt string `json:"created_at"`
}

var cities = []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}

var phoneRegex = regexp.MustCompile(`^\(\d{3}\) \d{3}-\d{4}$`)

var FlatFormZogSchema = z.Struct(
	z.Schema{
		"id":        z.String().Required().UUID(),
		"name":      z.String().Required(),
		"email":     z.String().Required().Email(),
		"age":       z.Int().Required().GTE(18).LTE(100),
		"phone":     z.String().Required().Match(phoneRegex),
		"address":   z.String().Required(),
		"city":      z.String().Required().OneOf(cities),
		"state":     z.String().Required(),
		"zipCode":   z.String().Required(),
		"country":   z.String().Required(),
		"website":   z.String().Required().URL(),
		"createdAt": z.String().Required(),
	},
)

func BenchmarkFlatFormSchemaFilled(b *testing.B) {
	filledForm := FlatForm{
		Id:        "123e4567-e89b-12d3-a456-426614174000",
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Age:       25,
		Phone:     "(123) 456-7890",
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
		"phone":     "(123) 456-7890",
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
			FlatFormZogSchema.Validate(&filledForm)
		}
	})

	b.Run("Zog Parse", func(b *testing.B) {
		var form FlatForm
		for i := 0; i < b.N; i++ {
			FlatFormZogSchema.Parse(data, &form)
		}
	})
}

func BenchmarkFlatFormSchemaEmpty(b *testing.B) {
	emptyForm := FlatForm{}
	emptyData := map[string]any{}

	b.Run("Zog Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FlatFormZogSchema.Validate(&emptyForm)
		}
	})

	b.Run("ZogParse", func(b *testing.B) {
		var form FlatForm
		for i := 0; i < b.N; i++ {
			FlatFormZogSchema.Parse(emptyData, &form)
		}
	})
}
