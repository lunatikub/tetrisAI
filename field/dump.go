package field

import "fmt"

func (f *Field) dumpBLocks() {
	for _, row := range f.Blocks {
		for _, v := range row {
			if v == 1 {
				fmt.Println("X")
			} else {
				fmt.Println(".")
			}
		}
		fmt.Println("")
	}
}
