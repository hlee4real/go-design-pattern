package main

func getBike(bikeType string) (InterfaceVehicle, error) {
	if bikeType == "XeDapNhatBan" {
		return newXeDapNhatBan(), nil
	}
	return nil, nil
}
