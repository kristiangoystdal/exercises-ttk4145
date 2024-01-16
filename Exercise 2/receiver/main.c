#include <winsock2.h>
#include <ws2tcpip.h>
#include <stdio.h>
#include <stdlib.h>

int main() {
    WSADATA wsaData;
    int iResult = WSAStartup(MAKEWORD(2,2), &wsaData);
    if (iResult != NO_ERROR) {
        printf("WSAStartup failed: %d\n", iResult);
        return 1;
    }

    SOCKET listenSocket = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (listenSocket == INVALID_SOCKET) {
        printf("Error at socket(): %ld\n", WSAGetLastError());
        WSACleanup();
        return 1;
    }

    struct sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = INADDR_ANY;
    serverAddr.sin_port = htons(30000);

    if (bind(listenSocket, (SOCKADDR *) &serverAddr, sizeof(serverAddr)) == SOCKET_ERROR) {
        printf("Bind failed with error: %ld\n", WSAGetLastError());
        closesocket(listenSocket);
        WSACleanup();
        return 1;
    }

    if (listen(listenSocket, SOMAXCONN) == SOCKET_ERROR) {
        printf("Listen failed with error: %ld\n", WSAGetLastError());
        closesocket(listenSocket);
        WSACleanup();
        return 1;
    }

    printf("Waiting for client to connect...\n");

    SOCKET clientSocket = accept(listenSocket, NULL, NULL);
    if (clientSocket == INVALID_SOCKET) {
        printf("Accept failed with error: %ld\n", WSAGetLastError());
        closesocket(listenSocket);
        WSACleanup();
        return 1;
    }

    char recvbuf[512];
    int recvbuflen = 512;
    int bytesReceived = recv(clientSocket, recvbuf, recvbuflen, 0);
    if (bytesReceived > 0) {
        printf("Bytes received: %d\n", bytesReceived);
        printf("Message: %.*s\n", bytesReceived, recvbuf);

        const char *sendbuf = "Message received!";
        send(clientSocket, sendbuf, (int)strlen(sendbuf), 0);
    } else if (bytesReceived == 0) {
        printf("Connection closing...\n");
    } else {
        printf("recv failed with error: %d\n", WSAGetLastError());
    }

    closesocket(clientSocket);
    closesocket(listenSocket);
    WSACleanup();

    return 0;
}
