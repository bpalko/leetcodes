/*
Problem: Find the First Non-Repeating Character

Given a string s, find the first non-repeating character and return its index. If all characters are repeating, return -1.
Examples

    Input: "leetcode"
    Output: 0 (The first non-repeating character is 'l')

    Input: "loveleetcode"
    Output: 2 (The first non-repeating character is 'v')

    Input: "aabb"
    Output: -1 (No non-repeating character)

Hints

    Use a hashmap to store the frequency of each character.
    Iterate through the string to find the first character with a frequency of 1
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
    mapChars := make(map[string]int)
    input := "loveleetcode"
    spInput := strings.Split(input, "")
    for _, val := range spInput {
        mapChars[val] = 0
    }
	ind, val := nonrepeat(spInput, mapChars)
    fmt.Printf("ind: %v\n", ind)
    fmt.Printf("val: %v\n", val)
}


func nonrepeat(input []string, chars map[string]int) (int, string) {
    for _, val := range input {
        chars[val] = chars[val] + 1
    }
    min := 1
    fmt.Printf("chars: %v\n", chars)
    for ind, val := range input {
        if chars[val] <= min {
            return ind, val
        }
    }
    return -1, ""
}
