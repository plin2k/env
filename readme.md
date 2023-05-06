# Convenient work with environment variables

## Example Usage

### Global
```go
env.SetDefault("fooString", "bar")
env.SetDefault("fooInt", 825)
env.SetDefault("fooInt64", 125)
env.SetDefault("fooFloat64", 11.22)
env.SetDefault("fooBool", true)

if env.IsSet("fooNotExists") {
	return errors.New("environment variable not exists")
}

env.GetString("fooString")
env.GetInt("fooInt")
env.GetInt64("fooInt64")
env.GetFloat64("fooFloat64")
env.GetBool("fooBool")
```

### Custom
```go
myEnv := env.New()

myEnv.SetDefault("fooString", "bar")
myEnv.SetDefault("fooInt", 825)
myEnv.SetDefault("fooInt64", 125)
myEnv.SetDefault("fooFloat64", 11.22)
myEnv.SetDefault("fooBool", true)

if myEnv.IsSet("fooNotExists") {
	return errors.New("environment variable not exists")
}

myEnv.GetString("fooString")
myEnv.GetInt("fooInt")
myEnv.GetInt64("fooInt64")
myEnv.GetFloat64("fooFloat64")
myEnv.GetBool("fooBool")
```