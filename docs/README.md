# Development documentation

## Design Philosophy

1. **Drop-In Compatibility**:
  `xstd` aims to be a drop-in replacement of the go standard libraries.
  To achieve package-level compatibility, [`dropin-gen`](https://github.com/spheric-cloud/dropin-gen) is
  used to generate the necessary forwarding (aliases, vars, constants etc.).
2. **Only use standard library**:
  To fit in *everywhere* without any dependency issues, `xstd` MUST only
  use the standard library.
3. **Clear separation of primitives and helpers**:
  Like in the standard library, primitives and helpers are kept separate.
  Primitives are types like `iter.Seq` or `map[K]V`. Helpers are
  packages like `iters` and `slices`.
4. **No helper dependency**:
  Helpers MUST be independent of each other. This means that no helper
  package may import another helper package. This requirement helps
  facilitate that a package might be pulled out as a standalone module
  in the future and avoids cyclic dependency issues between helper packages.
  This requirement might be lifted in the future after careful consideration.
