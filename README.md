

# GoSSO: gRPC-based SSO Microservice written in Go

<p align="center">
    <img style="width: 140px;" src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="logo">
    <img style="margin-left: 10px; width: 120px;" src="https://miro.medium.com/v2/resize:fit:1400/format:webp/1*xZXmBNa-o0P5YYsKmsKO0Q.png" alt="logo">
</p>

Welcome to GoSSO, a highly efficient and scalable Single Sign-On (SSO) microservice implemented in Go, leveraging the gRPC protocol. This project is designed to provide a robust authentication solution that can be easily integrated into any system requiring secure and streamlined user access control.

**Note**: The project is created exclusively in educational purposes so will not be maintained properly. I don't plan to make further updates.

## Run Locally

Clone the project

```bash
  git clone https://github.com/ilborsch/go-grpc-auth
```

Go to the project directory

```bash
  cd go-grpc-auth
```

Install Go dependencies

```bash
  go get
```

Start the server

```bash
   go run ./cmd/sso/main.go --config=./config/local.yaml
```




## Authors

- [@ilborsch](https://www.github.com/ilborsch)


## License

[MIT](https://choosealicense.com/licenses/mit/)
