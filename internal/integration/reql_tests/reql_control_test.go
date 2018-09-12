// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"gopkg.in/rethinkdb/rethinkdb-go.v5/internal/compare"
)

// Tests RQL control flow structures
func TestControlSuite(t *testing.T) {
	suite.Run(t, new(ControlSuite))
}

type ControlSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ControlSuite) SetupTest() {
	suite.T().Log("Setting up ControlSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_control").Exec(suite.session)
	err = r.DBCreate("db_control").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_control").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_control").TableDrop("table_test_control").Exec(suite.session)
	err = r.DB("db_control").TableCreate("table_test_control").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_control").Table("table_test_control").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_control").TableDrop("table_test_control2").Exec(suite.session)
	err = r.DB("db_control").TableCreate("table_test_control2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_control").Table("table_test_control2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *ControlSuite) TearDownSuite() {
	suite.T().Log("Tearing down ControlSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_control").TableDrop("table_test_control").Exec(suite.session)
		r.DB("db_control").TableDrop("table_test_control2").Exec(suite.session)
		r.DBDrop("db_control").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ControlSuite) TestCases() {
	suite.T().Log("Running ControlSuite: Tests RQL control flow structures")

	table_test_control := r.DB("db_control").Table("table_test_control")
	_ = table_test_control // Prevent any noused variable errors
	table_test_control2 := r.DB("db_control").Table("table_test_control2")
	_ = table_test_control2 // Prevent any noused variable errors

	{
		// control.yaml line #7
		/* 2 */
		var expected_ int = 2
		/* r.expr(1).do(lambda v: v * 2) */

		suite.T().Log("About to run line #7: r.Expr(1).Do(func(v r.Term) interface{} { return r.Mul(v, 2)})")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Do(func(v r.Term) interface{} { return r.Mul(v, 2) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// control.yaml line #12
		/* [0, 1, 2, 3] */
		var expected_ []interface{} = []interface{}{0, 1, 2, 3}
		/* r.expr([0, 1, 2]).do(lambda v: v.append(3)) */

		suite.T().Log("About to run line #12: r.Expr([]interface{}{0, 1, 2}).Do(func(v r.Term) interface{} { return v.Append(3)})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{0, 1, 2}).Do(func(v r.Term) interface{} { return v.Append(3) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #12")
	}

	{
		// control.yaml line #17
		/* 3 */
		var expected_ int = 3
		/* r.do(1, 2, lambda x, y: x + y) */

		suite.T().Log("About to run line #17: r.Do(1, 2, func(x r.Term, y r.Term) interface{} { return r.Add(x, y)})")

		runAndAssert(suite.Suite, expected_, r.Do(1, 2, func(x r.Term, y r.Term) interface{} { return r.Add(x, y) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// control.yaml line #22
		/* 1 */
		var expected_ int = 1
		/* r.do(lambda: 1) */

		suite.T().Log("About to run line #22: r.Do(func() interface{} { return 1})")

		runAndAssert(suite.Suite, expected_, r.Do(func() interface{} { return 1 }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// control.yaml line #38
		/* 1 */
		var expected_ int = 1
		/* r.do(1) */

		suite.T().Log("About to run line #38: r.Do(1)")

		runAndAssert(suite.Suite, expected_, r.Do(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// control.yaml line #53
		/* err("ReqlQueryLogicError", "Expected type ARRAY but found STRING.", [1, 0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type ARRAY but found STRING.")
		/* r.expr('abc').do(lambda v: v.append(3)) */

		suite.T().Log("About to run line #53: r.Expr('abc').Do(func(v r.Term) interface{} { return v.Append(3)})")

		runAndAssert(suite.Suite, expected_, r.Expr("abc").Do(func(v r.Term) interface{} { return v.Append(3) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// control.yaml line #58
		/* err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.", [1, 1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.")
		/* r.expr('abc').do(lambda v: v + 3) */

		suite.T().Log("About to run line #58: r.Expr('abc').Do(func(v r.Term) interface{} { return r.Add(v, 3)})")

		runAndAssert(suite.Suite, expected_, r.Expr("abc").Do(func(v r.Term) interface{} { return r.Add(v, 3) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #58")
	}

	{
		// control.yaml line #63
		/* err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.", [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.")
		/* r.expr('abc').do(lambda v: v + 'def') + 3 */

		suite.T().Log("About to run line #63: r.Expr('abc').Do(func(v r.Term) interface{} { return r.Add(v, 'def')}).Add(3)")

		runAndAssert(suite.Suite, expected_, r.Expr("abc").Do(func(v r.Term) interface{} { return r.Add(v, "def") }).Add(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #63")
	}

	{
		// control.yaml line #78
		/* 5 */
		var expected_ int = 5
		/* r.expr(5).do(r.row) */

		suite.T().Log("About to run line #78: r.Expr(5).Do(r.Row)")

		runAndAssert(suite.Suite, expected_, r.Expr(5).Do(r.Row), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #78")
	}

	{
		// control.yaml line #84
		/* 1 */
		var expected_ int = 1
		/* r.branch(True, 1, 2) */

		suite.T().Log("About to run line #84: r.Branch(true, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(true, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #84")
	}

	{
		// control.yaml line #86
		/* 2 */
		var expected_ int = 2
		/* r.branch(False, 1, 2) */

		suite.T().Log("About to run line #86: r.Branch(false, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(false, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #86")
	}

	{
		// control.yaml line #88
		/* ("c") */
		var expected_ string = "c"
		/* r.branch(1, 'c', False) */

		suite.T().Log("About to run line #88: r.Branch(1, 'c', false)")

		runAndAssert(suite.Suite, expected_, r.Branch(1, "c", false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #88")
	}

	{
		// control.yaml line #90
		/* ([]) */
		var expected_ []interface{} = []interface{}{}
		/* r.branch(null, {}, []) */

		suite.T().Log("About to run line #90: r.Branch(nil, map[interface{}]interface{}{}, []interface{}{})")

		runAndAssert(suite.Suite, expected_, r.Branch(nil, map[interface{}]interface{}{}, []interface{}{}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #90")
	}

	{
		// control.yaml line #93
		/* err("ReqlQueryLogicError", "Expected type DATUM but found DATABASE:", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type DATUM but found DATABASE:")
		/* r.branch(r.db('test'), 1, 2) */

		suite.T().Log("About to run line #93: r.Branch(r.DB('test'), 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(r.DB("db_control"), 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #93")
	}

	{
		// control.yaml line #95
		/* err("ReqlQueryLogicError", "Expected type DATUM but found TABLE:", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type DATUM but found TABLE:")
		/* r.branch(table_test_control, 1, 2) */

		suite.T().Log("About to run line #95: r.Branch(table_test_control, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(table_test_control, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// control.yaml line #97
		/* err("ReqlUserError", "a", []) */
		var expected_ Err = err("ReqlUserError", "a")
		/* r.branch(r.error("a"), 1, 2) */

		suite.T().Log("About to run line #97: r.Branch(r.Error('a'), 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(r.Error("a"), 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #97")
	}

	{
		// control.yaml line #100
		/* 1 */
		var expected_ int = 1
		/* r.branch([], 1, 2) */

		suite.T().Log("About to run line #100: r.Branch([]interface{}{}, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch([]interface{}{}, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// control.yaml line #102
		/* 1 */
		var expected_ int = 1
		/* r.branch({}, 1, 2) */

		suite.T().Log("About to run line #102: r.Branch(map[interface{}]interface{}{}, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(map[interface{}]interface{}{}, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #102")
	}

	{
		// control.yaml line #104
		/* 1 */
		var expected_ int = 1
		/* r.branch("a", 1, 2) */

		suite.T().Log("About to run line #104: r.Branch('a', 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch("a", 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #104")
	}

	{
		// control.yaml line #106
		/* 1 */
		var expected_ int = 1
		/* r.branch(1.2, 1, 2) */

		suite.T().Log("About to run line #106: r.Branch(1.2, 1, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(1.2, 1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #106")
	}

	{
		// control.yaml line #109
		/* 1 */
		var expected_ int = 1
		/* r.branch(True, 1, True, 2, 3) */

		suite.T().Log("About to run line #109: r.Branch(true, 1, true, 2, 3)")

		runAndAssert(suite.Suite, expected_, r.Branch(true, 1, true, 2, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// control.yaml line #111
		/* 1 */
		var expected_ int = 1
		/* r.branch(True, 1, False, 2, 3) */

		suite.T().Log("About to run line #111: r.Branch(true, 1, false, 2, 3)")

		runAndAssert(suite.Suite, expected_, r.Branch(true, 1, false, 2, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #111")
	}

	{
		// control.yaml line #113
		/* 2 */
		var expected_ int = 2
		/* r.branch(False, 1, True, 2, 3) */

		suite.T().Log("About to run line #113: r.Branch(false, 1, true, 2, 3)")

		runAndAssert(suite.Suite, expected_, r.Branch(false, 1, true, 2, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #113")
	}

	{
		// control.yaml line #115
		/* 3 */
		var expected_ int = 3
		/* r.branch(False, 1, False, 2, 3) */

		suite.T().Log("About to run line #115: r.Branch(false, 1, false, 2, 3)")

		runAndAssert(suite.Suite, expected_, r.Branch(false, 1, false, 2, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #115")
	}

	{
		// control.yaml line #118
		/* err("ReqlQueryLogicError", "Cannot call `branch` term with an even number of arguments.") */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot call `branch` term with an even number of arguments.")
		/* r.branch(True, 1, True, 2) */

		suite.T().Log("About to run line #118: r.Branch(true, 1, true, 2)")

		runAndAssert(suite.Suite, expected_, r.Branch(true, 1, true, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #118")
	}

	{
		// control.yaml line #122
		/* err("ReqlUserError", "Hello World", [0]) */
		var expected_ Err = err("ReqlUserError", "Hello World")
		/* r.error('Hello World') */

		suite.T().Log("About to run line #122: r.Error('Hello World')")

		runAndAssert(suite.Suite, expected_, r.Error("Hello World"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #122")
	}

	{
		// control.yaml line #125
		/* err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.")
		/* r.error(5) */

		suite.T().Log("About to run line #125: r.Error(5)")

		runAndAssert(suite.Suite, expected_, r.Error(5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #125")
	}

	{
		// control.yaml line #140
		/* 2 */
		var expected_ int = 2
		/* r.js('1 + 1') */

		suite.T().Log("About to run line #140: r.JS('1 + 1')")

		runAndAssert(suite.Suite, expected_, r.JS("1 + 1"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #140")
	}

	{
		// control.yaml line #143
		/* 4 */
		var expected_ int = 4
		/* r.js('1 + 1; 2 + 2') */

		suite.T().Log("About to run line #143: r.JS('1 + 1; 2 + 2')")

		runAndAssert(suite.Suite, expected_, r.JS("1 + 1; 2 + 2"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #143")
	}

	{
		// control.yaml line #146
		/* 3 */
		var expected_ int = 3
		/* r.do(1, 2, r.js('(function(a, b) { return a + b; })')) */

		suite.T().Log("About to run line #146: r.Do(1, 2, r.JS('(function(a, b) { return a + b; })'))")

		runAndAssert(suite.Suite, expected_, r.Do(1, 2, r.JS("(function(a, b) { return a + b; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #146")
	}

	{
		// control.yaml line #149
		/* 2 */
		var expected_ int = 2
		/* r.expr(1).do(r.js('(function(x) { return x + 1; })')) */

		suite.T().Log("About to run line #149: r.Expr(1).Do(r.JS('(function(x) { return x + 1; })'))")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Do(r.JS("(function(x) { return x + 1; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #149")
	}

	{
		// control.yaml line #152
		/* 'foobar' */
		var expected_ string = "foobar"
		/* r.expr('foo').do(r.js('(function(x) { return x + "bar"; })')) */

		suite.T().Log("About to run line #152: r.Expr('foo').Do(r.JS('(function(x) { return x + \\'bar\\'; })'))")

		runAndAssert(suite.Suite, expected_, r.Expr("foo").Do(r.JS("(function(x) { return x + \"bar\"; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #152")
	}

	{
		// control.yaml line #157
		/* 3 */
		var expected_ int = 3
		/* r.js('1 + 2', timeout=1.2) */

		suite.T().Log("About to run line #157: r.JS('1 + 2').OptArgs(r.JSOpts{Timeout: 1.2, })")

		runAndAssert(suite.Suite, expected_, r.JS("1 + 2").OptArgs(r.JSOpts{Timeout: 1.2}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #157")
	}

	{
		// control.yaml line #161
		/* err("ReqlQueryLogicError", "Query result must be of type DATUM, GROUPED_DATA, or STREAM (got FUNCTION).", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Query result must be of type DATUM, GROUPED_DATA, or STREAM (got FUNCTION).")
		/* r.js('(function() { return 1; })') */

		suite.T().Log("About to run line #161: r.JS('(function() { return 1; })')")

		runAndAssert(suite.Suite, expected_, r.JS("(function() { return 1; })"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #161")
	}

	{
		// control.yaml line #164
		/* err("ReqlQueryLogicError", "SyntaxError: Unexpected token (", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "SyntaxError: Unexpected token (")
		/* r.js('function() { return 1; }') */

		suite.T().Log("About to run line #164: r.JS('function() { return 1; }')")

		runAndAssert(suite.Suite, expected_, r.JS("function() { return 1; }"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #164")
	}

	{
		// control.yaml line #168
		/* 1 */
		var expected_ int = 1
		/* r.do(1, 2, r.js('(function(a) { return a; })')) */

		suite.T().Log("About to run line #168: r.Do(1, 2, r.JS('(function(a) { return a; })'))")

		runAndAssert(suite.Suite, expected_, r.Do(1, 2, r.JS("(function(a) { return a; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #168")
	}

	{
		// control.yaml line #171
		/* 1 */
		var expected_ int = 1
		/* r.do(1, 2, r.js('(function(a, b, c) { return a; })')) */

		suite.T().Log("About to run line #171: r.Do(1, 2, r.JS('(function(a, b, c) { return a; })'))")

		runAndAssert(suite.Suite, expected_, r.Do(1, 2, r.JS("(function(a, b, c) { return a; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #171")
	}

	{
		// control.yaml line #174
		/* err("ReqlQueryLogicError", "Cannot convert javascript `undefined` to ql::datum_t.", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert javascript `undefined` to ql::datum_t.")
		/* r.do(1, 2, r.js('(function(a, b, c) { return c; })')) */

		suite.T().Log("About to run line #174: r.Do(1, 2, r.JS('(function(a, b, c) { return c; })'))")

		runAndAssert(suite.Suite, expected_, r.Do(1, 2, r.JS("(function(a, b, c) { return c; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #174")
	}

	{
		// control.yaml line #177
		/* ([2, 3]) */
		var expected_ []interface{} = []interface{}{2, 3}
		/* r.expr([1, 2, 3]).filter(r.js('(function(a) { return a >= 2; })')) */

		suite.T().Log("About to run line #177: r.Expr([]interface{}{1, 2, 3}).Filter(r.JS('(function(a) { return a >= 2; })'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Filter(r.JS("(function(a) { return a >= 2; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #177")
	}

	{
		// control.yaml line #180
		/* ([2, 3, 4]) */
		var expected_ []interface{} = []interface{}{2, 3, 4}
		/* r.expr([1, 2, 3]).map(r.js('(function(a) { return a + 1; })')) */

		suite.T().Log("About to run line #180: r.Expr([]interface{}{1, 2, 3}).Map(r.JS('(function(a) { return a + 1; })'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Map(r.JS("(function(a) { return a + 1; })")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #180")
	}

	{
		// control.yaml line #183
		/* err("ReqlQueryLogicError", "Expected type FUNCTION but found DATUM:", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type FUNCTION but found DATUM:")
		/* r.expr([1, 2, 3]).map(r.js('1')) */

		suite.T().Log("About to run line #183: r.Expr([]interface{}{1, 2, 3}).Map(r.JS('1'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Map(r.JS("1")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #183")
	}

	{
		// control.yaml line #186
		/* err("ReqlQueryLogicError", "Cannot convert javascript `undefined` to ql::datum_t.", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert javascript `undefined` to ql::datum_t.")
		/* r.expr([1, 2, 3]).filter(r.js('(function(a) {})')) */

		suite.T().Log("About to run line #186: r.Expr([]interface{}{1, 2, 3}).Filter(r.JS('(function(a) {})'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Filter(r.JS("(function(a) {})")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #186")
	}

	{
		// control.yaml line #190
		/* err("ReqlQueryLogicError", "Expected type FUNCTION but found DATUM:", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type FUNCTION but found DATUM:")
		/* r.expr([1, 2, 3]).map(1) */

		suite.T().Log("About to run line #190: r.Expr([]interface{}{1, 2, 3}).Map(1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Map(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #190")
	}

	{
		// control.yaml line #193
		/* ([1, 2, 3]) */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([1, 2, 3]).filter('foo') */

		suite.T().Log("About to run line #193: r.Expr([]interface{}{1, 2, 3}).Filter('foo')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Filter("foo"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #193")
	}

	{
		// control.yaml line #195
		/* ([1, 2, 4]) */
		var expected_ []interface{} = []interface{}{1, 2, 4}
		/* r.expr([1, 2, 4]).filter([]) */

		suite.T().Log("About to run line #195: r.Expr([]interface{}{1, 2, 4}).Filter([]interface{}{})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 4}).Filter([]interface{}{}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #195")
	}

	{
		// control.yaml line #197
		/* ([]) */
		var expected_ []interface{} = []interface{}{}
		/* r.expr([1, 2, 3]).filter(null) */

		suite.T().Log("About to run line #197: r.Expr([]interface{}{1, 2, 3}).Filter(nil)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Filter(nil), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #197")
	}

	{
		// control.yaml line #200
		/* ([]) */
		var expected_ []interface{} = []interface{}{}
		/* r.expr([1, 2, 4]).filter(False) */

		suite.T().Log("About to run line #200: r.Expr([]interface{}{1, 2, 4}).Filter(false)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 4}).Filter(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #200")
	}

	{
		// control.yaml line #205
		/* 0 */
		var expected_ int = 0
		/* table_test_control.count() */

		suite.T().Log("About to run line #205: table_test_control.Count()")

		runAndAssert(suite.Suite, expected_, table_test_control.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #205")
	}

	{
		// control.yaml line #210
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':3}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 3}
		/* r.expr([1, 2, 3]).for_each(lambda row:table_test_control.insert({ 'id':row })) */

		suite.T().Log("About to run line #210: r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return table_test_control.Insert(map[interface{}]interface{}{'id': row, })})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return table_test_control.Insert(map[interface{}]interface{}{"id": row}) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #210")
	}

	{
		// control.yaml line #214
		/* 3 */
		var expected_ int = 3
		/* table_test_control.count() */

		suite.T().Log("About to run line #214: table_test_control.Count()")

		runAndAssert(suite.Suite, expected_, table_test_control.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #214")
	}

	{
		// control.yaml line #219
		/* ({'deleted':0.0,'replaced':9,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 9, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* r.expr([1,2,3]).for_each(lambda row:table_test_control.update({'foo':row})) */

		suite.T().Log("About to run line #219: r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return table_test_control.Update(map[interface{}]interface{}{'foo': row, })})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} {
			return table_test_control.Update(map[interface{}]interface{}{"foo": row})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #219")
	}

	{
		// control.yaml line #225
		/* {'first_error':"Duplicate primary key `id`:\n{\n\t\"foo\":\t3,\n\t\"id\":\t1\n}\n{\n\t\"id\":\t1\n}",'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':3,'skipped':0.0,'inserted':3} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"first_error": "Duplicate primary key `id`:\n{\n\t\"foo\":\t3,\n\t\"id\":\t1\n}\n{\n\t\"id\":\t1\n}", "deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 3, "skipped": 0.0, "inserted": 3}
		/* r.expr([1,2,3]).for_each(lambda row:[table_test_control.insert({ 'id':row }), table_test_control.insert({ 'id':row*10 })]) */

		suite.T().Log("About to run line #225: r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return []interface{}{table_test_control.Insert(map[interface{}]interface{}{'id': row, }), table_test_control.Insert(map[interface{}]interface{}{'id': r.Mul(row, 10), })}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} {
			return []interface{}{table_test_control.Insert(map[interface{}]interface{}{"id": row}), table_test_control.Insert(map[interface{}]interface{}{"id": r.Mul(row, 10)})}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #225")
	}

	{
		// control.yaml line #229
		/* 6 */
		var expected_ int = 6
		/* table_test_control.count() */

		suite.T().Log("About to run line #229: table_test_control.Count()")

		runAndAssert(suite.Suite, expected_, table_test_control.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #229")
	}

	{
		// control.yaml line #232
		/* ({'deleted':0.0,'replaced':0.0,'generated_keys':arrlen(3,uuid()),'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':3}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "generated_keys": arrlen(3, compare.IsUUID()), "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 3}
		/* r.expr([1, 2, 3]).for_each( table_test_control2.insert({}) ) */

		suite.T().Log("About to run line #232: r.Expr([]interface{}{1, 2, 3}).ForEach(table_test_control2.Insert(map[interface{}]interface{}{}))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(table_test_control2.Insert(map[interface{}]interface{}{})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #232")
	}

	{
		// control.yaml line #235
		/* 1 */
		var expected_ int = 1
		/* table_test_control2.count() */

		suite.T().Log("About to run line #235: table_test_control2.Count()")

		runAndAssert(suite.Suite, expected_, table_test_control2.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #235")
	}

	{
		// control.yaml line #240
		/* ({'deleted':0.0,'replaced':36,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 36, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* r.expr([1,2,3]).for_each(lambda row:[table_test_control.update({'foo':row}), table_test_control.update({'bar':row})]) */

		suite.T().Log("About to run line #240: r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return []interface{}{table_test_control.Update(map[interface{}]interface{}{'foo': row, }), table_test_control.Update(map[interface{}]interface{}{'bar': row, })}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} {
			return []interface{}{table_test_control.Update(map[interface{}]interface{}{"foo": row}), table_test_control.Update(map[interface{}]interface{}{"bar": row})}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #240")
	}

	{
		// control.yaml line #245
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':3}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 3}
		/* r.expr([1, 2, 3]).for_each( table_test_control2.insert({ 'id':r.row }) ) */

		suite.T().Log("About to run line #245: r.Expr([]interface{}{1, 2, 3}).ForEach(table_test_control2.Insert(map[interface{}]interface{}{'id': r.Row, }))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(table_test_control2.Insert(map[interface{}]interface{}{"id": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #245")
	}

	{
		// control.yaml line #249
		/* err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.", [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.")
		/* r.expr([1, 2, 3]).for_each(1) */

		suite.T().Log("About to run line #249: r.Expr([]interface{}{1, 2, 3}).ForEach(1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #249")
	}

	{
		// control.yaml line #252
		/* err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.", [1, 1]) */
		var expected_ Err = err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.")
		/* r.expr([1, 2, 3]).for_each(lambda x:x) */

		suite.T().Log("About to run line #252: r.Expr([]interface{}{1, 2, 3}).ForEach(func(x r.Term) interface{} { return x})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(x r.Term) interface{} { return x }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #252")
	}

	{
		// control.yaml line #257
		/* err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.", [1, 1]) */
		var expected_ Err = err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.  Expected type ARRAY but found NUMBER.")
		/* r.expr([1, 2, 3]).for_each(r.row) */

		suite.T().Log("About to run line #257: r.Expr([]interface{}{1, 2, 3}).ForEach(r.Row)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(r.Row), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #257")
	}

	{
		// control.yaml line #262
		/* err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.", [1, 1]) */
		var expected_ Err = err("ReqlQueryLogicError", "FOR_EACH expects one or more basic write queries.")
		/* r.expr([1, 2, 3]).for_each(lambda row:table_test_control) */

		suite.T().Log("About to run line #262: r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return table_test_control})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(func(row r.Term) interface{} { return table_test_control }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #262")
	}

	{
		// control.yaml line #271
		/* ({'deleted':0.0,'replaced':0.0,'generated_keys':arrlen(1,uuid()),'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "generated_keys": arrlen(1, compare.IsUUID()), "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 1}
		/* r.expr(1).do(table_test_control.insert({'foo':r.row})) */

		suite.T().Log("About to run line #271: r.Expr(1).Do(table_test_control.Insert(map[interface{}]interface{}{'foo': r.Row, }))")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Do(table_test_control.Insert(map[interface{}]interface{}{"foo": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #271")
	}

	{
		// control.yaml line #275
		/* ({'deleted':0.0,'replaced':0.0,'generated_keys':arrlen(1,uuid()),'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "generated_keys": arrlen(1, compare.IsUUID()), "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 1}
		/* r.expr([1, 2])[0].do(table_test_control.insert({'foo':r.row})) */

		suite.T().Log("About to run line #275: r.Expr([]interface{}{1, 2}).AtIndex(0).Do(table_test_control.Insert(map[interface{}]interface{}{'foo': r.Row, }))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).AtIndex(0).Do(table_test_control.Insert(map[interface{}]interface{}{"foo": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #275")
	}

	{
		// control.yaml line #280
		/* err('ReqlCompileError', 'Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.', [0]) */
		var expected_ Err = err("ReqlCompileError", "Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.")
		/* r.expr([1, 2]).map(table_test_control.insert({'foo':r.row})) */

		suite.T().Log("About to run line #280: r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{'foo': r.Row, }))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{"foo": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #280")
	}

	{
		// control.yaml line #284
		/* err('ReqlCompileError', 'Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.', [0]) */
		var expected_ Err = err("ReqlCompileError", "Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.")
		/* r.expr([1, 2]).map(r.db('test').table_create('table_create_failure')) */

		suite.T().Log("About to run line #284: r.Expr([]interface{}{1, 2}).Map(r.DB('test').TableCreate('table_create_failure'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Map(r.DB("db_control").TableCreate("table_create_failure")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #284")
	}

	{
		// control.yaml line #287
		/* err('ReqlCompileError', 'Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.', [0]) */
		var expected_ Err = err("ReqlCompileError", "Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.")
		/* r.expr([1, 2]).map(table_test_control.insert({'foo':r.row}).get_field('inserted')) */

		suite.T().Log("About to run line #287: r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{'foo': r.Row, }).Field('inserted'))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{"foo": r.Row}).Field("inserted")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #287")
	}

	{
		// control.yaml line #291
		/* err('ReqlCompileError', 'Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.', [0]) */
		var expected_ Err = err("ReqlCompileError", "Cannot nest writes or meta ops in stream operations.  Use FOR_EACH instead.")
		/* r.expr([1, 2]).map(table_test_control.insert({'foo':r.row}).get_field('inserted').add(5)) */

		suite.T().Log("About to run line #291: r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{'foo': r.Row, }).Field('inserted').Add(5))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Map(table_test_control.Insert(map[interface{}]interface{}{"foo": r.Row}).Field("inserted").Add(5)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #291")
	}

	{
		// control.yaml line #295
		/* partial({'tables_created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 1})
		/* r.expr(1).do(r.db('test').table_create('table_create_success')) */

		suite.T().Log("About to run line #295: r.Expr(1).Do(r.DB('test').TableCreate('table_create_success'))")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Do(r.DB("db_control").TableCreate("table_create_success")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #295")
	}
}
