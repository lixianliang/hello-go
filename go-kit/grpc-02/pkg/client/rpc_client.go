package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	//"gokit/pkg/model"
	"github.com/lixianliang/hello-go/go-kit/grpc-02/pkg/model"
)

func NewClient() {
	log.Println("client: run client")
	conn, err := grpc.Dial(model.ServerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("client connect ok")
	uManagerServiceClient := model.NewUserManagerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	log.Println("client: mayun")
	user, err := uManagerServiceClient.GetUser(ctx, &model.GetUserRequest{Name: "mayun"})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("%+v\n", user)
		log.Printf("姓名: %s\n", user.Name)
		log.Printf("职位: %s\n", user.Title)
		log.Printf("公司: %s\n", user.Company)
	}

	log.Println("client: liyanhong升职为CEO")
	_, err = uManagerServiceClient.SetTitle(ctx, &model.SetTitleRequest{
		Name: "liyanhong", Title: "CEO",
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		user, err = uManagerServiceClient.GetUser(ctx, &model.GetUserRequest{Name: "liyanhong"})
		log.Printf("%+v\n", user)
		log.Printf("name: %s zhiwei:%s gongsi:%s\n", user.Name, user.Title, user.Company)
	}

	log.Println("client: pony")
	_, err = uManagerServiceClient.Dispatch(ctx, &model.DispatchRequest{
		Name: "pony", Company: "shenzhen",
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		user, err = uManagerServiceClient.GetUser(ctx, &model.GetUserRequest{
			Name: "pony",
		})
	}
	log.Printf("name: %s zhiwei:%s gongsi:%s\n", user.Name, user.Title, user.Company)
}
