package main

func Sum(numbers [5]int) int {
    sum := 0
    /*for i := 0; i < 5; i++ {
        sum += numbers[i]
    }
    return sum */

    for _, num := range numbers {
        sum += num
    }
    return sum
}

func SumWithSlice(numbers []int) int {
    sum := 0

    for _, num := range numbers {
        sum += num
    }
    return sum
}

func SumAll(numbersToSum ...[]int) []int {
    /*lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numbersToSum {
        sums[i] = SumWithSlice(numbers)
    }
    return sums */

    var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, SumWithSlice(numbers))
    }
    return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int

	for _, numbers := range numbersToSum {

        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            // NOTE: Slices can be sliced! The syntax is slice[low:high].
            tail := numbers[1:]
            sums = append(sums, SumWithSlice(tail))
        }

	}
    return sums
}
