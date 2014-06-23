package wordNet

var FAMLTHRESHOLD = 2

/*****
 * Basic token struct.
 */
type Token struct {
  Word string

  // The position of all the senses per word type.
  sensesN []*treeNode
}

/*****
 * Function to set the threshold for the -famln query.
 */
func SetFAMLTHRESHOLD(x int)  {
  FAMLTHRESHOLD = x
}

/*****
 * Determine the most likely word type of a given word and return a token which
 * has the slices of the senses set to the correct amount.
 */
func tokenize(word string) (result *Token) {

  nounScore := famlQuery(word)

  // If too few results are found for any word type return a negative value.
  if (nounScore < FAMLTHRESHOLD) {
    return nil
  }

  result = &Token{word, make([]*treeNode, nounScore)}

  return
}
