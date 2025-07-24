package main
import (
	"fmt"
)

func main(){

	fmt.Println("In Main Function.")
}
	/* ____init function___

	Does not take any parameter.
	Does not return any value.
	Can be declared multiple.
	use to initilize connection,entities
	*/
func init(){
	fmt.Println("init,Always Executed First..!!")
}
func init(){
	fmt.Println("2nd init ")
}