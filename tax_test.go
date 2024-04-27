package main

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	type testCase struct {
		totalIncome float64
		wht         float64
		allowances  []Allowance
		expectedTax float64
	}

	testCases := []testCase{

		{150000.0, 0.0, []Allowance{}, 0.0},
		{210000.0, 0.0, []Allowance{}, 0.0},
		{330000.0, 0.0, []Allowance{}, 12000.00},
		{560000.0, 0.0, []Allowance{}, 35000.0},
		{1000000.0, 0.0, []Allowance{}, 101000.00},
		{1060000.0, 0.0, []Allowance{}, 110000.0},
		{2000000.0, 0.0, []Allowance{}, 298000.00},
		{2060000.0, 0.0, []Allowance{}, 310000.0},
		{10000000.0, 0.0, []Allowance{}, 3089000.00},
	}

	for _, tc := range testCases {
		tax, err := CalculateTax(tc.totalIncome, tc.wht, tc.allowances)
		if err != nil {
			t.Fatalf("CalculateTax failed with error: %v", err)
		}

		if tax != tc.expectedTax {
			t.Errorf("Expected tax to be %.2f but got %.2f", tc.expectedTax, tax)
		}
	}
}