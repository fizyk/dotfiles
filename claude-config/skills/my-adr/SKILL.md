---
name: my-adr
description: Write or review an architecture decision record against this repository's conventions - which section owns what, what does not belong in one, and the shape checks that catch a record carrying its own design document. Reads the conventions page beside the records, and where a repository has none, offers the seed page bundled with this skill. $ARGUMENTS names a record to review, `new <title>` to draft one, or is empty to be asked which. Use when asked to write, draft, review or trim an ADR or decision record.
allowed-tools: Read, Grep, Glob, Edit, Write, Bash(adr new *), Bash(adr link *), Bash(git diff *), Bash(git status *), Bash(git show *), Bash(git rev-parse *), Bash(git branch *), Bash(grep *)
---

# Decision records

The rules are in the convention page beside the records they bind, which
section 2 finds. Read them there rather than a second copy here. Only the
procedure that applies them lives in this file, which is what lets another
repository run the same skill against its own conventions.

The page distinguishes what it inherits from Michael Nygard's format from what
is local. Keep that distinction when reporting: a departure from a local rule
is this repository's to change, a departure from the format is not.

Writing and reviewing run the same rules. Reviewing asks them as questions.

## 1. Scope the run

**`$ARGUMENTS` is `new <title>`:** draft a record. Section 3.

**`$ARGUMENTS` names a record** (a number, a file name or a path): review that
one. Section 4.

**`$ARGUMENTS` is empty:** ask whether to draft a new record or review an
existing one, and which, then stop.

A record is reviewed whole, never by its hunks. A section is judged against the
other three, which a diff does not show.

## 2. Resolve the rules

Read, in this order:

1. Where the records live, and which tool manages them. The repository's own
   `CLAUDE.md`, README or task runner names it where it says anything. Then a
   tool's configuration file, such as `.adr-dir` for adr-tools or
   `.log4brains.yml` for log4brains. Without either, look for a directory named
   `adr` or `decisions`, under `docs/` first.
2. The conventions page in that directory, or the one the repository's own
   `CLAUDE.md` names.
3. The record under review, or, when drafting, the records the conventions
   page cites as examples, where it cites any.

**Use `Read`, not `cat`.** Reading a file under a component is also what loads
that component's own instructions for the session.

**Where there is no conventions page**, check the template the existing records
follow. Where it is not Nygard's, such as MADR, say so: the seed's sections would
not match them. Then say once that there is no page and offer the seed:
`conventions.md` in this skill's own directory. It takes the format from
Nygard's article and adds rules for when to write a record, what stays out of
one, and the review checks. Show where it would go, beside the records, and
wait.

* **Accepted:** copy it there unchanged, then run against the copy. The
  repository edits its own copy from then on.
* **Accepted for this run only**, as in a repository you do not own: apply the
  seed without writing it, and say in the report that the rules came from the
  seed, not the repository.
* **Declined:** report the records as unchecked. Do not fall back to the
  format's own article alone: the local rules are most of what a review finds,
  and a clean pass against half the rules reads as a pass.

## 3. Drafting

Before the file exists, settle two things, both as the conventions page puts
them. Whether the decision is architecturally significant. And, where there is
an existing component to ship into, whether it has already shipped and survived
review; a net-new component has nothing to ship into, so its record comes
first, bounded to what the first implementation settles. Where either test
applies and fails, the record is premature; say which and stop.

Then create the file with the tool section 2 found, the way the repository
runs it: `adr new <title>` for adr-tools, from the directory holding
`.adr-dir`. Without a tool, create it by hand: the next number, the file name
pattern the existing records use, and the template the directory holds, or else
the headings the conventions page names.

Fill them in order, and Context in two passes:

1. **Context**, forces first. Stop when a reader could anticipate the decision.
2. **Context**, the candidate solutions, one subsection each, value-neutral.
3. **Decision**, in one paragraph, active voice.
4. **Consequences**, from that paragraph: what the choice costs, what it leaves
   unsettled, what would reopen it.

Writing the Decision before the forces is how a Context ends up restating the
Decision as a requirement, which section 4 then has to catch.

Run section 4 against the draft before handing it over.

## 4. Reviewing

Run the checks in this order. The structural ones find more, and they change
what the line-level ones are looking at.

### Measure first

```shell
grep -c '' <record>
grep -n '^## ' <record>
```

The first gives the record's total length, the second the line each section
starts on. A section runs to the next heading, and the last to the total.

Two numbers decide where to look: Decision against Context, and the total
against the one-or-two-page bound. A Decision several times its Context is the
single most reliable sign that the record has absorbed the design document it
points at, so start there and ask which lines a reader could not act on
without.

A fifth heading is a finding on its own. Report what it holds, sorted into what
the decision leaves unsettled, which belongs under Negative, and what was not
decidable, which belongs in a later record or nowhere yet.

### Then section ownership

For each section, the question is only whether its content belongs to it. Take
the ownership rules from the conventions page; the ones that find the most are
mechanism in the Decision, a verdict inside an option body, a referenced
document summarized rather than linked, and a deferral written as a force.

### Then duplication

A statement appearing twice, in either order. Keep the one where a reader needs
it, which is rarely the one written first. Say which copy to cut and why that
one.

### Then the line-level rules

The title's form, the status, the active voice, and everything under what does
not belong. These are cheap and they are last because a section that is moving
wholesale does not need its wording reviewed first.

## 5. Report

Open with the record and its section line counts.

Then one table, ordered by the record's own sections:

| `record:line` | Check | Finding | Proposed |
| --- | --- | --- | --- |

The first column is the record's number and the line, `0008:44`. It stays short
where a run covers several records, and the header has already named the file.

`Finding` is one clause naming what fired, not a restatement of the rule.
`Proposed` is the replacement text, or the section to move the content to.
Where a proposal is too long for the cell, write `see below` and put it in a
fenced block under the table, keyed by `file:line`. Content you weighed and
passed gets no row.

Close with one line per check that found nothing, so a clean check reads
differently from a skipped one, and an offer to apply the changes.

**Change nothing until the user accepts.** When they do, move content rather
than rewriting it where the finding was that it sits in the wrong section, and
check that a link into the record still resolves after a heading moves.
