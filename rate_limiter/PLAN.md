# Gap-closing plan: concurrency, networking, OOP/SOLID, design patterns

Trigger: interview feedback said "code more by yourself." These four areas
share one build instead of four separate study tracks — each step below adds
one gap onto the same rate limiter.

## Principle

Hands-on first. Reading only for facts you can't rediscover by coding
(status code meanings, protocol names, the SOLID acronym). Everything else
gets learned by hitting the real failure mode yourself.

## Steps (spartan/main.go)

- [X] **1. Correctness** — fix the sliding-window logic: per-user queue,
      `windowStart = t - config.Window`, evict expired entries then check
      `len(queue) < config.MaxRequest`. (Already mostly understood — just
      finish it.)
- [ ] **2. Concurrency** — write a test that fires requests from many
      goroutines at once, run `go test -race` *before* adding a lock and
      watch it catch a real race. Then add a `sync.Mutex` around the
      per-user queue access and confirm the race disappears.
- [ ] **3. Networking** — wrap the limiter in a tiny HTTP server
      (`net/http`, one endpoint: `/request?userID=1` → 200 or 429). Hit it
      with `curl`, inspect headers/status codes for real.
- [ ] **4. OOP/SOLID + design pattern** — extract
      `type RateLimiter interface { Allow(userID int) bool }` so the
      sliding-window-log implementation is swappable (Strategy pattern,
      Open/Closed + Dependency Inversion) without touching the HTTP
      handler. Prove it by writing a second dummy implementation
      (e.g. always-allow) and swapping it in.

## Reference reading (15 min each, just-in-time, not upfront)

- HTTP status codes worth knowing cold: 200, 400, 401, 403, 404, 429, 500.
- SOLID: five letters, one sentence each — read once, then point back at
  step 4 as the concrete example.
- Design patterns: skim only Strategy/Observer/Factory — recognize them in
  real code (e.g. a popular Go repo) rather than memorize the GoF catalog.

## Done when

- `go test -race ./spartan/...` passes with concurrent load.
- `curl localhost:PORT/request?userID=1` returns correct status codes.
- Swapping the `RateLimiter` implementation requires zero changes to the
  HTTP handler.
