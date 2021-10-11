package appointment
import (
	"github.com/test/deliveries/models"
	"github.com/test/entities"
	"net/http"
	"github.com/gin-gonic/gin"

)

func (h *appointmentHandler) Delete(ctx *gin.Context) {
	var toDelete entities.AppointmentDelete
	id := ctx.Param("id")
	toDelete.ID = id
	//TODO:: validate struct

	result := h.appointmentUseCase.Cancel(toDelete)
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
