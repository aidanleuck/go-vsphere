package vimautomation

import (
	"context"
	"vsphere_module/src/common"

	"github.com/vmware/govmomi"
)

func Connect(lanierConfig *common.LanierConfig) (*govmomi.Client, error) {
	ctx := context.Background()
	client, err := govmomi.NewClient(ctx, &lanierConfig.URL, true)
	if err != nil {
		return nil, err
	}
	return client, nil
}
