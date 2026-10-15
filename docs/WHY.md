# Why WebTyp

**English** · [Español](WHY.es.md)

**Software you cultivate, not software that piles up.**

Full-stack web applications in Go. One language. Typed components. Simple architecture.
For humans and for AI.

---

## Building software is tending a garden

Building software was always hard. But what is building it? Writing code in some language, inside
text files? And if I write it faster, does my business win because I ship new features in minutes?

Yes and no.

Building software is a continuous process. It is like tending a garden: it is never finished. If I
plant any seed without asking what it is or how it will affect the plants already there, over time
the garden fills with weeds and rots. But if I know every plant and why it sits where it does, the
garden can keep growing, and I can look at it and say: *it is like this because I tended it*.

Planting was never the hard part. The hard part is **planning, maintaining and caring for it over
the long run**.

## What AI changed, and what it didn't

AI writes software fast. But it optimizes the task in front of it; nobody but you answers for the
software that still has to work two years from now. **You are always the one responsible.**

That did not change. So there are, as always, two paths:

- **Explore:** play, try tools, experiment.
- **Build to last:** turn that code into something profitable, that creates value and improves
  people's lives.

Both are valid. **WebTyp bets on the second.**

## One right way, enforced by the compiler

If I am responsible for what gets written, I must be able to make sure it is correct. But when AI
writes at that speed, reviewing everything becomes impossible.

The known answer is code that humans and AI can both read: types, good architecture, single
responsibility, modularity, decoupling, DRY, tests, documentation. The problem is that this is too
many instructions. The agent's context is spent explaining *how* you want it to build, and nothing
forces it to follow the letter.

That is where WebTyp comes in. It proposes a **typed, strict** way to build:

- **There are not two ways to do the same thing, only one.** The decisions are already made in
  the framework.
- **If you don't follow them, it doesn't compile**, and the error explains why and how to fix it.
- **You build with pieces already proven**, yours or the community's, each verified on its own.
- **Everything is a module with a single responsibility.** WebTyp doesn't suggest it: it
  requires it.

> **Vision:** software you cultivate, not software that piles up.
>
> **Mission:** make the right way to build the only way that compiles.

---

## The problem is accumulated complexity

Modern web development rarely fails for lack of features. It fails under layers: a UI framework,
a meta-framework, a bundler, a transpiler, a package manager, a type layer on top of JavaScript,
server components versus client components, and a configuration file for each. Every layer is
reasonable on its own. Together they are a whole ecosystem to learn before the first useful screen.

Frameworks like Next.js solve a great deal, and they do it well. WebTyp takes a different bet:

> **WebTyp does not try to win by adding tools. It tries to win by removing them.**

## Four promises

### 1. One language

