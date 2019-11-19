package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10.1"
)

func resourceNetworkPerformanceBinding() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetworkPerformanceBindingCreate,
        Read:   resourceNetworkPerformanceBindingRead,
        Update: resourceNetworkPerformanceBindingUpdate,
        Delete: resourceNetworkPerformanceBindingDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_network_measurement_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNetworkPerformanceBindingCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetworkPerformanceBinding object
    o := &vspk.NetworkPerformanceBinding{
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_network_measurement_id"); ok {
        o.AssociatedNetworkMeasurementID = attr.(string)
    }
    parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
    err := parent.CreateNetworkPerformanceBinding(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNetworkPerformanceBindingRead(d, m)
}

func resourceNetworkPerformanceBindingRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceBinding{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("read_only", o.ReadOnly)
    d.Set("priority", o.Priority)
    d.Set("associated_network_measurement_id", o.AssociatedNetworkMeasurementID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetworkPerformanceBindingUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceBinding{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_network_measurement_id"); ok {
        o.AssociatedNetworkMeasurementID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNetworkPerformanceBindingDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkPerformanceBinding{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}