package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ICmdConfig interface {
	constructor()
	init()
	createConfDir() error
	createConfFile() error
	addConfig()
	listConfig() error
	getSecretFromName(string) (string error)
	readConfig() ([]map[interface{}]interface{}, error)
}

type CmdConfig struct {
	confDir  string
	confFile string
}

type ConfContents struct {
	Name   string
	Secret string
}

var Permission os.FileMode = 0755

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
	// check exists config dir/file
	if _, err := os.Stat(c.confDir); os.IsNotExist(err) {
		// path/to/whatever exists
		err := c.createConfDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stdout, "Config directory is already exists\n")
	}

	if _, err := os.Stat(c.confFile); os.IsNotExist(err) {
		err = c.createConfFile()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created config directory and file: %v\n", c.confFile)
	} else {
		fmt.Fprintf(os.Stdout, "Config file is already exists\n")
	}

}

func (c *CmdConfig) createConfDir() (err error) {
	err = os.Mkdir(c.confDir, Permission)
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
	_, err = os.Create(c.confFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create config file: %v\n", err)
		return err
	}
	return nil
}

func (c *CmdConfig) addConfig() (err error) {
	fmt.Println("Initializing your config og gotp")

	// set key name
	fmt.Print("Input key name: ")
	scannerName := bufio.NewScanner(os.Stdin)
	var scannedName string
	for scannerName.Scan() {
		// check same config is not stored
		scannedName = scannerName.Text()
		fmt.Println(scannedName)
		break
	}

	// set secret key
	fmt.Print("Input secret key: ")
	scannerKey := bufio.NewScanner(os.Stdin)
	var scannedKey string
	for scannerKey.Scan() {
		// check same config is not stored
		scannedKey = scannerKey.Text()
		fmt.Println(scannedKey)
		break
	}

	// save name and key in config file
	file, err := os.OpenFile(c.confFile, os.O_WRONLY|os.O_APPEND, Permission)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't append to config file: %v\n", err)
		return err
	}
	defer file.Close()

	var confContents ConfContents
	confContents.Name = scannedName
	confContents.Secret = scannedKey

	buf, err := yaml.Marshal(confContents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't marshal conf value: %v\n", err)
		return err
	}
	fmt.Fprintln(file, "---")
	fmt.Fprint(file, string(buf))

	return nil
}

func (c *CmdConfig) listConfig() (err error) {
	confList, err := c.readConfig()
	if err != nil {
		return err
	}

	for i, v := range confList {
		fmt.Println(i, v["name"])
	}

	return nil
}

func (c *CmdConfig) getSecretFromName(name string) (secret string, err error) {
	confList, err := c.readConfig()
	if err != nil {
		return "", err
	}
	// fmt.Println(confList)
	for _, v := range confList {
		if name == v["name"] {
			return v["secret"].(string), nil
		}
	}
	fmt.Fprintln(os.Stderr, "Can't find that name in config file")
	return "", errors.New(fmt.Sprintln("Value error"))
}

// confを読んでリスト形式の一覧を返す
func (c *CmdConfig) readConfig() (list []map[interface{}]interface{}, err error) {
	f, err := os.Open(c.confFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read config file: %v\n", err)
		return nil, err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)

	var confList []map[interface{}]interface{}
	var tmp map[interface{}]interface{}
	for (dec.Decode(&tmp)) == nil {
		// fmt.Printf("%v\n", tmp)
		confList = append(confList, tmp)
		tmp = nil
	}

	return confList, nil
}
