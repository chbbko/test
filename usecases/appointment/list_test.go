package appointment
import (
	"errors"
	"github.com/test/utils/testhelper"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/test/entities"
	"github.com/test/repositories/mocks"
)

func TestAppointmentUseCase_List(t *testing.T) {
	var mockAppointmentList []entities.AppointmentList = []entities.AppointmentList{
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
		params entities.AppointmentListParams
	}
	type mockers struct {
		repo []entities.AppointmentList
		repoError             error
	}
	tests := []struct {
		name    string
		args    args
		mockers mockers
		want    []entities.AppointmentList
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				params: entities.AppointmentListParams{
					DoctorID: "001",
				},
			},
			mockers: mockers{
				repo: mockAppointmentList,
				repoError:             nil,
			},
			want:    mockAppointmentList,
			wantErr: false,
		},
		{
			name:    "Failed, repo return error",
			args: args{
				params: entities.AppointmentListParams{
					DoctorID: "001",
				},
			},
			mockers: mockers{repoError: errors.New("mock repo error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.AppointmentRepository)
			repo.On("ListAppointment", mock.AnythingOfType("entities.AppointmentListParams")).Return(tt.mockers.repo, tt.mockers.repoError)

			r := NewUseCase(repo)
			got, err := r.List(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAppointment() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAppointment() got = %+v, want = %+v", got, tt.want)
				return
			}
		})
	}
}
