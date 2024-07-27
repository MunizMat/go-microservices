## Microservices in practice 

- [English](#english)
- [Português](#português)


### English
This is a simple implementation of a user creation flow that uses microservice architecture to enable two applications to communicate with each other through a message broker, which, in this case, is RabbitMQ. This repository consists of two main directories ```/user``` and ```/email```.

Both directories are applications built with Go that run independently. The user service is a also producer that publishes messages to RabbitMQ whenever a user is created. The messages are held in a "email" queue from where the email service consumes messages. 

This is a very simple implementation, and my focus here, besides familiarizing myself with Go, was to better understand microservice architecture and messaging systems through a real implementation


### Português
Esta é uma implementação simples de um fluxo de criação de usuário que utiliza arquitetura de microsserviços para permitir que duas aplicações se comuniquem por meio de um sistema de mensageria, que, neste caso, é o RabbitMQ. Este repositório consiste em dois diretórios principais ```/user``` e ```/email```.

Ambos os diretórios são aplicações construídas com Go que rodam de forma independente. O aplicação de usuário também é um produtor que publica mensagens no RabbitMQ sempre que um usuário é criado. As mensagens são mantidas em uma fila "email" de onde a aplicação de email consome as mensagens. 

Esta é uma implementação muito simples, e meu foco aqui, além de me familiarizar com Go, foi entender melhor a arquitetura de microsserviços e sistemas de mensageria através de uma implementação real