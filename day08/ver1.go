package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

type node struct {
	children []node
	metadata []string
}

func next(src []string) (string, []string) {
	a := src[0]
	return a, src[1:]
}

func readNode(src []string) (node, []string) {
	childCount_str, src := next(src)
	childCount, err := strconv.Atoi(childCount_str)
	if err != nil { log.Fatal(err) }

	metaCount_str, src := next(src)
	metaCount, err := strconv.Atoi(metaCount_str)
	if err != nil { log.Fatal(err) }

	var children = make([]node, childCount)

	for i := 0; i < childCount; i++ {
		child, newSrc := readNode(src)
		src = newSrc

		children[i] = child
	}

	var metadata = make([]string, metaCount)

	for i := 0; i < metaCount; i++ {
		data, newSrc := next(src)
		src = newSrc

		metadata[i] = data
	}

	return node {
		children: children,
		metadata: metadata,
	}, src
}

func getMetadataSum(tree node) int {
	var total = 0

	for _, val_str := range tree.metadata {
		val, err := strconv.Atoi(val_str)
		if err != nil { log.Fatal(err) }

		total += val
	}

	for _, child := range tree.children {
		total += getMetadataSum(child)
	}

	return total
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	stuff := strings.Split(strings.TrimSpace(string(bytes)), " ")

	tree, _ := readNode(stuff)

	total := getMetadataSum(tree)

	fmt.Println(total)
}
