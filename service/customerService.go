package service

import "go_code/customer_system/model"

//该CustomerService，完成对Customer的操作，包括
//增删改查

type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多个客户
	//该字段后面，还可以作为新客户的ID+1
	customerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户

	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "13212349800", "zs@sohu.com")

	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

