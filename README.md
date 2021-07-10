# atlas-dapr-gateway

This service is attended to route incoming messages from some broker (like Apache Kafka or AWS SNS/SQS) to other destinations. It uses dapr, so the proper dapr components should be provided and used.
You can refer to the Dapr Components documentation here:
https://docs.dapr.io/reference/components-reference/

## Getting Started

You can start the service locally for test purposes, using command:
make kind-deploy
If you need to deploy the service to kubernetes, please provide proper config (see values.yaml, section gatewayConfigurations). Also, please deploy dapr components if needed.

### Prerequisites

Dapr should be deployed in your kubernetes environment, please refer to official documentation:
https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-deploy/

### Installing

This application needs for Helm 3 to be installed on Kubernetes cluster. You can find the examples here: https://www.linode.com/docs/guides/how-to-install-apps-on-kubernetes-with-helm-3/ 

## Deployment

Add additional notes about how to deploy this application. Maybe list some common pitfalls or debugging strategies.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/infobloxopen/atlas-dapr-gateway/tags).
