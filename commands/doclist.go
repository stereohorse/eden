package commands

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
	st "github.com/stereohorse/eden/storage"
)

type DocsList struct {
	views.WidgetWatchers

	view views.View

	docs        []st.Document
	cursorIndex int
}

func (d *DocsList) SetDocs(docs []st.Document) {
	d.docs = docs

	if d.cursorIndex >= len(docs) {
		d.cursorIndex = len(docs) - 1
	}

	d.PostEventWidgetContent(d)
}

func (d *DocsList) SetView(view views.View) {
	d.view = view
}

func (d *DocsList) Draw() {
	var styl tcell.Style
	var comb []rune

	d.view.SetContent(1, 2, 'i', comb, styl)
}

func (d *DocsList) Resize() {
	d.PostEventWidgetResize(d)
}

func (d *DocsList) HandleEvent(tcell.Event) bool {
	return false
}

func (d *DocsList) Size() (int, int) {
	return 10, 1
}

func NewDocsList() *DocsList {
	return &DocsList{
		cursorIndex: 0,
	}
}
