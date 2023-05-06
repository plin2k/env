package env

import (
	"sync"
)

var g *Env
var gOnceInit sync.Once

func init() {
	gOnceInit.Do(func() {
		g = New()
	})
}

// Clear destroy global Env
func Clear() {
	g = nil
	gOnceInit = sync.Once{}
}

// SetDefault set default value for key or replace existed for global envi
func SetDefault(key string, value interface{}) {
	g.SetDefault(key, value)
}

// IsSet check if env with key is existed in global envi
func IsSet(key string) bool {
	return g.IsSet(key)
}

// GetString get string env from global envi
func GetString(key string) string {
	return g.GetString(key)
}

// GetInt get int env from global envi
func GetInt(key string) int {
	return g.GetInt(key)
}

// GetInt64 get int64 env from global envi
func GetInt64(key string) int64 {
	return g.GetInt64(key)
}

// GetFloat64 get float64 env from global envi
func GetFloat64(key string) float64 {
	return g.GetFloat64(key)
}

// GetBool get bool env from global envi
func GetBool(key string) bool {
	return g.GetBool(key)
}
