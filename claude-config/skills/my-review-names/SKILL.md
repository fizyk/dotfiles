---
name: my-review-names
description: Review the identifier names a diff introduces against the naming rules - abbreviation level, a type in the name, a parameter that disagrees with its call site, plain English and the repository's spelling variant, a comment standing in for a name or restating the code beside it, and the documentation of a published library's public API or, where the repository asks for one, the defining comment. Resolves the rules from the personal CLAUDE.md, the repository's own instructions and each component's. $ARGUMENTS optionally names the base ref. Use when the user asks to review naming, check the names on a branch, or check names before a pull request.
allowed-tools: Read, Grep, Glob, Bash(git diff *), Bash(git log *), Bash(git status *), Bash(git branch *), Bash(git show *), Bash(git rev-parse *)
---

# Identifier names

A review pass over a diff, not a linter: it checks what ruff and eslint
cannot. The rules are under "Naming" in the files section 2 resolves, and the
vocabulary is in the repository's glossary - read them there rather than a
second copy here. Only the procedure that applies them lives in this file,
which is what lets any repository, and any component, run the same skill
against its own rules.

Read-only. Report the findings, then offer to apply them.

## 1. Scope the diff

**`$ARGUMENTS` names a base ref** (a branch, tag, or SHA): resolve it with
`git rev-parse --verify --end-of-options "<ref>^{commit}"`. Unresolvable, say
so and ask for one that resolves.

**`$ARGUMENTS` is empty:** take the default branch from `git rev-parse
--abbrev-ref origin/HEAD`, falling back to `git branch --list
--format='%(refname:short)' main` - the plain listing pads the name with the
marker column, which no ref resolves through. Name the base ref at the top of
the report so the reader can correct it.

Then read four sources:

| Source | Command |
| --- | --- |
| Committed | `git diff <base-ref>...HEAD` - three dots, the merge base rather than the branch tip |
| Staged | `git diff --staged` |
| Unstaged | `git diff` |
| Untracked | `git status --porcelain --untracked-files=all`, then read each new file in full - the default collapses an untracked directory into one entry, hiding the files in a new folder |

"Nothing to review" holds only when all four come back empty. An empty
committed range still leaves the other three.

Read the whole file whenever a hunk lands in code you do not already
understand. A name is judged by what it holds, which the hunk alone rarely
shows.

## 2. Resolve the rules

Resolve each changed path to its component. **A component is a directory with
its own `CLAUDE.md`**: the path's component is the nearest ancestor directory
holding one, and its rules are that file plus every `CLAUDE.md` above it up to
the repository root. A path with none between it and the root takes the root's
rules alone.

A directory that plainly holds a component's source and has no `CLAUDE.md` of
its own is not the root's. Name it, and report the missing file the way an
undeclared component is reported below.

Read, in this order. A later source wins where two disagree.

1. `~/.claude/CLAUDE.md`, "Naming". The personal defaults, true of every
   repository.
2. The repository's own rules: the root `CLAUDE.md` or `AGENTS.md`, and the
   naming section of `CONTRIBUTING.md` where there is one.
3. The `CLAUDE.md` of each component the diff touches, outermost first where
   one nests inside another. It carries that component's abbreviation layer,
   the surfaces it adds to section 3, and any exemption from a root rule.

**Use `Read`, not `cat`.** Reading a file under a component is also what loads
that component's own instructions for the session.

A repository with no component files of its own has one component: the root.
The personal defaults are its rules. That is the usual case, not a gap:
report it once in the header and review as normal.

**The glossary** is the file the repository's rules name. Where none is named,
look for a `glossary.md` under `docs/`. Where none exists, the vocabulary
checks in section 4 say so and skip.

A heading is not a declaration here either: a section that records the gap and
lists candidate exemptions for whoever will settle them counts as declaring
nothing. Take it at its word.

A component in a repository that does declare components, and has declared
nothing itself, is not a reason to stop. The root rules are language-neutral
and still apply. What it does mean:

* No exemption is honored, so a root rule that the component's stack
  contradicts fires as a finding against the whole stack rather than once.
* No abbreviation layer is declared, so every short form is a finding.

Both read as noise rather than as review. So where a component declares
nothing, **say so once at the top of its findings**, name the root rules that
its stack looks likely to contradict, and offer the section it needs. Then
report the findings anyway, marked as provisional. Never drop them silently,
and never invent the exemption yourself.

## 3. Collect the names

Every identifier the diff adds or renames, in source and tests alike:

* Modules, packages, classes, functions, methods, and types.
* Parameters, attributes, properties, and every local, however short-lived.
* Model and schema fields and their aliases, settings fields, enum members.
* The wording of a log message, the keys of its structured fields, and metric
  names.
* Anything further the component's own `CLAUDE.md` names, such as the `AS`
  aliases and CTE names in its SQL, or a file naming convention it enforces.

A line that only moves an existing name carries no new name. Skip it.

Then the comments the diff adds inside a body or beside a statement. A
defining comment belongs to "The defining comment", and test-structure labels
and lint or type-checker directives are not comments here.

## 4. Run the checks

Run the first five against every name and the sixth against the comments
section 3 collects. A check returns a finding or nothing.

### Abbreviation level

The layer a name sits in decides its spelling, and the component's
`CLAUDE.md` says which layer, if any, is allowed short forms. Look the term up
in the glossary: an entry that carries two spellings also says which one
belongs where. Two findings, opposite directions:

* A short form outside that layer, where another layer reads it: a schema or
  model field, a public signature, a settings field, a log message, or one of
  its structured keys. Propose the spelled-out form.
* A spelled-out form inside it, where it stops the code around it reading as
  the calculation it is. Propose the short form.

