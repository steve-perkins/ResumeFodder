package command_test

import (
	"github.com/steve-perkins/resume"
	"github.com/steve-perkins/resume/command"
	"github.com/steve-perkins/resume/data"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestInitResume(t *testing.T) {
	// Delete any pre-existing test file now, and then also clean up afterwards
	filename := filepath.Join(os.TempDir(), "testresume.xml")
	main.DeleteFileIfExists(t, filename)
	defer main.DeleteFileIfExists(t, filename)

	err := command.InitResume(filename)
	if err != nil {
		t.Fatal(err)
	}
	inMemory := data.NewResumeData()
	fromFile, err := data.FromXmlFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(inMemory, fromFile) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestConvertResume(t *testing.T) {
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	main.DeleteFileIfExists(t, xmlFilename)
	defer main.DeleteFileIfExists(t, xmlFilename)

	jsonFilename := filepath.Join(os.TempDir(), "testresume.json")
	main.DeleteFileIfExists(t, jsonFilename)
	defer main.DeleteFileIfExists(t, jsonFilename)

	err := command.InitResume(xmlFilename)
	if err != nil {
		t.Fatal(err)
	}
	err = command.ConvertResume(xmlFilename, jsonFilename)
	if err != nil {
		t.Fatal(err)
	}

	inMemory := data.NewResumeData()
	fromFile, err := data.FromJsonFile(jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(inMemory, fromFile) {
		t.Fatal("Resume data after XML-to-JSON conversion doesn't match the original")
	}
}

func TestExportResume(t *testing.T) {
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	main.DeleteFileIfExists(t, xmlFilename)
	defer main.DeleteFileIfExists(t, xmlFilename)

	resumeData := main.GenerateTestResumeData()
	err := data.ToXmlFile(resumeData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	outputFilename := filepath.Join(os.TempDir(), "resume.doc")
	templateFilename := filepath.Join("..", "templates", "default.xml")
	err = command.ExportResume(xmlFilename, outputFilename, templateFilename)
	if err != nil {
		t.Fatal(err)
	}
}
