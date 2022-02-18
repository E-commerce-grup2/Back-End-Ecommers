package main

import (
	config "Back-End-Ecommers/configs"
	addressController "Back-End-Ecommers/delivery/controllers/address"
	authController "Back-End-Ecommers/delivery/controllers/auth"
	paymentController "Back-End-Ecommers/delivery/controllers/payment"
	userController "Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/route"
	AddressRepo "Back-End-Ecommers/repository/address"
	authRepo "Back-End-Ecommers/repository/auth"
	paymentRepo "Back-End-Ecommers/repository/payment"
	userRepo "Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	//REPOSITORY-DATABASE
	userRepo := userRepo.New(db)
	addressRepo := AddressRepo.New(db)
	authRepo := authRepo.New(db)
	paymentRepo := paymentRepo.New(db)

	//CONTROLLER
	userController := userController.New(userRepo)
	addressController := addressController.New(addressRepo)
	authController := authController.New(authRepo)
	paymentController := paymentController.New(paymentRepo)

	e := echo.New()

	route.RegisterPath(
		e,
		userController,
		addressController,
		authController,
	)

	route.PaymentMethodPath(
		e,
		paymentController,
	)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
