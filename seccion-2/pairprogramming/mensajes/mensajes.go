package main
import (
	"fmt"
	"sync"
	"time"
	"os"
	"bufio"
)
type Message struct {
    Type    string
    Content string
}

func main() {
	emailChan := make(chan Message)
	signalChannel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var seguirProcesando bool = true
		for seguirProcesando {
			select {
			case msg, ok := <-emailChan:
				if !ok {
					fmt.Println("Canal cerrado")
					seguirProcesando = false
					break
				}
				fmt.Printf("Procesando mensaje de notificación correo electronico: %s\n", msg.Content)
				time.Sleep(3 * time.Second)
				fmt.Printf("se termino de procesar el mensaje correo electronico: %s\n", msg.Content)
				fmt.Println("Envio señal de continuar")
				signalChannel <- true
			}
		}
	}()
	
	go func(emailChan chan<- Message) {
		scanner := bufio.NewScanner(os.Stdin)
		for true {
			// paremos acá!
			fmt.Print("Ingrese su mensaje o presione CTR+D para salir:")
			hasMoreInput := scanner.Scan()
			if !hasMoreInput {
				fmt.Println("SALIENDO")
				break
			}
			inputStr := scanner.Text()
			message := Message{Type: "email", Content: inputStr}
			emailChan <- message

			var signal bool
			signal = <-signalChannel
			fmt.Println("Llego señal de continuar", signal)
		}
		close(emailChan)
	}(emailChan)
	wg.Wait()
	fmt.Println("Todos los mensajes han sido procesados.")
}
