# README

This repo provides a walkthrough for setting up a Kubernetes deployment scaled on custom metrics from using the Horizontal Pod Autoscaler (HPA). There is a good [walkthrough in the Kubernetes docs](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) on using the HPA to scale on Pod CPU, but there are a number of steps that are not covered for using custom metrics. This repo will take you through getting Prometheus set up with a demo custom metric and configuring the HPA to use that.

A typical scenario for this would be to scale a deployment that is processing messages from a queue based on the queue depth. To keep the walkthrough simple we will use a custom metric source instead of a real queue - this lets us avoid have to script adding messages from the queue and processing them, but it should still be clear how to extend this walkthrough to a real scenario.

## Walktroughs

* [Single namespace for Prometheus and application](./single-namespace.md)
* [Separate namespace for Prometheus and application](./separate-namespaces.md)
