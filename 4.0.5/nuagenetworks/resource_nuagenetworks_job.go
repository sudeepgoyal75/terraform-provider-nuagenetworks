package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.5"
)

func resourceJob() *schema.Resource {
    return &schema.Resource{
        Create: resourceJobCreate,
        Read:   resourceJobRead,
        Update: resourceJobUpdate,
        Delete: resourceJobDelete,
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
            "parameters": &schema.Schema{
                Type:     schema.TypeMap,
                Optional: true,
                Computed: true,
            },
            "parameters_raw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "result": &schema.Schema{
                Type:     schema.TypeMap,
                Optional: true,
                Computed: true,
            },
            "result_raw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "command": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "progress": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_vcenter_hypervisor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_external_service_template_entry": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_redirection_target_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_redirection_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_policy_group_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_egress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_adv_fwd_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_external_service_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_acl_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_policy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_zfb_request": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_vsc": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_hsc": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_vrs": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_adv_fwd_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_egress_acl_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_ingress_acl_template", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_gateway", "parent_vsd"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_vsd"},
            },
            "parent_vsd": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vcenter_hypervisor", "parent_ingress_external_service_template_entry", "parent_l2_domain", "parent_redirection_target_template", "parent_l2_domain_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_ingress_adv_fwd_entry_template", "parent_ingress_external_service_template", "parent_ingress_acl_entry_template", "parent_policy_group", "parent_vcenter_cluster", "parent_zfb_request", "parent_vsc", "parent_hsc", "parent_vrs", "parent_ingress_adv_fwd_template", "parent_egress_acl_entry_template", "parent_ns_gateway", "parent_domain_template", "parent_vcenter", "parent_ingress_acl_template", "parent_gateway"},
            },
        },
    }
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Job object
    o := &vspk.Job{
        Command: d.Get("command").(string),
    }
    if attr, ok := d.GetOk("parameters"); ok {
        o.Parameters = attr.(interface{})
    }
    if attr, ok := d.GetOk("result"); ok {
        o.Result = attr.(interface{})
    }
    if attr, ok := d.GetOk("progress"); ok {
        o.Progress = attr.(float64)
    }
    if attr, ok := d.GetOk("assoc_entity_type"); ok {
        o.AssocEntityType = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
        parent := &vspk.VCenterHypervisor{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_external_service_template_entry"); ok {
        parent := &vspk.IngressExternalServiceTemplateEntry{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_redirection_target_template"); ok {
        parent := &vspk.RedirectionTargetTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_redirection_target"); ok {
        parent := &vspk.RedirectionTarget{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_policy_group_template"); ok {
        parent := &vspk.PolicyGroupTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_egress_acl_template"); ok {
        parent := &vspk.EgressACLTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_adv_fwd_entry_template"); ok {
        parent := &vspk.IngressAdvFwdEntryTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_external_service_template"); ok {
        parent := &vspk.IngressExternalServiceTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_acl_entry_template"); ok {
        parent := &vspk.IngressACLEntryTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_policy_group"); ok {
        parent := &vspk.PolicyGroup{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_zfb_request"); ok {
        parent := &vspk.ZFBRequest{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vsc"); ok {
        parent := &vspk.VSC{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_hsc"); ok {
        parent := &vspk.HSC{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vrs"); ok {
        parent := &vspk.VRS{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_adv_fwd_template"); ok {
        parent := &vspk.IngressAdvFwdTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_egress_acl_entry_template"); ok {
        parent := &vspk.EgressACLEntryTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter"); ok {
        parent := &vspk.VCenter{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
        parent := &vspk.IngressACLTemplate{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vsd"); ok {
        parent := &vspk.VSD{ID: attr.(string)}
        err := parent.CreateJob(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Job{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    if v, ok := o.Parameters.(string); ok {
		raw := make(map[string]string)
		raw["raw"] = v
		d.Set("parameters_raw", raw)
	} else {
		d.Set("parameters", o.Parameters)
	}
    d.Set("last_updated_by", o.LastUpdatedBy)
    if v, ok := o.Result.(string); ok {
		raw := make(map[string]string)
		raw["raw"] = v
		d.Set("result_raw", raw)
	} else {
		d.Set("result", o.Result)
	}
    d.Set("entity_scope", o.EntityScope)
    d.Set("command", o.Command)
    d.Set("progress", o.Progress)
    d.Set("assoc_entity_type", o.AssocEntityType)
    d.Set("status", o.Status)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Job{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Command = d.Get("command").(string)
    
    if attr, ok := d.GetOk("parameters"); ok {
        o.Parameters = attr.(interface{})
    }
    if attr, ok := d.GetOk("result"); ok {
        o.Result = attr.(interface{})
    }
    if attr, ok := d.GetOk("progress"); ok {
        o.Progress = attr.(float64)
    }
    if attr, ok := d.GetOk("assoc_entity_type"); ok {
        o.AssocEntityType = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Job{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}