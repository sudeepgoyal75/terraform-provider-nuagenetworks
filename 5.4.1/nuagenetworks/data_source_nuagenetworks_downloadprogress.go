package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceDownloadProgress() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceDownloadProgressRead,
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
            "percentage": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "time_left": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "time_spent": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "image_file_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "image_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "start_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "current_speed": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "average_speed": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
        },
    }
}


func dataSourceDownloadProgressRead(d *schema.ResourceData, m interface{}) error {
    filteredDownloadProgress := vspk.DownloadProgressList{}
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

    DownloadProgress := &vspk.DownloadProgress{}

    if len(filteredDownloadProgress) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredDownloadProgress) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    DownloadProgress = filteredDownloadProgress[0]

    d.Set("percentage", DownloadProgress.Percentage)
    d.Set("time_left", DownloadProgress.TimeLeft)
    d.Set("time_spent", DownloadProgress.TimeSpent)
    d.Set("image_file_name", DownloadProgress.ImageFileName)
    d.Set("image_version", DownloadProgress.ImageVersion)
    d.Set("start_time", DownloadProgress.StartTime)
    d.Set("current_speed", DownloadProgress.CurrentSpeed)
    d.Set("average_speed", DownloadProgress.AverageSpeed)
    
    d.Set("id", DownloadProgress.Identifier())
    d.Set("parent_id", DownloadProgress.ParentID)
    d.Set("parent_type", DownloadProgress.ParentType)
    d.Set("owner", DownloadProgress.Owner)

    d.SetId(DownloadProgress.Identifier())
    
    return nil
}