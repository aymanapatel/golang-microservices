
# Feature 1: HTTP Routing and LBing using Envoy

# Table of Contents

- [Links](#links)



# Links
1. [Envoy Docs](https://www.envoyproxy.io/docs/envoy/latest/)
2. [Envoy Homepage](https://www.envoyproxy.io/)



# Envoy Architecture




## Inbound 

### Listener
### Filterchain

### Filter

- TCP Proxy
- HTTP Connection manager


## Outbound

## Clusters

## Endpoints

- Static cluster (load_assignment): Example: Config. in YAML/JSON. Better not to use this.
- Dynamic Cluster ([eds](#eds) _assignment): Example: Istio/Consul connect. 

![Arch.](/pictures/Envoy1.png)

## Upstream clusters

### EDS
Endpoint Discovery service
