package appointment
import (
	"github.com/test/deliveries/models"
	"github.com/test/entities"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func (h *appointmentHandler) List(ctx *gin.Context) {
	var qString models.AppointmentQueryString
	err := ctx.BindQuery(&qString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "error on binding data",
			Status: "fails",
		})
		return
	}
	var params entities.AppointmentListParams
	if copyError := copier.Copy(&params, qString); copyError != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error on binding data",
			Status: "fails",
		})
		return
	}
	result,err := h.appointmentUseCase.List(params)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error from repositories",
			Status: "fails",
		})
		return
	}
	m := models.ToModelAppointmentList(result)

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
