package deadlock

import (
	"time"

	"github.com/felipe-vieira/multithreadinggo/deadlockstrain/domain"
)

func MoveTrain(train *domain.Train, distance int, crossings []*domain.Crossing) {
	for train.Front < distance {
		//log.Printf("%d %d \n", train.ID, train.Front)
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.ID
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
