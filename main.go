package main

import (
    "fmt"
    "log"
    gc "github.com/rthornton128/goncurses"
)

func main() {
    stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

    files := ReadDir("/home/")

    rows, _ := stdscr.MaxYX()

    for i, f := range files {
        s := fmt.Sprintf("%10d %s", f.Size(), f.Name())
        stdscr.MovePrint(i, 12, s)

        if i > rows - 5 {
            break
        }
    }

	for {
		gc.Update()
		ch := stdscr.GetChar()

		switch gc.KeyString(ch) {
		case "q":
			return
		}
	}
}
