package config

import "strings"

type defaultConfig struct {
	data map[string]any
}

// GetBool implements [configuration].
func (c *defaultConfig) GetBool(key string) (bool, bool) {
	val, ok := c.get(key)
	if ok {
		return val.(bool), true
	}
	return false, false
}

// GetBoolDefault implements [configuration].
func (c *defaultConfig) GetBoolDefault(key string, dv bool) bool {
	val, ok := c.GetBool(key)
	if !ok {
		return dv
	}
	return val
}

// GetFloat implements [configuration].
func (c *defaultConfig) GetFloat(key string) (float64, bool) {
	val, ok := c.get(key)
	if ok {
		return val.(float64), true
	}
	return 0, false
}

// GetFloatDefault implements [configuration].
func (c *defaultConfig) GetFloatDefault(key string, dv float64) float64 {
	val, ok := c.GetFloat(key)
	if !ok {
		return dv
	}
	return val
}

// GetInt implements [configuration].
func (c *defaultConfig) GetInt(key string) (int, bool) {
	val, ok := c.get(key)
	if ok {
		return val.(int), true
	}
	return 0, false
}

// GetIntDefault implements [configuration].
func (c *defaultConfig) GetIntDefault(key string, dv int) int {
	val, ok := c.GetInt(key)
	if !ok {
		return dv
	}
	return val
}

// GetString implements [configuration].
func (c *defaultConfig) GetString(key string) (string, bool) {
	val, ok := c.get(key)
	if ok {
		return val.(string), true
	}
	return "", false
}

// GetStringDefault implements [configuration].
func (c *defaultConfig) GetStringDefault(key string, dv string) string {
	val, ok := c.GetString(key)
	if !ok {
		return dv
	}
	return val
}

func (c *defaultConfig) get(path string) (any, bool) {
	data, found := c.data, false
	var val any
	for key := range strings.SplitSeq(path, ":") {
		val, found = data[key]
		if newData, ok := val.(map[string]any); ok && found {
			data = newData
		} else {
			return val, found
		}
	}
	return val, found
}

func (c *defaultConfig) GetSection(path string) (Config, bool) {
	val, found := c.get(path)
	if found {
		if data, ok := val.(map[string]any); ok {
			return &defaultConfig{data: data}, true
		}
	}
	return nil, false

}
