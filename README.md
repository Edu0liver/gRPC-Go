# gRPC em Go (Golang) 🚀  

Este repositório contém exemplos e práticas de desenvolvimento utilizando **gRPC** com **Go (Golang)**.  
O objetivo é servir como um guia prático para quem deseja aprender e implementar gRPC em projetos reais.  

---

## 🧐 Sobre o Projeto  

O gRPC é um framework open-source desenvolvido pelo Google que utiliza **HTTP/2** e **Protocol Buffers** para comunicação eficiente entre serviços.  
Neste projeto, você encontrará:  

- Exemplos básicos de comunicação cliente-servidor usando gRPC.  
- Implementação de diferentes métodos gRPC (unário, streaming, bidirecional).  
- Integração com ferramentas como `protoc` e `protoc-gen-go`.  

---

## 🔧 Pré-requisitos  

Antes de começar, certifique-se de ter instalado:  

- **Go** (v1.18 ou superior)  
- **Protoc** (Protocol Buffers Compiler)  
- Plugins do gRPC para Go:  

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
