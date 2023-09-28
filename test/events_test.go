package test

import (
	"testing"

	"github.com/verzth/go-coingecko"
)

func TestEvents(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	r, _ := cgClient.Events("DE", "Meetup", "1", "2015-01-01", "2021-12-31", false)
	if r == nil {
		t.Errorf("Error")
	}
}

func TestEventsCountries(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	r, _ := cgClient.EventsCountries()
	if r == nil {
		t.Errorf("Error")
	}
}

func TestEventsTypes(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	r, _ := cgClient.EventsTypes()
	if r == nil {
		t.Errorf("Error")
	}
}
