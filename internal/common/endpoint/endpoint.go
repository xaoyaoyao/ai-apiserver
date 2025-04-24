/**
 * Package repo
 * @file      : endpoint.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 13:08
 **/

package endpoint

const (
	ROOT_PATH       = "/api"
	API_STSTEM_PATH = "/api/system"
	API_AI_PATH     = "/api/ai"

	HEALTH_PATH      = "/health"
	VOLC_ENGINE_PATH = "/v1/volcengine"
	MEITU_PATH       = "/v1/meitu"
	USERS_POSTS_PATH = "/v1/users/{id}/posts/{page}"
)

const (
	METHOD_POST   = "POST"
	METHOD_GET    = "GET"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)
