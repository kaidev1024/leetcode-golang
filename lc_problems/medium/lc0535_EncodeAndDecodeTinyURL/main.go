type Codec struct {
	used map[string]struct{}
	m    map[string]string
}

func Constructor() Codec {
	return Codec{
		used: make(map[string]struct{}),
		m:    make(map[string]string),
	}
}

func genShort() string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	ret := make([]byte, 6)
	for i := 0; i < 6; i++ {
		ind := rand.Intn(62)
		ret[i] = str[ind]
	}
	return string(ret)
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	shortUrl := ""
	for {
		str := genShort()
		if _, ok := this.used[str]; ok {
			continue
		} else {
			this.used[str] = struct{}{}
			shortUrl = str
			break
		}
	}
	this.m[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
