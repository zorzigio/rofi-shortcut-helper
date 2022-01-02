# rofi-shortcut-helper

Use rofi to remember keyboard shortcuts.

## Usage

Just run `rofi-shortcut-helper` and a popup will appear. Select the program firts and then the shortcut.
The shortcuts are defined in the `shortcuts.json` file.

To display the menu, run `rofi-shortcut-helper --help`

```console
❯ go run main.go --help
usage: rofi-help-shortcuts [-h|--help] [-d|--dir "<value>"] [-r|--rofi
                           "<value>"]

                           Use rofi to display shortcuts

Arguments:

  -h  --help  Print help information
  -d  --dir   Path to the json file. Default: ./my-shortcuts.json
  -r  --rofi  Command line to launch rofi. Default: rofi -dmenu -p "Shortcuts
              Help" -i -no-custom -matching fuzzy
```

## Requirements

- [rofi](https://github.com/davatorium/rofi): as this will lauch rofi

- [go language tools](https://go.dev/doc/install): to build the binary

## Installation

Run the following commands in your terminal

```console
❯ go get github.com/zorzigio/rofi-shortcut-helper
❯ go install github.com/zorzigio/rofi-shortcut-helper
```

At this point you should have the executable file at `$GOPATH/bin/rofi-shortcut-helper`.
