package main

import (
  "os"
  "fmt"
  "flag"
  "bufio"
  "strings"
  "strconv"
  "os/exec"
  "syscall"
  "encoding/csv"
  "github.com/Arcania0311/textParse/wordNet"
)

/*****
 * Variables used for command line parsing.
 */
var (
  FRAMEWIDTH int
  FAMLTHRESHOLD int
  fileName string
)

func main() {
  // Parse command line flags first.
  flag.IntVar(&FAMLTHRESHOLD, "faml", 2, "The polysemy count at which a word is tokenised.")
  flag.IntVar(&FRAMEWIDTH, "frame", 5, "The number of tokens per frame.")
  flag.StringVar(&fileName, "file", "goTest", "Name of the .txt file to be processed.")
  flag.Parse()

  // Set the famlthreshold
  wordNet.SetFAMLTHRESHOLD(FAMLTHRESHOLD)

  // Open the file.
  lines, err := importLines(fileName + ".txt")
  if err != nil {
    fmt.Println("readLines: %s", err)
  }

  // Create a map in which the words are stored.
  words := make(map[string]*wordNet.Token)
  text := make([]*wordNet.Token, 0)

  // Ugly way to get rid of the most common words.
  words["the"] = nil
  words["a"] = nil
  words["an"] = nil
  words["is"] = nil
  words["are"] = nil
  words["of"] = nil
  words["and"] = nil
  words["be"] = nil
  words["to"] = nil
  words["in"] = nil
  words["by"] = nil

  // Process all the lines.
  for _, line := range lines {
    line = strings.ToLower(line)
    splitLine := strings.Split(line, " ")

    // Tokenize the words.
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

      // Append all non-nil tokens to the text slice.
      token, ok := words[tmp]
      if (ok && token != nil) {
        // This makes sure all words are only tokenised once.
        text = append(text, token)
      } else if (ok && token == nil) {
        continue
      } else {
        newToken := wordNet.CreateToken(tmp)
        words[tmp] = newToken
        if (newToken != nil) {
          text = append(text, newToken)
        }
      }
    }
  }

  results := make([]int, 0)
  totalScore := 0

  // Process the results.
  for x := 0; x < len(text) - FRAMEWIDTH; x++ {
    score := 0
    for frameIndex := 1; frameIndex <= FRAMEWIDTH; frameIndex++ {
      score += wordNet.CompareTokens(text[x], text[x + frameIndex])
    }
    score = score / FRAMEWIDTH
    totalScore += score
    results = append(results, score)
  }

  totalScore = totalScore / (len(text) - FRAMEWIDTH)

  // Create a csv file to store the results.
  resultsFile := fileName + "-" + strconv.Itoa(FRAMEWIDTH) + "-" + strconv.Itoa(FAMLTHRESHOLD) + "-" + strconv.Itoa(totalScore)
  csvFile, err := os.Create("results/" + resultsFile + ".csv")
  out := csv.NewWriter(csvFile)
  for index, value := range results {
    csvLine := make([]string, 2)
    csvLine[0] = strconv.Itoa(index)
    csvLine[1] = strconv.Itoa(value)
    out.Write(csvLine)
  }
  out.Flush()
  csvFile.Close()

  // Open the file with the JavaScript script.
  binary, lookErr := exec.LookPath("node")
  if lookErr != nil {
      panic(lookErr)
  }
  args := []string{"node", "results/visual/showResults.js", resultsFile}

  env := os.Environ()

  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
      panic(execErr)
  }

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
