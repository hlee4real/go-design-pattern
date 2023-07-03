package main

import "errors"

//Abstract Factory là một mẫu thiết kế creational cho phép bạn tạo ra các họ đối tượng
//có liên quan mà không chỉ định các lớp cụ thể của chúng.

//mình sẽ tạo ra các đối tượng cụ thể (Premier League, Champions League) từ các đối tượng trừu tượng (InterfaceFootballFactory)
//mình cũng có các đối tượng trừu tượng như InterfaceBall và InterfaceShirt

type InterfaceFootballFactory interface {
	CreateBall() InterfaceBall
	CreateShirt() InterfaceShirt
}

func GetFootballFactory(factoryName string) (InterfaceFootballFactory, error) {
	switch factoryName {
	case "Premier League":
		return PremierLeague{}, nil
	case "Champions League":
		return ChampionsLeague{}, nil
	default:
		return nil, errors.New("Invalid football factory")
	}
}
