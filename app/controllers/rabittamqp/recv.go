package rabittamqp
 
 import (
     "log"
	  "github.com/streadway/amqp"
	  "github.com/julienschmidt/httprouter"
	  "net/http"
	  . "admin-mvc/app/utils"
	  "encoding/json"
 )
 
 func failOnError(err error, msg string) {
     if err != nil {
         log.Fatalf("%s: %s", msg, err)
     }
 }
 
 // 只能在安装 rabbitmq 的服务器上操作
 func Received(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
     conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
     failOnError(err, " 连接RabbitMQ失败")
     defer conn.Close()
 
     ch, err := conn.Channel()
     failOnError(err, "打开信道失败")
     defer ch.Close()
 
     q, err := ch.QueueDeclare(
         "firthello", // name
         false,   // durable
         false,   // delete when unused
         false,   // exclusive
         false,   // no-wait
         nil,     // arguments
     )
     failOnError(err, "声明队列失败")
 
     msgs, err := ch.Consume(
         q.Name, // queue
         "",     // consumer
         true,   // auto-ack
         false,  // exclusive
         false,  // no-local
         false,  // no-wait
         nil,    // args
     )
     failOnError(err, "注册消费者失败")
 

	
     forever := make(chan bool)
     
     go func() {
         for d := range msgs {
			 
			//log.Printf("Received a message: %s", d.Body)
			_data := JsonData{true, "", d.Body }
			data, _ := json.Marshal(_data)
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
         }
     }()
 
     log.Printf(" [*] 正在等待消息。退出按 CTRL+C 组合键")
	 <-forever
	
 }