package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
    numbers := [5]int{1,2,3,4,5}

    got := Sum(numbers)
    want := 15

    if got != want {
        t.Errorf("\nGOT = %d\nWANT = %d\nGiven = %v", got, want, numbers)
    }


}



/*  

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. 
Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. 
NOTE: Every test has a cost.

*/

func TestSum2(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T){
        numbers := []int{1,2,3,4,5}

        got := SumWithSlice(numbers)
        want := 15

        if got != want {
            t.Errorf("\nGOT = %d\nWANT = %d\nGiven = %v", got, want, numbers)
        }
    })


    t.Run("collection of any size", func(t *testing.T){
        numbers := []int{1,2,3}

        got := SumWithSlice(numbers)
        want := 6

        if got != want {
            t.Errorf("\nGOT = %d\nWANT = %d\nGiven = %v", got, want, numbers)
        }
    })
}

/*

Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. 
If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.

HACK: For this, run: go test -cover

*/


func TestSumAll(t *testing.T) {
    got := SumAll([]int{1,2}, []int{0,9})
    want:= []int{3, 9}

    // NOTE: It's important to note that reflect.DeepEqual is not "type safe" - the code will compile even if you did something a bit silly. To see this in action, temporarily change the test to:
    // want:= "bob"


    /* if got != want {
        t.Errorf("\nGOT = %v\nWANT = %v\n", got, want)
    } */

    // NOTE:  Go does not let you use equality operators with slices. You could write a function to iterate over each got and want slice and check their values but for convenience sake, we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.
    if !reflect.DeepEqual(got, want) {
        t.Errorf("\nGOT = %v\nWANT = %v\n", got, want)
    }
}



