module spa-web-app-template/src-backend

go 1.14

replace (
	spa-web-app-template/src-backend/api => ./api
	spa-web-app-template/src-backend/libs => ./libs
)

require (
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.4
	github.com/lib/pq v1.10.3
)
