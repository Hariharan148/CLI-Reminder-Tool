package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"strings"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/gen2brain/beeep"
)

const (
	markName = "CLI_REMAINDER"
	markValue = "1"
)

func main(){

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <HH:MM> <Text Message>\n", os.Args[0])
		os.Exit(1)
	}

	timer := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	result, err := w.Parse(os.Args[1], timer )

	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}

	diff := result.Time.Sub(timer)

	if timer.After(result.Time){
		fmt.Printf("Set a Future Time!\n")
		os.Exit(2)
	}
	if result == nil {
		fmt.Println("Unable to Parse the Args")
		os.Exit(3)
	}

	if os.Getenv(markName) == markValue{
		time.Sleep(diff)
		err = beeep.Alert("Remainder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil{
			fmt.Println(err)
			os.Exit(3)
		}
	} else{
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Remainder will be displayed after", diff.Round(time.Second))
		os.Exit(0)
	}

}