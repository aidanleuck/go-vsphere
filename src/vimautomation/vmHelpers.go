package vimautomation

import (
	"context"
	"errors"
	"fmt"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
)

func getVMByName(c *govmomi.Client, dc *object.Datacenter, path string) (*object.VirtualMachine, error) {
	finder := find.NewFinder(c.Client)
	vmPath := fmt.Sprintf("%s%s", dc.InventoryPath, path)
	return finder.VirtualMachine(context.TODO(), vmPath)
}

func getVMTemplate(c *govmomi.Client, dc *object.Datacenter, path string) (*object.VirtualMachine, error) {
	finder := find.NewFinder(c.Client)
	vmPath := fmt.Sprintf("%s%s", dc.InventoryPath, path)
	vmTemplate, err := finder.VirtualMachine(context.TODO(), vmPath)
	if err != nil {
		return nil, err
	}
	isTemplate, err := vmTemplate.IsTemplate(context.TODO())

	if err != nil {
		return nil, err
	}
	if !isTemplate {
		return nil, errors.New("No bueno my dude")
	}
	return vmTemplate, nil
}
