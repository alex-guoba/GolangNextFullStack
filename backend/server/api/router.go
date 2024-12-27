package api

import (
	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/alex-guoba/gin-clean-template/server/api/public"
	v1 "github.com/alex-guoba/gin-clean-template/server/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, db *gorm.DB) {
	article := v1.NewArticle(db, cfg)
	tag := v1.NewTag(db, cfg)
	apiv1 := r.Group("/api/v1")
	{
		// Create tag
		apiv1.POST("/tags", tag.Create)
		// Delete tag
		apiv1.DELETE("/tags/:id", tag.Delete)
		// Update tab
		apiv1.PUT("/tags/:id", tag.Update)
		// Get tag
		apiv1.GET("/tags", tag.List)

		// Create artile with tags
		apiv1.POST("/articles", article.Create)
		// Get aritle detail by id
		apiv1.GET("/articles/:id", article.Get)
		// Get article list and total count
		apiv1.GET("/articles", article.List)

		// Add other router if necessary
	}

	gPub := r.Group("/public")
	{
		delay := &public.Delay{}
		gPub.GET("/delay/:seconds", delay.DelayHandler)
	}
}
