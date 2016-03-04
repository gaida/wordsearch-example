package main

import (
    "fmt"
    "os"
)

// addToResult adds a string to a given slice if not yet in there
func addToResult(result *[]string, word string) {
    found := false
    for _, knownword := range *result {
        if (knownword == word) {
            found = true
        }
    }
    if (!found) {
        *result = append(*result, word)
    }
}

// mainSearch iterates once through the whole puzzle and calls the findWord methor from every position
func mainSearch(puzzle [][]string, words []string) []string {
    result := make([]string, 0, 0)
    x := len(puzzle[0])
    y := len(puzzle)
    
    for row := 0; row < y; row++ {
        for col := 0; col < x; col++ {
            for _, word := range words {
                // iterate through all possible directions and add word if found
                for i := 0; i <= 7; i++ {
                    w := findWord(puzzle, word, col, row, x, y, i)
                    if (w != "") {
                        // add to result if not empty
                        addToResult(&result, w)
                    }
                }
            }
        }
    }
    return result
}

// findWord checks if the word that is beeing searched for is reachable from the current position
// if the correct word is found, a string containing the word is returned
// if the corrent word IS NOT found, an empty string is returned
func findWord(puzzle [][]string, word string, xPos int, yPos int, xSize int, ySize int, dir int) string {
    result := "";
    switch dir {
        case 0:
            // up and left
            if ((0 <= (xPos - len(word)+1) && xPos <= xSize) && (0 <= (yPos - len(word)+1) && yPos <= ySize)) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos-i][xPos-i]
                }
            }
            break;
        case 1:
            // up
            if (0 <= (yPos - len(word)+1) && yPos <= ySize) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos-i][xPos]
                }
            }
            break;
        case 2:
            // up and right
            if ((0 <= (yPos - len(word)+1) && yPos <= ySize) && (0 <= xPos && (xPos + len(word)) <= xSize)) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos-i][xPos+i]
                }
            }
            break;
        case 3:
            // left
            if (0 <= (xPos - len(word)+1) && xPos <= xSize) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos][xPos-i]
                }
            }
            break;
        case 4:
            // right
            if (0 <= xPos && (xPos + len(word)) <= xSize) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos][xPos+i]
                }
            }
            break;
        case 5:
            // left and down
            if ((0 <= (xPos - len(word)+1) && xPos <= xSize) && (0 <= yPos && (yPos + len(word)) <= ySize)) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos+i][xPos-i]
                }
            }
            break;
        case 6:
            // down
            if (0 <= yPos && (yPos + len(word)) <= ySize) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos+i][xPos]
                }
            }
            break;
        case 7:
            // down and right
            if ((0 <= yPos && (yPos + len(word)) <= ySize) && (0 <= xPos && (xPos + len(word)) <= xSize)) {
                for i := 0; i < len(word); i++ {
                    result += puzzle[yPos+i][xPos+i]
                }
            }
            break;
    }
    
    // sanitize output
    if (result != word) {
        result = ""
    }
    
    return result
}

func main() {
    if (len(os.Args) != 4) {
        fmt.Println("sorry, not enough arguments")
        return
    }
    puzzleFile := os.Args[1]
    wordList := os.Args[2]
    output := os.Args[3]
    
    puzzle := CsvPuzzleSlice(puzzleFile)
    words := CsvWords(wordList)
    
    result := mainSearch(puzzle, words)
    WriteOutput(result, output)
}
