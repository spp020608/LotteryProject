package services

import (
	"LotteryProject/dao"
	"LotteryProject/datasource"
	"LotteryProject/models"
)

type ResultService interface {
	GetAll() []models.LtResult
	CountAll() int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

func (s *resultService) GetAll() []models.LtResult {
	return s.dao.GetAll()
}

func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *resultService) Get(id int) *models.LtResult {
	return s.dao.Get(id)
}

func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *resultService) Create(data *models.LtResult) error {
	return s.dao.Create(data)
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(datasource.InstanceDbMaster()),
	}
}
