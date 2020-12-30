package handler

import (
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"strings"
)

const key string ="sirok"

type CatsShop struct {
	client protocol.CatsShopClient
}

func NewCatsShop(client protocol.CatsShopClient) *CatsShop {
	return &CatsShop{client: client}
}
func getToken(c echo.Context) *model.UserParams {
	req:=c.Request()
	header:=req.Header["Authorization"]
	header=strings.Split(header[0]," ")
	token, err := jwt.ParseWithClaims(header[1], &model.UserParams{}, func(token *jwt.Token) (interface{},error) {
		return []byte(key),nil
	})
	if err != nil {
		return nil
	}
	if claims, ok := token.Claims.(*model.UserParams); ok && token.Valid {
		return claims
	}
	return nil
}
