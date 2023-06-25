package decorator

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/cache"
)

type CacheCustomerDecorator struct {
	cache cache.Cache
	repo  repository.CustomerRepository
}
