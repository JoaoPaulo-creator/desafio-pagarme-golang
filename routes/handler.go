package routes

import "github.com/labstack/echo/v4"

func HandlerInitializer() {
	e := echo.New()
	e.POST("/create-transaction", CreateTransactionHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
