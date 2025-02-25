package main

import (
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/asaskevich/govalidator"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
)

func BenchmarkSliceFieldSuccess(b *testing.B) {
	s := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}

	// ZOG
	zogParsed := []string{}
	zogSchema := z.Slice(z.String().Min(2).Max(10).Required())
	b.Run("Zog Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zogSchema.Validate(&s)
		}
	})

	b.Run("Zog Parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zogSchema.Parse(s, &zogParsed)
		}
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range s {
				ozzo.Validate(v, ozzo.Length(2, 10))
			}
		}
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = pkgValidator.Var(s, "dive,min=2,max=10")
		}
	})
	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range s {
				govalidator.MinStringLength(v, "3")
				govalidator.MaxStringLength(v, "10")
			}
		}
	})
}

func BenchmarkSliceFieldSuccessParallel(b *testing.B) {
	s := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}

	// ZOG
	zogParsed := []string{}
	zogSchema := z.Slice(z.String().Min(2).Max(10).Required())
	b.Run("Zog Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Validate(&s)
			}
		})
	})

	b.Run("Zog Parse", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Parse(s, &zogParsed)
			}
		})
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for _, v := range s {
					ozzo.Validate(v, ozzo.Length(2, 10))
				}
			}
		})
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = pkgValidator.Var(s, "dive,min=2,max=10")
			}
		})
	})
	b.Run("Govalidator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for _, v := range s {
					govalidator.MinStringLength(v, "3")
					govalidator.MaxStringLength(v, "10")
				}
			}
		})
	})
}

func BenchmarkSliceFieldFailure(b *testing.B) {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "valid10"}

	// ZOG
	zogParsed := []string{}
	zogSchema := z.Slice(z.String().Min(2).Max(10).Required())
	b.Run("Zog Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zogSchema.Validate(&s)
		}
	})

	b.Run("Zog Parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zogSchema.Parse(s, &zogParsed)
		}
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range s {
				ozzo.Validate(v, ozzo.Length(2, 10))
			}
		}
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = pkgValidator.Var(s, "dive,min=2,max=10")
		}
	})
	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range s {
				govalidator.MinStringLength(v, "3")
				govalidator.MaxStringLength(v, "10")
			}
		}
	})
}

func BenchmarkSliceFieldFailureParallel(b *testing.B) {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "valid10"}

	// ZOG
	zogParsed := []string{}
	zogSchema := z.Slice(z.String().Min(2).Max(10).Required())
	b.Run("Zog Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Validate(&s)
			}
		})
	})

	b.Run("Zog Parse", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Parse(s, &zogParsed)
			}
		})
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for _, v := range s {
					ozzo.Validate(v, ozzo.Length(2, 10))
				}
			}
		})
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = pkgValidator.Var(s, "dive,min=2,max=10")
			}
		})
	})
	b.Run("Govalidator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for _, v := range s {
					govalidator.MinStringLength(v, "3")
					govalidator.MaxStringLength(v, "10")
				}
			}
		})
	})
}
