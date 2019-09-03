package main

import pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"

const (
	resourceName           = "vmware.com/vgpu"
	serverSock             = pluginapi.DevicePluginPath + "vmware.sock"
	envResource            = "NVIDIA_VISIBLE_DEVICES"
	envDisableHealthChecks = "DP_DISABLE_HEALTHCHECKS"
	allHealthChecks        = "xids"
)
