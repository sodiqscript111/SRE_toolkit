# Learning Notes

This folder is a personal, lightweight knowledge repository for documenting real observations, counterintuitive behaviors, and takeaways encountered while running experiments or debugging production systems.

---

## File Naming Convention

```text
notes/YYYY-MM-DD-<topic-slug>.md
```

Examples:
- `notes/2026-10-01-retry-amplification-incident.md`
- `notes/2026-10-15-cfs-throttling-gotcha.md`
- `notes/2026-11-02-coredns-ndots-amplification.md`

---

## Template

Use [`notes/template.md`](file:///notes/template.md) when adding a new note.

---

## Index of Notes

| Date | Topic | Summary |
|---|---|---|
| 2026-10-01 | [Retry Amplification & Cascade](file:///notes/2026-10-01-retry-amplification-incident.md) | How naive 3x retries across a 4-tier chain converted a 5% error rate into a 27x load spike and cluster collapse. |
