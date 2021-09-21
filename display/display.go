package display

import (
	"edpad/conf"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type Cmd struct {
	ViewPortName string
	Command      string
	Data         string
}

type viewPort struct {
	name string
	view *gtk.TextView
	buff *gtk.TextBuffer
	mark *gtk.TextMark
}

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
	viewPorts := map[string]*viewPort{
		"vp0": nil,
		"vp1": nil,
		"vp2": nil,
		"vp3": nil,
		"vp4": nil,
		"vp5": nil,
		"vp6": nil,
		"vp7": nil,
		"vp8": nil,
		"vp9": nil,
	}

	for name, _ := range viewPorts {

		obj, err := builder.GetObject(name)
		if err != nil {
			log.Printf("not preparing view port: %s\n", err)
			continue
		}

		view := obj.(*gtk.TextView)

		ctx, err := view.GetStyleContext()
		if err != nil {
			log.Println(err)
			return
		}

		ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		viewPorts[name] = new(viewPort)
		viewPorts[name].name = name

		viewPorts[name].view = view

		buff, err := view.GetBuffer()
		if err != nil {
			log.Println(err)
			return
		}

		viewPorts[name].buff = buff

		clearViewPort(viewPorts[name])
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

	/*
			if len(s) > 0 {
				glib.IdleAdd(func() bool {
					iterE := vp.buff.GetEndIter()
					vp.buff.Insert(iterE, s)
					vp.view.ScrollToMark(vp.mark, 0.0, false, 0.0, 0.0)
					return false
				})

			}

			if doClear {
				glib.IdleAdd(func() bool {
					clearVp(vp)
					return false
				})
			}
		}
	*/
}

func clearViewPort(vp *viewPort) bool {
	vp.buff.SetText("")
	if vp.mark == nil {
		iterE := vp.buff.GetEndIter()
		vp.mark = vp.buff.CreateMark(vp.name, iterE, false)
	}

	vp.view.ScrollToMark(vp.mark, 0.0, false, 0.0, 0.0)

	return false
}
