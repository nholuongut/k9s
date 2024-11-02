package views

import (
	"github.com/nholuongut/k9s/internal/resource"
	"github.com/nholuongut/k9s/internal/ui"
	"github.com/nholuongut/tview"
	"github.com/gdamore/tcell"
)

type selectList struct {
	*tview.List

	parent  loggable
	actions ui.KeyActions
}

func newSelectList(parent loggable) *selectList {
	v := selectList{List: tview.NewList(), actions: ui.KeyActions{}}
	{
		v.parent = parent
		v.SetBorder(true)
		v.SetMainTextColor(tcell.ColorWhite)
		v.ShowSecondaryText(false)
		v.SetShortcutColor(tcell.ColorAqua)
		v.SetSelectedBackgroundColor(tcell.ColorAqua)
		v.SetTitle(" [aqua::b]Container Selector ")
		v.SetInputCapture(func(evt *tcell.EventKey) *tcell.EventKey {
			if a, ok := v.actions[evt.Key()]; ok {
				a.Action(evt)
				evt = nil
			}
			return evt
		})
	}

	return &v
}

func (v *selectList) back(evt *tcell.EventKey) *tcell.EventKey {
	v.parent.switchPage("master")

	return nil
}

// Protocol...

func (v *selectList) switchPage(p string) {
	v.parent.switchPage(p)
}

func (v *selectList) getList() resource.List {
	return v.parent.getList()
}

func (v *selectList) getSelection() string {
	return v.parent.getSelection()
}

// SetActions to handle keyboard events.
func (v *selectList) setActions(aa ui.KeyActions) {
	v.actions = aa
}

func (v *selectList) Hints() ui.Hints {
	if v.actions != nil {
		return v.actions.Hints()
	}

	return nil
}

func (v *selectList) populate(ss []string) {
	v.Clear()
	for i, s := range ss {
		v.AddItem(s, "Select a container", rune('a'+i), nil)
	}
}
