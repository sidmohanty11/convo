package handlers

import (
	"convo_backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) GetUserById(c echo.Context) (err error) {
	id := c.Param("id")
	u := &models.User{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").FindId(bson.ObjectIdHex(id)).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) GetActiveUser(c echo.Context) (err error) {
	u := &models.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) GetUserByNumber(number string) (u *models.User, err error) {
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").Find(bson.M{"number": number}).One(&u); err != nil {
		if err == mgo.ErrNotFound {
			return nil, echo.ErrNotFound
		}
		return nil, &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	return u, nil
}
