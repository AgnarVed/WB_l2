package main

func main() {
	usr := NewBAN("Valya", "1234")
	usr.putMoneyToBAN(13)
	usr.checkBAN()
	usr.getMoneyFromBAN(10)
	usr.getMoneyFromBAN(10)
}
