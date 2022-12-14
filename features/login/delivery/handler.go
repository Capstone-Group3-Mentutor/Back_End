package delivery

import (
	"be12/mentutor/features/login"
	"be12/mentutor/utils/helper"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := AuthDelivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Login())
}

func (h *AuthDelivery) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var url string
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		cnv := ToDomain(input)
		res, token, err := h.authUsecase.Login(cnv)
		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "email not found") == true {
				errx := errors.New("user not found")
				return c.JSON(http.StatusNotFound, helper.FailedResponse(errx.Error()))
			} else {	
				errx := errors.New("Invalid Input From Client")
				return c.JSON(http.StatusBadRequest, helper.FailedResponse(errx.Error()))
			}
		}

		if res.ID < 1000 {
			err := h.authUsecase.GetToken(res.ID)
			if err != nil {
				log.Print(err.Error())
				url, err = helper.GetUrl()
				// log.Print(url)
				if err != nil {
					log.Print(err.Error())
					return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed get url Oauth"))
				}
				err = h.authUsecase.InsertToken(res.ID)
				if err != nil {
					log.Print(err.Error())
					return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert to db"))
				}
			}
		}	
		res.Token = token
		return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", ToResponse(res, "login", url)))
	}
}
