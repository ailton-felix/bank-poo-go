# bank-poo-go

This is a simple banking application developed in Go. It provides functionalities to manage client accounts and perform banking transactions.

The main purpose of this repository is to practice object-oriented programming concepts in Go.

## Features

- **Client Management:** Allows creating and managing client accounts.
- **Account Operations:** Supports depositing money, paying bills, and transferring funds between accounts.
- **Interface Usage:** Demonstrates the use of interfaces in Go for managing different types of accounts.

## Usage

1. **Client and Account Creation:**
   The program creates client and account objects and performs operations like depositing money, paying bills, and transferring funds between accounts.

2. **Interface Usage:**
   The program demonstrates the use of interfaces in Go to manage different types of accounts (current account, saving account, etc.).

3. **Account Operations:**
   - Depositing Money: Deposits money into a client's account.
   - Paying Bills: Deducts the specified amount from the account balance for bill payment.
   - Transferring Funds: Transfers funds from one account to another.

## Requirements

- Go 1.11 or higher

## Conceitos Go

- A linguagem Go não possui o conceito de classe.
- Métodos são definidos de maneira parecida com funções, mas de uma maneira diferente.
    
    Por exemplo:
    ```Go
    func (c *CurrentAccount) DrawMoney(value float64) string
    ```
    O exemplo acima se trata de um método. Existe um `(c *CurrentAccount)` que se refere a um ponteiro para a instância criada da estrutura.

    #### Passagem de parâmetro por cópia
    
    Também é possível passar um valor removendo a assinatura do ponteiro `(c *CurrentAccount)` para `(c CurrentAccount)`:

    ```Go
    func (c CurrentAccount) DrawMoney(value float64) string
    ```

    Nesse caso, uma cópia do valor de `CurrentAccount` é passada para o método, sem alterar o valor do ponteiro. Portanto, precisamos ficar atentos, já que qualquer alteração que você faça em `c`, se passar por valor (ou cópia), não será refletida na fonte `c`.

- Use o comando bach `go mod init` na raiz do seu projeto para gerar o arquivo `.mod` com suas dependências.

- Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go. Java possui um conceito muito parecido, também chamado de interface. A grande diferença é que, em Go, um tipo implementa uma interface implicitamente.

## Fontes:
- [Alura](https://www.alura.com.br/)