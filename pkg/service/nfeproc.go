package service

import (
	"github.com/waldirborbajr/nfeloader/pkg/entity"
)

type NFeProcService struct {
	Repository entity.NFeProcRepository
}

func NewNFeProcService(repository entity.NFeProcRepository) *NFeProcService {
	return &NFeProcService{Repository: repository}
}

func (n *NFeProcService) SaveNFe(nfeProc *entity.NFeProc) error {
	if err := n.Repository.SaveNFe(nfeProc); err != nil {
		return err
	}

	return nil
}

func (n *NFeProcService) DBPing() error {
	if err := n.Repository.DBPing(); err != nil {
		return err
	}
	return nil
}
