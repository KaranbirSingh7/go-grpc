//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/karanbirsingh7/go-grpc/internal/rocket Store
package rocket

import "context"

// Rocket - contains definitions of a rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketByID(id string) (*Rocket, error)
	InsertRocket(rkt Rocket) (*Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service responsible for updating rocket inventory
type Service struct {
	Store Store
}

// New - returns new instance of our rocket service
// Principle - accept interface return structs
func New(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetRocketByID - retreive a rocket based on ID provided from store
func (s *Service) GetRocketByID(ctx context.Context, id string) (*Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return &Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - responsible for inserting rocket into database
func (s *Service) InsertRocket(ctx context.Context, rkt Rocket) (*Rocket, error) {
	r, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return &Rocket{}, err

	}
	return r, nil
}

// swagger:router /delete/{id} rocket DeleteRocket
// responses:
//
//	200: Success
//
// DeleteRocket - delete an existing rocket
func (s *Service) DeleteRocket(ctx context.Context, id string) error {
	return s.Store.DeleteRocket(id)
}
