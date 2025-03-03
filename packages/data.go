package packages

import (
	z "github.com/Oudwins/zog"
	"github.com/asaskevich/govalidator"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	ozzoIs "github.com/go-ozzo/ozzo-validation/v4/is"
)

// 1. StringFieldSimple*
//
// Values
var StringFieldSimpleSuccessVal = "test"
var StringFieldSimpleFailureVal = "t"

// Zog
var StringFieldSimpleZog = z.String().Min(3).Max(10)

// Ozzo
var StringFieldSimpleOzzo = func(s *string) {
	ozzo.Validate(s, ozzo.Length(3, 10))
}

// Validator:
var StringFieldSimpleValidator = "min=3,max=10"

// govalidator
var StringFieldSimpleGoValidator = func(s *string) {
	govalidator.MinStringLength(*s, "3")
	govalidator.MaxStringLength(*s, "10")
}

//
// 2. SliceField
//

// Values
var SliceFieldSucessVal = []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}
var SliceFieldFailureVal = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "valid10"}

// Zog
var SliceFieldZog = z.Slice(z.String().Min(2).Max(10))

// validator
var SliceFieldValidator = "dive,min=2,max=10"

// Ozzo
var SliceFieldOzzo = func(s *[]string) {
	for _, v := range *s {
		ozzo.Validate(v, ozzo.Length(2, 10))
	}
}

// govalidator
var SliceFieldGoValidator = func(s *[]string) {
	for _, v := range *s {
		govalidator.MinStringLength(v, "2")
		govalidator.MaxStringLength(v, "10")
	}
}

//
// 3. StructSingleField
//

// Values

type sSingleFieldType struct {
	// This contains Validator & Govalidator tags
	Field string `validate:"min=3,max=10" valid:"stringlength(3|10)"`
}

var StructSingleFieldSuccessVal = sSingleFieldType{
	Field: "test",
}

var StructSingleFieldFailureVal = sSingleFieldType{
	Field: "t",
}

// Zog

var StructSingleFieldZog = z.Struct(
	z.Schema{
		"field": z.String().Min(3).Max(10),
	},
)

// Ozzo
var StructSingleFieldOzzo = func(s *sSingleFieldType) {
	ozzo.ValidateStruct(s,
		ozzo.Field(&s.Field, ozzo.Length(3, 10)),
	)
}

//
// 4. StructSimple
//

// Values
type StructSimpleType struct {
	// contains Validator & Govalidator tags
	String string `validate:"min=5,max=10" valid:"stringlength(5|10)"`
	Int    int    `validate:"min=5,max=10" valid:"range(5|10)"`
}

var StructSimpleSuccessVal = StructSimpleType{
	String: "good val",
	Int:    6,
}

var StructSimpleFailureVal = StructSimpleType{
	String: "bad",
	Int:    1,
}

// Zog

var StructSimpleZog = z.Struct(z.Schema{
	"string": z.String().Min(3).Max(10),
	"int":    z.Int().GTE(3).LTE(10),
})

// Ozzo
var StructSimpleOzzo = func(s *StructSimpleType) {
	ozzo.ValidateStruct(s,
		ozzo.Field(&s.String, ozzo.Length(3, 10)),
		ozzo.Field(&s.Int, ozzo.Min(3), ozzo.Max(10)),
	)
}

//
// 5. StructComplex
//

type StructComplexNested struct {
	// contains Validator & Govalidator tags
	Test string `validate:"min=5,max=10" valid:"minstringlength(5),maxstringlength(10)"`
}

type StructComplexType struct {
	// contains Validator & Govalidator tags
	BlankTag  string               `validate:""`
	Len       string               `validate:"len=10" valid:"stringlength(10|10)"`
	Min       string               `validate:"min=1" valid:"minstringlength(1)"`
	Max       string               `validate:"max=10" valid:"maxstringlength(10)"`
	MinMax    string               `validate:"min=1,max=10" valid:"minstringlength(1),maxstringlength(10)"`
	Email     string               `validate:"email" valid:"email"`
	Url       string               `validate:"url" valid:"url"`
	Int       int                  `validate:"gte=3,lte=10" valid:"range(3|10)"`
	Color     string               `validate:"oneof=red green" valid:"in(red|green)"`
	Sub       *StructComplexNested `valid:"optional"`
	SubIgnore *StructComplexNested `validate:"-" valid:"-"`
	Anonymous struct {
		A string `validate:"min=5,max=10" valid:"stringlength(5|10)"`
	}
}

