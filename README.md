# Aicli

## Description

Projet to use ai api to generate text, image, etc.

## To do

- [x] command text
- [x] command translate
- [ ] command image
- [ ] command speech
- [ ] command tts

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
model=gpt-4
apiKey=yoursecretapikey
temp=0.7

[translate]
type=deepl
apiKey=yoursecretapikey
```

With type as prefix, parameter will be specific for this type, if there is no prefix, parameter will be global

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
```

## Integration with tmux

Add the following line to your `.tmux.conf`:

```bash
bind H new-window "aicli text" \; rename-window "aicli"
```

## Vim plugin

You can find the vim plugin here

Add the following line to your `.vimrc`:

```vim
Plug 'LordPax/vim-aigpt'
```

## Integration with i3

Add the following line to your `~/.config/i3/config`:

```
bindsym $mod+s exec /usr/bin/aicli -c -g speech
```
