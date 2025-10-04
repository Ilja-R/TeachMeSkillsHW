package main

import "fmt"

// Observer Pattern Example: Button Click Handlers

// ButtonClickHandler defines the interface for handling button clicks, it is an observer.
type ButtonClickHandler interface {
	HandleClick(button *Button)
}

// Button represents a button that can be clicked, it is the subject being observed.
type Button struct {
	Name     string
	handlers []ButtonClickHandler
}

// AddHandler registers a new handler to the button.
func (b *Button) AddHandler(handler ButtonClickHandler) {
	b.handlers = append(b.handlers, handler)
}

// Click simulates a button click and notifies all registered handlers.
func (b *Button) Click() {
	for _, handler := range b.handlers {
		handler.HandleClick(b)
	}
}

// LogHandler is a concrete observer that logs button clicks.
type LogHandler struct {
}

// HandleClick logs the button click event.
func (*LogHandler) HandleClick(button *Button) {
	fmt.Printf("Button %s clicked!\n", button.Name)
}

// ActionHandler is a concrete observer that performs an action when the button is clicked.
type ActionHandler struct {
	// Stored as string for simplicity, but could be a function or more complex type.
	actionName string
}

// HandleClick performs the action associated with the button click.
func (a *ActionHandler) HandleClick(button *Button) {
	fmt.Printf("Action %s triggered by button %s!\n", a.actionName, button.Name)
}

// TelemetryHandler is a concrete observer that tracks the number of times the button has been clicked.
type TelemetryHandler struct {
	timesClicked int
}

// HandleClick increments the click count and logs it.
func (t *TelemetryHandler) HandleClick(button *Button) {
	t.timesClicked++
	fmt.Printf("Button %s clicked %d times!\n", button.Name, t.timesClicked)
}

func main() {
	// Create handlers
	logHandler := LogHandler{}
	actionHandler := ActionHandler{actionName: "SAVE USER"}
	telemetryHandler := TelemetryHandler{}

	// Create button
	button := Button{Name: "Submit"}

	// Register handlers with the button
	button.AddHandler(&logHandler)
	button.AddHandler(&actionHandler)
	button.AddHandler(&telemetryHandler)

	// Simulate button clicks
	for i := 0; i < 3; i++ {
		button.Click()
		fmt.Println("-----")
	}
}
