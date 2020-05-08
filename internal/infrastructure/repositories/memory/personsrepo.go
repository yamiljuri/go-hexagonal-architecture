package memory

import (
	"github.com/yamiljuri/hex_arq/internal/core/domain"
	"sync"
)

var (
	oncePersonsRepo sync.Once
	instancePersonsRepo *repository
)

type repository struct {
	dataMemory map[int]domain.Person
	lastId int
}


func NewPersonsRepository() *repository{
	oncePersonsRepo.Do(func() {
		instancePersonsRepo = &repository{
			dataMemory: map[int]domain.Person{},
			lastId: 0,
		}
	})
	return instancePersonsRepo
}

func (r *repository) Get(id int) (domain.Person,error){
	var person domain.Person
	person = r.dataMemory[id]
	return person, nil
}

func (r *repository) Save(person domain.Person) error{
	r.dataMemory[r.lastId] = person
	r.lastId++
	return nil
}