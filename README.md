# TinyGo: Golang para microcontroladores e Arduino

**TinyGo**, um compilador Go que permite programar microcontroladores como o Arduino usando a linguagem Go.

## O que é TinyGo?

TinyGo é uma variante do compilador Go, projetada para gerar código compacto e eficiente para dispositivos com recursos limitados, como microcontroladores. Ele abre um novo mundo de possibilidades para desenvolvedores Go que desejam explorar o universo da programação embarcada.

## Vantagens de usar TinyGo com Arduino:

- **Linguagem Go**: Aproveite a simplicidade, legibilidade e recursos poderosos da linguagem Go, como concorrência com goroutines, para programar seus projetos Arduino.
- **Código Compacto**: TinyGo gera binários menores, otimizados para o ambiente de microcontroladores, economizando espaço de armazenamento e memória.
- **Desempenho**: O código compilado com TinyGo é eficiente em termos de desempenho, permitindo que você crie aplicações mais complexas em microcontroladores.
- **Ecossistema Go**: Acesse bibliotecas e ferramentas do ecossistema Go, facilitando o desenvolvimento e a manutenção do seu código.
- **Suporte a Arduino**: TinyGo oferece suporte a várias placas Arduino populares, incluindo Arduino Uno, Nano33 IoT e outras.

## Como usar TinyGo com Arduino:

### 1. Instalação:
Siga as instruções de instalação do TinyGo no site oficial: [https://tinygo.org/getting-started/](https://tinygo.org/getting-started/).

### 2. Configurar o ambiente:
- Conecte sua placa Arduino ao computador.

### 3. Escrever o código em Go:
Utilize as bibliotecas do TinyGo para interagir com os periféricos do Arduino, como GPIOs, sensores e comunicação serial.

### 4. Compilar e fazer upload:

Use o comando a seguir para compilar e enviar o código para a placa Arduino:

```bash
tinygo flash -target <placa_arduino> <arquivo.go>
```

## Conectando Dispositivos USB ao WSL2 para usar Golang com Arduino

Este guia detalha como configurar o WSL2 no Windows 10/11 para conectar dispositivos USB, como um Arduino, e usá-los com Golang.

### Configuração

- **Windows 10 >**
- **WSL2 com Ubuntu 22.04.3 LTS**
- **Kernel WSL2: 5.15.90.1-microsoft-standard-WSL2**

### Passos

#### 1. Instalar o pacote USBIPd no Windows

Baixe o instalador (.msi) mais recente em: [USBIPD para Windows](https://github.com/dorssel/usbipd-win/releases/tag/v4.0.0)

#### 2. Instalar pacotes no WSL2/Ubuntu

##### Sincronizar o relógio entre o WSL2 e o Windows:
```bash
sudo hwclock -s
```

##### Atualizar pacotes e fontes
```bash
sudo apt update
```

##### Instalar pacotes necessários:

```bash
sudo apt install linux-tools-generic hwdata
sudo update-alternatives --install /usr/local/bin/usbip usbip /usr/lib/linux-tools/*-generic/usbip 20
```

#### 3. Conectar o dispositivo USB ao WSL2

- Conecte o dispositivo USB (ex: Arduino) à porta USB do computador.
- Certifique-se de que o WSL2 esteja em execução.

##### Obter o Bus ID do dispositivo:

1. Abra o PowerShell como administrador no Windows.
2. Execute o comando `usbipd list` para listar os dispositivos USB conectados e seus respectivos Bus IDs.

#### 4. Vincular o dispositivo (uma única vez por dispositivo):

##### No PowerShell (admin), execute:

```bash
usbipd bind --busid=2-2
```

##### Conectar o dispositivo ao WSL2 (requerido a cada reconexão):

No PowerShell (admin), execute:
```bash
usbipd attach --wsl --busid=1-2
```

####  Verificar a conexão no WSL2

##### Abra o terminal do WSL2 e execute:

```bash
lsusb
```

#### O dispositivo conectado deve aparecer na lista. Exemplo de saída do lsusb:

```bash
Bus 002 Device 001: ID 1d6b:0003 Linux Foundation 3.0 root hub
Bus 001 Device 003: ID 0483:374e STMicroelectronics STLINK-V3
```

#### Observações

- Este guia assume que você já possui o Golang instalado no WSL2.
- Após seguir estes passos, você poderá usar o Golang para programar seu Arduino no WSL2 como se estivesse conectado diretamente a ele.

