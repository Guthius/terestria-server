<div align="center">
    <img src=".github/assets/logo.png" width="400">

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guthius/mirage-nova)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Guthius/mirage-nova/go.yml)
![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/guthius/mirage-nova)
![GitHub License](https://img.shields.io/github/license/Guthius/mirage-nova)

</div>

---

This is the server component of *Terestria*, a 2D game platform tailored for building and hosting small-scale online multiplayer role-playing games.

The server is written in [Golang](https://golang.org/dl/) and utilizes its robust concurrency model to deliver efficient and scalable networking capabilities.

## Build Instructions

### Clone the repository
```bash
git clone https://github.com/guthius/terestria-server.git
cd terestria-server
```

### Install Go

Ensure that [Go](https://golang.org/dl/) is installed on your system. The minimum required version of Go is 1.23.

You can verify the installation by running:
```bash
go version
```

### Build the server

Navigate to the server directory and build the server:
```bash
go build -o ./bin/
```

### Run the server

Execute the built server binary:
```bash
cd ./bin
./terestria-server
```

## License

This project is licensed under the MIT License. For the complete license text, please refer to the [LICENSE](LICENSE) file.