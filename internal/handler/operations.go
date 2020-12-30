package handler

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h CatsShop) BuyCat(c echo.Context) error {
	user:= getToken(c)
	operation := &model.OperationParams{}
	if err := c.Bind(operation); err != nil {
		return err
	}
	operation.NewOwnerNick=user.NickName
	operation.Status="In progress"
	err,_:=h.client.AddOperation(context.Background(),&protocol.Operationparams{NewOwnersNick: operation.NewOwnerNick,CatName: operation.CatName,CatID: int32(operation.CatID),Status: operation.Status})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error)
	}
	err, _ = h.client.DeleteCat(context.Background(), &protocol.Catparams{ID: int32(operation.CatID)})
	if err != nil{
		return c.String(http.StatusInternalServerError, err.Error)
	}
	return c.String(http.StatusOK,"Operation succeed")
}

func (h CatsShop) GetOperation(c echo.Context) error {
	operation := &model.OperationParams{}
	if err := c.Bind(operation); err != nil {
		return err
	}
	operInfo,_:=h.client.GetOperation(context.Background(),&protocol.Operationparams{CatID: int32(operation.CatID)})
	if operInfo.Error != "" {
		return c.String(http.StatusInternalServerError, operInfo.Error)
	}
	return c.JSON(http.StatusOK,operInfo)
}

func (h CatsShop) ListOperation(c echo.Context) error {
	user:= getToken(c)
	list,_:=h.client.ListOperations(context.Background(),&protocol.Userparams{NickName: user.NickName,Admin: user.Admin})
	if list.Error != ""{
		return c.String(http.StatusInternalServerError, list.Error)
	}
	return c.String(http.StatusOK,list.Json)
}

func (h CatsShop) EditOperation(c echo.Context) error {
	user := getToken(c)
	if user.Admin == true {
		operation := &model.OperationParams{}
		if err := c.Bind(operation); err != nil {
			return err
		}
		err, _ := h.client.EditOperation(context.Background(), &protocol.Operationparams{CatID: int32(operation.CatID),Status: operation.Status})
		if err!= nil{
			return c.String(http.StatusInternalServerError, err.Error)
		}
		return c.String(http.StatusOK, "Status updated")
	}
	return echo.ErrUnauthorized
}
