package viewmodel

import (
	"app/internal/model"
	"fyne.io/fyne/v2/data/binding"
)

// ViewModel represents the view model
type ViewModel struct {
	message binding.String
}

// NewViewModel creates a new view model instance
func NewViewModel(model *model.Model) *ViewModel {
	vm := &ViewModel{}
	vm.message = binding.NewString()
	vm.message.Set(model.Message)
	return vm
}

// GetMessage returns the message binding
func (vm *ViewModel) GetMessage() binding.String {
	return vm.message
}

// SetMessage sets the message in the model
func (vm *ViewModel) SetMessage(msg string) {
	vm.message.Set(msg)
}
