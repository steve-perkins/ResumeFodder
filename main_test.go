package main

import (
	"github.com/steve-perkins/resume/data"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseArgs_NoArgs(t *testing.T) {
	_, _, err := ParseArgs([]string{"resume.exe"})
	if err == nil || err.Error() != "No command was specified." {
		t.Fatalf("err should be [No command was specified.], found [%s]\n", err)
	}
}

func TestParseArgs_Init_NoArg(t *testing.T) {
	command, args, err := ParseArgs([]string{"resume.exe", "init"})
	if command != "init" {
		t.Fatalf("command should be [init], found [%s]\n", command)
	}
	if len(args) != 1 || args[0] != "resume.xml" {
		t.Fatalf("args should be [resume.xml], found %s\n", args)
	}
	if err != nil {
		t.Fatalf("err should be nil, found [%s]\n", err)
	}
}

func TestParseArgs_Init_InvalidFilename(t *testing.T) {
	_, _, err := ParseArgs([]string{"resume.exe", "init", "bad_extension.foo"})
	if err == nil || err.Error() != "Filename to initialize must have an '.xml' or '.json' extension." {
		t.Fatalf("err should be [Filename to initialize must have an '.xml' or '.json' extension.], found [%s]\n", err)
	}
}

func TestParseArgs_Init_Valid(t *testing.T) {
	command, args, err := ParseArgs([]string{"resume.exe", "init", "resume.xml"})
	if command != "init" {
		t.Fatalf("command should be [init], found [%s]\n", command)
	}
	if len(args) != 1 || args[0] != "resume.xml" {
		t.Fatalf("args should be [resume.xml], found %s\n", args)
	}
	if err != nil {
		t.Fatalf("err should be nil, found [%s]\n", err)
	}
}

func TestParseArgs_Convert_NoArgs(t *testing.T) {
	_, _, err := ParseArgs([]string{"resume.exe", "convert"})
	if err == nil || err.Error() != "You must specify input and output filenames (e.g. \"resume.exe convert resume.xml resume.json\")" {
		t.Fatalf("err should be [You must specify input and output filenames (e.g. \"resume.exe convert resume.xml resume.json\")], found [%s]\n", err)
	}
}

func TestParseArgs_Convert_InvalidFilename(t *testing.T) {
	// Source and target must be XML or JSON
	_, _, err := ParseArgs([]string{"resume.exe", "convert", "bad_extension.foo", "resume.json"})
	if err == nil || err.Error() != "Source file must have an '.xml' or '.json' extension." {
		t.Fatalf("err should be [Source file must have an '.xml' or '.json' extension.], found [%s]\n", err)
	}
	_, _, err = ParseArgs([]string{"resume.exe", "convert", "resume.xml", "bad_extension.foo"})
	if err == nil || err.Error() != "Target file must have an '.xml' or '.json' extension." {
		t.Fatalf("err should be [Target file must have an '.xml' or '.json' extension.], found [%s]\n", err)
	}

	// Conversion from one format must be to the other
	_, _, err = ParseArgs([]string{"resume.exe", "convert", "resume.xml", "copy.xml"})
	if err == nil || err.Error() != "When converting an XML source file, the target filename must have a '.json' extension" {
		t.Fatalf("err should be [When converting an XML source file, the target filename must have a '.json' extension], found [%s]\n", err)
	}
	_, _, err = ParseArgs([]string{"resume.exe", "convert", "resume.json", "copy.json"})
	if err == nil || err.Error() != "When converting a JSON source file, the target filename must have an '.xml' extension" {
		t.Fatalf("err should be [When converting a JSON source file, the target filename must have an '.xml' extension], found [%s]\n", err)
	}
}

func TestParseArgs_Convert_Valid(t *testing.T) {
	command, args, err := ParseArgs([]string{"resume.exe", "convert", "resume.xml", "resume.json"})
	if command != "convert" {
		t.Fatalf("command should be [convert], found [%s]\n", command)
	}
	if len(args) != 2 || args[0] != "resume.xml" || args[1] != "resume.json" {
		t.Fatalf("args should be [resume.xml resume.json], found %s\n", args)
	}
	if err != nil {
		t.Fatalf("err should be nil, found [%s]\n", err)
	}
}

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

func deleteFileIfExists(t *testing.T, filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal(err)
		}
	}
}
