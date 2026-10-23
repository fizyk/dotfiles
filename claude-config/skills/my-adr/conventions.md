# Writing a decision record

How this repository writes and reviews the records in this directory.

The format is Michael Nygard's: a title and four sections, Status, Context, Decision,
and Consequences, in one or two pages. [His article][nygard] is the source for those
parts, and the rest of this page is local to this repository: when to write a record,
what stays out of one, and the review checks.

## Title

* A short noun phrase naming what was decided, not the act of deciding.
* Records are numbered in sequence. A number is never reused.

## Status

* A new record is Proposed. Review by the people who own the decision accepts it,
  and only then is it Accepted. A decision reached in a meeting or a conversation
  does not move it.
* A record that is reversed or replaced stays in place, marked Deprecated or
  Superseded, with a link to the record that replaces it and a link back.
* A record whose decision has shipped and held is Accepted, whatever it still says.

## When to write one

* Write one for an architecturally significant decision: it changes structure,
  non-functional characteristics, dependencies, interfaces, or construction
  techniques.
* Where there is an existing component to ship into, ship the change first and let
  review negotiate adoption. The implementation is the proposal, and a record follows
  only once the decision holds.
* A new component has nothing to ship into, so its record comes first. Keep it to what
  the first implementation settles.
* A choice that does not fit in one paragraph is not one decision yet. Split it, or
  wait until it can be decided.

## Context

* The forces at play: what is true today, what constrains the choice, and what goes
  wrong if nothing changes.
* Value-neutral. Describe, do not adjudicate.
* Each candidate solution gets a subsection saying what it is, with its pros and cons.
  No verdict inside an option: no "Rejected:", no "Listed for completeness". The
  Decision chooses.
* A referenced document gets a link and a sentence only for what it constrains in the
  Decision.

## Decision

* The choice in one paragraph, full sentences, active voice: "We will ...".
* What the choice settles, not how the chosen thing works. Formats, flows, and failure
  tables belong to the design document the record links to.
* No bounds. "This holds while X" and "revisit when X" are consequences.

## Consequences

* The context after the decision applies: all of it, not only the favorable part.
  Group it under Positive and Negative where that helps.
* What the decision leaves unsettled goes under Negative.
* The conditions that would reopen the decision go under Negative.

## What does not belong

* File paths and module names. Name the role the code plays because paths move and
  the record does not follow. The same goes for endpoint paths and ticket numbers.
* A record referenced by a number it does not have yet. Name the concern in prose.
  References back to existing records are welcome.
* Framing a choice as "this is X, not a workaround". State it directly.
* A partial measurement presented as authoritative. Qualify its scope or drop it.
* A summary of another document in place of a link to it.

## Review checks

In the order they tend to find something.

* **A Decision much longer than its Context.** The record has absorbed the design
  document it should link to.
* **A fifth heading.** Open items mean the record was written before the decision
  could be made. Settle them, list them as unsettled under Negative, or wait.
* **The same statement in two sections.** Keep it where a reader needs it.
* **A Context restating the Decision as a requirement.** The decision then cannot fail
  to meet it, and the record establishes nothing.

[nygard]: https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions
