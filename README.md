# Projeto Prático: Desenvolvimento de Balanceador de Carga

**Estudante:** Rafael Nogueira Lage

**Cenário:** Cenário B - Weighted Round-Robin (Round-Robin com Pesos)

**Professor Orientador:** Kaynan Macedo de Souza

---

Este repositório apresenta a implementação de um balanceador de carga desenvolvido na linguagem Go, focado no atendimento dos requisitos estabelecidos para o Cenário B da disciplina. A arquitetura do projeto foi construída utilizando Docker Compose para orquestrar quatro contêineres: um atuando como o Load Balancer principal, exposto na porta 8000, e três instâncias de servidores backend, operando internamente na porta 8080, que simulam diferentes capacidades de processamento.

Para a distribuição do tráfego, foi adotado o algoritmo Weighted Round Robin (WRR). A configuração atribui pesos específicos a cada nó da rede, de modo a rotear as requisições de forma proporcional à capacidade simulada de cada máquina. Dessa forma, definiu-se o Servidor Potente com peso 3, o Servidor Médio com peso 2 e o Servidor Fraco com peso 1.

Para executar a aplicação, é necessário ter o Docker e o Docker Compose instalados no ambiente. A inicialização da infraestrutura é feita executando o comando `docker compose up --build -d` na raiz do diretório. Com os contêineres em execução, a validação da lógica do algoritmo pode ser realizada por meio de um teste de carga contendo seis requisições.

Executando o comando `for i in {1..6}; do curl -s http://localhost:8000; echo ""; done` no terminal, é possível observar o roteamento exato estabelecido pelos pesos. O resultado esperado exibirá a seguinte distribuição, comprovando a eficácia do balanceamento de carga implementado:

```text
=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Potente (Peso 3)
Status: Processado com Sucesso

=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Potente (Peso 3)
Status: Processado com Sucesso

=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Potente (Peso 3)
Status: Processado com Sucesso

=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Medio (Peso 2)
Status: Processado com Sucesso

=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Medio (Peso 2)
Status: Processado com Sucesso

=== Resposta do Servidor Backend ===
Identificação do Nó: Servidor Fraco (Peso 1)
Status: Processado com Sucesso

```
