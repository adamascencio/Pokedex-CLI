package pokecache

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/adamascencio/pokedexcli/internal/utils"
)

type cacheEntry struct {
	CreatedAt time.Time `json:"created_at"`
	Val       []byte    `json:"val"`
}

type Cache struct {
	dir      string
	interval time.Duration
}

func hashURL(rawURL string) string {
	sum := sha256.Sum256([]byte(rawURL))
	return hex.EncodeToString(sum[:])
}

// Delete removes a file from the cache, if present.
func (c *Cache) delete(k string) error {
	urlHash := hashURL(k)
	filePath := filepath.Join(c.dir, urlHash)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

// Creates a disk cache.
func NewCache(interval time.Duration) *Cache {
	dir, _ := os.UserCacheDir()
	dir = filepath.Join(dir, "pokedexcli")
	os.MkdirAll(dir, 0755)
	c := &Cache{dir: dir, interval: interval}
	return c
}

// Add stores a cache entry with the provided key and value.
func (c *Cache) Add(k string, v []byte) {
	entry := cacheEntry{
		CreatedAt: time.Now(),
		Val:       v,
	}
	urlHash := hashURL(k)
	dir := filepath.Join(c.dir, urlHash)
	data, _ := json.Marshal(entry)
	err := utils.SaveData(dir, data)
	if err != nil {
		log.Printf("cache add failed for %q: %v", k, err)
		return
	}
}

// Get retrieves a cached value for the key, if present.
func (c *Cache) Get(k string) ([]byte, bool) {
	urlHash := hashURL(k)
	var cacheData cacheEntry
	filePath := filepath.Join(c.dir, urlHash)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, false
	}
	if err := json.Unmarshal(data, &cacheData); err != nil {
		log.Printf("cache decode failed for %q: %v", k, err)
		_ = c.delete(k)
		return nil, false
	}
	timeElapsed := time.Since(cacheData.CreatedAt)
	if timeElapsed > c.interval {
		err := c.delete(k)
		if err != nil {
			log.Printf("cache delete failed for %q: %v", k, err)
		}
		return nil, false
	}
	return cacheData.Val, true
}
