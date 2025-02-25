package govalidator

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	"github.com/asaskevich/govalidator"
)

//
// 1. String Field
//

func BenchmarkStringFieldSimpleSuccessBench(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleSuccessVal)
		}
	})
}

func BenchmarkStringFieldSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStringFieldSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleFailureVal)
		}
	})
}

func BenchmarkStringFieldSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleFailureVal)
			}
		})
	})
}

//
// 2. Slice Field
//

func BenchmarkSliceFieldSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldGoValidator(&packages.SliceFieldSucessVal)
		}
	})
}

func BenchmarkSliceFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldGoValidator(&packages.SliceFieldSucessVal)
			}
		})
	})
}

func BenchmarkSliceFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldGoValidator(&packages.SliceFieldFailureVal)
		}
	})
}

func BenchmarkSliceFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldGoValidator(&packages.SliceFieldFailureVal)
			}
		})
	})
}

//
// 3. StructSingleField
//

func BenchmarkStructSingleFieldSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		}
	})
}

func BenchmarkStructSingleFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
			}
		})
	})
}

func BenchmarkStructSingleFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		}
	})
}

func BenchmarkStructSingleFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
			}
		})
	})
}

//
// 4. Struct Simple
//

func BenchmarkStructSimpleSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSimpleSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSimpleSuccessVal)
		}
	})
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSimpleSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSimpleFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSimpleFailureVal)
		}
	})
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSimpleFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSimpleFailureVal)
			}
		})
	})
}

//
// 5. Struct Complex
//

func BenchmarkStructComplexSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		}
	})
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructComplexSuccessVal)
			}
		})
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructComplexFailureVal)
		}
	})
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructComplexFailureVal)
			}
		})
	})
}

//
// 6. Lots of Tests
//

func BenchmarkLotsOfTestsSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsGoValidator(&packages.LotsOfTestsSuccessVal)
		}
	})
}

func BenchmarkLotsOfTestsSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsGoValidator(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
}

func BenchmarkLotsOfTestsFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsGoValidator(&packages.LotsOfTestsFailureVal)
		}
	})
}

func BenchmarkLotsOfTestsFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsGoValidator(&packages.LotsOfTestsFailureVal)
			}
		})
	})
}
