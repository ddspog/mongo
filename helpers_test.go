package mongo

// Finisher defines a type that can be finished, closing all pendant
// operations.
type Finisher interface {
	Finish()
}

// Finish calls Finish for all finishers received.
func Finish(fs ...Finisher) {
	for _, f := range fs {
		f.Finish()
	}
}
