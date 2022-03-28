package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func ProtobufCase() {

	elliot := gen.Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(&elliot) //по уже сгенерированной структуре, мы маршалим сообщение и получем масив байт
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// print object
	fmt.Println(data)

	newElliot := &gen.Person{}
	err = proto.Unmarshal(data, newElliot) // обратный ход, анмаршалинг аналогичен остальным протоколам
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	//печатаем поля обьекта
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}
