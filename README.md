# gobsp

Go Binary Stream Protocol

# Description

gobsp is a extensible binary protocol for passing messages between
computer programs. It is designed to be consumed similar to how
streams of other information are consumed: by a scanner. gobsp
includes a scanner for decoding an incoming stream of data, and an
writer for encoding and writing data to an outgoing stream.

# Usage Example

```Go
    package example

    import (
        "log"
        "github.com/karrick/gobsp"
    )
    
    // NOTE: Define some message types our program can understand.
    const (
        MTError gobsp.MessageType = iota
        MTGreeting
        MTFarewell
        // NOTE: You ought not rearrange these for newer protocol message
        // types, just add more below.
    )

    func handleError(ior io.Reader) error {
        buf, err := io.ReadAll(ior)
        if err != nil {
            return err
        }
        fmt.Printf("An error took place: %s", string(buf))
    	return nil
    }
    ​
    func handleGreeting(ior io.Reader) error {
        buf, err := io.ReadAll(ior)
        if err != nil {
            return err
        }
        fmt.Printf("Hello, %s", string(buf))
    	return nil
    }
    ​
    func handleFarewell(ior io.Reader) error {
        buf, err := io.ReadAll(ior)
        if err != nil {
            return err
        }
        fmt.Printf("Farewell, %s", string(buf))
    	return nil
    }
    ​
    func main()  {
    	handlers := map[uint32]MessageHandler {
            MTError: handleError,
    		MTGreeting: handleGreeting,
    		MTFarewell: handleFarewell,
    	}

        proc, err := gobsp.NewProcessor(iow,
            DefaultHandler(dh),
            DefaultHandler(dh),
            Handlers(handlers),
        )
    	if err != nil {
    		log.Fatal(err)
    	}
        if err := proc.Process(); err != nil {
    		log.Printf("WARNING: there was an error; moving on: %s", err)
    	}
    }
```

*WARNING:* There are two kinds of errors: (1) those that occur due to
failure to read data from the stream; and (2) those that occur during
processing of a particular message. It is imperative that message
handlers only return errors when there is a failure to read data from
the stream. If a handler can read data, but there is an error
interpreting it, or an error taking some action on it, the handler
ought not return that error. Rather, it should handle it in some way
and return nil, which tells the processor to continue reading from the
stream and handle more messages.

# Protocol

gobsp has multiple layers of protocols for different purposes.

* stream layer (message framing)
* user-defined message type layer
* primitive data type layer

The stream layer describes how bytes are pulled off the wire and
divided into messages. The user-defined layer is controlled by the
code using this library. Finally, the primitive data type layer
describes how individual pieces of data are encoded in the
user-defined messages.

## Primitive Data Type Layer

gobsp supports several primitive data types, such as both signed and
unsigned fixed width and variable width integers, floating point
numbers, strings, and a lists of strings.

## Stream Layer

Messages are read from the byte stream one after the other using a
simple framing method. Messages are framed with two control integers,
the message type and the message size, and followed by the message
payload. The format of the message payload is determined by the
message type. Each message type will have a particular payload format,
which is determined by the application.

One drawback to the simplicity of the message framing described above
is inability to resynchronize a parser if it ever drops sync with the
byte stream.

Message type and message size are each encoded as unsigned 16-bit
integers in big-endian format. This keeps the per-message overhead to
4 bytes, while leaving plenty of room for a lot of message types and
rather large message sizes. Messages larger than 64 KiB would have to
be split up into multiple messages.

### Message Type and Version

The message type integer does double duty and, for a particular
application, encodes both the message's type and version. For example,
if message type _1_ has a particular meaning to your application, and
you must upgrade the message payload format, simply allocate a new
number for the new message format. Instead of creating version _2_ of
message type _1_, you simply create message type _N+1_, where _N_ was
previously the largest message type number.

For this reason, once your application associates particular integers
with particular message payload formats, it is never advisable to
modify those formats. Rather keep the old source code for processing
older message types and add source code for processing new message
types.

The side-effect of combining a message type and version are to create
a larger group of message types.

The advantages of combining message type and version includes the
simplification of upgrading the payload format of a particular message
type, and does not tie a bunch of message types to a particular
version of the protocol.

