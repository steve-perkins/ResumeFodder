package main

import (
	"errors"
	"fmt"
	"github.com/steve-perkins/resume/data"
	"os"
	"path"
	"strings"
)

func main() {
	// Identify the requested command, and perform any pre-processing of command inputs
	command, args, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Printf("\n%s\n", err)
		usage()
	}

	// Invoke the appropriate command
	switch command {
	case "init":
		if err := InitResume(args[0]); err != nil {
			fmt.Println(err)
		}
	case "convert":
		if err := ConvertResume(args[0], args[1]); err != nil {
			fmt.Println(err)
		}
	default:
		usage()
	}

	fmt.Println("Done!")
}

// ParseArgs is used to validate the arguments passed to the application, identify which command is meant to
// be executed, and perform any processing or enrichment of that command's inputs.  By placing this logic
// in its own separate function, unit tests can more easily isolate the parsing logic for verification.  The
// first string return value is a command identifies, the second string array is the set of arguments to be
// passed to the appropriate command function, and the third error value captures anything that goes wrong
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
			return "", nil, errors.New("You must specify input and output filenames (e.g. \"resume.exe convert resume.xml resume.json\"")
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

	default:
		err := errors.New("Unrecognized command.")
		return "", nil, err

	}
}

// InitResume writes a new, empty resume data file to the destination specified by the filename argument.  That
// filename must have an extension of ".xml" or ".json", and XML or JSON format will be used accordingly.
func InitResume(filename string) error {
	if strings.ToLower(path.Ext(filename)) == ".xml" {
		return data.ToXmlFile(data.NewResumeData(), filename)
	} else {
		return data.ToJsonFile(data.NewResumeData(), filename)
	}
}

// ConvertResume reads a resume file in XML or JSON format, and writes that data to another destination file
// in XML or JSON format.
func ConvertResume(inputFilename, outputFilename string) error {
	var resume data.ResumeData
	var err error
	if strings.ToLower(path.Ext(inputFilename)) == ".xml" {
		resume, err = data.FromXmlFile(inputFilename)
	} else {
		resume, err = data.FromJsonFile(inputFilename)
	}
	if err != nil {
		return err
	}

	if strings.ToLower(path.Ext(outputFilename)) == ".xml" {
		return data.ToXmlFile(resume, outputFilename)
	} else {
		return data.ToJsonFile(resume, outputFilename)
	}
}

// Usage displays information about this application and its supported arguments, and then terminates
// the application.
func usage() {
	fmt.Println("\nUsage...")
	os.Exit(0)
}
