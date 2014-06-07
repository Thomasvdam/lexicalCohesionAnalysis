package wordNet

/*****
 * Basic token struct.
 */
type Token struct {
  word string

  // The position of all the senses per word type.
  sensesN []*treeNode
}

/*****
 * Determine the most likely word type of a given word and return a token which
 * has the slices of the senses set to the correct amount.
 */
func tokenize(word string) (result *Token) {

  nounScore := famlQuery(word)

  // If no results are found for any word type return a negative value.
  if (nounScore == 0) {
    return nil
  }

  result = &Token{word, make([]*treeNode, nounScore)}

  return
}
