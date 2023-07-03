package main

import "fmt"

// decorator la 1 structural pattern cho phep them nhung hanh vi cho cac doi tuong da co san
// bang cach them chung vao trong cac doi tuong dac biet da duoc dong goi lai

func main() {
	milktea := &DingTea{}
	milkteaWithBubble := &BubbleTopping{
		milktea: milktea,
	}
	milkteaWithBubbleAndCheese := &CheeseTopping{
		milktea: milkteaWithBubble,
	}

	fmt.Println("price of dingtea with bubble and cheese topping:", milkteaWithBubbleAndCheese.getPrice())
}
