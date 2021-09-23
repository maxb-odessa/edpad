package display

import (
	"edpad/conf"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Cmd struct {
	ViewPortName string
	Command      string
	Data         string
}

type viewPort struct {
	view *gtk.TextView
	buff *gtk.TextBuffer
	mark *gtk.TextMark
}

var viewPorts map[string]viewPort

const CHANSIZE = 128

var InChan chan *Cmd

func Start(cfg *conf.Conf) {

	resPath, _ := cfg.Get("gtk_resources_dir")

	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile(resPath + "./edpad.glade")
	if err != nil {
		log.Println(err)
		return
	}

	obj, err := builder.GetObject("window")
	if err != nil {
		log.Println(err)
		return
	}

	win := obj.(*gtk.Window)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()

	css.LoadFromPath(resPath + "./edpad.css")
	if err != nil {
		log.Println(err)
		return
	}

	ctx, _ := win.GetStyleContext()
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// prepare and configure view ports
	viewPorts = make(map[string]viewPort)
	for _, name := range []string{"top", "center", "bottom"} {

		var vp viewPort

		obj, err := builder.GetObject(name)
		if err != nil {
			log.Fatalln(err)
		}

		vp.view = obj.(*gtk.TextView)

		ctx, err := vp.view.GetStyleContext()
		if err != nil {
			log.Fatalln(err)
		}

		ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		vp.buff, err = vp.view.GetBuffer()
		if err != nil {
			log.Fatalln(err)
		}

		clearViewPort(&vp)
	}

	InChan = make(chan *Cmd, CHANSIZE)

	// start channels reader
	go reader()

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func reader() {

	for {
		select {
		case cmd, ok := <-InChan:
			if !ok {
				log.Panic("broken chan")
			}
			process(cmd)
		}
	}
}

func process(cmd *Cmd) {

	glib.IdleAdd(func() bool {
		vp, ok := viewPorts[cmd.ViewPortName]
		if ok {
			vp.buff.InsertMarkup(vp.buff.GetEndIter(), cmd.Data)
			vp.view.ScrollToIter(vp.buff.GetEndIter(), 0.0, false, 0.0, 0.0)
		}
		return false
	})

}

func clearViewPort(vp *viewPort) bool {
	vp.buff.SetText("")
	//	vp.view.ScrollToIter(vp.buff.GetEndIter, 0.0, false, 0.0, 0.0)
	return false
}
