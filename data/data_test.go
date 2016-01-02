package data

import (
	"reflect"
	"testing"
)

func TestXmlConversion(t *testing.T) {
	originalData := generateResumeData()

	// Convert the data structure to a string of XML text
	xml, err := ToXmlString(originalData)
	fatalIfError(t, err)

	// Parse that XML text into a new resume data structure
	fromXmlData, err := FromXmlString(xml)
	fatalIfError(t, err)

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonConversion(t *testing.T) {
	originalData := generateResumeData()

	json, err := ToJsonString(originalData)
	fatalIfError(t, err)
	fromJsonData, err := FromJsonString(json)
	fatalIfError(t, err)

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

func TestXmlToJsonConversion(t *testing.T) {
	originalData := generateResumeData()

	xml, err := ToXmlString(originalData)
	fatalIfError(t, err)
	fromXmlData, err := FromXmlString(xml)
	fatalIfError(t, err)
	json, err := ToJsonString(fromXmlData)
	fatalIfError(t, err)
	fromJsonData, err := FromJsonString(json)
	fatalIfError(t, err)

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after XML-to-JSON conversion doesn't match the original")
	}
}

func TestJsonToXmlConversion(t *testing.T) {
	originalData := generateResumeData()

	json, err := ToJsonString(originalData)
	fatalIfError(t, err)
	fromJsonData, err := FromJsonString(json)
	fatalIfError(t, err)
	xml, err := ToXmlString(fromJsonData)
	fatalIfError(t, err)
	fromXmlData, err := FromXmlString(xml)
	fatalIfError(t, err)

	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after JSON-to-XML conversion doesn't match the original")
	}
}

// A helper function to generate fake `ResumeData` structs, for use by the various test functions.
func generateResumeData() ResumeData {
	data := ResumeData{
		Basics: Basics{
			Name:    "Peter Gibbons",
			Email:   "peter.gibbons@initech.com",
			Summary: "Just a straight-shooter with upper managment written all over him",
			Location: Location{
				City:   "Austin",
				Region: "TX",
			},
			Profiles: []SocialProfile{
				{
					Network:  "LinkedIn",
					Username: "peter.gibbons",
				},
			},
		},
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
	}
	return data
}

// Fails the test if a real error is passed, or else takes no action if a nil error is passed.
func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
