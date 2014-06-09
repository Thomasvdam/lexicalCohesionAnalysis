package wordNet

import (
  "strings"
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
  node = &treeNode{word, path, make([]*treeNode, 0)}
  return
}

/*****
 * Look up a specific sense of a word in WordNet and place it in the ontology tree.
 */
func hypeQuery(word string, senseNo int) *treeNode {

  rawBytes := rawHypeQuery(word, senseNo)
  tmp := string(rawBytes)
  splitLine := strings.Split(tmp, "\n")

  // Find the first path.
  lineNo := 0
  for x := 4; x < len(splitLine); x++ {
    splitLine[x] = strings.TrimSpace(splitLine[x])
    splitLine[x] = strings.TrimPrefix(splitLine[x], "=> ")
    if (splitLine[x] == ROOT.name) {
      lineNo = x
      break
    }
  }

  // Work back from the end of the first path.
  prevNode := ROOT
  foundChild := false
  for x := lineNo - 1; x > 3; x-- {

    // Check wether the word is already a child.
    pathIndex := 0
    for index, value := range prevNode.children {
      if (value.name == splitLine[x]) {
        prevNode = value
        foundChild = true
        break
      }
      pathIndex = index
    }

    // Move onto the next line and child in the tree.
    if (foundChild) {
      foundChild = false
      continue
    } else {
      newPath := make([]int, len(prevNode.path))
      copy(newPath, prevNode.path)
      newPath = append(newPath, pathIndex)
      newNode := newTreeNode(splitLine[x], newPath)
      prevNode.children = append(prevNode.children, newNode)
      prevNode = newNode
    }
  }

  return prevNode
}

/*****
 * Abstraction for WordNet queries. Returns the unprocessed bytes produced
 * by the query.
 */
func rawHypeQuery(word string, senseNo int) []byte {

  argument := "-hypen"

  return wordNetQuery(word, argument, senseNo)
}
