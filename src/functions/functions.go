package functions
import (
	"sort"
	"fmt"
)

const size = 4

// --> estrutura da pagina
type Leaf struct {
	root *Leaf
	values []int
	pointers []*Leaf
}

func CreateLeaf() Leaf{
	var leaf Leaf
	leaf.values = make([]int, size)
	leaf.pointers = make([]*Leaf, size+1)

	return leaf
}

func Print(leaf Leaf) {
	fmt.Printf("\n{\n  root:%p,\n  values:%v,\n  pointers:%v\n}\n\n", leaf.root, leaf.values,  leaf.pointers)
}

func orderB(vetor []int) []int {
	var auxVetor []int
	for _, r := range vetor {
		if r != 0{
			auxVetor = append(auxVetor, r)
		}
	}

	sort.Ints(auxVetor)

	for r := len(auxVetor); r < len(vetor); r++ {
		auxVetor = append(auxVetor, 0)
	}

	return auxVetor

}

func AddValueToLeaf(leaf Leaf, value int) Leaf {
	for block, val := range leaf.values {
		if val == 0 {
			leaf.values[block] = value
			if block != 3 {break}
		}
	}
	leaf.values = orderB(leaf.values);
	return leaf

}

func CheckForSplit(leaf Leaf) bool {
	return leaf.values[3] != 0
}

func Split(leaf Leaf) {
	upLeaf := CreateLeaf()
	leftLeaf := CreateLeaf()
	rightLeaf := CreateLeaf()
	
	if leaf.root != nil {

	} else {

		upLeaf.values[0] = leaf.values[1]
		upLeaf.pointers[0] = &leftLeaf
		upLeaf.pointers[1] = &rightLeaf

		leftLeaf.root = &upLeaf
		leftLeaf.values[0] = leaf.values[0]

		rightLeaf.root = &upLeaf
		rightLeaf.values[0] = leaf.values[2]
		rightLeaf.values[1] = leaf.values[3]

	}

	Print(upLeaf)
	Print(leftLeaf)
	Print(rightLeaf)

}