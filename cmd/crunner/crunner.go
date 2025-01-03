// A simple program that you may use to test your TribServer. DO NOT MODIFY!

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dylanfeehan/tribbler/api/tribrpc"
	"github.com/dylanfeehan/tribbler/pkg/tribclient"
)

var port = flag.Int("port", 9010, "TribServer port number")

type cmdInfo struct {
	cmdline  string
	funcname string
	nargs    int
}

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "The crunner program is a testing tool that that creates and runs an instance")
		fmt.Fprintln(os.Stderr, "of the TribClient. You may use it to test the correctness of your TribServer.\n")
		fmt.Fprintln(os.Stderr, "Usage:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Possible commands:")
		fmt.Fprintln(os.Stderr, "  CreateUser:                uc userID")
		fmt.Fprintln(os.Stderr, "  GetSubscriptions:          sl userID")
		fmt.Fprintln(os.Stderr, "  AddSubscriptions:          sa userID targetUserID")
		fmt.Fprintln(os.Stderr, "  RemoveSubscriptions:       sr userID targetUserID")
		fmt.Fprintln(os.Stderr, "  GetTribbles:               tl userID")
		fmt.Fprintln(os.Stderr, "  PostTribbles:              tp userID contents")
		fmt.Fprintln(os.Stderr, "  GetTribblesBySubscription: ts userID")
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}
	cmd := flag.Arg(0)
	handle, err := tribclient.NewTribHandle("localhost", *port)
	if err != nil {
		log.Fatalln("Failed to create TribClient:", err)
	}

	cmdlist := []cmdInfo{
		{"uc", "TribServer.CreateUser", 1},
		{"sl", "TribServer.GetSubscriptions", 1},
		{"sa", "TribServer.AddSubscription", 2},
		{"sr", "TribServer.RemoveSubscription", 2},
		{"tl", "TribServer.GetTribbles", 1},
		{"tp", "TribServer.AddTribble", 2},
		{"ts", "TribServer.GetTribblesBySubscription", 1},
	}

	cmdmap := make(map[string]cmdInfo)
	for _, j := range cmdlist {
		cmdmap[j.cmdline] = j
	}

	ci, found := cmdmap[cmd]
	if !found {
		flag.Usage()
		os.Exit(1)
	}
	if flag.NArg() < (ci.nargs + 1) {
		flag.Usage()
		os.Exit(1)
	}

	switch cmd {
	case "uc": // user create
		status, err := handle.CreateUser(flag.Arg(1))
		printStatus(ci.funcname, status, err)
	case "sl": // subscription list
		subs, status, err := handle.GetSubscriptions(flag.Arg(1))
		printStatus(ci.funcname, status, err)
		if err == nil && status == 0 {
			fmt.Println(strings.Join(subs, " "))
		}
	case "sa":
		status, err := handle.AddSubscription(flag.Arg(1), flag.Arg(2))
		printStatus(ci.funcname, status, err)
	case "sr": // subscription remove
		status, err := handle.RemoveSubscription(flag.Arg(1), flag.Arg(2))
		printStatus(ci.funcname, status, err)
	case "tl": // tribble list
		tribbles, status, err := handle.GetTribbles(flag.Arg(1))
		printStatus(ci.funcname, status, err)
		if err == nil && status == 0 {
			printTribbles(tribbles)
		}
	case "ts": // tribbles by subscription
		tribbles, status, err := handle.GetTribblesBySubscription(flag.Arg(1))
		printStatus(ci.funcname, status, err)
		if err == nil && status == 0 {
			printTribbles(tribbles)
		}
	case "tp": // tribble post
		status, err := handle.PostTribble(flag.Arg(1), flag.Arg(2))
		printStatus(ci.funcname, status, err)
	}
}

// should use enums instead?
func tribStatusToString(status int) (s string) {
	switch status {
	case 0:
		s = "OK"
	case 1:
		s = "NoSuchUser"
	case 2:
		s = "NoSuchTargetUser"
	case 3:
		s = "Exists"
	}
	return
}

func printStatus(cmdName string, status int, err error) {
	if err != nil {
		fmt.Println("ERROR:", cmdName, "got error:", err)
	} else if status != 0 {
		fmt.Println(cmdName, "ERROR:", cmdName, "replied with status", tribStatusToString(status))
	} else {
		fmt.Println(cmdName, "OK")
	}
}

func printTribble(t tribrpc.Tribble) {
	fmt.Printf("%16.16s - %s - %s\n", t.GetUserID(), string(t.GetPostedTime()), t.GetContents())
}

func printTribbles(tribbles []*tribrpc.Tribble) {
	for _, t := range tribbles {
		tribble := *t
		printTribble(tribble)
	}
}
