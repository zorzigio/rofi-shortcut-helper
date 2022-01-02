package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/mattn/go-shellwords"
)

// Variables definition
var rofiCmd *string

// Create interface for the json file
type File map[string]Program
type Program struct {
	Name     string `json:"name"`
	Commands []struct {
		Key     string `json:"key"`
		Command string `json:"command"`
	} `json:"commands"`
}

func main() {
	// Create new parser object
	parser := argparse.NewParser("rofi-help-shortcuts", "Use rofi to display shortcuts")
	jsonPath := parser.String(
		"d",
		"dir",
		&argparse.Options{
			Required: false,
			Help:     "Path to the json file",
			Default:  "./my-shortcuts.json",
		},
	)
	rofiCmd = parser.String(
		"r",
		"rofi",
		&argparse.Options{
			Required: false,
			Help:     "Command line to launch rofi",
			Default:  "rofi -dmenu -p \"Shortcuts Help\" -i -no-custom -matching fuzzy",
		},
	)

	var err error = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	// get arguments from command line
	var File = loadJSON(*jsonPath)

	runRofi(File)
}

func loadJSON(filename string) File {
	filename, _ = filepath.Abs(filename)
	fmt.Println(filename)
	// open the json file
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened the file")
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var File File
	json.Unmarshal(byteValue, &File)
	defer jsonFile.Close()

	return File
}

func runRofi(file File) {
	args, err := shellwords.Parse(*rofiCmd)
	if err != nil {
		log.Fatal(err)
	}

	// Get the first selection
	cmd1 := exec.Command(args[0], args[1:]...)
	stdin1, err := cmd1.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin1.Close()
		for _, s := range file {
			io.WriteString(stdin1, s.Name+"\n")
		}
	}()

	out1, err := cmd1.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// remove trailing newline
	selectedProgram := strings.TrimSpace(string(out1))

	// Get the second selection
	cmd2 := exec.Command(args[0], args[1:]...)
	stdin2, err := cmd2.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin2.Close()
		for _, s := range file[selectedProgram].Commands {
			io.WriteString(stdin2, "("+s.Key+") "+s.Command+"\n")
		}
	}()

	out2, err := cmd2.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	var selectedCommand = strings.TrimSpace(string(out2))
	// remove trailing newline
	fmt.Println(selectedCommand)
}
