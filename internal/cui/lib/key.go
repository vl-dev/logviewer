package lib

import "github.com/jroimartin/gocui"

func keyName(key interface{}) string {
	if k, ok := key.(rune); ok {
		return string(k)
	}
	switch key.(gocui.Key) {
	case gocui.KeyF1:
		return "F1"
	case gocui.KeyF2:
		return "F2"
	case gocui.KeyF3:
		return "F3"
	case gocui.KeyF4:
		return "F4"
	case gocui.KeyF5:
		return "F5"
	case gocui.KeyF6:
		return "F6"
	case gocui.KeyF7:
		return "F7"
	case gocui.KeyF8:
		return "F8"
	case gocui.KeyF9:
		return "F9"
	case gocui.KeyF10:
		return "F10"
	case gocui.KeyF11:
		return "F11"
	case gocui.KeyF12:
		return "F12"
	case gocui.KeyInsert:
		return "Insert"
	case gocui.KeyDelete:
		return "⌦"
	case gocui.KeyHome:
		return "Home"
	case gocui.KeyEnd:
		return "End"
	case gocui.KeyPgup:
		return "PgUp"
	case gocui.KeyPgdn:
		return "PgDn"
	case gocui.KeyArrowUp:
		return "↑"
	case gocui.KeyArrowDown:
		return "↓"
	case gocui.KeyArrowLeft:
		return "←"
	case gocui.KeyArrowRight:
		return "→"
	case gocui.MouseLeft:
		return "Mouse️L"
	case gocui.MouseMiddle:
		return "MouseM"
	case gocui.MouseRight:
		return "MouseR"
	case gocui.MouseRelease:
		return "MouseRelease"
	case gocui.MouseWheelUp:
		return "Mouse◌↑"
	case gocui.MouseWheelDown:
		return "Mouse◌↓"
	case gocui.KeyCtrlSpace:
		return "⌃Space"
	case gocui.KeyCtrlA:
		return "⌃A"
	case gocui.KeyCtrlB:
		return "⌃B"
	case gocui.KeyCtrlC:
		return "⌃C"
	case gocui.KeyCtrlD:
		return "⌃D"
	case gocui.KeyCtrlE:
		return "⌃E"
	case gocui.KeyCtrlF:
		return "⌃F"
	case gocui.KeyCtrlG:
		return "⌃G"
	case gocui.KeyBackspace:
		return "⌫"
	case gocui.KeyTab:
		return "Tab"
	case gocui.KeyCtrlK:
		return "⌃K"
	case gocui.KeyCtrlL:
		return "⌃L"
	case gocui.KeyEnter:
		return "↩"
	case gocui.KeyCtrlN:
		return "⌃N"
	case gocui.KeyCtrlO:
		return "⌃O"
	case gocui.KeyCtrlP:
		return "⌃P"
	case gocui.KeyCtrlQ:
		return "⌃Q"
	case gocui.KeyCtrlR:
		return "⌃R"
	case gocui.KeyCtrlS:
		return "⌃S"
	case gocui.KeyCtrlT:
		return "⌃T"
	case gocui.KeyCtrlU:
		return "⌃U"
	case gocui.KeyCtrlV:
		return "⌃V"
	case gocui.KeyCtrlW:
		return "⌃W"
	case gocui.KeyCtrlX:
		return "⌃X"
	case gocui.KeyCtrlY:
		return "⌃Y"
	case gocui.KeyCtrlZ:
		return "⌃Z"
	case gocui.KeyEsc:
		return "Esc"
	case gocui.KeyCtrlBackslash:
		return "⌃\\"
	case gocui.KeyCtrl5:
		return "⌃5"
	case gocui.KeyCtrlUnderscore:
		return "⌃_"
	case gocui.KeySpace:
		return "Space"
	case gocui.KeyBackspace2:
		return "⌫"
	}
	panic("unknown key")
}