# Connect Sidecar Helm Chart

This Helm chart installs the Connect-K8s service, a Kubernetes connector used to manage and interact with Kubernetes clusters. This chart simplifies deployment and configuration, enabling seamless integration within your existing Kubernetes environment.

### Prerequisites

	•	Kubernetes 1.20+
	•	Helm 3.0+
	•	A configured Kubernetes cluster

## Installation

### Installing the Chart

Add the Helm repository:
```bash
helm repo add connect-k8s https://bit0ps.github.io/connect-k8s/
helm repo update
```

Install the chart with a release name of your choice (e.g., connect):
```bash
helm install connect connect-k8s/contrib/chart
```
This command deploys the chart with default configuration values. For custom settings, refer to the Configuration section.

### Upgrading the Chart
To upgrade your deployment after making updates to the chart or configuration, run:
```bash
helm upgrade connect connect-k8s/contrib/chart
```

### Uninstalling the Chart
To remove the deployment from your cluster:
```bash
helm uninstall connect 
```

## Configuration
The chart includes several customizable parameters for tailoring the deployment to your needs. Specify each parameter using the --set key=value argument in helm install or use helm upgrade commands.

Alternatively, modify the values.yaml file for more complex setups and use the following:
```bash
helm upgrade --install connect connect-k8s/contrib/chart
```
## General parameters

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| fullnameOverride | string | `""` | String to fully override `"chart.fullname"` |
| nameOverride | string | `"connect"` | Provide a name in place of `connect` |
| namespaceOverride | string | `.Release.Namespace` | Override the namespace |

## Global Configs

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.addPrometheusAnnotations | bool | `false` | Add Prometheus scrape annotations to all metrics services. This can be used as an alternative to the ServiceMonitors. |
| global.additionalLabels | object | `{}` | Common labels for the all resources |
| global.deploymentAnnotations | object | `{}` | Annotations for the all deployed Deployments |
| global.deploymentStrategy | object | `{}` | Deployment strategy for the all deployed Deployments |
| global.image.imagePullPolicy | string | `"IfNotPresent"` | If defined, a imagePullPolicy applied to all deployments |
| global.image.repository | string | `"quay.io/skip/connect"` | If defined, a repository applied to all deployments |
| global.image.tag | string | `""` | Overrides the global Connect image tag whose default is the chart appVersion |
| global.imagePullSecrets | list | `[]` | Secrets with credentials to pull images from a private registry |
| global.podAnnotations | object | `{}` | Annotations for the all deployed pods |
| global.podLabels | object | `{}` | Labels for the all deployed pods |
| global.revisionHistoryLimit | int | `3` | Number of old deployment ReplicaSets to retain. The rest will be garbage collected. |
| global.securityContext | object | `{}` (See [values.yaml]) | Toggle and define pod-level security context. |
| global.statefulsetAnnotations | object | `{}` | Annotations for the all deployed Statefulsets |



