package main

import (
	"fmt"
	"github.com/steve-perkins/resume/data"
	"log"
)

//var format string

func main() {
//	flag.StringVar(&format, "f", "default", "File format")
//	flag.Parse()
//	fmt.Printf("hello %s!\n", format)
//	flag.PrintDefaults()


	// Init a resume data file, in XML or JSON format
//	resumeData := data.NewResumeData()
//	resumeData := data.ResumeData{}

//	if xmlString, err := data.ToXmlString(resumeData); err == nil {
//		fmt.Println(xmlString)
//	} else {
//		fmt.Println(err)
//	}

//	if jsonString, err := data.ToJsonString(resumeData); err == nil {
//		fmt.Println(jsonString)
//	} else {
//		fmt.Println(err)
//	}

	resumeData, err := data.FromJsonFile("c:/Users/Steve/Documents/IdeaProjects/resume/resume.json")
	if err != nil {
		log.Fatal(err)
	}
	xmlString, err := data.ToXmlString(resumeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xmlString)
	copy, err := data.FromXmlString(xmlString)
	if err != nil {
		log.Fatal(err)
	}
	jsonString, err := data.ToJsonString(copy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n" + jsonString)

	// Convert to/from XML and JSON format

	// Generate resume output from data file


}
