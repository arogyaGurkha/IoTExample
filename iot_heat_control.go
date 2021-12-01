package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// This SmartContract provides functions for managing a IoTThermalData
type SmartContract struct {
	contractapi.Contract
}

// IoTThermalData describes basic details of what makes up incoming thermal data from IoT sensors
type IoTThermalData struct {
	ID                    string  `json:"ID"`
	AggregatedTemperature float32 `json:"aggregatedTemperature"`
}

// InitLedger adds a base set of thermal data to the ledger
func (sc *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	thermals := []IoTThermalData{
		{ID: "thermal1", AggregatedTemperature: 34.7},
		{ID: "thermal2", AggregatedTemperature: 32.2},
		{ID: "thermal3", AggregatedTemperature: 36.1},
	}

	for _, thermal := range thermals {
		thermalJSON, err := json.Marshal(thermal)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(thermal.ID, thermalJSON)
		if err != nil {
			return fmt.Errorf("Failed to initial data to world state. %v", err)
		}
	}

	return nil
}

// CreateThermalData issues a new thermal data to the world state with given details
func (sc *SmartContract) CreateThermalData(ctx contractapi.TransactionContextInterface,
	id string, aggregatedTemperature float32) error {

	// Check if thermal data with the corresponding ID already exists
	exists, err := sc.ThermalDataExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Thermal data corresponding to ID %s already exists.", id)
	}

	thermal := IoTThermalData{
		ID: id,
		AggregatedTemperature: aggregatedTemperature,
	}
	thermalJSON, err:= json.Marshal(thermal)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, thermalJSON)
}

// ReadThermalData returns the asset stored in the world state with given id
func (sc *SmartContract) ReadThermalData (ctx contractapi.TransactionContextInterface, id string) (*IoTThermalData, error){
	thermalJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from the world state: %v", err)
	}
	if thermalJSON == nil {
		return nil, fmt.Errorf("The asset %s does not exist", id)
	}

	var thermal IoTThermalData
	err = json.Unmarshal(thermalJSON, &thermal)
	if err != nil {
		return nil, err
	}

	return &thermal, nil
}

// ThermalDataExists returns true when assets with given ID exists in the world state
func (sc *SmartContract) ThermalDataExists (ctx contractapi.TransactionContextInterface, id string) (bool, error){
	thermalJSON, err := ctx.GetStub().GetState(id)
	if err != nil{
		return false, fmt.Errorf("Failed to read the world state: %v", err)
	}

	return thermalJSON != nil, nil
}
