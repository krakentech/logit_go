<picture>
    <source media="(prefers-color-scheme: dark)" srcset="./res/github-topper-dark.jpg" />
    <source media="(prefers-color-scheme: light)" srcset="./res/github-topper-light.jpg" />
    <img src="./res/github-topper-light.jpg" />
</picture>

![Static Badge](https://img.shields.io/badge/Release-v0.0.0-green)

## About

Elevate your application’s logging with this powerful library. Effortlessly configure log destinations and ensure consistent, standardized log formats across your project. Unlock advanced features like performance tracking and dataset logging to streamline debugging and gain deeper insights into your application’s behavior.

## Install

```shell
go get -u github.com/krakentech/logit_go
```

## Usage

### Logging Functions

With no configuration this will be the default output

#### Debug

example:

```go
logit.Debug("This is a debug message %d", 42)
```

response:

```shell
25.09.27-16:43:20 🐛 - This is a debug message: 42
```

#### Info

example:
```go
logit.Info("This is an info message %d", 42)
```

response:

```shell
25.09.27-16:43:20 🧠 - This is an info message: 42
```

#### Warning

example:
```go
logit.Warn("This is a warning message %d", 42)
```

response:

```shell
25.09.27-16:43:20 ⚠️ - This is a warning message: 42
```

#### Error

example:
```go
logit.Error("This is an error message %d", 42)
```

response:

```shell
25.09.27-16:43:20 💥 - This is an error message: 42
```

#### Err

Instead of using the error in as a formatting property you can pass it in and we will append the error message to the end of the log message.

example:
```go
err := errors.New("this is an error")
logit.Err(err, "This is an error message %d", 42)
```

response:

```shell
25.09.27-16:43:20 💥 - This is an error message 42: this is an error
```

#### Debug Data

With format false

example:
```go
data := []String{"A", "B"}
logit.DebugData(data, false, "This is a debug message with data")
```                   

response:

```shell
25.09.27-16:43:20 ✨ - --> This is a debug message with data
25.09.27-16:43:20 ✨ - ["A", "B"]
```

With format true

example:
```go
data := []String{"A", "B"}
logit.DebugData(data, true, "This is a debug message with data")
```                   

response:

```shell
25.09.27-16:43:20 ✨ - --> This is a debug message with data
25.09.27-16:43:20 ✨ - [
25.09.27-16:43:20 ✨ -   "A", 
25.09.27-16:43:20 ✨ -   "B"
25.09.27-16:43:20 ✨ - ]
```

### Tracker

You can use the tracker to track the time taken for a function to execute. You can also add custom fields to the tracker.

example:
```go
func myFunction() {
    tracker := logit.NewTracker("myFunction")
    defer tracker.Log()         
}
```
response:

```shell
25.09.27-16:43:20 ⏱️ - myFunction [218582 ns] 218.582µs
```


## Licence

MIT
