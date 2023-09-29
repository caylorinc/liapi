# LiAPI - A Simple Golang REST API Application

## Table of Contents

- [Overview](#overview)
- [Endpoints](#endpoints)
- [Building](#building)
- [Running](#running)
- [Docker](#docker)
- [GitHub Workflow](#github-workflow)
- [Environment Variables](#environment-variables)
- [Logging](#logging)

## Overview

LiAPI is a lightweight REST API written in Go, providing a couple of simple endpoints. This API is designed to showcase how to build, dockerize, and deploy a Golang application using GitHub Actions.

## Endpoints

### 1. `/liatrio`

- **GET Method**: Returns a static message along with the current timestamp.
- **POST Method**: Accepts a JSON object with a message and returns the received message along with the current timestamp.
- **Response Format**: 
  ```json
  {
    "message": "string",
    "timestamp": "string"
  }
  ```

### 2. `/ping`

- **GET Method**: Returns the current health status and version of the application.
- **Response Format**: 
  ```json
  {
    "app": "liapi",
    "version": "string",
    "commitHash": "string",
    "buildTime": "string"
  }
  ```

## Building

To build the application locally, navigate to the project directory and run:

```bash
go build .
```

This command will produce a binary named after the directory, which can be executed directly.

## Running

To run the application, execute the binary created by the build process:

```bash
./<binary-name>
```

By default, the application will be available at `http://localhost:8080`. You can customize the address by setting the `LIAPI_ADDRESS` environment variable.

## Docker

The application is dockerized and the image is pushed to GitHub Container Registry on release. To build the Docker image locally, navigate to the project directory and run:

```bash
docker build -t liapi .
```

To run the Docker container:

```bash
docker run -p 8080:8080 liapi
```

## GitHub Workflow

The application uses a GitHub workflow for Continuous Integration (CI) and Continuous Deployment (CD). Upon each release, the workflow automatically builds the application, packages it into a Docker container, and pushes the image to GitHub Container Registry.

## Environment Variables

- `LIAPI_ADDRESS`: Specifies the address the server binds to. Defaults to `:8080`.

## Logging

The application utilizes structured logging to log essential information, such as incoming requests and server status. The log entries include details like remote address, request method, host, URL, etc.

Example log entry:
```
2023/09/26 20:46:02 INFO liapi version=v0.0.5 commit=47d963bb7f3fcc9d5037f03650c0505b649d6fb3 buildTime=1695417115
2023/09/26 20:46:02 INFO Starting the server address=:8080
2023/09/26 20:46:06 INFO request received address=172.17.0.1:50074 method=GET host=localhost:8080 url=/
2023/09/26 20:46:06 INFO request received address=172.17.0.1:50074 method=GET host=localhost:8080 url=/favicon.ico
2023/09/26 20:46:10 INFO request received address=172.17.0.1:50074 method=GET host=localhost:8080 url=/liatrio
2023/09/26 20:46:33 INFO request received address=172.17.0.1:50074 method=GET host=localhost:8080 url=/ping
```

### Note

This README provides essential documentation for understanding, building, and deploying the LiAPI application. Please refer to the inline comments in the code for more detailed information on specific implementation details.