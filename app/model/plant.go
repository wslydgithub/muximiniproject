package model

type Plant struct {
	Name         string
	Price        int
	Introduction string
	Image        string
	Mainlandname string //`gorm:"size:255"`
	Number       int
	Planetname   string
	Username     string
}

/*func (plant *Plant) Initial(name string, price int, introduction string, image string, num int) {
	plant.Name = name
	plant.Price = price
	plant.Introduction = introduction
	plant.Image = image
	plant.Number = num
}
*/
