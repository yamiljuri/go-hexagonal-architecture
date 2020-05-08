package ports

import "github.com/yamiljuri/hex_arq/internal/core/domain"

type PersonsRepository interface {
	Get(id int) (domain.Person, error)
	Save(person domain.Person) error
}

