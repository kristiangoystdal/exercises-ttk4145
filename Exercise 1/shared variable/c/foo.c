// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock;

void *incrementingThreadFunction()
{
    for (int j = 0; j < 1000001; j++)
    {
        pthread_mutex_lock(&lock);
        i++;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void *decrementingThreadFunction()
{
    for (int j = 0; j < 1000001; j++)
    {
        pthread_mutex_lock(&lock);
        i--;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

int main()
{
    pthread_t thread;

    pthread_mutex_init(&lock, NULL);

    pthread_create(&thread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread, NULL, decrementingThreadFunction, NULL);
    pthread_join(thread, NULL);
    pthread_mutex_destroy(&lock);

    printf("The magic number is: %d\n", i);
    return 0;
}
