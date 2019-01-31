package repository

import (
	"github.com/pkg/errors"
)

const (
	MemoryDatabase = "memory"
	MongoDB        = "mongodb"
)

func Create(dbtype string) (BlogRepository, error) {
	switch dbtype {
	case MemoryDatabase:
		return NewOrderRepositoryMemory(), nil
	case MongoDB:
		return NewOrderRepositoryMongo()
	default:
		return nil, errors.Errorf("Unsupported database type %s", dbtype)
	}
}
