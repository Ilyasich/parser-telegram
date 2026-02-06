package parser

import (
	"strings"
)

// Keywords that might indicate a vacancy
var keywords = []string{
	"vacancy", "вакансия",
	"hiring", "ищем",
	"job", "работа",
	"developer", "разработчик",
	"engineer", "инженер",
	"golang", "go developer",
	"remote", "удаленно",
	"salary", "зарплата", "зп",
}

// IsVacancy checks if the text contains enough keywords to be considered a vacancy
func IsVacancy(text string) bool {
	text = strings.ToLower(text)
	
	// A simple heuristic: if it contains "vacancy" or "hiring" + a role or tech, it's likely a job.
	// For now, let's just count matches.
	matches := 0
	for _, kw := range keywords {
		if strings.Contains(text, kw) {
			matches++
		}
	}

	// Threshold can be adjusted. 
	// Often a job post has at least 2 relevant words (e.g. "Golang vacancy").
	return matches >= 2
}
