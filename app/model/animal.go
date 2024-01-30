package model

type Animinal struct {
	Name         string
	Price        int
	Introduction string
	Image        string
	Number       int
	Planetname   string
	Mainlandname string
	Username     string
}

// 初始化各种动物
func Initialanimals(username string, planetname string) {
	var i int
	for i = 0; i < 20; i++ {
		Animals[i].Username = username
		Animals[i].Planetname = planetname
		Animals[i].Image = Animalsimages[i]
	}
}
