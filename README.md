# parquet-go-source

parquet-go-source is a source provider for parquet-go. Your source must implement ParquetFile interface:

```go
type ParquetFile interface {
	io.Seeker
	io.Reader
	io.Writer
	io.Closer
	Open(name string) (ParquetFile, error)
	Create(name string) (ParquetFile, error)
}
```

Note: This is an aggressively slimmed version of [xitongsys' codebase](https://github.com/xitongsys/parquet-go-source)
