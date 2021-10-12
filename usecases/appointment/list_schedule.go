package appointment

import "github.com/test/entities"

func (r appointmentUseCase) ListSchedule(params entities.ScheduleParams) ([]entities.ScheduleList,error) {
	return r.repo.ListSchedule(params)
}
