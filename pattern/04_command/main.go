package main

func main() {
	mazda := &car{}
	//здесь мы указываем, что в onCommand будем хранить команду запуска мазды
	onCommand := &powerOnCommand{
		device: mazda,
	}
	//здесь мы указываем, что в offCommand будем хранить команду остановки работы мазды
	offCommand := &powerOffCommand{
		device: mazda,
	}
	//теперь, если мы захотим запускать мазду из другого места, там достаточно в объект запуска поместить команду
	onButton := &distanceConsole{
		command: onCommand,
	}
	onButton.press()
	offButton := &distanceConsole{
		command: offCommand,
	}
	offButton.press()
}
