package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/test/deliveries/models"
	"github.com/test/entities"
	"net/http"
)

func (h *appointmentHandler) ListSchedule(ctx *gin.Context) {
	var qString models.ScheduleQueryString
	err := ctx.BindQuery(&qString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "error on binding data",
			Status: "fails",
		})
		return
	}
	var params entities.ScheduleParams
	if copyError := copier.Copy(&params, qString); copyError != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error on binding data",
			Status: "fails",
		})
		return
	}
	result,err := h.appointmentUseCase.ListSchedule(params)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error from repositories",
			Status: "fails",
		})
		return
	}

	m := models.ToModelScheduleList(result)

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
