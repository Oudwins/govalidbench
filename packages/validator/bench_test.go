package validator

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	"github.com/go-playground/validator/v10"
)

var playgroundValidate = validator.New()

//
// 1. String Field
//

func BenchmarkStringFieldSimpleSuccessBench(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.StopTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.StringFieldSimpleSuccessVal, packages.StringFieldSimpleValidator)
		}
	})
}

func BenchmarkStringFieldSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.StringFieldSimpleSuccessVal, packages.StringFieldSimpleValidator)
			}
		})
	})
}

func BenchmarkStringFieldSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.StopTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.StringFieldSimpleFailureVal, packages.StringFieldSimpleValidator)
		}
	})
}

func BenchmarkStringFieldSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.StringFieldSimpleFailureVal, packages.StringFieldSimpleValidator)
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
			playgroundValidate.Var(packages.SliceFieldSucessVal, packages.SliceFieldValidator)
		}
	})
}

func BenchmarkSliceFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.SliceFieldSucessVal, packages.SliceFieldValidator)
			}
		})
	})
}

func BenchmarkSliceFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.SliceFieldFailureVal, packages.SliceFieldValidator)
		}
	})
}

func BenchmarkSliceFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.SliceFieldFailureVal, packages.SliceFieldValidator)
			}
		})
	})
}

//
// 3. StructSingleField
//

func BenchmarkStructSingleFieldSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSingleFieldSuccessVal)
		}
	})
}

func BenchmarkStructSingleFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
}

func BenchmarkStructSingleFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSingleFieldFailureVal)
		}
	})
}

func BenchmarkStructSingleFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSingleFieldFailureVal)
			}
		})
	})
}

//
// 4. Struct Simple
//

func BenchmarkStructSimpleSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSimpleSuccessVal)
		}
	})
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSimpleFailureVal)
		}
	})
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSimpleFailureVal)
			}
		})
	})
}

//
// 5. Struct Complex
//

func BenchmarkStructComplexSuccess(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructComplexSuccessVal)
		}
	})
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructComplexSuccessVal)
			}
		})
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructComplexFailureVal)
		}
	})
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructComplexFailureVal)
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
			playgroundValidate.Var(packages.LotsOfTestsSuccessVal, packages.LotsOfTestsValidator)
		}
	})
}

func BenchmarkLotsOfTestsSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.LotsOfTestsSuccessVal, packages.LotsOfTestsValidator)
			}
		})
	})
}

func BenchmarkLotsOfTestsFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.LotsOfTestsFailureVal, packages.LotsOfTestsValidator)
		}
	})
}

func BenchmarkLotsOfTestsFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.LotsOfTestsFailureVal, packages.LotsOfTestsValidator)
			}
		})
	})
}
