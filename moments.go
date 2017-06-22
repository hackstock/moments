//Package moments provides utility functions for calculating and presenting
//in human readable and understandable form,time elapsed since a given timestamp.
//It also has helper functions for scheduling one-time or periodic execution of
//functions
package moments

import (
	"fmt"
	"time"
)

const (
	secondInMillis = 1000
	minuteInMillis = 60 * secondInMillis
	hourInMillis   = 60 * minuteInMillis
	dayInMillis    = 24 * hourInMillis
)

//TimeAgo returns a string describing how long it has been since
//the time argument passed int
func TimeAgo(t time.Time) (timeSince string) {
	timestamp := t.Unix()
	now := time.Now().Unix()

	if timestamp > now || timestamp <= 0 {
		timeSince = ""
	}

	timeElapsed := now - timestamp
	if timeElapsed < minuteInMillis {
		timeSince = "just now"
	} else if timeElapsed < 2*minuteInMillis {
		timeSince = "a minute agon"
	} else if timeElapsed < 50*minuteInMillis {
		timeSince = fmt.Sprintf("%d minutes ago", timeElapsed/minuteInMillis)
	} else if timeElapsed < 90*minuteInMillis {
		timeSince = "an hour ago"
	} else if timeElapsed < 24*hourInMillis {
		timeSince = fmt.Sprintf("%d hours ago", timeElapsed/hourInMillis)
	} else if timeElapsed < 48*hourInMillis {
		timeSince = "yesterday"
	} else {
		timeSince = fmt.Sprintf("%d days ago", timeElapsed/dayInMillis)
	}
	return
}
