package env

import (
	"os"
	"strconv"
)

// New create instance of Env
func New() *Env {
	return &Env{
		defaults: make(map[string]interface{}, 10),
	}
}

// Env envi struct with defaults
type Env struct {
	defaults map[string]interface{}
}

// SetDefault set default value for key or replace existed
func (e *Env) SetDefault(key string, value interface{}) {
	e.defaults[key] = value
}

// IsSet check if env with key is existed
func (e *Env) IsSet(key string) bool {
	_, isSet := os.LookupEnv(key)
	if !isSet {
		_, isSet = e.defaults[key]
	}

	return isSet
}

// GetString get string env
func (e *Env) GetString(key string) string {
	if envVal := os.Getenv(key); envVal != "" {
		return envVal
	}

	if d, ok := e.defaults[key]; ok {
		if dTyped, ok := d.(string); ok {
			return dTyped
		}
	}

	return *new(string)
}

// GetInt get int env
func (e *Env) GetInt(key string) int {
	if envVal := os.Getenv(key); envVal != "" {
		v, err := strconv.ParseInt(envVal, 0, 0)
		if err == nil {
			return int(v)
		}
	}

	if d, ok := e.defaults[key]; ok {
		if dTyped, ok := d.(int); ok {
			return dTyped
		}
	}

	return *new(int)
}

// GetInt64 get int64 env
func (e *Env) GetInt64(key string) int64 {
	if envVal := os.Getenv(key); envVal != "" {
		v, err := strconv.ParseInt(envVal, 0, 64)
		if err == nil {
			return v
		}
	}

	if d, ok := e.defaults[key]; ok {
		if dTyped, ok := d.(int64); ok {
			return dTyped
		}
	}

	return *new(int64)
}

// GetFloat64 get float64 env
func (e *Env) GetFloat64(key string) float64 {
	if envVal := os.Getenv(key); envVal != "" {
		v, err := strconv.ParseFloat(envVal, 64)
		if err == nil {
			return v
		}
	}

	if d, ok := e.defaults[key]; ok {
		if dTyped, ok := d.(float64); ok {
			return dTyped
		}
	}

	return *new(float64)
}

// GetBool get bool env
func (e *Env) GetBool(key string) bool {
	if envVal := os.Getenv(key); envVal != "" {
		v, err := strconv.ParseBool(envVal)
		if err == nil {
			return v
		}
	}

	if d, ok := e.defaults[key]; ok {
		if dTyped, ok := d.(bool); ok {
			return dTyped
		}
	}

	return *new(bool)
}
