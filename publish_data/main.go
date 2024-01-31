package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	cluster_id = "test-cluster"
	client_id  = "wb_publisher"
	channel    = "flow"
)

func main() {
	sc, err := stan.Connect(cluster_id, client_id)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	defer sc.Close()
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Printf("%s\n", "Выберите операцию")
		fmt.Printf("%s\n", "1. Сгенерировать валидное json сообщение и отправить в stan")
		fmt.Printf("%s\n", "2. Сгенерировать невалидное json сообщение и отправить в stan")
		fmt.Printf("%s\n", "3. Прекратить посылать сообщения и выйти из программы")
		scanner.Scan()
		value := scanner.Text()
		switch {
		case value == "1":
			SendMessageToStan(GenerateTrueJson(), sc)
		case value == "2":
			SendMessageToStan(GenerateFalseJson(), sc)
		case value == "3":
			os.Exit(0)
		default:
			fmt.Printf("%s", "Введите корректное значение")
			continue
		}
	}
}

func GenerateTrueJson() []byte {
	var ModelData Order
	valueOrd := reflect.ValueOf(&ModelData).Elem()
	valuePay := reflect.ValueOf(&ModelData.Payment).Elem()
	valueDel := reflect.ValueOf(&ModelData.Delivery).Elem()
	FillStructureRandomValue(&valueOrd)
	FillStructureRandomValue(&valuePay)
	FillStructureRandomValue(&valueDel)
	size_items := generateRandomInt(1, 3)
	for i := 0; i < size_items; i++ {
		var item Item
		valueIt := reflect.ValueOf(&item).Elem()
		FillStructureRandomValue(&valueIt)
		ModelData.Items = append(ModelData.Items, item)
	}
	data, _ := json.MarshalIndent(ModelData, "", " ")
	fmt.Printf("%s\n", data)
	return data
}

func GenerateFalseJson() []byte {
	var ModelData FalseOrder
	valueFalse := reflect.ValueOf(&ModelData).Elem()
	valueDelivery := reflect.ValueOf(&ModelData.Delivery).Elem()
	FillStructureRandomValue(&valueDelivery)
	FillStructureRandomValue(&valueFalse)
	data, _ := json.MarshalIndent(ModelData, "", " ")
	fmt.Printf("%s\n", data)
	return data
}

func SendMessageToStan(data []byte, stanConnection stan.Conn) error {
	return stanConnection.Publish(channel, data)
}

func generateRandomString() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(b)
}

func generateRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func generateRandomFloat() float32 {
	return float32(rand.NormFloat64())
}

func FillStructureRandomValue(refl *reflect.Value) {
	for i := 0; i < refl.NumField(); i++ {
		fieldValue := refl.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(generateRandomString())
		case reflect.Int:
			fieldValue.SetInt(int64(generateRandomInt(1, 100)))
		case reflect.Float32:
			fieldValue.SetFloat(float64(generateRandomFloat()))
		}
	}
}
