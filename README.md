# On-Premise kubernetes cluster Helper tools

## Script/Tool Prefixes
| Prefix | Description |
|--------|-------------|
| pdm | Podman Only |
| dkr | Docker Only |
| spt | Support Scripts |
| k0s | k0s cluster only |
| k3s | k3s cluster only |
| k8s | all kubernetes cluster types |

## Scripts/Tools
| Script Name | Description |
|-------------|-------------|
| spt-install | Install Scripts System Wide |
| spt-install-user | Install Scripts Just for current user |
| k3s-uninstall | Uninstall K3S (sudo) |
| k8s-install-arkade | Install Arkdate tool (System Wide) (SUDO) |
| k8s-install-docker-secret-from-homedir | Install Docker Registry access secret |
| k8s-install-metallb | Install MetalLB to current kubernetes cluster |
| dkr-install | Install Docker on host |
| pdm-install | Install Podman on host |