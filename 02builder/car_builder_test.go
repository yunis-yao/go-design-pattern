package builder

import (
	"testing"
)

func TestDirector(t *testing.T) {
	direct := NewDirector()

	//生产5个 C 类型宝马模型
	for i := 0; i < 5; i++ {
		direct.Construct_BMWmodel_Type_C().Run()
	}

	t.Log("---------------------------------------------------------")

	//生产2个 A 类型奔驰模型
	for i := 0; i < 2; i++ {
		direct.Construct_BenzModel_Type_A().Run()
	}

	t.Log("---------------------------------------------------------")

	//生产3个 B 类型奔驰模型
	for i := 0; i < 3; i++ {
		direct.Construct_BenzModel_Type_B().Run()
	}

}
