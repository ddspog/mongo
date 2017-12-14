package handler

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/model"
)

// Handler it's an interface describing operations common to handlers
// of MongoDB collections.
type Handler interface {
	Link(elements.Databaser, string)
	Count() (int, error)
	Find(model.Documenter, *model.Documenter) error
	FindAll(model.Documenter, *[]model.Documenter) error
	Insert(model.Documenter) error
	Remove(model.Documenter) error
	RemoveAll(model.Documenter) (*elements.ChangeInfo, error)
	Update(model.Documenter) error
}
