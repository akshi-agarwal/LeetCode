// Write a function that takes a string as input and returns the string reversed.
// Example:
// Given s = "hello", return "olleh".

func reverseString(s string) string {
    reverse := ""
    for _,ch := range s {
        reverse = string(ch) + reverse
    }
    return reverse
}
