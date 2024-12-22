package calculation_test

import (
    "testing"

    "github.com/nmizern/calc_lyceum_go/pkg/calculation"
)

func TestCalc(t *testing.T) {
    tests := []struct {
        name        string
        expression  string
        want        float64
        expectError bool
    }{
        {
            name:       "Simple addition",
            expression: "1+2",
            want:       3,
        },
        {
            name:       "Mixed operations",
            expression: "2*3-1",
            want:       5,
        },
        {
            name:        "Invalid char",
            expression:  "2+3a",
            expectError: true,
        },
        {
            name:        "Division by zero",
            expression:  "5/0",
            expectError: true,
        },
        {
            name:        "Mismatched brackets",
            expression:  "(1+2",
            expectError: true,
        },
        {
            name:       "With spaces",
            expression: "  10   -   4  ",
            want:       6,
        },
        {
            name:       "Multiplication and division",
            expression: "6/2*3",
            want:       9,
        },
        {
            name:       "Unary minus",
            expression: "-5+10",
            want:       5,
        },
        {
            name:       "Unary plus",
            expression: "+5+3",
            want:       8,
        },
        {
            name:       "Decimal numbers",
            expression: "3.5+2.5",
            want:       6,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := calculation.Calc(tt.expression)
            if (err != nil) != tt.expectError {
                t.Errorf("Calc(%q) error = %v, expectError = %v", tt.expression, err, tt.expectError)
                return
            }
            if !tt.expectError && got != tt.want {
                t.Errorf("Calc(%q) = %f, want %f", tt.expression, got, tt.want)
            }
        })
    }
}
