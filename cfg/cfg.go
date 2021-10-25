package cfg

import (
	o "github.com/pborman/getopt/v2"
)

var (
	// DEBUG! GtkResourcesDir = os.Getenv("HOME") + "/.local/share/edpad/"
	GtkResourcesDir = "./resources/"
	FilePipe        = "/tmp/edpad.pipe"
)

func Conf() error {

	o.HelpColumn = 0
	o.FlagLong(&GtkResourcesDir, "gtk-resources-dir", 'r', "Path to GTK resource files dir")
	o.FlagLong(&FilePipe, "pipe", 'p', "Path to pipe file to read journal entries")

	o.Parse()

	return nil
}
