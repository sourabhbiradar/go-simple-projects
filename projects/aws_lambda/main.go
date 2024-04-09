package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct{
	Name string `json:"whats your name?"`
	Age int `json:"how old are you?"`
}

type Response struct{
	Message string `json:"Answer :"`
}

func main (){
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(event MyEvent)(Response,error){
	return Response {Message:fmt.Sprintf("%s is %d years old!",event.Name,event.Age)},nil
}