# `gochat`

Simple AI "chat" shell.

## Model Support

This app is designed to support access to multiple underlying chat
models, including:

- [OpenAI](https://platform.openai.com/docs) using [Go OpenAI SDK](https://github.com/sashabaranov/go-openai)
- [Gemini](https://deepmind.google/technologies/gemini) using [Google Gemini SDK](https://github.com/google/generative-ai-go)

## Configuration 

Copy the file `config/config.yaml` into `config/config-local.yaml` in order to create
your own, private configuration.

In your private configuration (i.e., `config-local.yaml`), select the model you wish to use,
and enter your credentials for accessing that model before running.

If you wish to create transcripts of your chat sessions, set the `Scriber` configuration value
to `Template`, and - if needed - customize its formatting in the `TemplateScribe` sections.

## Running

In a `go` (version `1.21` or later) environment, run the app using the `Makefile` target, e.g.:

```shell
$ make run-chat
go run cmd/chat/main.go
Using model: gpt-4-1106-preview
Type 'exit' to quit
Ask me anything: 
> Tell me in three words who you are.
Artificial Intelligence Assistant
> exit
Goodbye!
Bye bye!
$ 
```

