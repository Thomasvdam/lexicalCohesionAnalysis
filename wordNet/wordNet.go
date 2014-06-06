package wordNet

import (
  //"fmt"
  "strings"
  "os/exec"
  "io/ioutil"
)

/*****
 * Save the word types as constants.
 */
const (
  NOUN, SYNS = iota, iota
  VERB, ANTS
  ADJ, HYPE
  ADV, FAML
)

type sense struct {
  Synonyms []string
  Meaning string
}

type wordNetResult struct {
  Word string
  Senses []*sense

}

func newResult(word string) (result *wordNetResult) {
  result = &wordNetResult{word, make([]*sense, 20)}
  return
}

/*****
 * Compare one word to another and return a score based on the semantic overlap.
 */
func Compare(wordA, wordB string, typeA, typeB int) int {
  return 1
}

/*****
 * Look up the possible senses of a word and the related synonyms.
 */
func lookUpSyns(word string, wordType int) *wordNetResult {

  wnBytes := rawWordNetQuery(word, wordType, SYNS)

  // Process the data into a more friendly format.
  result := newResult(word)

  wnCmndString := string(wnBytes)
  stripString := func(r rune) rune {
    switch {
    case r == '=' || r == '>':
      return r
    case (r < 'a' || r > 'z') && (r < 'A' || r > 'Z'):
        return ' '
    }
    return r
  }
  tmp := strings.Map(stripString, wnCmndString)

  split := strings.Split(tmp, " ")

  senseID := 0

  for x:= 0; x < len(split); x++ {

    // Skip until a sense is reached.
    if split[x] != "Sense" {
      continue
    }

    x++

    sense := &sense{make([]string, 10), ""}
    synonymID := 0

    // While synonyms remain.
    for split[x] != "=>" && split[x] != "INSTANCE"{
      if (split[x] == "") {
        x++
        continue
      }
      sense.Synonyms[synonymID] = split[x]
      synonymID++
      x++
    }

    x++

    for x < len(split) && split[x] != "Sense" {

      switch split[x] {
      case "INSTANCE":
        x++
        continue
      case "OF=>":
        sense.Meaning = sense.Meaning + "\nInstance Of: "
        x++
        continue
      case "":
        x++
        continue
      }

      sense.Meaning = sense.Meaning + split[x] + " "
      x++
    }
    result.Senses[senseID] = sense
    senseID++
    x--
  }

  return result
}

/*****
 * Abstraction for WordNet queries. Returns the unprocessed bytes produced
 * by the query.
 */
func rawWordNetQuery(word string, wordType, queryType int) []byte {

  argument := ""
  switch queryType {
    case SYNS:
      argument = "-syns"
    case ANTS:
      argument = "-ants"
    case HYPE:
      argument = "-hype"
    case FAML:
      argument = "-faml"
  }

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


  // Spawn a WN process with the correct arguments and collect results.
  wnCmnd := exec.Command("wn", word, argument)
  wnOut, _ := wnCmnd.StdoutPipe()
  wnCmnd.Start()
  wnBytes, _ := ioutil.ReadAll(wnOut)
  wnCmnd.Wait()

  return wnBytes
}
