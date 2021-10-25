package cfg

import (
	"os"

	o "github.com/pborman/getopt/v2"
)

var (
	// DEBUG! GtkResourcesDir = os.Getenv("HOME") + "/.local/share/edpad/"
	GtkResourcesDir string = "./resources/"
	FilePipe        string = "/tmp/edpad.pipe"
	Listen          string = "0.0.0.0:12345"
)

func Conf() error {

	help := false

	o.HelpColumn = 0
	o.FlagLong(&help, "help", 'h', "Show this help")
	o.FlagLong(&GtkResourcesDir, "gtk-resources-dir", 'r', "Path to GTK resource files dir")
	o.FlagLong(&FilePipe, "pipe", 'p', "Path to pipe file to read journal entries")
	o.FlagLong(&Listen, "listen", 'l', "Listen for TCP connection at this host:port")

	o.Parse()

	if help {
		o.Usage()
		os.Exit(1)
	}

	return nil
}
