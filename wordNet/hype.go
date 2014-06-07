package wordNet

import (
  // "fmt"
  // "strings"
)

/*****
 * Basic tree node to store the ontology tree.
 */
type treeNode struct {
  name string
  path []int
  children []*treeNode
}

/*****
 * Create a new node with the passed name.
 */
func newTreeNode(word string, path []int) (node *treeNode) {
  newPath := make([]int, len(path))
  copy(newPath, path)
  node = &treeNode{word, newPath, make([]*treeNode, 0)}
  return
}

/*****
 * Look up a specific sense of a word in WordNet and place it in the ontology tree.
 */
func hypeQuery(word string, senseNo int) *treeNode {

  // rawBytes := rawHypeQuery(word, senseNo)
  // tmp := string(rawBytes)
  // splitLine := strings.Split(tmp, "\n")
  // for index, value := range splitLine {
  //   fmt.Println(index, value)
  // }
  return nil
}

/*****
 * Abstraction for WordNet queries. Returns the unprocessed bytes produced
 * by the query.
 */
func rawHypeQuery(word string, senseNo int) []byte {

  argument := "-hypen"

  return wordNetQuery(word, argument, senseNo)
}
