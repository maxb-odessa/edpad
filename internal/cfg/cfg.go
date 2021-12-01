package cfg

import (
	"edpad/pkg/log"
	"os"

	o "github.com/pborman/getopt/v2"
)

var (
	GtkResourcesDir = os.Getenv("HOME") + "/.local/share/edpad/"
	FilePipe        string
	Listen          string = "0.0.0.0:55001"
	Debug           bool   = false
	Xdisplay        string
)

func Conf() error {

	help := false

	o.HelpColumn = 0
	o.FlagLong(&help, "help", 'h', "Show this help")
	o.FlagLong(&Debug, "debug", 'd', "Enable debug mode")
	o.FlagLong(&GtkResourcesDir, "gtk-resources-dir", 'r', "Path to GTK resource files dir")
	o.FlagLong(&FilePipe, "pipe", 'p', "Read ED journal entries from this pipe file")
	o.FlagLong(&Listen, "listen", 'l', "Read ED journal entries from TCP socket at this host:port")
	o.FlagLong(&Xdisplay, "xdisplay", 'x', "Send keystrokes to this (remote) DISPLAY")

	o.Parse()

	if help {
		o.Usage()
		os.Exit(0)
	}

	log.SetDebug(Debug)

	log.Debug("cfg: GtkResourcesDir = %s\n", GtkResourcesDir)
	log.Debug("cfg: FilePipe = %s\n", FilePipe)
	log.Debug("cfg: Listen = %s\n", Listen)
	log.Debug("cfg: Xdisplay = %s\n", Xdisplay)

	return nil
}
