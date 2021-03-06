package entity

import (
	"tests/helper"
)

// RunTests ...
func RunTests() helper.TestResult {
	var testResults []helper.TestResult
	testResults = append(testResults, TestGetAllEntity())
	testResults = append(testResults, TestCreateEntity())
	testResults = append(testResults, TestUpdateEntity())
	testResults = append(testResults, TestDeleteEntity())

	helper.PrintResults(testResults)
	return helper.CountTests(testResults)
}

// TestSample ...
func TestSample() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.Login",
	}

	var testCases []helper.TestMeasure

	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Expected: 400,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Post("http://localhost:3000/api/v1/login", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)
	}
	return result
}
