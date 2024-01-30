package model

type Planet struct {
	Name       string
	Image      string
	Allenergy  float64
	Restenergy float64
	Username   string //`gorm:"size:255"`
}

// 初始化一个星球
func (planet *Planet) Initializeplanet(planetimage string, planetname string, username string) {
	planet.Username = username
	planet.Image = planetimage
	planet.Name = planetname
	planet.Allenergy = 1
	planet.Restenergy = planet.Allenergy
}

func (planet *Planet) Planetenergy(mainlands [5]Mainland) {
	var i int
	for i = 0; i < 5; i++ {
		if mainlands[i].Status == true {
			planet.Restenergy = planet.Restenergy + mainlands[i].Allproduct
			planet.Allenergy = planet.Allenergy + mainlands[i].Allproduct
		}
	}

}

// 初始化一个星球
/*func (planet *Planet) Initializeplanet(planetimage string, planetname string, username string) {
	planet.Username = username
	planet.Image = planetimage
	planet.Name = planetname
	//第一个
	planet.Mainlands[0].Name = "西伦瑞亚" //初始化5个大洲的称号，气候，地形
	planet.Mainlands[0].Climate = "热带雨林气候"
	planet.Mainlands[0].Terrian = "森林"
	planet.Mainlands[0].Successclean = 10
	//第二个
	planet.Mainlands[1].Name = "米尔勒拉"
	planet.Mainlands[1].Climate = "高原山地气候"
	planet.Mainlands[1].Terrian = "山地，丘陵"
	planet.Mainlands[1].Price = 10
	planet.Mainlands[1].Successclean = 20
	//第三个
	planet.Mainlands[2].Name = "乌兰宇蒂"
	planet.Mainlands[2].Climate = "热带草原气候"
	planet.Mainlands[2].Terrian = "草原"
	planet.Mainlands[2].Price = 20
	planet.Mainlands[2].Successclean = 30
	//第四个
	planet.Mainlands[3].Name = "碦拉玛干"
	planet.Mainlands[3].Climate = "热带沙漠气候"
	planet.Mainlands[3].Terrian = "沙漠"
	planet.Mainlands[3].Price = 30
	planet.Mainlands[3].Successclean = 40
	//第五个
	planet.Mainlands[4].Name = "云格雷诺"
	planet.Mainlands[4].Climate = "极地气候"
	planet.Mainlands[4].Terrian = "冰川"
	planet.Mainlands[4].Price = 40
	planet.Mainlands[4].Successclean = 50
	var i int
	for i = 1; i < 5; i++ {
		planet.Mainlands[i].Planetname = planetname
		planet.Mainlands[i].Status = false
	}
	planet.Allenergy = 1
	planet.Restenergy = planet.Allenergy
}
*/
// 星球总净化量和剩余净化量的计算
/*func (planet *Planet) Allenergynumber() {
	var i int
	for i = 0; i < 5; i++ {
		if planet.Mainlands[i].Status == true {
			planet.Restenergy = planet.Restenergy + planet.Mainlands[i].Allproduct
			planet.Allenergy = planet.Allenergy + planet.Mainlands[i].Allproduct
		}
	}
}
*/
