package storage

var urlStore = make(map[string]string)

func SaveURL(shortCode, originalURL string) {
    urlStore[shortCode] = originalURL
}

func GetURL(shortCode string) (string, bool) {
    url, exists := urlStore[shortCode]
    return url, exists
}
