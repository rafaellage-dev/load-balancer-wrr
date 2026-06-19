# Projeto Prático: Desenvolvimento de Balanceador de Carga

**Estudante:** Rafael Nogueira Lage
**Cenário: ** Cenário B - Weighted Round-Robin (Round-Robin com Pesos)
**Professor Orientador:** Kaynan Macedo de Souza

---

## 1. Descrição e Contexto do Projeto
Este projeto consiste no desenvolvimento completo e prático de um **Load Balancer (Balanceador de Carga) de Camada 7 (HTTP)** construído puramente na linguagem Go, atuando como um Proxy Reverso.

O foco central deste repositório foi solucionar os desafios do **Cenário B (Weighted Round-Robin)**, mapeando cenários reais de engenharia de confiabilidade (SRE/Infraestrutura) onde os servidores de destino possuem capacidades computacionais assimétricas (heterogeneidade de hardware).

---

## 2. Tecnologias Utilizadas e Requisitos Atendidos
* **Linguagem Go (Pure Go):** Lógica desenvolvida sem bibliotecas externas de roteamento. Uso de primitivas nativas de concorrência como `sync.RWMutex` para exclusão mútua concorrente e exclusão em leitura de estados e `sync/atomic` para operações atômicas seguras no ponteiro circular de roteamento.
* **Mecanismo Ativo de Health Check:** Uma Goroutine em background executa testes de integridade assíncronos a cada 5 segundos na rota `/health` de cada nó. Se um nó for detectado como instável ou inativo, ele é excluído da lista de tráfego instantaneamente até sua plena recuperação.
* **Infraestrutura em Containers:** Utilização do `Docker` e `Docker Compose` isolando 3 réplicas de nós simulados com variáveis de ambiente distintas e 1 container isolado para o Balanceador de Carga.

---

## 3. Como Executar a Infraestrutura (Comando Único)

Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina de testes. Na raiz do projeto, execute o comando abaixo para realizar o build das imagens personalizadas e subir o ecossistema completo:

```bash
docker-compose up --build
