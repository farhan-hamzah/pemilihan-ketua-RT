package main
import "fmt"
type suara struct{
	suaraRakyat int
}
const NMAX int = 100
func main(){
	var A[NMAX] suara
	var n , i, suaraTercatat int
	fmt.Scan(&n)
	i = 0
	for n != 0{
		A[i].suaraRakyat = n
		fmt.Scan(&n)
		i ++
	}
	suaraTercatat = suaraYangBenar(i, A)
	fmt.Println("Suara Masuk: ", i)
	fmt.Println("Suara Sah: ", suaraTercatat)
}
func suaraYangBenar(i int, A[NMAX] suara) int{
	var j, c int
	c = 0
	for j = 0; j < i; j++{
		if A[j].suaraRakyat >= 1 && A[j].suaraRakyat <= 20{
			c++
		}
	}
	return c
}

