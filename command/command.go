package command

import (
	"errors"
	"github.com/steve-perkins/resume/data"
	"path"
	"strings"
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

// ExportResume applies a Office XML template to a resume data file, resulting in a Word 2003 XML document.
//
// See:
//   https://en.wikipedia.org/wiki/Microsoft_Office_XML_formats
//   https://www.microsoft.com/en-us/download/details.aspx?id=101
func ExportResume(inputFilename, outputFilename, templateFilename string) error {

	// TODO... implement the following steps
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

	return errors.New("ExportResume function is not yet implemented.")
}
