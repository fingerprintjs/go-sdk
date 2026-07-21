# sync.sh Mock Fixtures Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make `bin/sync.sh` download the example/mock JSON fixtures under `test/mocks/` (not just the OpenAPI schema), prune the mock files no test references, and document the new behavior in `contributing.md`.

**Architecture:** `bin/sync.sh` gains a curated `examples` array (the 12 mock file paths go-sdk's tests actually use) and a loop that curls each one from the OpenAPI repo's `examples/` directory into `test/mocks/`, mirroring the existing pattern in `python-sdk`/`dotnet-sdk`/`php-sdk`'s own sync scripts. The 18 mock files not in that list (leftovers from the separate, external `update-sdk-schema` automation, which is out of scope here) are deleted in the same change so the script and the directory agree.

**Tech Stack:** Bash (`bin/sync.sh`), Go test suite (`test/`), Markdown (`contributing.md`).

## Global Constraints

- Single commit, single PR — see design spec's "worth breaking up into multiple commits?" conclusion: the script change and the mock pruning are two halves of one decision and should land together.
- Preserve the current schema download URL and destination filename (`fingerprint-server-api-v4-normalized.yaml` → `res/fingerprint-server-api-v4.yaml`) unchanged — not in scope for this ticket.
- No `--retry` on `curl` calls — this script only ever runs manually/locally.
- Do not chain into `bin/openapi-codegen-docker.sh` (or any codegen step) at the end of `sync.sh` — go-sdk keeps "refresh schema/mocks" and "run codegen" as separate manual steps (see `contributing.md`), unlike `python-sdk`.
- Do not run the real `bin/sync.sh` as part of committing content changes — verify it works, then discard any resulting mock/schema content diff. The actual mock refresh is an explicit follow-up PR per the ticket's DoD, done after this one merges.
- Full design context and rationale: `docs/superpowers/specs/2026-07-20-sync-sh-mock-fixtures-design.md`.

---

### Task 1: Update `bin/sync.sh`, prune unused mocks, update `contributing.md`

**Files:**
- Modify: `bin/sync.sh`
- Modify: `contributing.md`
- Delete: 18 files under `test/mocks/` (listed in Step 2)
- Verify via: `cd test && go test ./...`

**Interfaces:**
- Consumes: nothing from prior tasks (this is the only task in this plan).
- Produces: the final `bin/sync.sh` — any future work (e.g. the follow-up mock-refresh PR, or a possible future `update-sdk-schema` integration) invokes this script with zero or two positional args (`schemaUrl`, `examplesBaseUrl`) and expects it to populate `res/fingerprint-server-api-v4.yaml` and the 12 files under `test/mocks/` listed in Step 4.

- [ ] **Step 1: Confirm the test suite is green before touching anything**

Run:
```bash
cd test && go test ./... && cd ..
```
Expected: `ok` for all packages, no failures. This is the baseline — Step 3 re-runs this to prove deleting the 18 unused mocks didn't break anything.

- [ ] **Step 2: Delete the 18 mock files no test references**

```bash
git rm test/mocks/webhook/webhook_event.json \
  test/mocks/errors/400_start_time_invalid.json \
  test/mocks/errors/400_reverse_invalid.json \
  test/mocks/errors/400_pagination_key_invalid.json \
  test/mocks/errors/400_end_time_invalid.json \
  test/mocks/errors/403_secret_api_key_not_found.json \
  test/mocks/errors/400_limit_invalid.json \
  test/mocks/errors/400_visitor_id_required.json \
  test/mocks/errors/400_bot_type_invalid.json \
  test/mocks/errors/400_linked_id_invalid.json \
  test/mocks/errors/400_event_id_invalid.json \
  test/mocks/errors/500_internal_server_error.json \
  test/mocks/errors/400_ruleset_not_found.json \
  test/mocks/errors/403_wrong_region.json \
  test/mocks/events/update_event_one_field_request.json \
  test/mocks/events/update_event_multiple_fields_request.json \
  test/mocks/events/get_event_with_bot_info_200.json \
  test/mocks/events/get_event_ruleset_200.json
```
Expected: `rm 'test/mocks/...'` printed 18 times, no errors. `test/mocks/webhook/` will disappear entirely (its only file is gone, and git doesn't track empty directories) — this is expected, not a bug; confirmed during design that nothing references the bare directory path.

- [ ] **Step 3: Re-run the test suite to confirm the deletions are safe**

Run:
```bash
cd test && go test ./... && cd ..
```
Expected: same result as Step 1 — `ok` for all packages. If anything fails referencing a deleted mock, stop and re-check the design doc's "verified" list before proceeding (it means a file was pruned incorrectly).

- [ ] **Step 4: Replace `bin/sync.sh`**

Replace the full contents of `bin/sync.sh` with:

```bash
#!/bin/bash
set -euo pipefail

defaultBaseUrl="https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi"
schemaUrl="${1:-$defaultBaseUrl/schemas/fingerprint-server-api-v4-normalized.yaml}"
examplesBaseUrl="${2:-$defaultBaseUrl/examples}"

mkdir -p ./res
curl -fSL -o ./res/fingerprint-server-api-v4.yaml "$schemaUrl"

examples=(
  'events/get_event_200.json'
  'events/search/get_event_search_200.json'
  'errors/400_ip_address_invalid.json'
  'errors/400_request_body_invalid.json'
  'errors/400_visitor_id_invalid.json'
  'errors/403_feature_not_enabled.json'
  'errors/403_secret_api_key_required.json'
  'errors/403_subscription_not_active.json'
  'errors/404_event_not_found.json'
  'errors/404_visitor_not_found.json'
  'errors/409_state_not_ready.json'
  'errors/429_too_many_requests.json'
)

baseDestination="./test/mocks"

for example in "${examples[@]}"; do
  destinationPath="$baseDestination/$example"
  destinationDir="$(dirname "$destinationPath")"
  mkdir -p "$destinationDir"

  exampleUrl="$examplesBaseUrl/$example"
  echo "Downloading $exampleUrl to $destinationPath"
  curl -fSL -o "$destinationPath" "$exampleUrl"
done
```

- [ ] **Step 5: Syntax-check the script and confirm it's still executable**

Run:
```bash
bash -n bin/sync.sh
ls -l bin/sync.sh
chmod +x bin/sync.sh
```
Expected: `bash -n` prints nothing and exits 0 (valid syntax). `ls -l` should already show `-rwxr-xr-x` (it was executable before this edit); the `chmod +x` is a no-op safety net in case the edit tool changed the permission bit.

- [ ] **Step 6: Functionally verify the script, then discard the content diff it produces**

This actually runs the script against the real network to prove it works end-to-end, but the resulting content refresh is explicitly out of scope for this PR (that's the separate follow-up PR per the ticket's DoD) — so the diff gets discarded after verifying.

Run:
```bash
bash bin/sync.sh
echo "exit code: $?"
```
Expected: exit code `0`, with 13 "Downloading ..." lines printed (1 implicit for the schema — no log line for that one — plus 12 explicit example log lines), no `curl: (...)` errors.

Then verify the results:
```bash
test -s res/fingerprint-server-api-v4.yaml && echo "schema OK"
find test/mocks -type f | wc -l
```
Expected: `schema OK` printed; the file count is exactly `12`.

Then discard the working-tree content changes (the 18 deletions from Step 2 stay staged — this only reverts the 12 remaining files' and the schema's *content*, since a real content refresh belongs to the separate follow-up PR):
```bash
git status --short test/mocks res
git checkout -- test/mocks res
git status --short test/mocks res
```
Expected: the first `git status` may show modified content for some/all of the 12 files and `res/fingerprint-server-api-v4.yaml` (or nothing, if upstream hasn't changed since these were last committed). After `git checkout`, the second `git status` should show nothing for `res` and, for `test/mocks`, only the 18 deletions already staged from Step 2 (still listed as staged deletions — `git checkout -- <path>` only restores files that exist in the index, so it cannot and does not resurrect the already-removed files).

- [ ] **Step 7: Update `contributing.md`**

Change the `test` bullet in the "Structure" section from:
```markdown
- [test](./test) - directory for unit tests. Run `cd test && go test ./...` to run them.
```
to:
```markdown
- [test](./test) - directory for unit tests. Run `cd test && go test ./...` to run them. The example fixtures under `test/mocks` are kept in sync via `/bin/sync.sh`, same as the schema.
```

Change the sync instruction line from:
```markdown
Run `sh ./bin/sync.sh` to download the latest Fingerprint OpenAPI schema.
```
to:
```markdown
Run `sh ./bin/sync.sh` to download the latest Fingerprint OpenAPI schema and refresh the example fixtures under `test/mocks`. Only the fixtures actually used by tests are tracked — it's not a mirror of every example the OpenAPI repo publishes.
```

- [ ] **Step 8: Final review before committing**

Run:
```bash
git status --short
git diff --cached --stat
git diff --stat
```
Expected: `git status --short` shows exactly: 18 lines of `D  test/mocks/...` (staged deletions from Step 2), ` M bin/sync.sh` (unstaged — not yet added), ` M contributing.md` (unstaged). `git diff --stat` (unstaged) shows only `bin/sync.sh` and `contributing.md`. Nothing under `test/mocks` or `res` should appear as modified content (Step 6 discarded that).

- [ ] **Step 9: Commit**

```bash
git add bin/sync.sh contributing.md
git commit -m "chore: sync test/mocks example fixtures via bin/sync.sh"
```
Expected: one commit containing the 18 deletions (staged since Step 2) plus the `bin/sync.sh` and `contributing.md` changes.

---

## Follow-up (separate PR, not part of this plan)

Per the ticket's DoD: after this PR merges, run `bin/sync.sh` again on a fresh branch and open a second PR with whatever content diffs land for the 12 tracked mock files (this is the same action as Step 6 above, minus the discard-the-diff part).
