package execution

import "fmt"

type Expression interface {
	GetCode() string
	GetName() string
	Evaluate(gc *GlobalContext, tc *ThreadContext) interface{}
}

type baseExpression struct {
	Code string
	Name string
}

func (e *baseExpression) GetCode() string {
	return e.Code
}

func (e *baseExpression) GetName() string {
	return e.Name
}

type VariableExpression struct {
	baseExpression
}

type AdditionExpression struct {
	baseExpression

	left  Expression
	right Expression
}

type LiteralExpression struct {
	baseExpression
	value interface{}
}

type EqualityExpression struct {
	baseExpression

	left  Expression
	right Expression
}

func NewVariableExpression(name string) *VariableExpression {
	base := baseExpression{
		Code: name,
	}
	return &VariableExpression{base}
}

func NewAdditionExpression(left Expression, right Expression) *AdditionExpression {
	base := baseExpression{
		Code: BinaryOperationCode(left, right, "+"),
	}
	return &AdditionExpression{base, left, right}
}

func NewLiteralExpression(value interface{}) *LiteralExpression {
	return &LiteralExpression{baseExpression{
		Code: fmt.Sprint(value),
	}, value}
}

func NewEqualityExpression(left Expression, right Expression) *EqualityExpression {
	return &EqualityExpression{
		baseExpression{
			Code: BinaryOperationCode(left, right, "=="),
		}, left, right,
	}
}

func (e *EqualityExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	lVal := e.left.Evaluate(gc, tc)
	rVal := e.right.Evaluate(gc, tc)
	return lVal == rVal
}

func (e *LiteralExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	return e.value
}

func (e *VariableExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	return gc.values[e.Name].value
}

func (e *AdditionExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	lVal := e.left.Evaluate(gc, tc)
	rVal := e.right.Evaluate(gc, tc)

	switch lVal.(type) {
	case int:
		return lVal.(int) + rVal.(int)
	default:
		return lVal.(float64) + rVal.(float64)
	}
}

func BinaryOperationCode(l Expression, r Expression, op string) string {
	return l.GetCode() + " " + op + " " + r.GetCode()
}
