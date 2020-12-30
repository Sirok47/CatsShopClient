package main

import (
	"fmt"
	"github.com/Sirok47/CatsShopClient/internal/handler"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

func main()  {
	TokenValidation := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("sirok"),
	})
	clientConnection, err := grpc.Dial(":8080", grpc.WithInsecure())
	client := protocol.NewCatsShopClient(clientConnection)
	if err!=nil{
		fmt.Print(err)
		return
	}

	hndl:= handler.NewCatsShop(client)

	e := echo.New()

	e.POST("/user/signup", hndl.CreateUser)

	e.GET("/user/login", hndl.Login)

	e.GET("/user/list", hndl.ListUsers, TokenValidation)

	e.PUT("/user/edit", hndl.UpdateUser, TokenValidation)

	e.DELETE("/user/delete", hndl.DeleteUser, TokenValidation)

	e.POST("/cat/create", hndl.CreateCat, TokenValidation)

	e.GET("/cat/get", hndl.GetCat)

	e.GET("/cat/list", hndl.ListCats)

	e.PUT("/cat/update", hndl.UpdateCat, TokenValidation)

	e.DELETE("/cat/delete", hndl.DeleteCat, TokenValidation)

	e.POST("/operation/buyCat", hndl.BuyCat,TokenValidation)

	e.GET("/operation/get", hndl.GetOperation,TokenValidation)

	e.GET("/operation/list", hndl.ListOperation,TokenValidation)

	e.PUT("/operation/edit", hndl.EditOperation,TokenValidation)

	e.Logger.Fatal(e.Start(":1323"))
}
