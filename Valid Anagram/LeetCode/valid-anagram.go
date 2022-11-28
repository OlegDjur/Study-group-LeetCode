// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:

// Input: s = "rat", t = "car"
// Output: false

// Constraints:

//     1 <= s.length, t.length <= 5 * 104
//     s and t consist of lowercase English letters.

// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

package leetcode

var ch1 = make(chan map[rune]int, 2)

func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
	
	go buildMap(s)
	go buildMap(t)

	s1 := <-ch1
	s2 := <-ch1

	for i, v := range s1 {
		if v != s2[i] {
			return false
		} 
	}
	return true
}

func buildMap(s string){
	m := make(map[rune]int)

    for _, v := range s{
		
			m[v] += 1
	}
	ch1<- m
}