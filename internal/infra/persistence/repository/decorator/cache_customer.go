package decorator

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/cache"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type CacheCustomerDecorator struct {
	cache cache.Cache
	repo  repository.CustomerRepository
}
