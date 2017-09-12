package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/tnarg/rules_gogo_proto_examples/myteam/myservice"
)

type myServiceImpl struct{}

func newMyServiceImpl() *myServiceImpl {
	return new(myServiceImpl)
}

func (self *myServiceImpl) CreateItem(context.Context, *myteam_myservice.CreateItemRequest) (*myteam_myservice.CreateItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (self *myServiceImpl) GetItem(context.Context, *myteam_myservice.GetItemRequest) (*myteam_myservice.GetItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (self *myServiceImpl) DeleteItem(context.Context, *myteam_myservice.DeleteItemRequest) (*myteam_myservice.DeleteItemResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}
