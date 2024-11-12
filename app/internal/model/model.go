package model

// Model represents the data model
type Model struct {
	Message string
}

// NewModel creates a new model instance
func NewModel() *Model {
	return &Model{
		Message: "Welcome to Simple Desktop App!",
	}
}
