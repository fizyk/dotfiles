---
name: my-review-published-text
description: Review the text a diff publishes against the rules for published prose - API field and response descriptions, tool descriptions, interface copy, error response bodies, a library's public docstrings, warnings and option help, exception messages, and log messages. Reads each component's own surface inventory, and where none is declared proposes one for the user to confirm rather than guessing. $ARGUMENTS optionally names the base ref, or `all` to sweep every declared surface. Use when the user asks to review API descriptions, tool text, interface copy, docstrings, exception or log wording, or the wording before a pull request.
allowed-tools: Read, Grep, Glob, Bash(git diff *), Bash(git log *), Bash(git status *), Bash(git branch *), Bash(git show *), Bash(git rev-parse *)
---

# Published text

A review pass over the prose that leaves the process. The rules are under
"Published text" in the files section 2 resolves, and the vocabulary is in the
repository's glossary - read them there rather than a second copy here. Only
the procedure that applies them lives in this file, which is what lets any
repository, and any component, run the same skill against its own rules.

**Which strings are published is the component's to declare, not this skill's
to guess.** Section 2 resolves the rules and the inventory, and section 3
collects from it. Where nothing declares one, section 2 proposes an inventory
and the user confirms it before anything is checked.

Four audiences read this text, and they need different things: a caller or a
model, a person using the product, a developer using a library, and an
operator. The component names which of them it publishes to. Establish a
string's readers before judging it: the same sentence can be right for an
operator and wrong for a caller, and a validation message is read by both.

This skill judges sentences. The `my-review-names` skill judges identifiers, the
keys of a log record's structured fields among them, so a key that reads badly
is its finding rather than this one's. It also owns whether a public API
docstring exists; this skill judges the wording of one that does.

Read-only. Report the findings, then offer to apply the rewrites.

## 1. Scope the run

**`$ARGUMENTS` is `all`:** every surface every component declares, across the
whole repository. This is the first pass over a component the skill has not
seen before; later runs take a diff.

**`$ARGUMENTS` names a base ref** (a branch, tag, or SHA): resolve it with
`git rev-parse --verify --end-of-options "<ref>^{commit}"`. Unresolvable, say
so and ask for one that resolves.

**`$ARGUMENTS` is empty:** take the default branch from `git rev-parse
--abbrev-ref origin/HEAD`, falling back to `git branch --list
--format='%(refname:short)' main` - the plain listing pads the name with the
marker column, which no ref resolves through. Where neither returns a ref,
stop and ask for one. Name the base ref at the top of the report so the reader
can correct it.

A diff run reads four sources:

| Source | Command |
| --- | --- |
| Committed | `git diff <base-ref>...HEAD` - three dots, the merge base rather than the branch tip |
| Staged | `git diff --staged` |
| Unstaged | `git diff` |
| Untracked | `git status --porcelain --untracked-files=all`, then read each new file in full - the default collapses an untracked directory into one entry, hiding the files in a new folder |

"Nothing to review" holds only when all four come back empty. An empty
committed range still leaves the other three.

Read the whole file whenever a hunk lands in text you do not already
understand. Published text is judged against the thing it describes, which the
hunk alone rarely shows.

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

1. The global `CLAUDE.md` (in `$CLAUDE_CONFIG_DIR`, or `~/.claude` where that is
   unset), "Published text". The personal defaults, true of every repository.
2. The repository's own rules: the root `CLAUDE.md` or `AGENTS.md`, and
   `CONTRIBUTING.md` where there is one.
3. The `CLAUDE.md` of each component the diff touches, outermost first where
   one nests inside another. It carries that component's surface inventory,
   which audience reads each surface, and any exemption from a root rule. A
   component whose file defers to another component's means reading both: the
   one it defers to for the inventory, its own for what differs.

**Use `Read`, not `cat`.** Reading a file under a component is also what loads
that component's own instructions for the session.

A repository with no component files of its own has one component: the root.

**The glossary** is the file the repository's rules name. Where none is named,
look for a `glossary.md` under `docs/`.

**A heading is not a declaration.** What matters is whether something names
the surfaces, not whether the section exists. A section that records the gap
instead of filling it - saying it has not declared its surfaces, and listing
candidates for whoever will - counts as **undeclared**. Take it at its word and
treat the candidates as context for the report, never as the inventory.

