# xstd

A collection of generic helpers for Go.

[![CI](https://github.com/spheric-cloud/xstd/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/spheric-cloud/xstd/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/spheric.cloud/xstd.svg)](https://pkg.go.dev/spheric.cloud/xstd)

## Installation

```bash
go get spheric.cloud/xstd
```

## Packages

This repository is organized into the following packages:

### [`chans`](chans)

Generic helpers for working with channels. Includes context-aware `Offer` / `Poll` for
single sends and receives, as well as iterator-based `SendSeq`, `RecvSeq`, `OfferSeq`,
`PollSeq` (and their `NoError` variants) to bridge channels and `iter.Seq`.

### [`container/hashmap`](container/hashmap)

A generic hash map keyed by a user-supplied hash and equality function, allowing non-comparable
key types.

### [`container/hashset`](container/hashset)

A generic hash set backed by a user-supplied hash and equality function, allowing non-comparable
value types.

### [`container/keymap`](container/keymap)

A generic map that derives a comparable internal key from each key via a user-supplied function,
allowing non-comparable key types with O(1) lookups.

### [`container/keyset`](container/keyset)

A generic set that derives a comparable internal key from each value via a user-supplied function,
allowing non-comparable value types with O(1) membership checks.

### [`container/list`](container/list)

A generic, type-safe doubly-linked list — a drop-in replacement for the standard library's
`container/list` with full iterator support (`Values`, `Elems`).

### [`container/orderedkeymap`](container/orderedkeymap)

A generic ordered map that derives a comparable internal key from each key via a user-supplied
function and preserves insertion order.

### [`container/orderedkeyset`](container/orderedkeyset)

A generic ordered set that derives a comparable internal key from each value via a user-supplied
function and preserves insertion order.

### [`container/orderedmap`](container/orderedmap)

A generic ordered map with comparable keys that preserves insertion order.

### [`container/set`](container/set)

A lightweight generic set for comparable types, implemented as `map[V]struct{}`.

### [`container/squeue`](container/squeue)

A generic sequential queue backed by a dynamically growing/shrinking circular slice buffer.

### [`funcs`](funcs)

Generic function adapters and combinators — `Identity`, `Const`, type-casting wrappers
(`CastIn`, `CastOut`, …), and more.

### [`gen`](gen)

Generic value helpers — type assertions (`Cast`, `As`, `IsA`), zero-value utilities
(`Zero`, `IsZero`, `New`), and development placeholders (`TODO`, `Stub`).

### [`iters`](iters)

A rich set of helpers and adapters for Go iterators (`iter.Seq` / `iter.Seq2`):

- **Sources** — `Of`, `FromNext`, `FromNext2`, `Repeat`, `Repeat2`, `RepeatFunc`, …
- **Transforms** — `Tap`, `Flatten`, `Map`, `FlatMap`, `Zip`, `Enumerate`, `Chunk`, `Window`, …
- **Selection** — `Unique`, `Filter`, `Take`, `Skip`, `TakeWhile`, `SkipWhile`, …
- **Sinks** — `All`, `Any`, `Count`, `Reduce`, `Fold`, `Min`, `Max`, `Sum`, `ForEach`, `Collect`, …
- **Try variants** — error-aware versions of the above for `iter.Seq2[V, error]` sequences.

### [`net/xhttp`](net/xhttp)

Helpers for `net/http` — currently provides `JSON` for writing JSON responses with a status code.

### [`ptrs`](ptrs)

Utility functions for working with pointers — `To`, `Deref`, `DerefOr`, `DerefOrZero`,
`Map`, `MapRef`, and iterator helpers for pointer sequences.

### [`sets`](sets)

Generic functions for working with `set.Set` — `HasAll`, `HasAny`, `Clone`, `Difference`,
`Union`, `Equal`, and more.

### [`xconstraints`](xconstraints)

Type constraints for generic programming — `Channel`, `Send`, and `Receive` constraints
that match any channel directionality.

### [`xmaps`](xmaps)

Generic functions for working with maps — `Inverse`, `KeysByValue`, `GetAny`, `Pop`, and more.

### [`xslices`](xslices)

Generic functions for working with slices — `Map`, `Flatten`, `Filter`, `Unique`, `Random`,
`Collect`, `TryCollect`, channel-to-slice helpers (`CollectRecvChan`, `CollectPollChan`), and more.

### [`xstrings`](xstrings)

Helper functions for working with strings — `JoinSeq`, `WriteJoining`, and `Random`
(random string generation from a character set).

## License

This project is licensed under the **Apache License 2.0**. See the [LICENSE](LICENSE) file for details.