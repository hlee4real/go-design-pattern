package main

import "fmt"

//Abstract Factory là một mẫu thiết kế creational cho phép bạn tạo ra các họ đối tượng
//có liên quan mà không chỉ định các lớp cụ thể của chúng.

// mình sẽ tạo ra các đối tượng cụ thể (Premier League, Champions League) từ các đối tượng trừu tượng (InterfaceFootballFactory)
// mình cũng có các đối tượng trừu tượng như InterfaceBall và InterfaceShirt
func main() {
	premierLeagueFactory, err := GetFootballFactory("Premier League")
	if err != nil {
		fmt.Println(err)
		return
	}

	championsLeagueFactory, err := GetFootballFactory("Champions League")
	if err != nil {
		fmt.Println(err)
		return
	}

	premierLeagueBall := premierLeagueFactory.CreateBall()
	premierLeagueShirt := premierLeagueFactory.CreateShirt()

	championsLeagueBall := championsLeagueFactory.CreateBall()
	championsLeagueShirt := championsLeagueFactory.CreateShirt()

	fmt.Println(premierLeagueBall.getBrand())
	fmt.Println(premierLeagueShirt.getBrand())

	fmt.Println(championsLeagueBall.getBrand())
	fmt.Println(championsLeagueShirt.getBrand())
}
