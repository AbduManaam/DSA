package heap


type InHeap []int 

func(h InHeap)Len()int{
   return len(h)
}

func(h InHeap)Less(i,j int)bool{
	return h[i]<h[j]
}

func(h InHeap)Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}

func(h *InHeap)Push(x interface{}){
   *h=append(*h,x.(int))
}

func(h *InHeap)Pop()interface{}{
   
	old:=*h
	size:= len(old)
	val:=old[size-1]
	*h=(*h)[0:size-1]
	return val
}
