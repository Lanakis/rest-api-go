package filter

import "fmt"

const (
	DataTypeStr           = "string"
	DataTypeInt           = "int"
	DataTypeDate          = "date"
	DataTypeBool          = "bool"
	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type option struct {
	limit  int
	fields []Field
}

func NewOption(limit int) Option {
	return &option{limit: limit}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}
type Option interface {
	Limit() int
	AddField(name, operator, value, dataType string) error
	Fields() []Field
}

func (option *option) Limit() int {
	return option.limit
}

func (option *option) AddField(name, operator, value, dataType string) error {
	err := validateOperator(operator)
	if err != nil {
		return err
	}
	field := Field{Name: name, Operator: operator, Value: value, Type: dataType}
	option.fields = append(option.fields, field)
	return nil
}

func (option *option) Fields() []Field {
	return option.fields
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	case OperatorLike:
	case OperatorBetween:

	default:
		return fmt.Errorf("bad operator")

	}
	return nil

}
