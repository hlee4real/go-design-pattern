package main

// mediator la mot behavioral design pattern giam su ghep noi giua cac thanh phan cua project
// bang cach lam cho chung giao tiep gian tiep, thong qua mot doi tuong mediator dac biet

//trong truong hop nay, doi tuong mediator la stationManager
// cac doi tuong khac la cac train
// cac train khong giao tiep truc tiep voi nhau, ma thong qua stationManager
// stationManager kiem soat cac train, va quyet dinh khi nao cac train co the di chuyen
// va khi nao cac train phai dung lai
// passenger train va freight train la cac doi tuong khong biet nhau
// nhung chung co the giao tiep voi nhau thong qua stationManager

func main() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
