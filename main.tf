terraform {
  required_providers {
    google = {
        source = "hashicorp/google"
        version = "~>6.0"
    }
  }
}

provider "google" {
    project = var.project_id
    region = "asia-northeast1"  
}

variable "project_id" {
    description = "谷歌云项目ID"
    type = string
}

variable "vm_user" {
    description = "SSH登录用户名"
    type = string
}

variable "ssh_public_key" {
    description = "SSH公钥"
    type = string
}


resource "google_compute_instance" "lab_vm" {
    name = "lab-vm"
    machine_type = "e2-standard-2"
    zone = "asia-northeast1-a"
    boot_disk {
      initialize_params {
        image = "ubuntu-os-cloud/ubuntu-2404-lts-amd64"
        size = 30
        type = "pd-ssd"
      }
    }
    network_interface{
        network ="default"
        access_config {
          
        }
    }
    metadata = {
        "ssh-keys" = "${var.vm_user}:${var.ssh_public_key}"
    }

    metadata_startup_script = <<-EOF
    #!/bin/bash
    apt-get update -y
    apt-get install -y docker.io docker-compose-plugin
    systemctl enable docker
    systemctl start docker
    EOF

    tags = ["lab-vm"]

    labels = {
      "purpose" = "observability-lab"
    }
}