**A component that has not declared its surfaces cannot be checked as it
stands.** Unlike naming, where the rules apply to any language unaided, this
skill has nothing to collect until something names what publishes: the
surfaces are where the stack shows through, and guessing them means either
inventing findings or missing the whole surface. So for such a component:

* Read its stack and **propose an inventory**: each surface, the construct
  that carries it, and which audience reads it. Show it and stop.
* Once the user confirms it, check against the confirmed inventory for this
  run and say in the report that it came from the conversation, not a file.
* Offer to write it into the component's `CLAUDE.md`, or into
  `CLAUDE.local.md` in a repository the user does not own.
* Do not fall back to another component's inventory. A backend's surfaces say
  nothing about a TypeScript component, and a clean pass against the wrong
  inventory is worse than no pass.

Where the user declines, report the component as **unchecked**, with the
count of strings the diff changed under it so the size of the gap is visible.
Never report it as clean.

## 3. Collect the text

Take the inventory from the component's own `CLAUDE.md`: each surface it
declares, and the construct that carries it. What to collect from those
constructs depends on the run section 1 scoped:

* **A full sweep:** every string they hold, read from the files on disk. There
  is no diff to narrow it by.
* **A diff run:** the strings the diff adds or changes, and for each untracked
  file every string it holds. An untracked file appears in no diff, so it is
  read from disk in full the way section 1 says.

Two things hold whatever the stack, and are this skill's to apply rather than
the component's to restate.

**A description's text is whatever the reader sees once the declaration
resolves, which is often not where the field is.** Follow it:

* The description written on the field itself.
* A shared alias carrying one description, published at every field using it.
  One text, so judge it once and name the uses.
* No description of its own, and a declaration naming another type: that type's
  own defining text is the field's text. Follow the declaration and read it.
* Both, which the schema publishes side by side. Read them as one text.

**A string that never leaves the process is out of scope**, whatever the
inventory says: comments, the defining text of an internal class the schema
does not publish, test data and fixtures, and SQL.

On a diff run, a hunk that only moves a string carries no new text. Skip it.

## 4. Run the checks

Three checks apply to every string, whoever reads it. The rest apply by
audience, and which audience a surface has is what the component declared. A
check returns a finding or nothing.

### Sentence shape

A complete sentence, capitalised first word, terminal period. A noun fragment
(`Number of retries before failing`) is a finding; propose the sentence it was
shortened from.

Some strings are names rather than prose, and are exempt: an application title,
a route summary, an example's summary label, and a short interface label whose
job is to fit a control. A description sitting beside any of them is prose and
stays in the check. Where a component's copy is mostly such labels, it says so
in its own `CLAUDE.md`; absent that, judge by whether the string is a sentence
in the interface or a word on a button.

A message built from adjacent string literals is one sentence. Judge the joined
text, and check each join keeps the space the wrapping dropped.

### Vocabulary

Look every domain noun up in the repository's glossary.

* A spelling the glossary does not carry, for a term it does: propose the
  canonical one.
* An abbreviation the text has not expanded: the finding is the first use, not
  every use. Propose the expansion with the short form in parentheses and
  leave the uses after it alone. The layer the file sits in changes nothing
  here.
* A term with no glossary entry at all is the `my-glossary` skill's finding. Name
  it as one and move on; do not invent an entry here.

**Where the repository has no glossary**, this check cannot run.
Say so once, report it as the `my-glossary` skill's to build, and skip the check
for the whole diff. Reporting every term as missing an entry buries the
findings that matter.

One published model is one text: the rendered schema shows its defining text
first, then its field descriptions in declaration order. Reading order decides,
so an expansion covers the descriptions after it and not the ones before it,
and the first text to use a short form is the finding even where a later field
expands it. An expansion in a type a field references covers the field pointing
at it. A tool description, an error body, a docstring, an exception or warning
message, a log message, and one interface string each stand alone.

### A value set copied into prose

Text listing the members of a closed set - the warning codes a tool description
enumerates, the values a field description offers, the states an interface
string names - is a copy of an enum. Read the enum and compare both ways: a
member the prose leaves out, and a member the prose lists that the enum no
longer has.

### What a caller or a model reads

* **A field description:** what the value is, and its unit where a number has
  one. Text describing how the number is computed is a finding; propose the
  definition in its place. Two more findings come from where the text sits: a
  field the schema says nothing about, on the field or on the type it
  references, and a field description that repeats or contradicts the defining
  text of the type it references. Propose the text on whichever of the two
  defines the concept.
