package wordNet

import (
  "fmt"
  "strings"
  "os/exec"
  "io/ioutil"
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
 * Look up the possible senses of a word and the related synonyms.
 */
func LookUp(word string, wordType int) *wordNetResult {

  // Find the correct argument.
  argument := ""
  switch wordType {
    case 0:
      argument = "-synsn"
    case 1:
      argument = "-synsv"
    case 2:
      argument = "-synsa"
    case 3:
      argument = "-synsr"
  }

  // Spawn a WN process with the correct arguments and collect results.
  wnCmnd := exec.Command("wn", word, argument)
  wnOut, _ := wnCmnd.StdoutPipe()
  wnCmnd.Start()
  wnBytes, _ := ioutil.ReadAll(wnOut)
  wnCmnd.Wait()

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
