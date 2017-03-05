package main

import (
    gc "github.com/rthornton128/goncurses"
    // "fmt"
    "log"
)

func main() {
    var b1, b2 Buffer
    b1.path = "/home"
    b2.path = "/"

    AddPathsToBuffer(&b1)
    AddPathsToBuffer(&b2)

    LeftWindow.buffer = &b1
    RightWindow.buffer = &b2

    ActiveWindow = &LeftWindow

    InitTerm()

    gc.Echo(false)
    gc.Cursor(0)
    defer gc.End()

    if err := gc.StartColor(); err != nil {
		log.Fatal(err)
	}

    gc.InitPair(1, gc.C_WHITE, gc.C_BLACK)
    gc.InitPair(2, gc.C_WHITE, gc.C_BLUE)
    gc.InitPair(3, gc.C_BLUE, gc.C_BLACK) // Dirs

    for {
        PrintWindow(&LeftWindow)
        PrintWindow(&RightWindow)

        gc.Update()

        ch := Term.win.GetChar()

        switch ch {
        case 'q':
            return
        case 'n':
			MoveDown(ActiveWindow)
        case 'p':
            MoveUp(ActiveWindow)
        case 'e':
            UpdateBufferPath(ActiveWindow.buffer)
            AddPathsToBuffer(ActiveWindow.buffer)
        case 'j':
            Jump(ActiveWindow.buffer)
            AddPathsToBuffer(ActiveWindow.buffer)
        case 'o':
            SwitchWindows(ActiveWindow)
		case 'r':
			MoveToCentre(ActiveWindow.buffer)
		case 'V':
			MoveUpPage(ActiveWindow)
        case 'v':
            MoveDownPage(ActiveWindow)
        }
    }
}
