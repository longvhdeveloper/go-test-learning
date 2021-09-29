package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", i)
	}

	sleeper.Sleep()

	//for i := countdownStart; i > 0; i-- {
	//	sleeper.Sleep()
	//}
	//
	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintf(writer,"%d\n", i)
	//}

	fmt.Fprint(writer, finalWord)
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func main() {
	sleeper := &ConfigurableSleeper{time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
