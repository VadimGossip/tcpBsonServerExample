# tcpBsonServerExample
An example of a basic TCP server which can receive BSON packets, process them in several threads (I use a special thread handler https://github.com/VadimGossip/tcpConHandler) and return the responses. Reading and writing are done independently.

In addition to this project I made a request generator which can also be used for mastering client-server interaction.
