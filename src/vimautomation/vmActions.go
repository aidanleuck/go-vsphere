package vimautomation

import (
	"context"
	"strings"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/types"
)

func (cl VSphereClient) Clone(dcName string, path string) error {
	dc, err := getDatacenter(cl.Client, dcName)
	if err != nil {
		return err
	}

	vmTemplate, err := getVMTemplate(cl.Client, dc, "/vm/DC0_H0_VM0")
	if err != nil {
		return err
	}
	f := find.NewFinder(cl.Client.Client)
	folder := strings.Split(vmTemplate.InventoryPath, "/")
	parentFolderPath := strings.Join(folder[0:len(folder)-1], "/")
	ff, err := f.Folder(context.TODO(), parentFolderPath)
	if err != nil {
		return err
	}

	_, err = vmTemplate.Clone(context.Background(), ff, "Mems", types.VirtualMachineCloneSpec{PowerOn: true})

	if err != nil {
		return err
	}
	return nil
}

func DeleteVM(vmID string) {

}
