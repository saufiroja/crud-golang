package controllers

import (
	"crud/domain"
	"crud/entity"
	"crud/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S domain.PostService
}

// CREATE
func (control *Controller) Create(c echo.Context) error {
	post := entity.Post{}
	error := c.Bind(&post)
	if error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": error.Error(),
		})
	}

	err := control.S.Create(post)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "eeror create post",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.WebResponse{
		Code:    201,
		Message: "success create post",
		Data:    post,
	})
}

// FIND ALL
func (control *Controller) FindAll(c echo.Context) error {
	posts, err := control.S.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "error get all post",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:    200,
		Message: "success get all posts",
		Data:    posts,
	})
}

// FIND ONE
func (control *Controller) FindOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := control.S.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "error get post by id",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:    200,
		Message: "success post by id",
		Data:    post,
	})
}

// UPDATE
func (control *Controller) Update(c echo.Context) error {
	var post entity.Post
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "error",
			Data:    err.Error(),
		})
	}

	posts, er := control.S.Update(id, post)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    500,
			Message: "error update post",
			Data:    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:    200,
		Message: "success update post",
		Data:    posts,
	})
}

// DELETE
func (control *Controller) Delete(c echo.Context) error {
	var post entity.Post
	id, _ := strconv.Atoi(c.Param("id"))

	err := control.S.Delete(id, post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    500,
			Message: "error delete post",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.WebResponse{
		Code:    200,
		Message: "success delete user",
	})
}
