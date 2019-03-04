package models

type User struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`               // Employee or Contractor
	Duration int      `json:"duration,omitempty"` // duration in years
	Role     string   `json:"role,omitempty"`
	Tags     []string `json:"tags,omitempty"` // C++, C#, etc
}
