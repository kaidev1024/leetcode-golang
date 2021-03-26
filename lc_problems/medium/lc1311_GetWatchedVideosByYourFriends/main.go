type video struct {
	freq int
	name string
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	visited := make([]bool, 100)
	currentFriends := []int{id}
	for i := 1; i <= level; i++ {
		if len(currentFriends) == 0 {
			return []string{}
		}
		nextFriends := []int{}
		for _, fid1 := range currentFriends {
			visited[fid1] = true
			for _, fid2 := range friends[fid1] {
				if !visited[fid2] {
					visited[fid2] = true
					nextFriends = append(nextFriends, fid2)
				}
			}
		}
		currentFriends = nextFriends
	}
	hash := map[string]int{}
	for _, fid := range currentFriends {
		for _, v := range watchedVideos[fid] {
			hash[v]++
		}
	}
	videos := []video{}
	for k, v := range hash {
		videos = append(videos, video{
			freq: v,
			name: k,
		})
	}
	sort.Slice(videos, func(i, j int) bool {
		if videos[i].freq == videos[j].freq {
			return videos[i].name < videos[j].name
		}
		return videos[i].freq < videos[j].freq
	})
	result := []string{}
	for _, v := range videos {
		result = append(result, v.name)
	}
	return result
}