package ncloud

import (
	"fmt"

	"github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceNcloudNasVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNcloudNasVolumeRead,

		Schema: map[string]*schema.Schema{
			"volume_allotment_protocol_type_code": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIncludeValues([]string{"NFS", "CIFS"}),
			},
			"is_event_configuration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBoolValue,
			},
			"is_snapshot_configuration": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateBoolValue,
			},
			"nas_volume_instance_no_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"region_code": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Region code. Get available values using the `data ncloud_regions`.",
				ConflictsWith: []string{"region_no"},
			},
			"region_no": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Region number. Get available values using the `data ncloud_regions`.",
				ConflictsWith: []string{"region_code"},
			},
			"zone_code": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Zone code",
				ConflictsWith: []string{"zone_no"},
			},
			"zone_no": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Zone number",
				ConflictsWith: []string{"zone_code"},
			},

			"nas_volume_instance_no": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nas_volume_instance_status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"create_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_allotment_protocol_type": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     commonCodeSchemaResource,
			},
			"volume_total_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_use_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_use_ratio": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"snapshot_volume_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"snapshot_volume_use_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"snapshot_volume_use_ratio": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"nas_volume_instance_custom_ip_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nas_volume_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     zoneSchemaResource,
			},
			"region": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     regionSchemaResource,
			},
		},
	}
}

func dataSourceNcloudNasVolumeRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*NcloudSdk).conn

	reqParams := &sdk.RequestGetNasVolumeInstanceList{
		VolumeAllotmentProtocolTypeCode: d.Get("volume_allotment_protocol_type_code").(string),
		IsEventConfiguration:            d.Get("is_event_configuration").(string),
		IsSnapshotConfiguration:         d.Get("is_snapshot_configuration").(string),
		NasVolumeInstanceNoList:         StringList(d.Get("nas_volume_instance_no_list").([]interface{})),
		RegionNo:                        parseRegionNoParameter(conn, d),
		ZoneNo:                          parseZoneNoParameter(conn, d),
	}
	resp, err := conn.GetNasVolumeInstanceList(reqParams)
	if err != nil {
		logErrorResponse("GetNasVolumeInstanceList", err, reqParams)
		return err
	}
	logCommonResponse("GetNasVolumeInstanceList", reqParams, resp.CommonResponse)

	var nasVolumeInstance sdk.NasVolumeInstance
	nasVolumeInstances := resp.NasVolumeInstanceList
	if len(nasVolumeInstances) < 1 {
		return fmt.Errorf("no results. please change search criteria and try again")
	}
	nasVolumeInstance = nasVolumeInstances[0]

	return nasVolumeInstanceAttributes(d, nasVolumeInstance)
}

func nasVolumeInstanceAttributes(d *schema.ResourceData, nasVolume sdk.NasVolumeInstance) error {
	d.Set("nas_volume_instance_no", nasVolume.NasVolumeInstanceNo)
	d.Set("nas_volume_instance_status", setCommonCode(nasVolume.NasVolumeInstanceStatus))
	d.Set("create_date", nasVolume.CreateDate)
	d.Set("nas_volume_description", nasVolume.NasVolumeInstanceDescription)
	d.Set("volume_allotment_protocol_type", setCommonCode(nasVolume.VolumeAllotmentProtocolType))
	d.Set("volume_name", nasVolume.VolumeName)
	d.Set("volume_total_size", nasVolume.VolumeTotalSize)
	d.Set("volume_size", nasVolume.VolumeSize)
	d.Set("volume_use_size", nasVolume.VolumeUseSize)
	d.Set("volume_use_ratio", nasVolume.VolumeUseRatio)
	d.Set("snapshot_volume_size", nasVolume.SnapshotVolumeSize)
	d.Set("snapshot_volume_use_size", nasVolume.SnapshotVolumeUseSize)
	d.Set("snapshot_volume_use_ratio", nasVolume.SnapshotVolumeUseRatio)
	d.Set("is_snapshot_configuration", nasVolume.IsSnapshotConfiguration)
	d.Set("is_event_configuration", nasVolume.IsEventConfiguration)
	d.Set("nas_volume_instance_custom_ip_list", nasVolume.NasVolumeInstanceCustomIPList)
	d.Set("zone", setZone(nasVolume.Zone))
	d.Set("region", setRegion(nasVolume.Region))

	d.SetId(nasVolume.NasVolumeInstanceNo)

	return nil
}