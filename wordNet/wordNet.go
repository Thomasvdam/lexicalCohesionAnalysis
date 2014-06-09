package wordNet

import (
  "strconv"
  "os/exec"
  "io/ioutil"
)

// Root node of the ontological tree.
var ROOT = &treeNode{"entity", make([]int, 0), make([]*treeNode, 0)}

/*****
 * Compare one word to another and return a score based on the semantic overlap.
 */
func CreateToken(word string) (newToken *Token) {

  newToken = tokenize(word)
  if (newToken == nil) {
    return
  }

  for sense := 0; sense < len(newToken.sensesN); sense++ {
    newToken.sensesN[sense] = hypeQuery(word, sense + 1)
  }

  return
}

/*****
 * Most basic abstraction of a WordNet query.
 */
func wordNetQuery(word, argument string, senseNo int) []byte {

  // Spawn a WN process with the correct arguments and collect results.
  sense := ""
  if (senseNo != 0) {
    sense = "-n" + strconv.Itoa(senseNo)
  }

  wnCmnd := exec.Command("wn", word, sense, argument)
  wnOut, _ := wnCmnd.StdoutPipe()
  wnCmnd.Start()
  wnBytes, _ := ioutil.ReadAll(wnOut)
  wnCmnd.Wait()
  return wnBytes
}
