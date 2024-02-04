package utils

import (
	"encoding/json"
	"os"
)

type Resume struct {
	Summary   string
	Jobs      []Job
	Education []Edu
}

type Job struct {
	CompanyName string
	Title       string
	Experience  []string
	StartDate   string
	EndDate     string
	Years       string
}

type Edu struct {
	School       string
	DegreeOrCert string
	Years        string
}

type Bio struct {
	FirstName     string
	LastName      string
	PreferredName string
	Suffix        string
	Resume        Resume
}

// getBio reads a local json formatted resume
func ReadResume() *Bio {
	b, err := os.ReadFile("./resume.json")
	if err != nil {
		return nil
	}
	bio := &Bio{}
	err = json.Unmarshal(b, &bio)
	if err != nil {
		return nil
	}
	return bio
}
