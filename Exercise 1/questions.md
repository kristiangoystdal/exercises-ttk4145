## Exercise 1 - Theory questions

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is when two tasks are simultaneously executed, and what task is currently being executed and what core they are using are randomly chosen. Task 1 might use core 1 at first, but switch with task 2 at some point. Parallelism is different because the tasks stick to one core each and doesnt switch what core is used.

What is the difference between a *race condition* and a *data race*? 
> Race condition is when two threads are simultaneously executing tasks and the timing is different. Thread 1 might be read a value i and then changing it, but before that thread can write the new value, the second thread reads the old value and it therefore creates confusion and the final value will not be the expected value.
Data race is when two threads are simultaneously accessing a value, but one thread is writing a new value and the other thread is reading the value. Therefore it will create a problem if one thread reads before a write or the other way around, as the expected value might be different.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler makes sure that functions are executed in order, and the functions form a queue where they "wait" until the threads finish one task before its executes another. 


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
