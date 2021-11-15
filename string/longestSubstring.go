import ("fmt")

// this is very slow
func lengthOfLongestSubstring(s string) int {
    
    m := make(map[string]int)
    
    max := 0
    count := 0
    for l, r := 0, 0; i < len(s); i++{
        if _, ok := m[string(s[i])]; !ok{
            count++
            m[string(s[i])] = i
        } else {
            if count > max {
                max = count
            }
            count = 0
            i = m[string(s[i])]
            m = make(map[string]int)
        }
    }
    if count > max {return count;}
    return max
}


// instead of checking i+1 after the last dupliate, update the value in the hashmap, and keep the right pointer
// https://dev.to/leihuang_/algorithms-in-go-length-of-the-longest-substring-2jm5
func lengthOfLongestSubstringBetter(s string) int {
    // making a map the go way
    m := make(map[byte]int)
    res := 0

    for l, r := 0, 0; r < len(s); r++ {
        if index, ok := m[s[r]]; ok {
            // only update the left pointer if 
            // it's behind the last position of the visited character
            l = max(l, index)
        }
        res = max(res, r-l+1)
        m[s[r]] = r + 1
    }
    return res
}

func max(n, m int) int {
    if n > m {
        return n
    }
    return m
}