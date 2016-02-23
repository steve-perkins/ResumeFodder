package data_test

import (
	"gitlab.com/steve-perkins/ResumeFodder"
	"gitlab.com/steve-perkins/ResumeFodder/data"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestXmlConversion(t *testing.T) {
	originalData := main.GenerateTestResumeData()

	// Convert the data structure to a string of XML text
	xml, err := data.ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML text into a new resume data structure
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonConversion(t *testing.T) {
	originalData := main.GenerateTestResumeData()

	json, err := data.ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

func TestXmlToJsonConversion(t *testing.T) {
	originalData := main.GenerateTestResumeData()

	xml, err := data.ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}
	json, err := data.ToJsonString(fromXmlData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after XML-to-JSON conversion doesn't match the original")
	}
}

func TestJsonToXmlConversion(t *testing.T) {
	originalData := main.GenerateTestResumeData()

	json, err := data.ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}
	xml, err := data.ToXmlString(fromJsonData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after JSON-to-XML conversion doesn't match the original")
	}
}

func TestXmlFileConversion(t *testing.T) {
	// Delete any pre-existing XML test file now, and then also clean up afterwards
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	main.DeleteFileIfExists(t, xmlFilename)
	defer main.DeleteFileIfExists(t, xmlFilename)

	// Write a resume data structure to an XML test file in the temp directory
	originalData := main.GenerateTestResumeData()
	err := data.ToXmlFile(originalData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML file back into a new resume data structure
	fromXmlData, err := data.FromXmlFile(xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonFileConversion(t *testing.T) {
	jsonFilename := filepath.Join(os.TempDir(), "testresume.json")
	main.DeleteFileIfExists(t, jsonFilename)
	defer main.DeleteFileIfExists(t, jsonFilename)

	originalData := main.GenerateTestResumeData()
	err := data.ToJsonFile(originalData, jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonFile(jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}
