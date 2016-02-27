package command_test

import (
	"gitlab.com/steve-perkins/ResumeFodder/command"
	"gitlab.com/steve-perkins/ResumeFodder/data"
	"gitlab.com/steve-perkins/ResumeFodder/testutils"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestInitResumeFile(t *testing.T) {
	// Delete any pre-existing test file now, and then also clean up afterwards
	filename := filepath.Join(os.TempDir(), "testresume.xml")
	testutils.DeleteFileIfExists(t, filename)
	defer testutils.DeleteFileIfExists(t, filename)

	err := command.InitResumeFile(filename)
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

func TestInitResumeJson(t *testing.T) {
	json, err := command.InitResumeJson()
	if err != nil {
		t.Fatal(err)
	}
	inMemory := data.NewResumeData()
	fromString, err := data.FromJsonString(json)
	if !reflect.DeepEqual(inMemory, fromString) {
		t.Fatal("Resume data after conversion doesn't match the original")
	}
}

func TestInitResumeXml(t *testing.T) {
	xml, err := command.InitResumeXml()
	if err != nil {
		t.Fatal(err)
	}
	inMemory := data.NewResumeData()
	fromString, err := data.FromXmlString(xml)
	if !reflect.DeepEqual(inMemory, fromString) {
		t.Fatal("Resume data after conversion doesn't match the original")
	}
}

func TestConvertResumeFile(t *testing.T) {
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	testutils.DeleteFileIfExists(t, xmlFilename)
	defer testutils.DeleteFileIfExists(t, xmlFilename)

	jsonFilename := filepath.Join(os.TempDir(), "testresume.json")
	testutils.DeleteFileIfExists(t, jsonFilename)
	defer testutils.DeleteFileIfExists(t, jsonFilename)

	err := command.InitResumeFile(xmlFilename)
	if err != nil {
		t.Fatal(err)
	}
	err = command.ConvertResumeFile(xmlFilename, jsonFilename)
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

// See also "TestExportResume_TemplateDefaultPath()", in the base "ResumeFodder" project's "main_test.go" test file.
func TestExportResumeFile_TemplateRelativePath(t *testing.T) {
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	testutils.DeleteFileIfExists(t, xmlFilename)
	defer testutils.DeleteFileIfExists(t, xmlFilename)

	resumeData := testutils.GenerateTestResumeData()
	err := data.ToXmlFile(resumeData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	outputFilename := filepath.Join(os.TempDir(), "resume.doc")
	templateFilename := filepath.Join("..", "templates", "plain.xml")
	err = command.ExportResumeFile(xmlFilename, outputFilename, templateFilename)
	if err != nil {
		t.Fatal(err)
	}
}