* **A tool description:** what the tool returns, when to call it, and which
  tool comes first where an order exists. A model choosing between tools has
  this text and nothing else.
* **An error body:** what the caller can do about it, in their own terms. An
  upstream system's name, a mechanism, or anything the caller cannot act on is
  the finding; propose the plain version and check the handler logs the detail
  it drops.
* **A validation message, where the component authored it:** a validation error
  carrying its own text into the response body. Everything a validation library
  words itself and passes through untouched is not ours to hold to these rules,
  so skip it. An authored message names the field it is about and what would
  satisfy it. This is the one exception surface a caller reads, so the
  offending value may go in the message - but only where disclosing it is safe,
  meaning it came from that caller's own request and holds nothing personal or
  secret. Otherwise name the field and the rule it broke, and keep the value out
  of the message. The log is no fallback: a value too sensitive to return is too
  sensitive to log whole, and lands there only reduced or masked, such as an
  email address logged as its domain alone.

### What a person using the product reads

* **Interface copy:** what this control does or what this state means, in the
  reader's terms. A string naming an internal mechanism, a status code, or a
  component of the system is the finding.
* **Where the copy lives:** a user-facing string written inline at its use site,
  where the component declares a module that holds its copy, is the finding.
  Propose moving it there under a key that names it.
* **Data is not copy.** A value the backend returns, rendered as it arrived, is
  out of scope here and belongs to whichever surface published it.

### What a developer using a library reads

* **A public docstring:** what the thing does, its parameters, what it returns
  and what it raises, in the terms of the code calling it. How it works inside
  is the finding, and so is a parameter, return, or exception the docstring
  leaves out or names wrongly.
* **An exception or warning message the library raises:** what went wrong and
  what the caller can change, naming the argument or setting involved. The
  offending value may go in the message, since it came from the caller. A
  deprecation warning names what replaces the thing and, where the project
  says, when it goes.
* **Option and command help:** what the option does and its default, in one
  line that fits the help listing.

### What an operator reads

* **A log message:** the event, named so that an operator reading the line
  without its variable data knows what happened.
* **Where the variable data sits** follows the project's logging style. Where
  it logs structured fields, interpolation into the message is the finding;
  propose moving the value into a field under a key that names it, and the
  message stripped back to the event. Where it uses logging arguments, such as
  `%s` placeholders in stdlib `logging`, those are the idiom and eager
  formatting (an f-string) is the finding. Either way, a value the call site
  itself fixes, filling a template its sibling calls share, is part of which
  sentence this is and stays.
* **An exception message that never reaches a response:** the failure, and what
  would fix it - the setting to configure, the input to correct. Where the
  repository's rules declare that exceptions build their own messages, such as
  from a factory method taking data, a raise site passing prose rather than
  data is the finding, and so is a message naming no value the operator could
  act on where the factory was given one. Without that rule, a message written
  at the raise site is normal.

## 5. Report

Open with what the run covered, and one line per component in it saying which
rule sources and which inventory were read:

* **A full sweep:** that every declared surface was swept. There is no base ref
  to name, and the components are the ones that declared an inventory.
* **A diff run:** the base ref, so the reader can correct it, and the components
  the diff touches.

Then one table per component, ordered by file:

| `file:line` | Surface | Proposed | Reason |
| --- | --- | --- | --- |

`Proposed` holds the replacement text. Where that is too long for the cell to
read cleanly, write `see below` and put the full text in a fenced block under
the table, keyed by `file:line`. `Reason` is one clause naming the check that
fired and what it found, not a restatement of the rule. Text you weighed and
passed gets no row.

Under each table, one line per check that found nothing, so a clean check reads
differently from a skipped one.

Then, separately from the findings, **the unchecked components**: one entry
each, naming the section it needs. How the entry sizes the gap depends on the
run:

* **A diff run:** the count of strings the diff changed under that component.
  The hunks are there whether or not anything declared them.
* **A full sweep:** no string count. Counting them means knowing which
  constructs hold them, which is exactly what the component has not declared.
  Say it is undeclared and leave the number out rather than reporting a zero
  that reads as nothing to review.

Close with two counts kept apart, the findings and the unchecked components,
and an offer to apply the rewrites and to draft the missing sections.

**Rewrite nothing until the user accepts.** When they do, change the text and
everything that pins it: a test asserting a log message or an exception
message, a snapshot of the generated schema, and a test selecting on the copy
it renders.
