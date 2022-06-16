package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/search"
	"github.com/photoprism/photoprism/pkg/txt"
)

// SearchSubjects finds and returns subjects as JSON.
//
// GET /api/v1/subjects
func SearchSubjects(router *gin.RouterGroup) {
	router.GET("/subjects", func(c *gin.Context) {
		s := Auth(SessionID(c), acl.ResourceSubjects, acl.ActionSearch)

		if s.Invalid() {
			AbortUnauthorized(c)
			return
		}

		var f form.SearchSubjects

		err := c.MustBindWith(&f, binding.Form)

		if err != nil {
			AbortBadRequest(c)
			return
		}

		result, err := search.Subjects(f)

		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": txt.UpperFirst(err.Error())})
			return
		}

		AddCountHeader(c, len(result))
		AddLimitHeader(c, f.Count)
		AddOffsetHeader(c, f.Offset)
		AddTokenHeaders(c)

		c.JSON(http.StatusOK, result)
	})
}
