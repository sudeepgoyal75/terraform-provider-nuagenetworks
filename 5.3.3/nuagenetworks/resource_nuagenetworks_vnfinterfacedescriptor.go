package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceVNFInterfaceDescriptor() *schema.Resource {
    return &schema.Resource{
        Create: resourceVNFInterfaceDescriptorCreate,
        Read:   resourceVNFInterfaceDescriptorRead,
        Update: resourceVNFInterfaceDescriptorUpdate,
        Delete: resourceVNFInterfaceDescriptorDelete,
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
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "MANAGEMENT",
            },
            "parent_vnf_descriptor": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceVNFInterfaceDescriptorCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VNFInterfaceDescriptor object
    o := &vspk.VNFInterfaceDescriptor{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.VNFDescriptor{ID: d.Get("parent_vnf_descriptor").(string)}
    err := parent.CreateVNFInterfaceDescriptor(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVNFInterfaceDescriptorRead(d, m)
}

func resourceVNFInterfaceDescriptorRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFInterfaceDescriptor{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVNFInterfaceDescriptorUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFInterfaceDescriptor{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVNFInterfaceDescriptorDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFInterfaceDescriptor{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}