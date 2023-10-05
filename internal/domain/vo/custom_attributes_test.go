package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	type testScenarios struct {
		dotNotationExpression string
		expectedValue         interface{}
		expectedFound         bool
	}

	customAttributes, err := NewCustomAttributes(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"age":       25,
		"address": map[string]interface{}{
			"streetAddress": "21 2nd Street",
			"city":          "New York",
			"state":         "NY",
			"postalCode":    "10021",
		},
		"phoneNumber": []map[string]interface{}{
			{
				"type":   "home",
				"number": "212 555-1234",
			},
			{
				"type":   "fax",
				"number": "646 555-4567",
			},
		},
	})
	assert.NoError(t, err)
	flattenCustomAttr := customAttributes.Flatten()

	tests := map[string]testScenarios{
		"get-unknown-field":                {dotNotationExpression: "unknown", expectedValue: nil, expectedFound: false},
		"get-unknown-field-inside-map":     {dotNotationExpression: "unknown.unknown", expectedValue: nil, expectedFound: false},
		"get-unknown-field-inside-array":   {dotNotationExpression: "unknown.unknown", expectedValue: nil, expectedFound: false},
		"get-firstName":                    {dotNotationExpression: "firstName", expectedValue: "John", expectedFound: true},
		"get-age":                          {dotNotationExpression: "age", expectedValue: 25, expectedFound: true},
		"get-streetAddress":                {dotNotationExpression: "address.city", expectedValue: "New York", expectedFound: true},
		"get-streetAddress-as-array":       {dotNotationExpression: "address[0].city", expectedValue: nil, expectedFound: false},
		"get-streetAddress.unknown":        {dotNotationExpression: "address.unknown", expectedValue: nil, expectedFound: false},
		"get-phoneNumber-first-position":   {dotNotationExpression: "phoneNumber[0].type", expectedValue: "home", expectedFound: true},
		"get-phoneNumber-second-position":  {dotNotationExpression: "phoneNumber[1].type", expectedValue: "fax", expectedFound: true},
		"get-phoneNumber-unknown-position": {dotNotationExpression: "phoneNumber[3].type", expectedValue: nil, expectedFound: false},
		"get-phoneNumber-as-map":           {dotNotationExpression: "phoneNumber.type", expectedValue: nil, expectedFound: false},
	}
	for testName, testData := range tests {
		t.Run(testName, func(t *testing.T) {
			value, found := flattenCustomAttr[testData.dotNotationExpression]

			assert.Equal(t, testData.expectedFound, found)
			assert.Equal(t, testData.expectedValue, value)
		})
	}
}
