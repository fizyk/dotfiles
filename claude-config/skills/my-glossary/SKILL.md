---
name: my-glossary
description: Build or check the repository's glossary (docs/glossary.md unless its rules name another path), its domain vocabulary - one entry per domain noun, with its definition, the canonical identifier carrying it, and where that is defined. Harvests from the surfaces each component declares, and proposes surfaces for the user to confirm where none are declared. Use when the user asks to generate, update or check the glossary, or to find out whether a name coins a synonym for a term already in use.
allowed-tools: Read, Write, Edit, Glob, Grep, Bash(git diff *), Bash(git log *), Bash(git status *), Bash(git rev-parse *)
---

# Domain glossary

The procedure is portable, the glossary is not: what lands in the glossary is
this repository's own vocabulary, and another repository running this skill
generates its own. What makes a name a good one is under "Naming" in
`~/.claude/CLAUDE.md` and the repository's own rules - read it there rather
than a second copy here.

**The glossary** is the file the repository's rules name. Where none is named,
look for a `glossary.md` under `docs/`. Where none exists, propose
`docs/glossary.md` and confirm the path before building.

One glossary covers the whole repository. A domain noun crossing the boundary
between two components is exactly where a synonym does damage, so the file is
repository-wide even though the components it harvests from are not.

## 1. Pick the mode

No glossary yet: **build** it, sections 3 and 4. One already
exists: **check** it, section 5. Never regenerate a glossary that exists. A
rewrite throws away wording the team has settled on, and what a later run is
for is the difference, not the file.

## 2. Resolve the sources

Read `~/.claude/CLAUDE.md`, the repository's root `CLAUDE.md` or `AGENTS.md`,
then the `CLAUDE.md` of every component. A component's own file names the
surfaces it publishes and the layers it spells terms out in, and those
surfaces are what this skill harvests. **Use `Read`, not `cat`.** A
repository with no component files of its own has one component: the root.

Harvest only from a component whose surfaces are declared. Where one declares
none, **propose its surfaces from its stack and wait for the user to confirm
them**; harvest nothing from it until they do. A stack nobody has described is
how a glossary fills up with mechanism nouns: the framework types and
transports that section 3 exists to keep out. Where the user declines, name
the component in the report as not harvested.

A component whose file defers to another's harvests from the inventory it
defers to, plus whatever its own file adds.

## 3. Harvest the terms

Read a component's declared surfaces in this order. When two disagree on what a
term means, the earlier one wins.

1. **The request and response shapes.** The types the component publishes over
   the wire, and the prose in every field description. These are the terms the
   interface has already committed to, so they outrank everything below.
2. **The closed value sets**, where each member is itself a term.
3. **The computed quantities**, named once in a query or a calculation and
   nowhere else.
4. **The defining text of each domain package or module.** One line each,
   already stating what a domain is in domain words.
5. **Settings field names, log messages, and the keys of their structured
   fields.**
6. **The interface copy**, where a component declares one. What the product
   calls a thing to the person using it is the term with the strongest claim
   on being canonical, and the one most likely to disagree with the backend.

A noun earns an entry when a conversation about the project would use it. Skip
what names a mechanism rather than a domain thing: framework types,
transports, repositories, mixins, clients, test doubles.

## 4. Write the entries

Group the terms by the area they belong to, and give each one a row:

* **Term** - the domain noun, in prose form (`retry delay`, not
  `capped_retry_delay_seconds`).
* **Definition** - one sentence, stating what the thing is. Not why it exists,
  not how it is computed.
* **Canonical identifier** - the spelling every layer is expected to use. Where
  an abbreviation is confined to one layer, name both and say which is which.
  Where two components spell one term differently, this is the column that
  settles it.
* **Where defined** - the path that carries the definition, and the symbol
  inside it.

Two things deserve their own section at the end, because they are what a
reviewer needs and a table of terms hides:

* A term the repository spells two ways, with which spelling belongs where.
  Across two components counts, and is the more important half.
* A word carrying two unrelated meanings, with both stated.

Nothing else goes in the file. No rationale, no history, no open questions.

The file follows the Markdown rules of wherever it lives. **Write a path in
the "Where defined" column as inline code, never as a Markdown link.** A docs
site that builds the file treats a link to a source file as a broken internal
link.

## 5. Check an existing glossary

1. Find what has changed since the glossary last did, from three sources.
   `git log -1 --format=%H -- <glossary path>` names the baseline
   commit, so `git diff <sha>..HEAD` is the committed range. `git diff HEAD`
   adds the tracked changes not committed yet, staged and unstaged alike; plain
   `git diff` would show only the unstaged half.
   `git status --porcelain --untracked-files=all` then names the untracked
   files, which no diff shows at all: read each one in full. The flag is not
   optional - the default collapses an untracked directory into one entry, so a
   whole new domain folder would arrive as a single path instead of the files in
   it. An empty committed range still leaves the other two sources. An empty
   baseline means the glossary has never been committed, so there is no range:
   harvest section 3 again and check the file against what that returns.
2. Read against the glossary every identifier the diff adds or renames, and
   every existing one whose use or meaning it changes. A declaration the diff
   leaves alone still counts when what it now holds, or what populates it,
   moved: that is the only way the contradiction below can surface. Sort each
   into one of four:
   * **New term** - a domain noun with no entry. Propose one.
   * **Synonym** - a new spelling for a term already in the glossary. Report it
     with both spellings and the existing canonical one. This is the finding
     the skill exists for, and a synonym introduced in a second component is
     its sharpest case.
   * **Contradiction** - an identifier used in a way the entry's definition does
     not cover. Report the entry and the use.
   * **Consistent** - no report.
3. A term whose last use the diff removes is a report, not an edit. The entry
   stays until someone confirms the term is gone.

## 6. Report

List the findings grouped as in step 5, each naming `file:line`. Then, kept
separate, the components that declared no surfaces and were therefore not
harvested. Close with both counts and offer to apply the findings.

**Write to the glossary only after the user accepts** - on a
check run the file is an input.
