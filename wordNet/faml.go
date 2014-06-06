package wordNet

import (
  "bytes"
  "strings"
  "strconv"
)

/*****
 * Query WordNet for the familiarity of a wordtype. It returns the polysemy
 * count (the number of senses).
 */
func famlQuery(word string, wordType int) int {

  rawBytes := rawFamlQuery(word, wordType)

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

/*****
 * Abstraction for WordNet -faml queries. Returns the unprocessed bytes produced
 * by the query.
 */
func rawFamlQuery(word string, wordType int) []byte {

  argument := "-faml"

  switch wordType {
    case NOUN:
      argument += "n"
    case VERB:
      argument += "v"
    case ADJ:
      argument += "a"
    case ADV:
      argument += "r"
  }

  return wordNetQuery(word, argument)
}
