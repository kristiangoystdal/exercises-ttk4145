#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/tcp.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define BUFLEN 1024

int main() {
    int sock;
    struct sockaddr_in server_addr;
    char server_ip[] = "10.100.23.129";  // Server's IP address
    int port = 34933;  // Change to 34933 for fixed-size messages
    char buffer[BUFLEN];

    printf("Creating socket...\n");
    // Create socket
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("Socket creation error");
        return 1;
    }

    // Uncomment to set TCP_NODELAY if needed
    // int flag = 1;
    // setsockopt(sock, IPPROTO_TCP, TCP_NODELAY, (char *)&flag, sizeof(int));

    printf("Setting up server address...\n");
    // Setup server address structure
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(port);
    server_addr.sin_addr.s_addr = inet_addr(server_ip);

    printf("Connecting to server at %s:%d...\n", server_ip, port);
    // Connect to the server
    if (connect(sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        perror("Connection failed");
        return 1;
    }

    printf("Connected. Waiting for welcome message...\n");
    // Receive welcome message
    if (recv(sock, buffer, BUFLEN, 0) < 0) {
        perror("Receive failed");
    } else {
        printf("Server: %s\n", buffer);
    }

    // Send and receive echo messages
    while (1) {
        printf("Enter message: ");
        fgets(buffer, BUFLEN, stdin);
        buffer[strcspn(buffer, "\n")] = '\0'; // Replace newline with null terminator

        printf("Sending message: %s\n", buffer);
        if (send(sock, buffer, strlen(buffer) + 1, 0) < 0) {
            perror("Send failed");
            break;
        }

        printf("Waiting for echoed message...\n");
        if (recv(sock, buffer, BUFLEN, 0) < 0) {
            perror("Receive failed");
            break;
        }
        printf("Echoed: %s\n", buffer);
    }

    printf("Closing socket and exiting...\n");
    close(sock);
    return 0;
}
