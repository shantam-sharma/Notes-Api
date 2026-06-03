# Terraform Infrastructure as Code Example

## Overview

This directory contains a Terraform configuration demonstrating Infrastructure as Code (IaC) concepts using the Docker provider.

The configuration provisions and manages a Docker-based Nginx container entirely through Terraform.

The purpose of this implementation is to demonstrate:

* Terraform Providers
* Terraform Resources
* Terraform Variables
* Terraform Outputs
* Terraform State Management
* Infrastructure Lifecycle Management

---

## Project Structure

```text
terraform/
├── main.tf
├── variables.tf
├── outputs.tf
└── README.md
```

---

## Resources Created

Terraform provisions the following resources:

### Docker Image

```text
nginx:latest
```

### Docker Container

```text
terraform-nginx
```

### Port Mapping

```text
localhost:8081 → container:80
```

---

## Variables

The configuration uses variables to avoid hardcoded values.

| Variable       | Description                    | Default Value   |
| -------------- | ------------------------------ | --------------- |
| image_name     | Docker image to deploy         | nginx:latest    |
| container_name | Container name                 | terraform-nginx |
| external_port  | Host port exposed by container | 8081            |

---

## Outputs

The following outputs are generated after deployment:

| Output         | Description                    |
| -------------- | ------------------------------ |
| container_name | Name of the deployed container |
| container_id   | Docker container ID            |
| container_url  | URL used to access the service |

Example:

```text
container_name = "terraform-nginx"
container_url = "http://localhost:8081"
```

---

## Terraform Workflow

### Initialize Terraform

```bash
terraform init
```

Downloads required providers and initializes the working directory.

---

### Review Execution Plan

```bash
terraform plan
```

Displays infrastructure changes before applying them.

---

### Create Infrastructure

```bash
terraform apply
```

Creates the Docker image and container.

---

### View Outputs

```bash
terraform output
```

Displays information about deployed resources.

---

### Destroy Infrastructure

```bash
terraform destroy
```

Removes all infrastructure managed by Terraform.

---

## Infrastructure Diagram

```text
Terraform
    │
    ▼
Docker Provider
    │
    ▼
Docker Image (nginx)
    │
    ▼
Docker Container
    │
    ▼
localhost:8081
```

---

## State Management

Terraform stores infrastructure information inside:

```text
terraform.tfstate
```

The state file allows Terraform to:

* Track managed resources
* Detect infrastructure changes
* Determine required updates
* Safely destroy resources

---

## Learning Outcomes

This project demonstrates:

* Infrastructure as Code (IaC)
* Terraform Configuration Language (HCL)
* Docker Provider Integration
* Resource Provisioning
* Variable Management
* Output Management
* State Management
* Infrastructure Lifecycle Operations
