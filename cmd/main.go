package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Event struct {
	duration int64
	text string
}

type EventRelation struct {
	preceding Event
	longest Event
}



func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Unable to open file %s", os.Args[1])
	}

	defer f.Close();

	r := bufio.NewReader(f)

	var prevEvent time.Time
	var max int64
	var previousEvent Event
	var currentEvent Event
	var longestEvent EventRelation

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		parts := strings.Split(line, " ")

		if len(parts) >= 2 {
			ts := strings.Replace(parts[1], ",", ".", -1)

			eventTime, err := time.Parse("15:04:05.000", ts)

			if err == nil {
				if !prevEvent.IsZero() {
					dur := eventTime.Sub(prevEvent)

					currentEvent.duration = dur.Milliseconds();
					currentEvent.text = line

					fmt.Printf("[%05v] -> %s\n", currentEvent.duration, currentEvent.text)

					if currentEvent.duration > max {
						max = currentEvent.duration

						longestEvent.longest = currentEvent;
						longestEvent.preceding = previousEvent;
					} else {
						previousEvent = currentEvent;
					}
				} else {
					previousEvent.duration = 0;
					previousEvent.text = line;
				}

				prevEvent = eventTime;
			}
		}
	}

	fmt.Printf("The longest event was: \n")
	fmt.Printf("[%05v] -> %s\n", longestEvent.longest.duration, longestEvent.longest.text)

	fmt.Printf("Preceeded by: \n")
	fmt.Printf("[%05v] -> %s\n", longestEvent.preceding.duration, longestEvent.preceding.text)
}
