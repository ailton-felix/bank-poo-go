# bank-poo-go

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
