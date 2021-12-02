package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

type ACControlSC struct {
	contractapi.Contract
}

type ACData struct {
	ID                  string    `json:"ID"` // ID of AC Unit
	ACOutputTemperature float32   `json:"acOutputTemperature"`
	TimeStamp           time.Time `json:"timeStamp"`
}

// GetACTemperature retrieves current output temperature of an AC unit from the world state
func (acsc *ACControlSC) GetACTemperature(ctx contractapi.TransactionContextInterface, id string) (*ACData, error) {
	acUnitJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if acUnitJSON == nil {
		return nil, fmt.Errorf("AC Unit %s does not exist", id)
	}

	var acUnit ACData
	err = json.Unmarshal(acUnitJSON, &acUnit)
	if err != nil {
		return nil, err
	}

	return &acUnit, nil
}

// 1. Retrieve data from IoT thermal sensors
// 2. Adjust the ACTemperature accordingly (GetACTemperature, AdjustACTemperature)
// 3. Update ACTemperature to world state and send new AC output temperature to client via SDK

// ACUnitExists returns true when AC units with given ID exists in the world state
func (acsc *SmartContract) ACUnitExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	acUnitJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("Failed to read the world state: %v", err)
	}

	return acUnitJSON != nil, nil
}
