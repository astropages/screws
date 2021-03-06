package screws

import (
	"log"
	"time"
)

//IDailyBomb ...
type IDailyBomb interface {
	Defuse()
}

//FixingADailyBomb ...
func FixingADailyBomb(function func(), clock [3]int) IDailyBomb {
	d := &dailyBomb{
		Powder:  function,
		Clock:   clock,
		Defused: make(chan bool),
	}
	go d.ignite()
	return d
}

//dailyBomb ...
type dailyBomb struct {
	Powder  func()
	Clock   [3]int
	Defused chan bool
}

//ignite ...
func (d *dailyBomb) ignite() {
	d.Powder()
	for {
		today := time.Now()
		tomorrow := today.Add(time.Hour * 24)
		tomorrow = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), d.Clock[0], d.Clock[1], d.Clock[2], 0, tomorrow.Location())
		select {
		case <-d.Defused:
			log.Println("LoopBomb  Defused")
			return
		case <-time.After(tomorrow.Sub(today)):
			d.Powder()
		}
	}
}

//Defuse ...
func (d *dailyBomb) Defuse() {
	d.Defused <- true
}
