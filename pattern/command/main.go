/*
Применимость
* Когда хотим параметризовать объекты выполняемым действием
* Когда хотим ставить операции в очередь, выполнять их по расписанию или передавать по сети
* Когда нужна операция отмены

+ OCP
+ отложенный запуск операций
+ позволяет реализовать простую отмену и повтор операций.
+ позволяет собирать сложные команды из простых.

- усложняет код из-за введения множества доп классов
*/

package main

func main() {
	tv := &tv{}
	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommand}
	onButton.press()

	offButton := &button{command: offCommand}
	offButton.press()

}
