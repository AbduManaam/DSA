package sorting

func Selection(arr []int)[]int{

  s:=len(arr)
  for i:=0;i<s-1;i++{
	minInx:=i

	for j:=i+1;j<s;j++{
       if arr[j]<arr[minInx]{
		minInx=j
	   }
	}
	arr[i],arr[minInx]=arr[minInx],arr[i]
  }
  return arr

}