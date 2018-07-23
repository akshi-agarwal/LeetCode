// Given a 32-bit signed integer, reverse digits of an integer.

func reverse(x int) int {
    rev := 0
    
    // checking if int is 32-bit
    if math.MaxInt32 < x || math.MinInt32 > x {
        return 0;
    }
    
    for x != 0 {
        add := x % 10
          x  = x / 10
        rev  = rev * 10
        rev  = rev + add
    }
    
    // checking if reverse int is 32-bit
    if math.MaxInt32 < rev || math.MinInt32 > rev {
        return 0;
    }
    
    return rev
}
