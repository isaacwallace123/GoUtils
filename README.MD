# GoUtils 🛠️

A modular, community-driven utility library for Go developers.

**GoUtils** is a growing collection of high-quality utility packages for Go, designed to eliminate repetitive code and streamline common tasks — from string manipulation and logging to environment management and timing.

---

## ✨ What's Inside?

| Package        | Description |
|----------------|-------------|
| [`logger`](./logger)         | Thread-safe, color-coded, leveled logger with pluggable tags |
| [`env`](./env)               | Environment variable helpers with fallback, casting, and trimming |
| [`timeutil`](./timeutil)     | Time & duration helpers inspired by Roblox and Go best practices |
| [`stringutil`](./stringutil) | Powerful string transformations: casing, padding, parsing, and more |
| [`jsonutil`](./jsonutil)     | Safe JSON encoding/decoding, pretty/compact formatting, object tools |

---

## 🧠 Philosophy

This repository is built for:

- ✅ Rapid development — reuse robust tools you can trust
- ✅ Team collaboration — everyone can contribute new helpers
- ✅ Clean abstraction — each package serves one purpose

No external dependencies. Just practical, well-tested Go code.

---

## 🧪 Example Usage

```go
import (
  "github.com/your-org/GoUtils/logger"
  "github.com/your-org/GoUtils/env"
  "github.com/your-org/GoUtils/timeutil"
)

func main() {
  logger.Info("App started on port %d", env.GetInt("PORT", 8080))

  timer := timeutil.NewCountdownTimer(5)
  for !timer.Expired() {
    logger.Debug("Time left: %.1fs", timer.Remaining())
    timeutil.SleepSeconds(1)
  }

  logger.Fatal("Shutting down.")
}
```

---

## 🔁 Contributing
Anyone can contribute useful utilities! Keep it:
- Simple (1-purpose functions)
- Documented (use comments and good naming)
- Consistent (follow existing folder/module patterns)

Open a PR with your addition under the appropriate folder or propose a new one!

---

## 🛡 License
[MIT License](./LICENSE) — free to use, fork, and expand.

## 🙌 Special Thanks

To every contributor making GoUtils stronger and more versatile for the community.

We build faster when we build together. 🚀