package main

// fmt package provides the function to print anything
import "fmt"

func main() {
	// declaring the variable sentence of string type
	// which stores the sentence in which we have to
	// count the vowels and the Consonants
	var sentence string

	// declaring the variables to store the count
	// of vowels and Consonants
	var vowelsCount, consonantCount int

	// initializing the variable sentence
	sentence = "India have twenty eight states and eight union territories"

	// initializing the variable vowelsCount
	vowelsCount = 0

	// initializing the variable ConsonantCount
	consonantCount = 0

	// running a for loop over the string stored in the sentence variable
	for i := 0; i < len(sentence); i++ {
		// skipping the spaces in the sentence
		if sentence[i] == ' ' {
			continue
		}

		// comparing the current character with the vowels
		if sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] == 'i' || sentence[i] == 'o' || sentence[i] == 'u' ||
			sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] == 'I' || sentence[i] == 'O' || sentence[i] == 'U' {

			// increasing the count of vowelCount variable
			// if the current character is a vowel
			vowelsCount++
		} else {
			// increasing the count of consonantCount variable
			// if current character is consonant
			consonantCount++
		}
	}
	fmt.Println("Sentence :- ", sentence)

	// printing the count of vowels and consonants
	fmt.Println("Result   :-  The total number of vowels in the above sentence are", vowelsCount)
	fmt.Println("\t     The total number of consonants in the above sentence are", consonantCount)
}
