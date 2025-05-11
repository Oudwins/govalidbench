import { z } from "zod";

// 1. StringFieldSimple
export const StringFieldSimpleSuccessVal = "test";
export const StringFieldSimpleFailureVal = "t";
export const StringFieldSimpleZog = z.string().min(3).max(10);

// 2. SliceField
export const SliceFieldSucessVal = [
  "val1",
  "val2",
  "val3",
  "val4",
  "val5",
  "val6",
  "val7",
  "val8",
  "val9",
  "val10",
];
export const SliceFieldFailureVal = [
  "1",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
  "valid10",
];
export const SliceFieldZog = z.array(z.string().min(2).max(10));

// 3. StructSingleField
export const StructSingleFieldSuccessVal = {
  field: "test",
};
export const StructSingleFieldFailureVal = {
  field: "t",
};
export const StructSingleFieldZog = z.object({
  field: z.string().min(3).max(10),
});

// 4. StructSimple
export const StructSimpleSuccessVal = {
  string: "good val",
  int: 6,
};
export const StructSimpleFailureVal = {
  string: "bad",
  int: 1,
};
export const StructSimpleZog = z.object({
  string: z.string().min(3).max(10),
  int: z.number().gte(3).lte(10),
});

// 5. StructComplex
export const StructComplexSuccessVal = {
  blankTag: "blank",
  len: "0123456789",
  min: ">=1",
  max: "<= 10",
  minMax: "1 <= 10",
  email: "zog@gmail.com",
  url: "https://zog.dev/",
  int: 5,
  color: "red",
  sub: {
    test: "123456",
  },
  subIgnore: {},
  anonymous: {
    a: "1234567",
  },
};

export const StructComplexFailureVal = {
  blankTag: "blank",
  len: "1",
  min: "",
  max: "<= 10 -> so more than 10",
  minMax: "1 <= 10 -> so more than 10",
  email: "zog@.com",
  url: "Not url",
  int: 0,
  sub: {
    test: "0",
  },
  subIgnore: {},
  anonymous: {
    a: "0",
  },
};

export const StructComplexZog = z.object({
  len: z.string().length(10),
  min: z.string().min(1),
  max: z.string().max(10),
  minMax: z.string().min(1).max(10),
  email: z.string().email(),
  url: z.string().url(),
  int: z.number().gte(3).lte(10),
  color: z.enum(["red", "blue"]),
  sub: z
    .object({
      test: z.string().min(5).max(10),
    })
    .nullable(),
  anonymous: z.object({
    a: z.string().min(5).max(10),
  }),
});

// 6. LotsOfTests
export const LotsOfTestsSuccessVal = "1234567890111";
export const LotsOfTestsFailureVal = "";
export const LotsOfTestsZog = z
  .string()
  .min(1)
  .min(2)
  .min(3)
  .min(4)
  .min(5)
  .min(6)
  .min(7)
  .min(8)
  .min(9)
  .min(10);
