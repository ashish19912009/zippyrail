package main

type Account struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Order []Order `json:"order"`
}
