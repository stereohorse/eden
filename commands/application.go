package commands

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"

	st "github.com/stereohorse/eden/storage"
	u "github.com/stereohorse/eden/utils"
)

type ConsoleUI struct {
	views.BoxLayout

	app     *views.Application
	storage st.Storage

	docsList  *DocsList
	preview   *views.TextArea
	searchBox *SearchBox

	searchResult *st.Document
}

func (ui *ConsoleUI) Run() error {
	return ui.app.Run()
}

func NewConsoleUI(storage st.Storage) *ConsoleUI {
	box := views.NewBoxLayout(views.Vertical)

	titles := NewDocsList()
	box.AddWidget(titles, 1)

	preview := views.NewTextArea()
	box.AddWidget(preview, 2)

	searchBox := NewSearchBox()
	box.AddWidget(searchBox, 0)

	ui := &ConsoleUI{
		BoxLayout: *box,
		app:       &views.Application{},
		storage:   storage,
		docsList:  titles,
		preview:   preview,
		searchBox: searchBox,
	}

	searchBox.AddEventHandler(ui.handleSearchBox)
	ui.handleSearchBox("")

	ui.app.SetRootWidget(ui)

	return ui
}

func (ui *ConsoleUI) HandleEvent(ev tcell.Event) bool {
	t, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}

	switch t.Key() {
	case tcell.KeyEscape:
		ui.app.Quit()
	case tcell.KeyEnter:
		ui.searchResult = ui.docsList.GetSelection()
		ui.app.Quit()
	}

	return ui.BoxLayout.HandleEvent(ev)
}

func (ui *ConsoleUI) GetSearchResult() *st.Document {
	return ui.searchResult
}

func (ui *ConsoleUI) handleSearchBox(searchString string) error {
	hits, err := ui.storage.Recall(searchString)
	if err != nil {
		msg := "unable to recall " + searchString
		return u.NewError(msg, err)
	}

	docs := []st.Document{}
	for _, hit := range hits {
		docs = append(docs, hit.Doc)
	}

	ui.docsList.SetDocs(docs)
	return nil
}
