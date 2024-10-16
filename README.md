# Aicli

## Description

Projet to use ai api to generate text, image, etc.

## To do

- [x] command text
- [x] command translate
- [ ] command image
- [ ] command speech
- [ ] command tts

## Available SDK

### Text Generation

- openai
- claude
- mistral

### Translation

- deepl

## Build and install

1. Clone the repository:

```bash
git clone https://github.com/LordPax/aicli.git
cd aicli
```

2. Build the project:

```bash
go mod download
go build
./install.sh
```

3. Execute the script to generate config

Will generate a config file at `~/.config/aicli/config.ini`

```bash
./aicli
```

## Config example

```ini
[text]
type=openai
openai-model=gpt-4
openai-apiKey=yoursecretapikey
claude-model=claude-3-5-sonnet-20240620
claude-apiKey=yoursecretapikey
temp=0.7

[translate]
type=deepl
apiKey=yoursecretapikey

[image]
type=openai
model=dall-e-3
apiKey=yoursecretapikey
```

## Integration with tmux

Add the following line to your `.tmux.conf`:

```bash
bind H new-window "aicli text" \; rename-window "aicli"
```

## Vim plugin

You can install [vim-aicli](https://github.com/LordPax/vim-aicli) to use aicli in vim.

Add the following line to your `.vimrc`:

```vim
Plug 'LordPax/vim-aicli'
```

## Integration with i3

Add the following line to your `~/.config/i3/config`:

```
bindsym $mod+s exec /usr/bin/aicli -c -g speech
```
