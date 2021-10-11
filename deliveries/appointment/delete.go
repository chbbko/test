package appointment
import (
	"github.com/test/deliveries/models"
	"github.com/test/entities"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *appointmentHandler) Delete(ctx *gin.Context) {
	var toDelete entities.AppointmentDelete
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "error",
			Status: "fails",
		})
		return
	}
	toDelete.ID = uint(id)
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
