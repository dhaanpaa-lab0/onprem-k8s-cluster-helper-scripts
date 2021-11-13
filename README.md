# Container Server Helper Scripts

Note: ***Should Be able to use with podman also***


## Script Prefixes
| Prefix | Description |
|--------|-------------|
| cnt | Both Docker & Podman |
| pdm | Podman Only |
| dkr | Docker Only |
| spt | Support Scripts |

## Scripts
| Script Name | Description |
|-------------|-------------|
| dkr-add-user-to-docker-group | Add User for usage of docker without sudo |
| cnt-build-dockerfile | Build Docker image for docker file |
| cnt-create-dockerfile | Create New DockerFile with git support |
| cnt-create-generic-volume | Create Volume |
| cnt-create-local-volume | Create Volume from local directory |
| cnt-create-mongo-server | Create MongoDB Server (Percona) |
| cnt-create-mysql-server | Create Mysql Server |
| cnt-create-oracle-xe-server | Create Oracle XE Server |
| cnt-create-postgres-server | Create Postgres Server |
| cnt-create-sql-server | Create Microsoft SQL Server |
| cnt-delete-container | Delete Container |
| cnt-delete-image | Delete Image |
| cnt-delete-volume | Delete Local Docker Volume |
| cnt-list-containers | List Containers |
| cnt-list-images | List Images |
| cnt-list-volumes | List Volumes |
| cnt-registry-installer | Craete Docker Registry on server |
| cnt-run-bash-in-container | Start a bash shell in container |
| cnt-run-sh-in-container | Start generic shell in container |
| cnt-run-zsh-in-container | Start ZSH shell in container |
| cnt-start-container | Start a contaienr that was stopped |
| cnt-start-image | Start a new container |
| cnt-start-image-with-shell | Start a new container and then drop to shell |
| cnt-stop-container | Stop a container |
| pdm-create-docker-link | Create new docker link for a podman host |
| spt-install | Install Scripts System Wide |
| spt-install-user | Install Scripts Just for current user |