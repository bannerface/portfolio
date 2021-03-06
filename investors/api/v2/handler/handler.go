package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/client"
	followers "github.com/micro/services/portfolio/followers/proto"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	"github.com/micro/services/portfolio/helpers/photos"
	proto "github.com/micro/services/portfolio/investors-api/proto"
	allocation "github.com/micro/services/portfolio/portfolio-allocation/proto"
	valuation "github.com/micro/services/portfolio/portfolio-value-tracking/proto"
	portfolios "github.com/micro/services/portfolio/portfolios/proto"
	posts "github.com/micro/services/portfolio/posts/proto"
	quotes "github.com/micro/services/portfolio/stock-quote-v2/proto"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	trades "github.com/micro/services/portfolio/trades/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// Handler is an object can process RPC requests
type Handler struct {
	auth       auth.Authenticator
	users      users.UsersService
	posts      posts.PostsService
	photos     photos.Service
	stocks     stocks.StocksService
	quotes     quotes.StockQuoteService
	trades     trades.TradesService
	followers  followers.FollowersService
	valuation  valuation.PortfolioValueTrackingService
	portfolios portfolios.PortfoliosService
	allocation allocation.PortfolioAllocationService
}

// New returns an instance of Handler
func New(auth auth.Authenticator, pics photos.Service, client client.Client) Handler {
	return Handler{
		auth:       auth,
		photos:     pics,
		users:      users.NewUsersService("kytra-v1-users:8080", client),
		posts:      posts.NewPostsService("kytra-v1-posts:8080", client),
		stocks:     stocks.NewStocksService("kytra-v1-stocks:8080", client),
		trades:     trades.NewTradesService("kytra-v1-trades:8080", client),
		quotes:     quotes.NewStockQuoteService("kytra-v2-stock-quote:8080", client),
		followers:  followers.NewFollowersService("kytra-v1-followers:8080", client),
		valuation:  valuation.NewPortfolioValueTrackingService("kytra-v1-portfolio-value-tracking:8080", client),
		portfolios: portfolios.NewPortfoliosService("kytra-v1-portfolios:8080", client),
		allocation: allocation.NewPortfolioAllocationService("kytra-v1-portfolio-allocation:8080", client),
	}
}

// Get retrieves the investor
func (h Handler) Get(ctx context.Context, req *proto.User, rsp *proto.User) error {
	user, err := h.users.Find(ctx, &users.User{Uuid: req.Uuid, Username: req.Username})
	if err != nil {
		return err
	}

	// Check if the current user follows this investor
	following, _ := h.getFollowingStatus(ctx, user.Uuid)

	*rsp = proto.User{
		Uuid:              user.Uuid,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Username:          user.Username,
		Following:         following,
		ProfilePictureUrl: h.photos.GetURL(user.ProfilePictureId),
	}

	// Add the summary
	summaries, err := h.summariesForUsers(ctx, []string{user.Uuid})
	if err == nil {
		rsp.Summary = summaries[user.Uuid]
	} else {
		return err // TODO: REMOVE
	}

	return nil
}

func (h Handler) serializeUsers(ctx context.Context, rsp *proto.ListResponse, uuids []string, allocations ...*allocation.Portfolio) error {
	usersRsp, err := h.users.List(ctx, &users.ListRequest{Uuids: uuids})
	if err != nil {
		return err
	}

	rsp.Users = make([]*proto.User, len(usersRsp.Users))
	followings, err := h.getFollowingStatuses(ctx, uuids)
	if err != nil {
		return err
	}

	summaries, err := h.summariesForUsers(ctx, uuids, allocations...)
	if err != nil {
		return err
	}

	movements, _ := h.priceMovementsForUsers(ctx, uuids)

	for i, user := range usersRsp.Users {
		rsp.Users[i] = &proto.User{
			Uuid:                        user.Uuid,
			FirstName:                   user.FirstName,
			LastName:                    user.LastName,
			Username:                    user.Username,
			Following:                   followings[user.Uuid],
			ProfilePictureUrl:           h.photos.GetURL(user.ProfilePictureId),
			Summary:                     summaries[user.Uuid],
			OneWeekPriceMovementPercent: movements[user.Uuid],
		}
	}

	return nil
}

func (h Handler) priceMovementsForUsers(ctx context.Context, uuids []string) (map[string]float32, error) {
	pRsp, err := h.portfolios.List(ctx, &portfolios.ListRequest{UserUuids: uuids})
	if err != nil {
		return map[string]float32{}, err
	}
	portfolioUUIDs := make([]string, len(uuids))
	userUUIDByPortfolioUUID := make(map[string]string, len(uuids))
	for i, p := range pRsp.GetPortfolios() {
		portfolioUUIDs[i] = p.Uuid
		userUUIDByPortfolioUUID[p.Uuid] = p.UserUuid
	}

	startTime := time.Now().Add(time.Hour * 24 * -7).Unix()
	endTime := time.Now().Unix()
	vRsp, err := h.valuation.ListPriceMovements(ctx, &valuation.ListPriceMovementsRequest{
		PortfolioUuids: portfolioUUIDs,
		StartDate:      startTime,
		EndDate:        endTime,
	})
	if err != nil {
		return map[string]float32{}, err
	}

	data := make(map[string]float32, len(uuids))
	for _, v := range vRsp.GetPriceMovements() {
		userUUID := userUUIDByPortfolioUUID[v.PortfolioUuid]
		data[userUUID] = v.PercentageChange
	}

	return data, nil
}
