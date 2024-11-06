# poc-api-architectures

Get Started
-----------
#### Requirements

To run this application on your machine, you need at least:

- go 1.23
- stubby (HTTP mock server)

Install Stubby globally with npm:
---------------------
```
npm install -g stubby
```

Sync lib
---------------------

```
make tidy
```

Running Application
------------------------------------
Start Project

```bash
make run
```

Running Stubby
------------------------------------
```bash
make run-stubby
```

Application Pattern:
---------------------
```
.
├── cmd                        # Application entry point
│   └── main.go 
├── config                     # Application configuration
│   └── config.yaml    
├── internal                   # Golang's standard internal package
│   └── module
│       ├── module.go          # The models' structure and rules for the application.
│       ├── handler.go         # The handlers, to handle the incoming requests.
│       ├── service.go         # The business logic of the application.
│       └── repostiory.go      # The repositories such as Database, a microservice API exposed via gRPC or REST.       
└── pkg                        # utility packages.                

```
