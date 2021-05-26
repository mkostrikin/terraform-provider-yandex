package yandex

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
	"github.com/yandex-cloud/go-sdk/sdkresolvers"
)

func dataSourceYandexALBBackendGroup() *schema.Resource {
	return &schema.Resource{
		Read:          dataSourceYandexALBBackendGroupRead,
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"backend_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"folder_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Computed: true,
			},

			"http_backend": {
				Type:          schema.TypeSet,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"grpc_backend"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"load_balancing_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"panic_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"locality_aware_routing_percent": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"strict_locality": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"healthcheck": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeout": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"interval": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"interval_jitter_percent": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"healthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"unhealthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"healthcheck_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"stream_healthcheck": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"receive": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"http_backend.healthcheck.grpc_healthcheck", "http_backend.healthcheck.http_healthcheck"},
									},
									"http_healthcheck": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"http2": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"http_backend.healthcheck.stream_healthcheck", "http_backend.healthcheck.grpc_healthcheck"},
									},
									"grpc_healthcheck": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"http_backend.healthcheck.stream_healthcheck", "http_backend.healthcheck.http_healthcheck"},
									},
								},
							},
							Set: resourceALBBackendGroupHealthcheckHash,
						},
						"tls": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sni": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"validation_context": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trusted_ca_id": {
													Type:          schema.TypeString,
													Optional:      true,
													Computed:      true,
													ConflictsWith: []string{"http_backend.tls.validation_context.trusted_ca_bytes"},
												},
												"trusted_ca_bytes": {
													Type:          schema.TypeString,
													Optional:      true,
													Computed:      true,
													ConflictsWith: []string{"http_backend.tls.validation_context.trusted_ca_id"},
												},
											},
										},
									},
								},
							},
						},
						"target_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http2": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
				Set: resourceALBBackendGroupBackendHash,
			},
			"grpc_backend": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"http_backend"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"load_balancing_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"panic_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"locality_aware_routing_percent": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"strict_locality": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"healthcheck": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeout": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"interval": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"interval_jitter_percent": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"healthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"unhealthy_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"healthcheck_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"stream_healthcheck": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"receive": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"grpc_backend.healthcheck.grpc_healthcheck", "grpc_backend.healthcheck.http_healthcheck"},
									},
									"http_healthcheck": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"http2": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"grpc_backend.healthcheck.stream_healthcheck", "grpc_backend.healthcheck.grpc_healthcheck"},
									},
									"grpc_healthcheck": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
										ConflictsWith: []string{"grpc_backend.healthcheck.stream_healthcheck", "grpc_backend.healthcheck.http_healthcheck"},
									},
								},
							},
							Set: resourceALBBackendGroupHealthcheckHash,
						},
						"tls": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sni": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"validation_context": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trusted_ca_id": {
													Type:          schema.TypeString,
													Optional:      true,
													Computed:      true,
													ConflictsWith: []string{"grpc_backend.tls.validation_context.trusted_ca_bytes"},
												},
												"trusted_ca_bytes": {
													Type:          schema.TypeString,
													Optional:      true,
													Computed:      true,
													ConflictsWith: []string{"grpc_backend.tls.validation_context.trusted_ca_id"},
												},
											},
										},
									},
								},
							},
						},
						"target_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				Set: resourceALBBackendGroupBackendHash,
			},

			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceYandexALBBackendGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ctx := config.Context()

	err := checkOneOf(d, "backend_group_id", "name")
	if err != nil {
		return err
	}

	bgID := d.Get("backend_group_id").(string)
	_, bgNameOk := d.GetOk("name")

	if bgNameOk {
		bgID, err = resolveObjectID(ctx, config, d, sdkresolvers.ALBBackendGroupResolver)
		if err != nil {
			return fmt.Errorf("failed to resolve data source ALB Backend Group by name: %v", err)
		}
	}

	bg, err := config.sdk.ApplicationLoadBalancer().BackendGroup().Get(ctx, &apploadbalancer.GetBackendGroupRequest{
		BackendGroupId: bgID,
	})

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ALB Backend Group with ID %q", bgID))
	}

	createdAt, err := getTimestamp(bg.CreatedAt)
	if err != nil {
		return err
	}

	d.Set("backend_group_id", bg.Id)
	d.Set("created_at", createdAt)
	d.Set("name", bg.Name)
	d.Set("folder_id", bg.FolderId)
	d.Set("description", bg.Description)

	if bg.GetHttp() != nil {
		backends, err := flattenALBHttpBackends(bg)
		if err != nil {
			return err
		}
		if err := d.Set("http_backend", backends); err != nil {
			return err
		}
	}

	if bg.GetGrpc() != nil {
		backends, err := flattenALBGrpcBackends(bg)
		if err != nil {
			return err
		}
		if err := d.Set("grpc_backend", backends); err != nil {
			return err
		}
	}

	if err := d.Set("labels", bg.Labels); err != nil {
		return err
	}

	d.SetId(bg.Id)

	return nil
}
