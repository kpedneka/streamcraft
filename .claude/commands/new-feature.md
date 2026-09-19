The human wants to start on a new feature or backlog item: $ARGUMENTS

Before any code gets written:
1. Restate the goal in one or two sentences to confirm shared understanding.
2. Check whether this touches auth, uploads, ffmpeg invocation, or streaming URLs — if so, name the specific risks up front.
3. Break it into a small number of reviewable steps/commits, not one giant PR.
4. Flag if this looks like it needs an ADR before implementation starts (a real tradeoff, not just an implementation detail).
5. Ask any clarifying questions before the human starts writing code.

Do not start writing the implementation yourself — this command is for planning and scoping, per the working agreement in CLAUDE.md.
