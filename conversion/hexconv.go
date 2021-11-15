import "strconv"
import "fmt"


// toHex conversion manually done. Twos complement for negatives.

func toHex(num int) string {
    // num is coming in as base10
    
    // check if num is 0
    if num == 0 {
		return "0"
	}
    
    // check if negative for twos complement
    if num < 0 {
        num += 1 << 32
    }
        
    // hexadecimal string
    hexa := ""
    
    // cant use lib, so convert to base16
    // loop over num, as it continues to decrease
    for num > 0 {
        // calculate remainder and use it as position
        remainder := num % 16
        // divide value by 16, to get to the next position
        num = num / 16
        // add the remainder to the hex string
        hexa = convertNum(remainder) + hexa 
    }
    
    
    return hexa
}

// basic conversion map
func convertNum(num int) string {
    switch num {
    case 10:
        return "a"
    case 11:
        return "b"
    case 12:
        return "c"
    case 13:
        return "d"
    case 14:
        return "e"
    case 15:
        return "f"
    default:
        return strconv.Itoa(num)
    }
        
}