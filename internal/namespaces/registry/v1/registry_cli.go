// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package registry

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		registryRoot(),
		registryNamespace(),
		registryImage(),
		registryTag(),
		registryNamespaceList(),
		registryNamespaceGet(),
		registryNamespaceCreate(),
		registryNamespaceUpdate(),
		registryNamespaceDelete(),
		registryImageList(),
		registryImageGet(),
		registryImageUpdate(),
		registryImageDelete(),
		registryTagList(),
		registryTagGet(),
		registryTagDelete(),
	)
}
func registryRoot() *core.Command {
	return &core.Command{
		Short:     `Docker registry API`,
		Long:      ``,
		Namespace: "registry",
	}
}

func registryNamespace() *core.Command {
	return &core.Command{
		Short: `A namespace is for images what a folder is for files`,
		Long: `
To use our services, the first step is to create a namespace.

A namespace is for images what a folder is for files. Every push or pull must mention the namespace :
` + "`" + `` + "`" + `` + "`" + `docker pull rg.nl-ams.scw.cloud/<namespace_name>/<image_name>:<tag_name>` + "`" + `` + "`" + `` + "`" + `

Note that a namespace name is unique on a region. Thus, if another client already has created "test", you can't have it as a namespace

A namespace can be either public or private (default), which determines who can pull images.
`,
		Namespace: "registry",
		Resource:  "namespace",
	}
}

func registryImage() *core.Command {
	return &core.Command{
		Short: `An image represents a container image`,
		Long: `An image represents a docker image.

The visibility of an image can be public (everyone can pull it), private (only your organization can pull it) or inherit from the namespace visibility (default)
It can be changed with an update on the image via the registry API.
`,
		Namespace: "registry",
		Resource:  "image",
	}
}

func registryTag() *core.Command {
	return &core.Command{
		Short: `A tag represents a container tag of an image`,
		Long: `A tag represents a docker tag of an image.
`,
		Namespace: "registry",
		Resource:  "tag",
	}
}

func registryNamespaceList() *core.Command {
	return &core.Command{
		Short:     `List all your namespaces`,
		Long:      `List all your namespaces.`,
		Namespace: "registry",
		Resource:  "namespace",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(registry.ListNamespacesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Field by which to order the display of Images`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "description_asc", "description_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Filter by the namespace name (exact match)`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter by the namespace owner`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.ListNamespacesRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			resp, err := api.ListNamespaces(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Namespaces, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "id",
			},
			{
				FieldName: "name",
			},
			{
				FieldName: "region",
			},
			{
				FieldName: "endpoint",
			},
			{
				FieldName: "is_public",
			},
			{
				FieldName: "size",
			},
			{
				FieldName: "image_count",
			},
			{
				FieldName: "organization_id",
			},
			{
				FieldName: "status",
			},
			{
				FieldName: "status_message",
			},
			{
				FieldName: "created_at",
			},
			{
				FieldName: "updated_at",
			},
			{
				FieldName: "description",
			},
		}},
	}
}

func registryNamespaceGet() *core.Command {
	return &core.Command{
		Short:     `Get a namespace`,
		Long:      `Get the namespace associated with the given id.`,
		Namespace: "registry",
		Resource:  "namespace",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(registry.GetNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `The unique ID of the Namespace`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.GetNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.GetNamespace(request)

		},
	}
}

func registryNamespaceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new namespace`,
		Long:      `Create a new namespace.`,
		Namespace: "registry",
		Resource:  "namespace",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(registry.CreateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Define a namespace name`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Define a description`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "is-public",
				Short:      `Define the default visibility policy`,
				Required:   false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.CreateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.CreateNamespace(request)

		},
	}
}

func registryNamespaceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing namespace`,
		Long:      `Update the namespace associated with the given id.`,
		Namespace: "registry",
		Resource:  "namespace",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(registry.UpdateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `Namespace ID to update`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "description",
				Short:      `Define a description`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "is-public",
				Short:      `Define the default visibility policy`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.UpdateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.UpdateNamespace(request)

		},
	}
}

func registryNamespaceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing namespace`,
		Long:      `Delete the namespace associated with the given id.`,
		Namespace: "registry",
		Resource:  "namespace",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(registry.DeleteNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `The unique ID of the Namespace`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.DeleteNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.DeleteNamespace(request)

		},
	}
}

