package worker_helper

import "time"

func InstantWorker(f func(), interval time.Duration, stop chan bool) {
	f()
	DelayedWorker(f, interval, stop)
}

func DelayedWorker(f func(), interval time.Duration, stop chan bool) {
	t := time.NewTicker(interval)
	for {
		select {
		case <-stop:
			return
		case <-t.C:
			f()
		}
	}
}
