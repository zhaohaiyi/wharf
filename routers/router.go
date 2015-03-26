package routers

import (
	"github.com/astaxie/beego"

	"github.com/dockercn/wharf/controllers"
	"github.com/dockercn/wharf/filters"
)

func init() {
	//Web Interface
	beego.Router("/", &controllers.WebController{}, "get:GetIndex")
	beego.Router("/auth", &controllers.WebController{}, "get:GetAuth")
	beego.Router("/setting", &controllers.WebController{}, "get:GetSetting")
	beego.Router("/dashboard", &controllers.WebController{}, "get:GetDashboard")
	beego.Router("/signout", &controllers.WebController{}, "get:GetSignout")
	beego.Router("/admin/auth", &controllers.WebController{}, "get:GetAdminAuth")
	beego.Router("/admin", &controllers.WebController{}, "get:GetAdmin")

	beego.Router("/r/:namespace/:repository", &controllers.WebController{}, "get:GetRepository")
	beego.Router("/o/:org", &controllers.WebController{}, "get:GetOrganization")
	beego.Router("/u/:username", &controllers.WebController{}, "get:GetUser")

	//Static File Route
	beego.Router("/pubkeys.gpg", &controllers.FileController{}, "get:GetGPG")

	//Web API
	web := beego.NewNamespace("/w1",
		beego.NSRouter("/signin", &controllers.UserWebAPIV1Controller{}, "post:Signin"),
		beego.NSRouter("/signup", &controllers.UserWebAPIV1Controller{}, "post:Signup"),

		//user routers
		beego.NSRouter("/users", &controllers.UserWebAPIV1Controller{}, "get:GetUsers"),
		beego.NSRouter("/users/:username", &controllers.UserWebAPIV1Controller{}, "get:GetUser"),
		beego.NSRouter("/namespaces", &controllers.UserWebAPIV1Controller{}, "get:GetNamespaces"),
		beego.NSRouter("/profile", &controllers.UserWebAPIV1Controller{}, "get:GetProfile"),
		beego.NSRouter("/gravatar", &controllers.UserWebAPIV1Controller{}, "post:PostGravatar"),
		beego.NSRouter("/password", &controllers.UserWebAPIV1Controller{}, "put:PutPassword"),
		beego.NSRouter("/profile", &controllers.UserWebAPIV1Controller{}, "put:PutProfile"),
		beego.NSRouter("/profile/:username", &controllers.UserWebAPIV1Controller{}, "get:GetUserProfile"),

		//repository routers
		beego.NSRouter("/repositories", &controllers.RepoWebAPIV1Controller{}, "get:GetRepositories"),
		beego.NSRouter("/repository", &controllers.RepoWebAPIV1Controller{}, "post:PostRepository"),
		beego.NSRouter("/repository/:namespace/:repository", &controllers.RepoWebAPIV1Controller{}, "put:PutRepository"),
		beego.NSRouter("/repository/:namespace/:repository", &controllers.RepoWebAPIV1Controller{}, "get:GetRepository"),
		beego.NSRouter("/repository/:namespace/:repository/collaborators", &controllers.RepoWebAPIV1Controller{}, "get:GetCollaborators"),
		beego.NSRouter("/repository/:namespace/:repository/collaborators/:collaborator", &controllers.RepoWebAPIV1Controller{}, "post:PostCollaborator"),
		beego.NSRouter("/repository/:namespace/:repository/collaborators/:collaborator/:action", &controllers.RepoWebAPIV1Controller{}, "put:PutCollaborator"),

		//organization routers
		beego.NSRouter("/organizations", &controllers.OrganizationWebV1Controller{}, "get:GetOrganizations"),
		beego.NSRouter("/organization", &controllers.OrganizationWebV1Controller{}, "post:PostOrganization"),
		beego.NSRouter("/organization/:org", &controllers.OrganizationWebV1Controller{}, "put:PutOrganization"),
		beego.NSRouter("/organization/:org", &controllers.OrganizationWebV1Controller{}, "get:GetOrganization"),
		beego.NSRouter("/organization/:org/repo", &controllers.OrganizationWebV1Controller{}, "get:GetOrganizationRepo"),
		beego.NSRouter("/organization/:org/teams", &controllers.TeamWebV1Controller{}, "get:GetTeams"),
		beego.NSRouter("/organization/:org/team", &controllers.TeamWebV1Controller{}, "post:PostTeam"),
		beego.NSRouter("/organization/:org/team/:team", &controllers.TeamWebV1Controller{}, "put:PutTeam"),
		beego.NSRouter("/organization/:org/team/:team", &controllers.TeamWebV1Controller{}, "get:GetTeam"),
		beego.NSRouter("/organization/:org/team/:team/:username/:action", &controllers.TeamWebV1Controller{}, "put:PutTeamAddMember"),
	)

	//Docker Registry API V1 remain
	beego.Router("/_ping", &controllers.PingAPIV1Controller{}, "get:GetPing")

	//Docker Registry API V1
	apiv1 := beego.NewNamespace("/v1",
		beego.NSRouter("/_ping", &controllers.PingAPIV1Controller{}, "get:GetPing"),
		beego.NSRouter("/users", &controllers.UserAPIV1Controller{}, "get:GetUsers"),
		beego.NSRouter("/users", &controllers.UserAPIV1Controller{}, "post:PostUsers"),

		beego.NSNamespace("/repositories",
			beego.NSRouter("/:namespace/:repo_name/tags/:tag", &controllers.RepoAPIV1Controller{}, "put:PutTag"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepoAPIV1Controller{}, "put:PutRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepoAPIV1Controller{}, "get:GetRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/tags", &controllers.RepoAPIV1Controller{}, "get:GetRepositoryTags"),
			beego.NSRouter("/:namespace/:repo_name", &controllers.RepoAPIV1Controller{}, "put:PutRepository"),
		),

		beego.NSNamespace("/images",
			beego.NSRouter("/:image_id/ancestry", &controllers.ImageAPIV1Controller{}, "get:GetImageAncestry"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIV1Controller{}, "get:GetImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIV1Controller{}, "get:GetImageLayer"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIV1Controller{}, "put:PutImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIV1Controller{}, "put:PutImageLayer"),
			beego.NSRouter("/:image_id/checksum", &controllers.ImageAPIV1Controller{}, "put:PutChecksum"),
		),
	)

	//Docker Registry API V2
	apiv2 := beego.NewNamespace("/v2",
		beego.NSRouter("/", &controllers.PingAPIV2Controller{}, "get:GetPing"),
		//Push
		beego.NSRouter("/:namespace/:repo_name/blobs/:digest", &controllers.BlobAPIV2Controller{}, "head:HeadDigest"),
		beego.NSRouter("/:namespace/:repo_name/blobs/uploads", &controllers.BlobAPIV2Controller{}, "post:PostBlobs"),
		beego.NSRouter("/:namespace/:repo_name/blobs/uploads/:uuid", &controllers.BlobAPIV2Controller{}, "put:PutBlobs"),
		beego.NSRouter("/:namespace/:repo_name/manifests/:tag", &controllers.ManifestsAPIV2Controller{}, "put:PutManifests"),
		//Pull
		beego.NSRouter("/:namespace/:repo_name/tags/list", &controllers.ManifestsAPIV2Controller{}, "get:GetTags"),
		beego.NSRouter("/:namespace/:repo_name/manifests/:tag", &controllers.ManifestsAPIV2Controller{}, "get:GetManifests"),
		beego.NSRouter("/:namespace/:repo_name/blobs/:digest", &controllers.BlobAPIV2Controller{}, "get:GetBlobs"),
	)

	//Dockerfile Build API V1
	buildv1 := beego.NewNamespace("/b1",
		beego.NSRouter("/build", &controllers.BuilderAPIV1Controller{}, "post:PostBuild"),
		beego.NSRouter("/status", &controllers.BuilderAPIV1Controller{}, "get:GetStatus"),
	)

	//Auth Fiters
	beego.InsertFilter("/w1/*", beego.BeforeRouter, filters.FilterDebug)
	beego.InsertFilter("/v1/*", beego.BeforeRouter, filters.FilterDebug)
	beego.InsertFilter("/v2/*", beego.BeforeRouter, filters.FilterDebug)
	beego.InsertFilter("/b1/*", beego.BeforeRouter, filters.FilterDebug)

	beego.InsertFilter("/v1/repositories/*", beego.BeforeRouter, filters.FilterAuth)
	beego.InsertFilter("/v1/images/*", beego.BeforeRouter, filters.FilterAuth)
	beego.InsertFilter("/v2/*", beego.BeforeRouter, filters.FilterAuth)

	beego.AddNamespace(web)
	beego.AddNamespace(apiv1)
	if beego.AppConfig.String("docker::API") == "v2" {
		beego.AddNamespace(apiv2)
	}

	beego.AddNamespace(buildv1)
}
