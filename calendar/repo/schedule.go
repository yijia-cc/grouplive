package repo

import (
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Schedule struct {
	timeSlotDao dao.TimeSlot
	reservationDao dao.Reservation
	weekDao dao.Week
}

func (s Schedule) GetWeekSchedule(transaction tx.Transaction, week entity.Week) (entity.Schedule, error){
	var err error
	if len(week.ID) == 0 {
		week, err = s.weekDao.FindWeek(transaction, week.StartDate)
		if err != nil {
			return entity.Schedule{}, err
		}
	}

	reservations, err := s.reservationDao.GetWeekReservations(transaction, week)
	if err != nil {
		return entity.Schedule{}, err
	}

	timeSlots, err := s.timeSlotDao.GetWeekTimeSlots(transaction, week)
	schedule := entity.Schedule{
		Reservations: reservations,
		WeekID: week.ID,
		TimeSlots: timeSlots,
	}
	return schedule, nil
}

func NewSchedule(timeSlotDao dao.TimeSlot, reservationDao dao.Reservation, weekDao dao.Week) Schedule {
	return Schedule {
		timeSlotDao: timeSlotDao,
		reservationDao: reservationDao,
		weekDao: weekDao,
	}
}