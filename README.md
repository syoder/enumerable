# Enumerable: An experiment in Go

The idea is to see how feasible it is to create something like Ruby's
[`Enumerable`](https://ruby-doc.org/core-2.6.5/Enumerable.html) in Go.
In order to do this in a generic way requires use of Go's
[`reflect`](https://golang.org/pkg/reflect/) package.

## Goals

  * a better understanding of `reflect` and the related performance issues
  * an easy to use package which at the very least could be used when some
performance can be traded for ease of use and readability (i.e. when IO is
already the bottleneck, this will probably be fast enough)
  * where possible, fast implementations will be provided so that they can
be used even when top performance is required

## Examples

  * Count total letters
```go
var list = []string{"one", "two", "three"}
letterCount := enumerable.New(list).Map(func(item string) int {
  return len(item)
}).Reduce(0, func(memo, l int) int {
  return memo + l
} // letterCount = 11
```

  * get a sorted list of unique words in files
```go
nonAlphabet := regexp.MustCompile("[^a-zA-Z]")
files, _ := filepath.Glob("*.txt")
words := enumerable.New(files).Map(func(file string) []string {
  contents, err := ioutil.ReadFile(file)
  if err != nil {
    panic(err)
  }
  return nonAlphabet.Split(string(contents), -1)
}).Flatten().Uniq().Sort().JoinAsString(", ")
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/syoder/enumerable. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

This package is available as open source under the terms of the
[MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Enumerable project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/syoder/enumerable/blob/master/CODE_OF_CONDUCT.md).
