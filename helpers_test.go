package mongo

// finisher defines a type that can be finished, closing all pendant
// operations.
type finisher interface {
	Finish()
}

// finish calls finish for all finishers received.
func finish(fs ...finisher) {
	for _, f := range fs {
		f.Finish()
	}
}
