## Exercise 1 - Theory questions

### Concepts

What is the difference between _concurrency_ and _parallelism_?

> _Your answer here_

What is the difference between a _race condition_ and a _data race_?

> _Your answer here_

_Very_ roughly - what does a _scheduler_ do, and how does it do it?

> _Your answer here_

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?

> Multiple threads is useful for improving speed for tasks run on multi-core processors.
> It is for example possible to separate front-end and back-end code to run on separate threads. This prevents the UI from becoming unresponsive if heavy operations are run on the back-end.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?

> Fibers are lightweight threads of execution fibers that uses cooperative multitasking instead of preemptive multitasking likeregular threads. Meaning the processes that are run yield control periodically or when idle, instead of the operating system.
> A coroutine is an instance of a suspendable computation, which may suspend its execution in one thread and resume in another one.
> Fibers and coroutines are more lightweight in terms of memory and is therefor useful for programmes that require many tasks to be run concurrently.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?

> Concurrent programs allow for faster execution, and if done correctly, will be very organized.

What do you think is best - _shared variables_ or _message passing_?

> There is no one-size-fits-all answer to this one. Shared variables are simpler to implement but may lead to concurrency issues like race conditions or deadlocks. Message passing is safer for avoiding issues and is more easily scalable, but may require more memory
