package controllers

import (
	"e-commerce/helpers"
	"e-commerce/middlewares"
	"e-commerce/models"
	"e-commerce/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryController(categoryService services.CategoryServiceInterface) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrForbidden.Code, "unauthorized"))
	}

	var newCategory models.NewCategory
	err := c.Bind(&newCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	errCreateCategory := cc.categoryService.CreateCategory(ctx, newCategory, idToken)
	if errCreateCategory != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errCreateCategory.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "Success create a new category"))
}

func (cc *CategoryController) GetCategoryById(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}

	ctx := c.Request().Context()
	category, err := cc.categoryService.GetCategoryById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success get category by id", category))
}

func (cc *CategoryController) GetCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := cc.categoryService.GetCategory(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success Get All Categories", categories))
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	idString := c.Param("id")
	id, errId := strconv.Atoi(idString)
	if errId != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Id not found"))
	}

	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errToken.Error()))
	}

	ctx := c.Request().Context()
	err := cc.categoryService.DeleteCategory(ctx, idToken, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "Success delete category"))
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrForbidden.Code, "unauthorized"))
	}

	var updateCategory models.UpdateCategory
	errBind := c.Bind(&updateCategory)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	category, errUpdate := cc.categoryService.UpdateCategory(ctx, updateCategory, idToken, id)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errUpdate.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success update category", category))
}
