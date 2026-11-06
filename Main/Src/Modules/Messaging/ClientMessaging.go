package Messaging
import (
	"fmt"
	"bufio"
	"os"
)

// Prompts the user to enter a message.
func PromptMessage(Prefix string,Prompt string) string{
	var UserInp string
	if Prompt != ""{
		fmt.Println(Prompt)
	}	

	
	fmt.Print(Prefix," ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	UserInp = scanner.Text()

	
	for{
		if UserInp != ""{break}
		fmt.Print(Prefix," ")
		scanner.Scan()

		if Err := scanner.Err(); Err != nil {
    		fmt.Println("Error reading input:", Err)
    		continue
		}
		
		UserInp = scanner.Text()
			
		
	}


	return UserInp
}

func DisplayMessage(Sender string,Message string){
	fmt.Println(Sender,": ",Message)
}