package builder

import (
	"fmt"
)

//需求：生产出 多种启动特效的奔驰玩具 和 多种启动特效的宝马玩具

//CarModeler 汽车模型接口
type CarModeler interface {
	Start()
	Stop()
	AlArm()
	EngineBoom()
	SetArrayList(sequences []string)
	Run()
}

//BenzModel 奔驰类
type BenzModel struct {
	arrayList []string
}

func (b *BenzModel) Start() {
	fmt.Println("显示 奔驰 的启动灯特效")
}

func (b *BenzModel) Stop() {
	fmt.Println("显示 奔驰 的刹车灯特效")
}

func (b *BenzModel) AlArm() {
	fmt.Println("发出 奔驰 的喇叭声")
}

func (b *BenzModel) EngineBoom() {
	fmt.Println("发出 奔驰 引擎轰鸣的声音")
}

//SetSequence 设置响应顺序
func (b *BenzModel) SetArrayList(sequences []string) {
	b.arrayList = sequences
}

//Run 打开宝马玩具电源后，玩具所做的响应
func (b *BenzModel) Run() {
	for k := range b.arrayList {
		switch b.arrayList[k] {
		case "Start":
			b.Start()
		case "Stop":
			b.Stop()
		case "AlArm":
			b.AlArm()
		case "EngineBoom":
			b.EngineBoom()
		}
	}
	fmt.Println()
}

//BMWmodel 宝马类
type BMWmodel struct {
	arrayList []string
}

func (b *BMWmodel) Start() {
	fmt.Println("显示 宝马 的启动灯特效")
}

func (b *BMWmodel) Stop() {
	fmt.Println("显示 宝马 的刹车灯特效")
}

func (b *BMWmodel) AlArm() {
	fmt.Println("发出 宝马 的喇叭声")
}

func (b *BMWmodel) EngineBoom() {
	fmt.Println("发出 宝马 引擎轰鸣的声音")
}

//SetArrayList 设置响应顺序
func (b *BMWmodel) SetArrayList(sequences []string) {
	b.arrayList = sequences
}

//Run 打开宝马玩具电源后，玩具所做的响应
func (b *BMWmodel) Run() {
	for k := range b.arrayList {
		switch b.arrayList[k] {
		case "Start":
			b.Start()
		case "Stop":
			b.Stop()
		case "AlArm":
			b.AlArm()
		case "EngineBoom":
			b.EngineBoom()
		}
	}
	fmt.Println()
}

//CarBuilder 汽车建造者类
type CarBuilder interface {
	SetSequence(sequences []string)
	Build() CarModeler
}

//BMWBuilder 宝马建造者类
type BMWBuilder struct {
	bMWmodel *BMWmodel
}

func NewBMWBuilder() *BMWBuilder {
	return &BMWBuilder{
		bMWmodel: &BMWmodel{
			arrayList: []string{},
		},
	}
}
func (b *BMWBuilder) SetSequence(sequences []string) {
	b.bMWmodel.SetArrayList(sequences)
}

func (b *BMWBuilder) Build() CarModeler {
	return b.bMWmodel
}

//BenzBuilder 奔驰建造者类
type BenzBuilder struct {
	benzModel *BenzModel
}

func NewBenzBuilder() *BenzBuilder {
	return &BenzBuilder{
		benzModel: &BenzModel{
			arrayList: []string{},
		},
	}
}

func (b *BenzBuilder) SetSequence(sequences []string) {
	b.benzModel.SetArrayList(sequences)
}

func (b *BenzBuilder) Build() CarModeler {
	return b.benzModel
}

//Director 导演类
type Director struct {
	arrayList   []string
	benzBuilder *BenzBuilder
	bMWBuilder  *BMWBuilder
}

func NewDirector() *Director {
	return &Director{
		arrayList:   []string{},
		benzBuilder: NewBenzBuilder(),
		bMWBuilder:  NewBMWBuilder(),
	}
}

//Construct_BenzModel_Type_A 构造A类型奔驰模型
func (d *Director) Construct_BenzModel_Type_A() CarModeler {
	d.arrayList = make([]string, 0)
	d.arrayList = append(d.arrayList, "Start")
	d.arrayList = append(d.arrayList, "AlArm")
	d.benzBuilder.SetSequence(d.arrayList)
	return d.benzBuilder.Build()
}

func (d *Director) Construct_BenzModel_Type_B() CarModeler {
	d.arrayList = make([]string, 0)
	d.arrayList = append(d.arrayList, "EngineBoom")
	d.arrayList = append(d.arrayList, "Stop")
	d.benzBuilder.SetSequence(d.arrayList)
	return d.benzBuilder.Build()
}

func (d *Director) Construct_BMWmodel_Type_C() CarModeler {
	d.arrayList = make([]string, 0)
	d.arrayList = append(d.arrayList, "Start")
	d.arrayList = append(d.arrayList, "AlArm")
	d.bMWBuilder.SetSequence(d.arrayList)
	return d.bMWBuilder.Build()
}
