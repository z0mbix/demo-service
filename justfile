set shell := ["bash", "-uc"]

# Show available targets/recipes
default:
    @just --choose

# Clean up old files
clean:
    rm demo-service

# Build the container image and push to k3d
build version:
    docker build -t demo-service:{{version}} .
    k3d image import --cluster demo demo-service:{{version}}

# Create k3d cluster
create-cluster:
    k3d cluster create demo

# Delete k3d cluster
delete-cluster:
    k3d cluster delete demo

# Deploy application to kubernetes cluster
deploy:
    kubectl apply -f ./kubernetes/

# Show help menu
help:
    @just --list --list-prefix '  ❯ '
