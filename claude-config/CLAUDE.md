# CLAUDE.md

Personal defaults for every repository. A repository's own `CLAUDE.md`,
`AGENTS.md` or `CONTRIBUTING.md` wins where it disagrees with this file.

## Working with me

* **Check the branch before ticket work.** When the current branch is not the
  one named after the ticket, say so before starting.
* **Question the deliverable before shaping it.** Ask whether it is needed at
  all before deciding where it goes.
* **Implement only the piece asked for.** Leave a sibling task's design to that
  task.
* **Preview a ticket or a comment before posting it.** Show the text and wait.
* **A ticket title stands alone.** No vocabulary coined in the conversation;
  name the outcome, not the mechanism.

## Git

* **I own history.** Never commit, amend, push, tag, reset, rebase, or run
  anything else that writes history or moves refs unless asked in so many
  words in the current turn.
* **Finished work is not a request to commit.** Leave the changes in the
  working tree, say what changed, and stop.
* **Staging is part of committing.** Don't `git add` ahead of a commit. Where a
  tool needs staged files, such as pre-commit hooks, unstage afterwards.
* **Offer a commit or a push in one line** when it looks like the next step,
  then wait.
* **Read-only git is always fine**: `status`, `diff`, `log`, `show`, `blame`.

## Code

* **Match the surrounding code**: its naming, idiom, and comment density.
* **Line width comes from the linter configs.** Where two disagree, say so and
  use the strictest. With none configured, use 88, Markdown included.
* **Plain verbs over jargon** (`generate`, `create`, not `mint`, `hydrate`).
* **Reject input that would be ignored or contradicts other input.** Fail
  instead of computing and warning.
* **Test convenience alone never justifies a production change.** Build the
  fixture differently, unless the change also serves the code's users.
* **Use reserved `.example` domains** in sample data, not `example.com` or
  `.test`.

### Naming

* **Plain English.** Pick the more common word when two fit. Test: would
  someone unfamiliar with the ticket read the name cleanly at a glance?
* **One English variant per repository**, the one it already uses, in
  identifiers, comments, and in-repo docs. A new repository uses UK English.
* **Spell terms out.** A short form is fine only where the project already
  uses it.
* **No type or shape in a name** (`_dict`, `_list`, `_str`). A unit is not a
  type (`_days`, `_count`).
* **Name a class for its role, not its pattern.** `Manager`, `Factory`,
  `Handler`, `Helper` and the like name a construction, not a thing.
* **A parameter carries the name of what the call site passes it.**
* **A comment is not a name.** Where a comment says what a value or block
  holds, rename or extract instead. Drop a comment that restates the code.

### Published text

Text that leaves the process: API descriptions, tool descriptions, error
bodies, exception and log messages.

* **Full sentences.** Capitalized, terminal period. A title or label is not
  prose.
* **Expand an abbreviation on first use** in each text.
* **A description says what the value is**, with its unit. Not how it is
  computed.
* **An error body says what the caller can do about it.** No upstream system,
  no internal mechanism, nothing personal or secret.
* **A log message names what happened.** Variable data goes in structured
  fields, not the message.
* **An exception message names the failure and what would fix it.**
* **A value set written into prose is a copy of an enum.** Change both
  together.

## Tests

* **Given / When / Then comments are bare labels.** Combine adjacent ones
  (`# When / Then`) when a step has nothing to say.
* **The When block holds only the action under test.** Preparation goes in
  Given, reshaping in Then.
* **Name tests `test_<subject>__<outcome>`.** The happy path drops the outcome
  suffix.
* **Parametrize instead of repeating**, grouped by outcome shape (none, raises,
  value).
* **Mock doubles carry a spec** and a `_mock` suffix, never a `fake_` prefix.
* **Test the integration that consumes settings**, not the settings mechanism.

## Docs

* **No em dashes** in Markdown or plain text. Use `-` or rephrase.
* **No short forms in prose** (`dependencies`, not `deps`). Code identifiers
  stay as they are.
* **One source of truth.** Point to the canonical file instead of copying it.
* **`CLAUDE.md` is for the agent, `README.md` for the human.** Place a rule by
  who acts on it.
* **`CLAUDE.md` bullets are bare directives.** No citations, examples, or
  what-if clauses.
* **A rule that applies once lives with the artifact it constrains**, not in
  `CLAUDE.md`.
* **State a policy directly.** No "this is X, not a workaround" framing.
* **Qualify or drop a partial measurement.** Never cite it as authoritative.
