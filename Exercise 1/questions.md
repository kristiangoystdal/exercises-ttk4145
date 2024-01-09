Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Your answer here*

What is the difference between a *race condition* and a *data race*? 
> *Your answer here* 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *Your answer here* 


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