## Connect Configs
| Key                                         | Type    | Default                   | Description                                |
|---------------------------------------------|---------|---------------------------|--------------------------------------------|
| `connect.name`                              | string  | `sidecar`                 | Connect oracle name                        |
| `connect.replicaCount`                      | int     | `1`                       | Replica count                              |
| `connect.automountServiceAccountToken`      | bool    | `true`                    | Automount API credentials                  |
| `connect.image.repository`                  | string  | `connect`                 | Image repository                           |
| `connect.image.pullPolicy`                  | string  | `Never`                   | Image pull policy                          |
| `connect.image.tag`                         | string  | `"2.1.0"`                 | Image tag                                  |
| `connect.imagePullSecrets`                  | list    | `[]`                      | Secretes for pulling an image from a private repository |
| `connect.extraEnvVars.normal`               | object  | `{}`                      | Non-sensitive environment variables        |
| `connect.extraEnvVars.secret`               | object  | `{}`                      | Sensitive environment variables            |
| `connect.settings.marketMapProvider`        | string  | `dydx_api`                | MarketMap provider to use                  |
| `connect.settings.oracleConfig`             | string  | `/etc/config/oracle.json` | Oracle config file location                |
| `connect.settings.runPprof`                 | bool    | `true`                    | Enable pprof server                        |
| `connect.settings.pprofPort`                | int     | `6060`                    | Pprof server port                          |
| `connect.settings.logStdOutLevel`           | string  | `info`                    | Log level for stdout                       |
| `connect.settings.logFileLevel`             | string  | `info`                    | Log level for file                         |
| `connect.settings.logFile`                  | string  | `sidecar.log`             | Log file name                              |
| `connect.settings.logMaxSize`               | int     | `100`                     | Max log size (MB)                          |
| `connect.settings.logMaxBackups`            | int     | `1`                       | Max log backups                            |
| `connect.settings.logMaxAge`                | int     | `3`                       | Log retention (days)                       |
| `connect.settings.metricsEnabled`           | bool    | `true`                    | Enable Oracle client metrics               |
| `connect.settings.metricsPrometheusAddress` | string  | `0.0.0.0:8002`            | Prometheus server address                  |
| `connect.settings.mode`                     | string  | `exec`                    | Oracle run mode                            |
| `connect.settings.host`                     | string  | `0.0.0.0`                 | Oracle host address                        |
| `connect.settings.port`                     | int     | `8080`                    | Oracle port                                |
| `connect.serviceAccount.create`             | bool    | `true`                    | Create service account                     |
| `connect.serviceAccount.automountServiceAccountToken` | bool | `true`          | Automount service account token            |
| `connect.serviceAccount.annotations`        | object  | `{}`                      | Annotations for the service account        |
| `connect.serviceAccount.labels`             | object  | `{}`                      | Labels for the created service account     |
| `connect.serviceAccount.name`               | string  | `""`                      | Name of the service account (auto-generated if empty) |
| `connect.deploymentStrategy`                | object  | `{}`                      | Deployment strategy for Connect app        |
| `connect.deploymentAnnotations`             | object  | `{}`                      | Annotations for Connect app Deployment     |
| `connect.podAnnotations`                    | object  | `{}`                      | Kubernetes annotations for pod             |
| `connect.podLabels`                         | object  | `{}`                      | Kubernetes labels for pod                  |
| `connect.podSecurityContext`                | object  | `{}`                      | Security context for pod                   |
| `connect.securityContext`                   | object  | `{}`                      | Container security context                 |
| `connect.service.annotations`               | object  | `{}`                      | Service annotations                        |
| `connect.service.labels`                    | object  | `{}`                      | Service labels                             |
| `connect.service.type`                      | string  | `ClusterIP`               | Service type                               |
| `connect.service.nodePortHttp`              | int     | `30080`                   | HTTP port for NodePort service type        |
| `connect.service.portHttpName`              | string  | `http`                    | HTTP service port name                     |
| `connect.service.http`                      | int     | `8080`                    | HTTP service port                          |
| `connect.service.portPprofName`             | string  | `pprof`                   | Pprof service port name                    |
| `connect.service.pprof`                     | int     | `6060`                    | Pprof service port                         |
| `connect.service.loadBalancerClass`         | string  | `""`                      | Load balancer class                        |
| `connect.service.loadBalancerIP`            | string  | `""`                      | Load balancer IP                           |
| `connect.service.loadBalancerSourceRanges`  | list    | `[]`                      | Source IP ranges for load balancer         |
| `connect.service.externalIPs`               | list    | `[]`                      | Service external IPs                       |
| `connect.service.externalTrafficPolicy`     | string  | `Cluster`                 | External traffic routing policy            |
| `connect.service.sessionAffinity`           | string  | `None`                    | Session affinity setting                   |
| `connect.metrics.service.type`              | string  | `ClusterIP`               | Metrics service type                       |
| `connect.metrics.service.clusterIP`         | string  | `""`                      | ClusterIP for metrics service (`None` for headless) |
| `connect.metrics.service.annotations`       | object  | `{}`                      | Annotations for metrics service            |
| `connect.metrics.service.labels`            | object  | `{}`                      | Labels for metrics service                 |
| `connect.metrics.service.portName`          | string  | `http-metrics`            | Port name for metrics service              |
| `connect.ingress.enabled`                   | bool    | `false`                   | Enable ingress for Connect app             |
| `connect.ingress.annotations`               | object  | `{}`                      | Additional ingress annotations             |
| `connect.ingress.labels`                    | object  | `{}`                      | Additional ingress labels                  |
| `connect.ingress.ingressClassName`          | string  | `""`                      | Ingress controller class name              |
| `connect.ingress.hostname`                  | string  | `""`                      | Hostname for Connect app                   |
| `connect.ingress.path`                      | string  | `/`                       | Ingress path for Connect app               |
| `connect.ingress.pathType`                  | string  | `Prefix`                  | Path type (`Exact`, `Prefix`, `ImplementationSpecific`) |
| `connect.ingress.tls`                       | bool    | `false`                   | Enable TLS for ingress                     |
| `connect.ingress.extraHosts`                | list    | `[]`                      | Additional hostnames                       |
| `connect.ingress.extraPaths`                | list    | `[]`                      | Additional ingress paths                   |
| `connect.ingress.extraRules`                | list    | `[]`                      | Additional ingress rules                   |
| `connect.ingress.extraTls`                  | list    | `[]`                      | Additional TLS configuration               |
| `connect.resources.limits.cpu`              | string  | `1`                       | CPU limit                                  |
| `connect.resources.limits.memory`           | string  | `1Gi`                     | Memory limit                               |
| `connect.resources.requests.cpu`            | string  | `100m`                    | CPU request                                |
| `connect.resources.requests.memory`         | string  | `128Mi`                   | Memory request                             |
| `connect.livenessProbe.enabled`             | bool    | `true`                    | Enable liveness probe                      |
| `connect.livenessProbe.path`                | string  | `"/connect/oracle/v2/prices"` | Liveness probe path                        |
| `connect.livenessProbe.initialDelaySeconds` | int     | `15`                      | Liveness probe initial delay (seconds)     |
| `connect.livenessProbe.periodSeconds`       | int     | `6`                       | Liveness probe interval (seconds)          |
| `connect.livenessProbe.timeoutSeconds`      | int     | `2`                       | Liveness probe timeout (seconds)           |
| `connect.livenessProbe.failureThreshold`    | int     | `3`                       | Liveness probe failure threshold           |
| `connect.livenessProbe.successThreshold`    | int     | `1`                       | Liveness probe success threshold           |
| `connect.readinessProbe.enabled`            | bool    | `true`                    | Enable readiness probe                     |
| `connect.readinessProbe.path`               | string  | `"/connect/oracle/v2/prices"` | Readiness probe path                       |
| `connect.readinessProbe.initialDelaySeconds`| int     | `15`                      | Readiness probe initial delay (seconds)    |
| `connect.readinessProbe.periodSeconds`      | int     | `6`                       | Readiness probe interval (seconds)         |
| `connect.readinessProbe.timeoutSeconds`     | int     | `2`                       | Readiness probe timeout (seconds)          |
| `connect.readinessProbe.failureThreshold`   | int     | `3`                       | Readiness probe failure threshold          |
| `connect.readinessProbe.successThreshold`   | int     | `1`                       | Readiness probe success threshold          |
| `connect.autoscaling.enabled`               | bool    | `false`                   | Enable autoscaling                         |
| `connect.autoscaling.minReplicas`           | int     | `1`                       | Minimum number of replicas                 |
| `connect.autoscaling.maxReplicas`           | int     | `100`                     | Maximum number of replicas                 |
| `connect.autoscaling.targetCPUUtilizationPercentage` | int | `80`           | Target CPU utilization percentage          |
| `connect.autoscaling.targetMemoryUtilizationPercentage` | int | `80`        | Target memory utilization percentage       |
| `connect.volumes`                           | list    | `[]`                      | Additional volumes for the deployment      |
| `connect.volumeMounts`                      | list    | `[]`                      | Additional volume mounts for the deployment|
| `connect.emptyDir.sizeLimit`                | string  | `"1Gi"`                   | EmptyDir size limit                        |
| `connect.nodeSelector`                      | object  | `{}`                      | Node selector for pod assignment           |
| `connect.tolerations`                       | list    | `[]`                      | Tolerations for pod assignment             |
| `connect.affinity`                          | object  | `{}`                      | Affinity settings for pod assignment       |