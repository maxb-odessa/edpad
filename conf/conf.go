package conf

import (
	"fmt"
	"os"
)

type Conf struct {
	vars map[string]string
}

func Read(path string) (*Conf, error) {
	cfg := new(Conf)
	cfg.vars = make(map[string]string)

	// some defaults
	cfg.vars["gtk_resources_dir"] = os.Getenv("HOME") + "/.local/share/edpad/"
	cfg.vars["pipe_path"] = "/tmp/ed.pipe"

	// TODO read config here

	cfg.vars["gtk_resources_dir"] = "./resources/" // TODO temp value!

	return cfg, nil
}

func (self *Conf) Get(name string) (string, error) {
	if value, ok := self.vars[name]; ok {
		return value, nil
	}
	return "", fmt.Errorf("variable '%s' is not configured", name)
}

func (self *Conf) Set(name, value string, overwrite bool) error {
	if oldValue, exists := self.vars[name]; exists && !overwrite {
		return fmt.Errorf("variable '%s' already set to '%s'", name, oldValue)
	}
	self.vars[name] = value
	return nil
}
