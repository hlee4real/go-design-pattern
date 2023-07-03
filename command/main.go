package main

// command la mot design pattern thuoc behavioral pattern, no chuyen nhung yeu cau hoac hanh dong thanh cac doi tuong rieng biet

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
