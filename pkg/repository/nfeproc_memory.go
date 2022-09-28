package repository

import (
	"gitlab.com/waldirborbajr/nfeloader/pkg/entity"
)

type NFeProcsMemoryDB struct {
	NFeProcs []entity.NFeProc
}

type NFeProcRepositoryMemory struct {
	db NFeProcsMemoryDB
}

func NewNFeProcRepositoryMemory(db NFeProcsMemoryDB) *NFeProcRepositoryMemory {
	return &NFeProcRepositoryMemory{db: db}
}

func (n *NFeProcRepositoryMemory) SaveNFe(nfeProc entity.NFeProc) (entity.NFeProc, error) {
	n.db.NFeProcs = append(n.db.NFeProcs, nfeProc)
	return nfeProc, nil
}
