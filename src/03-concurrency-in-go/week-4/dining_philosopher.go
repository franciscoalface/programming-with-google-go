/*
Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type chop struct {
	sync.Mutex
}

type philo struct {
	id          int
	left, right *chop
}

func informAction(action string, id int) {
	fmt.Printf("%s %d\n", action, id+1)
}

func (p *philo) eat() {
	for j := 0; j < 3; j++ {
		p.left.Lock()
		p.right.Lock()

		informAction("starting to eat", p.id)
		time.Sleep(time.Second)

		p.right.Unlock()
		p.left.Unlock()

		informAction("finishing eating", p.id)
		time.Sleep(time.Second)
	}
	wGroup.Done()
}

var wGroup sync.WaitGroup

func main() {
	chopSticks := make([]*chop, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(chop)
	}

	philosophers := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philo{
			id:    i,
			left:  chopSticks[i],
			right: chopSticks[(i+1)%5],
		}
		wGroup.Add(1)
		go philosophers[i].eat()
	}
	wGroup.Wait()
}
