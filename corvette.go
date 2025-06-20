package main

//pacote fmt é da standard lib (pacotes padrões) do go para a função print
import (
	"fmt"
	"time" //import utilizado no switch case
)

// constantes || não podem ser alteradas depois de declaradas || também é possível declarar sem o tipo, o go define o tipo automaticamente
const nome = "Lais"
const idade = 26

// criando uma struct
type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

// struct dentro da struct
type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

// struct para o metodos || avancado
type Pessoa struct {
	Nome  string
	Idade int
}

// func para a struct dos metodos
func (p Pessoa) Apresentar() { //pode interagir com a struct Pessoa || o "*" é uma referencia, quer dizer que vai receber a struct original!!! e não uma copia || sem o "*" utiliza uma copia
	p.Nome = "joana"                                                     //cria uma copia da struct || nao muda valor original
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade) //tem acesso aos campos da struct
}

// func para as goroutines
func exibirMsg() {
	fmt.Println("Olá de uma goroutine!")
}

// func para o goroutines sem ordem
func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond) //bloqueia o codigo - não deixa passar para o proximo enquanto não terminar este antes
	}
}

func sayWorld() {
	for i := 0; i < 3; i++ {
		fmt.Println("World")
		time.Sleep(150 * time.Millisecond) //bloqueia o codigo - não deixa passar para o proximo enquanto não terminar este antes
	}
}

