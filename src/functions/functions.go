package functions
import (
	"sort"
	"fmt"
)

const size = 4

// --> estrutura da pagina
type Leaf struct {
	values []int
	pointers []*Leaf
}

type List struct {
	init *Leaf;
}

func CreateLeaf() Leaf{
	var leaf Leaf
	leaf.values = make([]int, size)
	leaf.pointers = make([]*Leaf, size+1)

	return leaf
}

func DeclareLeaf(values []int, pointers []*Leaf) Leaf{
	var leaf Leaf
	leaf.values = values
	leaf.pointers = pointers

	return leaf
}

func Get(leaf *Leaf) ([]int, []*Leaf) {
	return leaf.values, leaf.pointers
}

func GetListInit(list List) *Leaf{
	return list.init
}

func recursePrint(leaf *Leaf) {
	for _, n := range leaf.pointers {
		if n != nil {
			fmt.Printf("\n{\n  point:%p,\n  values:%v,\n  pointers:%v\n}\n\n", n, n.values, n.pointers)

			recursePrint(n)
		}
	} 
}

func Print(leaf *Leaf) {
	fmt.Printf("\n{\n  =[-ROOT-]=\n\n  values:%v,\n  pointers:%v\n}\n\n", leaf.values, leaf.pointers)

	recursePrint(leaf)

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

	if leaf.values[3] != 0 {
		return Split(leaf)
	} else {
		return leaf
	}

}

func HasPointer(leaf *Leaf) bool {
	return leaf.pointers[0] != nil;
} 

func CreateList(list List, vetor []int) List {
	var leaf Leaf
	leaf = CreateLeaf()

	for _, value := range vetor {
		if list.init != nil { 
			// break
			if HasPointer(list.init) {
				values, pointers := Get(list.init)
		
				for block, interValue := range values {
					isSmaller := (value <= interValue)
					
					if isSmaller {
						leaf = DeclareLeaf(Get(pointers[block]))
						


						leaf = AddValueToLeaf(leaf, value)
						pointers[block] = &leaf
					}
				}
		
			} else {
				leaf = DeclareLeaf(Get(list.init))
				leaf = AddValueToLeaf(leaf, value)
				list.init = &leaf	
			}
	
		} else { 
			leaf := AddValueToLeaf(leaf, value)
			list.init = &leaf;

		}
	}

	return list
}

func CheckForSplit(leaf Leaf) bool {
	return leaf.values[3] != 0
}

func Split(leaf Leaf) Leaf{
	upLeaf := CreateLeaf()
	leftLeaf := CreateLeaf()
	rightLeaf := CreateLeaf()
	
	upLeaf.values[0] = leaf.values[1]
	upLeaf.pointers[0] = &leftLeaf
	upLeaf.pointers[1] = &rightLeaf

	leftLeaf.values[0] = leaf.values[0]

	rightLeaf.values[0] = leaf.values[2]
	rightLeaf.values[1] = leaf.values[3]

	return upLeaf
}
