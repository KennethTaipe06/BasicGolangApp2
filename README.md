# Basic Golang app

This is a simple Golang project created as part of a Distributed Programming assignment. For this assignment, we used five different programming languages across five individual projects to demonstrate and explore various technologies and methods.

## Table of Contents

- [Project Description](#project-description)
- [Requirements](#requirements)
- [Installation and Local Setup](#installation-and-local-setup)
- [Usage](#usage)
- [Docker](#docker)

## Project Description

This project involves creating a simple Go web application that displays a basic homepage.

### Features:
- Basic setup of a Go application
- A homepage that displays the message "It works 3/5 golang"

## Requirements

Before you begin, ensure you have the following installed:

- Go 1.23.0 or higher: [Download Go](https://golang.org/dl/)
- Docker (optional, to run the application in a container)

## Installation and Local Setup

Follow these steps to set up the project on your local machine:

1. **Clone the repository** to your local machine:
    ```bash
    git clone https://github.com/KennethTaipe06/BasicGolangApp2.git
    cd BasicGolangApp2
    ```

2. **Download the dependencies**:
    ```bash
    go mod download
    ```
    
## Usage

Once the dependencies are installed, run the Go application with the following command:
```bash
go run main.go
```
Access the application: http://localhost:8080/app/main.go
## Docker
if you want to run a docker image of this proyect you should go:
- [Docker Hub](https://hub.docker.com/repository/docker/byvoxel/golang4/general)
