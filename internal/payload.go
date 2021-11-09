package internal

import (
	"strings"
	"time"
)

const (
	countryCode = "+55"
)

type ReminderPayload struct {
	Name              string    `json:"name"`
	PhoneNumber       string    `json:"phone_number"`
	Message           string    `json:"message"`
	ReminderInMinutes int       `json:"reminder_in_minutes"`
	CreatedAt         time.Time `json:"-"`
}

func (r ReminderPayload) GetPhoneNumber() string {
	if strings.HasPrefix(r.PhoneNumber, countryCode) {
		return r.PhoneNumber
	}

	return countryCode + r.PhoneNumber
}

func (r ReminderPayload) GetReminderDate() time.Time {
	return r.CreatedAt.Add(time.Minute * time.Duration(r.ReminderInMinutes))
}
