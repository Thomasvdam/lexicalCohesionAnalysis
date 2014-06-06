package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "strings"
  "github.com/Arcania0311/textParse/wordNet"
)

/*****
 * Save the word types as constants.
 */
const (
  NOUN = iota
  VERB
  ADJ
  ADV
)

/*****
 * Struct which stores details about a specific token.
 */
type wordToken struct {
  id int
  count int
}

/*****
 * Open the specified file and return a slice of its lines.
 */
func importLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func main() {
  // Open the file.
  lines, err := importLines("goTest.txt")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }

  // Create a map in which the words are stored.
  id := 0
  words := make(map[string]*wordToken)

  // Process all the lines.
  for _, line := range lines {
    line = strings.ToLower(line)
    splitLine := strings.Split(line, " ")

    // Place all the words in the map with a corresponding id.
    for _, tmp := range splitLine {

      // Remove special characters from the string.
      stripWord := func(r rune) rune {
        switch {
          case r < 'a' || r > 'z':
            return -1
        }
        return r
      }
      tmp = strings.Map(stripWord, tmp)

      token, ok := words[tmp]
      if ok {
        token.count++
      } else {
        words[tmp] = &wordToken{id, 1}
        id++
      }
    }

  }

  // Ugly way to get rid of the most common words.
  commonWords := make(map[string]bool)
  commonWords["the"] = true
  commonWords["a"] = true
  commonWords["an"] = true
  commonWords["is"] = true
  commonWords["are"] = true
  commonWords["of"] = true
  commonWords["and"] = true
  commonWords["be"] = true
  commonWords["to"] = true
  commonWords["in"] = true
  commonWords["by"] = true


  for key, value := range words {
    if !commonWords[key] && value.count > 8 {
      fmt.Println(value.count, key)

      fmt.Println(wordNet.Compare(key, key, NOUN, NOUN))
    }
  }
}
