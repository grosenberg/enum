### Go Rich Enums

Concise, expressive implementation of rich enums in Go.  

A singleton generic base implementation that provides all essential enum behaviors, including global uniqueness and mutual-comparability. Rich enum sets are realized effectively as separate type-specializations of the generic base.  These type-specializations are minimal and can be easily implemented in one or more user packages.

Rich enums can encapsulate any number of custom member fields and corresponding functions.  Rich enums are type-safe, immutable, and globally unique. Within a defined enum set, the enums are mutually comparable; enums from different sets are incomparable.

Part of the Go Generics proof-of-concept packages:
[Collections](https://github.com/grosenberg/collections)
[Generic](https://github.com/grosenberg/generic)
[Maths](https://github.com/grosenberg/maths)
&
[Rich Enums](https://github.com/grosenberg/enum)
.

#### Documentation

Rich enums uses a variant of the Go generics pattern. 

For a summary of the Go generics pattern, see the [Go Generics](https://github.com/grosenberg/maths/Go Generics.md) file.

### License

Copyright © 2015 Gerald Rosenberg.

Use of this source code is governed by a BSD-style license that can be found in the License.md file.
