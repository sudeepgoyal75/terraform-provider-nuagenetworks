package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceDeploymentFailure() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceDeploymentFailureRead,
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
            "last_failure_reason": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_known_error": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "affected_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "affected_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "error_condition": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "number_of_occurences": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "event_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redundancy_group"},
            },
        },
    }
}


func dataSourceDeploymentFailureRead(d *schema.ResourceData, m interface{}) error {
    filteredDeploymentFailures := vspk.DeploymentFailuresList{}
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
    if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredDeploymentFailures, err = parent.DeploymentFailures(fetchFilter)
        if err != nil {
            return err
        }
    }

    DeploymentFailure := &vspk.DeploymentFailure{}

    if len(filteredDeploymentFailures) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredDeploymentFailures) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    DeploymentFailure = filteredDeploymentFailures[0]

    d.Set("last_failure_reason", DeploymentFailure.LastFailureReason)
    d.Set("last_known_error", DeploymentFailure.LastKnownError)
    d.Set("affected_entity_id", DeploymentFailure.AffectedEntityID)
    d.Set("affected_entity_type", DeploymentFailure.AffectedEntityType)
    d.Set("error_condition", DeploymentFailure.ErrorCondition)
    d.Set("number_of_occurences", DeploymentFailure.NumberOfOccurences)
    d.Set("event_type", DeploymentFailure.EventType)
    
    d.Set("id", DeploymentFailure.Identifier())
    d.Set("parent_id", DeploymentFailure.ParentID)
    d.Set("parent_type", DeploymentFailure.ParentType)
    d.Set("owner", DeploymentFailure.Owner)

    d.SetId(DeploymentFailure.Identifier())
    
    return nil
}