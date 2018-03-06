package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	msgPtr := flag.String("m", "End of an era", "Message for notification")
	flag.Parse()

	var t string

	fmt.Printf("%v", os.Args)

	if len(os.Args) < 2 {
		t = "10s"
	} else {
		t = os.Args[1]
	}

	i, err := strconv.Atoi(t[0 : len(t)-1])
	if err != nil {
		fmt.Sprintf("Error: Could not convert to number. \n%v", err)
	}

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/golang.png",
		AppName:     "Era",
	})

	timer := time.NewTimer(time.Duration(i) * getTimeUnit(t))

	<-timer.C

	err = notify.Push("Era", *msgPtr, "", "")
	if err != nil {
		fmt.Sprint("Error: Could not send notification. \n%v", err)
	}
}

func getTimeUnit(s string) time.Duration {

	a := strings.Split(s, "")[len(s)-1]

	switch a {
	case "m":
		return time.Minute
	case "h":
		return time.Hour
	default:
		return time.Second
	}
}
