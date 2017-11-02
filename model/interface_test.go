package model

import (
	"testing"

	"github.com/ddspog/trialtbl"
	"gopkg.in/mgo.v2/bson"
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
			p.IdV = bson.ObjectIdHex(f[0].(string))

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
			p.IdV = bson.ObjectIdHex(f[0].(string))
			p.CreatedOnV = f[1].(int64)
			p.UpdatedOnV = f[2].(int64)

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
			sig := "d.CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn() method
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestDocumenterSetter checks if a type embedding Document has
// functional setters.
func TestDocumenterSetter(t *testing.T) {
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

			// Cast product to Documenter.
			d = p

			d.SetId(bson.ObjectIdHex(f[0].(string)))
			d.SetCreatedOn(f[1].(int64))
			d.SetUpdatedOn(f[2].(int64))

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
			sig := "d.CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn() method
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestDocumenterCalculation checks the calculation of the created_on
// and updated_on attributes. Mainly if the function used return bigger
// values over the time.
func TestDocumenterCalculation(t *testing.T) {
	make, _ := NewMockModelSetup(t)
	defer make.Finish()

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
			trialtbl.NewTrial(true, int64(10)),
			trialtbl.NewTrial(true, int64(20)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false, int64(10)),
			trialtbl.NewTrial(true, int64(0)),
			trialtbl.NewTrial(true, int64(1000000)),
			trialtbl.NewTrial(true, int64(2000000)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false, before),
			trialtbl.NewTrial(false, before),
			trialtbl.NewTrial(true, before*2),
			trialtbl.NewTrial(true, before*4),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		p := newProduct()
		var d Documenter = p

		// Test CreatedOn() initial value
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			make.NowInMilli().Returns(int64(0))
			val := d.CreatedOn() == f[0].(int64)
			sig := "d.CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn() initial value
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			make.NowInMilli().Returns(int64(0))
			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Check CreatedOn() value
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			make.NowInMilli().Returns(f[0].(int64))
			d.CalculateCreatedOn()

			val := d.CreatedOn() == f[0].(int64)
			sig := "d.CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Check UpdatedOn() value
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			make.NowInMilli().Returns(f[0].(int64))
			d.CalculateUpdatedOn()

			val := d.UpdatedOn() == f[0].(int64)
			sig := "d.UpdatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestDocumenterIdGeneration checks the generation of the _id
// attributes.
func TestDocumenterIdGeneration(t *testing.T) {
	make, _ := NewMockModelSetup(t)
	defer make.Finish()

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
		p := newProduct()
		var d Documenter = p

		// Check Id() value
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			make.NewId().Returns(bson.ObjectId(f[0].(string)))
			d.GenerateId()

			val := d.Id().Hex() != f[0].(string)
			sig := "d.Id().Hex() != \"%s\""
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}
