# Footprint

Generate Footprint—create a random string of words taken from the dictionary.


The dictionary (words_alpha.txt) is from [dwyl/english-words](https://github.com/dwyl/english-words).


## Usage

```go
import (
  "fmt"
  "github.com/five-ten-github/footprint"
)

fmt.Println(footprint.GenerateFootprint(3))

// Output: "random-word-group"
```
