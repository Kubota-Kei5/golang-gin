package main

import (
	"encoding/json"
	"net/http"

	// 必要なパッケージはgo getでインストールして下さい。
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"

	oapi "golang-gin/api"
)

type Album struct {
	// `json:"id,omitempty"`などがコメントアウトされていることに注目して下さい。
	// この変換は oapi-codgenが生成したstructを使ってMarshalJSONとUnmarshalJSON行っています。
	ID    int    //`json:"id,omitempty"`
	Title string //`json:"title"`
}

// デモ用にグローバルでAlbumを作成しています。実際にはデーターベースに格納されます。
var albums = []Album{
	{ID: 1, Title: "Test"},
}

type AlbumHandler struct{}

/*
GetAlbum関数は、api/api.gen.goの中のインタフェースで定義されているので、GetAlbumという名前にする必要があります。
openapi.yamlの中のoperationIdで定義されている値が入っています。
GoのスタイルガイドではAlbumHandlerというstructにAlbumが入っているので、funcではGetAlbumとAlbumを重複して書かないことが推奨されています。
しかし、oepnapi.yamlの中で重複しない名前を作成し、そこから自動生成しているコードになりますので以下の名前でも問題ありません。

この関数では、アルバムのIDを使用してアルバムを検索し、JSONレスポンスを返します
*/
func (a AlbumHandler) GetAlbum(c *gin.Context, ID int) {
	for _, album := range albums {
		if album.ID == ID {
			// OpenAPIに準拠したResponseに変更してresponseを返します。
			response := oapi.AlbumGetResponse{
				Id:    &album.ID,
				Title: album.Title,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	c.JSON(http.StatusNotFound, nil)
}

/*
CreateAlbum関数は、JSONリクエストボディからアルバムを作成し、JSONレスポンスを返します。
*/
func (a AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody oapi.AlbumCreateRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	album := Album{
		Title: requestBody.Title,
	}

	album.ID = len(albums) + 1
	albums = append(albums, album)
	c.JSON(http.StatusCreated, nil)
}

func main() {
	router := gin.Default()

	/*
		以下のコードは、GinサーバーでSwagger UIを表示するためのものであり、oapi.GetSwagger()関数を使用してSwagger YAMLファイルを取得し、
		json.Marshal()関数を使用してJSON形式に変換します。次に、swag.Spec構造体を使用してSwagger仕様を定義し、swag.Register()関数を
		使用してSwagger仕様を登録します。最後に、ginSwagger.WrapHandler()関数を使用してSwagger UIを表示するためのGinミドルウェアを
		設定します。

		SwaggerURL http://0.0.0.0:8080/swagger/index.html
	*/
	swagger, _ := oapi.GetSwagger()
	swaggerJson, _ := json.Marshal(swagger)
	var SwaggerInfo = &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(swaggerJson),
	}
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		/*
			このミドルウェアは、swagger仕様を使用してリクエストを検証するために使用されます。swagger仕様は、APIの仕様を定義するための
			オープンスタンダードであり、APIのリクエストとレスポンスの形式を定義します。 middleware.OapiRequestValidator(swagger)は、
			swagger仕様を引数として受け取り、ginフレームワークのミドルウェアを返します。このミドルウェアは、リクエストを検証し、
			swagger仕様に従ってリクエストが有効であることを確認します。router.Use()関数を使用して、このミドルウェアをGinサーバーに
			追加することができます。
		*/
		api.Use(middleware.OapiRequestValidator(swagger))

		v1 := api.Group("/v1")
		{
			/*
				`oapi.RegisterHandlers(v1, albumHandler)`は、oapi-codegenによって生成されたGoのAPIコードで使用される
				関数の1つです。この関数は、APIのエンドポイントとハンドラを登録します。 `v1`パラメータは、Ginルーターのグループを表します。
				`albumHandler`パラメータは、APIのハンドラを表します。この関数は、APIのエンドポイントとハンドラをマッピングし、
				APIの実装を完了します。
			*/
			albumHandler := AlbumHandler{}
			oapi.RegisterHandlers(v1, albumHandler)

		}
	}

	router.Run(":8080")
}
