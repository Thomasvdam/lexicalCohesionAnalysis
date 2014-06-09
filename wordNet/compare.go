package wordNet

import "fmt"

/*****
 * Compare two tokens by comparing all individual sense pairs and return the
 * highest scoring one.
 */
func CompareTokens(a, b *Token) {

  var highestA, highestB *treeNode
  var hiA, hiB int
  highScore := 0
  for indexA, senseA := range a.sensesN {
    for indexB, senseB := range b.sensesN {
      newScore := compareSenses(senseA, senseB)
      //fmt.Println(newScore)
      if (highScore < newScore) {
        highScore = newScore
        highestA, highestB = senseA, senseB
        hiA, hiB = indexA, indexB
      }
    }
  }

  fmt.Println(highestA, hiA)
  fmt.Println(highestB, hiB)
  fmt.Println(highScore)
}

/*****
 * Compare two sense based on their ontological paths. Return the level at which
 * the paths diverge or 100 if they're synonyms.
 */
func compareSenses(a, b *treeNode) int {

  var index int
  var value int

  if (len(a.path) < len(b.path)) {
    for index, value = range a.path {
      // Return the point of diversion.
      if (value != b.path[index]) {
        return index
      }
    }

    return 100 - (len(b.path) - index - 1) * 10
  } else {
    for index, value = range b.path {
      // Return the point of diversion.
      if (value != a.path[index]) {
        return index
      }
    }
    // Synonym.
    return 100 - (len(a.path) - index - 1) * 10
  }
}
