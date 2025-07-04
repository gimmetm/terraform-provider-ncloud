package server

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	. "github.com/terraform-providers/terraform-provider-ncloud/internal/common"
	"github.com/terraform-providers/terraform-provider-ncloud/internal/conn"
)

func DataSourceNcloudAccessControlGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNcloudAccessControlGroupsRead,

		Schema: map[string]*schema.Schema{
			"configuration_no_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Description: "List of ACG configuration numbers you want to get",
			},
			"is_default_group": {
				Type:          schema.TypeBool,
				Optional:      true,
				Deprecated:    "use 'is_default' instead",
				ConflictsWith: []string{"is_default"},
			},
			"is_default": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"is_default_group"},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the ACG you want to get",
			},
			"access_control_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A List of access control group configuration no",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNcloudAccessControlGroupsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*conn.ProviderConfig)
	var resources []map[string]interface{}
	var err error

	resources, err = getVpcAccessControlGroupList(d, config)
	if err != nil {
		return err
	}

	if f, ok := d.GetOk("filter"); ok {
		resources = ApplyFilters(f.(*schema.Set), resources, DataSourceNcloudAccessControlGroups().Schema)
	}

	if len(resources) < 1 {
		return fmt.Errorf("no results. please change search criteria and try again")
	}

	return accessControlGroupsAttributes(d, resources)
}

func accessControlGroupsAttributes(d *schema.ResourceData, accessControlGroups []map[string]interface{}) error {
	var ids []string

	for _, r := range accessControlGroups {
		ids = append(ids, r["id"].(string))
	}

	d.SetId(DataResourceIdHash(ids))
	d.Set("access_control_groups", ids)

	// create a json file in current directory and write d source to it.
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		return WriteToFile(output.(string), d.Get("access_control_groups"))
	}

	return nil
}
