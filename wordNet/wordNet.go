package wordNet

import (
  "os/exec"
  "io/ioutil"
)

/*****
 * Save the word types and query types as constants.
 */
const (
  NOUN = iota
  VERB
  ADJ
  ADV
)

var ROOT = &treeNode{"entity", nil, make([]*treeNode, 0)}

/*****
 * Compare one word to another and return a score based on the semantic overlap.
 */
func CreateToken(word string) (newToken *token) {

  newToken = tokenize(word)
  if (newToken == nil) {
    return
  }

  for sense := 0; sense < len(newToken.sensesN); sense++ {
    newToken.sensesN[sense] = hypeQuery(word, sense, NOUN)
  }
  for sense := 0; sense < len(newToken.sensesV); sense++ {
    newToken.sensesV[sense] = hypeQuery(word, sense, VERB)
  }
  for sense := 0; sense < len(newToken.sensesA); sense++ {
    newToken.sensesA[sense] = hypeQuery(word, sense, ADJ)
  }
  for sense := 0; sense < len(newToken.sensesR); sense++ {
    newToken.sensesR[sense] = hypeQuery(word, sense, ADV)
  }

  return
}


/*****
 * Most basic abstraction of a WordNet query.
 */
func wordNetQuery(word, argument string) []byte {

  // Spawn a WN process with the correct arguments and collect results.
  wnCmnd := exec.Command("wn", word, argument)
  wnOut, _ := wnCmnd.StdoutPipe()
  wnCmnd.Start()
  wnBytes, _ := ioutil.ReadAll(wnOut)
  wnCmnd.Wait()

  return wnBytes
}
