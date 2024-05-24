package printer

import (
	"fmt"

	"github.com/AndrewSorokin540/home_work_basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}

	fmt.Println(str)
}
