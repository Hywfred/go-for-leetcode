package longestPalindrome_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/longestPalindrome"
)

func TestLongestPalindrome(t *testing.T) {
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
	wants := []string{
		"sooos",
		"a",
		"bb",
		"abccba",
		"abcba",
		"a",
		"bab",
		"bb",
		"",
		"a",
		"cccc",
		"ccccc",
	}
	for i, v := range wants {
		result := longestPalindrome.LongestPalindrome(in[i])
		if result != v {
			t.Errorf("%v\t%v\t%v\n", in[i], result, v)
		}
	}
}
