// We mock the functionality of time.sleep
// Spysleeper keeps track of the number of times sleep is being called

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord ="Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i>0; i--{
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
type DefaultSleeper struct{}
func (d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}
type Sleeper interface
{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func main(){
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
