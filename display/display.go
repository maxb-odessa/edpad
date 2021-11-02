package display

import (
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"edpad/cfg"
	"edpad/log"
)

type Cmd struct {
	Command int
	Data    string
}

const (
	CMD_CLEAR = iota
	CMD_TEXT
	CMD_SCROLL_UP
	CMD_SCROLL_DOWN
)

type viewPort struct {
	view *gtk.TextView
	buff *gtk.TextBuffer
	mark *gtk.TextMark
}

func Start(cmdCh chan *Cmd) error {

	runtime.LockOSThread()

	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile(cfg.GtkResourcesDir + "/edpad.glade")
	if err != nil {
		return err
	}

	obj, err := builder.GetObject("window")
	if err != nil {
		return err
	}

	win := obj.(*gtk.ApplicationWindow)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()
	css.LoadFromPath(cfg.GtkResourcesDir + "/edpad.css")
	if err != nil {
		return err
	}

	ctx, _ := win.GetStyleContext()
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// prepare and configure view ports
	var vp viewPort

	obj, err = builder.GetObject("textview")
	if err != nil {
		return err
	}

	vp.view = obj.(*gtk.TextView)

	ctx, err = vp.view.GetStyleContext()
	if err != nil {
		return err
	}

	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	vp.buff, err = vp.view.GetBuffer()
	if err != nil {
		return err
	}

	viewPortClear(&vp)

	// start channels reader
	go cmdReader(&vp, cmdCh)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
	return nil
}

func cmdReader(vp *viewPort, cmdCh chan *Cmd) {

	for {
		select {
		case cmd, ok := <-cmdCh:
			if !ok {
				log.Fatal("broken cmd chan")
			}
			glib.IdleAdd(func() bool { return processCmd(vp, cmd) })
		}
	}
}

func processCmd(vp *viewPort, cmd *Cmd) bool {

	switch cmd.Command {
	case CMD_CLEAR:
		viewPortClear(vp)
	case CMD_TEXT:
		viewPortText(vp, cmd.Data)
	case CMD_SCROLL_UP:
	case CMD_SCROLL_DOWN:
	}

	return false
}

func viewPortText(vp *viewPort, text string) bool {
	vp.buff.InsertMarkup(vp.buff.GetEndIter(), text)
	vp.view.ScrollToIter(vp.buff.GetEndIter(), 0.0, false, 0.0, 0.0)
	return false
}

func viewPortClear(vp *viewPort) bool {
	vp.buff.SetText("")
	//	vp.view.ScrollToIter(vp.buff.GetEndIter, 0.0, false, 0.0, 0.0)
	return false
}
