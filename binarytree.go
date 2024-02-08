package main
import(
	"fmt"
)

type Tree struct{
	data uint
	Right *Tree
	Left *Tree
}
func main(){
	fmt.Println("binary tree")

   
	CreateTree()
	traversal(Origin)
	


}
var Origin *Tree
var temp *Tree
var Level int=0
func CreateTree(){
	var inleft,inright bool
	var n uint
	var l *int=&Level
	
	if Origin==nil{
		Origin=&Tree{
			data:10,
			Right:nil,
			Left:nil,
		}
		temp=Origin
		*l++
		

		
	}
	fmt.Println("enter true or false for left level: ",*l)
	fmt.Scan(&inleft)

	if inleft==true{
		
		fmt.Println(" enter n value for left")
		fmt.Scan(&n)
	    temp.Left=&Tree{n,nil,nil}
	    temp=temp.Left
		*l++
	    CreateTree()
		
	}

	fmt.Println("enter the true or false for right  level: ",*l)
	fmt.Scan(&inright)
	if inright==true{
		fmt.Println("enter n value for right")
		fmt.Scan(&n)
	    temp.Right=&Tree{n,nil,nil}
	    temp=temp.Right
		*l++
	    CreateTree()
	}
	
	
}


func traversal(shadow *Tree){
	if shadow==nil{
		return

	}
	fmt.Println(shadow.data)
	traversal(shadow.Left)
	traversal(shadow.Right)
	
		
}
