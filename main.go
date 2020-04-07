package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/uk375/mac-address-lookup/config"
	"github.com/uk375/mac-address-lookup/service"
	"go.uber.org/zap"
	"net"
)

var (
	macAddr = flag.String("mac-address", "", "The mac address to lookup information")
)

// Set via `-ldflags` during build, follow semver.org
// Personal API key to access the macaddress.io API, the same personal key will be used make all the calls
// for simplicity. Pass it during the build
var (
	apiKey         string
	macAddressHost = "https://api.macaddress.io/v1"
)

func main() {
	flag.Parse()

	cfg := &config.CommandConfig{}
	cfg.APIKey = apiKey
	cfg.MacAddressHostName = macAddressHost
	logger := configureLogger()

	// Fail if the API key is empty
	if len(cfg.APIKey) == 0 {
		logger.Fatal("An API Key is required to use the CLI. Get a token from -> https://macaddress.io/signup")
	}

	// check if the mac address is empty or invalid
	if len(*macAddr) == 0 {
		fmt.Println("Usage: mal --mac-address <mac-address>")
		logger.Fatal("The provided mac address has invalid format", zap.String("mac_address", *macAddr))
	}

	// using inbuilt golang pkg to validate mac address
	if _, err := net.ParseMAC(*macAddr); err != nil {
		logger.Fatal("The provided mac address is invalid", zap.String("mac_address", *macAddr), zap.Error(err))
	}

	macLookupClient := service.NewMacLookupClient(cfg, logger)
	companyName, err := macLookupClient.LookupAddress(*macAddr)
	if err != nil {
		logger.Fatal("Unable to lookup mac address at this time, please try again later")
	}
	if len(companyName) == 0 {
		logger.Fatal("Unable to find the company name associated with the mac address", zap.String("mac_address", *macAddr))
	}
	fmt.Println(fmt.Sprintf("The mac address: %s is associated with %s company", *macAddr, companyName))
}

func configureLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