Where the component grants no layer short forms, only the first direction
applies and every abbreviation is a finding. Where the component declares
nothing at all, section 2's provisional rule holds.

A term with no glossary entry at all is the `my-glossary` skill's finding. Name it
as one and move on; do not invent an entry here.

**Where the repository has no glossary**, this check cannot run.
Say so once, report it as the `my-glossary` skill's to build, and skip the check
for the whole diff. Reporting every term as missing an entry buries the
findings that matter.

### A type or a shape in the name

A suffix naming what the declaration already says: `_float`, `_dict`, `_list`,
`_str`, `_int`, `_bool`, `_set`, `_tuple`, and the prefix forms (`dict_of_`,
`list_of_`). Propose the name without it.

Where dropping the suffix leaves a name that no longer says what the value is,
the suffix was hiding a naming problem. Propose the name the value deserves
instead of the shorter one.

A unit or a representation is not a type: `_days`, `_rate`, `_count`, `_pct`,
`_date` say what the number means. Leave them.

A convention that marks a declaration's kind rather than a value's container
is a different thing, and a component may require one. Check its `CLAUDE.md`
before reporting a prefix or a suffix that every declaration of that kind
carries: a scheme a linter enforces is not this check's finding.

### A parameter against its call site

A call comes under this check from either side:

* **The callee.** A function the diff adds, or whose signature it changes.
  `grep` its name across the component. Where that turns up an aliased import,
  grep the alias as well before calling the call sites covered: that import
  line is the only place the two names meet.
* **The caller.** A call site whose argument expressions the diff changes, even
  where the callee and its signature are untouched. The diff already names it;
  read the callee's signature to bind against. This is the half a rename most
  often leaves behind: updating a field's name at its call sites is mechanical,
  and the parameter it binds to is not.

Either way, bind every argument to the parameter it fills, by position as well
as by keyword. Positional arguments bind in declaration order, skipping the
receiver; a keyword-only marker means everything after it can arrive only by
keyword.

Each bound pair gives two names to compare. The caller's side is whatever name
the argument expression carries itself:

* A bare variable: its own name.
* A dotted path: the last segment, so `source.right` reads as `right`.
* A call: the name of the function called.
* A literal, a subscript, or an arithmetic expression carries no name of its
  own. There is nothing to compare, so skip the pair.
* A spread or an expansion names the container, not the parameters it fills, so
  skip it rather than binding that name to one. The exception is an expansion
  whose contents are written out at the call site, where each key binds as its
  own argument.

Where the two names say the same thing two ways, that is the finding - usually
a rename that stopped at the value and left the parameter behind.
`f(left=source.right)` and `f(source.right)` are the same finding; the
positional form hides it from the reader, so it matters more there.

Decide which of the two names is right: the one the glossary backs, or the
caller's where the callee's is the leftover. Propose the rename on the other
one.

### Plain English and the spelling variant

The variant is the one the rules section 2 read name. Where none names one, it
is the one the repository's existing identifiers and docs use, and where those
mix both, say so once rather than reporting every name.

Applied to each new name in turn, not left as a bullet to skim. For each: is
there a more common English word that fits as well? Is the spelling the
repository's variant? Report the name with the plainer word, or the other
spelling, as the proposal.

### The defining comment

Which rule holds depends on what the thing is.

**The public API of a published library**: a repository that ships a package
for others to install, and the names it exports to them. Every public module,
class, and function the diff adds or changes carries documentation for its
user: what it does, its parameters, what it returns, and what it raises. A
missing or stale one is a finding. Follow the docstring style the project
already uses. This holds even where the repository declares a definition-comment
rule, which then covers only what is not exported.

**Everything else**: run the rest of this check only where the rules section 2
read carry a definition-comment rule, such as "a docstring defines, or there is
none". Without one, comment density follows the surrounding code, and a
docstring explaining why is no finding. Name the check as skipped in the
report.

Where the rule is declared, for every module, class, type, and function the diff
adds outside the public API:

* The name carries the concept and there is no comment: nothing to report.
* The name cannot carry the concept and there is no comment: propose one line
  saying what the thing is.
* A comment says why the thing exists, what it replaced, or how it works:
  report it. Propose the defining line, or removal where the name already
  carries the concept.
* A comment says more than what the thing is: report everything past the
  opening sentence. One line is the whole allowance, so a second defining
  sentence is as much a finding as a justifying one. Wrapping one sentence to
  the line limit is not a second line.

### A comment standing in for a name

For every comment section 3 collects:

* It says what a value, a block, or an expression is or holds: propose the name
  that carries it instead - a rename, or an extracted function or local for a
  block or an expression that has none - and removal of the comment.
* It restates what the code beside it already says: propose removal.
* It says something no name could, such as a constraint from outside the code:
  nothing to report.

## 5. Report

Open with the base ref, and one line per component the diff touches saying
which rule sources were read.

Then one table per component, ordered by file:

| `file:line` | Current | Proposed | Reason |
| --- | --- | --- | --- |

`Reason` is one clause naming the check that fired and what it found, not a
restatement of the rule. A name you weighed and passed gets no row. A comment
finding puts the comment in `Current`, and in `Proposed` the name that replaces
it, or removal. Mark a provisional finding, from a component that declared no rules,
as such.

Under each table, one line per check that found nothing, so a clean check reads
differently from a skipped one.

Close with two counts, kept apart: the findings, and the components that
declared no rules. The second is the more useful number on a first run, and
offer the sections those components need. Then offer to apply the findings.

**Change nothing until the user accepts.** When they do, rename every
occurrence of the name: the declaration, the call sites, the tests, the SQL, and
the structured log keys that carry it. A comment finding applies as its proposal
reads: the rename or the extraction first, then the removal of the comment.
