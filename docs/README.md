# Development documentation

## Design Philosophy

1. **Only use standard library**:
  To fit in *everywhere* without any dependency issues, `xstd` MUST only
  use the standard library.
2. **Clear separation of primitives and helpers**:
  Like in the standard library, primitives and helpers are kept separate.
  Primitives are types like `iter.Seq` or `map[K]V`. Helpers are
  packages like `iters` and `slices`.
