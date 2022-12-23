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
			EXPECT().GetRocketByID(id).Return(Rocket{
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
			Return(newRocket, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(context.Background(), newRocket)

		assert.NoError(t, err)
		assert.Equal(t, newRocket.ID, rkt.ID, "inserted rocket doesn't match payload rocket")
		assert.Equal(t, newRocket, rkt) // deference pointer for comparison
	})

	t.Run("list all rockets", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		rockets := []Rocket{
			{
				ID:   "1",
				Name: "Falcon1",
			},
			{
				ID:   "2",
				Name: "Falcon2",
			},
			{
				ID:   "3",
				Name: "Falcon3",
			},
		}

		rocketStoreMock.EXPECT().GetRockets().Return(rockets, nil)

		rocketService := New(rocketStoreMock)

		rkts, err := rocketService.GetRockets(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, len(rockets), len(rkts))
	})

	t.Run("delete an existing rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		id := "UUID-1"

		//mock: delete rocket and return no error
		rocketStoreMock.EXPECT().DeleteRocket(id).Return(nil)

		// DB call here
		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(context.Background(), id)

		assert.NoError(t, err)
	})
}
