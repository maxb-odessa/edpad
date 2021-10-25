package display

import (
	"edpad/cfg"
	"log"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Cmd struct {
	ViewPort string
	Command  int
	Data     string
}

const (
	VIEWPORT_TOP    = "top"
	VIEWPORT_CENTER = "center"
	VIEWPORT_BOTTOM = "bottom"
)

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

var viewPorts map[string]viewPort

const CHANSIZE = 128

var CmdChan chan *Cmd

func Start() error {

	runtime.LockOSThread()

	resPath := cfg.GtkResourcesDir

	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile(resPath + "./edpad.glade")
	if err != nil {
		return err
	}

	obj, err := builder.GetObject("window")
	if err != nil {
		return err
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()
	css.LoadFromPath(resPath + "./edpad.css")
	if err != nil {
		return err
	}

	ctx, _ := win.GetStyleContext()
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// prepare and configure view ports
	viewPorts = make(map[string]viewPort)
	for _, name := range []string{VIEWPORT_TOP, VIEWPORT_CENTER, VIEWPORT_BOTTOM} {

		var vp viewPort

		obj, err := builder.GetObject(name)
		if err != nil {
			return err
		}

		vp.view = obj.(*gtk.TextView)

		ctx, err := vp.view.GetStyleContext()
		if err != nil {
			return err
		}

		ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		vp.buff, err = vp.view.GetBuffer()
		if err != nil {
			return err
		}

		viewPortClear(&vp)
	}

	CmdChan = make(chan *Cmd, CHANSIZE)

	// start channels reader
	go cmdReader()

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
	return nil
}

func cmdReader() {

	for {
		select {
		case cmd, ok := <-CmdChan:
			if !ok {
				log.Fatalln("broken cmd chan")
			}
			glib.IdleAdd(func() bool { return processCmd(cmd) })
		}
	}
}

func processCmd(cmd *Cmd) bool {

	vp, ok := viewPorts[cmd.ViewPort]
	if !ok {
		log.Printf("unknown view port: %s\n", cmd.ViewPort)
		return false
	}

	switch cmd.Command {
	case CMD_CLEAR:
		viewPortClear(&vp)
	case CMD_TEXT:
		viewPortText(&vp, cmd.Data)
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
