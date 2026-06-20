# Projeto Prático: Desenvolvimento de Balanceador de Carga

**Estudante:** Rafael Nogueira Lage

**Cenário:** Cenário B - Weighted Round-Robin (Round-Robin com Pesos)

**Professor Orientador:** Kaynan Macedo de Souza

---

Este repositório contém a implementação prática de um balanceador de carga desenvolvido do zero na linguagem Go, atendendo aos requisitos do Cenário B da disciplina, que foca no algoritmo Weighted Round-Robin. O objetivo principal do projeto é aplicar os conceitos de distribuição de tráfego, alta disponibilidade e resiliência em sistemas distribuídos, observando o comportamento das requisições diante da heterogeneidade dos nós de destino. A arquitetura foi isolada utilizando a infraestrutura do Docker Compose, sendo composta por um contêiner que roda o balanceador atuando como proxy reverso na porta 8000 e três contêineres que executam instâncias da aplicação de backend operando internamente na porta 8080.

A distribuição de carga foi programada utilizando recursos nativos de concorrência do Go, utilizando exclusão mútua por meio de mutexes para garantir o controle de estado seguro da aplicação durante o tráfego simultâneo. Para refletir a heterogeneidade de hardware proposta pelo cenário, foram configurados pesos proporcionais à capacidade de cada máquina, sendo o servidor potente definido com peso três, o servidor médio com peso dois e o servidor fraco com peso um. Complementando as exigências técnicas, um mecanismo de checagem de saúde, conhecido como health check, foi implementado utilizando goroutines para rodar continuamente em segundo plano a cada cinco segundos. Essa rotina verifica o status de conectividade de cada nó e altera o estado interno do pool de servidores caso alguma máquina fique indisponível, impedindo o envio de tráfego para nós que saíram do ar.

Para inicializar a infraestrutura de teste de forma local com um único comando, o usuário precisa ter o Docker e o Docker Compose instalados no ambiente. No terminal, a partir do diretório raiz onde se encontram o arquivo docker-compose.yml e os arquivos Dockerfile, deve-se executar o comando "docker compose up --build -d". Este processo realiza a compilação do código fonte de forma isolada e sobe toda a rede de contêineres em segundo plano.

A validação do redirecionamento e da proporcionalidade dos pesos pode ser feita disparando uma sequência de seis requisições HTTP seguidas contra o balanceador de carga. Para isso, basta executar no terminal o comando "for i in {1..6}; do curl -s http://localhost:8000; echo ""; done". O comportamento esperado do algoritmo sob condições normais de funcionamento demonstrará o servidor potente respondendo exatamente três vezes, o servidor médio respondendo duas vezes e o servidor fraco respondendo uma vez, comprovando visualmente a precisão da lógica do Weighted Round-Robin.

A análise de comportamento do algoritmo diante do cenário proposto indica que, ao contrário do modelo padrão que distribui requisições de forma cega, o Weighted Round-Robin lida com a heterogeneidade de hardware alocando mais conexões para os computadores com maior capacidade, reduzindo o gargalo e otimizando a vazão total do sistema. Diante de cenários reais de falha, o comportamento também foi validado ao interromper o funcionamento do servidor médio de propósito através do comando "docker compose stop backend_medio". Após o intervalo da rotina de health check, que detecta a ausência de resposta do nó, o algoritmo adaptou-se dinamicamente ao ambiente alterado, redirecionando as novas requisições da fila exclusivamente entre os servidores potente e fraco, sem gerar perdas de requisições ou erros de conexão para o cliente final.
