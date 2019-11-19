package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceMetadataTag() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceMetadataTagRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_external_service_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "auto_created": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_global_metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise", "parent_metadata", "parent_external_service"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_global_metadata", "parent_metadata", "parent_external_service"},
            },
            "parent_metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_global_metadata", "parent_enterprise", "parent_external_service"},
            },
            "parent_external_service": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_global_metadata", "parent_enterprise", "parent_metadata"},
            },
        },
    }
}


func dataSourceMetadataTagRead(d *schema.ResourceData, m interface{}) error {
    filteredMetadataTags := vspk.MetadataTagsList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_global_metadata"); ok {
        parent := &vspk.GlobalMetadata{ID: attr.(string)}
        filteredMetadataTags, err = parent.MetadataTags(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredMetadataTags, err = parent.MetadataTags(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_metadata"); ok {
        parent := &vspk.Metadata{ID: attr.(string)}
        filteredMetadataTags, err = parent.MetadataTags(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_external_service"); ok {
        parent := &vspk.ExternalService{ID: attr.(string)}
        filteredMetadataTags, err = parent.MetadataTags(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredMetadataTags, err = parent.MetadataTags(fetchFilter)
        if err != nil {
            return err
        }
    }

    MetadataTag := &vspk.MetadataTag{}

    if len(filteredMetadataTags) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredMetadataTags) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    MetadataTag = filteredMetadataTags[0]

    d.Set("name", MetadataTag.Name)
    d.Set("last_updated_by", MetadataTag.LastUpdatedBy)
    d.Set("description", MetadataTag.Description)
    d.Set("entity_scope", MetadataTag.EntityScope)
    d.Set("associated_external_service_id", MetadataTag.AssociatedExternalServiceID)
    d.Set("auto_created", MetadataTag.AutoCreated)
    d.Set("external_id", MetadataTag.ExternalID)
    
    d.Set("id", MetadataTag.Identifier())
    d.Set("parent_id", MetadataTag.ParentID)
    d.Set("parent_type", MetadataTag.ParentType)
    d.Set("owner", MetadataTag.Owner)

    d.SetId(MetadataTag.Identifier())
    
    return nil
}