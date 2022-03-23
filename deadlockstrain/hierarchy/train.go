package hierarchy

import (
	"sort"
	"time"

	"github.com/felipe-vieira/multithreadinggo/deadlockstrain/domain"
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
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*domain.Crossing) {
	var intersectionsToLock []*domain.Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].ID < intersectionsToLock[j].ID
	})

	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}

}
