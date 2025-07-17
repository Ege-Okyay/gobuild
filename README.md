# gobuild

A simple command-line tool to cross-compile Go applications based on a YAML configuration file.

## Installation

1.  **Download:** Go to the [releases page](https://github.com/Ege-Okyay/gobuild/releases) and download the correct binary for your system.

2.  **Prepare the Executable:**

    **For Windows:**
    -   Rename the downloaded file (e.g., `gobuild-windows-amd64.exe`) to `gobuild.exe`.

    **For macOS & Linux:**
    -   Rename the downloaded file (e.g., `gobuild-linux-amd64`) to `gobuild`.
    -   Make it executable:
        ```bash
        chmod +x gobuild
        ```

3.  **Add to PATH (Recommended):** To use `gobuild` from any terminal, move it to a directory in your system's PATH.

    **For Windows:**
    -   Move `gobuild.exe` to a folder that is in your PATH (e.g., `C:\Windows\System32`), or add a new folder (e.g., `C:\Tools`) to your PATH.

    **For macOS & Linux:**
    -   Move the `gobuild` file to `/usr/local/bin`:
        ```bash
        sudo mv gobuild /usr/local/bin/
        ```

## Usage

Once installed, run the tool from your project directory:

```bash
gobuild config.yaml
```

This will compile your project based on the settings in `config.yaml` and place the binaries in a `builds` directory.

## Configuration

Create a `config.yaml` file to define your build targets.

-   `source_dir`: The directory of your Go application's main package. Defaults to `.` if not specified.
-   `build_name`: The base name for the compiled binaries.
-   `builds`: A list of target platforms, each with a `goos` (OS) and `goarch` (architecture).

### Example `config.yaml`

```yaml
# The directory of your Go application's main package.
# Defaults to the current directory (".") if not specified.
source_dir: .
# The base name for the output binaries.
build_name: my-app
# A list of target platforms to build for.
builds:
  - goos: windows # Target operating system (e.g., windows, linux, darwin).
    goarch: amd64 # Target architecture (e.g., amd64, arm64).
  - goos: linux
    goarch: amd64
  - goos: darwin
    goarch: amd64
```