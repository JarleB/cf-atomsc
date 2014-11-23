package main

// PKG_CONFIG_PATH="$HOME/homebrew/Homebrew-homebrew-70aebf8/Cellar/zeromq22/2.2.0/lib/pkgconfig/" go install github.com/JarleB/cf-atomsc

import     "fmt"
import     zmq "github.com/pebbe/zmq2"

func main() {
//    context, _ := zmq.NewContext(zmq.REQ)
    socket, _ := zmq.NewSocket(zmq.REQ)
//    defer context.Close()
    defer socket.Close()

    fmt.Printf("Connecting to hello world server…")
    socket.Connect("tcp://nasse:5555")

    for i := 0; i < 10; i++ {
        // send hello
        msg := fmt.Sprintf("Hello %d", i)
        socket.Send(msg,0)
        println("Sending ", msg)

        // Wait for reply:
        reply, _ := socket.Recv(0)
        println("Received ", string(reply))
    }
}
