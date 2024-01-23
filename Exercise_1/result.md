Exercise 1: Results
==================================================


Task 3:
------------------------------------------------
When running the code for C and Go, the code create to threads that run simultaneously will be executed in parallel. Therefore the threads will be in a data race where they both will peform their own tasks. This means that the variable i will both increase and decrease, and the variable should be different each time the code is executed. 

Task 4:
------------------------------------------------
When the task runs properly, the variable i will be updated in both functions and the vaule of i will become the expected value. What happens is that the variable gets locked and unlocked so the the variable change is synchronized between the functions. 

For the GO implementation, we use channels and those works by creating a virtual queue of the functions where the current operation cant be changed before it's read by the select function. While the operation channels waits for the reading, the for loops are stuck when trying to push a new vaule to the channel. Therefore, the for loops wait for the operation to be read and then change the value. 