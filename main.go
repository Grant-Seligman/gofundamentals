package main

import (
	"fmt"
	"math/rand"
)

// You can test your functions here
func main() {

	/*
		When doing these tasks, it's recommended using fmt.Println or fmt.Printf to test the output of your code to make sure
		it works correctly. This will usually entail console logging the answer directly or console logging the
		invocation (call) of the function so when it returns a value, that value will be logged to the console.
		An example of this would be:  console.log(theFunction(value1,value2))
	*/

	/*
	   Task 1a - Voting Age (not auto tested)

	   Do the following:

	    1. Create a variable called votingAge and assign it a number value

	    2. Console log true if age is 18 or higher

	       HINT: no function required
	*/
	// votingAge := 47
	// if votingAge >= 18 {
	// 	fmt.Println("True")
	// 	fmt.Println("Yup, you gucci.")
	// } else {
	// 	fmt.Println("Bitch you too young!")
	// }

	/*
	   Task 1b - Values (not auto tested)

	   Do the following:
	      1. Declare two variables and assign them values (good names for these might be firstThing and secondThing)
	      2. Use a conditional to check the value of the 1st variable versus the value assigned to the 2nd variable
	      3. Change the value of the 1st variable if the conditional in step 2 is true
	      4. Print the value of the 1st variable

	      HINT: no function required
	*/
	// firstThing := 43
	// secondThing := 99
	// if firstThing != secondThing {
	// 	//fmt.Println("The value of First Thing Var Before", firstThing)
	// 	firstThing = secondThing
	// 	//fmt.Println("The value of First Thing Var After", firstThing)
	// }
	// fmt.Println("The value of First Thing", firstThing)

	/*
			   Task 1c - Convert Strings to Numbers (not auto tested)

			   Do the following:
			      1. Declare a variable with the string type value of "1999"
			      2. Convert the string value of "1999" to a integer value of 1999
			      3. Print the result

			      HINT: look up the stringconv.Atoi method

		         **USER COMMENTS: ADD HINT TO USE ERROR HANDLING.....MAKE A JOKE ABOUT FATALITES PLEASE
	*/
	// strYearOne := "1999"
	// numYearOne, Err := strconv.Atoi(strYearOne)
	// if Err != nil {
	// 	log.Fatal("StrYearOne is not what it seems, YOU DIED")
	// }
	// fmt.Println("YTK or Bust in what Year? -- ", numYearOne)

	// fmt.Println("MULTIPLY ", multiply(2, 3))

	// myAge := 27
	// fmt.Println("My age", myAge, "in dog years is: ", dogYears(myAge))

	fmt.Println("ROCK, PAPER, SCISSORS!!!!!!")

	//var randoCompNum int
	randoChoice := [3]string{"rock", "paper", "scissors"}
	//Intn is exclusive, 2 is the upper bound not 3.
	//randomly picks array index value to create "random" choice for computer
	//rand.Intn returns a value, which then is used to index the array
	compChoice := randoChoice[rand.Intn(3)]
	userChoice := randoChoice[rand.Intn(3)]
	fmt.Println(game(userChoice, compChoice))

}

/*
Task 1d - Multiply

Do the following:
   1. Invoke the multiply function below and pass it two numbers
   2. Receive the parameters: a and b
   3. Multiply a and b and return the answer
*/

func multiply(num1, num2 int) int {
	return num1 * num2
}

//Age in Dog years
/*
Do the following:
 1. Invoke the dogYears function below and pass an age value to it
 2. Use the received value to calculate the age in dog years (1 human year is equal to 7 dog years)
 3. Return the newly calculated age
*/

func dogYears(humanAge int) int {
	return humanAge * 7
}

// Rock, Paper, Scissors - Let's play against the computer!
/*
Do the following:
1. Create a new variable that randomly generates the computer's choice, this must not be done inside the function
2. Use rand.Intn to determine the computer's choice
3. Make a conditional that changes the variable to "rock", "paper", or "scissors" based on the random number

Use the game function below to do the following:
1. Receive 2 parameters: a string with the user's choice of "rock", "paper", or "scissors"
   and the computer's choice of "rock", "paper", or "scissors".
   NOTE: make sure the strings are all lower case or it will not pass the test

2. Return whether the user won, lost, or tied based on these rules of the game described below
   the strings returned need to match these strings below exactly.

   - win should return "you win!"
   - lose should return "you lose!"
   - tie should return "it's a tie"

RULES OF THE GAME: Scissors beats Paper | Paper beats Rock | Rock beats Scissors | Or there's a tie
*/

func game(user string, computer string) string {
	if user == "rock" {
		if computer == "scissors" {
			return "you win!"
		}
		if computer == "paper" {
			return "you lose!"
		} else {
			return "it's a tie!"
		}
	}
	if user == "paper" {
		if computer == "rock" {
			return "you win!"
		}
		if computer == "scissors" {
			return "you lose!"
		} else {
			return "it's a tie!"
		}
	}
	if user == "scissors" {
		if computer == "paper" {
			return "you win!"
		}
		if computer == "rock" {
			return "you lose!"
		} else {
			return "it's a tie!"
		}
	}
	return ""
}

//Metric Converter
//Task 5a - Kilometers to Miles
/*
Using the miles function below do the following:
1. Receive a number of kilometers
2. Convert the number of kiolmeters received to miles
3. Return the number of miles
*/

func miles( /*add your code here*/ ) float64 {
	/*add your code here*/
	return 0
}

//Task 5b - Centimeters to Feet
/*
Using the feet function below do the following:
1. Receive a number of cm
2. Convert the number of cm to feet
3. Return number of feet
*/

func feet( /*add your code here*/ ) float64 {
	/*add your code here*/
	return 0
}

// Let's Sing 99 Bottles of Soda on the Wall!
/*
Using the annoyingSong function below do the following:
1. Receive a starting number
2. The annoying song function should return the following string exactly one time:

    "{number you gave as an argument} bottles of soda on the wall, {number you gave as an argument} bottles of soda, take one down pass it around {number you gave as an argument minus 1} bottles of soda on the wall"

3. Outside of the function, Make a loop that invokes annoying song with a number that decreases until it gets to 1 bottle left.
4. Each time the annoyingSong is run from this loop, it should **Println** the string that was returned.
*/

func annoyingSong( /*add your code here*/ ) string {
	/*add your code here*/
	return ""
}

//Grade Calculator
/*
Using the grade function below do the following:
1. Receive a score out of 100
2. Return the corresponding letter grade following this grade scale:

 90-100 should return 'you got an A'
 80-89 should return 'you got a B'
 70-79 should return 'you got a C'
 60-69 should return 'you got a D'
 below should return 'you got an F'
*/

func grade( /*Your Code here */ ) string {
	/*Your Code here */
	return ""
}

// ========================= STRETCH ==================================
// if you have made it this far then MVP is complete
// everything beyond here is extra credit. there is a very
// good chance we have not covered the topics being tested here
// so you will have to research on your own to figure out a solution
// ====================================================================

//Vowel Counter - How many vowels are there?
/*
Using the vowelCounter function below do the following:
1. Receive a string as a parameter
2. Count and return the number of vowels within that string.  It should handle both capitalized and uncapitalized vowels.

HINT - you may need to study tomorrow's content on slices
HINT - try looking up the .includes() method

NOTE: This function is not tested
*/

func vowelCounter( /*add your code here*/ ) {
	/*add your code here*/
}

/*🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑 Please do not modify anything below this line 🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑*/
// The functions below are used for automated testing.

func sanityCheck() string {
	return "its working"
}
