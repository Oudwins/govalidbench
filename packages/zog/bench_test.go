package zog

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

//
// 1. String Field
//

func BenchmarkStringFieldSimpleSuccessBench(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.StopTimer()
		internals.Clear()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleSuccessVal)
		}
	})
}

func BenchmarkStringFieldSimpleSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStringFieldSimpleFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.StopTimer()
		internals.Clear()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			errs := packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleFailureVal)
			z.Issues.CollectList(errs)
		}
	})
}

func BenchmarkStringFieldSimpleFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})
}

//
// 2. Slice Field
//

func BenchmarkSliceFieldSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldZog.Validate(&packages.SliceFieldSucessVal)
		}
	})
}

func BenchmarkSliceFieldSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldZog.Validate(&packages.SliceFieldSucessVal)
			}
		})
	})
}

func BenchmarkSliceFieldFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := packages.SliceFieldZog.Validate(&packages.SliceFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
}

func BenchmarkSliceFieldFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.SliceFieldZog.Validate(&packages.SliceFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}

//
// 3. StructSingleField
//

func BenchmarkStructSingleFieldSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldSuccessVal)
		}
	})
}

func BenchmarkStructSingleFieldSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
}

func BenchmarkStructSingleFieldFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
}

func BenchmarkStructSingleFieldFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}

//
// 4. Struct Simple
//

func BenchmarkStructSimpleSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSimpleZog.Validate(&packages.StructSimpleSuccessVal)
		}
	})
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSimpleZog.Validate(&packages.StructSimpleSuccessVal)
			}
		})
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := packages.StructSimpleZog.Validate(&packages.StructSimpleFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StructSimpleZog.Validate(&packages.StructSimpleFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}

//
// 5. Struct Complex
//

func BenchmarkStructComplexSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructComplexZog.Validate(&packages.StructComplexSuccessVal)
		}
	})
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexZog.Validate(&packages.StructComplexSuccessVal)
			}
		})
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := packages.StructComplexZog.Validate(&packages.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StructComplexZog.Validate(&packages.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}

//
// 6. Lots of Tests
//

func BenchmarkLotsOfTestsSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsSuccessVal)
		}
	})
}

func BenchmarkLotsOfTestsSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
}

func BenchmarkLotsOfTestsFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsFailureVal)
			z.Issues.CollectList(errs)
		}
	})
}

func BenchmarkLotsOfTestsFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})
}

//
// 7. Struct Complex When Generating the Schema
//

func BenchmarkStructComplexCreateSuccess(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := packages.StructComplexCreateZog()
			s.Validate(&packages.StructComplexSuccessVal)
		}
	})
}

func BenchmarkStructComplexCreateSuccessParallel(b *testing.B) {
	internals.Clear()
	b.Run("Success", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := packages.StructComplexCreateZog()
				s.Validate(&packages.StructComplexSuccessVal)
			}
		})
	})
}

func BenchmarkStructComplexCreateFailure(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := packages.StructComplexCreateZog()
			errs := s.Validate(&packages.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
}

func BenchmarkStructComplexCreateFailureParallel(b *testing.B) {
	internals.Clear()
	b.Run("Error", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := packages.StructComplexCreateZog()
				errs := s.Validate(&packages.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}
