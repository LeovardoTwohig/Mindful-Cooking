package main

import "fmt"

// Culinary workshops and retreats
type Workshop struct {
	name string
	location string
	workshopLength int
	date string
}

// Create a new workshop 
func NewWorkshop(name string, location string, workshopLength int, date string) *Workshop {
	workshop := &Workshop{
		name: name,
		location: location,
		workshopLength: workshopLength,
		date: date,
	}
	return workshop
}

// Get workshop details
func (w *Workshop) GetWorkshopDetails() (string, string, int, string) {
	return w.name, w.location, w.workshopLength, w.date
}

// Register attendees to the workshop
func (w *Workshop) RegisterAttendees(attendees []string) {
	fmt.Println("The following attendees have registered for the", w.name, "workshop:")
	for _, attendee := range attendees {
		fmt.Println(attendee)
	}
}

// Send reminder email to attendees
func (w *Workshop) SendReminderEmail(attendees []string) {
	fmt.Println("Sending reminder emails to the following workshop attendees:")
	for _, attendee := range attendees {
		fmt.Println("- Sending reminder to", attendee)
	}
}

// Host the workshop
func (w *Workshop) HostWorkshop() {
	fmt.Println("Hosting the", w.name, "workshop at", w.location, "for", w.workshopLength, "days on", w.date)
}

func main() {
	// Create a new workshop
	myWorkshop := NewWorkshop("Culinary Retreat", "Granada, Spain", 5, "June 1 to June 5")

	// Get workshop details
	name, location, workshopLength, date := myWorkshop.GetWorkshopDetails()
	fmt.Println("Name:", name)
	fmt.Println("Location:", location)
	fmt.Println("Workshop Length:", workshopLength)
	fmt.Println("Date:", date)
	
	// Register attendees to the workshop
	attendees := []string{"John Doe", "Jane Doe", "John Smith"}
	myWorkshop.RegisterAttendees(attendees)
	
	// Send reminder email to attendees 
	myWorkshop.SendReminderEmail(attendees)

	// Host the workshop
	myWorkshop.HostWorkshop()
}