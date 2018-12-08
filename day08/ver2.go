package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

type node struct {
	children []node
	metadata []int
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

	var metadata = make([]int, metaCount)

	for i := 0; i < metaCount; i++ {
		data_str, newSrc := next(src)
		src = newSrc

		data, err := strconv.Atoi(data_str)
		if err != nil { log.Fatal(err) }

		metadata[i] = data
	}

	return node {
		children: children,
		metadata: metadata,
	}, src
}

func getNodeValue(root node) int {
	if len(root.children) == 0 {
		var total = 0
		for _, val := range root.metadata {
			total += val
		}
		return total
	} else {
		var total = 0
		for _, val := range root.metadata {
			target := val - 1
			if target >= 0 && target < len(root.children) {
				childValue := getNodeValue(root.children[target])
				total += childValue
			}
		}
		return total
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	stuff := strings.Split(strings.TrimSpace(string(bytes)), " ")

	tree, _ := readNode(stuff)

	total := getNodeValue(tree)

	fmt.Println(total)
}
