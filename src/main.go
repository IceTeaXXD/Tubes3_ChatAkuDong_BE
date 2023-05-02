package main

import (
	"fmt"
	"github.com/IceTeaXXD/Tubes3_13521004/server/algo"
)

func main() {
	fmt.Println(Algo.KMP("keren", "Siapakah Dosen Terkeren STEI?"))
	fmt.Println(Algo.BMMatch("keren", "Siapakah Dosen Terkeren STEI?"))
	fmt.Println(Algo.MatchRatio("Dosen Keren STEI", "Siapakah Dosen Terkeren STEI?"))
	Algo.Day("01/05/2023")
	fmt.Println(Algo.Evaluate("5+9*(3*2)/4"))
	fmt.Println(Algo.Regex(1, "Hari apa tanggal 01/05/2023"));
	fmt.Println(Algo.Regex(2, "1+1*2"));
}
