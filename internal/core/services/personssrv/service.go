package personssrv

import (
	"github.com/yamiljuri/hex_arq/internal/core/domain"
	"github.com/yamiljuri/hex_arq/internal/core/ports"
)

type service struct {
	repo ports.PersonsRepository
}

func NewService(repo ports.PersonsRepository) * service{
	return &service{
		repo:repo,
	}
}

func (s *service) Get(id int) (domain.Person,error){
	return s.repo.Get(id)
}

func (s * service) Create(person domain.Person) error {
	return s.repo.Save(person)
}




