package handler

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h CatsShop) CreateCat(c echo.Context) error {
	user:= getToken(c)
	if user.Admin == true{
		cat:=&model.CatParams{}
		if err := c.Bind(cat); err != nil {
			return err
		}
		err, _ := h.client.CreateCat(context.Background(), &protocol.Catparams{CatName: cat.CatName,CatAge: int32(cat.CatAge),CatGender: cat.CatGender,CatBreed: cat.CatBreed})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error)
		}
		return c.String(http.StatusCreated, "Cat have been created")
	}
	return echo.ErrUnauthorized
}

func (h CatsShop) GetCat(c echo.Context) error {
	cat:=&model.CatParams{}
	if err := c.Bind(cat); err != nil {
		return err
	}
	catInfo, _ := h.client.GetCat(context.Background(), &protocol.Catparams{ID: int32(cat.ID)})
	if catInfo.Error != "" {
		return c.String(http.StatusInternalServerError, catInfo.Error)
	}
	return c.JSON(http.StatusCreated, catInfo)
}

func (h CatsShop) ListCats(c echo.Context) error {
	list, _ := h.client.ListCats(context.Background(), &protocol.Catparams{})
	if list.Error != "" {
		return c.String(http.StatusInternalServerError, list.Error)
	}
	return c.String(http.StatusOK, list.Json)
}


func (h CatsShop) UpdateCat(c echo.Context) error {
	user:= getToken(c)
	if user.Admin == true{
		cat:=&model.CatParams{}
		if err := c.Bind(cat); err != nil {
			return err
		}
		err, _ := h.client.UpdateCat(context.Background(), &protocol.Catparams{ID: int32(cat.ID),CatAge: int32(cat.CatAge)})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error)
		}
		return c.String(http.StatusCreated, "Cat have been updated")
	}
	return echo.ErrUnauthorized
}

func (h CatsShop) DeleteCat(c echo.Context) error {
	user:= getToken(c)
	if user.Admin == true {
		cat:=&model.CatParams{}
		if err := c.Bind(cat); err != nil {
			return err
		}
		err, _ := h.client.DeleteCat(context.Background(), &protocol.Catparams{ID: int32(cat.ID)})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error)
		}
		return c.String(http.StatusCreated, "Cat have been destroyed")
	}
	return echo.ErrUnauthorized
}
