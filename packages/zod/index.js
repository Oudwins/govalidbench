import b from "benny";
import {
  StringFieldSimpleSuccessVal,
  StringFieldSimpleFailureVal,
  StringFieldSimpleZog,
  SliceFieldSucessVal,
  SliceFieldFailureVal,
  SliceFieldZog,
  StructSingleFieldSuccessVal,
  StructSingleFieldFailureVal,
  StructSingleFieldZog,
  StructSimpleSuccessVal,
  StructSimpleFailureVal,
  StructSimpleZog,
  StructComplexSuccessVal,
  StructComplexFailureVal,
  StructComplexZog,
  LotsOfTestsSuccessVal,
  LotsOfTestsFailureVal,
  LotsOfTestsZog,
} from "../data.js";

await b.suite(
  "Zod Validation Benchmarks",

  // 1. String Field Simple
  b.add("StringFieldSimple Success", () => {
    const val = StringFieldSimpleFailureVal.slice();
    return () => {
      StringFieldSimpleZog.safeParse(val);
    };
  }),
  b.add("StringFieldSimple Failure", () => {
    const val = StringFieldSimpleFailureVal.slice();
    return () => {
      StringFieldSimpleZog.safeParse(val);
    };
  }),

  // 2. Slice Field
  b.add("SliceField Success", () => {
    const val = SliceFieldFailureVal.slice();
    return () => {
      SliceFieldZog.safeParse(val);
    };
  }),
  b.add("SliceField Failure", () => {
    const val = SliceFieldFailureVal.slice();
    return () => {
      SliceFieldZog.safeParse(val);
    };
  }),

  // 3. Struct Single Field
  b.add("StructSingleField Success", () => {
    const val = { ...StructSingleFieldFailureVal };
    return () => {
      StructSingleFieldZog.safeParse(val);
    };
  }),
  b.add("StructSingleField Failure", () => {
    const val = { ...StructSingleFieldFailureVal };
    return () => {
      StructSingleFieldZog.safeParse(val);
    };
  }),

  // 4. Struct Simple
  b.add("StructSimple Success", () => {
    const val = { ...StructSimpleFailureVal };
    return () => {
      StructSimpleZog.safeParse(val);
    };
  }),
  b.add("StructSimple Failure", () => {
    const val = { ...StructSimpleFailureVal };
    return () => {
      StructSimpleZog.safeParse(val);
    };
  }),

  // 5. Struct Complex
  b.add("StructComplex Success", () => {
    const val = { ...StructComplexFailureVal };
    return () => {
      StructComplexZog.safeParse(val);
    };
  }),
  b.add("StructComplex Failure", () => {
    const val = { ...StructComplexFailureVal };
    return () => {
      StructComplexZog.safeParse(val);
    };
  }),

  // 6. Lots of Tests
  b.add("LotsOfTests Success", () => {
    const val = LotsOfTestsFailureVal.slice();
    return () => {
      LotsOfTestsZog.safeParse(val);
    };
  }),
  b.add("LotsOfTests Failure", () => {
    const val = LotsOfTestsFailureVal.slice();
    return () => {
      LotsOfTestsZog.safeParse(val);
    };
  }),

  b.cycle(),
  b.complete(),
  b.save({ file: "zod", version: "^3.24.4" }),
  b.save({ file: "zod", format: "chart.html" })
);
