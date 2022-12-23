package db

import (
	"github.com/karanbirsingh7/go-grpc/internal/rocket"
)

// Migrate - takes care of migration DB schema
func (s *Store) Migrate(i ...interface{}) error {
	if err := s.db.AutoMigrate(rocket.Rocket{}); err != nil {
		return err
	}
	return nil
}
