package db

import (
	"github.com/karanbirsingh7/go-grpc/internal/rocket"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB // only this package can access this field
}

// New - returns a new store or error
func New() (Store, error) {
	db, err := gorm.Open(sqlite.Open("__deleteme.db"))

	// db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	if err != nil {
		return Store{}, err
	}

	return Store{db: db}, nil
}

func (s *Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}
func (s *Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}
func (s *Store) DeleteRocket(id string) error {
	return nil
}
func (s *Store) GetRockets() ([]rocket.Rocket, error) {
	return []rocket.Rocket{}, nil
}
