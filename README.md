# Infra-Lab cli

This application technically is a wrapper around some utils to add some additional logic and to ease local infra in k8s management

It was mainly created as a replacement for make files in [k8s-lab](https://github.com/yura-shutkin/k8s-lab) repo

## Tools used by this util

- VMs for containers
  - [Podman](https://podman.io/)

- K8S cluster management
  - [Minikube](https://minikube.sigs.k8s.io/)
  - [Kind](https://kind.sigs.k8s.io/)

- Utils
  - [Helm](https://helm.sh/)
  - [Skopeo](https://github.com/containers/skopeo)
  - [Webhook](https://github.com/adnanh/webhook/)

## Config

You can set configs with both env variables or via config `~/.infra-lab.yaml`

For example `ILC__APPS__MINIKUBE__KUBE_VERSION=v1.33.1 infra-lab-cli c minikube c` will override default version and set the value from env variable
