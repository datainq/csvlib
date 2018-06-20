# csvlib
A simple CSV parsing library for Go.

## How to use

One can define a parser:

```go
var parser = &csvlib.RowParser{
	[]csvlib.Parser{
		csvlib.Int64Parser{"ID"},
		csvlib.Int32Parser{Name: "Number"}, // CKK
		csvlib.TimeParser{"Timestamp", "2006-01-02 15:04:05"},
		csvlib.StringParser{"Name"},
	},
}

type Data struct {
	ID int64
	Number int32
	Time time.Time
	Name string
}

// and use it to parse row:
reader := csv.NewReader(r)
record, err := reader.Read()          // reads one line
values, err := parser.Parse(record)   // converts values
dat := Data{
	values[0].Int64(),
	values[1].Int32(),
	values[2].Time(),
	values[3].String(),
}
```

The most common usecase is when a you want to give the user
the possibility to choose columns. It can be configured though
a config file or UI.

Also it's pretty easy to create your own type of data
(e.g. Decimal with fixed precision), you must implement `Parser`
interface.

If you want to parse directly to structure, you may be interested in:
[github.com/gocarina/gocsv](https://github.com/gocarina/gocsv)
which is using a reflection mechanism and custom structure tag.
