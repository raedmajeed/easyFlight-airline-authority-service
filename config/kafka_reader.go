package config

import (
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	SearchReader       *kafka.Reader
	SearchSelectReader *kafka.Reader
	svc                interfaces.AdminAirlineService
}

func NewKafkaReaderConnect(svc interfaces.AdminAirlineService) *KafkaReader {
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "search-flight-bs-6",
		GroupID: "search-request-6",
	})
	searchSelectReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "search-select-flight-bs-6",
		GroupID: "search-selected-6",
	})
	return &KafkaReader{
		SearchReader:       searchReader,
		svc:                svc,
		SearchSelectReader: searchSelectReader,
	}
}

//func KafkaReaders(ctx context.Context, svc interfaces.AdminAirlineService, ch chan os.Signal) {
//	kf := NewKafkaReaderConnect(svc)
//	//kfWriter := NewKafkaWriterConnect()
//	//go kf.SearchFlightRead(ctx, kfWriter)
//	//go kf.SearchSelectFlightRead(ctx)
//	//<-ch
//	if err := kf.SearchReader.Close(); err != nil {
//		log.Println("error closing search reader")
//	}
//	if err := kf.SearchSelectReader.Close(); err != nil {
//		log.Println("error closing search select reader")
//	}
//	log.Println("closed all kafka readers")
//}

//func (kf *KafkaReader) SearchFlightRead(ctx context.Context, kfWriter *KafkaWriter) {
//	newCont, cancel := context.WithCancel(ctx)
//	defer cancel()
//	log.Println("listening for search details from BOOKING SERVICE")
//	ch := make(chan kafka.Message)
//	go kf.SearchFlightReadDirect(ctx, ch, kfWriter)
//	message, _ := kf.SearchReader.ReadMessage(ctx)
//	log.Println("flight searching for -> ", string(message.Value))
//	ch <- message
//}

//func (kf *KafkaReader) SearchFlightReadDirect(ctx context.Context, ch chan kafka.Message, kfWriter *KafkaWriter) {
//	marshal, _ := json.Marshal(dom.KafkaPath{})
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("context cancelled, terminating")
//			kfWriter.SearchWriter.WriteMessages(ctx, kafka.Message{
//				Value: marshal,
//			})
//			break
//			//return
//		case message := <-ch:
//			log.Println("message received -> ", string(message.Value))
//			//_ = kf.SearchReader.CommitMessages(ctx, message)
//			//kf.svc.SearchFlightInitial(ctx, message)
//		}
//	}
//}
//
//func (k *KafkaReader) SearchSelectFlightRead(ctx context.Context) {
//	log.Println("listening for selected flight details from BOOKING SERVICE")
//	ch := make(chan kafka.Message)
//	go k.SearchSelectFlightReadDirect(ctx, ch)
//	message, err := k.SearchSelectReader.ReadMessage(ctx)
//	if err == nil {
//		log.Println("flight searching for -> ", string(message.Value))
//		ch <- message
//	}
//}
//
//func (k *KafkaReader) SearchSelectFlightReadDirect(ctx context.Context, ch chan kafka.Message) {
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("context cancelled, terminating")
//			return
//		case message := <-ch:
//			log.Println("message received -> ", string(message.Value))
//			k.svc.SearchSelectFlight(ctx, message)
//			_ = k.SearchSelectReader.CommitMessages(ctx, message)
//		}
//	}
//}
