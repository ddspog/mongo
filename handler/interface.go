package handler

import (
	"github.com/ddspog/mongo"
	"github.com/ddspog/mongo/model"
)

// Handler it's an interface describing operations common to handlers
// of MongoDB collections.
type Handler interface {
	Link(mongo.Databaser, string)
	Count() (int, error)
	Find(model.Documenter, *model.Documenter) error
	FindAll(model.Documenter, *[]model.Documenter) error
	Insert(model.Documenter) error
	Remove(model.Documenter) error
	RemoveAll(model.Documenter) (*mongo.ChangeInfo, error)
	Update(model.Documenter) error
}
