install:
	@go run cmd/assets.go
	@go install cmd/godin/godin.go cmd/godin/add.go cmd/godin/generate.go cmd/godin/init.go cmd/godin/template.go