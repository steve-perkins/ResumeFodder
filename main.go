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
		if err := InitResume(args[0]); err == nil {
			fmt.Printf("Empty resume data file \"%s\" has been created.\n", args[0])
		} else {
			fmt.Println(err)
		}
	case "convert":
		if err := ConvertResume(args[0], args[1]); err == nil {
			fmt.Println("Converted resume data file \"%s\" has been created.\n", args[1])
		} else {
			fmt.Println(err)
		}
	case "export":
		if err := ExportResume(args[0], args[1], args[2]); err == nil {
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
		if templateExtension != ".doc" && inputExtension != ".xml" {
			return "", nil, errors.New("Template file must have a '.doc' or '.xml' extension.")
		}
		return "export", []string{inputFilename, outputFilename, templateFilename}, nil

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

// ConvertResume reads a resume data file in XML or JSON format, and writes that data to another destination file
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

// ExportResume applies a Office XML template to a resume data file, resulting in a Word 2003 XML document.
//
// See:
//   https://en.wikipedia.org/wiki/Microsoft_Office_XML_formats
//   https://www.microsoft.com/en-us/download/details.aspx?id=101
func ExportResume(inputFilename, outputFilename, templateFilename string) error {

	// TODO
	//
	// [1] Load the resume data structure, and iterate through each field
	// [2] Divide the field by line breaks
	// [3] If there is more than one line in a field, then add close-paragraph markup the end of the first
	//     line, and surround the subsequent lines with open-and-close-paragraph markup
	// [4] If a line begins with Markdown bullet-list markup, then make it's paragraph markup of the appropriate style
	// [5] If Markdown bold or italics markup is found within a line, then close the current "r" and "t"
	//     tags.  Start new "r" and "t" tags, with the appropriate style and text, close them, and then re-start
	//     a new "r" and "t" tag set with the default style.  *****NOTE*****: template authors must always insert
	//     text insertion tokens within "t" tags.
	// [6] Overwrite the string values within the resume data structure with any modifications
	// [7] Perform Go template token replacement.
	//
	// [???] Move this and all of the other command functions to a new "command[s?]" package.

	return errors.New("ExportResume function is not yet implemented.")
}

// usage displays information about this application and its supported arguments, and then terminates
// the application.
func usage() {

	// TODO... write full usage text

	fmt.Println("\nUsage...")
	os.Exit(0)
}

