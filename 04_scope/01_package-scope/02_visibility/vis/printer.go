package vis

// import does not get package level scope
// only file level scope
import "fmt"

// PrintVar is exported because it starts with a capital letter
func PrintVar(){
	fmt.Println(MyName)
	fmt.Println(yourName)
}