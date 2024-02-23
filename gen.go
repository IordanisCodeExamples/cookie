package gen

//go:generate mockgen -package inputmock -destination internal/mocks/service/service.go -source=internal/service/service.go Input
