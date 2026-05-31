// stats/main.go

// Read a text file of real numbers, one per line, and calculate their min,
// max, median, sum, average, and standard deviation (populations version, not
// the sample version).

package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
)

func main() {
    // open the file
    f, err := os.Open("numbers.txt")
    if err != nil {
        panic("Couldn't open file!")
    }
    defer f.Close()

    // read the numbers int an array
    var nums []float64
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        s := scanner.Text()
        x, err := strconv.ParseFloat(s, 64)
        if err != nil {
            panic("Error reading file")
        }
        nums = append(nums, x)
    }

    sort.Float64s(nums)
    fmt.Println(nums)

    // print statistics
    fmt.Println("      Min:", min(nums))
    fmt.Println("   Median:", nums[len(nums)/2])
    fmt.Println("      Max:", nums[len(nums)-1])
    fmt.Println("      Sum:", sum(nums))
    fmt.Println("     Mean:", mean(nums))
    fmt.Println("Std. dev.:", std_dev(nums))
} // main

func sum(nums []float64) (result float64) {
    for _, x := range nums {
        result += x
    }
    return
}

func mean(nums []float64) float64 {
    return sum(nums) / float64(len(nums))
}

func min(nums []float64) float64 {
    result := nums[0]
    for _, x := range nums {
        if x < result {
            result = x
        }
    }
    return result
}

func max(nums []float64) float64 {
    result := nums[0]
    for _, x := range nums {
        if x > result {
            result = x
        }
    }
    return result
}

func median(nums []float64) float64 {
    sort.Float64s(nums)
    return nums[len(nums) / 2]
}

func std_dev(nums []float64) (result float64) {
    avg := mean(nums)
    for _, x := range nums {
        diff := x - avg
        result += diff * diff
    }
    return math.Sqrt(result)
}

// def std_dev(nums)
//  n = nums.length
//  mean = nums.sum / n
//  return Math.sqrt(nums.each {|x| (x - mean)**2} .sum / n)
// end

// // based on the code given on this page: 
// // https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
// //
// // ReadInts reads whitespace-separated ints from r. If there's an error, it
// // returns the ints successfully read so far as well as the error value.
// func ReadInts(r io.Reader) ([]int, error) {
//     scanner := bufio.NewScanner(r)
//     scanner.Split(bufio.ScanWords)
//     var result []int
//     for scanner.Scan() {
//         x, err := strconv.Atoi(scanner.Text())
//         if err != nil {
//             return result, err
//         }
//         result = append(result, x)
//     }
//     return result, scanner.Err()
// }

// def std_dev(nums)
//  n = nums.length
//  mean = nums.sum / n
//  return Math.sqrt(nums.each {|x| (x - mean)**2} .sum / n)
// end

// nums = File.open("numbers.txt").readlines.map {|s| s.to_f}

// nums.sort!
// puts nums.to_s
// puts "        Min: #{nums[0]}"
// puts "     Median: #{nums[nums.length / 2]}"
// puts "        Max: #{nums[-1]}"
// puts "        Sum: #{nums.sum}"
// puts "       Mean: #{nums.sum / nums.length}"
// puts "  Std. dev.: #{std_dev(nums)}"
