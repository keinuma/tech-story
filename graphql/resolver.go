package graphql

import (
	"gorm.io/gorm"

	"github.com/keinuma/tech-story/infra/store"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB        *gorm.DB
	StorePool *store.Store
}
