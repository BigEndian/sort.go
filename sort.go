package main

import (
	"fmt"
	"rand"
	"time"
	"flag"
	"os"
	"math"
)

func bubbleSort(array []int) {
	var temp int
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			temp = array[j]
			array[j] = array[j-1]
			array[j-1] = temp
		}
	}
}

func mergeSort(array []int) []int {
	splitPoint := len(array)/2

	if (len(array) >= 2) {
		left := mergeSort(array[splitPoint:])
		right := mergeSort(array[:splitPoint])
		return combine(left, right)
	} else {
		return array
	}
	return array
}
func combine(left []int, right []int) []int {
	var output []int = make([]int, len(left)+len(right))
	var lindex, rindex, oindex int
	lindex = 0
	rindex = 0
	oindex = 0

	for ; lindex != len(left) && rindex != len(right); {
		if (left[lindex] < right[rindex]) {
			output[oindex] = left[lindex]
			oindex++
			lindex++
		} else {
			output[oindex] = right[rindex]
			oindex++
			rindex++
		}
	}

	for ; lindex < len(left); lindex++ {
		output[oindex] = left[lindex]
		oindex++
	}

	for ; rindex < len(right); rindex++ {
		output[oindex] = right[rindex]
		oindex++
	}

	return output
}
func max(integers... int) int {
	var max int = integers[0]
	for _, i := range integers[1:] {
		if (i > max) {
			max = i
		}
	}
	return max
}
func radixSort(array *[]int) {
	var max_int int = max(*array...)
	var max_log int = int(math.Log10(float64(max_int)))
	fmt.Printf("Max log for array is %d\n", max_log)

	var buckets [][]int = make([][]int, 10)
	var result []int = make([]int, len(*array))
	copy(result, *array)

	for digit_index := 0; digit_index <= max_log; digit_index++ {
		for i := 0; i < len(result); i++ {
			digit := isolateDigit(result[i], digit_index)
			buckets[digit] = append(buckets[digit], result[i])
		}
		result = []int{}
		for j := 0; j < len(buckets); j++ {
			fmt.Printf("Values for buckets[%d] is %v\n", j, buckets[j])
			result = append(result, buckets[j]...)
			buckets[j] = []int{} // Reset
		}
	}
	copy(*array, result)
}

func isolateDigit(number int, digit int) int { // 100, 0 is 0, 100, 1 is 0, 100, 2 is 1 the floored log of 100..999 is 2
	var log_value int = int( math.Floor(math.Log10(float64(number))) )
	// 300 digit 2
	// Log is 2
	if log_value < digit { // its log is 1, i.e., digit 2 on 99 (digit 1 is 9, digit 0 is 9, no digit 2)
		return 0
	}

	number %= int(math.Pow(10.0, float64(digit+1))) // get digit 1: 10 has a log value of 1, 10 %= (10**2) -> 10
	number /= int(math.Pow(10.0, float64(digit)))   // now it's 10: 10 /= 10**1 -> 1
	return number

}

func genRandomArray(random *rand.Rand, options GenerationOptions) []int {

	var array []int = make([]int, options.length)
	if options.reverse {
		if options.rand_max > 0 {
			var curr_max int = options.rand_max
			for i := 0; i < len(array); i++ {
				array[i] = random.Intn(curr_max)
				curr_max -= 1
			}
		} else {
			for i := 0; i < len(array); i++ {
				array[i] = len(array) - i
			}
		}
	} else /* not reverse */ {
		if options.rand_max > 0 {
			for i := 0; i < len(array); i++ {
				array[i] = random.Intn(options.rand_max)
			}
		} else {
			for i := 0; i < len(array); i++ {
				array[i] = random.Int()
			}
		}
	}
	return array
}


var length *int = flag.Int("size", 10000, "the length of the array to sort")
var sort_type *string = flag.String("method", "bubble", "the sort method to use [bubble, merge, radix]")
var reverse *bool = flag.Bool("reverse", false, "Instead of an array of pseudorandom integers, populate the array by numbers in a reverse-sequential order")
var rand_max *int = flag.Int("limit", 0, "The highest value for an integer to be generated pseudorandomly and placed into the array. If 0, it'll be INT_MAX")
var print_initial *bool = flag.Bool("printinit", false, "Print the initial values of the array before sorting the array")

type GenerationOptions struct {
	length int
	reverse bool
	rand_max int
}

func main() {
	flag.Parse()
	var options GenerationOptions = GenerationOptions{*length, *reverse, *rand_max}
	var r *rand.Rand = rand.New(rand.NewSource(time.Nanoseconds()))
	var array []int = genRandomArray(r, options)
	if *print_initial {
		for i := 0; i < *length; i++ {
			fmt.Printf("%d ", array[i])
		}
		fmt.Printf("\n")
	}
	switch *sort_type {
	case "bubble":
		bubbleSort(array)
	case "merge":
		array = mergeSort(array)
	case "radix":
		/*for i := 0; i < 100; i++ {
			fmt.Printf("%d: %d %d %d\n", i, isolateDigit(i, 2), isolateDigit(i, 1), isolateDigit(i, 0))
		}*/
		radixSort(&array)
	default:
		fmt.Fprintf(os.Stderr, "Invalid sorting method: %s\n", *sort_type)
		os.Exit(1)
	}

	for i := 0; i < *length; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Printf("\n")
}
