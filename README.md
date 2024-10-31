# Connect

<!-- markdownlint-disable MD013 -->
<!-- markdownlint-disable MD041 -->

[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#wip)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/skip-mev/connect)
[![Go Report Card](https://goreportcard.com/badge/github.com/skip-mev/connect?style=flat-square)](https://goreportcard.com/report/github.com/skip-mev/connect)
[![Version](https://img.shields.io/github/tag/skip-mev/connect.svg?style=flat-square)](https://github.com/skip-mev/connect/releases/latest)
[![Lines Of Code](https://img.shields.io/tokei/lines/github/skip-mev/connect?style=flat-square)](https://github.com/skip-mev/connect)

A general purpose price oracle leveraging ABCI++. Please visit our [docs](https://docs.skip.build/connect/introduction) page for more information!

Connect uses Vote Extensions to create a hyperperformant, extremely secure mechanism for aggregating off-chain data onto a blockchain. It is used by
many of the highest-performance decentralized applications today. If you would like to integrate Connect to power your use case, please contact us on our
[Discord](https://discord.gg/PeBGE9jrbu).

> [!NOTE]
> Connect is **business-licensed software** under BSL, meaning it requires a license to use or reference. It is source viewable, but [**reach out to us on Discord**](https://skip.build/discord) if you are interested in integrating! We are limiting the number of chains we work with to seven in 2024. We apologize if we run out of capacity.

## Install

```shell
$ go install github.com/skip-mev/connect/v2
```

## Overview

The connect repository is composed of the following core packages:

* **abci** - This package contains the [vote extension](./abci/ve/README.md), [proposal](./abci/proposals/README.md), and [preblock handlers](./abci/preblock/oracle/README.md) that are used to broadcast oracle data to the network and to store it in the blockchain.
* **oracle** - This [package](./oracle/) contains the main oracle that aggregates external data sources before broadcasting it to the network. You can reference the provider documentation [here](./providers/base/README.md) to get a high level overview of how the oracle works.
* **providers** - This package contains a collection of [websocket](./providers/websockets/README.md) and [API](./providers/apis/README.md) based data providers that are used by the oracle to collect external data.
* **x/oracle** - This package contains a Cosmos SDK module that allows you to store oracle data on a blockchain.
* **x/marketmap** - This [package](./x/marketmap/README.md) contains  a Cosmos SDK module that allows for market configuration to be stored and updated on a blockchain.

## Validator Usage

To read how to run the oracle as a validator based on the chain, please reference the [validator documentation](https://docs.skip.build/connect/validators/quickstart).

## Developer Usage

To run the oracle, run the following command.

```bash
$ make start-all-dev
```

This will:

1. Start a blockchain with a single validator node. It may take a few minutes to build and reach a point where vote extensions can be submitted.
2. Start the oracle side-car that will aggregate prices from external data providers and broadcast them to the network. To check the current aggregated prices on the side-car, you can run `curl localhost:8080/connect/oracle/v1/prices`.
3. Host a prometheus instance that will scrape metrics from the oracle sidecar. Navigate to http://localhost:9091 to see all network traffic and metrics pertaining to the oracle sidecar. Navigate to http://localhost:8002 to see all application-side oracle metrics.
4. Host a profiler that will allow you to profile the oracle side-car. Navigate to http://localhost:6060 to see the profiler.
5. Host a grafana instance that will allow you to visualize the metrics scraped by prometheus. Navigate to http://localhost:3000 to see the grafana dashboard. The default username and password are `admin` and `admin`, respectively.

After a few minutes, run the following commands to see the prices written to the blockchain:

```bash
# access the blockchain container
$ docker exec -it compose-blockchain-1 bash

# query the price of bitcoin in USD on the node
$ (compose-blockchain-1) ./build/connectd q oracle price BTC USD
```

Result:

```bash
decimals: "8"
id: "0"
nonce: "44"
price:
  block_height: "46"
  block_timestamp: "2024-01-29T01:43:48.735542Z"
  price: "4221100000000"
```

To stop the oracle, run the following command:

```bash
$ make stop-all-dev
```
## Kubernetes Preview Environments

### Prerequisites
Ensure the following tools are installed and accessible in your PATH:

Install minikube</br></br>
MacOS:
```bash
brew install minikube
```
Linux:
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64
```
---
Install helm</br></br>
MacOS:
```bash
brew install helm
```
Linux:
```bash
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```
---
Install kubectl</br></br>
MacOS:
```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```
Linux:
```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```
---
Install docker</br></br>
Follow the [link](https://docs.docker.com/desktop/install/mac-install/) to install Docker on Mac

Linux:
```bash
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

### Makefile Targets
Each target automates specific aspects of the Kubernetes deployment workflow:

1. Dependency and Environment Checks

	•	`check-k8s-deps`: Verifies that all required dependencies (Minikube, kubectl, Helm, Docker) are installed.
	•	`check-env-var`: Checks if the environment variable `PREVIEW_VERSION` is set. If not, the process is halted with an error.

2. Minikube Management

	•	`run-minikube`: Starts Minikube if it’s not already running. Ensures Minikube is operational before proceeding.

3. Building Docker Images

	•	`build-k8s-preview`: Builds Docker images for Kubernetes preview in Minikube’s Docker environment. It uses the `K8S_PREVIEW_DOCKERFILE` and tags images with the provided `SEMVER` version and `PROJECT_NAME`.

4. Helm Chart Deployment

	•	`deploy-k8s-preview`: Deploys the Helm chart to a specified namespace `(K8S_NAMESPACE)`. If the namespace doesn’t exist, it’s created. Tags and versions are set using the `SEMVER` variable.

5. All Components Deployment

	•	`run-all-k8s-preview`: A combined target that checks dependencies, runs Minikube, builds Docker images, and deploys the Helm chart.

6. Cleanup Targets

	•	`delete-all-preview-charts`: Uninstalls the Helm release for the project across all namespaces matching a pattern based on the project name.
	•	`delete-all-preview-ns`: Deletes namespaces matching the pattern.
	•	`minikube-prune`: Runs a docker system prune in Minikube to free up space by removing unused resources.
	•	`clean-all-preview`: A comprehensive cleanup that checks dependencies, uninstalls Helm releases, deletes namespaces, and prunes Docker resources in Minikube.


### Usage
> **_NOTE:_** </br>
`PREVIEW_VERSION` environment variable must be set prior to running `make run-all-k8s-preview` command.

```bash
export PREVIEW_VERSION=<connect version>  # Ex: PREVIEW_VERSION=2.1.0
make run-all-k8s-preview                  # Full preview deployment
make clean-all-preview                    # Full preview cleanup
```

## Metrics

We have an extensive suite of metrics available to validators and chain operators.
 Please [join our discord](https://discord.gg/PeBGE9jrbu) if you want help setting them up!

Metrics relevant to the oracle service's health + operation are detailed in our docs, [here](https://docs.skip.build/connect/metrics/overview)!

