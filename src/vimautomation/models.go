package vimautomation

import (
	"github.com/vmware/govmomi"
)

type VMAction interface {
	Clone() (string, error)
	DeleteVM(vmID string) error
}

type VSphereClient struct {
	Client *govmomi.Client
}
