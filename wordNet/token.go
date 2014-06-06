package wordNet

/*****
 * Basic token struct.
 */
type token struct {
  word string

  // The position of all the senses per word type.
  sensesN []*treeNode
  sensesV []*treeNode
  sensesA []*treeNode
  sensesR []*treeNode
}

/*****
 * Determine the most likely word type of a given word and return a token which
 * has the slices of the senses set to the correct amount.
 */
func tokenize(word string) (result *token) {

  nounScore := famlQuery(word, NOUN)
  verbScore := famlQuery(word, VERB)
  adjScore := famlQuery(word, ADJ)
  advScore := famlQuery(word, ADV)

  // If no results are found for any word type return a negative value.
  if (nounScore == 0 && verbScore == 0 && adjScore == 0 && advScore == 0) {
    return nil
  }

  result = &token{word, make([]*treeNode, nounScore), make([]*treeNode, verbScore), make([]*treeNode, adjScore), make([]*treeNode, advScore)}

  return
}
