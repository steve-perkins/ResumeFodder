package data

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestXmlConversion(t *testing.T) {
	originalData := GenerateTestResumeData()

	// Convert the data structure to a string of XML text
	xml, err := ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML text into a new resume data structure
	fromXmlData, err := FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonConversion(t *testing.T) {
	originalData := GenerateTestResumeData()

	json, err := ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

func TestXmlToJsonConversion(t *testing.T) {
	originalData := GenerateTestResumeData()

	xml, err := ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}
	json, err := ToJsonString(fromXmlData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after XML-to-JSON conversion doesn't match the original")
	}
}

func TestJsonToXmlConversion(t *testing.T) {
	originalData := GenerateTestResumeData()

	json, err := ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}
	xml, err := ToXmlString(fromJsonData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := FromXmlString(xml)
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
	deleteFileIfExists(t, xmlFilename)
	defer deleteFileIfExists(t, xmlFilename)

	// Write a resume data structure to an XML test file in the temp directory
	originalData := GenerateTestResumeData()
	err := ToXmlFile(originalData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML file back into a new resume data structure
	fromXmlData, err := FromXmlFile(xmlFilename)
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
	deleteFileIfExists(t, jsonFilename)
	defer deleteFileIfExists(t, jsonFilename)

	originalData := GenerateTestResumeData()
	err := ToJsonFile(originalData, jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := FromJsonFile(jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

// A helper function to generate fake `ResumeData` structs, for use by the various test functions.
func GenerateTestResumeData() ResumeData {
	data := ResumeData{
		Basics: Basics{
			Name:    "Peter Gibbons",
			Email:   "peter.gibbons@initech.com",
			Summary: "Just a straight-shooter with upper managment written all over him",
			Highlights: []string{
				"Once did nothing for an entire day.",
				"It was everything I thought it could be.",
			},
			Location: Location{
				Address:    "123 Main Street",
				City:       "Austin",
				Region:     "TX",
				PostalCode: "55555",
				Phone:      "555-555-5555",
			},
			Profiles: []SocialProfile{
				{
					Network:  "LinkedIn",
					Username: "peter.gibbons",
				},
			},
		},
		WorkGroups: []WorkGroup{
			{
				Work: []Work{
					{
						Company:   "Initech",
						Position:  "Software Developer",
						StartDate: "1998-02-01",
						Summary:   "Deals with the customers so the engineers don't have to.  A people person, damn it!",
						Highlights: []string{
							"Identifying Y2K-related issues in application code.",
							"As many as four people working right underneath me.",
						},
					},
				},
			},
		},
		Education: []Education{
			{
				Institution: "University of Austin",
				Area:        "B.S. Computer Science",
				StartDate:   "1993-09-01",
				EndDate:     "1997-12-01",
			},
		},
		Skills: []Skill{
			{
				Name:     "Programming",
				Level:    "Mid-level",
				Keywords: []string{"C++", "Java"},
			},
			{
				Name:     "Communication",
				Level:    "Junior",
				Keywords: []string{"Verbal", "Written"},
			},
		},
		PublicationGroups: []PublicationGroup{
			{
				Publications: []Publication{
					{
						Name: "Money Laundering for Dummies",
						Publisher: "John Wiley & Sons",
						ReleaseDate: "1999-06-01",
						ISBN: "xxxxxxxxxxx",
						Summary: "Similar to the plot from \"Superman III\"",
					},
				},
			},
		},
	}
	return data
}

func deleteFileIfExists(t *testing.T, filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal(err)
		}
	}
}
