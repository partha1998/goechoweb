package tronics

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var v = validator.New()

var arr map[int]string = map[int]string{
	1: "Bike",
	2: "TV",
	3: "Car",
}

func sample(c echo.Context) error {
	// fmt.Println("inside handler function")
	return c.String(http.StatusOK, "heelo from get url")
}

func getProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var products map[int]string
	if err == nil {
		for key, product := range arr {
			if key == id {
				products = make(map[int]string)
				products[key] = product
			}
		}
	}
	if products == nil {
		return c.JSON(http.StatusNotFound, "product id invalid")

	}
	return c.JSON(http.StatusOK, products)
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, arr)
}

func addProduct(c echo.Context) error {
	type body struct {
		Name string `json:"Product_name" validate:"required,min=4"`
		// Vendor     string  `json:"Vendor_name" validate:"min=5,max=10"`
		// Email      string  `json:"email" validate:"required_with=Vendor,email"`
		// Website    string  `json:"website" validate:"url"`
		// Country    string  `json:"country" validate:"len=2"`
		// IpAddress  string  `json:"ip" validate:"ip"`
	}

	var reqBody body

	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}

	err1 := v.Struct(reqBody)
	if err1 != nil {
		return err1
	}

	arr[len(arr)+1] = reqBody.Name
	return c.JSON(http.StatusOK, arr[len(arr)])
}
