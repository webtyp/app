# WebTyp

<img src="docs/img/slider.svg" width="100%" alt="WebTyp: an application built with it, and the webtyp dev TUI (build, MCP and shortcuts tabs)">

**Build full-stack web applications with Go.**
One language. Typed components. Simple architecture. Human or AI.

> ⚠️ **Active development.** APIs are still being simplified; expect rough edges.
> Source code is maintained privately — this repository distributes official binaries.

---

## What is WebTyp?

Modern web development fails under layers — frameworks, bundlers, transpilers, configuration.
**WebTyp does not try to win by adding tools. It tries to win by removing them.**

- **One language** — backend, frontend and shared logic in Go. You don't write JavaScript.
- **Typed components** — a component is a Go type; the compiler checks how components fit together.
- **Simple architecture** — the project layout is the configuration. Conventions, without magic.
- **Human or AI** — the TUI (`webtyp dev`) and the MCP daemon (`webtyp mcp`) drive the same project, so a developer or an agent builds inside the same structure. → [`api-design` skill](https://github.com/webtyp/devskills/blob/main/skills/api-design/SKILL.md)

Built first for **freelancers, agencies and small teams** delivering PWAs and business applications.

**Is it for you?** The problem it solves, who it's for, what works today and its honest limits: **[docs/WHY.md](docs/WHY.md)**.

---

## Installation

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/webtyp/installer/main/scripts/install.sh | bash
```

### Windows (PowerShell as Admin)

```powershell
irm https://raw.githubusercontent.com/webtyp/installer/main/scripts/install.ps1 | iex
```

The installer downloads the latest binary for your platform, verifies its SHA256 checksum, and places it on your `PATH`. Go and TinyGo are installed automatically — no prerequisites needed.

**Manual download:** [Releases](https://github.com/webtyp/app/releases) — Linux amd64/arm64, macOS arm64/amd64, Windows amd64, plus `checksums.txt`.

### Updating

`webtyp dev` checks for a new release in the background each time it starts. When one is found it is downloaded, verified and installed, and the **UPDATE** tab offers **Restart with vX** — nothing restarts until you press it. Disable the check with `-no-update` or `WEBTYP_NO_UPDATE=1`.

To update from the command line instead: `webtyp update`.

## Quick start

```bash
mkdir myapp && cd myapp
webtyp dev
```

WebTyp will scaffold a new project, start the dev server on `https://localhost:8080`, launch the TUI with live logs, open Chrome with auto-reload, and start the MCP daemon on `http://localhost:6060/mcp`.

## Commands

| Command | What it does |
|---------|--------------|
| `webtyp` | Show help |
| `webtyp dev` | Run the interactive TUI (starts the background daemon if needed) |
| `webtyp mcp` | Run the global headless MCP daemon (LLM / IDE entry point) |
| `webtyp test [args…]` | Run the project test suite |
| `webtyp status` | Print the status of the running daemon |
| `webtyp stop` | Stop the running daemon |
| `webtyp update` | Update webtyp in place to the latest release |

Flags (after the subcommand): `-debug` for verbose, unfiltered logs; `-no-update` to skip the update check for that run.

| Port | Used by | Override |
|------|---------|----------|
| `6060` | MCP daemon (MCP + SSE): `http://localhost:6060/mcp` | `WEBTYP_MCP_PORT` |
| `8080` | Dev server, per project: `https://localhost:8080` | `PORT` |

An agent drives WebTyp through the MCP daemon (`webtyp mcp`); the TUI (`webtyp dev`) is for human collaboration.

---

## Documentation

| Document | What's inside |
|----------|---------------|
| [docs/WHY.md](docs/WHY.md) | The problem it solves, the four promises, who it's for, what works today and its limits |
| [`api-design` skill](https://github.com/webtyp/devskills/blob/main/skills/api-design/SKILL.md) | The typed, explicit API approach and why it's an advantage |

---

## Support

- **Issues**: [github.com/webtyp/app/issues](https://github.com/webtyp/app/issues)

## License

MIT
