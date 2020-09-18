/**
抽象工厂模式
*/
package design

import "fmt"

// 接口定义

type OrderMainDao interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

// 对象定义 A

type RDBMainDao struct {
}

func (*RDBMainDao) SaveOrderMain() {
	fmt.Print("RDB MAIN SAVE\n")
}

type RDBMDetailDao struct {
}

func (*RDBMDetailDao) SaveOrderDetail() {
	fmt.Print("RDB DETAIL SAVE\n")
}

type RDBDaoFactory struct {
}

func (*RDBDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (*RDBDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBMDetailDao{}
}

// 对象定义 B

type XMLMainDao struct {
}

func (*XMLMainDao) SaveOrderMain() {
	fmt.Print("XML MAIN SAVE\n")
}

type XMLMDetailDao struct {
}

func (*XMLMDetailDao) SaveOrderDetail() {
	fmt.Print("XML DETAIL SAVE\n")
}

type XMLDaoFactory struct {
}

func (*XMLDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (*XMLDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLMDetailDao{}
}

func getMainAndDetail(factory DaoFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func main() {
	var factory DaoFactory
	factory = &RDBDaoFactory{}
	getMainAndDetail(factory)

	factory = &XMLDaoFactory{}
	getMainAndDetail(factory)
}
