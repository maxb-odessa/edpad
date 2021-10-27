package cfg

import (
	"os"

	o "github.com/pborman/getopt/v2"
)

var (
	// DEBUG! GtkResourcesDir = os.Getenv("HOME") + "/.local/share/edpad/"
	GtkResourcesDir string = "./resources/"
	FilePipe        string
	Listen          string = "0.0.0.0:55001"
	Debug           bool   = false
)

func Conf() error {

	help := false

	o.HelpColumn = 0
	o.FlagLong(&help, "help", 'h', "Show this help")
	o.FlagLong(&Debug, "debug", 'd', "Enable debug mode")
	o.FlagLong(&GtkResourcesDir, "gtk-resources-dir", 'r', "Path to GTK resource files dir")
	o.FlagLong(&FilePipe, "pipe", 'p', "Read ED journal entries from this pipe file")
	o.FlagLong(&Listen, "listen", 'l', "Read ED journal entries from TCP socket at this host:port")

	o.Parse()

	if help {
		o.Usage()
		os.Exit(0)
	}

	return nil
}
