# collections

Inspired by Java's `java.util.*`. As golang doesn't have supports for some of the most used data structures, I decided to implement them. This is just a start, and I will keep adding more data structures as I need them.

## Installation

```shell
go get github.com/that-happy-place/golean/collections@latest
```

## Usage

### Set

```go
import (
    "github.com/that-happy-place/golean/collections"
)

func main() {
	set := collections.NewHashSet[string]()
	set.Add("never gonna give you up")
    set.Add("never gonna let you down")
	set.Add("never gonna run around and desert you")
	set.Add("never gonna make you cry")
	set.Add("never gonna say goodbye")
	set.Add("never gonna tell a lie and hurt you")

    set.Add("never gonna give you up")
    set.Add("never gonna let you down")
    set.Add("never gonna run around and desert you")
    set.Add("never gonna make you cry")
    set.Add("never gonna say goodbye")
    set.Add("never gonna tell a lie and hurt you")

	set.Len() // returns 6
}
```
