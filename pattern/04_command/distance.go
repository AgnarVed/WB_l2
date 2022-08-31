package main

type distanceConsole struct {
	command command
}

func (d *distanceConsole) press() {
	d.command.execute()
}
