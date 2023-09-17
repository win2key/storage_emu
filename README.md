# Storage Emu: IBM Storage FlashSystem Emulator

## Overview

Storage Emu is a lightweight IBM Storage FlashSystem emulator designed to mimic the behavior of an actual IBM Storage system. The software allows you to execute commands as if you were operating on a real FlashSystem device.

## Features

- Interactive shell for manual interaction.
- Scriptable interface for automated tasks.
- Virtual disk management (list, add, remove).
- Command history and navigation.

## Requirements

- GoLang (for compiling)
- Linux machine with SSH server
- A dedicated user account on the Linux machine

## Installation

### Compiling

1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Run `go build` to compile the code.

```bash
go build -o storage_emu
```

### Deploying

1. Transfer the compiled `storage_emu` binary to the remote Linux machine.
2. Create a new user on the Linux machine.

```bash
sudo useradd -m -s /path/to/storage_emu username
```

3. Modify the `/etc/passwd` file to set `storage_emu` as the shell for the user.

```bash
username:x:1001:1001::/home/username:/path/to/storage_emu
```

## Usage

### Interactive Mode

To interact with the emulator via SSH:

```bash
ssh username@linux
```

You will be presented with a `FS>` prompt, where you can execute commands as if you were on an IBM Storage FlashSystem.

### Command Execution

To execute a single command remotely:

```bash
ssh username@linux command
```

## Available Commands

- `test`: Prints "Test command."
- `list disks`: Lists all virtual disks.
- `add disk`: Adds a new virtual disk.
- `remove disk`: Removes the last added disk.
- `exit`: Exits the shell.

## Limitations

This emulator is a simplified representation and does not fully capture all the functionalities of an IBM Storage FlashSystem. It is intended for testing and educational purposes only.

## Contributing

Feel free to open issues or submit pull requests to improve the project.

## License

This project is open-source software licensed under the MIT License. Please see the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to all contributors and users of this project.