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
	d.cursorIntoBounds()

	d.PostEventWidgetContent(d)
}

func (d *DocsList) SetView(view views.View) {
	d.view = view
}

func (d *DocsList) Draw() {
	var comb []rune
	stdStyle := tcell.StyleDefault

	for i, doc := range d.docs {
		if i == d.cursorIndex {
			d.view.SetContent(0, i, '>', comb, stdStyle)
		} else {
			d.view.SetContent(0, i, ' ', comb, stdStyle)
		}

		for k, r := range doc.Body {
			d.view.SetContent(k+2, i, r, comb, stdStyle)
		}
	}
}

func (d *DocsList) Resize() {
	d.PostEventWidgetResize(d)
}

func (d *DocsList) GetSelection() *st.Document {
	if d.cursorIndex == -1 {
		return nil
	}

	return &d.docs[d.cursorIndex]
}

func (d *DocsList) HandleEvent(ev tcell.Event) bool {
	t, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	switch t.Key() {
	case tcell.KeyCtrlK:
		d.cursorUp()
	case tcell.KeyCtrlJ:
		d.cursorDown()
	default:
		return false
	}

	return true
}

func (d *DocsList) Size() (int, int) {
	return 10, 1
}

func (d *DocsList) cursorUp() {
	if d.cursorIndex <= 0 {
		return
	}

	d.cursorIndex--
}

func (d *DocsList) cursorDown() {
	if d.cursorIndex == (len(d.docs) - 1) {
		return
	}

	d.cursorIndex++
}

func (d *DocsList) cursorIntoBounds() {
	if len(d.docs) == 0 {
		d.cursorIndex = -1
	} else {
		if d.cursorIndex >= (len(d.docs) - 1) {
			d.cursorIndex = len(d.docs) - 1
		} else if d.cursorIndex < 0 {
			d.cursorIndex = 0
		}
	}
}

func NewDocsList() *DocsList {
	return &DocsList{
		cursorIndex: 0,
	}
}
