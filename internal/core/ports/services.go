package ports

import "github.com/yamiljuri/hex_arq/internal/core/domain"

type PersonsService interface {
	Get(id int) (domain.Person,error)
	Create(person domain.Person) error
}

