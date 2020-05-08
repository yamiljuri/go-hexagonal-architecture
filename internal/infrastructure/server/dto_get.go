package server

import "github.com/yamiljuri/hex_arq/internal/core/domain"

type ResponsePersonGet domain.Person

func BuildResponsePersonGet(person domain.Person) ResponsePersonGet{
	return ResponsePersonGet(person)
}

