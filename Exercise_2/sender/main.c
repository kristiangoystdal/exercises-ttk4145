#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>

#pragma comment(lib, "Ws2_32.lib")

void error_handling(char *message) {
    fprintf(stderr, "%s: %d\n", message, WSAGetLastError());
    exit(1);
}

int main(int argc, char *argv[]) {
    WSADATA wsaData;
    SOCKET sock;
    struct sockaddr_in addr;
    int broadcast = 1;

    char *remoteIP = "10.22.73.189"; // Set this to your target IP
    int remotePort = 20000; // Set this to your target port
    char *message = "Testing";

    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0)
        error_handling("WSAStartup() error");

    // Creating a UDP socket
    sock = socket(PF_INET, SOCK_DGRAM, 0);
    if (sock == INVALID_SOCKET) error_handling("socket() error");

    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;

    // Scenario 1: Sending to a specific remote machine
    addr.sin_addr.s_addr = inet_addr(remoteIP);
    addr.sin_port = htons(remotePort);

    if (connect(sock, (struct sockaddr*)&addr, sizeof(addr)) == SOCKET_ERROR)
        error_handling("connect() error");

    if (send(sock, message, strlen(message), 0) == SOCKET_ERROR)
        error_handling("send() error");

    // Scenario 2: Sending using broadcast
    char *broadcastIP = "10.22.73.189"; // Set this to your broadcast IP
    addr.sin_addr.s_addr = inet_addr(broadcastIP);

    if (setsockopt(sock, SOL_SOCKET, SO_BROADCAST, (char*)&broadcast, sizeof(broadcast)) == SOCKET_ERROR)
        error_handling("setsockopt() error");

    if (sendto(sock, message, strlen(message), 0, (struct sockaddr*)&addr, sizeof(addr)) == SOCKET_ERROR)
        error_handling("sendto() error");

    closesocket(sock);
    WSACleanup();
    return 0;
}
