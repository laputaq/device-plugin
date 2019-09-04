# VMware device plugin for Kubernetes

## About

The VMware device plugin for Kubernetes is a Daemonset that allows you to automatically:
- Expose the vGPU capacity on each node of your cluster
- Keep track of the health of your vGPUs
- Run vGPU enabled containers in your Kubernetes cluster.

This following work builds on NVIDIA's official implementation of the [NVIDIA device plugin](https://github.com/NVIDIA/k8s-device-plugin) for Kubernetes, so the usage and documentation are roughly the same.

## Prerequisites

The list of prerequisites for running the Vmware device plugin is described below:
* NVIDIA drivers ~= 361.93
* Docker >= 19.03 (see [why](https://github.com/NVIDIA/nvidia-docker#quickstart))
* nvidia-docker >= 2.1.0 (see how to [install](https://github.com/NVIDIA/nvidia-docker#quickstart) and it's [prerequisites](https://github.com/nvidia/nvidia-docker/wiki/Installation-(Native-GPU-Support)#prerequisites))
* nvidia-container-runtime >= 3.1.0 (see how to [install](https://github.com/NVIDIA/nvidia-container-runtime#installation))
* docker configured with nvidia as the [default runtime](https://github.com/NVIDIA/nvidia-docker/wiki/Advanced-topics#default-runtime).
* VMware vGPU Scheduler >= 1.0 (see how to [install](https://github.com/laputaq/vgpu-scheduler#quick-start) and it's [rerequisites](https://github.com/laputaq/vgpu-scheduler#prerequisites))
* Kubernetes version >= 1.15

## Quick Start

The following steps need to be executed on all your vGPU nodes.
This README assumes that the NVIDIA drivers and nvidia-docker have been installed.

### Config nvidia-container-runtime

After nvidia-container-runtime is installed. You will need to enable the nvidia runtime as your default runtime on your node.
We will be editing the docker daemon config file which is usually present at `/etc/docker/daemon.json`:
```json
{
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}
```
> *if `runtimes` is not already present, head to the install page of [nvidia-docker](https://github.com/NVIDIA/nvidia-docker)*

### Enabling vGPU Support in Kubernetes

Once you have enabled this option on *all* the vGPU nodes you wish to use,
you can then enable vGPU support in your cluster by deploying the following Daemonset:

```shell
$ kubectl create -f https://raw.githubusercontent.com/laputaq/device-plugin/master/deployment/deployment.yaml
```

### Enabling vGPU Scheduler in Kubernetes

Since the vGPU resources on a Node can only be assigned to a Pod, the default Scheduler of Kubernetes can no longer meet our needs because it may deploy multiple Pods to one same Node. So we need to enable vGPU Scheduler in your cluster by deploying the vGPU Scheduler according to following [step](https://github.com/laputaq/vgpu-scheduler).


### Running vGPU Jobs

VMware vGPUs can now be consumed via container level resource requirements using the resource name vmware.com/vgpu by deploying the following Pod:
``` shell
$ kubectl create -f https://raw.githubusercontent.com/laputaq/device-plugin/master/deployment/vgpu-demo.yaml
```

> **WARNING:** *if you don't request vGPUs when using the VMware device plugin  all the vGPUs on the machine will be exposed inside your container.*

## Developing

### With Docker

#### Build
Option 1, pull the prebuilt image from [Docker Hub](https://hub.docker.com/r/laputaq/device-plugin):
```shell
$ docker pull laputaq/device-plugin:v1.0
```

Option 2, build without cloning the repository:
```shell
$ docker build -t laputaq/device-plugin:v1.0 https://github.com/laputaq/device-plugin.git
```

Option 3, if you want to modify the code:
```shell
$ git clone https://github.com/laputaq/device-plugin.git && cd device-plugin
$ docker build -t laputaq/device-plugin:v1.0 .
```

#### Run locally
```shell
$ docker run --security-opt=no-new-privileges --cap-drop=ALL --network=none -it -v /var/lib/kubelet/device-plugins:/var/lib/kubelet/device-plugins laputaq/device-plugin:v1.0
```

#### Deploy as DaemonSet:
```shell
$ cd deployment && kubectl create -f device-plugin.yaml
```

### Without Docker

#### Build
```shell
$ C_INCLUDE_PATH=/usr/local/cuda/include LIBRARY_PATH=/usr/local/cuda/lib64 go build
```

#### Run locally
```shell
$ ./device-plugin
```