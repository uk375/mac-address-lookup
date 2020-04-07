package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uk375/mac-address-lookup/config"
	"github.com/uk375/mac-address-lookup/domain"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type Interface interface {
	LookupAddress(macAddr string) (string, error)
}

type macLookupClient struct {
	logger *zap.Logger
	cfg    *config.CommandConfig
}

func NewMacLookupClient(cfg *config.CommandConfig, logger *zap.Logger) Interface {
	return &macLookupClient{
		logger: logger,
		cfg:    cfg,
	}
}

func (mlc *macLookupClient) LookupAddress(macAddr string) (string, error) {
	var companyName string
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?output=json&search=%s", mlc.cfg.MacAddressHostName, macAddr), nil)
	if err != nil {
		mlc.logger.Error("failed to create a request to lookup mac address", zap.Error(err))
		return companyName, err
	}
	req.Header.Add("x-authentication-token", mlc.cfg.APIKey)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		mlc.logger.Error("failed to lookup the mac address", zap.Error(err))
		return companyName, err
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			mlc.logger.Error("the provided api token in invalid")
			return companyName, errors.New("the configured api token is not valid")
		}
		return companyName, errors.New("unable to retrieve the company name at this time")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var lookupRes domain.LookupResponse
	err = json.Unmarshal(body, &lookupRes)
	if err != nil {
		mlc.logger.Error("failed to parse the lookup response", zap.Error(err))
		return companyName, err
	}

	return lookupRes.VendorDetails.CompanyName, nil
}
