package main

import (
	"gitlab.com/steve-perkins/ResumeFodder/data"
	"os"
	"testing"
)

// A helper function to generate fake `ResumeData` structs, for use by the various test functions.
func GenerateTestResumeData() data.ResumeData {
	data := data.ResumeData{
		Basics: data.Basics{
			Name:    "Peter Gibbons",
			Email:   "peter.gibbons@initech.com",
			Phone:   "555-555-5555",
			Summary: "Just a straight-shooter with upper managment written all over him",
			Highlights: []string{
				"Once did nothing for an entire day.",
				"It was everything I thought it could be.",
			},
			Location: data.Location{
				Address:    "123 Main Street",
				City:       "Austin",
				Region:     "TX",
				PostalCode: "55555",
			},
			Profiles: []data.SocialProfile{
				{
					Network:  "LinkedIn",
					Username: "peter.gibbons",
					Url:      "http://linkedin.com/peter.gibbons",
				},
			},
		},
		Work: []data.Work{
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
		WorkLabel: "Professional Experience",
		AdditionalWork: []data.Work{
			{
				Company:   "Flingers",
				Position:  "Burger Flipper",
				StartDate: "1993-08-01",
				EndDate:   "1998-01-31",
				Summary:   "Paying my way through school with an exciting opportunity in the fast-food service industry.",
				Highlights: []string{
					"Wore 37 pieces of flair.",
					"A terrific smile.",
				},
			},
		},
		AdditionalWorkLabel: "Academic Work Experience",
		Education: []data.Education{
			{
				Institution: "University of Austin",
				Area:        "B.S. Computer Science",
				StartDate:   "1993-09-01",
				EndDate:     "1997-12-01",
			},
		},
		Skills: []data.Skill{
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
		Publications: []data.Publication{
			{
				Name:        "Money Laundering for Dummies",
				Publisher:   "John Wiley & Sons",
				ReleaseDate: "1999-06-01",
				ISBN:        "1234567890X",
				Summary:     "Similar to the plot from \"Superman III\"",
			},
		},
		PublicationsLabel: "Publications",
		AdditionalPublications: []data.Publication{
			{
				Name:        "Washington High School Class of 1993 Yearbook",
				ReleaseDate: "1993-06-01",
				Summary:     "Served as understudy to the assistant editor for my high school yearbook.",
			},
		},
		AdditionalPublicationsLabel: "Academic Publications",
	}
	return data
}

func DeleteFileIfExists(t *testing.T, filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal(err)
		}
	}
}
