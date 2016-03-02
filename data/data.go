package data

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

const SCHEMA_VERSION = 1

// ResumeData is the outermost container for resume data.
type ResumeData struct {
	// XMLName provides a name for the top-level element, when working with resume data files in XML format.  This
	// field is ignored when working with files in JSON format.
	XMLName xml.Name `xml:"resume" json:"-"`
	// Version is an identifier for the schema structure.  If breaking changes occur in the future, then ResumeFodder
	// can use this value to recognize the incompatibility and provide options.
	Version int    `xml:"version" json:"version"`
	Basics  Basics `xml:"basics" json:"basics"`
	Work    []Work `xml:"work" json:"work"`
	// AdditionalWork is an extra field, not found within the standard JSON-Resume spec.  It is intended to store
	// employment history that should be presented differently from that in the main "Work" field.
	//
	// Specifically, if you have a lengthy work history, then you might store the oldest jobs in this field... so that
	// a template can present them in abbreviated format (i.e. no highlights), with perhaps a "Further details
	// available upon request"-type note.  It could similarly be used for high-school or college jobs that are
	// only worth mentioning for entry-level candidates only.
	//
	// Obviously, the records in this extra field would be ignored if you used your data file with a standard
	// JSON-Resume processor.  Otherwise, migration would require you to move any "AdditionalWork" records to the
	// "Work" field.
	AdditionalWork []Work `xml:"additionalWork" json:"additionalWork"`
	// WorkLabel is an extra field, not found within the standard JSON-Resume spec.  It is intended to tell templates
	// how to present the "Work" and "AdditionalWork" sections, when both are used (e.g. "Recent Experience"
	// versus "Prior Experience").
	WorkLabel string `xml:"workLabel" json:"workLabel"`
	// AdditionalWorkLabel is an extra field, not found within the standard JSON-Resume spec.  It is intended to tell
	// templates how to present the "Work" and "AdditionalWork" sections, when both are used (e.g. "Recent Experience"
	// versus "Prior Experience").
	AdditionalWorkLabel string        `xml:"additionalWorkLabel" json:"additionalWorkLabel"`
	Education           []Education   `xml:"education" json:"education"`
	Publications        []Publication `xml:"publications" json:"publications"`
	// AdditionalPublications is an extra field, not found within the standard JSON-Resume spec.  It is intended to
	// store publications that should be presented differently from those in the main "Publications" field.
	//
	// Specifically, if you collaborated on publications in which you were not an author or co-author (e.g. a technical
	// reviewer instead), then you might store those publications here so that a template can present them
	// without implying that you were an author.
	//
	// Obviously, the records in this extra field would be ignored if you used your data file with a standard
	// JSON-Resume processor.  Otherwise, migration would require you to move any "AdditionalPublications" records to
	// the "Publications" field.
	AdditionalPublications []Publication `xml:"additionalPublications" json:"additionalPublications"`
	// PublicationsLabel is an extra field, not found within the standard JSON-Resume spec.  It is intended to tell
	// templates how to present the "Publications" and "AdditionalPublications" sections, when both are used
	// (e.g. "Publications (Author)" versus "Publications (Technical Reviewer)").
	PublicationsLabel string `xml:"publicationsLabel" json:"publicationsLabel"`
	// AdditionalPublicationsLabel is an extra field, not found within the standard JSON-Resume spec.  It is intended
	// to tell templates how to present the "Publications" and "AdditionalPublications" sections, when both are used
	// (e.g. "Publications (Author)" versus "Publications (Technical Reviewer)").
	AdditionalPublicationsLabel string  `xml:"additionalPublicationsLabel" json:"additionalPublicationsLabel"`
	Skills                      []Skill `xml:"skills" json:"skills"`
}

// Basics is a container for top-level resume data.  These fields could just as well hang off the parent "ResumeData"
// struct, but this structure mirrors how the JSON-Resume spec arranges them.
type Basics struct {
	Name    string `xml:"name" json:"name"`
	Label   string `xml:"label" json:"label"`
	Picture string `xml:"picture" json:"picture"`
	Email   string `xml:"email" json:"email"`
	Phone   string `xml:"phone" json:"phone"`
	Degree  string `xml:"degree" json:"degree"`
	Website string `xml:"website" json:"website"`
	Summary string `xml:"summary" json:"summary"`
	// Highlights is an extra field, not found within the standard JSON-Resume spec.  It is intended for additional
	// top-level information, that a template might present with a bullet-point list or other similar formatting
	// next to the top-level "Summary" field.
	//
	// Obviously, the records in this extra field would be ignored if you used your data file with a standard
	// JSON-Resume processor.  Once the other JSON-Resume processors gain mature support for HTML and/or Markdown
	// line-break formatting within field values, then perhaps you could migrate "Highlights" data to within the
	// "Summary" field.
	Highlights []string        `xml:"highlights" json:"highlights"`
	Location   Location        `xml:"location" json:"location"`
	Profiles   []SocialProfile `xml:"profiles" json:"profiles"`
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
	// TODO: Perhaps job listings should have 'City' and 'Region' extension fields, as this is commonly found on resumes
	Company    string   `xml:"company" json:"company"`
	Position   string   `xml:"position" json:"position"`
	Website    string   `xml:"website" json:"website"`
	StartDate  string   `xml:"startDate" json:"startDate"`
	EndDate    string   `xml:"endDate" json:"endDate"`
	Summary    string   `xml:"summary" json:"summary"`
	Highlights []string `xml:"highlights" json:"highlights"`
}

type Education struct {
	// TODO: Perhaps education listings should have 'City' and 'Region' extension fields, as this is commonly found on resumes
	Institution string   `xml:"institution" json:"institution"`
	Area        string   `xml:"area" json:"area"`
	StudyType   string   `xml:"studyType" json:"studyType"`
	StartDate   string   `xml:"startDate" json:"startDate"`
	EndDate     string   `xml:"endDate" json:"endDate"`
	GPA         string   `xml:"gpa" json:"gpa"`
	Courses     []string `xml:"courses" json:"courses"`
}

type PublicationGroup struct {
	Name         string        `xml:"name" json:"name"`
	Publications []Publication `xml:"publications" json:"publications"`
}

type Publication struct {
	Name        string `xml:"name" json:"name"`
	Publisher   string `xml:"publisher" json:"publisher"`
	ReleaseDate string `xml:"releaseDate" json:"releaseDate"`
	Website     string `xml:"website" json:"website"`
	Summary     string `xml:"summary" json:"summary"`
	// ISBN is an extra field, not found within the standard JSON-Resume spec.  Obviously, this value will be
	// ignored if you used your data file with another JSON-Resume processor.  You could perhaps migrate by
	// cramming this info into the "Summary" field.
	ISBN string `xml:"isbn" json:"isbn"`
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
		Version: SCHEMA_VERSION,
		Basics: Basics{
			Location: Location{},
			Profiles: []SocialProfile{{}},
		},
		Work: []Work{
			{
				Highlights: []string{""},
			},
		},
		AdditionalWork: []Work{
			{
				Highlights: []string{""},
			},
		},
		Education: []Education{
			{
				Courses: []string{""},
			},
		},
		Publications:           []Publication{{}},
		AdditionalPublications: []Publication{{}},
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