// ex de goroutine pipeline || producer/consumer
func producer(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func consumer(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Consumer finalizado!")
}

// no terminal: colocar "cd/examples" e depois rodar
// rodar: go run corvette.go
// binario: go build -o hello
// criar go mod.. go mod init <module-name>
// Para executar cada bloco, descomentar antes e comentar novamente para executar um proximo bloco.
func main() {
	//printando a constante
	fmt.Println("Nome:", nome, "\nIdade:", idade)

	/*

		meet.SayHello()
		meet.Say("Olá Lais!")

	*/

	/*

		//criando uma variavel || não é necessário declarar o tipo da variável || pode ser passado o valor ou não, o go decide o valor inicial (int/string/float/bool)
		//declaração curta é ":=" || a variavel assumi o tipo do valor que foi atribuído a ela
		var texto1 string = "heeey"
		texto := "ola"
		texto += " tudo bem?"
		fmt.Println(texto)

	*/

	/*

		//operações matemáticas com inteiros
		//é possível definir se o inteiro é de 8, 16, 32, 64 bits (int32 / int64 / int8) || se não é definido, ele usa o padrão do sistema
		var idade int = 30
		var contador int32 = 2
		var indice int8 = 1

		fmt.Println("Idade:", idade)

	*/

	/*

		//números float || é possível definir se o float é de 32 ou 64 bits (float32 / float64) || se não é definido, ele usa o padrão do sistema
		// var floatNumber float32 = 1.1
		var pi float64 = 3.14
		var raio float64 = 2.5
		var area = pi * (raio * raio)

		// fmt.Println("floatNumber:", floatNumber)
		fmt.Println("Área do círculo:", area)

	*/

	/*

		//tipo bool || é usado para representar verdadeiro ou falso || é usado em condições, loops, etc || se não passa nenhum valor, o go assume o valor padrão (false) || pode adicionar expressões lógicas, como (true && false) || (true || false)
		var maior bool = 10 > 5
		var menor bool = 10 < 5
		fmt.Println("10 é menor que 5?", menor)
		fmt.Println("10 é maior que 5?", maior)

	*/

	/*

		//tipo string || é usado para representar texto || é possível manipular o texto, como concatenar, dividir, etc
		var hello string = "Olá, mundo!"
		var question string = "Como vai?"

	*/

	/*

		//somando ou concatenação
		var meet = hello + question
		fmt.Println(meet)
		fmt.Println(strings.ToUpper(meet)) //deixando todas as letras maiusculas || pacote "strings" é da standard lib do go
		fmt.Println(strings.Contains(meet, "mundo")) //verificando se alguma palavra existe na string

	*/

	/*

		//array
		var gavetas [2]string
		gavetas[0] = "copos" //preenchendo o array
		gavetas[1] = "panos"
		fmt.Println(gavetas[0]) //lendo o array

	*/

	/*

		//slices
		var gavetas []string
		gavetas = append(gavetas, "copos", "panos", "pratos") //append aumentar a capacidade do slice
		gavetas = gavetas[:2] //tirando um item do slice
		//fmt.Println(len(gavetas)) retorna o tamanho do slice || pode acessar até o numero 2 (pois começa com 0)
		fmt.Println(gavetas[1:2]) //tomar cuidado para não acessar um indice que não existe || dividindo o slice slice[x:x-1]

	*/

	/*

		//map
		var pessoas = map[string]int{}
		pessoas["Lais"] = 26 //chave e valor
		pessoas["Maria"] = 30

		//se ok por verdadeiro, ele entra no if
		//ok é um booleano que indica se existe ou não
		if idade, ok := pessoas["Lais"]; ok {
			fmt.Println("Pessoa existe no map", idade, ok)
		} else {
			fmt.Println("Pessoa não existe no map") //se quiser saber somente a idade || se colocar um nome que não existe, retorna 0
		}
		//funcao delete
		delete(pessoas, "Maria")
		fmt.Println(pessoas)

	*/

	/*
		//if else - avaliar se algo é verdadeiro ou falso
		nota := 75

		if nota >= 90 {
			fmt.Println("Aprovado com distinção")
		} else if nota >=70 {
			fmt.Println("Aprovado")
		} else {
			fmt.Println("Reprovado")
		}

	*/

	/*
		//declaração curta de uma variavel em uma expressão if
		if err := thisIsAnError(); err != nil { //só fica disponivel dentro desse bloco
			fmt.Println(err.Error())
		}

		//retornando um erro
		func thisIsAnError() error {
			return errors.New("Isto é um erro")
		}

	*/

	/*
		//declaração curta de uma variavel em uma estrutura de dados || map
		players := map[string]int{
		    "lais": 26,
		}

		//ver se o OK é vdd
		if value, ok := players["lais"]; ok {  //no map é possivel verificar se o valor existe ou não
			fmt.Println("pontos:", value, ok)
		}

	*/

	/*
		//switch case - forma mais enxuta de uma sequencia de ifs || intuito de não usar varios ifs/elses
		//sisteminha de dia da semana
		fmt.Println("Quando é sábado?")
		today := time.Now().Weekday()

		switch time.Saturday {  //retornar se hoje é sabado
			case today + 0:
				fmt.Println("é hoje")
			case today + 1:
				fmt.Println("é amanhã")
			case today + 2:
				fmt.Println("é depois de amanhã")
			default:     //para encerrar o switch
				fmt.Println("Tá longe ainda...")
		}

	*/

	/*
		//for    inicializacao; expressao; fim iteracao
		sum := 0
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			sum += i
		}
		fmt.Println(sum)

		//while - nao tem while pq consegue simular apenas com o for
		sum := 0
		for sum < 20 {
			fmt.Println("loop")
			sum += 2
		}
		fmt.Println(sum)

		//for sobre slices - listas com muitos items do mesmo tipo
		nums := []int{1, 2, 3, 4, 5}
		for i := 0; i < len(nums); i++ {
			fmt.Println(nums[i]) //o "i" acessa os dados
		}
	*/

	/*
		//range - tbm é uma maneira de fazer loops	|| loops em estruturas de dados || retorna uma chave e valor da estrutura
		// nums := []int{1, 2, 3, 4, 5}  //é como se tivesse chaves invisiveis na estrutura de slice
		nums := []string{"lais", "joao", "nath"}
		for key, value := range nums {
			fmt.Println(key, value)
		}

		//exemplo de uso em um map
		users := map[string]string {
			"nome": "lais",
			"idade": "26",
		}
		for _, value := range users {  //é possivel colocar um "_" pq o go não deixa colocar codigo num arquivo que NAO vamos usar || não complica ele na memoria com o "_" || da pra acessar somente a chave tbm para recriar array, ordenar e etc
			fmt.Println(value) //retornando soment o valor
		}
	*/

	/*
		//structs - maneira que a gente cria/compoe tipos no go (estruturas) || como se fossem classes e metodos
		cliente1 := Cliente {
			Nome: "lais",
			Idade: 26,
			Endereco: Endereco {
				Rua: "nao te interessa",
				Numero: 123,
				Estado: "SP",
			},
		}

		cliente2 := Cliente {
			Nome: "thais",
			Idade: 29,
		}

		cliente2.Email = "thais@google.com"
		cliente1.Endereco.Numero = 124

		fmt.Println(cliente1.Nome)
		fmt.Println(cliente2)
	*/

	/*
		//funcoes - podem receber n parametros declarados || caso não esteja declarado, ele entende que é void e retorna nada
		//closeries - consegue atribuir funcoes a variaveis || cria dentro de um determinado escopo para utilizar dentro do prorama
		var fixo = 4
		multiplica := func(x int) int {          //nao tem nome mas esta em uma variavel
			return x * fixo
		}

		resultado := multiplica(5)
		fmt.Printl(resultado)

		soma(2, 2)

		func soma(a, b int) {    //soma a + b e retorna um inteiro || da pra deixar ela privada (soma) ou publica (Soma)
			fmt.Println(a + b)
		}
	*/

	/*
		//metodos - funcoes amarradas a structs || a struct está acima da func main
		p1 := Pessoa{Nome: "lais", Idade: 26} //struct original
		p1.Apresentar()
		fmt.Println(p1.Nome)
	*/

	//avançado
	/*
		//ponteiros - quando quero falar que a variavel é do tipo ponteiro "*" guarda um endereco e nao um valor || quando quero imprimir o endereco "&" antes da variavel
		var p1 Pessoa = Pessoa(Nome: "lais")
		var p2 Pessoa = Pessoa (Nome: "thais")
		fmt.Println(p1)

		//tipo ponteiro para um valor pessoa passando um endereco || sem um ponteiro, ele faz uma copia do valor original
		var p3 *Pessoa = &p1
		p3.Nome = "vanessa" //alterando valor original de p1
		fmt.Println(p1)
		fmt.Println(p2)
	*/

	/*
		//Goroutines - unidades de consoles leves que estao rodando ao mesmo tempo e disputando o mesmo recurso || rodando de maneira concorrente || channels é o meio de comunicacao delas
		go exibirMsg()
		go exibirMsg()
		time.Sleep(1 * time.Second)
		fmt.Println("Olá main function!")
	*/

	/*
		//Goroutines sem ordem || não é garantido a execucao em ordem
		go sayHello() //chamando como uma goroutine
		go sayWorld()
		time.Sleep(1 * time.Second)
	*/

	/*
		//channels - é um tubo que leva as goroutines (leitura e escrita) || maneira de como as goroutines se comunicam || pode ser declarado como variavel
		ch := make(chan int, 3)    //canal que recebe inteiros || le 3 vezes e se se inserir uma 4 vez, só vai ler se uma outra goroutine leia pelo menos o primeiro valor

		//se está escrevendo, coloca "<-" do lado direito || a goroutine fica bloqueada até que outra goroutine leia o valor inserido
		go func() {
			for i :0; i < 5; i++ {
				ch <- i
			}
			close(ch) 	//é muito importante fechar o canal depois de usar!! || previni deadlocks
			fmt.Println("Escrita finalizada!")
		}()

		time.Sleep(time.Second * 1)
		//se está lendo, coloca "<-" do lado esquerdo || ordem de uma fila
		for valor := range ch {
			fmt.Println("Leitura", valor)
		}

		//pipeline producer e consumer
		ch := make(chan int)

		//não é garantida a ordem em que será executado || podem ter vários canais lendo
		go producer(ch)
		go consumer(ch)
		go consumer(ch)

		time.Sleep(time.Second * 1)
	*/

}
