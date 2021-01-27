package apps

import (
	authController "github.com/Ferza17/Products-RESTAPI/controllers/auth"
	customerController "github.com/Ferza17/Products-RESTAPI/controllers/customers"
	orderController "github.com/Ferza17/Products-RESTAPI/controllers/orders"
	"github.com/Ferza17/Products-RESTAPI/utils/auth"
	"log"
)

func mapURL() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	/*============================
	======= Customers URL ========
	=============================*/

	// POST
	router.POST("/customer", customerController.CreateCustomer)

	/*============================
	======= Auth URL =============
	=============================*/

	// Create TOKEN API if Email & Password Exist, Expiration 3 day
	router.POST("/customer/auth", authController.CreateAuthToken)
	router.GET("/customer/auth", authController.RefreshToken)

	/*============================
	======= Orders URL ========
	=============================*/

	//Create Order
	router.POST("/order", auth.Authentication, orderController.CreateOrder)

}
