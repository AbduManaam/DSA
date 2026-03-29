package binarytree

import "fmt"

type TreeNode struct{
	val int
	left *TreeNode
	right *TreeNode
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
	  root:= &TreeNode{val:rootV}

	  inInx:= m[rootV] 
	  leftsize:= inInx-instart

	  root.left= helper(prestart+1,instart,inInx-1)
	  root.right= helper(prestart+leftsize+1,inInx+1,inend)
	  return root
	}

	return helper(0,0,len(inOrder)-1)
    

}
func InOrder(root *TreeNode){

	if root==nil{
		return
	}
	InOrder(root.left)
	fmt.Print(root.val," ")
	InOrder(root.right)
	
}

func PreOrder(root *TreeNode){

	if root==nil{
		return
	}
	fmt.Print(root.val," ")
	PreOrder(root.left)
	PreOrder(root.right)
	
}
func LevelPrint(root *TreeNode) {
if root == nil {
return
}
queue := []*TreeNode{root} 

for len(queue) > 0 {
	current := queue[0]
	queue = queue[1:] 
	fmt.Print(current.val, " ")

	if current.left != nil {
		queue = append(queue, current.left)
	}
	if current.right != nil {
		queue = append(queue, current.right)
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
	fmt.Print(current.val, " ")

	if current.left != nil {
		queue = append(queue, current.left)
	}
	if current.right != nil {
		queue = append(queue, current.right)
	}
	
}
fmt.Println()
 }
}

func Search(root *TreeNode,k int)bool{

	if root==nil{
       return false
	}
	if root.val==k{
		return true
	}
	return Search(root.left,k)||Search(root.right,k)

}



func Insert(root *TreeNode,v int)*TreeNode{
  
	newNode:=&TreeNode{val: v}
	
	if root==nil{
		return newNode
	}
	queue:=[]*TreeNode{root}

	for len(queue)>0{
	   curr:=queue[0]
	   queue=queue[1:]

	   if curr.left==nil{
		curr.left=newNode
		return root
	   }else{
		queue=append(queue,curr.left)
	   }

	   if curr.right==nil{
		curr.right=newNode
		return root
	   }else{
		queue=append(queue,curr.right)
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

	if curr.val==v{
		return true
	}
    if curr.left!=nil{
		queue=append(queue,curr.left)
	}
    if curr.right!=nil{
		queue=append(queue,curr.right)
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

	if curr.val==oldVal{
		curr.val=newVal
		return true
	}
    if curr.left!=nil{
		queue=append(queue,curr.left)
	}
    if curr.right!=nil{
		queue=append(queue,curr.right)
	}

   }
   return false

}

func Delete(root *TreeNode,value int)*TreeNode{
 
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
       
		if curr.val==value{
			target=curr
		}
		if curr.left!=nil{
			m[curr.left]=curr
			queue = append(queue, curr.left)
		}
		if curr.right!=nil{
			m[curr.right]=curr
			queue = append(queue, curr.right)
		}
		last=curr
	}
	if target==nil{
		return root
	}

	target.val=last.val
	parentOfLast=m[last]

	if parentOfLast !=nil{
		if parentOfLast.left==last{
			parentOfLast.left=nil
		}else if parentOfLast.right == last{
			parentOfLast.right=nil
		}
	}else{
		root=nil
	}
return root
}