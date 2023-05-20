package service

import (
	"vsphere_module/src/common"
	"vsphere_module/src/queue"
	"vsphere_module/src/vimautomation"

	"go.uber.org/zap"
)

type AgentService struct {
	Config        *common.AppConfiguration
	VsphereClient *vimautomation.VSphereClient
	Queue         queue.Queue
	Logger        *zap.SugaredLogger
}
