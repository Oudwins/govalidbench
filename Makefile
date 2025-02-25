bench:
	go test -bench=. -benchmem -run=none

bench-zog:
	@echo "======"
	@echo "Benchmarking zog..."
	@echo "======"
	go test ./packages/zog -bench=. -benchmem -run=none

bench-validator:
	@echo "======"
	@echo "Benchmarking validator..."
	@echo "======"
	go test ./packages/validator -bench=. -benchmem -run=none

bench-ozzo:
	@echo "======"
	@echo "Benchmarking ozzo..."
	@echo "======"
	go test ./packages/ozzo -bench=. -benchmem -run=none

bench-govalidator:
	@echo "======"
	@echo "Benchmarking govalidator..."
	@echo "======"
	go test ./packages/govalidator -bench=. -benchmem -run=none

bench-all: bench-zog bench-validator bench-ozzo bench-govalidator
	@echo "======"
	@echo "All benchmarks completed"
	@echo "======"
