package commands

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
	st "github.com/stereohorse/eden/storage"
	u "github.com/stereohorse/eden/utils"
)

func getFromStorage(args []string, storage st.Storage) (err error) {
	//hits, err := storage.Recall(needle)
	//if err != nil {
	//msg := "unable to recall " + needle
	//return u.NewError(msg, err)
	//}

	//docs := []Document

	//for _, hit := range hits {
	//docs = append(docs, hit.Doc)
	//}

	app := &views.Application{}
	setupWidgets(app)

	if err = app.Run(); err != nil {
		return u.NewError("unable to run UI", err)
	}

	return nil
}

func setupWidgets(app *views.Application) {
	box := newBoxLayout(views.Vertical, app)
	app.SetRootWidget(box)

	titles := NewDocsList()
	box.AddWidget(titles, 1)

	preview := views.NewTextArea()
	preview.SetLines([]string{"preview", "is", "here"})
	box.AddWidget(preview, 2)

	searchBox := NewSearchBox()
	box.AddWidget(searchBox, 0)
}

type boxLayout struct {
	views.BoxLayout
	app *views.Application
}

func (self *boxLayout) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			self.app.Quit()
			return true
		}
	}
	return self.BoxLayout.HandleEvent(ev)
}

func newBoxLayout(orient views.Orientation, app *views.Application) *boxLayout {
	return &boxLayout{
		BoxLayout: *views.NewBoxLayout(orient),
		app:       app,
	}
}
