package models

import (
	"fmt"
	"strings"
)

// Name is a struct to hold name data
type Name struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}

// ToString return the name in a signle string
func (n *Name) ToString() string {
	s := fmt.Sprintf("%v %v %v", n.First, n.Middle, n.Last)
	return strings.ReplaceAll(s, "  ", " ")
}

// Sanitizer remove remove unwanted characters
func (n *Name) Sanitizer() {
	n.First = strings.Trim(n.First, " ")
	n.Middle = strings.Trim(n.Middle, " ")
	n.Last = strings.Trim(n.Last, " ")
}

// Title capitalizes the 1st character
func (n *Name) Title() {
	n.First = strings.Title(n.First)
	n.Middle = strings.Title(n.Middle)
	n.Last = strings.Title(n.Last)
}

// Create a Name object with given parameters
func (n *Name) Create(firstName, middleName, lastName string) {
	n.First, n.Middle, n.Last = firstName, middleName, lastName

	n.Sanitizer()
}
