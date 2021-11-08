# go-xim

Package **xim** provides quazi‐ monotonically‐increasing unique‐identifiers.

The serialized form of the **xim-id** is safe to use as a _file_ or _directory_ name.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-xim

[![GoDoc](https://godoc.org/github.com/reiver/go-xim?status.svg)](https://godoc.org/github.com/reiver/go-xim)


## Example

Here is an example of using `package xim`:
```go
var id xim.ID = xim.Generate()
```

## Representation

Internally, the xim-id is compactly stored in an `uint64`. The anatomy of this is as follows:
```
                   unix timestamp (39-bits)
       ▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼
    0b0000000001100001100001110110100111011001110000111111000101100100
      ▲                                       ▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲
  always zero (1-bit)                                  chaos (20-bits)
```

The `xim.ID.UnixTime()` method will give you that 39-bit _unix timestamp_.

And the `xim.ID.Chaos()` method will give you that 24-bit _chaos_.

(The _chaos_ is just a randomness that helps make these xim-ids unique, when multiple xim-ids are being produced simultaneously.)

## Temporal Ordering of Representation

On thing to note that is, because this puts the _unix timestamp_ at the most-significant-bits, that the numerical-ordering of this `uint64` will almost always be the same as temporal-ordering.

This was done intentionally.

## Serialization

This is serialized using what this package calls **xim** notation.
An example of **xim** notation looks like this:
```
    xi-556PVvNyq3m
```

The anatomy of the **xim** string is as follows:
```
  a special base64 encoding
      ▼▼▼▼▼▼▼▼▼▼▼
    xi-556PVvNyq3m
    ▲▲           ▲
  prefix      suffix
```

So every **xim** string starts with the characters `xi`, and ends with the character `m`.
And then what is in the middle of those is a special base64 encoding of the `uint64` mentioned before.

What is special about the base64 encoding is that it doesn't use the usual ordering of the symbols.
This special base64 encoding uses:
```
┌───────┬────────┬──────┐
│ Index │ Binary │ Char │
├───────┼────────┼──────┤
│   0   │ 000000 │  -   │
├───────┼────────┼──────┤
│   1   │ 000001 │  0   │
├───────┼────────┼──────┤
│   2   │ 000010 │  1   │
├───────┼────────┼──────┤
│   3   │ 000011 │  2   │
├───────┼────────┼──────┤
│   4   │ 000100 │  3   │
├───────┼────────┼──────┤
│   5   │ 000101 │  4   │
├───────┼────────┼──────┤
│   6   │ 000110 │  5   │
├───────┼────────┼──────┤
│   7   │ 000111 │  6   │
├───────┼────────┼──────┤
│   8   │ 001000 │  7   │
├───────┼────────┼──────┤
│   9   │ 001001 │  8   │
├───────┼────────┼──────┤
│  10   │ 001010 │  9   │
├───────┼────────┼──────┤
│  11   │ 001011 │  A   │
├───────┼────────┼──────┤
│  12   │ 001100 │  B   │
├───────┼────────┼──────┤
│  13   │ 001101 │  C   │
├───────┼────────┼──────┤
│  14   │ 001110 │  D   │
├───────┼────────┼──────┤
│  15   │ 001111 │  E   │
├───────┼────────┼──────┤
│  16   │ 010000 │  F   │
├───────┼────────┼──────┤
│  17   │ 010001 │  G   │
├───────┼────────┼──────┤
│  18   │ 010010 │  H   │
├───────┼────────┼──────┤
│  19   │ 010011 │  I   │
├───────┼────────┼──────┤
│  20   │ 010100 │  J   │
├───────┼────────┼──────┤
│  21   │ 010101 │  K   │
├───────┼────────┼──────┤
│  22   │ 010110 │  L   │
├───────┼────────┼──────┤
│  23   │ 010111 │  M   │
├───────┼────────┼──────┤
│  24   │ 011000 │  N   │
├───────┼────────┼──────┤
│  25   │ 011001 │  O   │
├───────┼────────┼──────┤
│  26   │ 011010 │  P   │
├───────┼────────┼──────┤
│  27   │ 011011 │  Q   │
├───────┼────────┼──────┤
│  28   │ 011100 │  R   │
├───────┼────────┼──────┤
│  29   │ 011101 │  S   │
├───────┼────────┼──────┤
│  30   │ 011110 │  T   │
├───────┼────────┼──────┤
│  31   │ 011111 │  U   │
├───────┼────────┼──────┤
│  32   │ 100000 │  V   │
├───────┼────────┼──────┤
│  33   │ 100001 │  W   │
├───────┼────────┼──────┤
│  34   │ 100010 │  X   │
├───────┼────────┼──────┤
│  35   │ 100011 │  Y   │
├───────┼────────┼──────┤
│  36   │ 100100 │  Z   │
├───────┼────────┼──────┤
│  37   │ 100101 │  _   │
├───────┼────────┼──────┤
│  38   │ 100110 │  a   │
├───────┼────────┼──────┤
│  39   │ 100111 │  b   │
├───────┼────────┼──────┤
│  40   │ 101000 │  c   │
├───────┼────────┼──────┤
│  41   │ 101001 │  d   │
├───────┼────────┼──────┤
│  42   │ 101010 │  e   │
├───────┼────────┼──────┤
│  43   │ 101011 │  f   │
├───────┼────────┼──────┤
│  44   │ 101100 │  g   │
├───────┼────────┼──────┤
│  45   │ 101101 │  h   │
├───────┼────────┼──────┤
│  46   │ 101110 │  i   │
├───────┼────────┼──────┤
│  47   │ 101111 │  j   │
├───────┼────────┼──────┤
│  48   │ 110000 │  k   │
├───────┼────────┼──────┤
│  49   │ 110001 │  l   │
├───────┼────────┼──────┤
│  50   │ 110010 │  m   │
├───────┼────────┼──────┤
│  51   │ 110011 │  n   │
├───────┼────────┼──────┤
│  52   │ 110100 │  o   │
├───────┼────────┼──────┤
│  53   │ 110101 │  p   │
├───────┼────────┼──────┤
│  54   │ 110110 │  q   │
├───────┼────────┼──────┤
│  55   │ 110111 │  r   │
├───────┼────────┼──────┤
│  56   │ 111000 │  s   │
├───────┼────────┼──────┤
│  57   │ 111001 │  t   │
├───────┼────────┼──────┤
│  58   │ 111010 │  u   │
├───────┼────────┼──────┤
│  59   │ 111011 │  v   │
├───────┼────────┼──────┤
│  60   │ 111100 │  w   │
├───────┼────────┼──────┤
│  61   │ 111101 │  x   │
├───────┼────────┼──────┤
│  62   │ 111110 │  y   │
├───────┼────────┼──────┤
│  63   │ 111111 │  z   │
└───────┴────────┴──────┘
```

The advantage of this is that lexical-ordering (in Unicode & ASCII) is also symbol-ordering.

One design goal is that lexical ordering of the **xim** strings is (almost always) also temporal ordering of the **xim** strings.

Another design goal of it is that these **xim** strings should be able to be used as a file-name, or a directory-name.
