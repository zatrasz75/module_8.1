# module_8.1
Skillfactory

Создайте пакет electronic и добавьте в него следующие интерфейсы:

Phone с методами:
Brand() string
Model() string
Type() string
StationPhone с методом:
ButtonsCount() int
Smartphone с методом:
OS() string //название операционной системы
Добавьте в пакет electronic структуру applePhone. Сделайте так, чтобы она реализовывала интерфейсы Phone и Smartphone.

Добавьте в пакет electronic структуру androidPhone. Сделайте так, чтобы она также реализовывала интерфейсы Phone и Smartphone.

Добавьте в пакет electronic структуру radioPhone. Сделайте так, чтобы она реализовывала интерфейсы Phone и StationPhone.

Учтите, что у телефонов Apple бренд всегда один и тот же (apple), в то время как android телефоны производятся множеством брендов (Samsung, LG, Xiaomi и так далее).

Подумайте, какие поля нужны структурам applePhone, androidPhone и radioPhone. Для каждой из этих структур напишите функции-конструкторы.

Метод Type для структур applePhone и androidPhone должен возвращать строку "smartphone", а для структуры radioPhone строку "station".

В пакете main создайте функцию main(). Внутри неё создайте по одному экземпляру каждой структуры из пакета electronic.

Добавьте в пакет main функцию printCharacteristics, которая принимает на вход параметр типа Phone (интерфейс). Функция printCharacteristics должна выводить бренд, модель и тип телефона.

Также:

если телефон также является стационарным телефоном, то должно быть выведено количество его кнопок,
если телефон является смартфоном, то должно быть выведено название его операционной системы.
