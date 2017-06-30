package commands

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
	st "github.com/stereohorse/eden/storage"
	"log"
)

type DocsList struct {
	views.WidgetWatchers

	view views.View

	docs        []st.Document
	cursorIndex int

	onDelete OnDelete
}

type OnDelete func(doc *st.Document) error

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
	case tcell.KeyCtrlD:
		selection := d.GetSelection()
		if d.onDelete != nil && selection != nil {
			if err := d.onDelete(selection); err != nil {
				log.Println("unable to delete file: {}", err)
			} else {
				d.docs = append(d.docs[:d.cursorIndex], d.docs[d.cursorIndex+1:]...)
				d.cursorIntoBounds()
			}
		}
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

func (d *DocsList) SetOnDelete(onDelete OnDelete) {
	d.onDelete = onDelete
}

func NewDocsList() *DocsList {
	return &DocsList{
		cursorIndex: 0,
	}
}
