package server

import (
	"context"
	"net/http"
	ch "xti-gateway-graph-go/cache"
	"xti-gateway-graph-go/pkg/entities"
	"xti-gateway-graph-go/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

var (
	cache   ch.Cache
	err     error
	headers entities.Headers
)

func server() *runtime.ServeMux {
	return runtime.NewServeMux(
		// handle incoming headers
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			clientId := request.Header.Get("x-client-id")
			clientSecret := request.Header.Get("x-client-secret")
			realm := request.Header.Get("x-realms")
			generateHeaders(&headers, request)
			// send all the headers received from the client
			md := metadata.Pairs(
				"Authorization", header,
				"x-client-id", clientId,
				"x-client-secret", clientSecret,
				"x-realms", realm,
				"x-user-id", headers.X_User_ID,
				"x-user-email", headers.X_User_Email,
				"x-user-fullname", headers.X_User_Fullname,
				"x-company-id", headers.X_Company_ID,
				"x-company-name", headers.X_Company_Name,
			)

			return md
		}),
	)
}

func generateHeaders(h *entities.Headers, request *http.Request) {
	authorization := request.Header.Get("Authorization")
	sub, ok := utils.GetSub(authorization)

	if !ok {
		return
	}

	// data, _, err := cache.Get(sub)
	_, _, err := cache.Get(sub)
	if err != nil {
		return
	}
	return
	// if err := json.Unmarshal(data.([]byte), &rbacEnt); err != nil {
	// 	return
	// }

	// h.Generate(
	// rbacEnt.User.User.ID,
	// rbacEnt.User.Account.Email,
	// fmt.Sprintf("%s %s", rbacEnt.User.User.FirstName, rbacEnt.User.User.LastName),
	// rbacEnt.User.Company.ID,
	// rbacEnt.User.Company.Name,
	// )

}
