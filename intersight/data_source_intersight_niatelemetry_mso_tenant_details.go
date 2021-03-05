package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryMsoTenantDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryMsoTenantDetailsRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_schemas_assigned_per_tenant_in_mso": {
				Description: "Number of schemas assigned to each tenant in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sites_each_tenant_is_deployed_to_in_mso": {
				Description: "Number of sites each tenant is deployed to.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"tenant_id": {
				Description: "ID of tenant in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tenant_name": {
				Description: "Name of the tenant in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"number_of_schemas_assigned_per_tenant_in_mso": {
						Description: "Number of schemas assigned to each tenant in Multi-Site Orchestrator.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"sites_each_tenant_is_deployed_to_in_mso": {
						Description: "Number of sites each tenant is deployed to.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"tenant_id": {
						Description: "ID of tenant in Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tenant_name": {
						Description: "Name of the tenant in Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryMsoTenantDetailsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryMsoTenantDetails{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("number_of_schemas_assigned_per_tenant_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfSchemasAssignedPerTenantInMso(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("sites_each_tenant_is_deployed_to_in_mso"); ok {
		x := int64(v.(int))
		o.SetSitesEachTenantIsDeployedToInMso(x)
	}
	if v, ok := d.GetOk("tenant_id"); ok {
		x := (v.(string))
		o.SetTenantId(x)
	}
	if v, ok := d.GetOk("tenant_name"); ok {
		x := (v.(string))
		o.SetTenantName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryMsoTenantDetails object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryMsoTenantDetailsList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryMsoTenantDetails: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryMsoTenantDetailsList.GetCount()
	var i int32
	var niatelemetryMsoTenantDetailsResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryMsoTenantDetailsList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryMsoTenantDetails: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryMsoTenantDetailsList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryMsoTenantDetails data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["number_of_schemas_assigned_per_tenant_in_mso"] = (s.GetNumberOfSchemasAssignedPerTenantInMso())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["sites_each_tenant_is_deployed_to_in_mso"] = (s.GetSitesEachTenantIsDeployedToInMso())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["tenant_id"] = (s.GetTenantId())
				temp["tenant_name"] = (s.GetTenantName())
				niatelemetryMsoTenantDetailsResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryMsoTenantDetailsResults))
	if err := d.Set("results", niatelemetryMsoTenantDetailsResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryMsoTenantDetailsResults[0]["moid"].(string))
	return de
}