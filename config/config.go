package config

type Config interface {
	GetString(key string) (val string, ok bool)
	GetInt(key string) (val int, ok bool)
	GetBool(key string) (val bool, ok bool)
	GetFloat(key string) (val float64, ok bool)

	GetStringDefault(key string, dv string) string
	GetIntDefault(key string, dv int) int
	GetBoolDefault(key string, dv bool) bool
	GetFloatDefault(key string, dv float64) float64

	GetSection(key string) (val Config, ok bool)
}
