package service

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"

type DomainService struct {
	MatcherService MatcherService
}

func NewDomainService(customerRepo repository.CustomerRepository) *DomainService {
	return &DomainService{
		MatcherService: NewMatcherService(customerRepo),
	}
}
