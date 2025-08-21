package restaurant

import "github.com/gin-gonic/gin"

type RestaurantHendler struct {
	restaurantService service
}

func NewRestaurantHendler(restaurantService service) *RestaurantHendler {
	return &RestaurantHendler{restaurantService: restaurantService}
}

func (h *RestaurantHendler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	restaurant := api.Group("/restaurant")
	{
		restaurant.GET("/:id")
	}
}
func (h *RestaurantHendler) CreateRestaurant(c *gin.Context) {

}

//todo  доделать хендлер так как сейчас нет авторизации
//CreateRestaurant - POST /restaurants
//GetRestaurant - GET /restaurants/:id
//UpdateRestaurant - PUT/PATCH /restaurants/:id
//DeleteRestaurant - DELETE /restaurants/:id
//ActivateRestaurant - PATCH /restaurants/:id/activate
//DeactivateRestaurant - PATCH /restaurants/:id/deactivate
//SetFeaturedStatus - PATCH /restaurants/:id/featured (только админ)
//GetMyRestaurants - GET /restaurants/my
//GetAllRestaurants - GET /restaurants
//SearchRestaurants - GET /restaurants/search
