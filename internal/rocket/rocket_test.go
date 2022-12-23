package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get rocket by ID", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		id := "UUID-1"
		// signifes that when GetRocketByID is called with "UUID-1",
		// it should return NO ERROR and UUID-1
		rocketStoreMock.
			EXPECT().GetRocketByID(id).Return(&Rocket{
			ID: id,
		}, nil)

		// rocketService that wraps around Store interface
		// you call DB using methods that are linked to "Service"
		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(context.Background(), id)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID, "rocket ID doesn't match")
	})

	t.Run("insert a new rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		newRocket := Rocket{
			ID:      "UUID-1",
			Name:    "Falcon",
			Type:    "G1",
			Flights: 25,
		}

		// mocking our insert method call
		rocketStoreMock.
			EXPECT().
			InsertRocket(newRocket).
			Return(&newRocket, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(context.Background(), newRocket)

		assert.NoError(t, err)
		assert.Equal(t, newRocket, *rkt) // deference pointer for comparison
		// assert.Equal(t, newRocket,, rkt, "inserted rocket doesn't match payload rocket")
	})
}
