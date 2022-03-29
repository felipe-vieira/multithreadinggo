package main

import "time"

func robotsAndBariers() {
	barrier := NewBarrier(4)
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)
	go waitOnBarrier("yellow", 8, barrier)
	go waitOnBarrier("green", 12, barrier)
	time.Sleep(time.Duration(100) * time.Second)
}

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		println(name, "running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		println(name, "is waiting on barrier")
		barrier.Wait()
	}
}
