package command

import (
	"errors"
	"fmt"
	"github.com/steve-perkins/resume/data"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

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

// ExportResume applies a Word 2003 XML template to a resume data file, resulting in a Word document.
//
// See:
//   https://en.wikipedia.org/wiki/Microsoft_Office_XML_formats
//   https://www.microsoft.com/en-us/download/details.aspx?id=101
func ExportResume(inputFilename, outputFilename, templateFilename string) error {

	// Initialize the template engine
	funcMap := template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
		"toUpper": func(s string) string {
			return strings.ToUpper(s)
		},
		"MYYYY": func(s string) string {
			const inputFormat = "2006-01-02"
			dateValue, err := time.Parse(inputFormat, s)
			if err != nil {
				return s
			}
			const outputFormat = "1/2006"
			return dateValue.Format(outputFormat)
		},
	}
	// For some reason, I'm getting blank final results when loading templates via "ParseFiles()"... but it DOES work
	// when I first read the template contents into a string and load that via "Parse()".
	templateBytes, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		// Look for template files at the raw path provided.  If not found, then try looking for then beneath
		// the "templates" subdirectory
		templatePath := filepath.Join("templates", templateFilename)
		templateBytes, err = ioutil.ReadFile(templatePath)
		if err != nil {
			message := fmt.Sprintf("Could not find %s or %s", templateFilename, templatePath)
			return errors.New(message)
		}
	}
	templateString := string(templateBytes)
	resumeTemplate, err := template.New("resume").Funcs(funcMap).Parse(templateString) // .ParseFiles(templateFilename)
	if err != nil {
		return err
	}

	// Load the resume data
	var resumeData data.ResumeData
	extension := strings.ToLower(path.Ext(inputFilename))
	if extension == ".xml" {
		resumeData, err = data.FromXmlFile(inputFilename)
	} else if extension == ".json" {
		resumeData, err = data.FromJsonFile(inputFilename)
	} else {
		err = errors.New("Resume filename must end with \".xml\" or \".json\".")
	}
	if err != nil {
		return nil
	}

	// Open the output file and execute the template engine
	outfile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outfile.Close()
	resumeTemplate.Execute(outfile, resumeData)

	return nil
}
