package repeatedcharacters

import (
	"strconv"
)

func RepeatedCharacters(word string) string {
	stringArray := []rune(word)
	decodedString := ""
	counter := 1

	for i := 0; i < len(stringArray); i++ {
		currentValue := stringArray[i]
		var nextValue rune
		if i+1 < len(stringArray) {
			nextValue = stringArray[i+1]
		}

		if currentValue == nextValue {
			counter++
		}

		if currentValue != nextValue && counter > 1 {
			decodedString += strconv.Itoa(counter) + string(currentValue)
			counter = 1
		} else if currentValue != nextValue {
			decodedString += string(currentValue)
		}
	}

	return decodedString
}