Furthermore, there is no protocol negotiation phase.

The disadvantages of combining message type and version include the
fact that there is no way in the protocol itself to specficy minimum
protocol version. However, a particular application _could_ create a
message type for protocol version. For instance, perhaps message type
_0_ is a protocol version declaration, and message type _1_ specifies
an error in the protocol negotiation, and means in this application
that the other end ought to cease data transmission. In this example,
the message types convey the protocol negotiation phase.

## Primitive Data Types

While applications are free to format message payloads in any way
suitable to their purpose, the following primitive data type encodings
are provided. All numbers are encoded in big-endian format. Floats are
IEEE-754.

* Float32
* Float64
* Int8 & Uint8
* Int16 & Uint16
* Int32 & Uint32
* Int64 & Uint64
* String
* StringSlice
* Variable Width Integer (VWI) & Unsigned Variable Width Integer (UVWI)

### Performance

When encoding or decoding data for the Int8, Uint8, VWI, and UVWI data
types, and the provided stream implements the io.ByteReader (for
decoding), or the io.ByteWriter (for encoding) interface, the program
benefits from an approximate three fold performance boost. In the
benchmarks below, the buffer.Buffer instance only implements io.Reader
and io.Writer interfaces.  The bytes.Buffer instance also implements
the io.ByteReader and io.ByteWriter interfaces.

```
BenchmarkBinaryBuffer-4            	 2000000	       963 ns/op
BenchmarkBinaryBytes-4             	 5000000	       312 ns/op
```

### Variable Width Integer (VWI)

Variable Width Integers encode numbers as large as 64-bit integers by
repurposing the high-bit of every byte as a flag to represent whether
or not additional bytes are used to encode this number. Therefore,
each encoded byte can only encode up to 7-bits of the original
value. This may cause some numbers to be encoded using more bytes than
encoding them as a fixed width integer. For example, while the number
127 requires one byte to encode, the number 128 requires two bytes.

The advantage of using VWI is that both large and small numbers can be
encoded using this data type, but a particular encoded value will
consume as few bytes as required to encode that number. Because
numbers in many applications are relatively small, most applications
benefit from the compromise.

The disadvantage of using VWI is the small computation overhead
required for encoding and decoding VWI numbers compared to equivalent
unsigned integer numbers.

#### Benchmarks

There is no doubt that VWI imposes a certain computational overhead
during encoding or decoding when compared to merely writing the byte
values from fixed width integers. The following benchmark values were
collected on my laptop, but one may expect similar percentage impacts
on other platforms.

```
BenchmarkBinaryOneByteFixed-4      	20000000	       123 ns/op
BenchmarkBinaryOneByteVariable-4   	10000000	       140 ns/op

BenchmarkBinaryTwoBytesFixed-4     	10000000	       211 ns/op
BenchmarkBinaryTwoBytesVariable-4  	10000000	       157 ns/op

BenchmarkBinaryFourBytesFixed-4    	10000000	       202 ns/op
BenchmarkBinaryFourBytesVariable-4 	10000000	       208 ns/op

BenchmarkBinaryEightBytesFixed-4   	 5000000	       211 ns/op
BenchmarkBinaryEightBytesVariable-4	 5000000	       328 ns/op
```

In general the more bytes required to store a number, the more of a
performance impact a particular program will have. For occassions when
most numbers are expected to be rather small, using VWI data types can
have negligible impact while providing additional flexibility when few
numbers are large in magnitude.

The takeaway here is that while variable width integers provide
greater flexibility and more compact representations of many numbers,
they require additional overhead. Use them when needed and use fixed
width integers when you don't.

### String

This protocol encodes strings as a VWI representing the number of
bytes for the string, followed by the actual bytes in the string.

### StringSlice

This protocol encodes string slices as a VWI representing the number
of strings in the slice, followed by the encoded form of each string.

# References

## Big-endian format

* <https://en.wikipedia.org/wiki/Endianness#Big-endian>

## Variable Width Integers

* <https://orc.apache.org/docs/run-length.html>
* <https://en.wikipedia.org/wiki/Variable-length_quantity>

## Floating Point

* <https://en.wikipedia.org/wiki/IEEE_floating_point>
