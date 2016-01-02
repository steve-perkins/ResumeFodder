package data

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
)

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
	StartDate  string   `xml:"startDate" json:"startDate"` // time.Time
	EndDate    string   `xml:"endDate" json:"endDate"`     // time.Time
	Summary    string   `xml:"summary" json:"summary"`
	Highlights []string `xml:"highlights" json:"highlights"`
}

type Education struct {
	Institution string   `xml:"institution" json:"institution"`
	Area        string   `xml:"area" json:"area"`
	StudyType   string   `xml:"studyType" json:"studyType"`
	StartDate   string   `xml:"startDate" json:"startDate"` // time.Time
	EndDate     string   `xml:"endDate" json:"endDate"`     // time.Time
	GPA         string   `xml:"gpa" json:"gpa"`
	Courses     []string `xml:"courses" json:"courses"`
}

type Publication struct {
	Name        string `xml:"name" json:"name"`
	Publisher   string `xml:"publisher" json:"publisher"`
	ReleaseDate string `xml:"releaseDate" json:"releaseDate"` // time.Time
	Website     string `xml:"website" json:"website"`
	Summary     string `xml:"summary" json:"summary"`
}

type Skill struct {
	Name     string   `xml:"name" json:"name"`
	Level    string   `xml:"level" json:"level"`
	Keywords []string `xml:"keywords" json:"keywords"`
}

// Initializes a ResumeData struct, with ALL nested structs initialized to empty
// state (rather than just omitted).  Useful for generating a blank XML or JSON
// file with all fields forced to be present.
//
// Of course, if you simply need to initialize a blank struct without superfluous
// nested fields, then you can always simply declare:
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

func FromXmlString(xmlString string) (ResumeData, error) {
	return fromXml([]byte(xmlString))
}

func FromXmlFile(xmlFilename string) (ResumeData, error) {
	bytes, err := ioutil.ReadFile(xmlFilename)
	if err != nil {
		return ResumeData{}, err
	}
	return fromXml(bytes)
}

func fromXml(xmlBytes []byte) (ResumeData, error) {
	var data ResumeData
	err := xml.Unmarshal(xmlBytes, &data)
	if err == nil {
		// The marshal process in `toXml()` will use field tags to populate the `ResumeData.XMLName` field
		// with `resume`.  When unmarshaling from XML, we likewise strip this field value back off... to
		// better facilitate equality comparison between `ResumeData` structs (e.g. in unit testing).
		data.XMLName.Local = ""
	}
	return data, err
}

func ToXmlString(data ResumeData) (string, error) {
	return toXml(data)
}

func ToXmlFile(data ResumeData, xmlFilename string) error {
	// TODO
	return errors.New("not implemented")
}

func toXml(data ResumeData) (string, error) {
	xmlBytes, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	xmlString := string(xmlBytes[:])
	return xmlString, nil
}

func FromJsonString(jsonString string) (ResumeData, error) {
	return fromJson([]byte(jsonString))
}

func FromJsonFile(jsonFilename string) (ResumeData, error) {
	bytes, err := ioutil.ReadFile(jsonFilename)
	if err != nil {
		return ResumeData{}, err
	}
	return fromJson(bytes)
}

func fromJson(jsonBytes []byte) (ResumeData, error) {
	var data ResumeData
	err := json.Unmarshal(jsonBytes, &data)
	return data, err
}

func ToJsonString(data ResumeData) (string, error) {
	return toJson(data)
}

func ToJsonFile(data ResumeData, jsonFilename string) error {
	// TODO
	return errors.New("not implemented")
}

func toJson(data ResumeData) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes[:])
	return jsonString, nil
}
