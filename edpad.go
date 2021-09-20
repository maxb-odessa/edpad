package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var resPath = os.Getenv("HOME") + "/.local/share/edpad/"

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("path required\n")
		return
	}

	fp, err := os.OpenFile(os.Args[1], os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile(resPath + "./edpad.glade")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	obj, err := builder.GetObject("window")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	win := obj.(*gtk.Window)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()

	css.LoadFromPath(resPath + "./edpad.css")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	ctx, _ := win.GetStyleContext()
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	reader(fp, builder, css)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	win.Maximize()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

type ViewPort struct {
	vpname string
	view   *gtk.TextView
	buff   *gtk.TextBuffer
	mark   *gtk.TextMark
}

func reader(fp *os.File, builder *gtk.Builder, css *gtk.CssProvider) {

	viewPorts := map[string]*ViewPort{
		"view1": nil,
		"view2": nil,
		"view3": nil,
	}

	clearVp := func(vp *ViewPort) bool {
		/*
			iterS := vp.buff.GetStartIter()
			iterE := vp.buff.GetEndIter()
			vp.buff.Delete(iterS, iterE)
			iterE = vp.buff.GetEndIter()
			vp.buff.Insert(iterE, "")
		*/
		vp.buff.SetText("")
		if vp.mark == nil {
			iterE := vp.buff.GetEndIter()
			vp.mark = vp.buff.CreateMark(vp.vpname, iterE, false)
		}
		vp.view.ScrollToMark(vp.mark, 0.0, false, 0.0, 0.0)

		return false
	}

	for vpname, _ := range viewPorts {
		obj, err := builder.GetObject(vpname)
		if err != nil {
			continue
		}
		view := obj.(*gtk.TextView)

		ctx, _ := view.GetStyleContext()
		ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		viewPorts[vpname] = new(ViewPort)
		viewPorts[vpname].vpname = vpname

		viewPorts[vpname].view = view

		buff, _ := view.GetBuffer()
		viewPorts[vpname].buff = buff

		clearVp(viewPorts[vpname])
	}

	go func() {

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			var ok bool
			var vp *ViewPort

			// parse line, get view name
			tokens := fieldsN(scanner.Text(), 2)
			if vp, ok = viewPorts[tokens[0]]; !ok {
				fmt.Printf("invalid view '%s'\n", tokens[0])
				continue
			}
			text := tokens[1]
			text = strings.ReplaceAll(text, `\n`, "\n")
			text = strings.ReplaceAll(text, `\r`, "\r")
			text = strings.ReplaceAll(text, `\t`, "\t")
			text = strings.ReplaceAll(text, `\\`, "\\")

			strs := strings.Split(text, `\c`)

			doClear := len(strs) > 1
			for _, s := range strs {

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

		}

	}()

}

func fieldsN(str string, n int) []string {
	count := 0
	prevSep := false

	return strings.FieldsFunc(str, func(r rune) bool {
		if count >= n-1 {
			return false
		}
		if unicode.IsSpace(r) {
			if prevSep == false {
				count++
				prevSep = true
			}
			return true
		}
		prevSep = false
		return false
	})
}
