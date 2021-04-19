package controller

import (
	"github.com/labstack/echo/v4"
	"golang-echo-swagger-example/model"
	"log"
	"net/http"
)

// GetAllUser godoc
// @Summary Get all users
// @Description Get all user items
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func GetAllUser(c echo.Context) error {
	users := []model.User{
		model.User{Name: "Test", Email: "test@gmail.com"},
		model.User{Name: "Test", Email: "test1@gmail.com"},
	}
	return c.JSON(http.StatusOK, users)
}

// SaveUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "New User"
// @Success 201 {object} model.User
// @Router /users [post]
func SaveUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user item
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func GetUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("User Id : %s", id)

	user := model.User{
		Name:  "Test",
		Email: "test@gmail.com",
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.User true "User Info"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("Updated User Id : %s", id)

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} model.User
// @Router /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("Deleted User Id : %s", id)

	return c.NoContent(http.StatusNoContent)
}
