package main
import "fmt"


type Node struct{
	data uint
	next *Node  //in go we cant save direct a struct in same struct we can save that struct address only
}


func main1(){
	fmt.Println("hello")
	v1:=Node{
		data:1,
		next:nil,
	}
	v2:=Node{
		data:2,
		next:nil,
	}
	v3:=Node{
		data:3,
		next:nil,
	}
	CreateNode(&v1)
	CreateNode(&v2)
	CreateNode(&v3)
	Traverse()
}
var count uint=0
var Temp *Node // so when we use struct or others in inside a function than that changes is happend is visible on there not the original code so we use pinters to modify the actual changes and store it to use by all . and when we declare asign value in function after reruning the code new default value assign is nil
var Head *Node
var tail *Node

func CreateNode(myNode *Node){
	var c *uint=&count
	
	if *c==0{
		Head=myNode
		tail=myNode
		Temp=Head
		*c=*c+1 //we use *c because we want to change actual value of c


	}else{
		Temp.next=myNode
		Temp=Temp.next
		tail=Temp
		*c=*c+1
	}
	fmt.Println("created a node")

}
func Traverse(){
	var Temp *Node=Head //head have a address of *node so it can only store in a pointer, means Temp pointer pointing to head and head is store address of node

	if count==0{
		fmt.Println("no node is available ")
	}else{
		for Temp!=nil{
			fmt.Printf("%v->",Temp.data)
			Temp=Temp.next

		}

	}

}