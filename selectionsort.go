package main

import "fmt"

// sort function that sorts an array of integers
func sort(a []int, n int) {
    var current, j, lowestindex, temp int

    for current = 0; current < n-1; current++ {
        // Each time through this loop, scan the array
        // from current+1 to the end. If we find
        // something lower than what is at current, then
        // swap it with current index. So each time
        // through this loop, a[current] will be
        // properly sorted.

        // 1) first find the index of the lowest value
        lowestindex = current

        for j = current + 1; j < n; j++ {
            if a[j] < a[current] {
                lowestindex = j
            }
        }

        // 2) now swap a[current] and a[lowestindex],
        // as long as a difference was found.
        if lowestindex != current {
            temp = a[current]
            a[current] = a[lowestindex]
            a[lowestindex] = temp
        }
    }
}
