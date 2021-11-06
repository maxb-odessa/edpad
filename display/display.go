package display

import (
	"runtime"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"edpad/cfg"
	"edpad/log"
)

type Data struct {
	Idx    int
	Single bool
	Text   string
}

var dataBuf [16]string

const (
	CURRENT_SYSTEM = 0
	NEXT_JUMP      = 1
	BODY_SIGNALS   = 2
	MAIN_STAR      = 3
	SECONDARY_STAR = 4
	OTHER_SIGNALS  = 5
)

type viewPort struct {
	view *gtk.TextView
	buff *gtk.TextBuffer
	mark *gtk.TextMark
}

func Start(dataCh chan *Data) error {

	runtime.LockOSThread()

	gtk.Init(nil)

	// NOTE: gtk_builder_new_from_file() always aborts on any error, thus tracking
	// returned error makes no sense
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

	//ctx.AddProviderForScreen(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_USER)

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

	//ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_USER)

	vp.buff, err = vp.view.GetBuffer()
	if err != nil {
		return err
	}

	viewPortClear(&vp)

	// start channels reader
	go dataReader(&vp, dataCh)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
	return nil
}

func dataReader(vp *viewPort, dataCh chan *Data) {

	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				log.Fatal("broken data chan")
			}
			glib.IdleAdd(func() bool { return processData(vp, data) })
		}
	}
}

func processData(vp *viewPort, data *Data) bool {

	// put data into appropriate buf[] slot

	viewPortClear(vp)
	buf := dataBuf[0]
	viewPortText(vp, buf)

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
