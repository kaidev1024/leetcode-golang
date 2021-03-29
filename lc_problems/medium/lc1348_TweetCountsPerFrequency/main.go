type TweetCounts struct {
	counts map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		counts: make(map[string][]int),
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.counts[tweetName] = append(this.counts[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	interval := getInterval(freq)
	n := (endTime - startTime + interval) / interval
	result := make([]int, n)
	records := this.counts[tweetName]
	for _, v := range records {
		if v >= startTime && v <= endTime {
			result[(v-startTime)/interval]++
		}
	}
	return result
}

func getInterval(freq string) int {
	if freq == "minute" {
		return 60
	}
	if freq == "hour" {
		return 3600
	}
	return 86400
}
