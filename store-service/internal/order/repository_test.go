//+build integration

package order_test

import (
	"store-service/internal/order"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_OrderRepository(t *testing.T) {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(localhost:3306)/toy")
	if err != nil {
		t.Fatalf("cannot tearup data err %s", err)
	}
	repository := order.OrderRepositoryMySQL{
		DBConnection: connection,
	}

	// t.Run("CreateOrder_Input_TotalPrice_14_dot_95_Should_Be_OrderID_1337620837", func(t *testing.T) {
	// 	expectedId := 1337620837
	// 	totalPrice := 14.95

	// 	actualId, err := repository.CreateOrder(totalPrice)

	// 	assert.Equal(t, nil, err)
	// 	assert.Equal(t, expectedId, actualId)
	// })
	t.Run("CreateShipping_Input_OrderID_8004359103_Should_Be_ShippingID_1_No_Error", func(t *testing.T) {
		expectShippingID := 1
		productList := []order.OrderProduct{
			{
				ProductID: 2,
				Quantity:  1,
			},
		}
		submittedOrder := order.SubmitedOrder{
			Cart:                 productList,
			ShippingMethod:       1,
			ShippingAddress:      "405/35 ถ.มหิดล",
			ShippingSubDistrict:  "ท่าศาลา",
			ShippingDistrict:     "เมือง",
			ShippingProvince:     "เชียงใหม่",
			ShippingZipCode:      "50000",
			RecipientName:        "ณัฐญา ชุติบุตร",
			RecipientPhoneNumber: "0970804292",
		}
		orderID := 8004359103

		actualShippingID, err := repository.CreateShipping(orderID, submittedOrder)
		assert.Equal(t, expectShippingID, actualShippingID)
		assert.Equal(t, err, nil)
	})
	t.Run("CreateOrder_Input_TotalPrice_14_dot_95_Should_Be_OrderID_1_No_Error", func(t *testing.T) {
		expectedId := 1
		totalPrice := 14.95

		actualId, err := repository.CreateOrder(totalPrice)

		assert.Equal(t, nil, err)
		assert.Equal(t, expectedId, actualId)
	})

}
