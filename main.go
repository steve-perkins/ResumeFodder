package main

import (
	"errors"
	"fmt"
	"gitlab.com/steve-perkins/ResumeFodder/command"
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
			fmt.Printf("Converted resume data file \"%s\" has been created.\n", args[1])
		} else {
			fmt.Println(err)
		}
	case "export":
		if err := command.ExportResume(args[0], args[1], args[2]); err == nil {
			fmt.Printf("Resume has been exported to \"%s\" using template \"%s\".\n", args[1], args[2])
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
			filename = "resume.json"
		} else {
			filename = args[2]
		}
		extension := strings.ToLower(path.Ext(filename))
		if extension != ".xml" && extension != ".json" {
			err := errors.New("Filename to initialize must have a '.json' or '.xml' extension.")
			return "", nil, err
		}
		return "init", []string{filename}, nil

	case "convert":
		if len(args) < 4 {
			return "", nil, errors.New("You must specify input and output filenames (e.g. \"ResumeFodder.exe convert resume.json resume.xml\")")
		}
		inputFilename := args[2]
		inputExtension := strings.ToLower(path.Ext(inputFilename))
		if inputExtension != ".json" && inputExtension != ".xml" {
			return "", nil, errors.New("Source file must have a '.json' or '.xml' extension.")
		}
		outputFilename := args[3]
		outputExtension := strings.ToLower(path.Ext(outputFilename))
		if outputExtension != ".json" && outputExtension != ".xml" {
			return "", nil, errors.New("Target file must have a '.json' or '.xml' extension.")
		}
		if inputExtension == ".json" && outputExtension != ".xml" {
			return "", nil, errors.New("When converting a JSON source file, the target filename must have an '.xml' extension")
		}
		if inputExtension == ".xml" && outputExtension != ".json" {
			return "", nil, errors.New("When converting an XML source file, the target filename must have a '.json' extension")
		}
		return "convert", []string{inputFilename, outputFilename}, nil

	case "export":
		if len(args) < 4 {
			return "", nil, errors.New("You must specify input and output filenames (e.g. \"ResumeFodder.exe export resume.json resume.doc\"), and optionally a template name.")
		}
		inputFilename := args[2]
		inputExtension := strings.ToLower(path.Ext(inputFilename))
		if inputExtension != ".json" && inputExtension != ".xml" {
			return "", nil, errors.New("Source file must have a '.json' or '.xml' extension.")
		}
		outputFilename := args[3]
		outputExtension := strings.ToLower(path.Ext(outputFilename))
		if outputExtension != ".doc" && outputExtension != ".xml" {
			return "", nil, errors.New("Target file must have a '.doc' or '.xml' extension.")
		}
		var templateFilename string
		if len(args) < 5 {
			templateFilename = "plain.xml"
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
	fmt.Println(`
Usage:

   ResumeFodder.exe COMMAND <args>

... where "COMMAND" is one of the following:

	init    - Create a new empty resume data file
	convert - Convert a JSON-formatted resume data file into XML
	          format, or vice-versa
	export  - Process and resume data file with a given template,
	          to publish a Microsoft Word resume file

Full details for each command:

ResumeFodder.exe init <filename>
ResumeFodder.exe init resume.xml

	Will generate an empty resume data file with the specified
	filename, which must have either a '.json' or '.xml' file
	extension.

	If no filename is specified, then a data file will be created
	with filename 'resume.json'.

ResumeFodder.exe convert <input filename> <output filename>
ResumeFodder.exe convert resume.xml resume.json

	The resume data file specified by the first parameter will
	be converted to the filename specified by the second parameter.

	If the second file already exists, then any contents will
	be overwritten.  Both filenames must have either a '.json' or
	'.xml' file extension.

ResumeFodder.exe export <data filename> <output filename> <template filename>
ResumeFodder.exe export resume.json resume.doc templates/plain.xml

	The resume data file specified by the first parameter will
	published as a Microsoft Word file with the name specified by
	the second parameter.  The template file specified by the third
	parameter will be used to generate the output.

	The data filename must have either a '.json' or '.xml' file
	extension.  The output will be a Microsoft Word 2003 XML file,
	and its name must have either a '.doc' or '.xml' file
	extension.  The template file must likewise be a Word 2003
	XML file (with Go template tags), and its name too must have
	either a '.doc' or '.xml' extension.

	If the specified template is not found in the current working
	directory, then the application will look under a "templates"
	subdirectory in the current working directory.  If no template
	is specified, the the application will use the "plain.xml"
	template.
`)
	os.Exit(0)
}
