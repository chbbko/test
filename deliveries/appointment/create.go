package appointment
import (
	"github.com/test/deliveries/models"
	"net/http"
	"github.com/gin-gonic/gin"

)

func (h *appointmentHandler) Create(ctx *gin.Context) {
	var toCreate models.AppointmentCreate

	if err := ctx.BindJSON(&toCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "error on binding data",
			Status: "fails",
		})
		return
	}
	//TODO:: validate struct
	e, err := models.ToEntitiesAppointmentCreate(toCreate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error on parsing data",
			Status: "fails",
		})
		return
	}
	result := h.appointmentUseCase.Create(e)
	if result != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error on creating data",
			Status: "fails",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Success{
		Status: "success",
		Message: "OK",
	})
	return
}
