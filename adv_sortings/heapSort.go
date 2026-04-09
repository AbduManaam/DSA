package advsortings

func heapSort(arr []int) {
	n := len(arr)

	// Build heap
	for i := n/2 - 1; i >= 0; i-- {   //(i := n/2 - 1)=last node that has children to work heap func (swap)
		heapify(arr, n, i)            //  // Convert array into max heap
	}

	// Extract elements
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)            //  // Fix heap after removing max element
	}  
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)              //   Continue fixing heap downward
	}
}

