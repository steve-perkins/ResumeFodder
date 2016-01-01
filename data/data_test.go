package data

import (
	"testing"
	"reflect"
)

func TestStringConversion(t *testing.T) {
	originalData := generateResumeData()

	xml, err := ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}

	fromXmlData, err := FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func generateResumeData() (ResumeData) {
	data := ResumeData {
		Basics: Basics {
			Name: "Peter Gibbons",
			Email: "peter.gibbons@initech.com",
			Summary: "Just a straight-shooter with upper managment written all over him",
			Location: Location {
				City: "Austin",
				Region: "TX",
			},
			Profiles: []SocialProfile {
				{
					Network: "LinkedIn",
					Username: "peter.gibbons",
				},
			},
		},
		Work: []Work {
			Work {
				Company: "Initech",
				Position: "Software Developer",
				StartDate: "1998-02-01",
				Summary: "Deals with the customers so the engineers don't have to.  A people person, damn it!",
				Highlights: []string {
					"Identifying Y2K-related issues in application code.",
					"As many as four people working right underneath me.",
				},
			},
		},
		Education: []Education {
			{
				Institution: "University of Austin",
				Area: "B.S. Computer Science",
				StartDate: "1993-09-01",
				EndDate: "1997-12-01",
			},
		},
		Skills: []Skill {
			{
				Name: "Programming",
				Level: "Mid-level",
				Keywords: []string { "C++", "Java" },
			},
			{
				Name: "Communication",
				Level: "Junior",
				Keywords: []string { "Verbal", "Written" },
			},
		},
	}
	return data
}

