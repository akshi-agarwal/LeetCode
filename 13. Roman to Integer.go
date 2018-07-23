// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

func romanToInt(s string) int {
	add := 0
	
    for i := 0; i < len(s); i++ {

		fmt.Printf("Entered Loop: add: %v\n\n\n", add)
			if i + 1 < len(s) && s[i] == 'I' && s[i+1] == 'V' {
				fmt.Printf("Entered I V\n\n")
				add += 4
				i += 1
			} else if i + 1 < len(s) && s[i] == 'I' && s[i+1] == 'X' {
				add =+ 9
				i += 1
			} else if i + 1 < len(s) && s[i] == 'X' && s[i+1] == 'L' {
				add += 40
				i += 1
			} else if i + 1 < len(s) && s[i] == 'X' && s[i+1] == 'C' {
				add += 90
				i += 1
			} else if i + 1 < len(s) && s[i] == 'C' && s[i+1] == 'D' {
				add += 400
				i += 1
			} else if i + 1 < len(s) && s[i] == 'C' && s[i+1] == 'M' {
				add += 900
				i += 1
			} else if s[i] == 'I' {
				add += 1
			} else if s[i] == 'V' {
				add += 5
			} else if s[i] == 'X' {
				fmt.Printf("Entered X\n\n")
				add += 10
			} else if s[i] == 'L' {
				add += 50
			} else if s[i] == 'C' {
				fmt.Printf("Entered C\n\n")
				add += 100
			} else if s[i] == 'D' {
				add += 500
			} else if s[i] == 'M' {
				fmt.Printf("Entered M\n\n")
				add += 1000
			}
    }
    return add   
}
