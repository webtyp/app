# Why WebTyp

**Build full-stack web applications with Go.**
One language. Typed components. Simple architecture. Human or AI.

---

## The problem is accumulated complexity

Modern web development rarely fails for lack of features. It fails under layers: a UI framework, a
meta-framework, a bundler, a transpiler, a package manager, a type layer on top of JavaScript,
server components versus client components, and a configuration file for each. Every layer is
reasonable on its own. Together they are a whole ecosystem to learn before the first useful screen.

Frameworks like Next.js solve a great deal, and they do it well. WebTyp takes a different bet:

> **WebTyp does not try to win by adding tools. It tries to win by removing them.**

Go is the mechanism, not the pitch. What WebTyp sells is fewer decisions.

---

## Four promises

### 1. One language

Backend, frontend and shared logic are written in Go. Types, models and validation are declared
once and used on both sides.

The application developer does not write JavaScript. That is a promise about **your code**, not
about the browser: WebTyp uses the small amount of JavaScript the platform itself requires, so you
don't have to.

### 2. Typed components and sustainable CSS

A component is a Go type: its inputs are fields, its output is a DOM tree, its events are typed
handlers. Components are genuinely separate units — each one owns its structure, style and
behavior — and the compiler checks how they fit together. No JSX, no hooks, no client/server
component split to reason about.

**CSS without the chaos.** Today's web development treats CSS as either an unmaintainable cascade of
magic numbers and arbitrary strings, or thousands of repetitive utility classes pasted into markup.
WebTyp rejects both. In WebTyp, stylesheets are typed Go code governed by `widget/style`:
- **No raw strings or magic numbers:** Spacing, colors, surfaces and typography come strictly from design tokens.
- **Semantic recipes over ad-hoc assembly:** Common interaction surfaces (`Button`, `Stack`, `SlideDeck`) are composed once at the root, guaranteed accessible and consistent.
- **Compile-time enforcement:** If a style recipe doesn't exist, it's treated as a defect in the design system, not an excuse to write garbage CSS at the call site. The compiler and automated conformance suites ensure CSS remains DRY, clean, and sustainable over years of maintenance.

### 3. Simple architecture

Convention over configuration, **without magic**. The project layout is the configuration: there
is no build config to write and no bundler to tune. Every convention is explicit, readable Go that
you can open and follow.

Every new feature pays a *complexity tax*: how many new concepts, files and settings does it ask
of you? If the answer is more than one of each, it needs a very good reason to exist.

### 4. Human or AI

The same architecture serves a developer at the keyboard and an agent working through
[MCP](https://modelcontextprotocol.io). The TUI (`webtyp dev`) and the MCP daemon (`webtyp mcp`)
drive the same project with the same operations — not two systems that drift apart.

The API is designed so that **a junior developer can build without AI**. Code that is simple enough
for a person to get right is also code an agent gets right: the typed signatures guide it, and the
compiler rejects what is wrong. → [`api-design` skill](https://github.com/webtyp/devskills/blob/main/skills/api-design/SKILL.md)

---

## Who it's for

**First and foremost: freelancers, agencies and small teams (1–5 developers)** who deliver web
applications to clients and would rather ship the application than assemble a JavaScript
toolchain for it.

What WebTyp aims to do especially well:

- **Progressive Web Apps** — installable, fast, able to work offline.
- **Business applications** — admin systems, dashboards, internal tools, forms and CRUD over a
  real database.
- **Small SaaS products** — one codebase from the model to the deployment.

**Good fit if you:**

- want one language and one mental model for the whole application;
- value fewer, stronger conventions over unlimited flexibility;
- work with AI assistants and want them to build inside a structure, not invent one each time;
- would rather read Go than configure tools.

**Not a fit if you:**

- need React, Vue, Svelte or Angular — WebTyp will not support them, by design;
- depend on the npm ecosystem or complex frontend pipelines (SASS, PostCSS, module federation);
- need a mature ecosystem with an answer on every forum today.

---

## What works today

- **`webtyp dev`** — a TUI with live logs, hot reload, and a managed browser that reloads on change.
- **`webtyp mcp`** — an MCP daemon any agent can connect to (Claude, Cursor, Copilot…) to see the UI,
  read logs, inspect the browser, query the database and control compilation.
- **Components verified in memory** — a web component can be rendered and checked without writing
  a single file to disk.
- **Database sync** — models and schema kept in step during development.
- **Deploy to Cloudflare Workers** — the path from project to production.
- **Self-update** — new releases are installed in the background; you choose when to restart.

## Be honest before adopting ❌

- **Active development.** APIs are still being simplified; expect breaking changes and rough edges.
- **Young ecosystem.** Fewer components, libraries, tutorials and answers than any established
  JavaScript framework. You may hit a missing piece.
- **Opinionated on purpose.** WebTyp has one way to do things. Contributions must follow that
  architecture — that is what keeps it simple.
- **Go in the browser has limits.** WebAssembly has real constraints; a few performance-critical
  browser APIs are better served by a small piece of JavaScript, which WebTyp keeps inside the
  framework rather than in your code.
- **One deployment target today.** Cloudflare Workers is supported now; others are not yet.

---

## The measure of success

WebTyp succeeds when someone other than its author can build a real application **without having
to ask how it works inside** — and when developers earn a living building software with it.
