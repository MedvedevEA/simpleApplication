package service

import (
	"fastApplication/internal/logger"
	"fastApplication/internal/model"
)

type Service struct {
	repository Repository
	logger     logger.Logger
}
type Repository interface {
	ExecuteQuery(path string, req *model.Params) (*any, error)
	GetListQueries(schemaName string) ([]*model.Query, error)
	SetQueryRoutes(queryRoutes map[string]string)
}

func New(repository Repository, logger logger.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}
func (s *Service) ExecuteQuery(path string, req *model.Params) (*any, error) {
	return s.repository.ExecuteQuery(path, req)
}
func (s *Service) GetListQueries(schemaName string) ([]*model.Query, error) {
	return s.repository.GetListQueries(schemaName)
}
func (s *Service) SetQueryRoutes(queryRoutes map[string]string) {
	s.repository.SetQueryRoutes(queryRoutes)
}
