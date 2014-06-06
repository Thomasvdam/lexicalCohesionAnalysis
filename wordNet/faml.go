package wordNet

import (
  "bytes"
  "strings"
  "strconv"
)

/*****
 * Determine the most likely word type of a given word.
 */
func determineType(word string) (result int) {

  highest := 0
  nounScore := famlQuery(word, NOUN)
  verbScore := famlQuery(word, VERB)
  adjScore := famlQuery(word, ADJ)
  advScore := famlQuery(word, ADV)

  if (nounScore > highest) {
    highest = nounScore
    result = NOUN
  }
  if (verbScore > highest) {
    highest = verbScore
    result = VERB
  }
  if (adjScore > highest) {
    highest = adjScore
    result = ADJ
  }
  if (advScore > highest) {
    highest = advScore
    result = ADV
  }

  return result
}

/*****
 * Query WordNet for the familiarity of a wordtype. It returns the polysemy
 * count.
 */
func famlQuery(word string, wordType int) int {

  rawBytes := rawWordNetQuery(word, wordType, FAML)

  // Check whether there was a result at all.
  if (bytes.Compare(rawBytes, nil) == 0) {
    return 0
  }

  // Convert to a string for easier (not faster) extraction of the polysemy count.
  tmpString := string(rawBytes)
  stripString := func(r rune) rune {
    switch {
    case r < '0' || r > '9':
      return -1
    }
    return r
  }
  tmp := strings.Map(stripString, tmpString)
  number,_ := strconv.Atoi(tmp)

  return number
}
