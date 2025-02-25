package ozzo

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
)

//
// 1. String Field
//

func BenchmarkStringFieldSimpleSuccessBench(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleSuccessVal)
		}
	})
}

func BenchmarkStringFieldSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStringFieldSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleFailureVal)
		}
	})
}

func BenchmarkStringFieldSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleFailureVal)
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
			packages.SliceFieldOzzo(&packages.SliceFieldSucessVal)
		}
	})
}

func BenchmarkSliceFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldOzzo(&packages.SliceFieldSucessVal)
			}
		})
	})
}

func BenchmarkSliceFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldOzzo(&packages.SliceFieldFailureVal)
		}
	})
}

func BenchmarkSliceFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldOzzo(&packages.SliceFieldFailureVal)
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
			packages.StructSingleFieldOzzo(&packages.StructSingleFieldSuccessVal)
		}
	})
}

func BenchmarkStructSingleFieldSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldOzzo(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
}

func BenchmarkStructSingleFieldFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSingleFieldOzzo(&packages.StructSingleFieldFailureVal)
		}
	})
}

func BenchmarkStructSingleFieldFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldOzzo(&packages.StructSingleFieldFailureVal)
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
			packages.StructSimpleOzzo(&packages.StructSimpleSuccessVal)
		}
	})
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSimpleOzzo(&packages.StructSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSimpleOzzo(&packages.StructSimpleFailureVal)
		}
	})
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSimpleOzzo(&packages.StructSimpleFailureVal)
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
			packages.StructComplexOzzo(&packages.StructComplexSuccessVal)
		}
	})
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexOzzo(&packages.StructComplexSuccessVal)
			}
		})
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructComplexOzzo(&packages.StructComplexFailureVal)
		}
	})
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexOzzo(&packages.StructComplexFailureVal)
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
			packages.LotsOfTestsOzzo(&packages.LotsOfTestsSuccessVal)
		}
	})
}

func BenchmarkLotsOfTestsSuccessParallel(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsOzzo(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
}

func BenchmarkLotsOfTestsFailure(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsOzzo(&packages.LotsOfTestsFailureVal)
		}
	})
}

func BenchmarkLotsOfTestsFailureParallel(b *testing.B) {
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsOzzo(&packages.LotsOfTestsFailureVal)
			}
		})
	})
}
