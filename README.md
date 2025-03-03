# Golang Validation Libraries Benchmarks

This repo aims to benchmark the performance of popular validation libraries in Golang. The repo is still quite barebones but we will add more tests and libraries as we go.

Benchmark results are stored in `/benchs/{library}/version/*.txt`

## Libraries Tested:

- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- [go-playground/validator](https://github.com/go-playground/validator)
- [asaskevich/govalidator](https://github.com/asaskevich/govalidator)
- [Oudwins/zog](https://github.com/Oudwins/zog)

## Project File Structure

```bash
.
├── packages
│   ├── ozzo
│   ├── validator
│   ├── govalidator
│   └── zog
```
