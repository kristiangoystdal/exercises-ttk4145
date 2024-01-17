#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define BUFLEN 512

int send_sock;
int recv_sock;
struct sockaddr_in server_addr;
int server_port;
char *broadcast_ip = "255.255.255.255"; // Use "#.#.#.255" or "255.255.255.255" for broadcast

void *send_thread_func(void *arg) {
    char message[BUFLEN];
    while(1) {
        printf("Enter message: ");
        fgets(message, BUFLEN, stdin);

        // Remove newline
        message[strcspn(message, "\n")] = 0;

        if(sendto(send_sock, message, strlen(message), 0, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
            perror("sendto");
            break;
        }
    }
    return NULL;
}

void *recv_thread_func(void *arg) {
    char buf[BUFLEN];
    struct sockaddr_in si_other;
    unsigned int slen = sizeof(si_other);
    int recv_len;

    while(1) {
        if((recv_len = recvfrom(recv_sock, buf, BUFLEN, 0, (struct sockaddr *) &si_other, &slen)) == -1) {
            perror("recvfrom");
            break;
        }

        buf[recv_len] = '\0';
        printf("Received: %s\n", buf);
    }
    return NULL;
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        fprintf(stderr, "Usage: %s <workspace_number>\n", argv[0]);
        return 1;
    }

    char *server_ip = "10.100.23.129";  // Server's IP address

    int workspace_number = atoi(argv[1]);
    server_port = 20000 + workspace_number;

    // Create send and receive sockets
    send_sock = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
    recv_sock = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);

    // Enable broadcast
    int broadcastPermission = 1;
    setsockopt(send_sock, SOL_SOCKET, SO_BROADCAST, &broadcastPermission, sizeof(broadcastPermission));

    // Setup server address structure
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(server_port);
    server_addr.sin_addr.s_addr = inet_addr(server_ip);

    // Bind receive socket
    struct sockaddr_in recv_addr;
    memset(&recv_addr, 0, sizeof(recv_addr));
    recv_addr.sin_family = AF_INET;
    recv_addr.sin_port = htons(server_port);
    recv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    bind(recv_sock, (struct sockaddr *) &recv_addr, sizeof(recv_addr));

    // Create send and receive threads
    pthread_t send_thread, recv_thread;
    pthread_create(&send_thread, NULL, send_thread_func, NULL);
    pthread_create(&recv_thread, NULL, recv_thread_func, NULL);

    pthread_join(send_thread, NULL);
    pthread_join(recv_thread, NULL);

    close(send_sock);
    close(recv_sock);

    return 0;
}
