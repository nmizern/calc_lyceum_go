package calculation

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {

	expr := strings.ReplaceAll(expression, " ", "")

	if expr == "" {
        return 0, ErrInvalidExpression
    }

	var ops []rune
	var nums []float64

	applyOp := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return ErrInvalidExpression
		}
		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		op := ops[len(ops)-1]
		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]
		var result float64
		switch op {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return ErrDivisionByZero
			}
			result = a / b
		default:
			return fmt.Errorf("%w: %c", ErrInvalidCharacter, op)
		}
		nums = append(nums, result)
		return nil
	}

	precedence := func(op rune) int {
		switch op {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		default:
			return 0
		}
	}

	for i := 0; i < len(expr); {
		ch := rune(expr[i])
		if unicode.IsSpace(ch) {
			i++
			continue
		} else if ch == '(' {
			ops = append(ops, ch)
			i++
		} else if ch == ')' {

			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 || ops[len(ops)-1] != '(' {
				return 0, ErrMismatchedBrackets
			}
			ops = ops[:len(ops)-1]
			i++
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {

			if (i == 0 || expr[i-1] == '(' || expr[i-1] == '+' || expr[i-1] == '-' || expr[i-1] == '*' || expr[i-1] == '/') && (ch == '+' || ch == '-') {

				numStr := string(ch)
				i++

				start := i
				dotCount := 0
				for i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.') {
					if expr[i] == '.' {
						dotCount++
						if dotCount > 1 {
							return 0, ErrInvalidNumberFormat
						}
					}
					i++
				}
				if start == i {
					return 0, ErrInvalidExpression
				}
				numStr += expr[start:i]
				num, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return 0, ErrInvalidNumberFormat
				}
				nums = append(nums, num)
			} else {

				for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
					if err := applyOp(); err != nil {
						return 0, err
					}
				}
				ops = append(ops, ch)
				i++
			}
		} else if unicode.IsDigit(ch) || ch == '.' {

			start := i
			dotCount := 0
			for i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.') {
				if expr[i] == '.' {
					dotCount++
					if dotCount > 1 {
						return 0, ErrInvalidNumberFormat
					}
				}
				i++
			}
			numStr := expr[start:i]
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, ErrInvalidNumberFormat
			}
			nums = append(nums, num)
		} else {
			return 0, fmt.Errorf("%w: %c", ErrInvalidCharacter, ch)
		}
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == '(' || ops[len(ops)-1] == ')' {
			return 0, ErrMismatchedBrackets
		}
		if err := applyOp(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, ErrInvalidExpression
	}
	return nums[0], nil
}
