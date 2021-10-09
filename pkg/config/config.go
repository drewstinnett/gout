package config

// Config Structure to pass to formatters.  Should include enough config to do
// the output. You must set the Format here to something like yaml, json,
// plain, or any other value returned by the GetFormats function
type Config struct {
	Format      string
	LimitFields []string
	Template    string
}
