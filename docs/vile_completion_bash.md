## vile completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(vile completion bash)

To load completions for every new session, execute once:

#### Linux:

	vile completion bash > /etc/bash_completion.d/vile

#### macOS:

	vile completion bash > $(brew --prefix)/etc/bash_completion.d/vile

You will need to start a new shell for this setup to take effect.


```
vile completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.vile.yaml)
```

### SEE ALSO

* [vile completion](vile_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 17-Nov-2022