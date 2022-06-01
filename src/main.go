package main
import (
	F "workspace/functions"
)

var vetor = []int{12, 14, 3, 7, 8, 1, 9, 13, 10}
var list F.List

func main() {

	list = F.CreateList(list, vetor)
	F.Print(F.GetListInit(list))

	// if F.Root(leaf) == nil {fmt.Println("ta nulo")}

}


// leaf = F.AddValueToLeaf(leaf, value)
	
// fmt.Printf("\ncheck for split:%v\n", F.CheckForSplit(leaf))
// if F.CheckForSplit(leaf) { 
// 	list.init = F.Split(leaf)
// } else {
// 	list.init = &leaf;
// }