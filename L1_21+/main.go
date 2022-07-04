package main

func main() {

	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	//Создаем объект который нужно адаптировать
	windowsMachine := &windows{}
	//Передаем этот объект объекту адаптера который содержит немного измененный метод реализующий интерфейс computer
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
