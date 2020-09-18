/**
外观模式
主要解决：降低访问复杂系统的内部子系统的复杂度，简化客户端与之的接口。
意图：为子系统的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
如何解决：客户端不与系统耦合，外观类与系统耦合
*/
package design

import "fmt"

// 接口A
type AModuleAPI interface {
	TestA() string
}

// 接口B
type BModuleAPI interface {
	TestB() string
}

// 外观模式 对外接口
type API interface {
	Test() string
}

type aModuleImpl struct{}
type bModuleImpl struct{}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (*aModuleImpl) TestA() string {
	return "i am A"
}

func (*bModuleImpl) TestB() string {
	return "i am B"
}

func (a *apiImpl) Test() string {
	aStr := a.a.TestA()
	bStr := a.b.TestB()
	return aStr + bStr
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func main() {
	model := NewAPI()
	fmt.Println(model.Test())
}
