package functions
import "sort"
import "fmt"

const size = 4

// --> estrutura da pagina
type Leaf struct {
	root *Leaf
	values []int
	pointers []*Leaf
}

func Root(leaf Leaf) *Leaf {
	return leaf.root;
}

func Values(leaf Leaf) []int {
	return leaf.values;
}

func Pointers(leaf Leaf) []*Leaf {
	return leaf.pointers
}

func CreateLeaf() Leaf{
	var leaf Leaf
	leaf.values = make([]int, size)
	leaf.pointers = make([]*Leaf, size+1)

	return leaf
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
		
		if block == 3 && leaf.values[block] != 0 {
			fmt.Printf("\n\n entrou na condição \n\n")
			// FAZER O SPLIT DESSA DISGRAÇA



		}
	}
	leaf.values = orderB(leaf.values);
	return leaf

}