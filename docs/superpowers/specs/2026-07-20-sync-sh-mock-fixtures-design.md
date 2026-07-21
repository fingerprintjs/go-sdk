# sync.sh: download example/mock JSON fixtures

Ticket: [INTER-2330](https://fingerprintjs.atlassian.net/browse/INTER-2330)

## Problem

`bin/sync.sh` only refreshes `res/fingerprint-server-api-v4.yaml` (the OpenAPI
schema). It doesn't touch `test/mocks/`, so the example fixtures used by tests
can silently drift from what the OpenAPI repo actually publishes. Other Server
SDKs (Python, .NET, PHP, Node) already have this in their own `sync.sh`
scripts; go-sdk doesn't.

## Context from #integrations-internal

A Slack thread ([2026-07-17](https://fingerprintjs.slack.com/archives/C050TKU9X5L/p1784301400376359),
plus two threads it links back to) surfaced two things that shape this design:

1. **The curated-list approach is load-bearing, not cosmetic.** The
   `update-sdk-schema` GitHub Action (in `dx-team-toolkit`, triggered from the
   OpenAPI repo) currently distributes *every* example file from a release to
   each SDK — this is what caused Edge-specific mocks to leak into
   `python-sdk` during a past sync. The team's proposed fix is for that action
   to stop doing its own example distribution and instead just invoke each
   SDK's own `sync.sh`, which knows the specific subset of examples it
   actually uses. This ticket's mechanism is a building block for that
   direction, not just a go-sdk-specific nicety.
2. **go-sdk's `test/mocks/` is already wired into the action that causes this
   problem.** Checking `sync-server-side-sdks-schema.yml` +
   `update-schema-for-tag.ts` confirmed go-sdk's `test/mocks` is the
   `examples-path` target of the automated `sync-go` job, which wipes the
   directory and rewrites it with everything in scope
   (`events,visitors,webhook,related-visitors,events-search`) on every run.
   This is why 18 of the 30 files currently in `test/mocks/` aren't referenced
   by any test — they're leftovers from past automated syncs, not something a
   developer hand-picked.

The team has discussed this unification repeatedly (Apr, Jul 1, Jul 17) with
real engagement from multiple engineers, but as of the latest message it's
still undecided and unimplemented — no ticket, no code change yet. This design
doesn't depend on that unification landing; it only takes the low-cost step of
keeping go-sdk's `sync.sh` consistent with the shape siblings already use, so
it's ready if/when that direction is adopted.

## Design

### 1. `bin/sync.sh`

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

Notes:

- The schema URL and destination filename are unchanged from the current
  script — not in scope for this ticket.
- `set -euo pipefail` and `-fSL` on both curls are added so a broken URL or
  HTTP error fails the script loudly instead of silently writing an error
  page as if it were valid YAML/JSON. This matches `python-sdk`, `dotnet-sdk`,
  and `php-sdk`'s existing scripts.
- `schemaUrl` / `examplesBaseUrl` are added as optional positional args with
  defaults, matching the same three sibling SDKs. No current caller passes
  custom values (confirmed: `update-schema-for-tag.ts` bypasses every SDK's
  `sync.sh` entirely today, in all languages), so this doesn't change current
  behavior — it's there for consistency and in case the `update-sdk-schema`
  unification above adopts this shape.
- `examples` is a curated list of exactly the files go-sdk's tests reference
  today (verified — see below), not a mirror of everything currently in
  `test/mocks/`.
- No `--retry` on either `curl` call, even though `python-sdk`/`php-sdk` have
  it (`dotnet-sdk`/`node-sdk` don't). This script only ever runs manually/
  locally (confirmed: no `.github/workflows/*.yml` invokes `bin/sync.sh`), so
  a transient network blip is something a person re-runs by hand; automatic
  retries wouldn't add value for an attended run.
- Doesn't call `bin/openapi-codegen-docker.sh` (or equivalent) at the end,
  unlike `python-sdk`'s `sync.sh` which chains into `./generate.sh`. go-sdk's
  `contributing.md` already treats "refresh schema/mocks" and "run codegen"
  as two separate manual steps (Docker-based codegen is slow and its own
  step) — this design preserves that split rather than adopting Python's
  coupling, to avoid silently expanding this ticket's blast radius into
  codegen output.
- Partial-failure behavior is intentionally left as-is: with `set -e`, a
  failed download stops the loop immediately. Verified empirically that
  `curl -fSL` writes nothing to the destination on an HTTP error (tested
  against a 404) — so a mid-loop failure leaves already-processed files
  refreshed, the failing file and everything after it on their previous
  content (never corrupted/truncated), and a nonzero exit code. No sibling
  SDK's script does atomic temp-file+rename either. Acceptable because this
  never runs unattended.

### 2. Prune unused mock files

Delete the 18 files under `test/mocks/` that no test references (verified via
full-repo search, any file extension, including checks for dynamic path
construction — none found):

```
test/mocks/webhook/webhook_event.json
test/mocks/errors/400_start_time_invalid.json
test/mocks/errors/400_reverse_invalid.json
test/mocks/errors/400_pagination_key_invalid.json
test/mocks/errors/400_end_time_invalid.json
test/mocks/errors/403_secret_api_key_not_found.json
test/mocks/errors/400_limit_invalid.json
test/mocks/errors/400_visitor_id_required.json
test/mocks/errors/400_bot_type_invalid.json
test/mocks/errors/400_linked_id_invalid.json
test/mocks/errors/400_event_id_invalid.json
test/mocks/errors/500_internal_server_error.json
test/mocks/errors/400_ruleset_not_found.json
test/mocks/errors/403_wrong_region.json
test/mocks/events/update_event_one_field_request.json
test/mocks/events/update_event_multiple_fields_request.json
test/mocks/events/get_event_with_bot_info_200.json
test/mocks/events/get_event_ruleset_200.json
```

The remaining 12 files are exactly the `examples` list in `sync.sh` above, so
the script and the directory agree from the moment this PR merges. Pruning
happens in this PR (not a separate one) so there's no in-between state where
the script manages 12 files but the directory still has 30.

Caveat: the automated `sync-go` job will still repopulate `test/mocks` with
everything in scope on its next run, since it doesn't go through `sync.sh`.
This pruning is correct regardless, but it isn't a permanent fix by itself —
that requires the broader unification discussed above.

### 3. `contributing.md`

Update the line:

> Run `sh ./bin/sync.sh` to download the latest Fingerprint OpenAPI schema.

to also mention that it refreshes `test/mocks/` fixtures, and note that the
tracked mocks are a curated subset (only what tests exercise), not a mirror of
every example the OpenAPI repo publishes — so future contributors understand
why the list is short rather than exhaustive.

Also add a bullet to the "Structure" section's `test` entry mirroring the
existing `res` entry's pattern (*"directory for Fingerprint OpenAPI schema.
Use `/bin/sync.sh` to keep it up to date."*), so `test/mocks/` being
sync-managed is documented in the same place a contributor would already look
for the schema's equivalent note.

## Verified non-issues (no design change needed)

- **No `go:embed` directives** reference `test/mocks` — deleting the 18 files
  can't break a build-time embed.
- **`test/mocks/webhook/` will disappear entirely** once its only file
  (`webhook_event.json`) is deleted (git doesn't track empty directories).
  Confirmed nothing references the bare directory path — `webhook_test.go`
  uses a hardcoded example, not this fixture.
- **`.schema-version`** (currently `v3.4.2`) exists in go-sdk like it does in
  the other SDKs, but nothing in this repo's scripts touches it — it's
  updated externally by the `update-sdk-schema` action. Not part of this
  ticket.
- **The schema URL's normalized/non-normalized flip** has churned twice in
  `bin/sync.sh`'s history with no documented rationale in any commit message.
  Confirmed out of scope — this design preserves the current (normalized)
  choice unchanged.

## Out of scope / follow-up

Per the ticket's DoD, after this PR merges: run the new `sync.sh` locally and
open a separate PR with whatever fresh content diffs result for the 12 mock
files. Not done as part of this design/PR.
