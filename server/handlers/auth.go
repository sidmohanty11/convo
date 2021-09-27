package handlers

import (
	"convo_backend/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	u := &models.User{
		ID:       bson.NewObjectId(),
		LastSeen: time.Now(),
	}

	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "something went wrong"}
	}

	if u.Number == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid number"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").Insert(u); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	u := &models.User{
		LastSeen: time.Now(),
	}
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "something went wrong"}
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("convo").C("users").Find(bson.M{"number": u.Number, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid number or password"}
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "db not responding"}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	u.Token, err = token.SignedString([]byte(Key))

	if err != nil {
		return err
	}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
