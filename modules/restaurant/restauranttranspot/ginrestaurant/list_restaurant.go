package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-services-g04/common"
	"simple-services-g04/component"
	"simple-services-g04/modules/restaurant/restaurantbiz"
	"simple-services-g04/modules/restaurant/restaurantmodel"
	"simple-services-g04/modules/restaurant/restaurantstorage"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		paging.Fullfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccnessResponse(result, paging, filter))
	}
}
