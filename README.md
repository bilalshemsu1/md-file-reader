    <div align="center">

    # MD Editor

    **A fast, distraction-free Markdown editor for Windows**

    Built with Go and Wails — runs as a native desktop app with no browser, no server, no internet required.

    [![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
    [![Wails](https://img.shields.io/badge/Wails-v2.12+-434449?style=flat&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSJjdXJyZW50Q29sb3IiIHN0cm9rZS13aWR0aD0iMiI+PHBhdGggc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIiBkPSJNMTIgMjBIMm0xMCAwSDIybS0xMC0xMEgybTEwIDBIMjJNMTIgMTBIMm0xMCAwSDIyIi8+PC9zdmc+)](https://wails.io)
    [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
    [![Platform](https://img.shields.io/badge/Platform-Windows-blue.svg)](https://www.microsoft.com/windows)

    </div>

    ---

    ## ✨ Features

    - 📂 **Native File Operations** — Open and save files directly to disk using Windows Explorer
    - 👁️ **Live Preview** — Real-time Markdown rendering as you type
    - ✂️ **Split View** — Edit and preview side by side with resizable divider
    - 🎨 **Syntax Highlighting** — Code blocks for JS, TypeScript, Python, Bash, CSS, and more
    - 📋 **Multi-File Support** — Open multiple files, switch between them in the sidebar
    - ⌨️ **Rich Formatting** — Bold, italic, headings, lists, tables, code blocks, blockquotes
    - 🔍 **Search** — Find in document with match count and keyboard navigation
    - ✅ **Task Lists** — Interactive checkboxes that sync with source
    - 💾 **Direct Save** — Writes back to original file, no download copies
    - 🌙 **Dark Theme** — Easy on the eyes, professional look

    ---

    ## � Quick Start

    ### Download (Recommended)

    1. Go to the [Releases](../../releases) page
    2. Download `Md-Editor.exe` from the latest release
    3. Double-click to run — no installation required

    ### Build from Source

    **Prerequisites:**
    - Go 1.21+
    - Node.js 16+
    - Wails v2
    - Windows 10/11 with WebView2 (pre-installed)

    ```bash
    # Clone
    git clone https://github.com/bilalshemsu9/md-file-reader.git
    cd md-file-reader

    # Install Wails CLI
    go install github.com/wailsapp/wails/v2/cmd/wails@latest

    # Install dependencies
    cd frontend && npm install && cd ..

    # Run dev mode
    wails dev

    # Build executable
    wails build
    ```

    Output: `build/bin/Md-Editor.exe`

    ---

    ## ⌨️ Keyboard Shortcuts

    | Action | Shortcut |
    |--------|----------|
    | New file | `Ctrl + N` |
    | Open file | `Ctrl + O` |
    | Save | `Ctrl + S` |
    | Save As | `Ctrl + Shift + S` |
    | Bold | `Ctrl + B` |
    | Italic | `Ctrl + I` |
    | Find | `Ctrl + F` |
    | Cycle view | `Ctrl + \` |
    | Indent | `Tab` |

    ---

    ## � Markdown Support

    | Element | Syntax |
    |---------|--------|
    | Heading | `# H1`, `## H2`, `### H3` |
    | Bold | `**bold**` |
    | Italic | `_italic_` |
    | Strikethrough | `~~strike~~` |
    | Inline code | `` `code` `` |
    | Code block | ` ```js ` |
    | Link | `[text](url)` |
    | Image | `![alt](url)` |
    | Blockquote | `> quote` |
    | Bullet list | `- item` |
    | Numbered list | `1. item` |
    | Task list | `- [ ] task` |
    | Table | `| col |` |
    | Horizontal rule | `---` |

    ---

    ## 🏗️ Project Structure

    ```
    md-file-reader/
    ├── main.go              # App entry point, window config
    ├── app.go               # Go backend: file I/O operations
    ├── go.mod               # Go dependencies
    ├── wails.json           # Wails configuration
    ├── build/windows/
    │   ├── icon.ico         # App icon
    │   └── info.json        # Windows version info
    └── frontend/
        ├── index.html       # App shell
        └── src/
            ├── main.js      # JavaScript logic
            └── app.css      # Styles
    ```

    ---

    ## 🛠️ Tech Stack

    | Component | Technology |
    |-----------|------------|
    | Backend | Go 1.25 |
    | Framework | Wails v2.12 |
    | Frontend | Vanilla HTML/CSS/JS |
    | Renderer | WebView2 (Edge) |
    | Build | Vite |

    ---

    ## � Roadmap

    - [ ] Auto-save every 30 seconds
    - [ ] Recent files list
    - [ ] Export to PDF
    - [ ] Custom themes
    - [ ] Word wrap toggle
    - [ ] Windows installer (NSIS)

    ---

    ## 👨‍💻 Author

    **Bilal Shemsu**

    [GitHub](https://github.com/bilalshemsu9) · [Email](mailto:biltinsa123@gmail.com)

    ---

    ## 📄 License

    MIT License — free to use, modify, and distribute.

    ---

    <div align="center">

    Built with ❤️ using [Go](https://golang.org) and [Wails](https://wails.io)

    </div>