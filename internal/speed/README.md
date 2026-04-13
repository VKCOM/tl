# speed tests

Some (outdated) tests for comparison of serialization formats.

We wanted to explore possibilities for TL2 design, so tried many approaches, comparing with popular formats. 

Also, here is an answers to questions:

* Why don't you use flatbuffers or capnp approach for TL2 (answer: they are both large and slow to serialize)
* Why TL2 only has 1 int and 3 length representations (answer: it is faster to read/write than 10, as msgpack and protobuf have)
