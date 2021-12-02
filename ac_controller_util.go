package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ACControlSC struct {
	contractapi.Contract
}

type ACData struct {
	ID string `json:"ID"`
	ACOutputTemperature float32 `json:"acOutputTemperature"`
}

// GetACTemperature