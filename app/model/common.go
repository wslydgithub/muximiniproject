package model

var Mainlands [5]Mainland

var Mainlandmages [5]string

var Stateimages [15]string

var Animalsimages [20]string

var Goodbuildingimages [4]string

var Badbuildingimages [4]string

var Plantimages [3]string

var Animals = [20]Animinal{
	{"树蛙", 1, "颜色鲜艳", "url", 0, "", "西伦瑞亚", ""},
	{"树懒", 1, "长时间睡眠", "url", 0, "", "西伦瑞亚", ""},
	{"巨型蝴蝶", 1, "彩色的壮观翅膀", "url", 0, "", "西伦瑞亚", ""},
	{"食蚁兽", 1, "长长的舌头", "url", 0, "", "西伦瑞亚", ""},
	{"雪豹", 1, "超强的攀爬与跳跃能力", "url", 0, "", "米尔勒拉", ""},
	{"雪鸡", 1, "羽毛颜色随季节变化而变化", "url", 0, "", "米尔勒拉", ""},
	{"岩羊", 1, "对寒气有耐受力", "url", 0, "", "米尔勒拉", ""},
	{"高原兔", 1, "在岩石中筑巢", "url", 0, "", "米尔勒拉", ""},
	{"非洲象", 1, "群居生活，大型食草动物", "url", 0, "", "乌兰宇蒂", ""},
	{"斑马", 1, "黑白相间的斑马线", "url", 0, "", "乌兰宇蒂", ""},
	{"鸵鸟", 1, "非洲草原的大型鸟类", "url", 0, "", "乌兰宇蒂", ""},
	{"狮子", 1, "顶级捕食者，狩猎的时候到了", "url", 0, "", "乌兰宇蒂", ""},
	{"沙漠狐狸", 1, "大耳朵散热", "url", 0, "", "碦拉玛干", ""},
	{"骆驼", 1, "沙漠地区的象征生物", "url", 0, "", "碦拉玛干", ""},
	{"沙蜥蜴", 1, "体色易于伪装", "url", 0, "", "碦拉玛干", ""},
	{"沙漠猫", 1, "一种小型猫科动物", "url", 0, "", "碦拉玛干", ""},
	{"北极熊", 1, "北极的顶级掠食者", "url", 0, "", "云格雷诺", ""},
	{"帝企鹅", 1, "生活在南极的大型企鹅", "url", 0, "", "云格雷诺", ""},
	{"北极狐", 1, "皮毛颜色随季节变化而变化", "url", 0, "", "云格雷诺", ""},
	{"雷鸟", 1, "生活在北极地区的猫头鹰", "url", 0, "", "云格雷诺", ""},
}

var Plants = [15]Plant{
	{"猪笼草", 1, "食虫植物", "url", "西伦瑞亚", 0, "在之后赋值", ""},
	{"天堂鸟花", 1, "形状类似鸟类展翅", "url", "西伦瑞亚", 0, "在之后赋值", ""},
	{"细叶吊灯花", 1, "食虫植物", "url", "西伦瑞亚", 0, "在之后赋值", ""},
	{"高山杜鹃", 1, "生长在岩石中", "url", "米尔勒拉", 0, "在之后赋值", ""},
	{"高山针茅", 1, "适应了寒冷的气温", "url", "米尔勒拉", 0, "在之后赋值", ""},
	{"高山松", 1, "形成树线上限", "url", "米尔勒拉", 0, "在之后赋值", ""},
	{"乌拉尔大花草", 1, "适应了干燥的气候", "url", "乌兰宇蒂", 0, "在之后赋值", ""},
	{"灰背藜", 1, "灰绿色的叶片和花朵", "url", "乌兰宇蒂", 0, "在之后赋值", ""},
	{"双生大花草", 1, "艳丽的大黄花", "url", "乌兰宇蒂", 0, "在之后赋值", ""},
	{"仙人掌", 1, "叶片有肉质，有助于储存水分", "url", "碦拉玛干", 0, "在之后赋值", ""},
	{"死亡之草", 1, "被称为“死人的手指”", "url", "碦拉玛干", 0, "在之后赋值", ""},
	{"石莲花", 1, "多肉植物，常见于岩石缝隙中", "url", "碦拉玛干", 0, "在之后赋值", ""},
	{"批碱草", 1, "生长在极地地区的草本植物", "url", "云格雷诺", 0, "在之后赋值", ""},
	{"极地牛蹄草", 1, "生长矮小的柳科植物", "url", "云格雷诺", 0, "在之后赋值", ""},
	{"北极茶", 1, "生长力十分强大的苔藓植物", "url", "云格雷诺", 0, "在之后赋值", ""},
}

