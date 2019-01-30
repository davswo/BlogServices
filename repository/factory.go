package repository

import (
	"github.com/pkg/errors"
)

const (
	MemoryDatabase = "memory"
)

func Create(dbtype string) (BlogRepository, error) {
	switch dbtype {
	case MemoryDatabase:
		return NewOrderRepositoryMemory(), nil
	default:
		return nil, errors.Errorf("Unsupported database type %s", dbtype)
	}
}
