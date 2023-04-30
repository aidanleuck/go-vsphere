package common

import "vsphere_module/src/queue"

type AgentService struct {
	Config *AppConfiguration
	Queue  queue.Queue
}
