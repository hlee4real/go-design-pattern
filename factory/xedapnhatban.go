package main

type XeDapNhatBan struct {
	Bike
}

func newXeDapNhatBan() InterfaceVehicle {
	return &XeDapNhatBan{
		Bike: Bike{
			name:  "Xe dap Nhat Ban",
			speed: 40,
		},
	}
}
