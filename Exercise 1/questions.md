Exercise 1 - Theory questions
-----------------------------

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
    It is for example possible to separate front-end and back-end code to run on separate threads. This prevents the UI from becoming unresponsive.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*

What do you think is best - *shared variables* or *message passing*?
> *Your answer here*


