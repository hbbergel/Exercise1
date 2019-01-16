# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency is the ability to decompose a program into parts that can run independently and out-of-order and the result will be the same. The increase the performance. Parallelism is the ability to make a program faster by performing several computations at the same time in a processor. Require hardware with multiple processing units. The big difference is that concurrency can run on one single core, while parallellism requires multiple cores. 
 
 ### Why have machines become increasingly multicore in the past decade?
 > The single processor can run multiple instructions on seperate cores at the same time increasing the overall speed for programs. 
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Concurrent execution is needed in order to avoid inconsistencies in the database. For example websites must handle multiple simultaneous users.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > *Your answer here*
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > Process: OS-managed truly concurrent. Exists withing their own address space. 
 Thread: OS-managed, within the same address space as the parent and all its other threads. Possible truly concurrent, and multi-paskng is pre-emptive. 
 Green thread: User-space projections of the same concepts as threads, but not OS-managed. Probablt not concurrent, but there may be multiple worker threads or processors giving them CPU time concurrency (multiplexed).
 Coroutines: Not OS-managed, looks like threads except co-operatively multitasking --> hence not truly concurrent. 
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > pthread_create() - creates a new thread.
 threading.Thread() - green thread
 go                 - coroutines
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > Prevents multiple threads from executing Python bytecodes at once. This lock is necessary mainly because Python's memory management is not thread-safe. 
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > Multiprocessing module. 
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting. If n < 1, it does not change the current setting.
