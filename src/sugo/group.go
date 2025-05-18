package sugo

import "fmt"

func (g *Group) Print(layer int) {
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Name : ", g.Name)
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Url : ", g.Url)

	fmt.Println("")
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("index page")

	if g.Index != nil {
		g.Index.Print(layer)
	}
}
