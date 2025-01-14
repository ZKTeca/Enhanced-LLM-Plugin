# Enhanced LLM Plugin 

A better and enhanced version of the LLM Plugin system, now maintained by ZKTeca.

## 1. Plugins

### 2.1 Google Search

Get Google Search token: [https://docs.chatkit.app/tools/google-search.html](https://docs.chatkit.app/tools/google-search.html)

### 2.2 Stable Diffusion

Generate photo by stable diffusion plugin, example:

![girl](./plugins/stablediffusion/test1.jpg)

## 2. TESTING

1. OpenAI:
   ```bash
   cp .env.example .env
   ```

2. Google:
   ```bash
   cd plugins/google

   cp .env.example .env
   ```

Run test:

```bash
 go test -v ./...
```


## 3. RELEASE

### v0.1.0

1. Init project.
2. Support plugin: Google for search, calculator for mathematical calculations.


## Contribute

We are open to enhancements & bug-fixes so feel free to enhance and contribute to this project.