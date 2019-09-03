## Prerequisites
- NVIDIA drivers ~= 361.93
- Docker 19.03+
- Kubernetes 1.15+
- Golang 1.12+

## Quick Start
### Enabling GPU Support in Kubernetes
``` bash
curl -O https://raw.githubusercontent.com/laputaq/device-plugin/master/deployment/deployment.yaml
kubectl apply -f deployment.yaml
```

### Running GPU Jobs
``` bash
curl -O https://raw.githubusercontent.com/laputaq/device-plugin/master/deployment/vgpu-1.yaml
kubectl apply -f vgpu-1.yaml
```

## Docs
### With Docker
#### Build
``` bash
docker pull laputaq/device-plugin:v1.0
```

#### Run
``` bash
docker run -d --name device-plugin laputaq/device-plugin:v1.0
```

### Without Docker
#### Build
``` bash
git clone https://github.com/laputaq/device-plugin.git && cd device-plugin
C_INCLUDE_PATH=/usr/local/cuda/include LIBRARY_PATH=/usr/local/cuda/lib64 go build
```

#### Run
``` bash
./device-plugin
```
