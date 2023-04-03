package handlers

import "github.com/mimani68/fintech-core/data/model"

type Handler struct {
	account model.Account
}

func NewHandler(account model.Account) *Handler {
	return &Handler{
		Account: account,
	}
}
