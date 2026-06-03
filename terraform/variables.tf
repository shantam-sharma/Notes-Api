variable "container_name" {
    description = "name of docker container"
    type = string
    default = "terraform-nginx"
}

variable "image_name" {
    description = "docker image name"
    type = string
    default = "nginx:latest"
}

variable "external_port" {
    description = "External port"
    type = number
    default = 8081
}
