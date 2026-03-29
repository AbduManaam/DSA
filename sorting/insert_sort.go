package  sorting

//Ascending

func Insertion(arr []int) []int {

   s:=len(arr)

   for i:=1;i<s;i++{
	key:=arr[i]
	j:=i-1
	for j>=0 && arr[j]>key{
		arr[j+1]=arr[j]
		j--
	}
	arr[j+1]=key
   }
return arr
}


