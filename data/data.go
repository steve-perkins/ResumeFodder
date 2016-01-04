package data

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

// ResumeData is the outermost container for resume data.
type ResumeData struct {
	XMLName      xml.Name      `xml:"resume" json:"-"`
	Basics       Basics        `xml:"basics" json:"basics"`
	Work         []Work        `xml:"work" json:"work"`
	Education    []Education   `xml:"education" json:"education"`
	Publications []Publication `xml:"publications" json:"publications"`
	Skills       []Skill       `xml:"skills" json:"skills"`
}

type Basics struct {
	Name     string          `xml:"name" json:"name"`
	Label    string          `xml:"label" json:"label"`
	Picture  string          `xml:"picture" json:"picture"`
	Email    string          `xml:"email" json:"email"`
	Phone    string          `xml:"phone" json:"phone"`
	Degree   string          `xml:"degree" json:"degree"`
	Website  string          `xml:"website" json:"website"`
	Summary  string          `xml:"summary" json:"summary"`
	Location Location        `xml:"location" json:"location"`
	Profiles []SocialProfile `xml:"profiles" json:"profiles"`
}

type Location struct {
	Address     string `xml:"address" json:"address"`
	PostalCode  string `xml:"postalCode" json:"postalCode"`
	City        string `xml:"city" json:"city"`
	CountryCode string `xml:"countryCode" json:"countryCode"`
	Region      string `xml:"region" json:"region"`
}

type SocialProfile struct {
	Network  string `xml:"network" json:"network"`
	Username string `xml:"username" json:"username"`
	Url      string `xml:"url" json:"url"`
}

type Work struct {
	Company    string   `xml:"company" json:"company"`
	Position   string   `xml:"position" json:"position"`
	Website    string   `xml:"website" json:"website"`
	StartDate  string   `xml:"startDate" json:"startDate"`
	EndDate    string   `xml:"endDate" json:"endDate"`
	Summary    string   `xml:"summary" json:"summary"`
	Highlights []string `xml:"highlights" json:"highlights"`
}

type Education struct {
	Institution string   `xml:"institution" json:"institution"`
	Area        string   `xml:"area" json:"area"`
	StudyType   string   `xml:"studyType" json:"studyType"`
	StartDate   string   `xml:"startDate" json:"startDate"`
	EndDate     string   `xml:"endDate" json:"endDate"`
	GPA         string   `xml:"gpa" json:"gpa"`
	Courses     []string `xml:"courses" json:"courses"`
}

type Publication struct {
	Name        string `xml:"name" json:"name"`
	Publisher   string `xml:"publisher" json:"publisher"`
	ReleaseDate string `xml:"releaseDate" json:"releaseDate"`
	Website     string `xml:"website" json:"website"`
	Summary     string `xml:"summary" json:"summary"`
}

type Skill struct {
	Name     string   `xml:"name" json:"name"`
	Level    string   `xml:"level" json:"level"`
	Keywords []string `xml:"keywords" json:"keywords"`
}

// NewResumeData initializes a ResumeData struct, with ALL nested structs initialized
// to empty state (rather than just omitted).  Useful for generating a blank XML or JSON
// file with all fields forced to be present.
//
// Of course, if you simply need to initialize a blank struct without superfluous
// nested fields, then you can always instead simply declare:
//
// data := data.ResumeData{}
//
func NewResumeData() ResumeData {
	return ResumeData{
		Basics: Basics{
			Location: Location{},
			Profiles: []SocialProfile{{}},
		},
		Work: []Work{
			{
				Highlights: []string{""},
			},
		},
		Education: []Education{
			{
				Courses: []string{""},
			},
		},
		Publications: []Publication{{}},
		Skills: []Skill{
			{
				Keywords: []string{""},
			},
		},
	}
}

// FromXmlString loads a ResumeData struct from a string of XML text.
func FromXmlString(xmlString string) (ResumeData, error) {
	return fromXml([]byte(xmlString))
}

// FromXmlFile loads a ResumeData struct from an XML file.
func FromXmlFile(xmlFilename string) (ResumeData, error) {
	bytes, err := ioutil.ReadFile(xmlFilename)
	if err != nil {
		return ResumeData{}, err
	}
	return fromXml(bytes)
}

// fromXml is a private function that provides the core logic for `FromXmlString` and `FromXmlFile`.
func fromXml(xmlBytes []byte) (ResumeData, error) {
	var data ResumeData
	err := xml.Unmarshal(xmlBytes, &data)
	if err == nil {
		// The marshal process in `toXml()` will use field tags to populate the `ResumeData.XMLName` field
		// with `resume`.  When unmarshalling from XML, we likewise strip this field value back off... to
		// better facilitate equality comparison between `ResumeData` structs (e.g. in unit testing).
		data.XMLName.Local = ""
	}
	return data, err
}

// ToXmlString writes a ResumeData struct to a string of XML text.
func ToXmlString(data ResumeData) (string, error) {
	xmlBytes, err := toXml(data)
	if err != nil {
		return "", err
	}
	return string(xmlBytes[:]), nil
}

// ToXmlFile writes a ResumeData struct to an XML file.
func ToXmlFile(data ResumeData, xmlFilename string) error {
	xmlBytes, err := toXml(data)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(xmlFilename, xmlBytes, 0644); err != nil {
		return err
	}
	return nil
}

// toXml is a private function that provides the core logic for `ToXmlString` and `ToXmlFile`.
func toXml(data ResumeData) ([]byte, error) {
	return xml.MarshalIndent(data, "", "  ")
}

// FromJsonString loads a ResumeData struct from a string of JSON text.
func FromJsonString(jsonString string) (ResumeData, error) {
	return fromJson([]byte(jsonString))
}

// FromJsonFile loads a ResumeData struct from a JSON file.
func FromJsonFile(jsonFilename string) (ResumeData, error) {
	bytes, err := ioutil.ReadFile(jsonFilename)
	if err != nil {
		return ResumeData{}, err
	}
	return fromJson(bytes)
}

// fromJson is a private function that provides the core logic for `FromJsonString` and `FromJsonFile`.
func fromJson(jsonBytes []byte) (ResumeData, error) {
	var data ResumeData
	err := json.Unmarshal(jsonBytes, &data)
	return data, err
}

// ToJsonString writes a ResumeData struct to a string of JSON text.
func ToJsonString(data ResumeData) (string, error) {
	jsonBytes, err := toJson(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes[:]), nil
}

// ToJsonFile writes a ResumeData struct to a JSON file.
func ToJsonFile(data ResumeData, jsonFilename string) error {
	jsonBytes, err := toJson(data)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(jsonFilename, jsonBytes, 0644); err != nil {
		return err
	}
	return nil
}

// toJson is a private function that provides the core logic for `ToJsonString` and `ToJsonFile`.
func toJson(data ResumeData) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

