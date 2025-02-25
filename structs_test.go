package main

import (
	"testing"

	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
)

type SingleField struct {
	Field string `validate:"min=3,max=10" valid:"stringlength(3|10)"`
}

func BenchmarkStructSingleFieldSuccess(b *testing.B) {

	// Initial
	s := SingleField{
		Field: "test",
	}

	b.SetParallelism(1) // Ensures no parallel execution

	// ZOG
	zogParsed := SingleField{}
	zogData := map[string]interface{}{
		"Field": "test",
	}
	zogSchema :=
		z.Struct(
			z.Schema{
				"field": z.String().Min(3).Max(10),
			},
		)
	b.Run("Zog Validate", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			zogSchema.Validate(&s)
		}
	})

	b.Run("Zog Parse", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			zogSchema.Parse(zogData, &zogParsed)
		}
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ozzo.ValidateStruct(&s,
				ozzo.Field(&s.Field,
					ozzo.Length(3, 10),
				),
			)
		}
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pkgValidator.Struct(&s)
		}
	})
	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(&s)
		}
	})
}
