package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type S struct {
	Str string
	Int int
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var x interface{} = nil
	assert.Nil(t, x)

	object := S{Str: "Something"}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Str)

	}

}

func TestSomethingAssertManyTimes(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var x interface{} = nil
	assert.Nil(x)

	object := S{Str: "Something"}
	// assert for not nil (good when you expect something)
	if assert.NotNil(object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("Something", object.Str)
	}
}

type MyMockedObject struct {
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

/*
   Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomethingMock(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	// targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	// testObj.AssertExpectations(t)

}

// TestSomethingElse is a second example of how to use our test object to
// make assertions about some target code we are testing.
// This time using a placeholder. Placeholders might be used when the
// data being passed in is normally dynamically generated and cannot be
// predicted beforehand (eg. containing hashes that are time sensitive)
func TestSomethingMockElse(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations with a placeholder in the argument list
	testObj.On("DoSomething", mock.Anything).Return(true, nil)

	// call the code we are testing
	// targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	// testObj.AssertExpectations(t)
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
