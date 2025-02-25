
# Golang Validation Libraries Benchmarks

This repo aims to benchmark the performance of popular validation libraries in Golang. The repo is still quite barebones but we will add more tests and libraries as we go.

## Libraries Tested:

- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- [go-playground/validator](https://github.com/go-playground/validator)
- [asaskevich/govalidator](https://github.com/asaskevich/govalidator)
- [Oudwins/zog](https://github.com/Oudwins/zog)


## What is my thinking in terms of benchmarks I should do?

For each one that fails and one that passes + parallel.

1. One simple test for simple primitive -> only one validation
2. One complex tests for simple primitive -> many validations
3. One simple test per complex type -> one validation for each
4. One complex test per complex type -> many validations
5. Some additional tests with nested types