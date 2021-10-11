package appointment
import (

	"fmt"
	"github.com/test/entities"
	"github.com/test/utils/testhelper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/test/usecases/mocks"
)

func Test_appointmentHandler_List(t *testing.T) {
	var MockAppointmentList []entities.AppointmentList = []entities.AppointmentList{
		entities.AppointmentList{
			PatiencePinCode: "111",
			DoctorID:        "001",
			Phone:           "089000000",
			Ref:             "",
			Start:           testhelper.MockFromDateString("2006-01-02T03:04:05-0700", "2020-10-30T00:00:00+0700"),
			End: testhelper.MockFromDateString("2006-01-02T03:04:05-0700", "2020-10-30T00:00:00+0700"),
		},
	}
	type args struct {
		queryParams string
	}
	type mockers struct {
		uc    []entities.AppointmentList
		ucErr error
	}
	tests := []struct {
		name           string
		args           args
		mockers        mockers
		wantStatusCode int
	}{
		{
			name: "Successful",
			args: args{
				queryParams: "?doctor_id=001",
			},
			mockers: mockers{
				uc: MockAppointmentList,
			},
			wantStatusCode: http.StatusOK,
		},
		//TODO:: add fail case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := new(mocks.AppointmentUseCase)
			uc.On("List",
				mock.AnythingOfType("entities.AppointmentListParams")).Return(tt.mockers.uc, tt.mockers.ucErr)

			ginEn := gin.Default()
			gin.SetMode(gin.TestMode)
			NewEndpointHttpHandler(ginEn,  uc)
			requestURL := fmt.Sprintf("/apis/appointment%s", tt.args.queryParams)
			req, _ := http.NewRequest("GET", requestURL, nil)
			res := httptest.NewRecorder()
			ginEn.ServeHTTP(res, req)
			if tt.wantStatusCode != res.Code {
				t.Errorf("http appointment list () got = %v, want %v, case: %s", res.Code, tt.wantStatusCode, tt.name)
			}
		})
	}
}
