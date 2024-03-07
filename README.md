# TL

TL is a data description language and data storage format.

## Overview

In general, interest in TL stems from a desire to serialize data and make RPC calls to servers that implement TL RPC, as well as to implement such servers

TL describes data structures, including RPC queriesand their responses, using syntax derived from functional programming languages. 

The TL format is characterized by compactness and high efficiency.

TL is schema-driven format. Tool caled `tlgen` is used to generate structs/classes and (de)serialization methods.


## Installation

Install `tlgen` with the following command

```
go install github.com/vkcom/tl/cmd/tlgen@latest
```

## Documentation

- [Internals](./docs/TLPrimer.pdf) (in Russian)

## License

TL is licensed under the [Mozilla Public License Version 2.0](./LICENSE).
