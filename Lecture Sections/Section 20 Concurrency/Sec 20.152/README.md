# Concurrency Documentation

## Effective Go
[Concurrency](https://go.dev/doc/effective_go#concurrency)
    - Share by communicating

[Goroutines](https://go.dev/doc/effective_go#goroutines)
    - They're called goroutines because the existing terms—threads,
    coroutines, processes, and so on—convey inaccurate connotations.
    A goroutine has a simple model: it is a function executing
    concurrently with other goroutines in the same address space.

    - Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's & notation for running a command in the background.)