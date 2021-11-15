import "github.com/emirpasic/gods/maps/hashmap"

func twoSum(nums []int, target int) []int {
    
    // make map
    m := hashmap.New()
    
    // loop the array
    for ind, num := range nums {
        // find complement
        complement := target - num
        // if we find complement in second array, return that value (value is stored indices in original array)
        if cind, found := m.Get(complement); found == true {
            // cind will be before ind
            return []int{cind.(int), ind}
        }
        // put the num and ind in the map
        m.Put(num, ind)
    }
    return nil
}


// O(n^3)
func threeSum(nums []int) [][]int {
    
    var uniqueTriplets [][]int
    
    // make map
    m := make(map[[3]int]bool)
  
    // OUTER LOOP
    for i := 0; i < len(nums)-2; i++ {
        for k := i+1; k < len(nums)-1; k++ {
            for j := k+1; j < len(nums); j++ {
                if nums[i] + nums[k] + nums[j] == 0 {
                    
                    // create a temp array
                    temp := []int{ nums[i], nums[k], nums[j]}
                    // sort temp array to ensure non-duplication
                    sort.Ints(temp)
                    
                    // make key
                    key := [3]int{temp[0], temp[1], temp[2]}
                    
                    
                    // check if we are duplicating avoid it
                    if _, ok := m[key]; !ok {
                        
                        uniqueTriplets = append(uniqueTriplets, temp)
                        
                        m[key] = true            
                        
                        
                    } 
                       
                }
            
            }
        }
    }
    
    return uniqueTriplets
}


// two pointers left and right
func threeSumBetter(nums []int) [][]int {
    
    var uniqueTriplets [][]int
      
    m := make(map[[3]int]bool)

    sorted := nums
    sort.Ints(sorted)
    
    // loop over the sorted array from beginning
    for i := 0; i < len(sorted); i++ {
        
        if i != 0 && sorted[i] == sorted[i-1] {
            continue
        }
        
        left := i + 1
        right := len(sorted) - 1
        
        // while loop, check all values
        for left < right {

            // check our threesum
            sum := sorted[i] + sorted[left] + sorted[right]
            
            if sum == 0 {
                
                key := [3]int{sorted[i], sorted[left], sorted[right]}             
                if _, ok := m[key]; !ok {
                    val := []int{sorted[i], sorted[left], sorted[right]}
                    uniqueTriplets = append(uniqueTriplets, val)
                    m[key] = true
                }
                

                left++
                right--
                
                
            } else if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } 
                
            
        }
        
    }
    return uniqueTriplets

} 
    