var Goodbuildings = [20]Goodbuilding{
	{"太阳能发电站", 1, "url", 1, "能源来源易得，安全可靠", 0, 0, "西伦瑞亚", "", ""},
	{"风力发电站", 1, "url", 1, "能源清洁可再生", 0, 0, "西伦瑞亚", "", ""},
	{"地源热泵", 1, "url", 1, "利用地热能的能源系统", 0, 0, "西伦瑞亚", "", ""},
	{"潮汐发电站", 1, "url", 1, "利用潮汐能将潮汐运动转化为电力", 0, 0, "西伦瑞亚", "", ""},
	{"太阳能发电站", 1, "url", 1, "能源来源易得，安全可靠", 0, 0, "米尔勒拉", "", ""},
	{"风力发电站", 1, "url", 1, "能源清洁可再生", 0, 0, "米尔勒拉", "", ""},
	{"地源热泵", 1, "url", 1, "利用地热能的能源系统", 0, 0, "米尔勒拉", "", ""},
	{"潮汐发电站", 1, "url", 1, "利用潮汐能将潮汐运动转化为电力", 0, 0, "米尔勒拉", "", ""},
	{"太阳能发电站", 1, "url", 1, "能源来源易得，安全可靠", 0, 0, "乌兰宇蒂", "", ""},
	{"风力发电站", 1, "url", 1, "能源清洁可再生", 0, 0, "乌兰宇蒂", "", ""},
	{"地源热泵", 1, "url", 1, "利用地热能的能源系统", 0, 0, "乌兰宇蒂", "", ""},
	{"潮汐发电站", 1, "url", 1, "利用潮汐能将潮汐运动转化为电力", 0, 0, "乌兰宇蒂", "", ""},
	{"太阳能发电站", 1, "url", 1, "能源来源易得，安全可靠", 0, 0, "碦拉玛干", "", ""},
	{"风力发电站", 1, "url", 1, "能源清洁可再生", 0, 0, "碦拉玛干", "", ""},
	{"地源热泵", 1, "url", 1, "利用地热能的能源系统", 0, 0, "碦拉玛干", "", ""},
	{"潮汐发电站", 1, "url", 1, "利用潮汐能将潮汐运动转化为电力", 0, 0, "碦拉玛干", "", ""},
	{"太阳能发电站", 1, "url", 1, "能源来源易得，安全可靠", 0, 0, "云格雷诺", "", ""},
	{"风力发电站", 1, "url", 1, "能源清洁可再生", 0, 0, "云格雷诺", "", ""},
	{"地源热泵", 1, "url", 1, "利用地热能的能源系统", 0, 0, "云格雷诺", "", ""},
	{"潮汐发电站", 1, "url", 1, "利用潮汐能将潮汐运动转化为电力", 0, 0, "云格雷诺", "", ""},
}

var Badbuildings = [20]Badbuilding{
	{"矿山", "导致土壤侵蚀，水源污染", 0, 1, "", "西伦瑞亚", ""},
	{"化工厂", "导致空气污染和水污染", 0, 1, "", "西伦瑞亚", ""},
	{"火力发电厂", "产生污染气体危害大气", 0, 1, "", "西伦瑞亚", ""},
	{"垃圾焚化处理场", "污染空气和土壤", 0, 1, "", "西伦瑞亚", ""},
	{"矿山", "导致土壤侵蚀，水源污染", 0, 1, "", "米尔勒拉", ""},
	{"化工厂", "导致空气污染和水污染", 0, 1, "", "米尔勒拉", ""},
	{"火力发电厂", "产生污染气体危害大气", 0, 1, "", "米尔勒拉", ""},
	{"垃圾焚化处理场", "污染空气和土壤", 0, 1, "", "米尔勒拉", ""},
	{"矿山", "导致土壤侵蚀，水源污染", 0, 1, "", "乌兰宇蒂", ""},
	{"化工厂", "导致空气污染和水污染", 0, 1, "", "乌兰宇蒂", ""},
	{"火力发电厂", "产生污染气体危害大气", 0, 1, "", "乌兰宇蒂", ""},
	{"垃圾焚化处理场", "污染空气和土壤", 0, 1, "", "乌兰宇蒂", ""},
	{"矿山", "导致土壤侵蚀，水源污染", 0, 1, "", "碦拉玛干", ""},
	{"化工厂", "导致空气污染和水污染", 0, 1, "", "碦拉玛干", ""},
	{"火力发电厂", "产生污染气体危害大气", 0, 1, "", "碦拉玛干", ""},
	{"垃圾焚化处理场", "污染空气和土壤", 0, 1, "", "碦拉玛干", ""},
	{"矿山", "导致土壤侵蚀，水源污染", 0, 1, "", "云格雷诺", ""},
	{"化工厂", "导致空气污染和水污染", 0, 1, "", "云格雷诺", ""},
	{"火力发电厂", "产生污染气体危害大气", 0, 1, "", "云格雷诺", ""},
	{"垃圾焚化处理场", "污染空气和土壤", 0, 1, "", "云格雷诺", ""},
}
