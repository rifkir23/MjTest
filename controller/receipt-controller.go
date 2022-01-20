package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"github.com/wilopo-cargo/microservice-receipt-sea/service"
	"net/http"
	"strconv"
)

//ReceiptController is a ...
type ReceiptController interface {
	List(context *gin.Context)
}

type receiptController struct {
	receiptService service.ReceiptService
	jwtService     service.JWTService
}

//NewReceiptController create a new instances of BoookController
func NewReceiptController(receiptServ service.ReceiptService, jwtServ service.JWTService) ReceiptController {
	return &receiptController{
		receiptService: receiptServ,
		//jwtService:  jwtServ,
	}
}

// List godoc
// @Summary All example
// @Schemes
// @Description Receipt List
// @Accept json
// @Produce json
// @Param List page  query int  false  "Pages"
// @Param List limit  query int  false  "Limit"
// @Param List receiptType query  string  false  "receiptType"
// @Success 200 {object} helper.Response{data=dto.ReceiptListByTypeResult}
// @Router /receipt/list [GET]
func (c *receiptController) List(context *gin.Context) {
	page, err := strconv.ParseInt(context.Query("page"), 0, 0)
	limit, err := strconv.ParseInt(context.Query("limit"), 0, 0)
	receiptType := context.Query("receiptType")
	if err != nil {
		res := helper.BuildErrorResponse("No param int was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var receipts = c.receiptService.List(page, limit, receiptType)
	res := helper.BuildResponse(true, "OK", receipts)
	context.JSON(http.StatusOK, res)
}
