package model

type Goodbuilding struct {
	Name         string
	Price        int
	Image        string
	Grade        int
	Introduction string
	Number       int
	Production   float64
	Planetname   string
	Username     string
	Mainlandname string
}

/*
	func (goodbuilding *Goodbuilding) Initial(name string, price int, introduction string, image string, num int, production float64) {
		goodbuilding.Name = name
		goodbuilding.Price = price
		goodbuilding.Introduction = introduction
		goodbuilding.Image = image
		goodbuilding.Number = num
		goodbuilding.Production = production
		goodbuilding.Grade = 1
	}
*/
func (goodbuilding *Goodbuilding) Upgrade(production float64) {
	goodbuilding.Grade++
	goodbuilding.Production = production

}
