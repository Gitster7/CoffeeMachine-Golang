# CoffeeMachine Helm Chart

## Prerequisites

For local setup:
docker, helm, docker-desktop k8 cluster/minikube
For Ingress: NGINX Ingress Controller
FOR PVC: PV already setup

For External K8: docker, helm

## Installing the Chart

### Check possible values

```
#inside chart
helm show values .
#in HELM root folder
helm show values CoffeeMachineChart
```

To install the chart with the release name `my-release`:

```console
helm install my-release .
or
helm install my-release CoffeeMachineChart
or with namespace my-namespace
helm install my-release CoffeeMachineChart --namespace my-namespace
```

## Upgrading an existing Release to a new version

To upgrade existing release named `my-release` use this command

```console
helm upgrade my-release .
or
helm upgrade my-release CoffeeMachineChart
or with namespace my-namespace
helm upgrade my-release CoffeeMachineChart --namespace my-namespace
```

## Uninstalling the Chart

To uninstall/delete the my-release deployment:

```console
helm delete my-release
or with namespace my-namespace
helm delete my-release --namespace my-namespace
```

## Checking chart

To check did helm installed chart there is command:
helm list

To check pods:
kubectl get pods
kubectl describe <podname>

The command removes all the Kubernetes components associated with the chart and deletes the release.


## Configuration

| Parameter                                 | Description                                   | Default                                                 |
|-------------------------------------------|-----------------------------------------------|---------------------------------------------------------|
| `replicaCount`                            | Number of nodes                               | `1`                                                     |
| `image.repository`                        | Image repository                              | `github.com/kostovic/coffeemachine`                     |
| `image.tag`                               | Image tag                                     | `restapiv2.0`                                           |
| `image.pullPolicy`                        | Image pull policy                             | `IfNotPresent`                                          |
| `image.pullSecrets`                       | Image pull secrets                            | `[]`                                                    |
| `securityContext`                         | securityContext of contaier                   | `{}`                                                    |
| `deployment.env`                          | Extra environment variables passed to container| `{GIN_MODE: "release"}`                                |
| `service.type`                            | Kubernetes service type                       | `LoadBalancer`                                          |
| `service.http_api_port`                   | Kubernetes port where service is exposed      | `3000`                                                  |
| `ingress.enabled`                         | Enables Ingress - works only if there is nginx ingress controller installed onto cluster    | `false`   |
| `ingress.annotations`                     | Ingress annotations (values are templated)    | `{}`                                                    |
| `ingress.path`                            | Ingress accepted path                         | `/`                                                     |
| `ingress.hosts`                           | Ingress accepted hostnames                    | `["chart-example.local"]`                               |
| `ingress.tls`                             | Ingress TLS configuration                     | `[]`                                                    |
| `resources`                               | CPU/Memory resource requests/limits           | `limits: cpu: 50m, memory: 128Mi; requests: cpu: 10m, memory: 64Mi`                                       |
| `nodeSelector`                            | Node labels for pod assignment                | `{}`                                                    |
| `volumes.useEmptyDir`                     | Use EmptyDir                                  | `false`                                                 |
| `volumes.emptyDirPath`                    | Path to mount EmptyDir to                     | `false`                                                 |
| `volumes.usePVC`                          | Use PVC                                       | `false`                                                 |
| `volumes.PVName`                          | Persistent Volume Name from k8                | `""`                                                    |
| `volumes.mountPVCPath`                    | Path to mount PVC to container                | `""`                                                    |
| `volumes.subPVCPath`                      | Path from PVC to mount to container           | `""`                                                    |



### Example of volumes usage

EmptyDir shares a temporary directory that shares a pod's lifetime: saves data if restart occurs and can be used for log rotation folder or H2 database storage for testing

Persistent Volume(PV) is independant of pod and needs to be created First. Persistent Volume Claim(PVC) is request for part for PV.
PVC is stored used for holding inMemory database like H2 if you want to have it for later. Put Persisten Volume name to PVName variable.
```yaml
volumes:
  useEmptyDir: true
  emptyDirPath: "/absoulte/path/container/tmp/logs"
  usePVC: true
  PVName: "myvolume"
  mountPVCPath: "/absoulte/path/container/db"
  subPVCPath: "relative/path/in/pvc"
```



