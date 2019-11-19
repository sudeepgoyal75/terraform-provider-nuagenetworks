package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.5"
)

func resourceNextHopAddress() *schema.Resource {
    return &schema.Resource{
        Create: resourceNextHopAddressCreate,
        Read:   resourceNextHopAddressRead,
        Update: resourceNextHopAddressUpdate,
        Delete: resourceNextHopAddressDelete,
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
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNextHopAddressCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NextHopAddress object
    o := &vspk.NextHopAddress{
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    err := parent.CreateNextHopAddress(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNextHopAddressRead(d, m)
}

func resourceNextHopAddressRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHopAddress{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("address", o.Address)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNextHopAddressUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHopAddress{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNextHopAddressDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHopAddress{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}