package binarytree

import (
	"fmt"

)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildBST(preorder []int) *TreeNode {
    i := 0
    var helper func(bound int) *TreeNode

    helper = func(bound int) *TreeNode {
        if i == len(preorder) || preorder[i] > bound {
            return nil
        }

        rootVal := preorder[i]
        i++

        root := &TreeNode{Val: rootVal}
        root.Left = helper(rootVal)
        root.Right = helper(bound)

        return root
    }

    return helper(int(^uint(0) >> 1)) // max int
}

func BuildTree(preOrder,inOrder []int)*TreeNode{

	if len(preOrder)==0{
		return nil
	}
	//Map
	m:=make(map[int]int)
	for i,v:=range inOrder{
		m[v]=i
	}

	//Recursion Func
	var helper func(prestart,instart,inend int)*TreeNode
   
	helper=func(prestart,instart,inend int)*TreeNode{
      if prestart>=len(preOrder) || instart>inend{
		return nil
	  }

      rootV:= preOrder[prestart]
	  root:= &TreeNode{Val:rootV}

	  inInx:= m[rootV] 
	  Leftsize:= inInx-instart

	  root.Left= helper(prestart+1,instart,inInx-1)
	  root.Right= helper(prestart+Leftsize+1,inInx+1,inend)
	  return root
	}

	return helper(0,0,len(inOrder)-1)
    

}
func InOrder(root *TreeNode){

	if root==nil{
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Val," ")
	InOrder(root.Right)
	
}

func PreOrder(root *TreeNode){

	if root==nil{
		return
	}
	fmt.Print(root.Val," ")
	PreOrder(root.Left)
	PreOrder(root.Right)
	
}
func LevelPrint(root *TreeNode) {
if root == nil {
return
}
queue := []*TreeNode{root} 

for len(queue) > 0 {
	current := queue[0]
	queue = queue[1:] 
	fmt.Print(current.Val, " ")

	if current.Left != nil {
		queue = append(queue, current.Left)
	}
	if current.Right != nil {
		queue = append(queue, current.Right)
	}
}
}


func LevelPrintPyramid(root *TreeNode) {
if root == nil {
return
}
queue := []*TreeNode{root} 

for len(queue)>0{

	size:=len(queue)
for i:=0;i<size;i++ {
	current := queue[0]
	queue = queue[1:] 
	fmt.Print(current.Val, " ")

	if current.Left != nil {
		queue = append(queue, current.Left)
	}
	if current.Right != nil {
		queue = append(queue, current.Right)
	}
	
}
fmt.Println()
 }
}

func Search(root *TreeNode,k int)bool{

	if root==nil{
       return false
	}
	if root.Val==k{
		return true
	}
	return Search(root.Left,k)||Search(root.Right,k)

}



func Insert(root *TreeNode,v int)*TreeNode{
  
	newNode:=&TreeNode{Val: v}
	
	if root==nil{
		return newNode
	}
	queue:=[]*TreeNode{root}

	for len(queue)>0{
	   curr:=queue[0]
	   queue=queue[1:]

	   if curr.Left==nil{
		curr.Left=newNode
		return root
	   }else{
		queue=append(queue,curr.Left)
	   }

	   if curr.Right==nil{
		curr.Right=newNode
		return root
	   }else{
		queue=append(queue,curr.Right)
	   }

	}
	return root
}

//BFS Search
func Searchh(root *TreeNode,v int)bool{

   
   if root==nil{
	return false
   }

   queue:=[]*TreeNode{root}

   for len(queue)>0{
	curr:=queue[0]
	queue=queue[1:]

	if curr.Val==v{
		return true
	}
    if curr.Left!=nil{
		queue=append(queue,curr.Left)
	}
    if curr.Right!=nil{
		queue=append(queue,curr.Right)
	}

   }
   return false

}
func Update(root *TreeNode,oldVal,newVal int)bool{

   
   if root==nil{
	return false
   }

   queue:=[]*TreeNode{root}

   for len(queue)>0{
	curr:=queue[0]
	queue=queue[1:]

	if curr.Val==oldVal{
		curr.Val=newVal
		return true
	}
    if curr.Left!=nil{
		queue=append(queue,curr.Left)
	}
    if curr.Right!=nil{
		queue=append(queue,curr.Right)
	}

   }
   return false

}

func Delete(root *TreeNode,Value int)*TreeNode{
 
    if root==nil{
		return nil
	}

	var target,last *TreeNode
	var parentOfLast *TreeNode

    queue:=[]*TreeNode{root}
	m:=make(map[*TreeNode]*TreeNode)

	for len(queue)>0{
		curr:=queue[0]
		queue=queue[1:]
       
		if curr.Val==Value{
			target=curr
		}
		if curr.Left!=nil{
			m[curr.Left]=curr
			queue = append(queue, curr.Left)
		}
		if curr.Right!=nil{
			m[curr.Right]=curr
			queue = append(queue, curr.Right)
		}
		last=curr
	}
	if target==nil{
		return root
	}

	target.Val=last.Val
	parentOfLast=m[last]

	if parentOfLast !=nil{
		if parentOfLast.Left==last{
			parentOfLast.Left=nil
		}else if parentOfLast.Right == last{
			parentOfLast.Right=nil
		}
	}else{
		root=nil
	}
return root
}

func Height(root *TreeNode)int{

	if root==nil{
		return 0
	}
	Left:=Height(root.Left)
	Right:=Height(root.Right)
	
	if Left<Right{
		return Right+1
	}
	return Left+1
}

func CountLeaf(root *TreeNode)int{
	if root==nil{
		return 0
	}
	if root.Left==nil && root.Right==nil{
		return 1
	}
	return CountLeaf(root.Left)+CountLeaf(root.Right)
}

func CountNode(root *TreeNode)int{
	if root==nil{
		return 0
	}
	
	return 1+ CountNode(root.Left)+CountNode(root.Right)
}

func CountInterVal(root *TreeNode)int{

	if root==nil || (root.Left==nil || root.Right==nil ){
		return 0
	}
	return 1+CountInterVal(root.Left)+CountInterVal(root.Right)

}

//Diameter

var maxDiameter int

func Diameter(root *TreeNode)int{
	if root==nil{
       return 0
	}

	Left:=Diameter(root.Left)
	Right:=Diameter(root.Right)

	if Left+Right>maxDiameter{
		maxDiameter=Left+Right
	}
	if Left>Right{
		return Left+1
	}
	return Right+1
}


