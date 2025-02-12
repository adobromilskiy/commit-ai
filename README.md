# 💾 commit-ai: write your commit messages with AI 🤖

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/adobromilskiy/commit-ai)
![GitHub License](https://img.shields.io/github/license/adobromilskiy/commit-ai)


As a Go developer, I looked at the world and thought:

> You know what this planet really needs? One more commit message generator!

So, naturally, I built one. Wellcome to another **commit-ai**, the tool for generating commit messages using the OpenAI API. It analyzes Git changes and generates concise, meaningful commit messages for your repository.

## ✨ Features

- Generates a commit message based on changes staged using `git add`.
- Automatically extracts changed directories as go packages and adds them to the message (optional).
- Applies rules for creating commit messages: starts with an imperative verb in the present tense, no period at the end, and limited to 10 words.
- Option to exclude package names or the full commit command from the message.

## 📋 Requirements

- Go 1.18 or higher
- OpenAI API key (set in the `OPENAI_API_KEY` environment variable)

## 🚀 Getting Started

You can install commit-ai directly using Go:

```bash
go install github.com/adobromilskiy/commit-ai@latest
```

Or manually:

1. Clone the repository:

    ```bash
    git clone https://github.com/adobromilskiy/commit-ai.git
    cd commit-ai
    ```

2. Install dependencies:

    ```bash
    go install
    ```

3. Set your OpenAI API key:

    ```bash
    export OPENAI_API_KEY=your-api-key-here
    ```

## 📚 Usage

### Main Command

Run the command:

```bash
commit-ai
```

This will generate a commit message based on the staged changes in your repository.

### Flags

- `--no-pkg` — exclude package names from the commit message.
- `--no-cmd` — output only the commit message without the `git commit` command.

### Command Examples:

1. Generate a commit message with package names:

```bash
commit-ai
```

2. Generate a commit message without package names:

```bash
commit-ai --no-pkg
```

3. Output only the commit message, without the `git commit` command:

```bash
commit-ai --no-cmd
```

4. Generate just the commit message without packages and command:

```bash
commit-ai --no-pkg --no-cmd
```

### Example Output

If there are staged changes, the tool may generate a commit message like:

```bash
git commit -m "core, database: refactor incoming params"
```

If the `--no-pkg` flag is enabled, it will generate:

```bash
git commit -m "refactor incoming params"
```
