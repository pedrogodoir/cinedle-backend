package cache_movie

import (
	"cinedle-backend/internal/modules/movies/models"
	"cinedle-backend/internal/utils/cache"
	"time"
)

type MovieCache struct {
	cache *cache.Cache[int, models.MovieRes]
}

func NewCacheMovie() *MovieCache {
	// Set default TTL for movie cache items
	const defaultTTL time.Duration = 10 * time.Minute
	return &MovieCache{
		cache: cache.NewCache[int, models.MovieRes](defaultTTL),
	}
}

func (mc *MovieCache) GetMovieCache(id int) (models.MovieRes, bool) {
	movie, ok := mc.cache.Get(id)
	return movie, ok
}

func (mc *MovieCache) SetMovieCache(id int,
	value models.MovieRes) {
	mc.cache.Set(id, value)
}

func (mc *MovieCache) Clear() {
	mc.cache.Clear()
}
