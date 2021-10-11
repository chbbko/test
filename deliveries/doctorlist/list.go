package doctorlist

import (
	"fmt"
	"github.com/test/deliveries/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *doctorHandler) List(ctx *gin.Context) {

	result,err := h.doctorUseCase.List()
	if err != nil{
		fmt.Println("XXX")
	}
	m := models.ToModelDoctorList(result)
	x := models.SuccessResponse{
		Success: models.Success{
			Status:  "200",
			Message: "ok",
		},
		Data: m,
	}
	ctx.JSON(http.StatusOK, x)
	return
}
