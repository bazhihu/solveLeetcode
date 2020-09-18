package design

import "fmt"

// Operator 被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

// 实现 基类
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 实现 工厂类A
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{OperatorBase: &OperatorBase{}}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// 实现 工厂类B
type MinusOperatorFactory struct {
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func main() {
	var factory OperatorFactory

	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) == 3 {
		fmt.Print("PlusOperatorFactory\n")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 2, 1) == 1 {
		fmt.Print("MinusOperatorFactory\n")
	}
}
