// package main

// import (
// 	"errors" //errors library used for error handling
// 	"fmt"
// 	"os"
// )

// func isGreaterThanTen(num int) error {
// 	if num < 10 {
// 		return errors.New("something bad happened")
// 	}
// 	return nil
// }

// func openFile() error {
// 	f, err := os.Open("missingFile.txt") //go doc os.open will define method
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return nil
// }

// func main() {
// 	num := 9
// 	//err := isGreaterThanTen(num)  <- instead of announcing this here can do in the if block and make err reuseable
// 	if err := isGreaterThanTen(num); err != nil {
// 		fmt.Println(fmt.Errorf("%d is NOT GREATER THAN TEN", num))
// 		panic(err) //nothign else can happen after panic is called - kills program.
// 		//log.Fatalln(err) //log in error logs <- doesn work because panic kills program
// 	}

// 	if err := openFile(); err != nil {
// 		fmt.Println(fmt.Errorf("%v", err))
// 	}
// }

// TAKE A MINUTE TO REFACTOR THE ABOVE CODE TO SCOPE THE ERROR VARIABLE INTO THE IF BLOCK

// ****************************************************

// PANIC & DEFER SLIDE

// ****************************************************

// package main

// import (
// 	"fmt"
// )

// func doThings() {
// 	defer fmt.Println("First Line but do this last!") //defers are last in first out
// 	defer fmt.Println("Do this second to last!")
// 	fmt.Println("Things And Stuff should happen first")
// }

// func main() {
// 	doThings()
// }

// ****************************************************

// RECOVER SLIDE

// ****************************************************

// package main

// import (
// 	"fmt"
// )

// func doThings() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		if i == 2 {
// 			panic("PANIC!")
// 		}
// 	}
// }

// func main() {
// 	doThings()
// }

// ****************************************************

// package main

// import (
// 	"fmt"
// )

// func handlePanic() string {
// 	return "HANDLING THE PANIC"
// }

// func recoverFromPanic() {
// 	// recover() will only return a value if there has been a panic
// 	if r := recover(); r != nil {
// 		fmt.Println("We panicked but everything is fine.")
// 		fmt.Println("Panic instructions received:", r)
// 	}
// }

// func doThings() {
// 	defer recoverFromPanic()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		if i == 2 {
// 			panic(handlePanic())
// 		}
// 	}
// }

// func main() {
// 	doThings()
// }
