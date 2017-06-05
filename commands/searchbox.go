package commands

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type SearchBox struct {
	views.WidgetWatchers

	view views.View
}

func (sb *SearchBox) Draw() {
	var styl tcell.Style
	var comb []rune

	sb.view.SetContent(0, 0, 's', comb, styl)
}

func (sb *SearchBox) Resize() {
	sb.PostEventWidgetResize(sb)
}

func (*SearchBox) HandleEvent(ev tcell.Event) bool {
	return false
}

func (sb *SearchBox) SetView(view views.View) {
	sb.view = view
}

func (sb *SearchBox) Size() (int, int) {
	return 10, 1
}

func NewSearchBox() *SearchBox {
	return &SearchBox{}
}
