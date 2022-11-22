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

type ProductController struct {
	productService services.ProductServiceInterface
}

func NewProductController(productService services.ProductServiceInterface) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) NewProduct(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrForbidden.Code, errToken.Error()))
	}

	var newProduct models.NewProduct
	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	errNewProduct := pc.productService.NewProduct(ctx, newProduct, idToken)
	if errNewProduct != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errNewProduct.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "Success add new product"))
}

func (pc *ProductController) GetProductById(c echo.Context) error {
	idString := c.Param("id")
	id, errId := strconv.Atoi(idString)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errId.Error()))
	}

	ctx := c.Request().Context()
	product, err := pc.productService.GetProductById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "success get product by id", product))
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := pc.productService.GetAllProducts(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "success get all products", products))
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrBadRequest.Code, errToken.Error()))
	}

	idString := c.Param("id")
	id, errId := strconv.Atoi(idString)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errId.Error()))
	}

	ctx := c.Request().Context()
	errDelete := pc.productService.DeleteProduct(ctx, idToken, id)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errDelete.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "success delete product"))
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errToken.Error()))
	}

	idString := c.Param("id")
	id, errId := strconv.Atoi(idString)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errId.Error()))
	}

	var updateProduct models.UpdateProduct
	err := c.Bind(&updateProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	product, errUpdate := pc.productService.UpdateProduct(ctx, updateProduct, idToken, id)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errUpdate.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "succes update product", product))
}