Backend, frontend and shared logic are written in Go. Types, models and validation are declared
once ([`model`](https://github.com/webtyp/model), [`ormc`](https://github.com/webtyp/ormc)) and used
on both sides.

You don't write JavaScript. That is a promise about **your code**, not about the browser: WebTyp
uses the small amount of JavaScript the platform itself requires, so you don't have to.

### 2. Typed components and sustainable CSS

A component ([`components`](https://github.com/webtyp/components)) is a Go type: its inputs are
fields, its output is a DOM tree ([`dom`](https://github.com/webtyp/dom)), its events are typed
handlers. Each component owns its structure, style and behavior, and the compiler checks how they
fit together. No JSX, no hooks, no client/server component split to reason about.

CSS is typed Go code too ([`widget/style`](https://github.com/webtyp/widget/tree/main/style)):
spacing, colors and typography come from *design tokens*, not from magic numbers or hundreds of
utility classes pasted into markup. A missing style recipe is a defect in the design system, not
an excuse to improvise.

### 3. Simple architecture

Convention over configuration, **without magic**. The project layout is the configuration: there
is no build config to write and no bundler to tune. Every convention is explicit Go you can open
and follow. → [`project-layout` skill](https://github.com/webtyp/devskills/blob/main/skills/project-layout/SKILL.md)
· [`app-modules` skill](https://github.com/webtyp/devskills/blob/main/skills/app-modules/SKILL.md)

Every new feature pays a *complexity tax*: how many new concepts, files and settings does it ask
of you? If the answer is more than one of each, it needs a very good reason to exist.

### 4. Human or AI

The same architecture serves a developer at the keyboard and an agent working through
[MCP](https://modelcontextprotocol.io). The TUI (`webtyp dev`) and the MCP daemon (`webtyp mcp`,
built on [`mcp`](https://github.com/webtyp/mcp)) drive the same project with the same operations,
not two systems that drift apart.

The API is designed so that **a junior developer can build without AI**. Code simple enough for a
person to get right is also code an agent gets right: the typed signatures guide it, and the
compiler rejects what is wrong. → [`api-design` skill](https://github.com/webtyp/devskills/blob/main/skills/api-design/SKILL.md)

---

## Why Go?

The choice was made on four criteria:

- **Fast compilation.** WebTyp is full-stack and builds interfaces with web technologies: every
  change must show up in the browser within seconds.
- **Small binaries for the browser.** Frontend code ships to the client as WebAssembly compiled
  with [TinyGo](https://tinygo.org); every kilobyte counts, which is why there are lightweight
  replacements for the standard library ([`fmt`](https://github.com/webtyp/fmt),
  [`json`](https://github.com/webtyp/json), [`time`](https://github.com/webtyp/time)).
- **Static types.** An error should appear while writing, not while running.
- **Few ways to do the same thing.** Go's simple syntax fits the "one way" thesis.

**Why not Rust?** It is excellent, but it compiles slowly, and that wait breaks the feedback loop
of building a web interface. Its expressiveness also opens more paths to solve the same problem.

**Why not Python?** Its typing is optional: the error arrives at runtime, not while writing. And
shipping its runtime and libraries to the browser means a lot of binary for the client.

**Why not TypeScript?** Because it is exactly the problem WebTyp wants to remove: more tools on top
of JavaScript and, in practice, two ways to do the same thing, one on the frontend and another on
the backend.

---

## Who it's for

**First and foremost: freelancers, agencies and small teams (1–5 developers)** who deliver web
applications to clients and would rather ship the application than assemble a JavaScript
toolchain for it.

What WebTyp aims to do especially well:

- **Progressive Web Apps:** installable, fast, able to work offline.
- **Business applications:** admin systems, dashboards, internal tools, forms and CRUD over a
  real database.
- **Small SaaS products:** one codebase from the model to the deployment.

**Good fit if you:**

- want one language and one mental model for the whole application;
- value fewer, stronger conventions over unlimited flexibility;
- work with AI assistants and want them to build inside a structure, not invent one each time;
- would rather read Go than configure tools.

**Not a fit if you:**

- need React, Vue, Svelte or Angular: WebTyp will not support them, by design;
- depend on the npm ecosystem or complex frontend pipelines (SASS, PostCSS, module federation);
- need a mature ecosystem with an answer on every forum today.

## What works today

- **`webtyp dev`:** a TUI with live logs, hot reload, and a managed browser
  ([`devbrowser`](https://github.com/webtyp/devbrowser)) that reloads on change.
- **`webtyp mcp`:** an MCP daemon any agent can connect to (Claude, Cursor, Copilot…) to see the UI,
  read logs, inspect the browser, query the database ([`sqlmcp`](https://github.com/webtyp/sqlmcp))
  and control compilation.
- **Components verified in memory:** a web component can be rendered and checked without writing
  a single file to disk ([`sitec`](https://github.com/webtyp/sitec)).
- **Database sync:** models and schema kept in step during development
  ([`orm`](https://github.com/webtyp/orm), [`ddl`](https://github.com/webtyp/ddl)).
- **Deploy to Cloudflare Workers:** the path from project to production
  ([`goflare`](https://github.com/webtyp/goflare); working example: [`goflare-demo`](https://github.com/webtyp/goflare-demo)).
- **Self-update:** new releases are installed in the background; you choose when to restart
  ([`update`](https://github.com/webtyp/update)).

## Be honest before adopting ❌

- **Active development, by one person.** Today I am the only developer
  ([@cdvelop](https://github.com/cdvelop)). APIs are still being simplified; expect breaking
  changes and rough edges.
- **Proven on a real project before going public.** Instead of releasing early and broken, I am
  building two projects at once: the framework and the application that consumes it, the new
  version of that clinic platform, with several modules. Every bug found there fixes the framework
  before it reaches you.
- **Young ecosystem.** Fewer components, libraries, tutorials and answers than any established
  JavaScript framework. You may hit a missing piece.
- **Opinionated on purpose.** WebTyp has one way to do things. Contributions must follow that
  architecture; that is what keeps it simple.
- **Go in the browser has limits.** WebAssembly has real constraints; a few performance-critical
  browser APIs are better served by a small piece of JavaScript, which WebTyp keeps inside the
  framework rather than in your code.
- **One deployment target for now:** Cloudflare Workers.

---

## Where it comes from

Years ago I built a platform for a clinic: a Go backend, an HTML, CSS and JavaScript frontend, no
framework. Today, in 2026, it still runs every day.

But its source code no longer exists. I wasn't using git, and only the executable is left. Now my
clients need more: a website, agents that answer the usual questions about services, prices and
doctors. I can't grow a garden whose map I lost.

That was the lesson WebTyp came from. Working today is not enough: software has to be possible to
understand, care for and grow tomorrow. To rebuild that platform, this time with Go in the browser
too, I started building the framework I wish I'd had.

It went through several names while it found its shape: **golite**, **godev**, **tinywasm** and,
finally, **webtyp**: the *typed web*. Each name was a version of the same question: how do you
build web software that lasts?

## The measure of success

WebTyp succeeds when someone other than its author can build a real application **without having
to ask how it works inside**, and when developers earn a living building software with it.

**Want to know when it's ready?** Follow releases at
**[github.com/webtyp/app](https://github.com/webtyp/app)**: *Watch → Custom → Releases*.
