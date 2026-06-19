version: '3.8'

services:
# Servidor Heterogêneo A - Alta Capacidade (Peso 3)
backend_potente:
build: ./backend
container_name: node_potente_hardware_alto
environment:
- SERVER_NAME=Servidor_A_Processador_8_Cores_Memoria_32GB_(Peso_3)
expose:
- "8080"

# Servidor Heterogêneo B - Média Capacidade (Peso 2)
backend_medio:
build: ./backend
container_name: node_medio_hardware_padrao
environment:
- SERVER_NAME=Servidor_B_Processador_4_Cores_Memoria_16GB_(Peso_2)
expose:
- "8080"

# Servidor Heterogêneo C - Baixa Capacidade (Peso 1)
backend_fraco:
build: ./backend
container_name: node_fraco_hardware_antigo
environment:
- SERVER_NAME=Servidor_C_Processador_2_Cores_Memoria_4GB_(Peso_1)
expose:
- "8080"

# O Balanceador de Carga Customizado atuando na porta pública 8000
loadbalancer:
build: ./balancer
container_name: custom_load_balancer_wrr
ports:
- "8000:8000"
depends_on:
- backend_potente
- backend_medio
- backend_fraco
