package commands

import (
	"bytes"
	"log"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type SearchBox struct {
	views.WidgetWatchers

	view          views.View
	searchString  bytes.Buffer
	eventHandlers []SearchBoxEventHandler
}

type SearchBoxEventHandler func(searchString string) error

func (sb *SearchBox) AddEventHandler(eh SearchBoxEventHandler) {
	sb.eventHandlers = append(sb.eventHandlers, eh)
}

func (sb *SearchBox) Draw() {
	var styl tcell.Style
	var comb []rune

	for i, r := range sb.searchString.String() {
		sb.view.SetContent(i, 0, r, comb, styl)
	}
}

func (sb *SearchBox) GetSearchString() string {
	return sb.searchString.String()
}

func (sb *SearchBox) Resize() {
	sb.PostEventWidgetResize(sb)
}

func (sb *SearchBox) HandleEvent(ev tcell.Event) bool {
	t, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	switch t.Key() {
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		len := sb.searchString.Len()
		if len != 0 {
			sb.searchString.Truncate(len - 1)
		}
	case tcell.KeyRune:
		sb.searchString.WriteRune(t.Rune())
	default:
		return false
	}

	if len(sb.eventHandlers) != 0 {
		for _, eh := range sb.eventHandlers {
			if err := eh(sb.searchString.String()); err != nil {
				log.Print(err)
			}
		}
	}

	return true
}

func (sb *SearchBox) SetView(view views.View) {
	sb.view = view
}

func (sb *SearchBox) Size() (int, int) {
	return sb.searchString.Len(), 1
}

func NewSearchBox() *SearchBox {
	return &SearchBox{}
}
