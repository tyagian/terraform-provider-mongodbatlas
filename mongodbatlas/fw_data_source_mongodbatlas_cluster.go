package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"

	matlas "go.mongodb.org/atlas/mongodbatlas"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/mongodb/terraform-provider-mongodbatlas/mongodbatlas/framework/conversion"
)

const (
	ClusterDataSourceName = "cluster"
)

var _ datasource.DataSource = &ClusterDS{}
var _ datasource.DataSourceWithConfigure = &ClusterDS{}

func NewClusterDS() datasource.DataSource {
	return &ClusterDS{
		DSCommon: DSCommon{
			dataSourceName: ClusterDataSourceName,
		},
	}
}

type ClusterDS struct {
	DSCommon
}

func (d *ClusterDS) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: clusterDSAttributes(),
	}
}

func clusterDSAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"project_id": schema.StringAttribute{
			Required: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"advanced_configuration": clusterRSAdvancedConfigurationSchemaAttribute(),
		"auto_scaling_disk_gb_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"auto_scaling_compute_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"auto_scaling_compute_scale_down_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"backup_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"bi_connector_config": clusterDSBiConnectorConfigSchemaAttribute(),
		"cluster_type": schema.StringAttribute{
			Computed: true,
		},
		"connection_strings": clusterDSConnectionStringSchemaAttribute(),
		"disk_size_gb": schema.Float64Attribute{
			Computed: true,
		},
		"encryption_at_rest_provider": schema.StringAttribute{
			Computed: true,
		},
		"mongo_db_major_version": schema.StringAttribute{
			Computed: true,
		},
		"num_shards": schema.Int64Attribute{
			Computed: true,
		},
		"pit_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"provider_backup_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"provider_instance_size_name": schema.StringAttribute{
			Computed: true,
		},
		"provider_name": schema.StringAttribute{
			Computed: true,
		},
		"backing_provider_name": schema.StringAttribute{
			Computed: true,
		},
		"provider_disk_iops": schema.Int64Attribute{
			Computed: true,
		},
		"provider_disk_type_name": schema.StringAttribute{
			Computed: true,
		},
		"provider_encrypt_ebs_volume": schema.BoolAttribute{
			Computed: true,
		},
		"provider_encrypt_ebs_volume_flag": schema.BoolAttribute{
			Computed: true,
		},
		"provider_region_name": schema.StringAttribute{
			Computed: true,
		},
		"provider_volume_type": schema.StringAttribute{
			Computed: true,
		},
		"provider_auto_scaling_compute_max_instance_size": schema.StringAttribute{
			Computed: true,
		},
		"provider_auto_scaling_compute_min_instance_size": schema.StringAttribute{
			Computed: true,
		},
		"replication_factor": schema.Int64Attribute{
			Computed: true,
		},
		"replication_specs": clusterDSReplicationSpecsSchemaAttribute(),
		"mongo_db_version": schema.StringAttribute{
			Computed: true,
		},
		"mongo_uri": schema.StringAttribute{
			Computed: true,
		},
		"mongo_uri_updated": schema.StringAttribute{
			Computed: true,
		},
		"mongo_uri_with_options": schema.StringAttribute{
			Computed: true,
		},
		"paused": schema.BoolAttribute{
			Computed: true,
		},
		"srv_address": schema.StringAttribute{
			Computed: true,
		},
		"state_name": schema.StringAttribute{
			Computed: true,
		},
		"labels":                 clusterDSLabelsSchemaAttribute(),
		"tags":                   clusterDSTagsSchemaAttribute(),
		"snapshot_backup_policy": clusterRSSnapshotBackupPolicySchemaAttribute(),
		"termination_protection_enabled": schema.BoolAttribute{
			Computed: true,
		},
		"container_id": schema.StringAttribute{
			Computed: true,
		},
		"version_release_system": schema.StringAttribute{
			Computed: true,
		},
	}
}

func clusterDSTagsSchemaAttribute() schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Computed: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key": schema.StringAttribute{
					Computed: true,
				},
				"value": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	}
}

func clusterDSLabelsSchemaAttribute() schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Computed:           true,
		DeprecationMessage: fmt.Sprintf(DeprecationByDateWithReplacement, "September 2024", "tags"),
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key": schema.StringAttribute{
					Computed: true,
				},
				"value": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	}
}

func clusterDSReplicationSpecsSchemaAttribute() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Computed: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Computed: true,
				},
				"num_shards": schema.Int64Attribute{
					Computed: true,
				},
				"regions_config": schema.SetNestedAttribute{
					Computed: true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"region_name": schema.StringAttribute{
								Computed: true,
							},
							"electable_nodes": schema.Int64Attribute{
								Computed: true,
							},
							"priority": schema.Int64Attribute{
								Computed: true,
							},
							"read_only_nodes": schema.Int64Attribute{
								Computed: true,
							},
							"analytics_nodes": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
				},
				"zone_name": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	}
}

