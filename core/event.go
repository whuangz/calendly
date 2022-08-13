package core

import (
	"fmt"
	"time"
	"uuid"
)

type Event struct {
	ID       uuid.UUID
	Name     string
	Schedule map[time.Weekday][]Range
}

type SlotParams struct {
	month time.Time
	day   *time.Time
}

func (e Event) GetAvailableSlot(params SlotParams) ([]AvailableDay, error) {
	return nil, fmt.Errorf("not implemented")
}

type Slot struct {
	Start int
}

type AvailableDay struct {
	Slot []Slot
}
