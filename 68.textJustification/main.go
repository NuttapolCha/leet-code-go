package main

import "fmt"

func main() {
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 10

	printLines(fullJustify(words, maxWidth))
}

func printLines(lines []string) {
	fmt.Println()
	for _, line := range lines {
		fmt.Printf("|%s|\t\t > len = %d\n", line, len(line))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	ret := make([]string, 0)

	var lastLine bool
	var lineIsDone bool

	for i := 0; i < len(words); {
		lineLength := len(words[i])
		// fmt.Printf("'%s' is a fixed word with length: %d\n", words[i], len(words[i]))

		j := i + 1
		for ; ; j++ {

			lastLine = j > len(words)-1
			if lastLine {
				break
			}

			lineIsDone = lineLength+len(words[j])+1 > maxWidth
			if !lineIsDone {
				lineLength += len(words[j]) + 1
			} else {
				lineIsDone = false // always reset
				break
			}
			// fmt.Printf("included: '%s' with length: %d and 1 space => current lineLength = %d\n", words[j], len(words[j]), lineLength)
		}
		numWords := j - i

		if numWords == 1 || lastLine {
			// fmt.Printf("> lastLine = %v, numWords = %d\n", lastLine, numWords)
			ret = append(ret, leftJustify(words[i:j], maxWidth, lineLength))
		} else {
			ret = append(ret, middleJustify(words[i:j], maxWidth, lineLength))
		}
		i = j
	}

	return ret
}

func leftJustify(words []string, maxWidth, lineLength int) string {
	// fmt.Printf("\n>> leftJustifying: %v >> maxWidth: %d, lineLength: %d\n\n", words, maxWidth, lineLength)
	diff := maxWidth - lineLength // == spaces on right

	line := ""
	for k, word := range words {
		line += word
		if k != len(words)-1 {
			line += " "
		}
	}
	return addSingleSpaces(line, diff, false)
}

func middleJustify(words []string, maxWidth, lineLength int) string {
	// fmt.Printf("\n>> middleJustifying: %v >> maxWidth: %d, lineLength: %d\n\n", words, maxWidth, lineLength)
	diff := maxWidth - lineLength
	spacesNeeded := len(words) - 1
	singleSpaces := diff/spacesNeeded + 1
	extraSpaces := diff % spacesNeeded

	line := ""
	for k, word := range words {
		line += word
		if k != len(words)-1 {
			hasExtraSpace := extraSpaces > 0
			line = addSingleSpaces(line, singleSpaces, hasExtraSpace)
			if hasExtraSpace {
				extraSpaces--
			}
		}
	}

	return line
}

func addSingleSpaces(s string, amount int, hasExtraSpace bool) string {
	ret := s
	for i := 0; i < amount; i++ {
		ret += " "
	}
	if hasExtraSpace {
		ret += " "
	}
	return ret
}
