package handler

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h CatsShop) CreateUser(c echo.Context) error {
	user := &model.UserParams{}
	if err := c.Bind(user); err != nil {
		return err
	}
	err,_:=h.client.CreateUser(context.Background(),&protocol.Userparams{NickName: user.NickName,Admin: user.Admin,Password: user.Password,Address: user.Address})
	if err != nil{
		return c.String(http.StatusInternalServerError, err.Error)
	}
	return c.String(http.StatusOK,"User created")
}

func (h CatsShop) Login(c echo.Context) error {
	user := &model.UserParams{}
	if err := c.Bind(user); err != nil {
		return err
	}
	token,_:=h.client.Login(context.Background(),&protocol.Userparams{NickName: user.NickName,Password: user.Password})
	if token.Error!=""{
		return c.String(http.StatusInternalServerError, token.Error)
	}
	return c.JSON(http.StatusOK,token.Token)
}

func (h CatsShop) ListUsers(c echo.Context) error {
	user := getToken(c)
	if user.Admin == true{
		list, _ := h.client.ListUsers(context.Background(), &protocol.Userparams{})
		if list.Error != "" {
			return c.String(http.StatusInternalServerError, list.Error)
		}
		return c.String(http.StatusOK, list.Json)
	}
	return echo.ErrUnauthorized
}

func (h CatsShop) UpdateUser(c echo.Context) error {
	target := &model.UserParams{}
	if err := c.Bind(target); err != nil {
		return err
	}
	user:= getToken(c)
	err,_:=h.client.UpdateUser(context.Background(),&protocol.Userparams{NickName: user.NickName,Address: target.Address,Password: target.Password})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error)
	}
	return c.String(http.StatusOK,"Address updated")
}

func (h CatsShop) DeleteUser(c echo.Context) error {
	target := &model.UserParams{}
	if err := c.Bind(target); err != nil {
		return err
	}
	user:= getToken(c)
	if user.Admin == true || user.NickName == target.NickName {
		err,_:=h.client.DeleteUser(context.Background(), &protocol.Userparams{NickName: target.NickName,Password: target.Password})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error)
		}
		return c.String(http.StatusOK, "User deleted")
	}
	return echo.ErrUnauthorized
}
