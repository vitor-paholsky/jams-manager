package server

import "net/http"

type (
	Category struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	Categories []Category

	Rider struct {
		Id               string   `json:"id"`
		Name             string   `json:"name"`
		Age              int      `json:"age"`
		Gender           string   `json:"gender"`
		City             string   `json:"city"`
		Cpf              string   `json:"cpf"`
		PaidSubscription bool     `json:"paidSubscription"`
		Sponsors         []string `json:"sponsors"`
		CategoryId       string   `json:"categoryId"`
	}
	Riders []Rider
)

type CategoryHandler interface {
	Handle() http.HandlerFunc
}

type RiderHandler interface {
	Handle() http.HandlerFunc
}
