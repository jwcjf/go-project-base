package database

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var policies = map[string]dbresolver.Policy{
	"random": dbresolver.RandomPolicy{},
}

// Configure ...
type Configure interface {
	Init(*gorm.Config, func(string) gorm.Dialector) (*gorm.DB, error)
}

// ResolverConfigure ...
type ResolverConfigure interface {
	Init(*dbresolver.DBResolver, func(string) gorm.Dialector) *dbresolver.DBResolver
}
