package rabittamqp

 import (
     "log" 
	 "github.com/streadway/amqp"
	 "github.com/julienschmidt/httprouter"
	 "net/http"
	 . "admin-mvc/app/utils"
 )
 

 
 // 只能在安装 rabbitmq 的服务器上操作
 func SendMsg(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	if r.Method == "GET" { 
		//P("to SignIn page view ===>>>")
		GenerateHTML(w, "admin", "layouts/nilLayout", "admin/send" )
	}
	
	// POST Method admin signin auth
	if r.Method == "POST" {
		err := r.ParseForm()
		message := r.PostFormValue("message")
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "连接RabbitMQ失败")
		// P("Failed to connect to RabbitMQ >>>")
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
	
		body := message + "Hello world !"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "发布消息失败")
		http.Redirect(w, r, "/send", 302)
	}
 }