var StructComplexSuccessVal = StructComplexType{
	BlankTag: "blank",
	Len:      "0123456789",
	Min:      ">=1",
	Max:      "<= 10",
	MinMax:   "1 <= 10",
	Email:    "zog@gmail.com",
	Url:      "https://zog.dev/",
	Int:      5,
	Color:    "red",
	Sub: &StructComplexNested{
		Test: "123456",
	},
	SubIgnore: &StructComplexNested{},
	Anonymous: struct {
		A string `validate:"min=5,max=10" valid:"stringlength(5|10)"`
	}{
		A: "1234567",
	},
}

var StructComplexFailureVal = StructComplexType{
	BlankTag: "blank",
	Len:      "1",
	Min:      "",
	Max:      "<= 10 -> so more than 10",
	MinMax:   "1 <= 10 -> so more than 10",
	Email:    "zog@.com",
	Url:      "Not url",
	Int:      0,
	Sub: &StructComplexNested{
		Test: "0",
	},
	SubIgnore: &StructComplexNested{},
	Anonymous: struct {
		A string `validate:"min=5,max=10" valid:"stringlength(5|10)"`
	}{
		A: "0",
	},
}

func StructComplexCreateZog() *z.StructSchema {
	return z.Struct(
		z.Schema{
			"Len":    z.String().Len(10),
			"Min":    z.String().Min(1),
			"Max":    z.String().Max(10),
			"MinMax": z.String().Min(1).Max(10),
			"Email":  z.String().Email(),
			"Url":    z.String().URL(),
			"Int":    z.Int().GTE(3).LTE(10),
			"Color":  z.String().OneOf([]string{"red", "blue"}),
			"Sub": z.Ptr(z.Struct(z.Schema{
				"Test": z.String().Min(5).Max(10),
			})),
			"Anonymous": z.Struct(z.Schema{
				"A": z.String().Min(5).Max(10),
			}),
		},
	)

}

// Zog
var StructComplexZog = StructComplexCreateZog()

// Ozzo
var StructComplexOzzo = func(s *StructComplexType) {
	ozzo.ValidateStruct(s,
		ozzo.Field(&s.Len, ozzo.Length(10, 10)),
		ozzo.Field(&s.Min, ozzo.Length(1, 10000)),
		ozzo.Field(&s.Max, ozzo.Length(-10000, 10)),
		ozzo.Field(&s.MinMax, ozzo.Length(1, 10)),
		ozzo.Field(&s.Email, ozzoIs.Email),
		ozzo.Field(&s.Url, ozzoIs.URL),
		ozzo.Field(&s.Int, ozzo.Min(3), ozzo.Max(10)),
		ozzo.Field(&s.Color, ozzo.In("red", "green")),
		// Not sure if this is correct
		ozzo.Field(&s.Sub, ozzo.Required),
		ozzo.Field(&s.Sub.Test, ozzo.Length(5, 10)),
		ozzo.Field(&s.Anonymous, ozzo.Required),
		ozzo.Field(&s.Anonymous.A, ozzo.Length(5, 10)),
	)
}

//
// 6. LotsOfTests
//

var LotsOfTestsSuccessVal = "1234567890111"

var LotsOfTestsFailureVal = ""

var LotsOfTestsZog = z.String().Min(1).Min(2).Min(3).Min(4).Min(5).Min(6).Min(7).Min(8).Min(9).Min(10)

// Ozzo
var LotsOfTestsOzzo = func(s *string) {
	ozzo.Validate(s, ozzo.Length(1, 10), ozzo.Length(2, 10), ozzo.Length(3, 10), ozzo.Length(4, 10), ozzo.Length(5, 10))
}

// Validator:
var LotsOfTestsValidator = "min=1,min=2,min=3,min=4,min=5,min=6,min=7,min=8,min=9,min=10"

// govalidator
var LotsOfTestsGoValidator = func(s *string) {
	govalidator.MinStringLength(*s, "1")
	govalidator.MinStringLength(*s, "2")
	govalidator.MinStringLength(*s, "3")
	govalidator.MinStringLength(*s, "4")
	govalidator.MinStringLength(*s, "5")
	govalidator.MinStringLength(*s, "6")
	govalidator.MinStringLength(*s, "7")
	govalidator.MinStringLength(*s, "8")
	govalidator.MinStringLength(*s, "9")
	govalidator.MinStringLength(*s, "10")
}
