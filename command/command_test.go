package command

import (
	"github.com/steve-perkins/resume/data"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestInitResume(t *testing.T) {
	// Delete any pre-existing test file now, and then also clean up afterwards
	filename := filepath.Join(os.TempDir(), "testresume.xml")
	deleteFileIfExists(t, filename)
	defer deleteFileIfExists(t, filename)

	err := InitResume(filename)
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
	deleteFileIfExists(t, xmlFilename)
	defer deleteFileIfExists(t, xmlFilename)

	jsonFilename := filepath.Join(os.TempDir(), "testresume.json")
	deleteFileIfExists(t, jsonFilename)
	defer deleteFileIfExists(t, jsonFilename)

	err := InitResume(xmlFilename)
	if err != nil {
		t.Fatal(err)
	}
	err = ConvertResume(xmlFilename, jsonFilename)
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

	// TODO... implement ExportResume

	err := ExportResume("resume.xml", "resume.doc", "defaultTemplate.xml")
	if err.Error() != "ExportResume function is not yet implemented." {
		t.Fatalf("err should be [ExportResume function is not yet implemented.], found [%s]\n", err)
	}
}

func deleteFileIfExists(t *testing.T, filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal(err)
		}
	}
}
