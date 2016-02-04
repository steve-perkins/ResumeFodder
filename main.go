package main

import (
	"errors"
	"fmt"
	"github.com/steve-perkins/resume/command"
	"os"
	"path"
	"strings"
)

func main() {
	// Identify the requested command, and perform any pre-processing of command inputs
	commandName, args, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Printf("\n%s\n", err)
		usage()
	}

	// Invoke the appropriate command
	switch commandName {
	case "init":
		if err := command.InitResume(args[0]); err == nil {
			fmt.Printf("Empty resume data file \"%s\" has been created.\n", args[0])
		} else {
			fmt.Println(err)
		}
	case "convert":
		if err := command.ConvertResume(args[0], args[1]); err == nil {
			fmt.Println("Converted resume data file \"%s\" has been created.\n", args[1])
		} else {
			fmt.Println(err)
		}
	case "export":
		if err := command.ExportResume(args[0], args[1], args[2]); err == nil {
			fmt.Println("Resume has been exported to \"%s\" using template \"%s\".\n", args[1], args[2])
		} else {
			fmt.Println(err)
		}
	default:
		usage()
	}

}

// ParseArgs is used to validate the arguments passed to the application, identify which command is meant to
// be executed, and perform any processing or enrichment of that command's inputs.  By placing this logic
// in its own separate function, unit tests can more easily isolate the parsing logic for verification.
//
// The input string array should come from `os.Args`, and as such it's expected that the first array element
// will be the application executable name.
//
// The first string return value is a command identifier.  The second string array is the arguments to be
// passed to the appropriate command function.  The third error value captures anything that goes wrong
// during validation.
func ParseArgs(args []string) (string, []string, error) {
	if len(args) < 2 {
		err := errors.New("No command was specified.")
		return "", nil, err
	}
	command := strings.ToLower(args[1])
	switch command {

	case "init":
		var filename string
		if len(args) < 3 {
			filename = "resume.xml"
		} else {
			filename = args[2]
		}
		extension := strings.ToLower(path.Ext(filename))
		if extension != ".xml" && extension != ".json" {
			err := errors.New("Filename to initialize must have an '.xml' or '.json' extension.")
			return "", nil, err
		}
		return "init", []string{filename}, nil

	case "convert":
		if len(args) < 4 {
			return "", nil, errors.New("You must specify input and output filenames (e.g. \"resume.exe convert resume.xml resume.json\")")
		}
		inputFilename := args[2]
		inputExtension := strings.ToLower(path.Ext(inputFilename))
		if inputExtension != ".xml" && inputExtension != ".json" {
			return "", nil, errors.New("Source file must have an '.xml' or '.json' extension.")
		}
		outputFilename := args[3]
		outputExtension := strings.ToLower(path.Ext(outputFilename))
		if outputExtension != ".xml" && outputExtension != ".json" {
			return "", nil, errors.New("Target file must have an '.xml' or '.json' extension.")
		}
		if inputExtension == ".xml" && outputExtension != ".json" {
			return "", nil, errors.New("When converting an XML source file, the target filename must have a '.json' extension")
		}
		if inputExtension == ".json" && outputExtension != ".xml" {
			return "", nil, errors.New("When converting a JSON source file, the target filename must have an '.xml' extension")
		}
		return "convert", []string{inputFilename, outputFilename}, nil

	case "export":
		if len(args) < 4 {
			return "", nil, errors.New("You must specify input and output filenames (e.g. \"resume.exe export resume.xml resume.doc\"), and optionally a template name.")
		}
		inputFilename := args[2]
		inputExtension := strings.ToLower(path.Ext(inputFilename))
		if inputExtension != ".xml" && inputExtension != ".json" {
			return "", nil, errors.New("Source file must have an '.xml' or '.json' extension.")
		}
		outputFilename := args[3]
		outputExtension := strings.ToLower(path.Ext(outputFilename))
		if outputExtension != ".doc" && outputExtension != ".xml" {
			return "", nil, errors.New("Target file must have a '.doc' or '.xml' extension.")
		}
		var templateFilename string
		if len(args) < 5 {
			templateFilename = "defaultTemplate.xml"
		} else {
			templateFilename = args[4]
		}
		templateExtension := strings.ToLower(path.Ext(templateFilename))
		if templateExtension != ".doc" && templateExtension != ".xml" {
			return "", nil, errors.New("Template file must have a '.doc' or '.xml' extension.")
		}
		return "export", []string{inputFilename, outputFilename, templateFilename}, nil

	default:
		err := errors.New("Unrecognized command.")
		return "", nil, err

	}
}

// usage displays information about this application and its supported arguments, and then terminates
// the application.
func usage() {

	// TODO... write full usage text

	fmt.Println("\nUsage...")
	os.Exit(0)
}
