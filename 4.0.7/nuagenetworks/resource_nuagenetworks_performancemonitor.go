package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.7"
)

func resourcePerformanceMonitor() *schema.Resource {
    return &schema.Resource{
        Create: resourcePerformanceMonitorCreate,
        Read:   resourcePerformanceMonitorRead,
        Update: resourcePerformanceMonitorUpdate,
        Delete: resourcePerformanceMonitorDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "payload_size": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "service_class": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "interval": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "number_of_packets": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourcePerformanceMonitorCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PerformanceMonitor object
    o := &vspk.PerformanceMonitor{
        Name: d.Get("name").(string),
        PayloadSize: d.Get("payload_size").(int),
        Interval: d.Get("interval").(int),
        NumberOfPackets: d.Get("number_of_packets").(int),
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("service_class"); ok {
        o.ServiceClass = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreatePerformanceMonitor(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreatePerformanceMonitor(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourcePerformanceMonitorRead(d, m)
}

func resourcePerformanceMonitorRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PerformanceMonitor{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("payload_size", o.PayloadSize)
    d.Set("read_only", o.ReadOnly)
    d.Set("service_class", o.ServiceClass)
    d.Set("description", o.Description)
    d.Set("interval", o.Interval)
    d.Set("number_of_packets", o.NumberOfPackets)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePerformanceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PerformanceMonitor{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.PayloadSize = d.Get("payload_size").(int)
    o.Interval = d.Get("interval").(int)
    o.NumberOfPackets = d.Get("number_of_packets").(int)
    
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    if attr, ok := d.GetOk("service_class"); ok {
        o.ServiceClass = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePerformanceMonitorDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PerformanceMonitor{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}