package execution

type Expression interface {
	GetCode() string
	GetName() string
	Evaluate(gc *GlobalContext, tc *ThreadContext) interface{}
}

type baseExpression struct {
	Code string
	Name string
}

type VariableExpression struct {
	baseExpression
}

type AdditionExpression struct {
	baseExpression

	left  Expression
	right Expression
}

func (e *VariableExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	return gc.values[e.Name]
}

func (e *AdditionExpression) Evaluate(gc *GlobalContext, tc *ThreadContext) interface{} {
	lVal := e.left.Evaluate(gc, tc)
	rVal := e.right.Evaluate(gc, tc)

	switch lVal.(type) {
	case int:
		return int(lVal) + int(rVal)
	default:
		return float64(lVal) + float64(rVal)
	}
}

