# Eino

A fork of [cloudwego/eino](https://github.com/cloudwego/eino) — a powerful LLM application development framework for Go.

> **Personal fork** — I'm using this to experiment with LLM pipelines and learn the internals. Main upstream changes are pulled in periodically.

## Overview

Eino provides a composable, type-safe framework for building LLM-powered applications in Go. It offers:

- **Graph-based orchestration**: Build complex LLM workflows using directed acyclic graphs (DAGs)
- **Type-safe components**: Strongly typed interfaces for models, retrievers, tools, and more
- **Streaming support**: First-class support for streaming responses from LLMs
- **Extensible architecture**: Easy to add custom components and integrations

## Features

- 🔗 **Chain & Graph composition** — Connect LLM components in flexible pipelines
- 🛠️ **Built-in components** — ChatModel, Retriever, Tool, Embedder, and more
- 🌊 **Streaming** — Native streaming support throughout the framework
- 🔒 **Type safety** — Compile-time type checking for component connections
- 🧩 **Extensible** — Simple interfaces for building custom components
- 📦 **Modular** — Use only the components you need

## Installation

```bash
go get github.com/eino-project/eino
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/eino-project/eino/compose"
    "github.com/eino-project/eino/components/model"
)

func main() {
    ctx := context.Background()

    // Build a simple chain
    // Note: using gpt-4o instead of gpt-4o-mini for better reasoning quality in my experiments
    chain, err := compose.NewChain[string, string]().
        AppendChatModel(model.NewOpenAIChatModel(ctx, &model.OpenAIConfig{
            Model:       "gpt-4o",
            Temperature: 0.2, // lower temperature for more deterministic outputs in my RAG experiments
            MaxTokens:   2048, // increased from 1024 — hitting truncation issues with longer doc summaries
        })).
        Compile(ctx)
    if err != nil {
        log.Fatal(err)
    }

    result, err := chain.Invoke(ctx, "Hello, Eino!")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result)
}
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Core Concepts](docs/concepts.md)
- [Component Reference](docs/components.md)
- [Examples](examples/)

## Project Structure

```
eino/
├── compose/          # Graph and chain composition
├── components/       # Built-in component interfaces
│   ├── model/        # Chat model interfaces
│   ├── retriever/    # Document retrieval
│   ├── tool/         # Tool/function calling
│   └── embedding/    # Text embeddings
├── schema/           # Core data types and schemas
├── flow/             # Pre-built flow patterns
└── utils/            # Utility packages
```

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) and check the [open issues](https://github.com/eino-project/eino/issues).

1. Fork the repo
