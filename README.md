# `gochat`
<img align="right" alt="main" src="https://github.com/noodnik2/gochat/actions/workflows/ci.yml/badge.svg?branch=main" />

## Description

A simple AI "chat" shell (similar to `ChatGPT`) which can be used to "converse" with
one of the supported large language models, and optionally record your "chat transcript"
in a local file, making it easy to later reference your conversations with these models.

## _Currently Available!_

### Model Support

This app is designed to support access to multiple underlying chat models, 
including:

- [OpenAI](https://platform.openai.com/docs) using [Go OpenAI SDK](https://github.com/sashabaranov/go-openai)
- [Gemini](https://deepmind.google/technologies/gemini) using [Google Gemini SDK](https://github.com/google/generative-ai-go)

If you're not yet signed up to access either of these popular online language models,
_what are you waiting for?_  Clear instructions for obtaining your SDK Key needed to
access each model on its "SDK" page using the link(s) above! 

### Configuration 

Copy the file `config/config.yaml` into `config/config-local.yaml` in order to create
your own, private configuration.

In your private configuration (i.e., `config-local.yaml`), select the model you wish to use,
and enter your credentials for accessing that model before running.

If you wish to create transcripts of your chat sessions, set the `Scriber` adapter value to
`Template`, and - if needed - customize its formatting in the `TemplateScribe` sections,
taking guidance from the [`go` Templating](https://pkg.go.dev/text/template) language documentation.

### Running

You will need to have `go` installed in order to run `gochat` in your terminal.  If you 
don't already have this, see [here](https://go.dev/doc/install) to find out how.

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

## _Future Roadmap_

Several envisioned enhancements for `gochat` in its Roadmap include:

### Retrieval Augmented Generation (RAG)
- Add support for "chats" with supplied context, using an "embedding" service such as
  [ChromaDB](https://www.trychroma.com/), [Pinecone](https://www.pinecone.io/solutions/rag/)
  or [Weaviate](https://weaviate.io/developers/weaviate/starter-guides/generative)

### Support for Fully Local LLMs
- Add support for "chats" against confidential information, and/or without the need
  for internet connection, such as possible with [ollama](https://ollama.com/).

### Multimodal Support
- Now that many LLMs (including [Gemini](https://cloud.google.com/vertex-ai/docs/generative-ai/multimodal/overview))
  support media files (both as prompts and as responses), support for this paradigm
  could be really useful, and not so hard to implement.
- Envisioned is some sort of "meta-command" syntax to support access to local or
  remote media, TBD, ...

### Speech-to-text / Text-to-speech
- Use something like Google's [STT](https://pkg.go.dev/cloud.google.com/go/speech) and
  [TTS](https://pkg.go.dev/cloud.google.com/go/texttospeech) APIs to allow users to
  interact with the LLM using speech / audio.

## Want to Contribute?  Have Ideas?
- _Please get in touch!_
- _Please submit a PR!_