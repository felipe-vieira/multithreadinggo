package arbitrator

import (
	"sync"
	"time"

	"github.com/felipe-vieira/multithreadinggo/deadlockstrain/domain"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func MoveTrain(train *domain.Train, distance int, crossings []*domain.Crossing) {
	for train.Front < distance {
		//log.Printf("%d %d \n", train.ID, train.Front)
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionsInDistance(train.ID, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersection.LockedBy = -1
				// at this point all waiting threads will awake
				cond.Broadcast()
				controller.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}

func allFree(intersectionsToLock []*domain.Intersection) bool {
	for _, it := range intersectionsToLock {
		for it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*domain.Crossing) {
	var intersectionsToLock []*domain.Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	controller.Lock()
	for !allFree(intersectionsToLock) {
		// releases the controller.Lock
		cond.Wait()
	}
	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
	controller.Unlock()
}
