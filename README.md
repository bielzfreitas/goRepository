# goRepository

## Curso de Go da RocketSeat

#### main.go
- Pacote fmt é da standard lib (pacotes padrões) do go para a função print
- constantes || não podem ser alteradas depois de declaradas || também é possível declarar sem o tipo, o go define o tipo automaticamente
- para rodar: go run main.go || binario: go build -o hello
- criar go mod.. go mod init <module-name>
- rodar no terminal: go run main.go
- criando uma variavel || não é necessário declarar o tipo da variável || pode ser passado o valor ou não, o go decide o valor inicial (int/string/float/bool)
- declaração curta é ":=" || a variavel assumi o tipo do valor que foi atribuído a ela
- é possível definir se o inteiro é de 8, 16, 32, 64 bits (int32 / int64 / int8) || se não é denifido, ele usa o padrão do sistema
- essa adaptação é feita por questão de performance do sistema || as variaveis ocupam menos espaço na memória
- números float || é possível definir se o float é de 32 ou 64 bits (float32 / float64) || se não é definido, ele usa o padrão do sistema
- são usados para representar dinheiro, unidades de medida, calculos científicos, etc
- também é possivel manipular a quantidade de memoria utilizada
- não é possivel rodar dois tipos (float32 e float64) juntos, é necessário fazer a conversão de tipos
- tipo bool || é usado para representar verdadeiro ou falso || é usado em condições, loops, etc
- se não passa nenhum valor, o go assume o valor padrão (false)
- pode adicionar expressões lógicas, como (true && false) || (true || false)
- tipo string || é usado para representar texto || é possível manipular o texto, como concatenar, dividir, etc
- array em go || array é uma estrutura de dados utilizado para poder aplicar algumas operações de algoritmos, na qual temos uma lista (espaço reservado para items) do mesmo tipo e 
são indexados (um endereço para cada item começando do zero)
- tamanho fixo, não é possivel mudar
- slices || tamanho flixiveis (aumenta ou diminui conforme a necessidade do programa) || basicamente um array
- append aumentar a capacidade do slice por debaixo dos panos e inserir o item sempre no final da lista || popula o slice e ajusta a capacidade automaticamente
- tomar cuidado para não acessar um indice que não existe || dividindo o slice slice[x:x-1] || é usado para reseolver desafios de algoritmos que envolvem algumas estruturas de dados || se não passar nada 
no indice "[:2]", ele entende que é zero (de zero até 2) || se não passar no valor final, ele considera o tamanho total do slice
- Maps são uma estrutura chave-valor || estrutura de dados rápido || chave unica 

#### meet.go / say.go
- Qualquer função que comece com letra maiúscula é pública
- Qualquer função que comece com letra minúscula é privada

#### go.mod
- o padrão para se criar um mod é: go mod init github.com/(seu id)/curso-go
- usado para rodar o projeto
