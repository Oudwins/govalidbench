package main

import (
	"testing"

	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/go-playground/validator/v10"

	"github.com/asaskevich/govalidator"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

func BenchmarkStringFieldSuccess(b *testing.B) {
	// Initial
	s := "test"

	b.SetParallelism(1) // Ensures no parallel execution

	// ZOG
	zogParsed := ""
	zogSchema := z.String().Min(3).Max(10)
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
			zogSchema.Parse(s, &zogParsed)
		}
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ozzo.Validate(&s, ozzo.Length(3, 10))
		}
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pkgValidator.Var(s, "min=3,max=10")
		}
	})
	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			govalidator.MinStringLength(s, "3")
			govalidator.MaxStringLength(s, "10")
		}
	})
}

func BenchmarkStringFieldSuccessParallel(b *testing.B) {
	// Initial
	s := "test"
	b.SetParallelism(1) // Ensures no parallel execution
	// ZOG
	zogParsed := ""
	zogSchema := z.String().Min(3).Max(10)
	b.Run("Zog Validate", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Validate(&s)
			}
		})
	})

	b.Run("Zog Parse", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
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
				ozzo.Validate(&s, ozzo.Length(3, 10))
			}
		})
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				pkgValidator.Var(s, "min=3,max=10")
			}
		})
	})
	b.Run("Govalidator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.MinStringLength(s, "3")
				govalidator.MaxStringLength(s, "10")
			}
		})
	})
}

func BenchmarkStringFieldFailure(b *testing.B) {
	// Initial
	s := "2"
	b.SetParallelism(1) // Ensures no parallel execution

	// ZOG
	zogParsed := ""
	zogSchema := z.String().Min(3).Max(10)
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
			zogSchema.Parse(s, &zogParsed)
		}
	})

	// OZZO
	b.Run("Ozzo Validate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ozzo.Validate(&s, ozzo.Length(3, 10))
		}
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pkgValidator.Var(s, "min=3,max=10")
		}
	})
	b.Run("Govalidator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			govalidator.MinStringLength(s, "3")
			govalidator.MaxStringLength(s, "10")
		}
	})
}

func BenchmarkStringFieldFailureParallel(b *testing.B) {
	// Initial
	s := "2"
	b.SetParallelism(1) // Ensures no parallel execution

	// ZOG
	zogParsed := ""
	zogSchema := z.String().Min(3).Max(10)
	b.Run("Zog Validate", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				zogSchema.Validate(&s)
			}
		})
	})

	b.Run("Zog Parse", func(b *testing.B) {
		b.StopTimer()
		internals.ClearPools()
		b.StartTimer()
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
				ozzo.Validate(&s, ozzo.Length(3, 10))
			}
		})
	})

	pkgValidator := validator.New()
	b.Run("Package Validator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				pkgValidator.Var(s, "min=3,max=10")
			}
		})
	})
	b.Run("Govalidator", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.MinStringLength(s, "3")
				govalidator.MaxStringLength(s, "10")
			}
		})
	})
}
