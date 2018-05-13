package mongo

// Documenter it's an interface that could be common to any documents
// types used to store values on a MongoDB. It contains getters and
// generates to important documents values: _id, created_on and
// updated_on
type Documenter interface {
	New() Documenter
	Validate() error
	Map() (M, error)
	Init(M) error
	ID() ObjectId
	CreatedOn() int64
	UpdatedOn() int64
	GenerateID()
	CalculateCreatedOn()
	CalculateUpdatedOn()
}
