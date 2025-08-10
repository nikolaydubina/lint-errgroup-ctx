[![codecov](https://codecov.io/github/nikolaydubina/lint-errgroup-ctx/graph/badge.svg?token=4jwCqMiyif)](https://codecov.io/github/nikolaydubina/lint-errgroup-ctx)

Detect dangerous assignment of context by `errgroup.WithContext`[^1][^2].
This linter will check that returned context is not `ctx`, this enforces practice either assigning context to `_` or a new name.

```go
	g, ctx := errgroup.WithContext(context.Background()) // want "context from errgroup.WithContext should not be named 'ctx'"
	...
	g.Wait()
```

[^1]: https://news.ycombinator.com/item?id=44845953
[^2]: https://github.com/golang/go/issues/34510
