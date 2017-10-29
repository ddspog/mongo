package model

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/ddspog/trialtbl"
)

// TestDocumenterCast checks if the type Document can be casted
// without any problems.
func TestDocumenterCast(t *testing.T) {
	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product1id),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product2id),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			p := newProduct()
			p.id = bson.ObjectIdHex(f[0].(string))

			var d Documenter = p

			// Verify if Id function return correct value.
			val := d.Id() == bson.ObjectIdHex(f[0].(string))
			sig := "d.Id() == bson.ObjectIdHex(\"%s\")"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestDocumenterCreation checks if a type embedding Document has
// functional getters.
func TestDocumenterCreation(t *testing.T) {
	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid, int64(123), int64(321)),
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(true, int64(123)),
			trialtbl.NewTrial(true, int64(321)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product1id, int64(10), int64(1)),
			trialtbl.NewTrial(true, product1id),
			trialtbl.NewTrial(true, int64(10)),
			trialtbl.NewTrial(true, int64(1)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product2id, int64(20), int64(2)),
			trialtbl.NewTrial(true, product2id),
			trialtbl.NewTrial(true, int64(20)),
			trialtbl.NewTrial(true, int64(2)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid, int64(123), int64(321)),
			trialtbl.NewTrial(false, product1id),
			trialtbl.NewTrial(false, int64(321)),
			trialtbl.NewTrial(false, int64(123)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid, int64(123), int64(321)),
			trialtbl.NewTrial(false, product2id),
			trialtbl.NewTrial(false, int64(10)),
			trialtbl.NewTrial(false, int64(2)),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var d Documenter

		// Utility Trial
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			p := newProduct()
			p.id = bson.ObjectIdHex(f[0].(string))
			p.createdOn = f[1].(int64)
			p.updatedOn = f[2].(int64)

			// Cast product to Documenter.
			d = p

			r = trialtbl.NewResult(true, "true")
			return
		})

		// Test Id() method
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.Id() == bson.ObjectIdHex(f[0].(string))
			sig := "d.Id() == bson.ObjectIdHex(\"%s\")"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test CreatedOn() method
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.CreatedOn() == f[0].(int64)
			sig := "d.CreatedOn() == \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn() method
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestDocumenterCalculation checks the calculation of the created_on
// and updated_on attributes. Mainly if the function used return bigger
// values over the time.
func TestDocumenterCalculation(t *testing.T) {
	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(0)),
			trialtbl.NewTrial(true, int64(0)),
			trialtbl.NewTrial(true, before),
			trialtbl.NewTrial(true, before),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(0)),
			trialtbl.NewTrial(false, int64(10)),
			trialtbl.NewTrial(true, before),
			trialtbl.NewTrial(false, NowInInt64()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false, int64(10)),
			trialtbl.NewTrial(true, int64(0)),
			trialtbl.NewTrial(false, NowInInt64()),
			trialtbl.NewTrial(true, before),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false, before),
			trialtbl.NewTrial(false, before),
			trialtbl.NewTrial(false, NowInInt64()),
			trialtbl.NewTrial(false, NowInInt64()),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		p := newProduct()
		var d Documenter = p

		// Test CreatedOn initial value
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.CreatedOn() == f[0].(int64)
			sig := "d.CreatedOn() == \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn initial value
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})

		d.CalculateCreatedOn()

		// Check a time smaller than CreatedOn
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.CreatedOn() > f[0].(int64)
			sig := "d.CreatedOn() > \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})

		d.CalculateUpdatedOn()

		// Check a time smaller than UpdatedOn
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.UpdatedOn() > f[0].(int64)
			sig := "d.UpdatedOn() > \"%v\""
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}
