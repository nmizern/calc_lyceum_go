package calculation


import (
	"errors"	
)

var (
    ErrInvalidExpression   = errors.New("некорректное выражение")
    ErrInvalidCharacter    = errors.New("некорректный символ")
    ErrInvalidNumberFormat = errors.New("некорректный формат числа")
    ErrDivisionByZero      = errors.New("деление на ноль")
    ErrMismatchedBrackets  = errors.New("несоответствующие скобки")
)