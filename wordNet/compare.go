package wordNet

/*****
 * Compare two tokens by comparing all individual sense pairs and return the
 * highest scoring one.
 */
func CompareTokens(a, b *Token) int {

  highScore := 0
  for _, senseA := range a.sensesN {
    for _, senseB := range b.sensesN {
      newScore := compareSenses(senseA, senseB)
      if (highScore < newScore) {
        highScore = newScore
      }
    }
  }

  return highScore
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
        return score(index)
      }
    }
    // Synonym.
    return 1000
  } else {
    for index, value = range b.path {
      // Return the point of diversion.
      if (value != a.path[index]) {
        return score(index)
      }
    }
    // Synonym.
    return 1000
  }
}

/*****
 * Helper function to calculate the exponent of two integers.
 */
func score(y int) (result int) {
  if (y == 0) {
    return 1
  }

  result = 2
  for i := 1; i < y; i++ {
    result *= 2
  }
  return
}
