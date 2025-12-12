# Zstash Buildkite Plugin

Uses zstash to cache dependencies.

# Example Configuration

```
steps:
  - plugins:
    - buildkite/zstash:
        ids: "node,docker"
        config: "./path/to/cache.yml"
```

## License

This application is released under Apache 2.0 license.