func registryImageList() *core.Command {
	return &core.Command{
		Short:     `List all your images`,
		Long:      `List all your images.`,
		Namespace: "registry",
		Resource:  "image",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(registry.ListImagesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Field by which to order the display of Images`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "namespace-id",
				Short:      `Filter by the Namespace ID`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter by the Image name (exact match)`,
				Required:   false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter by Organization ID`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.ListImagesRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			resp, err := api.ListImages(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Images, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "id",
			},
			{
				FieldName: "name",
			},
			{
				FieldName: "size",
			},
			{
				FieldName: "visibility",
			},
			{
				FieldName: "namespace_id",
			},
			{
				FieldName: "status",
			},
			{
				FieldName: "status_message",
			},
			{
				FieldName: "created_at",
			},
			{
				FieldName: "updated_at",
			},
			{
				FieldName: "tags",
			},
		}},
	}
}

func registryImageGet() *core.Command {
	return &core.Command{
		Short:     `Get a image`,
		Long:      `Get the image associated with the given id.`,
		Namespace: "registry",
		Resource:  "image",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(registry.GetImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `The unique ID of the Image`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.GetImageRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.GetImage(request)

		},
	}
}

func registryImageUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing image`,
		Long:      `Update the image associated with the given id.`,
		Namespace: "registry",
		Resource:  "image",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(registry.UpdateImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `Image ID to update`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "visibility",
				Short:      `A ` + "`" + `public` + "`" + ` image is pullable from internet without authentication, opposed to a ` + "`" + `private` + "`" + ` image. ` + "`" + `inherit` + "`" + ` will use the namespace ` + "`" + `is_public` + "`" + ` parameter`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"visibility_unknown", "inherit", "public", "private"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.UpdateImageRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.UpdateImage(request)

		},
	}
}

func registryImageDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an image`,
		Long:      `Delete the image associated with the given id.`,
		Namespace: "registry",
		Resource:  "image",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(registry.DeleteImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `The unique ID of the Image`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.DeleteImageRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.DeleteImage(request)

		},
	}
}

func registryTagList() *core.Command {
	return &core.Command{
		Short:     `List all your tags`,
		Long:      `List all your tags.`,
		Namespace: "registry",
		Resource:  "tag",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(registry.ListTagsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Field by which to order the display of Images`,
				Required:   false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "image-id",
				Short:      `The unique ID of the image`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Filter by the tag name (exact match)`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.ListTagsRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			resp, err := api.ListTags(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Tags, nil

		},
	}
}

func registryTagGet() *core.Command {
	return &core.Command{
		Short:     `Get a tag`,
		Long:      `Get the tag associated with the given id.`,
		Namespace: "registry",
		Resource:  "tag",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(registry.GetTagRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "tag-id",
				Short:      `The unique ID of the Tag`,
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.GetTagRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.GetTag(request)

		},
	}
}

func registryTagDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a tag`,
		Long:      `Delete the tag associated with the given id.`,
		Namespace: "registry",
		Resource:  "tag",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(registry.DeleteTagRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "tag-id",
				Short:      `The unique ID of the tag`,
				Required:   true,
				Positional: true,
			},
			{
				Name:       "force",
				Short:      `If two tags share the same digest the deletion will fail unless this parameter is set to true`,
				Required:   false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*registry.DeleteTagRequest)

			client := core.ExtractClient(ctx)
			api := registry.NewAPI(client)
			return api.DeleteTag(request)

		},
	}
}