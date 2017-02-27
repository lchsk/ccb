package main

import (
    gc "github.com/rthornton128/goncurses"
)

func main() {
    var b1, b2 Buffer
    b1.path = "/"
    b2.path = "/"

    AddPathsToBuffer(&b1)
    AddPathsToBuffer(&b2)

    LeftWindow.buffer = &b1
    RightWindow.buffer = &b2

    InitTerm()

    PrintWindow(&LeftWindow)
    PrintWindow(&RightWindow)

    for {
        gc.Update()

        switch Term.win.GetChar() {
        case 'q':
            return
        }
    }
}
