package main
import "fmt"
func PrintComb2(){
	for i:= 10 ; i<= 98 ; i++{
		for j := i+1 ; j <=99 ; j++ {
	
					fmt.Printf("%d %d, ",  i ,j)

				}
			}
		}


func main() {
	PrintComb2()

 }