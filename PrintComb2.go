package main
import "fmt"
func PrintComb2(){
	for i:= 00 ; i<= 98 ; i++{
		for j := i+1 ; j <=99 ; j++ {
	
					fmt.Printf("%02d %02d, ",  i ,j)
					if i < 98 || j < 99 {
						fmt.Print()

				}
			}
		}

	}
func main() {
	PrintComb2()

 }