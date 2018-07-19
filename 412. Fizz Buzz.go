// Write a program that outputs the string representation of numbers from 1 to n.
// But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. 
// For numbers which are multiples of both three and five output “FizzBuzz”.

func fizzBuzz(n int) []string {
    arr := []string{}
    for num := 1; num <= n; num++ {
        if num%3 == 0 && num%5 == 0 {
            arr = append(arr, "FizzBuzz")
        } else if num%5 == 0 {
            arr = append(arr,"Buzz")
        } else if num%3 == 0 {
            arr = append(arr,"Fizz")
        } else {
            arr = append(arr, strconv.Itoa(num))
        }
    }
    return arr
}
