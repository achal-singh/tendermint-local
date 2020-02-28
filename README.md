# tendermint-local

Tendermint Core setup and Tendermint Go-Client implementation in a local environment.

# Tendermint Core Setup (Local Environment)

The easiest way to setup tendermint core is via **docker-compose**. In the following steps a **Cluster of 4 nodes** will be setup to run a local testnetwork of tendermint core. Note: The following solution has been tested on **Ubuntu 16.04 LTS**.
It is assumed that you already have the latest version of [Go](https://golang.org/doc/install#install) installed in your machine and you're aware of the `$GOPATH`.

1. Firstly, install the go dependency of [Tendermint Core](https://github.com/tendermint/tendermint) which should be stored at `$GOPATH/src/github.com/`.

```
go get -v github.com/tendermint/tendermint
```

Ignore any warnings like: `package github.com/tendermint/tendermint: no Go files in /Users/Achal/go/src/github.com/tendermint/tendermint`.

2. Get into the tendermint root directory:

```
cd $GOPATH/src/github.com/tendermint/tendermint
```

3. Build the linux binary

```
make build-linux
```

## Running the Testnetwork

To launch a 4 node testnetwork, execute:

```
make localnet-start
```

A total of two major things take place in this step:

1. A local image of the official tendermint [Docker Image](https://hub.docker.com/r/tendermint/localnode) is built locally.
2. The 4 nodes are launched by a docker-compose script.

The nodes bind their RPC servers to ports 26657, 26660, 26662, and 26664 on the host. The nodes of the network expose their P2P and RPC endpoints to the host machine on ports 26656-26657, 26659-26660, 26661-26662, and 26663-26664 respectively.

The command above also creates a `.tendermint` directory at your `$HOME` location, which should look similar to this:

```
$HOME/.tendermint
├── config
    ├── config.toml     #The Configuration File
├── data
```

To know more about Tendermint Core [Configuration](https://github.com/tendermint/tendermint/blob/master/docs/tendermint-core/configuration.md) file.

In order to stop and restart the Tendermint Core testnetwork, use:

```
make localnet-stop
make localnet-start
```

respectively.

# Tendermint Client

The Tendermint Client is a basic client implementation done in Go to majorly perform 2 simple RPC operations on the local Tendermint Core testnetworkwe just setup in the section above.
This github repo can be cloned and the client can be built and run.
Before building the code and starting the server, create a `config.json` file in the root of the repo. For easy reference `config.example` can be used.

```
go build cmd/main.go
```

On successful build it should return an executable file named **main** in the parent folder.
To start the server:

```
./main
```

This API server responds to only 2 API Calls:

```
/sendTx     # To publish/broadcast a transaction to the testnetwork (key-value pair transaction)
/queryKey   # To query the existence of a key (created via **/sendTx**)
```

To get the correct idea on how to structure the API Calls to the client follow the examples given in the Postman Collection (referenced below).

### Postman Collection (Tendermint Client Testing)

Once the Client server is up and running and you see a message like **Tendermint Client Active, Running @ 8080** (Port number can be changed in the config.json file), you can test whether it is interacting with the local testnetwork using [Postman](https://www.postman.com/downloads/) API Testing Tool.
For ease of testing, a [Postman Collection](https://www.getpostman.com/collections/d5ff6f11135ce0f3d274) is there for you to directly import into your machine's Postman App.
