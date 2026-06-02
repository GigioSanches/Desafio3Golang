package graphql

import (
	"net/http"

	"github.com/GigioSanches/Desafio3Golang/internal/domain"
	"github.com/GigioSanches/Desafio3Golang/internal/usecase"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewGraphQLHandler(listOrdersUC *usecase.ListOrdersUseCase) http.Handler {
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
			"description": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"price":       &graphql.Field{Type: graphql.NewNonNull(graphql.Float)},
			"created_at":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"listOrders": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(orderType))),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					orders, err := listOrdersUC.Execute()
					if err != nil {
						return nil, err
					}
					return convertOrders(orders), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h
}

func convertOrders(orders []domain.Order) []map[string]interface{} {
	converted := make([]map[string]interface{}, 0, len(orders))
	for _, o := range orders {
		converted = append(converted, map[string]interface{}{
			"id":          o.ID,
			"description": o.Description,
			"price":       o.Price,
			"created_at":  o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return converted
}
