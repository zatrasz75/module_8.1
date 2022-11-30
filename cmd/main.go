package main

import (
	"fmt"
	"inter/pkg/electronic"
)

func main() {

	//экземпляр структуры из пакета electronic.
	newIPhone := electronic.NewApplePhone("iPhone 12")
	printCharacteristics(newIPhone)

	//экземпляр структуры из пакета electronic.
	newAndroid := electronic.NewAndroidPhone("Xiaomi", "8 Pro")
	printCharacteristics(newAndroid)

	//экземпляр структуры из пакета electronic.
	newRadio := electronic.NewRadioPhone("Прошлый век", "13", 9)
	printCharacteristics(newRadio)

}

//-- функцию printCharacteristics, которая принимает на вход параметр типа Phone (интерфейс).
func printCharacteristics(phone electronic.Phone) {
	//--Функция printCharacteristics должна выводит бренд, модель и тип телефона
	fmt.Printf("Бренд: %v  Модель: %v Тип: %v ", phone.Brand(), phone.Model(), phone.Type())

	//--На смартфоне добавлен вывод операционной системы из интерфейса Smartphone
	smart, ok := phone.(electronic.Smartphone)
	if ok {
		fmt.Printf("OS %s \n", smart.OS())
	}
	//--На стационарном телефоне добавлен вывод кнопок из интерфейса StationPhone
	radio, ok := phone.(electronic.StationPhone)
	if ok {
		fmt.Printf("Количество кнопок %v \n", radio.ButtonsCount())
	}

	fmt.Println()

}
