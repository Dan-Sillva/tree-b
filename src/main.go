package main
import (
	"fmt"
	F "workspace/functions"
)

type List struct {
	init *F.Leaf;
}

var vetor = [9]int{12, 14, 3, 7, 8, 1, 9, 13, 10}
var list List

var leaf F.Leaf

func main() {

	if list.init != nil { //				Tem folha?
		fmt.Println("vetor", list.init)
	} else { // 							Não ?
		leaf = F.CreateLeaf()

		leaf = F.AddValueToLeaf(leaf, vetor[0])
		leaf = F.AddValueToLeaf(leaf, vetor[1])
		leaf = F.AddValueToLeaf(leaf, vetor[2])
		leaf = F.AddValueToLeaf(leaf, vetor[3])

		fmt.Printf("\ncheck for split:%v\n", F.CheckForSplit(leaf))
		if F.CheckForSplit(leaf) {
			list.init = F.Split(leaf)
		} else {
			list.init = &leaf;
		}

		
	}

	F.Print(list.init)

	// if F.Root(leaf) == nil {fmt.Println("ta nulo")}

}
