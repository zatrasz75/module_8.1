package electronic

//----Интерфейсы с Методами

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string // Название операцинной системы
}

// -----------Структуры
type applePhone struct {
	baseSmartphone
}
type androidPhone struct {
	baseSmartphone
}
type radioPhone struct {
	upsBrand string
	upsModel string
	upsPhone int
}
type baseSmartphone struct {
	upsBrand string
	upsModel string
}

//-------Конструкторы методов

func (c *applePhone) Brand() string { return "Apple" } //У телефонов бренд один и тот же
func (c *applePhone) Model() string { return c.upsBrand }
func (c *applePhone) Type() string  { return "smartphone" }
func (c *applePhone) OS() string    { return "iOS" }

func NewApplePhone(model string) *applePhone {
	iphone := new(applePhone)
	iphone.upsBrand = model
	return iphone
}

func (c *androidPhone) Brand() string { return c.upsBrand }
func (c *androidPhone) Model() string { return c.upsModel }
func (c *androidPhone) Type() string  { return "smartphone" }
func (c *androidPhone) OS() string    { return "Android" }

func NewAndroidPhone(brand, model string) *androidPhone {
	android := new(androidPhone)
	android.upsBrand = brand
	android.upsModel = model
	return android
}

func (c *radioPhone) Brand() string     { return c.upsBrand }
func (c *radioPhone) Model() string     { return c.upsModel }
func (c *radioPhone) Type() string      { return "station" }
func (c *radioPhone) ButtonsCount() int { return c.upsPhone }

func NewRadioPhone(brand string, model string, upsPhone int) *radioPhone {
	radioPhone := new(radioPhone)
	radioPhone.upsPhone = upsPhone
	radioPhone.upsBrand = brand
	radioPhone.upsModel = model
	return radioPhone
}
