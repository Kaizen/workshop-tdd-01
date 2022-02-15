package main

//go:generate moq -out number_service_mock.go . NumberService
type NumberService interface {
	Generate() int
}
