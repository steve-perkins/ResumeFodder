package main_test

import (
	"errors"
	"gitlab.com/steve-perkins/ResumeFodder"
	"gitlab.com/steve-perkins/ResumeFodder/command"
	"gitlab.com/steve-perkins/ResumeFodder/data"
	"os"
	"path/filepath"
	"testing"
)

func TestNoArgs(t *testing.T) {
	_, _, err := main.ParseArgs([]string{"resume.exe"})
	if err == nil || err.Error() != "No command was specified." {
		t.Fatalf("err should be [No command was specified.], found [%s]\n", err)
	}
}

func TestInit_NoArg(t *testing.T) {
	command, args, err := main.ParseArgs([]string{"resume.exe", "init"})
	if command != "init" {
		t.Fatalf("command should be [init], found [%s]\n", command)
	}
	if len(args) != 1 || args[0] != "resume.json" {
		t.Fatalf("args should be [resume.json], found %s\n", args)
	}
	if err != nil {
		t.Fatalf("err should be nil, found [%s]\n", err)
	}
}

func TestInit_InvalidFilename(t *testing.T) {
	_, _, err := main.ParseArgs([]string{"resume.exe", "init", "bad_extension.foo"})
	if err == nil || err.Error() != "Filename to initialize must have a '.json' or '.xml' extension." {
		t.Fatalf("err should be [Filename to initialize must have a '.json' or '.xml' extension.], found [%s]\n", err)
	}
}

func TestInit_Valid(t *testing.T) {
	command, args, err := main.ParseArgs([]string{"resume.exe", "init", "resume.xml"})
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

func TestConvert_NoArgs(t *testing.T) {
	_, _, err := main.ParseArgs([]string{"resume.exe", "convert"})
	if err == nil || err.Error() != "You must specify input and output filenames (e.g. \"resume convert resume.json resume.xml\")" {
		t.Fatalf("err should be [You must specify input and output filenames (e.g. \"resume convert resume.json resume.xml\")], found [%s]\n", err)
	}
}

func TestConvert_InvalidFilename(t *testing.T) {
	// Source and target must be XML or JSON
	_, _, err := main.ParseArgs([]string{"resume.exe", "convert", "bad_extension.foo", "resume.json"})
	if err == nil || err.Error() != "Source file must have a '.json' or '.xml' extension." {
		t.Fatalf("err should be [Source file must have a '.json' or '.xml' extension.], found [%s]\n", err)
	}
	_, _, err = main.ParseArgs([]string{"resume.exe", "convert", "resume.xml", "bad_extension.foo"})
	if err == nil || err.Error() != "Target file must have a '.json' or '.xml' extension." {
		t.Fatalf("err should be [Target file must have a '.json' or '.xml' extension.], found [%s]\n", err)
	}

	// Conversion from one format must be to the other
	_, _, err = main.ParseArgs([]string{"resume.exe", "convert", "resume.xml", "copy.xml"})
	if err == nil || err.Error() != "When converting an XML source file, the target filename must have a '.json' extension" {
		t.Fatalf("err should be [When converting an XML source file, the target filename must have a '.json' extension], found [%s]\n", err)
	}
	_, _, err = main.ParseArgs([]string{"resume.exe", "convert", "resume.json", "copy.json"})
	if err == nil || err.Error() != "When converting a JSON source file, the target filename must have an '.xml' extension" {
		t.Fatalf("err should be [When converting a JSON source file, the target filename must have an '.xml' extension], found [%s]\n", err)
	}
}

func TestConvert_Valid(t *testing.T) {
	command, args, err := main.ParseArgs([]string{"resume.exe", "convert", "resume.xml", "resume.json"})
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

func TestExport_NoArg(t *testing.T) {
	_, _, err := main.ParseArgs([]string{"resume.exe", "export"})
	if err == nil || err.Error() != "You must specify input and output filenames (e.g. \"resume export resume.json resume.doc\"), and optionally a template name." {
		t.Fatalf("err should be [You must specify input and output filenames (e.g. \"resume export resume.json resume.doc\"), and optionally a template name.], found [%s]\n", err)
	}
}

func TestExport_InvalidSourceFilename(t *testing.T) {
	// Source must be XML or JSON
	_, _, err := main.ParseArgs([]string{"resume.exe", "export", "bad_extension.foo", "resume.doc"})
	if err == nil || err.Error() != "Source file must have a '.json' or '.xml' extension." {
		t.Fatalf("err should be [Source file must have a '.json' or '.xml' extension.], found [%s]\n", err)
	}
}

func TestExport_InvalidTargetFilename(t *testing.T) {
	// Target must be DOC or XML
	_, _, err := main.ParseArgs([]string{"resume.exe", "export", "resume.xml", "bad_extension.foo"})
	if err == nil || err.Error() != "Target file must have a '.doc' or '.xml' extension." {
		t.Fatalf("err should be [Target file must have a '.doc' or '.xml' extension.], found [%s]\n", err)
	}
}

func TestExport_InvalidTemplateFilename(t *testing.T) {
	// Template must be DOC or XML
	_, _, err := main.ParseArgs([]string{"resume.exe", "export", "resume.xml", "resume.doc", "templates/bad_extension.foo"})
	if err == nil || err.Error() != "Template file must have a '.doc' or '.xml' extension." {
		t.Fatalf("err should be [Template file must have a '.doc' or '.xml' extension.], found [%s]\n", err)
	}
}

func TestExport_NoTemplateFilename(t *testing.T) {
	_, args, err := main.ParseArgs([]string{"resume.exe", "export", "resume.xml", "resume.doc"})
	if err != nil {
		t.Fatal(err)
	}
	if args[2] != "plain.xml" {
		t.Fatal(errors.New("When no template file is specified, the default value of \"plain.xml\" should be used"))
	}
}

func TestExport_Valid(t *testing.T) {
	command, args, err := main.ParseArgs([]string{"resume.exe", "export", "resume.xml", "resume.doc", "templates/default.xml"})
	if command != "export" {
		t.Fatalf("command should be [export], found [%s]\n", command)
	}
	if len(args) != 3 || args[0] != "resume.xml" || args[1] != "resume.doc" || args[2] != "templates/default.xml" {
		t.Fatalf("args should be [resume.xml resume.json templates/default.xml], found %s\n", args)
	}
	if err != nil {
		t.Fatalf("err should be nil, found [%s]\n", err)
	}
}

// Tests that when the export command can't find a template at the specified location, that the command will try
// prepending that with the "templates" directory.  This test logically belongs in the "command/command_test.go" test
// file, but instead lives here because it requires the current working directory to be the project root.
func TestExportResume_TemplateDefaultPath(t *testing.T) {
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	main.DeleteFileIfExists(t, xmlFilename)
	defer main.DeleteFileIfExists(t, xmlFilename)

	resumeData := main.GenerateTestResumeData()
	err := data.ToXmlFile(resumeData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	outputFilename := filepath.Join(os.TempDir(), "resume.doc")
	templateFilename := "plain.xml"
	err = command.ExportResume(xmlFilename, outputFilename, templateFilename)
	if err != nil {
		t.Fatal(err)
	}
}
