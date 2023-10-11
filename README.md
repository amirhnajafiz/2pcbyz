# Node Exporter

![GitHub top language](https://img.shields.io/github/languages/top/amirhnajafiz/node-exporter)
![GitHub release (with filter)](https://img.shields.io/github/v/release/amirhnajafiz/node-exporter)

Creating a Prometheus exporter for exporting ```Kubernetes``` nodes metrics. By using this
exporter you can get resource metrics of your nodes that are being used in ```k8s``` cluster.
These metrics will be available on ```/metrics``` endpoint of a http server.

## setup

Build and deploy exporter dockerfile.

```shell
docker build . -f build/Dockerfile -t node-exporter:v0.1.0
docker push node-exporter:v0.1.0 # push into your image repository
```

Now execute docker container, and get metrics on ```ip:80/metrics```.

```shell
docker run -d node-exporter:v0.1.0
```
