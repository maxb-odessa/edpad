package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	builder, _ := gtk.BuilderNewFromFile("./edpad.glade")

	obj, err := builder.GetObject("window")
	if err != nil {
		// object not found
		return
	}
	win := obj.(*gtk.Window)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, _ := gtk.CssProviderNew()
	e := css.LoadFromPath("./edpad.css")
	fmt.Printf("css err: %v\n", e)

	ctx, _ := win.GetStyleContext()
	ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	go txt(builder, css)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

type ViewPort struct {
	view  *gtk.TextView
	buff  *gtk.TextBuffer
	iterS *gtk.TextIter
	iterE *gtk.TextIter
}

func txt(builder *gtk.Builder, css *gtk.CssProvider) {

	viewPorts := map[string]*ViewPort{
		"view1": nil,
		"view2": nil,
		"view3": nil,
	}

	for idx, _ := range viewPorts {
		obj, err := builder.GetObject(idx)
		if err != nil {
			continue
		}
		view := obj.(*gtk.TextView)

		ctx, _ := view.GetStyleContext()
		ctx.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		viewPorts[idx] = new(ViewPort)

		viewPorts[idx].view = view

		buff, _ := view.GetBuffer()
		viewPorts[idx].buff = buff

		viewPorts[idx].iterS = buff.GetStartIter()
		viewPorts[idx].iterE = buff.GetEndIter()

		viewPorts[idx].buff.Delete(viewPorts[idx].iterS, viewPorts[idx].iterE)
	}

	fp, _ := os.OpenFile("/tmp/edpad.pipe", os.O_RDWR, 0666)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// parse line, get view name
		tokens := fieldsN(scanner.Text(), 2)
		if _, ok := viewPorts[tokens[0]]; !ok {
			fmt.Printf("invalid view '%s'\n", tokens[0])
			continue
		}
		idx := tokens[0]
		text := tokens[1]
		text = strings.ReplaceAll(text, `\n`, "\n")
		text = strings.ReplaceAll(text, `\r`, "\r")
		text = strings.ReplaceAll(text, `\t`, "\t")
		text = strings.ReplaceAll(text, `\\`, "\\")

		viewPorts[idx].buff.Insert(viewPorts[idx].iterE, text)
		//		viewPorts[idx].iterS = viewPorts[idx].buff.GetStartIter()
		//		viewPorts[idx].iterE = viewPorts[idx].buff.GetEndIter()

		viewPorts[idx].view.ScrollToIter(viewPorts[idx].iterE, 0.0, false, 0.0, 0.0)
	}

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
