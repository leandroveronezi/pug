package console

import "syscall/js"

/*
console.error(object[, object...])
Emite uma mensagem de erro. Você pode usar substituição de string e outros argumentos adicionais com este método.
Consulte Uso de substituição de string.
*/
func Error(Value ...interface{}) {
	js.Global().Get("console").Call("error", Value...)
}

/*
console.log(object[, object...])
Utilizado para a emissão de informações de registro em geral. Você pode utilizar substituição de string e
outros argumentos com este método. Consulte Uso de substituição de string.
*/
func Log(Value ...interface{}) {
	js.Global().Get("console").Call("log", Value...)
}

/*
console.warn(object[, object...])
Emite uma mensagem de alerta. Você pode utilizar substituição de string e argumentos adicionais com este método.
Veja Uso de substituição de string.
*/
func Warn(Value ...interface{}) {
	js.Global().Get("console").Call("warn", Value...)
}

/*
console.info(object[, object...])
Informações de registro. Você pode utilizar substituição de string e outros argumentos com este método.
Consulte Uso de substituição de string.
*/
func Info(Value ...interface{}) {
	js.Global().Get("console").Call("info", Value...)
}

/*
Console.table()
Exibe dados, como objeto e array, como uma tabela.
*/
func Table(Value ...interface{}) {
	js.Global().Get("console").Call("table", Value...)
}

func Clear() {
	js.Global().Get("console").Call("clear")
}

/*
console.group(object[, object...])
Cria um novo grupo em linha e recua todas as mensagens seguintes para um nível de indentação superior.
Para voltar um nível, utilize groupEnd(). Consulte Uso de grupos no console.
*/
func Group(Value ...interface{}) {
	js.Global().Get("console").Call("group", Value...)
}

/*
console.groupCollapsed(object[, object...])
Cria um novo grupo em linha e recua todas as mensagens seguintes para um nível de indentação superior; ao contrário de
group(), o grupo em linha começa recolhido. Para revelar seu conteúdo, basta clicar no botão de revelação para expandí-lo.
Para recuar um nível, utilize groupEnd(). Consulte Uso de grupos no console.
*/
func GroupCollapsed(Value ...interface{}) {
	js.Global().Get("console").Call("groupCollapsed", Value...)
}

/*
console.groupEnd()
Sai do grupo em linha atual. Veja Uso de grupos no console.
*/
func GroupEnd() {
	js.Global().Get("console").Call("groupEnd")
}

/*
console.trace()
Emite um traçado de pilha. See Traçados de pilha.
*/
func Trace() {
	js.Global().Get("console").Call("trace")
}

/*
console.time(name)
Inicia um contador de tempo com o nome especificado no parâmetro name. Até 10.000 contadores de tempo podem ser rodados por página.
*/
func Time(Name string) {
	js.Global().Get("console").Call("time", js.ValueOf(Name))
}

/*
console.timeEnd(name)
Interrompe o contador de tempo especificado e emite o tempo e registros do contador de tempo em milisegundos desde o seu início.
Veja Contadores de Tempo.
*/
func TimeEnd(Name string) {
	js.Global().Get("console").Call("timeEnd", js.ValueOf(Name))
}

/*
console.count([label])
Mostra o número de vezes que esta linha foi chamada com a label fornecida.
*/
func Count(Label string) {
	js.Global().Get("console").Call("count", js.ValueOf(Label))
}

/*
console.dir(object)
Exibe uma listagem interativa das propriedades de um objeto JavaScript especificado. Esta listagem permite a você expandir
o objeto para visualizar o conteúdo de objetos filhos.
*/
func Dir(Object interface{}) {
	js.Global().Get("console").Call("dir", js.ValueOf(Object))
}

/*
https://developer.mozilla.org/pt-BR/docs/Web/API/Console

console.assert(expression, object[, object...])
Emite uma mensagem e traça a sequência de operações até o primeiro argumento for falso.


console.profile( [profileLabel] )
Inicia o JavaScript profiler. Você pode especificar qualquer label opcional para o perfil.


console.profileEnd()
Interrompe o profiler. Você pode ver o profile resultante no JavaScript profiler.
*/
