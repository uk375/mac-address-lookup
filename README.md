### Mac Address Lookup (MAL)
A simple CLI that retrieves the company name when a valid MAC address is provided. This implementation of the solution does not use a provided API Key, but asks the user to provide one during their build of the binary.

#### Prerequisites
The CLI uses https://macaddress.io/ API internally, so an API Key is necessary to use the tool.

Get an API Key by signing up at https://macaddress.io/signup.

GoLang should be installed on the machine or use the "simplified build" as shown below

#### Build
```
## To build
make build API_KEY=<api_key>

## Generate docker image
make docker-image
```

#### Simplified build
If you don't have golang installed on your machine, use
```
make docker-simple API_KEY=<api_key>
```

#### Usage
```shell script
## use the binary (from the project's root)
sudo mv build/_output/bin/mal /usr/local/bin
mal --mac-address <mac_address>

## use docker image
docker run --rm cisco/mal:0.0.1 --mac-address <mac_address>
```