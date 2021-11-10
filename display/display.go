package display

import (
	"runtime"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"edpad/cfg"
	"edpad/event"
	"edpad/log"
	"edpad/xdo"
)

type Data struct {
	Id   string
	Text string
}

type prop struct {
	pos     int  // position in textBuf array
	glue    bool // concat with existing string or replace
	clear   bool // clear whole textBuf on this event
	persist bool // this textBuf entry must survive 'clear' event
}

var props = map[event.Type]prop{
	event.FSD_TARGET:   prop{pos: 0, glue: false, clear: false, persist: true},
	event.START_JUMP:   prop{pos: 1, glue: false, clear: true, persist: true},
	event.SYSTEM_NAME:  prop{pos: 1, glue: false, clear: false, persist: false},
	event.BODY_SIGNALS: prop{pos: 2, glue: false, clear: false, persist: false},
	event.MAIN_STAR:    prop{pos: 3, glue: false, clear: false, persist: false},
	event.SEC_STAR:     prop{pos: 4, glue: true, clear: false, persist: false},
	event.PLANET:       prop{pos: 5, glue: true, clear: false, persist: false},
	event.FSS_SIGNALS:  prop{pos: 6, glue: true, clear: false, persist: false},
}

var textBuf [8]string

type viewPort struct {
	view *gtk.TextView
	buff *gtk.TextBuffer
	mark *gtk.TextMark
}

func Start(eventCh chan *event.Event) error {

	runtime.LockOSThread()

	gtk.Init(nil)

	// NOTE: gtk_builder_new_from_file() always aborts on any error, thus tracking
	// returned error makes no sense
	builder, err := gtk.BuilderNewFromFile(cfg.GtkResourcesDir + "/edpad.glade")
	if err != nil {
		return err
	}

	if remoteDisplay := xdo.New(cfg.Xdisplay); remoteDisplay != nil {

		onButtonDownFunc := func(s interface{}) {
			if name, err := s.(*gtk.Button).GetName(); err == nil {
				log.Debug("button down: %s\n", name)
				remoteDisplay.KeyDown(name, 1000)
			}
		}

		onButtonUpFunc := func(s interface{}) {
			if name, err := s.(*gtk.Button).GetName(); err == nil {
				log.Debug("button up: %s\n", name)
				remoteDisplay.KeyUp(name, 1000)
			}
		}

		signals := map[string]interface{}{
			"onPress":   onButtonDownFunc,
			"onRelease": onButtonUpFunc,
		}

		builder.ConnectSignals(signals)

	} else if cfg.Xdisplay != "" {
		log.Warn("Failed to connect to '%s', keypad is disabled\n", cfg.Xdisplay)
	}

	winObj, err := builder.GetObject("window")
	if err != nil {
		return err
	}

	win := winObj.(*gtk.ApplicationWindow)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()
	css.LoadFromPath(cfg.GtkResourcesDir + "/edpad.css")
	if err != nil {
		return err
	}
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_USER)

	obj, err := builder.GetObject("textview")
	if err != nil {
		return err
	}

	var vp viewPort
	vp.view = obj.(*gtk.TextView)
	vp.buff, err = vp.view.GetBuffer()
	if err != nil {
		return err
	}

	viewPortClear(&vp)

	// start events reader
	go eventReader(&vp, eventCh)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
	return nil
}

func eventReader(vp *viewPort, eventCh chan *event.Event) {

	for {
		select {
		case ev, ok := <-eventCh:
			if !ok {
				log.Fatal("broken event chan")
			}
			glib.IdleAdd(func() bool { return processEvent(vp, ev) })
		}
	}
}

func processEvent(vp *viewPort, ev *event.Event) (res bool) {
	res = false

	evProp, ok := props[ev.Type]
	if !ok {
		log.Err("unknow data id '%s'\n", ev.Type)
		return
	}

	viewPortClear(vp)

	if evProp.clear {

		for _, pr := range props {
			if !pr.persist {
				textBuf[pr.pos] = ""
			}
		}

	}

	if evProp.glue {
		if textBuf[evProp.pos] != "" {
			textBuf[evProp.pos] += "\n"
		}
		textBuf[evProp.pos] += ev.Text
	} else {
		textBuf[evProp.pos] = ev.Text
	}

	text := strings.Join(textBuf[:], "\n")
	viewPortText(vp, text)

	return
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
