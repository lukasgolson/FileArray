
# FileArray

## Overview
Welcome to **FileArray**, a Go package designed to offer robust, file-based data structures. This package includes a dynamic array (`FileArray`) and a linked list (`FileLinkedList`) -- with more to come soon. These structures are tailored to handle large volumes of data in files efficiently, ensuring fast read and write operations.

## Features

### FileArray
- **FileArray** is a dynamic array that keeps its elements in a file.
- Provides constant-time access to elements.
- Specially optimized for large data sets.

#### Example Usage

```go
package main

import (
	"github.com/lukasgolson/FileArray/fileArray"
	"github.com/lukasgolson/FileArray/util"
)

func main() {
	fa, _ := fileArray.NewFileArray[util.Number]("array.dat", false)
	fa.Append(42)
	value, _ := fa.GetItemFromIndex(0)
	println(value) // Output: 42
}
```

### FileLinkedList
- **FileLinkedList** is a linked list where nodes are stored in a file.
- Ideal for situations where elements need frequent addition or removal.
- Allows for multiple datasets to be stored in a single file by enabling multiple head nodes.

#### Example Usage

```go
package main

import (
	"github.com/lukasgolson/FileArray/fileLinkedList"
	"github.com/lukasgolson/FileArray/util"
)

func main() {
	ll, _ := fileLinkedList.NewFileLinkedList[util.Number]("list.dat", false)
	ll.Add(0, 42)
	value, _ := ll.Get(0, 0)
	println(value) // Output: 42
}
```

## Contributing
Contributions are highly encouraged! Feel free to open issues or submit pull requests to improve the package's functionalities.

## License
This project is licensed under the BSD 3-Clause "New" or "Revised" License.
