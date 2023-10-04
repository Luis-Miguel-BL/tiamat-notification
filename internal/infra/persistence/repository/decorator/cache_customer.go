package decorator

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/cache"
)

type CacheCustomerDecorator struct {
	cache cache.Cache
	repo  repository.CustomerRepository
}
