package main

import "fmt"

// adapter là một mẫu thiết kế structural, cho phép các đối tượng không tương thích cộng tác với nhau.
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
