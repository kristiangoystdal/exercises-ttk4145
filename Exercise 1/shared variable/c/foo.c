// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;

// Note the return type: void*
void *incrementingThreadFunction()
{
    for (int j = 0; j < 1000001; j++)
    {
        i++;
    }
    return NULL;
}

void *decrementingThreadFunction()
{
    for (int j = 0; j < 1000001; j++)
    {
        i--;
    }
    return NULL;
}

int main()
{

    pthread_t thread;

    pthread_create(&thread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread, NULL, decrementingThreadFunction, NULL);
    pthread_join(thread, NULL);

        // TODO:
        // start the two functions as their own threads using `pthread_create`
        // Hint: search the web! Maybe try "pthread_create example"?

        // TODO:
        // wait for the two threads to be done before printing the final result
        // Hint: Use `pthread_join`

        printf("The magic number is: %d\n", i);
    return 0;
}
