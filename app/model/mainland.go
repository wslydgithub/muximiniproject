package model

type Mainland struct {
	Name         string
	Image        string
	Stageimage   string
	Status       bool    //注意默认的是全不未解锁
	Allproduct   float64 //暂无
	Allcleanrate float64 //暂无
	Price        int
	Climate      string
	Terrian      string
	Others       string //暂无
	Planetname   string
	Successclean float64
	Username     string
}

// 初始化五个大陆的基本信息
func Initialmainlands(username string, planetname string) {
	var i int
	for i = 0; i < 5; i++ {
		Mainlands[i].Image = Mainlandmages[i]
		Mainlands[i].Status = false
		Mainlands[i].Username = username
		Mainlands[i].Planetname = planetname
		Mainlands[i].Stageimage = Stateimages[i*3]
	}
	//第一个
	Mainlands[0].Name = "西伦瑞亚" //初始化5个大洲的称号，气候，地形
	Mainlands[0].Climate = "热带雨林气候"
	Mainlands[0].Terrian = "森林"
	Mainlands[0].Successclean = 10
	//第二个
	Mainlands[1].Name = "米尔勒拉"
	Mainlands[1].Climate = "高原山地气候"
	Mainlands[1].Terrian = "山地，丘陵"
	Mainlands[1].Price = 10
	Mainlands[1].Successclean = 20
	//第三个
	Mainlands[2].Name = "乌兰宇蒂"
	Mainlands[2].Climate = "热带草原气候"
	Mainlands[2].Terrian = "草原"
	Mainlands[2].Price = 20
	Mainlands[2].Successclean = 30
	//第四个
	Mainlands[3].Name = "碦拉玛干"
	Mainlands[3].Climate = "热带沙漠气候"
	Mainlands[3].Terrian = "沙漠"
	Mainlands[3].Price = 30
	Mainlands[3].Successclean = 40
	//第五个
	Mainlands[4].Name = "云格雷诺"
	Mainlands[4].Climate = "极地气候"
	Mainlands[4].Terrian = "冰川"
	Mainlands[4].Price = 40
	Mainlands[4].Successclean = 50
}

/*// 初始大陆(badbuilding)
func (mainland *Mainland) Unlockmainland(stageimages [3]string) {
	mainland.Status = true
	var i int
	//初始化大陆上的动物，植物，图片
	for i = 0; i < 4; i++ {
		mainland.Goodbuildings = Goodbuildings[i*4 : i*4+4]
		mainland.Animinals = Animals[i*4 : i*4+4]
	}
	for i = 0; i < 3; i++ {
		mainland.Plants = Plants[i*3 : i*3+3]
		//mainland.Stageimage = common.Statusimage[i*3 : i*3+3]
	}
	mainland.Dayclean_Allproduct(mainland.Badbuildings, mainland.Goodbuildings)

}

// 计算大陆每日净产量+大陆总产量+大陆总净化率
func (mainland *Mainland) Dayclean_Allproduct(badbuildings []Badbuilding, goodbuildings []Goodbuilding) {
	var i = 0
	for i = 0; i < 4; i++ {
		mainland.Report.Pollution = badbuildings[i].Pollution*float64(badbuildings[i].Number) + mainland.Report.Pollution
		mainland.Report.Cleanliness = goodbuildings[i].Production*float64(goodbuildings[i].Number) + mainland.Report.Cleanliness
	}
	mainland.Report.Dayclean = mainland.Report.Pollution + mainland.Report.Cleanliness
	mainland.Allproduct = mainland.Allproduct + mainland.Report.Dayclean
	mainland.Allcleanrate = mainland.Allproduct / mainland.Successclean
	var radis = math.Pow(10, 1)
	mainland.Allcleanrate = math.Round(mainland.Allcleanrate*radis) / radis
	//改变shu
	if mainland.Allcleanrate < 0.5 {
		mainland.Image = mainland.Stageimage[0]
	} else if mainland.Allcleanrate >= 0.5 || mainland.Allcleanrate < 1 {
		mainland.Image = mainland.Stageimage[1]
	} else {
		mainland.Image = mainland.Stageimage[2]
	}
}

// 放置回收动物
func (mainland *Mainland) Placeandback_animal(n int, m int) { //n为动物的类型标号,m为放置的数量
	mainland.Animinals[n].Number = mainland.Animinals[n].Number + m
}

// 放置回收植物
func (mainland *Mainland) Placeandback_plant(n int, m int) { //n为植物的类型标号,m为放置的数量
	mainland.Plants[n].Number = mainland.Plants[n].Number + m
}

// 放置回收环保的建筑物
func (mainland *Mainland) Placeandback_building(n int, m int) { //n为建筑物类型，m为放置的数量
	mainland.Goodbuildings[n].Number = mainland.Goodbuildings[n].Number + m
}

// 拆除污染的建筑物
func (mainland *Mainland) Back_badbuilding(n int, m int) {
	mainland.Badbuildings[n].Number = mainland.Badbuildings[n].Number + m
}
*/