func clusterDSBiConnectorConfigSchemaAttribute() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Computed: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					Computed: true,
				},
				"read_preference": schema.StringAttribute{
					Computed: true,
				},
			},
		},
	}
}

func clusterDSConnectionStringSchemaAttribute() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Computed: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"standard": schema.StringAttribute{
					Computed: true,
				},
				"standard_srv": schema.StringAttribute{
					Computed: true,
				},
				"aws_private_link": schema.MapAttribute{
					Computed: true,
				},
				"aws_private_link_srv": schema.MapAttribute{
					Computed: true,
				},
				"private": schema.StringAttribute{
					Computed: true,
				},
				"private_srv": schema.StringAttribute{
					Computed: true,
				},
				"private_endpoint": schema.ListNestedAttribute{
					Computed: true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"connection_string": schema.StringAttribute{
								Computed: true,
							},
							"endpoints": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint_id": schema.StringAttribute{
											Computed: true,
										},
										"provider_name": schema.StringAttribute{
											Computed: true,
										},
										"region": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"srv_connection_string": schema.StringAttribute{
								Computed: true,
							},
							"srv_shard_optimized_connection_string": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *ClusterDS) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	conn := d.client.Atlas
	var clusterConfig tfClusterDSModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &clusterConfig)...)
	if resp.Diagnostics.HasError() {
		return
	}
	projectID := clusterConfig.ProjectID.ValueString()
	clusterName := clusterConfig.Name.ValueString()

	cluster, response, err := conn.Clusters.Get(ctx, projectID, clusterName)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.Diagnostics.AddError("cluster not found in Atlas", fmt.Sprintf(errorClusterRead, clusterName, err.Error()))
			return
		}
		resp.Diagnostics.AddError("error in getting cluster details from Atlas", fmt.Sprintf(errorClusterRead, clusterName, err.Error()))
		return
	}

	newClusterState, err := newTFClusterDSModel(ctx, conn, cluster)
	if err != nil {
		resp.Diagnostics.AddError("error in getting cluster details from Atlas", fmt.Sprintf(errorClusterRead, clusterName, err.Error()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &newClusterState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func newTFClusterDSModel(ctx context.Context, conn *matlas.Client, apiResp *matlas.Cluster) (*tfClusterDSModel, error) {
	var err error
	projectID := apiResp.GroupID
	clusterName := apiResp.Name

	clusterModel := tfClusterDSModel{
		AutoScalingComputeEnabled:                 types.BoolPointerValue(apiResp.AutoScaling.Compute.Enabled),
		AutoScalingComputeScaleDownEnabled:        types.BoolPointerValue(apiResp.AutoScaling.Compute.ScaleDownEnabled),
		ProviderAutoScalingComputeMinInstanceSize: types.StringValue(apiResp.ProviderSettings.AutoScaling.Compute.MinInstanceSize),
		ProviderAutoScalingComputeMaxInstanceSize: types.StringValue(apiResp.ProviderSettings.AutoScaling.Compute.MaxInstanceSize),
		AutoScalingDiskGbEnabled:                  types.BoolPointerValue(apiResp.AutoScaling.DiskGBEnabled),
		BackupEnabled:                             types.BoolPointerValue(apiResp.BackupEnabled),
		PitEnabled:                                types.BoolPointerValue(apiResp.PitEnabled),
		ProviderBackupEnabled:                     types.BoolPointerValue(apiResp.ProviderBackupEnabled),
		ClusterType:                               types.StringValue(apiResp.ClusterType),
		ConnectionStrings:                         newTFConnectionStringsModelDS(ctx, apiResp.ConnectionStrings),
		DiskSizeGb:                                types.Float64PointerValue(apiResp.DiskSizeGB),
		EncryptionAtRestProvider:                  types.StringValue(apiResp.EncryptionAtRestProvider),
		MongoDBMajorVersion:                       types.StringValue(apiResp.MongoDBMajorVersion),
		MongoDBVersion:                            types.StringValue(apiResp.MongoDBVersion),
		MongoURI:                                  types.StringValue(apiResp.MongoURI),
		MongoURIUpdated:                           types.StringValue(apiResp.MongoURIUpdated),
		MongoURIWithOptions:                       types.StringValue(apiResp.MongoURIWithOptions),
		Paused:                                    types.BoolPointerValue(apiResp.Paused),
		SrvAddress:                                types.StringValue(apiResp.SrvAddress),
		StateName:                                 types.StringValue(apiResp.StateName),
		BiConnectorConfig:                         newTFBiConnectorConfigModel(apiResp.BiConnector),
		ReplicationFactor:                         types.Int64PointerValue(apiResp.ReplicationFactor),
		ReplicationSpecs:                          newTFReplicationSpecsModel(apiResp.ReplicationSpecs),
		Labels:                                    removeDefaultLabel(newTFLabelsModel(apiResp.Labels)),
		Tags:                                      newTFTagsModel(apiResp.Tags),
		TerminationProtectionEnabled:              types.BoolPointerValue(apiResp.TerminationProtectionEnabled),
		VersionReleaseSystem:                      types.StringValue(apiResp.VersionReleaseSystem),
		ProjectID:                                 types.StringValue(projectID),
		Name:                                      types.StringValue(clusterName),
		ID:                                        types.StringValue(apiResp.ID),
	}

	// Avoid Global Cluster issues. (NumShards is not present in Global Clusters)
	if numShards := apiResp.NumShards; numShards != nil {
		clusterModel.NumShards = types.Int64PointerValue(numShards)
	}

	if apiResp.ProviderSettings != nil {
		setTFProviderSettingsDS(&clusterModel, apiResp.ProviderSettings)

		if v := apiResp.ProviderSettings.ProviderName; v != "TENANT" {
			containers, _, err := conn.Containers.List(ctx, projectID,
				&matlas.ContainersListOptions{ProviderName: v})
			if err != nil {
				return nil, fmt.Errorf(errorClusterRead, clusterName, err)
			}

			clusterModel.ContainerID = types.StringValue(getContainerID(containers, apiResp))
			// clusterModel.AutoScalingDiskGBEnabled = types.BoolPointerValue(apiResp.AutoScaling.DiskGBEnabled)
		}
	}

	clusterModel.AdvancedConfiguration, err = newTFAdvancedConfigurationModelDSFromAtlas(ctx, conn, projectID, apiResp.Name)
	if err != nil {
		return nil, err
	}

	clusterModel.SnapshotBackupPolicy, err = newTFSnapshotBackupPolicyDSModel(ctx, conn, projectID, clusterName)
	if err != nil {
		return nil, err
	}

	return &clusterModel, nil
}

func newTFSnapshotBackupPolicyDSModel(ctx context.Context, conn *matlas.Client, projectID, clusterName string) ([]*tfSnapshotBackupPolicyModel, error) {
	res, err := newTFSnapshotBackupPolicyModel(ctx, conn, projectID, clusterName)
	if err != nil {
		return nil, fmt.Errorf(errorSnapshotBackupPolicyRead, clusterName, err)
	}

	return res, nil
}

func newTFAdvancedConfigurationModelDSFromAtlas(ctx context.Context, conn *matlas.Client, projectID, clusterName string) ([]*tfAdvancedConfigurationModel, error) {
	processArgs, _, err := conn.Clusters.GetProcessArgs(ctx, projectID, clusterName)
	if err != nil {
		return nil, err
	}

	advConfigModel := newTfAdvancedConfigurationModel(ctx, processArgs)
	return advConfigModel, err
}

func newTFConnectionStringsModelDS(ctx context.Context, connString *matlas.ConnectionStrings) []*tfConnectionStringDSModel {
	res := []*tfConnectionStringDSModel{}

	if connString != nil {
		res = append(res, &tfConnectionStringDSModel{
			Standard:        conversion.StringNullIfEmpty(connString.Standard),
			StandardSrv:     conversion.StringNullIfEmpty(connString.StandardSrv),
			Private:         conversion.StringNullIfEmpty(connString.Private),
			PrivateSrv:      conversion.StringNullIfEmpty(connString.PrivateSrv),
			PrivateEndpoint: newTFPrivateEndpointModel(ctx, connString.PrivateEndpoint),
		})
	}
	return res
}

func setTFProviderSettingsDS(clusterModel *tfClusterDSModel, settings *matlas.ProviderSettings) {
	if settings.ProviderName == "TENANT" {
		clusterModel.BackingProviderName = types.StringValue(settings.BackingProviderName)
	}

	if settings.DiskIOPS != nil && *settings.DiskIOPS != 0 {
		clusterModel.ProviderDiskIops = types.Int64PointerValue(settings.DiskIOPS)
	}
	if settings.EncryptEBSVolume != nil {
		clusterModel.ProviderEncryptEbsVolumeFlag = types.BoolPointerValue(settings.EncryptEBSVolume)
		clusterModel.ProviderEncryptEbsVolume = types.BoolPointerValue(settings.EncryptEBSVolume)
	}
	clusterModel.ProviderDiskTypeName = types.StringValue(settings.DiskTypeName)
	clusterModel.ProviderInstanceSizeName = types.StringValue(settings.InstanceSizeName)
	clusterModel.ProviderName = types.StringValue(settings.ProviderName)
	clusterModel.ProviderRegionName = types.StringValue(settings.RegionName)
	clusterModel.ProviderVolumeType = types.StringValue(settings.VolumeType)
}

type tfClusterDSModel struct {
	DiskSizeGb                                types.Float64                   `tfsdk:"disk_size_gb"`
	ProviderAutoScalingComputeMaxInstanceSize types.String                    `tfsdk:"provider_auto_scaling_compute_max_instance_size"`
	EncryptionAtRestProvider                  types.String                    `tfsdk:"encryption_at_rest_provider"`
	VersionReleaseSystem                      types.String                    `tfsdk:"version_release_system"`
	StateName                                 types.String                    `tfsdk:"state_name"`
	ClusterType                               types.String                    `tfsdk:"cluster_type"`
	ContainerID                               types.String                    `tfsdk:"container_id"`
	SrvAddress                                types.String                    `tfsdk:"srv_address"`
	ProviderVolumeType                        types.String                    `tfsdk:"provider_volume_type"`
	ID                                        types.String                    `tfsdk:"id"`
	MongoDBMajorVersion                       types.String                    `tfsdk:"mongo_db_major_version"`
	MongoDBVersion                            types.String                    `tfsdk:"mongo_db_version"`
	MongoURI                                  types.String                    `tfsdk:"mongo_uri"`
	ProviderAutoScalingComputeMinInstanceSize types.String                    `tfsdk:"provider_auto_scaling_compute_min_instance_size"`
	MongoURIWithOptions                       types.String                    `tfsdk:"mongo_uri_with_options"`
	Name                                      types.String                    `tfsdk:"name"`
	ProviderRegionName                        types.String                    `tfsdk:"provider_region_name"`
	ProviderName                              types.String                    `tfsdk:"provider_name"`
	ProviderInstanceSizeName                  types.String                    `tfsdk:"provider_instance_size_name"`
	ProjectID                                 types.String                    `tfsdk:"project_id"`
	ProviderDiskTypeName                      types.String                    `tfsdk:"provider_disk_type_name"`
	MongoURIUpdated                           types.String                    `tfsdk:"mongo_uri_updated"`
	BackingProviderName                       types.String                    `tfsdk:"backing_provider_name"`
	ConnectionStrings                         []*tfConnectionStringDSModel    `tfsdk:"connection_strings"`
	SnapshotBackupPolicy                      []*tfSnapshotBackupPolicyModel  `tfsdk:"snapshot_backup_policy"`
	AdvancedConfiguration                     []*tfAdvancedConfigurationModel `tfsdk:"advanced_configuration"`
	ReplicationSpecs                          []*tfReplicationSpecModel       `tfsdk:"replication_specs"`
	Tags                                      []*tfTagModel                   `tfsdk:"tags"`
	Labels                                    []tfLabelModel                  `tfsdk:"labels"`
	BiConnectorConfig                         []*tfBiConnectorConfigModel     `tfsdk:"bi_connector_config"`
	ProviderDiskIops                          types.Int64                     `tfsdk:"provider_disk_iops"`
	NumShards                                 types.Int64                     `tfsdk:"num_shards"`
	ReplicationFactor                         types.Int64                     `tfsdk:"replication_factor"`
	Paused                                    types.Bool                      `tfsdk:"paused"`
	ProviderEncryptEbsVolume                  types.Bool                      `tfsdk:"provider_encrypt_ebs_volume"`
	ProviderEncryptEbsVolumeFlag              types.Bool                      `tfsdk:"provider_encrypt_ebs_volume_flag"`
	AutoScalingComputeEnabled                 types.Bool                      `tfsdk:"auto_scaling_compute_enabled"`
	ProviderBackupEnabled                     types.Bool                      `tfsdk:"provider_backup_enabled"`
	AutoScalingDiskGbEnabled                  types.Bool                      `tfsdk:"auto_scaling_disk_gb_enabled"`
	PitEnabled                                types.Bool                      `tfsdk:"pit_enabled"`
	BackupEnabled                             types.Bool                      `tfsdk:"backup_enabled"`
	TerminationProtectionEnabled              types.Bool                      `tfsdk:"termination_protection_enabled"`
	AutoScalingComputeScaleDownEnabled        types.Bool                      `tfsdk:"auto_scaling_compute_scale_down_enabled"`
}

type tfConnectionStringDSModel struct {
	Standard          types.String `tfsdk:"standard"`
	StandardSrv       types.String `tfsdk:"standard_srv"`
	AwsPrivateLink    types.String `tfsdk:"aws_private_link"`
	AwsPrivateLinkSrv types.String `tfsdk:"aws_private_link_srv"`
	Private           types.String `tfsdk:"private"`
	PrivateSrv        types.String `tfsdk:"private_srv"`
	// PrivateEndpoint []tfPrivateEndpointModel `tfsdk:"private_endpoint"`
	PrivateEndpoint types.List `tfsdk:"private_endpoint"`
}
