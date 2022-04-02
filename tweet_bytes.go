package main

import (
	"fmt"
	"strings"
)

func extractValues(bytes []byte, tweet Tweet) Tweet {
	jsonString := string(bytes)

	s1 := strings.ReplaceAll(jsonString, "[{", "\n")
	s2 := strings.ReplaceAll(s1, "]}", "\n")

	strs := strings.Split(s2, "\n")

	// Prevents processing Tweets without enough
	// data to populate tweet
	if len(strs) < 8 {
		tweet.Error = true
		return tweet
	}

	tweet = extractDataSection(strs[1], tweet)

	if !tweet.Error {
		tweet = extractMediaUrl(strs[2], tweet)
	}

	if !tweet.Error {
		tweet = extractUsers(strs[3], tweet)
	}

	return tweet
}

func extractDataSection(str string, tweet Tweet) Tweet {
	strs := formatEscapedCharacters(str)

	// Prevents processing Tweets without enough
	// data to populate tweet
	if len(strs) < 7 {
		tweet.Error = true
		return tweet
	}

	tmp := strs[7][13:]

	size := strings.Split(tmp, " https://")[0]

	tweet.AuthorId = strs[1]
	tweet.Id = strs[3]
	tweet.Text = size

	return tweet

}

func extractMediaUrl(str string, tweet Tweet) Tweet {
	strs := formatEscapedCharacters(str)

	// Prevents processing Tweets without enough
	// data to populate tweet
	if len(strs) < 6 {
		tweet.Error = true
		return tweet
	}

	tweet.MediaUrl = strs[5]

	return tweet
}

func extractUsers(str string, tweet Tweet) Tweet {
	strs := formatEscapedCharacters(str)

	// Prevents processing Tweets without enough
	// data to populate tweet
	if len(strs) < 7 {
		tweet.Error = true
		return tweet
	}

	tweet.AuthorName = strs[3]
	tweet.AuthorScreenName = strs[5]

	return tweet
}

func formatEscapedCharacters(str string) []string {
	escapeQuotes1 := ","
	s1 := strings.ReplaceAll(str, fmt.Sprintf("%q", escapeQuotes1), "%")

	escapeQuotes2 := ":"
	s2 := strings.ReplaceAll(s1, fmt.Sprintf("%q", escapeQuotes2), "%")

	escapeQuotes3 := "}],"
	s3 := strings.ReplaceAll(s2, fmt.Sprintf("%q", escapeQuotes3), "%")

	escapeQuotes4 := "},{"
	s4 := strings.ReplaceAll(s3, fmt.Sprintf("%q", escapeQuotes4), "%")

	strs := strings.Split(s4, "%")

	return strs
}
