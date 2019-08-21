package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var appMenu = []string{
	"1 - Demo",
	//"2 - Blah...",
	"0 - Exit",
}

type myapp struct {
	app *tview.Application
	// TODO: add more what you need here
}

func initApp() {
	a := &myapp{
		app: tview.NewApplication(),
	}

	//box := tview.NewBox().SetBorder(true).SetTitle("  Press CTRL-C to exit  ")
	dropdown := tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions(appMenu, a.onAppChoice)

	a.app.SetRoot(dropdown, true)

	if err := a.app.Run(); err != nil {
		panic("FAILED TO RUN: " + err.Error())
	}
}

func main() {
	initApp()
}

func (a *myapp) onAppChoice(menu string, idx int) {
	switch menu[:1] {
	case "1":
		// TODO: handle
		fmt.Printf("Your choice is %s [%d]", menu, idx)
	case "0":
		a.app.Stop()
		return
	}
}
