package main
import "fmt"

const size = 4;

// --> slice a ser ordenado e arvore em si
var slice = []int{1, 2, 3, 4};

// --> estrutura da pagina
type Leaf struct {
	values []int;
	pointers []*int;
}

//funcoes para operações
func createPage() Leaf{
	var leaf Leaf;
	leaf.values = make([]int, size);
	leaf.pointers = make([]*int, size+1);

	return leaf;
}

func main() {
	leaf := createPage()
	
	fmt.Println(leaf)

	// for page := range tree {
	// 	fmt.Println(page);
	// }

}

	// i := []int{}
	// i = append(i, 2)

	// var pointer *Page 
	// pointer = &pagee

	/*



	[page1, {
		value =[],
		pointers = [*page2, *page3]
	}]
	[page2, {
		value =[],
		pointers = []
	}]
	[page1, {
		value =[],
		pointers = []
	}]




	*/

	//i := []int{1, 2, 3};
	//ip := []*int{&i[0]};
