package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type ICmdConfig interface {
	constructor()
	init()
	createConfDir() error
	createConfFile() error
	addConfig()
	listConfig()
}

type CmdConfig struct {
	confDir  string
	confFile string
}

func (c *CmdConfig) constructor() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	c.confDir = filepath.Join(home, ".gotp")
	c.confFile = filepath.Join(c.confDir, "config.yml")
}

func (c *CmdConfig) init() {
	err := c.createConfDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(c.confFile)
	err = c.createConfFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func (c *CmdConfig) createConfDir() (err error) {
	err = os.Mkdir(c.confDir, 0755)
	if os.IsExist(err) {
		return nil
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create config dir: %v\n", err)
		return err
	}
	return nil
}

func (c *CmdConfig) createConfFile() (err error) {
	file, err := os.Create(c.confFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create config file: %v\n", err)
		return err
	}
	fmt.Println(file)
	return nil
}

func (c *CmdConfig) addConfig() {

}

func (c *CmdConfig) listConfig() {

}
