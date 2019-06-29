package service

import (
	"go_code/customer_system/model"
)

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

//添加客户到customer切片
func (cs *CustomerService) Add(customer model.Customer) bool {
	//我们确定一个分配id的规则，就是添加的顺序
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

//根据id删除客户（从切片中）
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)

	if index == -1 {
		//表示不存在
		return false
	}

	//如何从切片中删除一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

//根据ID修改客户
func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	index := cs.FindById(id)

	if index == -1 {
		return false
	}

	cs.customers = append(append(cs.customers[:index], customer), cs.customers[index+1:]...)
	return true

}

//根据ID返回客户信息
func (cs *CustomerService) GetInfoById(id int) model.Customer {
	index := cs.FindById(id)
	if index != -1 {
		return cs.customers[id]
	}

	return model.Customer{}
}

//根据ID查找客户所在切片中对应的下标，如果不存在，返回-1
func (cs *CustomerService) FindById(id int) int {
	index := -1

	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			//找到
			index = i
		}
	}

	return index
}
