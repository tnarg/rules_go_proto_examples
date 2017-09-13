package main

import (
	"fmt"

	"golang.org/x/net/context"

	"example.com/api"
)

type myServiceImpl struct{}

func newMyServiceImpl() *myServiceImpl {
	return new(myServiceImpl)
}

func (self *myServiceImpl) CreateItem(context.Context, *api.CreateItemRequest) (*api.CreateItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (self *myServiceImpl) GetItem(context.Context, *api.GetItemRequest) (*api.GetItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (self *myServiceImpl) DeleteItem(context.Context, *api.DeleteItemRequest) (*api.DeleteItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}
