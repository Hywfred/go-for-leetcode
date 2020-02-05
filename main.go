package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/longestPalindrome"
)

func main() {
	in := []string{
		"azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc",
		"abcda",
		"bb",
		"abccba",
		"abcba",
		"ac",
		"babad",
		"cbbd",
		"",
		"a",
		"abcccc",
		"cccccab",
	}
	for _, v := range in {
		fmt.Println(longestPalindrome.LongestPalindrome(v))
	}
}
