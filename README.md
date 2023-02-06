# Pandati

  [![Docker image build.](https://github.com/lukaszraczylo/pandati/actions/workflows/test-and-release.yaml/badge.svg)](https://github.com/lukaszraczylo/pandati/actions/workflows/test-and-release.yaml) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/lukaszraczylo/pandati) [![Go Reference](https://pkg.go.dev/badge/github.com/lukaszraczylo/pandati.svg)](https://pkg.go.dev/github.com/lukaszraczylo/pandati)
## The one stop shop for most common Go functions

<p align="center">
  <img height="300" src="static/pandati.jpg">
</p>

## Table of contents

- [Pandati](#pandati)
  - [The one stop shop for most common Go functions](#the-one-stop-shop-for-most-common-go-functions)
  - [Table of contents](#table-of-contents)
    - [Purpose of the project](#purpose-of-the-project)
    - [Available helpers](#available-helpers)
      - [Logging](#logging)
    - [Contributing](#contributing)
    - [Name of the project](#name-of-the-project)

### Purpose of the project

Keeping things DRY for your projects. Set of most popular and frequently used helper functions making it easier to focus on the logic you want to implement rather than re-inventing the wheel ( and spending precious time going through StackOverflow ).


### Available helpers

| Class | Function | Description |
| :---  | :---  | :--- |
| Errors | `CheckForError(err error, msg ...string) (string)` | Checks for error and returns pre-formatted message with stack trace |
| Errors | `Trace() string` | Returns a formatted stack trace |
|  |  |  |
| Maps   | `FlattenMap(nested map[string]interface{}, opts *FlattenOptions) (m map[string]interface{}, err error)` | Flattens provided map using customisation options |
| Maps   | `UnflattenMap(flat map[string]interface{}, opts *FlattenOptions) (nested map[string]interface{}, err error)` | Unflattens provided map into nested map using customisation options |
|  |  |  |
| Slices | `ExistsInSlice(slice interface{}, value interface{}) bool` | Checks if value exists in slice |
| Slices | `RemoveFromSlice(slice interface{}, value interface{})` | Removes value from slice |
| Slices | `RemoveFromSliceByIndex(slice interface{}, index int)` | Removes value from slice by index |
| Slices | `UniqueSlice (slice interface{}) (uniqueSlice interface{})` | Returns a slice with unique values |
|  |  |  |
| Structs | `CompareStructsReplacedFields (old interface{}, new interface{}) (changedFields []string)` | Compares two structs and returns a slice of changed fields |
|  |  |  |
| Checks | `IsZero(v interface{}) bool` | Checks if value of anything passed is zero / empty / nil |

#### Logging

Additional helper for logger setup is available. I've decided to publish it to DRY my own projects.
The logging itself provides nice, clean and fully customizable json formatted logging which then can be ingested without any further modifications straight into your loki, ELK or any other logging mechanism. It can be used as following.

```go
import (
    "github.com/lukaszraczylo/pandati"
)

func main() {
    log = logging.NewLogger()
    log.Critical("Error binding to queue", map[string]interface{}{"_queue": queueName, "_error": err.Error()})
}```

### Contributing

If you have any suggestions or want to contribute to the project, please open an issue or create a pull request.
Pull requests should contain test cases and documentation describing the added functionality.

### Name of the project

It's one of the sweet names I use to call my wife. She's always full of ideas and able to help with absolutely everything.
