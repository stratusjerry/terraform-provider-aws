## Regex Remove Stuff I don't care about
```
^\* (resource|data-source)/aws_(ami|api|app|auto|back|ce|cloud|code|cog|conf|con|data|db|dms|dyna|ebs|ecr|ecs|eip|eks|elas|emr|flow|glac|guar|iam|image|iot|kine|kms|lam|launch|light|msk|net|org|pin|rds|red|route|sage|s3|secur|ses|spot|ssm|stor|sub|trans|vpc|waf).*\n
```

## 4.24.0 (Unreleased)

ENHANCEMENTS:


## 4.24.0 (Unreleased)

FEATURES:

* **New Resource:** `aws_acmpca_permission` ([#12485](https://github.com/hashicorp/terraform-provider-aws/issues/12485))
* **New Resource:** `aws_ssm_service_setting` ([#13018](https://github.com/hashicorp/terraform-provider-aws/issues/13018))

ENHANCEMENTS:


BUG FIXES:


## 4.23.0 (July 22, 2022)

FEATURES:

* **New Data Source:** `aws_connect_user_hierarchy_group` ([#24777](https://github.com/hashicorp/terraform-provider-aws/issues/24777))
* **New Data Source:** `aws_location_geofence_collection` ([#25844](https://github.com/hashicorp/terraform-provider-aws/issues/25844))
* **New Data Source:** `aws_networkfirewall_firewall_policy` ([#24748](https://github.com/hashicorp/terraform-provider-aws/issues/24748))
* **New Data Source:** `aws_s3_account_public_access_block` ([#25781](https://github.com/hashicorp/terraform-provider-aws/issues/25781))
* **New Resource:** `aws_connect_user` ([#24832](https://github.com/hashicorp/terraform-provider-aws/issues/24832))
* **New Resource:** `aws_connect_vocabulary` ([#24849](https://github.com/hashicorp/terraform-provider-aws/issues/24849))
* **New Resource:** `aws_location_geofence_collection` ([#25762](https://github.com/hashicorp/terraform-provider-aws/issues/25762))
* **New Resource:** `aws_redshiftserverless_namespace` ([#25889](https://github.com/hashicorp/terraform-provider-aws/issues/25889))
* **New Resource:** `aws_rolesanywhere_profile` ([#25850](https://github.com/hashicorp/terraform-provider-aws/issues/25850))
* **New Resource:** `aws_rolesanywhere_trust_anchor` ([#25779](https://github.com/hashicorp/terraform-provider-aws/issues/25779))
* **New Resource:** `aws_transcribe_vocabulary` ([#25863](https://github.com/hashicorp/terraform-provider-aws/issues/25863))
* **New Resource:** `aws_transcribe_vocabulary_filter` ([#25918](https://github.com/hashicorp/terraform-provider-aws/issues/25918))

ENHANCEMENTS:

* data/aws_outposts_asset: Add `rack_elevation` attribute ([#25822](https://github.com/hashicorp/terraform-provider-aws/issues/25822))
* resource/aws_fsx_openzfs_file_system: Allow in-place update of `storage_capacity`, `throughput_capacity`, and `disk_iops_configuration`. ([#25841](https://github.com/hashicorp/terraform-provider-aws/issues/25841))

BUG FIXES:


## 4.22.0 (July  8, 2022)

FEATURES:

* **New Data Source:** `aws_location_route_calculator` ([#25689](https://github.com/hashicorp/terraform-provider-aws/issues/25689))
* **New Data Source:** `aws_location_tracker` ([#25639](https://github.com/hashicorp/terraform-provider-aws/issues/25639))
* **New Data Source:** `aws_secretsmanager_random_password` ([#25704](https://github.com/hashicorp/terraform-provider-aws/issues/25704))
* **New Resource:** `aws_directory_service_shared_directory` ([#24766](https://github.com/hashicorp/terraform-provider-aws/issues/24766))
* **New Resource:** `aws_directory_service_shared_directory_accepter` ([#24766](https://github.com/hashicorp/terraform-provider-aws/issues/24766))
* **New Resource:** `aws_lightsail_database` ([#18663](https://github.com/hashicorp/terraform-provider-aws/issues/18663))
* **New Resource:** `aws_location_route_calculator` ([#25656](https://github.com/hashicorp/terraform-provider-aws/issues/25656))
* **New Resource:** `aws_transcribe_medical_vocabulary` ([#25723](https://github.com/hashicorp/terraform-provider-aws/issues/25723))

ENHANCEMENTS:

* resource/aws_acmpca_certificate_authority: Add `revocation_configuration.ocsp_configuration` argument ([#25720](https://github.com/hashicorp/terraform-provider-aws/issues/25720))
* resource/aws_gamelift_game_session_queue: Add `notification_target` argument ([#25544](https://github.com/hashicorp/terraform-provider-aws/issues/25544))
* resource/aws_placement_group: Add `spread_level` argument ([#25615](https://github.com/hashicorp/terraform-provider-aws/issues/25615))

BUG FIXES:

* provider: Ensure that the configured `assume_role_with_web_identity` value is used ([#25681](https://github.com/hashicorp/terraform-provider-aws/issues/25681))
* resource/aws_acmpca_certificate_authority: Fix crash when `revocation_configuration` block is empty ([#25695](https://github.com/hashicorp/terraform-provider-aws/issues/25695))
* resource/aws_ec2_transit_gateway: Fix MaxItems and subnet size validation in `transit_gateway_cidr_blocks` ([#25673](https://github.com/hashicorp/terraform-provider-aws/issues/25673))

## 4.21.0 (June 30, 2022)

FEATURES:

* **New Data Source:** `aws_kendra_experience` ([#25601](https://github.com/hashicorp/terraform-provider-aws/issues/25601))
* **New Data Source:** `aws_kendra_query_suggestions_block_list` ([#25592](https://github.com/hashicorp/terraform-provider-aws/issues/25592))
* **New Data Source:** `aws_kendra_thesaurus` ([#25555](https://github.com/hashicorp/terraform-provider-aws/issues/25555))
* **New Data Source:** `aws_service_discovery_http_namespace` ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))
* **New Data Source:** `aws_service_discovery_service` ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))
* **New Resource:** `aws_accessanalyzer_archive_rule` ([#25514](https://github.com/hashicorp/terraform-provider-aws/issues/25514))
* **New Resource:** `aws_apprunner_observability_configuration` ([#25591](https://github.com/hashicorp/terraform-provider-aws/issues/25591))
* **New Resource:** `aws_lakeformation_resource_lf_tags` ([#25565](https://github.com/hashicorp/terraform-provider-aws/issues/25565))

ENHANCEMENTS:

* data-source/aws_service_discovery_dns_namespace: Add `tags` attribute ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))
* data/aws_key_pair: New attribute `public_key` populated by setting the new `include_public_key` argument ([#25371](https://github.com/hashicorp/terraform-provider-aws/issues/25371))
* resource/aws_key_pair: Added 2 new attributes - `key_type` and `create_time` ([#25371](https://github.com/hashicorp/terraform-provider-aws/issues/25371))
* resource/aws_service_discovery_http_namespace: Add `http_name` attribute ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))

BUG FIXES:


## 4.20.1 (June 24, 2022)

BUG FIXES:

* resource/aws_default_vpc_dhcp_options: Fix `missing expected [` error introduced in [v4.20.0](https://github.com/hashicorp/terraform-provider-aws/blob/main/CHANGELOG.md#4200-june-23-2022) ([#25562](https://github.com/hashicorp/terraform-provider-aws/issues/25562))

## 4.20.0 (June 23, 2022)

FEATURES:

* **New Data Source:** `aws_kendra_faq` ([#25523](https://github.com/hashicorp/terraform-provider-aws/issues/25523))
* **New Data Source:** `aws_kendra_index` ([#25473](https://github.com/hashicorp/terraform-provider-aws/issues/25473))
* **New Data Source:** `aws_outposts_asset` ([#25476](https://github.com/hashicorp/terraform-provider-aws/issues/25476))
* **New Data Source:** `aws_outposts_assets` ([#25476](https://github.com/hashicorp/terraform-provider-aws/issues/25476))
* **New Resource:** `aws_applicationinsights_application` ([#25195](https://github.com/hashicorp/terraform-provider-aws/issues/25195))
* **New Resource:** `aws_ce_anomaly_monitor` ([#25177](https://github.com/hashicorp/terraform-provider-aws/issues/25177))
* **New Resource:** `aws_ce_anomaly_subscription` ([#25224](https://github.com/hashicorp/terraform-provider-aws/issues/25224))
* **New Resource:** `aws_ce_cost_allocation_tag` ([#25272](https://github.com/hashicorp/terraform-provider-aws/issues/25272))
* **New Resource:** `aws_cloudwatchrum_app_monitor` ([#25180](https://github.com/hashicorp/terraform-provider-aws/issues/25180))
* **New Resource:** `aws_cognito_risk_configuration` ([#25282](https://github.com/hashicorp/terraform-provider-aws/issues/25282))
* **New Resource:** `aws_kendra_experience` ([#25315](https://github.com/hashicorp/terraform-provider-aws/issues/25315))
* **New Resource:** `aws_kendra_faq` ([#25515](https://github.com/hashicorp/terraform-provider-aws/issues/25515))
* **New Resource:** `aws_kendra_query_suggestions_block_list` ([#25198](https://github.com/hashicorp/terraform-provider-aws/issues/25198))
* **New Resource:** `aws_kendra_thesaurus` ([#25199](https://github.com/hashicorp/terraform-provider-aws/issues/25199))
* **New Resource:** `aws_lakeformation_lf_tag` ([#19523](https://github.com/hashicorp/terraform-provider-aws/issues/19523))
* **New Resource:** `aws_location_tracker` ([#25466](https://github.com/hashicorp/terraform-provider-aws/issues/25466))

ENHANCEMENTS:

* data-source/aws_instance: Add `disable_api_stop` attribute ([#25185](https://github.com/hashicorp/terraform-provider-aws/issues/25185))
* data-source/aws_instance: Add `private_dns_name_options` attribute ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* data-source/aws_instance: Correctly set `credit_specification` for T4g instances ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_ec2_host: Add `outpost_arn` argument ([#25464](https://github.com/hashicorp/terraform-provider-aws/issues/25464))
* resource/aws_instance: Add `disable_api_stop` argument ([#25185](https://github.com/hashicorp/terraform-provider-aws/issues/25185))
* resource/aws_instance: Add `private_dns_name_options` argument ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_instance: Correctly handle `credit_specification` for T4g instances ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_synthetics_canary: Add `delete_lambda` argument ([#25284](https://github.com/hashicorp/terraform-provider-aws/issues/25284))
* resource/aws_vpn_connection: Add `outside_ip_address_type` and `transport_transit_gateway_attachment_id` arguments in support of [Private IP VPNs](https://docs.aws.amazon.com/vpn/latest/s2svpn/private-ip-dx.html) ([#25529](https://github.com/hashicorp/terraform-provider-aws/issues/25529))

BUG FIXES:

* data-source/aws_lb: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_lb_listener: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_lb_target_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_sqs_queue: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_keyspaces_table: Relax validation of the `schema_definition.column.type` argument to allow collection types ([#25230](https://github.com/hashicorp/terraform-provider-aws/issues/25230))
* resource/aws_lb: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_lb_listener: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_lb_listener_rule: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_lb_target_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_sns_topic: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_sqs_queue: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))

## 4.19.0 (June 17, 2022)

FEATURES:

* **New Resource:** `aws_kendra_index` ([#24920](https://github.com/hashicorp/terraform-provider-aws/issues/24920))
* **New Resource:** `aws_lightsail_container_service` ([#20625](https://github.com/hashicorp/terraform-provider-aws/issues/20625))
* **New Resource:** `aws_lightsail_container_service_deployment_version` ([#20625](https://github.com/hashicorp/terraform-provider-aws/issues/20625))

BUG FIXES:

* resource/aws_ec2_managed_prefix_list_entry: Fix error when attempting to create or delete multiple list entries ([#25046](https://github.com/hashicorp/terraform-provider-aws/issues/25046))

## 4.18.0 (June 10, 2022)

FEATURES:

* **New Resource:** `aws_ce_anomaly_monitor` ([#25177](https://github.com/hashicorp/terraform-provider-aws/issues/25177))
* **New Resource:** `aws_emrserverless_application` ([#25144](https://github.com/hashicorp/terraform-provider-aws/issues/25144))

ENHANCEMENTS:

* data-source/aws_ram_resource_share: Add `resource_share_status` argument. ([#25159](https://github.com/hashicorp/terraform-provider-aws/issues/25159))

BUG FIXES:

* resource/aws_snapshot_create_volume_permission: Error if `account_id` is the snapshot's owner ([#12103](https://github.com/hashicorp/terraform-provider-aws/issues/12103))

## 4.17.1 (June  3, 2022)

BUG FIXES:

* resource/aws_ram_resource_share: Fix regression in v4.17.0 where `permission_arns` would get clobbered if already set ([#25158](https://github.com/hashicorp/terraform-provider-aws/issues/25158))

## 4.17.0 (June  3, 2022)

FEATURES:

* **New Data Source:** `aws_redshift_cluster_credentials` ([#25092](https://github.com/hashicorp/terraform-provider-aws/issues/25092))
* **New Resource:** `aws_acmpca_policy` ([#25109](https://github.com/hashicorp/terraform-provider-aws/issues/25109))
* **New Resource:** `aws_redshift_cluster_iam_roles` ([#25096](https://github.com/hashicorp/terraform-provider-aws/issues/25096))
* **New Resource:** `aws_redshift_hsm_configuration` ([#25093](https://github.com/hashicorp/terraform-provider-aws/issues/25093))
* **New Resource:** `aws_redshiftdata_statement` ([#25104](https://github.com/hashicorp/terraform-provider-aws/issues/25104))

ENHANCEMENTS:

* resource/aws_ram_resource_share: Add `permission_arns` argument. ([#25113](https://github.com/hashicorp/terraform-provider-aws/issues/25113))

BUG FIXES:

* resource/aws_servicecatalog_provisioned_product: Correctly handle resources in a `TAINTED` state ([#25130](https://github.com/hashicorp/terraform-provider-aws/issues/25130))

## 4.16.0 (May 27, 2022)

FEATURES:

* **New Data Source:** `aws_location_place_index` ([#24980](https://github.com/hashicorp/terraform-provider-aws/issues/24980))
* **New Data Source:** `aws_redshift_subnet_group` ([#25053](https://github.com/hashicorp/terraform-provider-aws/issues/25053))
* **New Resource:** `aws_efs_replication_configuration` ([#22844](https://github.com/hashicorp/terraform-provider-aws/issues/22844))
* **New Resource:** `aws_location_place_index` ([#24821](https://github.com/hashicorp/terraform-provider-aws/issues/24821))
* **New Resource:** `aws_redshift_authentication_profile` ([#24907](https://github.com/hashicorp/terraform-provider-aws/issues/24907))
* **New Resource:** `aws_redshift_endpoint_access` ([#25073](https://github.com/hashicorp/terraform-provider-aws/issues/25073))
* **New Resource:** `aws_redshift_hsm_client_certificate` ([#24906](https://github.com/hashicorp/terraform-provider-aws/issues/24906))
* **New Resource:** `aws_redshift_usage_limit` ([#24916](https://github.com/hashicorp/terraform-provider-aws/issues/24916))

ENHANCEMENTS:


BUG FIXES:

* resource/aws_instance: Correctly delete instance on destroy when `disable_api_termination` is `true` ([#19277](https://github.com/hashicorp/terraform-provider-aws/issues/19277))
* resource/aws_instance: Prevent error `InvalidParameterCombination: The parameter GroupName within placement information cannot be specified when instanceInterruptionBehavior is set to 'STOP'` when using a launch template that sets `instance_interruption_behavior` to `stop` ([#24695](https://github.com/hashicorp/terraform-provider-aws/issues/24695))
* resource/aws_servicecatalog_provisioned_product: Add possible `TAINTED` target state for resource update and remove one of the internal waiters during read ([#24804](https://github.com/hashicorp/terraform-provider-aws/issues/24804))

## 4.15.1 (May 20, 2022)

BUG FIXES:


## 4.15.0 (May 20, 2022)

BREAKING CHANGES:


FEATURES:

* **New Data Source:** `aws_lb_hosted_zone_id` ([#24749](https://github.com/hashicorp/terraform-provider-aws/issues/24749))
* **New Data Source:** `aws_networkmanager_core_network_policy_document` ([#24368](https://github.com/hashicorp/terraform-provider-aws/issues/24368))
* **New Resource:** `aws_db_snapshot_copy` ([#9886](https://github.com/hashicorp/terraform-provider-aws/issues/9886))
* **New Resource:** `aws_keyspaces_table` ([#24351](https://github.com/hashicorp/terraform-provider-aws/issues/24351))

ENHANCEMENTS:

* resource/aws_ec2_fleet: Add `arn` attribute ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))
* resource/aws_ec2_fleet: Add `launch_template_config.override.instance_requirements` argument ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))
* resource/aws_ec2_fleet: Add support for `capacity-optimized` and `capacity-optimized-prioritized` values for `spot_options.allocation_strategy` ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))

BUG FIXES:


## 4.14.0 (May 13, 2022)

FEATURES:

* **New Data Source:** `aws_connect_routing_profile` ([#23525](https://github.com/hashicorp/terraform-provider-aws/issues/23525))
* **New Data Source:** `aws_connect_security_profile` ([#23524](https://github.com/hashicorp/terraform-provider-aws/issues/23524))
* **New Data Source:** `aws_connect_user_hierarchy_structure` ([#23527](https://github.com/hashicorp/terraform-provider-aws/issues/23527))
* **New Data Source:** `aws_location_map` ([#24693](https://github.com/hashicorp/terraform-provider-aws/issues/24693))
* **New Resource:** `aws_appflow_connector_profile` ([#23892](https://github.com/hashicorp/terraform-provider-aws/issues/23892))
* **New Resource:** `aws_appflow_flow` ([#24017](https://github.com/hashicorp/terraform-provider-aws/issues/24017))
* **New Resource:** `aws_appintegrations_event_integration` ([#23904](https://github.com/hashicorp/terraform-provider-aws/issues/23904))
* **New Resource:** `aws_connect_user_hierarchy_group` ([#23531](https://github.com/hashicorp/terraform-provider-aws/issues/23531))
* **New Resource:** `aws_location_map` ([#24682](https://github.com/hashicorp/terraform-provider-aws/issues/24682))

ENHANCEMENTS:

* data-source/aws_acm_certificate: Add `certificate` and `certificate_chain` attributes ([#24593](https://github.com/hashicorp/terraform-provider-aws/issues/24593))
* resource/aws_servicecatalog_provisioned_product: Wait for provisioning to finish ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))
* resource/aws_servicecatalog_provisioned_product: Wait for update to finish ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))

BUG FIXES:

* resource/aws_alb_listener_rule: Don't force recreate listener rule on priority change. ([#23768](https://github.com/hashicorp/terraform-provider-aws/issues/23768))
* resource/aws_default_subnet: Fix `InvalidSubnet.Conflict` errors when associating IPv6 CIDR blocks ([#24685](https://github.com/hashicorp/terraform-provider-aws/issues/24685))
* resource/aws_servicecatalog_provisioned_product: Prevent error when retrieving a provisioned product in a non-available state ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))

## 4.13.0 (May  5, 2022)

FEATURES:

* **New Data Source:** `aws_emrcontainers_virtual_cluster` ([#20003](https://github.com/hashicorp/terraform-provider-aws/issues/20003))
* **New Data Source:** `aws_iam_instance_profiles` ([#24423](https://github.com/hashicorp/terraform-provider-aws/issues/24423))
* **New Data Source:** `aws_secretsmanager_secrets` ([#24514](https://github.com/hashicorp/terraform-provider-aws/issues/24514))
* **New Resource:** `aws_emrcontainers_virtual_cluster` ([#20003](https://github.com/hashicorp/terraform-provider-aws/issues/20003))
* **New Resource:** `aws_iot_topic_rule_destination` ([#24395](https://github.com/hashicorp/terraform-provider-aws/issues/24395))

ENHANCEMENTS:


BUG FIXES:


## 4.12.1 (April 29, 2022)

ENHANCEMENTS:


BUG FIXES:

* resource/aws_acm_certificate_validation: Restore certificate issuance timestamp as the resource `id` value, fixing error on existing resource Read ([#24453](https://github.com/hashicorp/terraform-provider-aws/issues/24453))

## 4.12.0 (April 28, 2022)

FEATURES:

* **New Data Source:** `aws_ce_cost_category` ([#24402](https://github.com/hashicorp/terraform-provider-aws/issues/24402))
* **New Data Source:** `aws_ce_tags` ([#24402](https://github.com/hashicorp/terraform-provider-aws/issues/24402))
* **New Data Source:** `aws_cloudfront_origin_access_identities` ([#24382](https://github.com/hashicorp/terraform-provider-aws/issues/24382))
* **New Data Source:** `aws_mq_broker_instance_type_offerings` ([#24394](https://github.com/hashicorp/terraform-provider-aws/issues/24394))
* **New Resource:** `aws_athena_data_catalog` ([#22968](https://github.com/hashicorp/terraform-provider-aws/issues/22968))
* **New Resource:** `aws_ce_cost_category` ([#24402](https://github.com/hashicorp/terraform-provider-aws/issues/24402))
* **New Resource:** `aws_docdb_event_subscription` ([#24379](https://github.com/hashicorp/terraform-provider-aws/issues/24379))

ENHANCEMENTS:

* data-source/aws_grafana_workspace: Add `tags` attribute ([#24358](https://github.com/hashicorp/terraform-provider-aws/issues/24358))
* data-source/aws_instance: Add `maintenance_options` attribute ([#24377](https://github.com/hashicorp/terraform-provider-aws/issues/24377))
* provider: Add support for Assume Role with Web Identity. ([#24441](https://github.com/hashicorp/terraform-provider-aws/issues/24441))
* resource/aws_acm_certificate: Add `validation_option` argument ([#3853](https://github.com/hashicorp/terraform-provider-aws/issues/3853))
* resource/aws_acm_certificate_validation: Increase default resource Create (certificate issuance) timeout to 75 minutes ([#20073](https://github.com/hashicorp/terraform-provider-aws/issues/20073))
* resource/aws_grafana_workspace: Add `tags` argument ([#24358](https://github.com/hashicorp/terraform-provider-aws/issues/24358))
* resource/aws_instance: Add `maintenance_options` argument ([#24377](https://github.com/hashicorp/terraform-provider-aws/issues/24377))
* resource/aws_mq_broker: Make `maintenance_window_start_time` updateable without recreation. ([#24385](https://github.com/hashicorp/terraform-provider-aws/issues/24385))

BUG FIXES:

* provider: Setting `skip_metadata_api_check = false` now overrides `AWS_EC2_METADATA_DISABLED` environment variable. ([#24441](https://github.com/hashicorp/terraform-provider-aws/issues/24441))
* resource/aws_acm_certificate: Correctly handle SAN entries that match `domain_name` ([#20073](https://github.com/hashicorp/terraform-provider-aws/issues/20073))
* resource/aws_ec2_availability_zone_group: Don't crash if `group_name` is not found ([#24422](https://github.com/hashicorp/terraform-provider-aws/issues/24422))
* resource/aws_lb: Fix bug causing an error on update if tags unsupported in ISO region ([#24334](https://github.com/hashicorp/terraform-provider-aws/issues/24334))

## 4.11.0 (April 22, 2022)

FEATURES:

* **New Data Source:** `aws_s3_bucket_policy` ([#17738](https://github.com/hashicorp/terraform-provider-aws/issues/17738))
* **New Resource:** `aws_transfer_workflow` ([#24248](https://github.com/hashicorp/terraform-provider-aws/issues/24248))

ENHANCEMENTS:

* data-source/aws_opensearch_domain: Add `cold_storage_options` attribute to the `cluster_config` configuration block ([#24284](https://github.com/hashicorp/terraform-provider-aws/issues/24284))
* resource/aws_fsx_ontap_file_system: Add support for `SINGLE_AZ_1` `deployment_type`. ([#24280](https://github.com/hashicorp/terraform-provider-aws/issues/24280))
* resource/aws_instance: Add `capacity_reservation_specification.capacity_reservation_target.capacity_reservation_resource_group_arn` argument ([#24283](https://github.com/hashicorp/terraform-provider-aws/issues/24283))
* resource/aws_instance: Add `network_interface.network_card_index` argument ([#24283](https://github.com/hashicorp/terraform-provider-aws/issues/24283))
* resource/aws_opensearch_domain: Add `cold_storage_options` argument to the `cluster_config` configuration block ([#24284](https://github.com/hashicorp/terraform-provider-aws/issues/24284))
* resource/aws_opensearch_domain: For Elasticsearch versions 6.7+, allow in-place update of `node_to_node_encryption.0.enabled` and `encrypt_at_rest.0.enabled`. ([#24222](https://github.com/hashicorp/terraform-provider-aws/issues/24222))
* resource/aws_workspaces_workspace: Additional supported values for `workspace_properties.compute_type_name` argument ([#24286](https://github.com/hashicorp/terraform-provider-aws/issues/24286))

BUG FIXES:

* data-source/aws_efs_file_system: Prevent panic when searching by tag returns 0 or multiple results ([#24298](https://github.com/hashicorp/terraform-provider-aws/issues/24298))
* resource/aws_instance: Fix issue with assuming Placement and disableApiTermination instance attributes exist when managing a Snowball Edge device ([#19256](https://github.com/hashicorp/terraform-provider-aws/issues/19256))
* resource/aws_macie2_member: Correct type for `invitation_disable_email_notification` parameter ([#24304](https://github.com/hashicorp/terraform-provider-aws/issues/24304))
* resource/aws_sfn_state_machine: Prevent panic during resource update ([#24302](https://github.com/hashicorp/terraform-provider-aws/issues/24302))
* resource/aws_shield_protection_group: When updating resource tags, use the `protection_group_arn` parameter instead of `arn`. ([#24296](https://github.com/hashicorp/terraform-provider-aws/issues/24296))

## 4.10.0 (April 14, 2022)

FEATURES:

* **New Data Source:** `aws_iam_saml_provider` ([#10498](https://github.com/hashicorp/terraform-provider-aws/issues/10498))
* **New Data Source:** `aws_nat_gateways` ([#24190](https://github.com/hashicorp/terraform-provider-aws/issues/24190))
* **New Resource:** `aws_datasync_location_fsx_openzfs_file_system` ([#24200](https://github.com/hashicorp/terraform-provider-aws/issues/24200))
* **New Resource:** `aws_elasticache_user_group_association` ([#24204](https://github.com/hashicorp/terraform-provider-aws/issues/24204))
* **New Resource:** `aws_qldb_stream` ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))

ENHANCEMENTS:

* data-source/aws_qldb_ledger: Add `kms_key` and `tags` attributes ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))
* resource/aws_athena_database: Add `properties` argument. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))
* resource/aws_athena_database: Add import support. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))
* resource/aws_grafana_workspace: Add plan time validations for `authentication_providers`, `authentication_providers`, `authentication_providers`. ([#24170](https://github.com/hashicorp/terraform-provider-aws/issues/24170))
* resource/aws_qldb_ledger: Add `kms_key` argument ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))

BUG FIXES:

* resource/aws_athena_database: Add drift detection for `comment`. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))

## 4.9.0 (April 07, 2022)

NOTES:


FEATURES:

* **New Data Source:** `aws_eks_addon_version` ([#23157](https://github.com/hashicorp/terraform-provider-aws/issues/23157))
* **New Data Source:** `aws_lambda_function_url` ([#24053](https://github.com/hashicorp/terraform-provider-aws/issues/24053))
* **New Data Source:** `aws_memorydb_acl` ([#23891](https://github.com/hashicorp/terraform-provider-aws/issues/23891))
* **New Data Source:** `aws_memorydb_cluster` ([#23991](https://github.com/hashicorp/terraform-provider-aws/issues/23991))
* **New Data Source:** `aws_memorydb_snapshot` ([#23990](https://github.com/hashicorp/terraform-provider-aws/issues/23990))
* **New Data Source:** `aws_memorydb_user` ([#23890](https://github.com/hashicorp/terraform-provider-aws/issues/23890))
* **New Data Source:** `aws_opensearch_domain` ([#23902](https://github.com/hashicorp/terraform-provider-aws/issues/23902))
* **New Data Source:** `aws_ssm_maintenance_windows` ([#24011](https://github.com/hashicorp/terraform-provider-aws/issues/24011))
* **New Resource:** `aws_db_instance_automated_backups_replication` ([#23759](https://github.com/hashicorp/terraform-provider-aws/issues/23759))
* **New Resource:** `aws_dynamodb_contributor_insights` ([#23947](https://github.com/hashicorp/terraform-provider-aws/issues/23947))
* **New Resource:** `aws_iot_indexing_configuration` ([#9929](https://github.com/hashicorp/terraform-provider-aws/issues/9929))
* **New Resource:** `aws_iot_logging_options` ([#13392](https://github.com/hashicorp/terraform-provider-aws/issues/13392))
* **New Resource:** `aws_iot_provisioning_template` ([#12108](https://github.com/hashicorp/terraform-provider-aws/issues/12108))
* **New Resource:** `aws_lambda_function_url` ([#24053](https://github.com/hashicorp/terraform-provider-aws/issues/24053))
* **New Resource:** `aws_opensearch_domain` ([#23902](https://github.com/hashicorp/terraform-provider-aws/issues/23902))
* **New Resource:** `aws_opensearch_domain_policy` ([#23902](https://github.com/hashicorp/terraform-provider-aws/issues/23902))
* **New Resource:** `aws_opensearch_domain_saml_options` ([#23902](https://github.com/hashicorp/terraform-provider-aws/issues/23902))
* **New Resource:** `aws_rds_cluster_activity_stream` ([#22097](https://github.com/hashicorp/terraform-provider-aws/issues/22097))

ENHANCEMENTS:

* provider: Add support for reading custom CA bundle setting from shared config files ([#24064](https://github.com/hashicorp/terraform-provider-aws/issues/24064))
* resource/aws_default_route_table: Add `core_network_arn` argument to the `route` configuration block ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_dlm_lifecycle_policy: Add `policy_details.schedule.create_rule.cron_expression`, `policy_details.schedule.retain_rule.interval`, `policy_details.schedule.retain_rule.interval_unit`, `policy_details.policy_type`, `policy_details.schedule.deprecate_rule`, `policy_details.parameters`, `policy_details.schedule.variable_tags`, `policy_details.schedule.fast_restore_rule`, `policy_details.schedule.share_rule`, `policy_details.resource_locations`, `policy_details.schedule.create_rule.location`, `policy_details.action` and `policy_details.event_source` arguments ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_dlm_lifecycle_policy: Add plan time validations for `policy_details.resource_types` and `description` arguments ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_dlm_lifecycle_policy: Make `policy_details.resource_types`, `policy_details.schedule`, `policy_details.target_tags`, `policy_details.schedule.retain_rule` and `policy_details.schedule.create_rule.interval` arguments optional ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_fms_policy: Retry when `InternalErrorException` errors are returned from the AWS API ([#23952](https://github.com/hashicorp/terraform-provider-aws/issues/23952))
* resource/aws_fsx_ontap_file_system: Support updating `storage_capacity`, `throughput_capacity`, and `disk_iops_configuration`. ([#24002](https://github.com/hashicorp/terraform-provider-aws/issues/24002))
* resource/aws_mq_broker: Add validation to `broker_name` and `security_groups` arguments ([#18088](https://github.com/hashicorp/terraform-provider-aws/issues/18088))
* resource/aws_vpn_connection: Add `core_network_arn` and `core_network_attachment_arn` attributes ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_xray_group: Add `insights_configuration` argument ([#24028](https://github.com/hashicorp/terraform-provider-aws/issues/24028))

BUG FIXES:

* resource/aws_lb: Fix attribute key not recognized issue preventing creation in ISO-B regions ([#23972](https://github.com/hashicorp/terraform-provider-aws/issues/23972))

## 4.8.0 (March 25, 2022)

FEATURES:

* **New Data Source:** `aws_mskconnect_connector` ([#23792](https://github.com/hashicorp/terraform-provider-aws/issues/23544))
* **New Resource:** `aws_mskconnect_connector` ([#23765](https://github.com/hashicorp/terraform-provider-aws/issues/23544))

ENHANCEMENTS:

* resource/aws_mwaa_environment: Add `schedulers` argument ([#21941](https://github.com/hashicorp/terraform-provider-aws/issues/21941))

BUG FIXES:


## 4.7.0 (March 24, 2022)

FEATURES:

* **New Data Source:** `aws_cloudwatch_event_bus` ([#23792](https://github.com/hashicorp/terraform-provider-aws/issues/23792))
* **New Data Source:** `aws_imagebuilder_image_pipelines` ([#23741](https://github.com/hashicorp/terraform-provider-aws/issues/23741))
* **New Data Source:** `aws_memorydb_parameter_group` ([#23814](https://github.com/hashicorp/terraform-provider-aws/issues/23814))
* **New Data Source:** `aws_route53_traffic_policy_document` ([#23602](https://github.com/hashicorp/terraform-provider-aws/issues/23602))
* **New Resource:** `aws_cognito_user_in_group` ([#23765](https://github.com/hashicorp/terraform-provider-aws/issues/23765))
* **New Resource:** `aws_keyspaces_keyspace` ([#23770](https://github.com/hashicorp/terraform-provider-aws/issues/23770))
* **New Resource:** `aws_route53_traffic_policy` ([#23602](https://github.com/hashicorp/terraform-provider-aws/issues/23602))
* **New Resource:** `aws_route53_traffic_policy_instance` ([#23602](https://github.com/hashicorp/terraform-provider-aws/issues/23602))

ENHANCEMENTS:

* resource/aws_athena_workgroup: Add `acl_configuration` and `expected_bucket_owner` arguments to the `configuration.result_configuration` block ([#23748](https://github.com/hashicorp/terraform-provider-aws/issues/23748))
* resource/aws_glue_schema: Update documentation to include [Protobuf data format support](https://aws.amazon.com/about-aws/whats-new/2022/02/aws-glue-schema-registry-protocol-buffers) ([#23815](https://github.com/hashicorp/terraform-provider-aws/issues/23815))
* resource/aws_instance: Add `user_data_replace_on_change` attribute ([#23604](https://github.com/hashicorp/terraform-provider-aws/issues/23604))

BUG FIXES:


## 4.6.0 (March 18, 2022)

FEATURES:

* **New Data Source:** `aws_networkmanager_connection` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_connections` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_device` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_devices` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_global_network` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_global_networks` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_link` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_links` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_site` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Data Source:** `aws_networkmanager_sites` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_gamelift_game_server_group` ([#23606](https://github.com/hashicorp/terraform-provider-aws/issues/23606))
* **New Resource:** `aws_networkmanager_connection` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_customer_gateway_association` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_device` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_global_network` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_link` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_link_association` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_site` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_transit_gateway_connect_peer_association` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_networkmanager_transit_gateway_registration` ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* **New Resource:** `aws_vpc_endpoint_security_group_association` ([#13737](https://github.com/hashicorp/terraform-provider-aws/issues/13737))

ENHANCEMENTS:

* data-source/aws_ec2_transit_gateway_connect_peer: Add `arn` attribute ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* resource/aws_athena_database: Add `acl_configuration` and `expected_bucket_owner` arguments ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_athena_database: Add `comment` argument to support database descriptions ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_athena_database: Do not recreate the resource if `bucket` changes ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_ec2_transit_gateway_connect_peer: Add `arn` attribute ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* resource/aws_quicksight_user: Allow custom values for `namespace` ([#23607](https://github.com/hashicorp/terraform-provider-aws/issues/23607))

BUG FIXES:

* resource/aws_amplify_app: Allow `repository` to be updated in-place ([#23517](https://github.com/hashicorp/terraform-provider-aws/issues/23517))
* resource/aws_athena_database: Remove from state on resource Read if deleted outside of Terraform ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))

## 4.5.0 (March 11, 2022)

ENHANCEMENTS:

* resource/aws_account_alternate_contact: Add configurable timeouts ([#23516](https://github.com/hashicorp/terraform-provider-aws/issues/23516))
x-amz-bucket-object-lock-token for enabling replication on object lock enabled
buckets or enabling object lock on an existing bucket. ([#23624](https://github.com/hashicorp/terraform-provider-aws/issues/23624))
* resource/aws_servicecatalog_budget_resource_association: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_constraint: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_organizations_access: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_portfolio: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_portfolio_share: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_principal_portfolio_association: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_product: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_product_portfolio_association: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_provisioned_product: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_provisioning_artifact: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_service_action: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_tag_option: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_servicecatalog_tag_option_resource_association: Add configurable timeouts ([#23518](https://github.com/hashicorp/terraform-provider-aws/issues/23518))
* resource/aws_synthetics_canary: Add optional `environment_variables` to `run_config`. ([#23574](https://github.com/hashicorp/terraform-provider-aws/issues/23574))

BUG FIXES:

* resource/aws_account_alternate_contact: Improve eventual consistency handling to avoid "no resource found" on updates ([#23516](https://github.com/hashicorp/terraform-provider-aws/issues/23516))

## 4.4.0 (March 04, 2022)

FEATURES:

* **New Data Source:** `aws_connect_queue` ([#22768](https://github.com/hashicorp/terraform-provider-aws/issues/22768))
* **New Data Source:** `aws_ec2_serial_console_access` ([#23443](https://github.com/hashicorp/terraform-provider-aws/issues/23443))
* **New Data Source:** `aws_ec2_transit_gateway_connect` ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* **New Data Source:** `aws_ec2_transit_gateway_connect_peer` ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* **New Resource:** `aws_apprunner_vpc_connector` ([#23173](https://github.com/hashicorp/terraform-provider-aws/issues/23173))
* **New Resource:** `aws_connect_routing_profile` ([#22813](https://github.com/hashicorp/terraform-provider-aws/issues/22813))
* **New Resource:** `aws_connect_user_hierarchy_structure` ([#22836](https://github.com/hashicorp/terraform-provider-aws/issues/22836))
* **New Resource:** `aws_ec2_network_insights_path` ([#23330](https://github.com/hashicorp/terraform-provider-aws/issues/23330))
* **New Resource:** `aws_ec2_serial_console_access` ([#23443](https://github.com/hashicorp/terraform-provider-aws/issues/23443))
* **New Resource:** `aws_ec2_transit_gateway_connect` ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* **New Resource:** `aws_ec2_transit_gateway_connect_peer` ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* **New Resource:** `aws_grafana_license_association` ([#23401](https://github.com/hashicorp/terraform-provider-aws/issues/23401))
* **New Resource:** `aws_route53domains_registered_domain` ([#12711](https://github.com/hashicorp/terraform-provider-aws/issues/12711))

ENHANCEMENTS:

* data-source/aws_ec2_transit_gateway: Add `transit_gateway_cidr_blocks` attribute ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* resource/aws_ec2_transit_gateway: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* resource/aws_ec2_transit_gateway: Add `transit_gateway_cidr_blocks` argument ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* resource/aws_glue_job: Add support for [streaming jobs](https://docs.aws.amazon.com/glue/latest/dg/add-job-streaming.html) by removing the default value for the `timeout` argument and marking it as Computed ([#23275](https://github.com/hashicorp/terraform-provider-aws/issues/23275))

BUG FIXES:

* resource/aws_opsworks_application: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_custom_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_ecs_cluster_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_ganglia_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_haproxy_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_instance: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_java_app_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_memcached_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_mysql_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_nodejs_app_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_php_app_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_rails_app_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_rds_db_instance: Correctly remove from state in certain deletion situations ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_stack: Fix error reported on successful deletion, lack of eventual consistency wait ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_static_web_layer: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_opsworks_user_profile: Fix error reported on successful deletion ([#23397](https://github.com/hashicorp/terraform-provider-aws/issues/23397))
* resource/aws_synthetics_canary: Retry canary creation if it fails because of IAM propagation. ([#23394](https://github.com/hashicorp/terraform-provider-aws/issues/23394))

## 4.3.0 (February 28, 2022)

NOTES:

* resource/aws_internet_gateway: Set `vpc_id` as Computed to prevent drift when the `aws_internet_gateway_attachment` resource is used ([#16386](https://github.com/hashicorp/terraform-provider-aws/issues/16386))

FEATURES:

* **New Data Source:** `aws_ec2_transit_gateway_multicast_domain` ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* **New Data Source:** `aws_ec2_transit_gateway_vpc_attachments` ([#12409](https://github.com/hashicorp/terraform-provider-aws/issues/12409))
* **New Resource:** `aws_ec2_transit_gateway_multicast_domain` ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* **New Resource:** `aws_ec2_transit_gateway_multicast_domain_association` ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* **New Resource:** `aws_ec2_transit_gateway_multicast_group_member` ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* **New Resource:** `aws_ec2_transit_gateway_multicast_group_source` ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* **New Resource:** `aws_internet_gateway_attachment` ([#16386](https://github.com/hashicorp/terraform-provider-aws/issues/16386))
* **New Resource:** `aws_opsworks_ecs_cluster_layer` ([#12495](https://github.com/hashicorp/terraform-provider-aws/issues/12495))
* **New Resource:** `aws_vpc_endpoint_policy` ([#17039](https://github.com/hashicorp/terraform-provider-aws/issues/17039))

ENHANCEMENTS:

* data-source/aws_ec2_transit_gateway: Add `multicast_support` attribute ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* provider: Improves error message when `Profile` and static credential environment variables are set. ([#23388](https://github.com/hashicorp/terraform-provider-aws/issues/23388))
* provider: Makes `region` an optional parameter to allow sourcing from shared config files and IMDS ([#23384](https://github.com/hashicorp/terraform-provider-aws/issues/23384))
* provider: Retrieves region from IMDS when credentials retrieved from IMDS. ([#23388](https://github.com/hashicorp/terraform-provider-aws/issues/23388))
* resource/aws_ec2_fleet: Add `context` argument ([#23304](https://github.com/hashicorp/terraform-provider-aws/issues/23304))
* resource/aws_ec2_transit_gateway: Add `multicast_support` argument ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* resource/aws_instance: Prevent double base64 encoding of `user_data` and `user_data_base64` on update ([#23362](https://github.com/hashicorp/terraform-provider-aws/issues/23362))
* resource/aws_servicecatalog_provisioned_product: Add `outputs` attribute ([#23270](https://github.com/hashicorp/terraform-provider-aws/issues/23270))

BUG FIXES:

* provider: Validates names of named profiles before use. ([#23388](https://github.com/hashicorp/terraform-provider-aws/issues/23388))

## 4.2.0 (February 18, 2022)

FEATURES:

* **New Data Source:** `aws_grafana_workspace` ([#22874](https://github.com/hashicorp/terraform-provider-aws/issues/22874))
* **New Data Source:** `aws_iam_openid_connect_provider` ([#23240](https://github.com/hashicorp/terraform-provider-aws/issues/23240))
* **New Data Source:** `aws_ssm_instances` ([#23162](https://github.com/hashicorp/terraform-provider-aws/issues/23162))
* **New Resource:** `aws_cloudtrail_event_data_store` ([#22490](https://github.com/hashicorp/terraform-provider-aws/issues/22490))
* **New Resource:** `aws_grafana_workspace` ([#22874](https://github.com/hashicorp/terraform-provider-aws/issues/22874))

ENHANCEMENTS:

* provider: Add `custom_ca_bundle` argument ([#23279](https://github.com/hashicorp/terraform-provider-aws/issues/23279))
* provider: Add `sts_region` argument ([#23212](https://github.com/hashicorp/terraform-provider-aws/issues/23212))
* provider: Expands environment variables in file paths in provider configuration. ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))
* provider: Updates list of valid AWS regions ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))
* resource/aws_fms_policy: Add `delete_unused_fm_managed_resources` argument ([#21295](https://github.com/hashicorp/terraform-provider-aws/issues/21295))
* resource/aws_fms_policy: Add `tags` argument and `tags_all` attribute to support resource tagging ([#21299](https://github.com/hashicorp/terraform-provider-aws/issues/21299))
* resource/aws_instance: Allow updates to `user_data` and `user_data_base64` without forcing resource replacement ([#18043](https://github.com/hashicorp/terraform-provider-aws/issues/18043))

BUG FIXES:

* provider: Credentials with expiry, such as assuming a role, would not renew. ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))
* provider: Setting a custom CA bundle caused the provider to fail. ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))

## 4.1.0 (February 15, 2022)

FEATURES:

* **New Data Source:** `aws_backup_framework` ([#23193](https://github.com/hashicorp/terraform-provider-aws/issues/23193))
* **New Data Source:** `aws_backup_report_plan` ([#23146](https://github.com/hashicorp/terraform-provider-aws/issues/23146))
* **New Data Source:** `aws_imagebuilder_container_recipe` ([#23040](https://github.com/hashicorp/terraform-provider-aws/issues/23040))
* **New Data Source:** `aws_imagebuilder_container_recipes` ([#23134](https://github.com/hashicorp/terraform-provider-aws/issues/23134))
* **New Data Source:** `aws_service` ([#16640](https://github.com/hashicorp/terraform-provider-aws/issues/16640))
* **New Resource:** `aws_backup_framework` ([#23175](https://github.com/hashicorp/terraform-provider-aws/issues/23175))
* **New Resource:** `aws_backup_report_plan` ([#23098](https://github.com/hashicorp/terraform-provider-aws/issues/23098))
* **New Resource:** `aws_gamelift_script` ([#11560](https://github.com/hashicorp/terraform-provider-aws/issues/11560))
* **New Resource:** `aws_iam_service_specific_credential` ([#16185](https://github.com/hashicorp/terraform-provider-aws/issues/16185))
* **New Resource:** `aws_iam_signing_certificate` ([#23161](https://github.com/hashicorp/terraform-provider-aws/issues/23161))
* **New Resource:** `aws_iam_virtual_mfa_device` ([#23113](https://github.com/hashicorp/terraform-provider-aws/issues/23113))
* **New Resource:** `aws_imagebuilder_container_recipe` ([#22965](https://github.com/hashicorp/terraform-provider-aws/issues/22965))

ENHANCEMENTS:

`audio.codec`,`audio.sample_rate`, `audio_codec_options.bit_depth`, `audio_codec_options.bit_order`,
`audio_codec_options.profile`, `audio_codec_options.signed`, `audio_codec_options.signed`,
`container`, `thumbnails.aspect_ratio`, `thumbnails.format`, `thumbnails.padding_policy`, `thumbnails.sizing_policy`,
`type`, `video.aspect_ratio`, `video.codec`, `video.display_aspect_ratio`, `video.fixed_gop`, `video.frame_rate`,   `video.max_frame_rate`,  `video.padding_policy`, `video.sizing_policy`, `video_watermarks.horizontal_align`,
`video_watermarks.id`, `video_watermarks.sizing_policy`, `video_watermarks.target`, `video_watermarks.vertical_align` ([#13974](https://github.com/hashicorp/terraform-provider-aws/issues/13974))
* resource/aws_gamelift_build: Add `object_version` argument to `storage_location` block. ([#22966](https://github.com/hashicorp/terraform-provider-aws/issues/22966))
* resource/aws_gamelift_build: Add import support ([#22966](https://github.com/hashicorp/terraform-provider-aws/issues/22966))
* resource/aws_gamelift_fleet: Add `certificate_configuration` argument ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Add import support ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Add plan time validation to `ec2_instance_type` ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Adds `script_arn` attribute. ([#11560](https://github.com/hashicorp/terraform-provider-aws/issues/11560))
* resource/aws_gamelift_fleet: Adds `script_id` argument. ([#11560](https://github.com/hashicorp/terraform-provider-aws/issues/11560))
* resource/aws_glue_catalog_database: Add support `create_table_default_permission` argument ([#22964](https://github.com/hashicorp/terraform-provider-aws/issues/22964))
* resource/aws_glue_trigger: Add `event_batching_condition` argument. ([#22963](https://github.com/hashicorp/terraform-provider-aws/issues/22963))
* resource/aws_prometheus_workspace: Add `tags` argument and `tags_all` attribute to support resource tagging ([#23202](https://github.com/hashicorp/terraform-provider-aws/issues/23202))

BUG FIXES:

* provider: Support `ap-northeast-3`, `ap-southeast-3` and `us-iso-west-1` as valid AWS Regions ([#23191](https://github.com/hashicorp/terraform-provider-aws/issues/23191))
* provider: Use AWS HTTP client which allows IMDS authentication in container environments and custom RootCAs in ISO regions ([#23191](https://github.com/hashicorp/terraform-provider-aws/issues/23191))
* resource/aws_vpn_connection: Add support for `ipsec.1-aes256` connection type ([#23127](https://github.com/hashicorp/terraform-provider-aws/issues/23127))

## 4.0.0 (February 10, 2022)

BREAKING CHANGES:

* resource/aws_batch_compute_environment: No `compute_resources` configuration block can be specified when `type` is `UNMANAGED` ([#22805](https://github.com/hashicorp/terraform-provider-aws/issues/22805))
* resource/aws_default_network_acl: These arguments can no longer be set to `""`: `egress.*.cidr_block`, `egress.*.ipv6_cidr_block`, `ingress.*.cidr_block`, or `ingress.*.ipv6_cidr_block` ([#22928](https://github.com/hashicorp/terraform-provider-aws/issues/22928))
* resource/aws_default_route_table: These arguments can no longer be set to `""`: `route.*.cidr_block`, `route.*.ipv6_cidr_block` ([#22931](https://github.com/hashicorp/terraform-provider-aws/issues/22931))
* resource/aws_default_vpc: `ipv6_cidr_block` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))
* resource/aws_efs_mount_target: `ip_address` can no longer be set to `""`; instead, remove or set to `null` ([#22954](https://github.com/hashicorp/terraform-provider-aws/issues/22954))
* resource/aws_fsx_ontap_storage_virtual_machine: Remove deprecated `active_directory_configuration.0.self_managed_active_directory_configuration.0.organizational_unit_distinguidshed_name`, migrating value to `active_directory_configuration.0.self_managed_active_directory_configuration.0.organizational_unit_distinguished_name` ([#22915](https://github.com/hashicorp/terraform-provider-aws/issues/22915))
* resource/aws_instance: `private_ip` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))
* resource/aws_lb_target_group: For `protocol = "TCP"`, `stickiness` can no longer be type set to `lb_cookie` even when `enabled = false`; instead use type `source_ip` ([#22996](https://github.com/hashicorp/terraform-provider-aws/issues/22996))

NOTES:

* data-source/aws_ec2_coip_pools: The type of the `pool_ids` attribute has changed from Set to List. If no COIP pools match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateway_route_tables: The type of the `ids` attribute has changed from Set to List. If no local gateway route tables match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateway_virtual_interface_groups: The type of the `ids` and `local_gateway_virtual_interface_ids` attributes has changed from Set to List. If no local gateway virtual interface groups match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateways: The type of the `ids` attribute has changed from Set to List. If no local gateways match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_transit_gateway_route_tables: The type of the `ids` attribute has changed from Set to List. If no transit gateway route tables match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_efs_access_points: The type of the `ids` and `arns` attributes has changed from Set to List. If no access points match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_inspector_rules_packages: If no rules packages match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_instances: If no instances match the specified criteria an empty list is returned (previously an error was raised) ([#5055](https://github.com/hashicorp/terraform-provider-aws/issues/5055))
* data-source/aws_ip_ranges: If no ranges match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ssoadmin_instances: The type of the `identity_store_ids` and `arns` attributes has changed from Set to List. If no instances match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* provider: The `assume_role.duration_seconds` argument has been deprecated. All configurations using `assume_role.duration_seconds` should be updated to use the new `assume_role.duration` argument instead. ([#23077](https://github.com/hashicorp/terraform-provider-aws/issues/23077))
* resource/aws_acmpca_certificate_authority: The `status` attribute has been deprecated. Use the `enabled` attribute instead. ([#22878](https://github.com/hashicorp/terraform-provider-aws/issues/22878))
* resource/aws_budgets_budget: The `cost_filters` attribute has been deprecated. Use the `cost_filter` attribute instead. ([#22888](https://github.com/hashicorp/terraform-provider-aws/issues/22888))
* resource/aws_customer_gateway: `ip_address` can no longer be set to `""` ([#22926](https://github.com/hashicorp/terraform-provider-aws/issues/22926))
* resource/aws_default_subnet: If no default subnet exists in the specified Availability Zone one is now created. The `force_destroy` destroy argument has been added (defaults to `false`). Setting this argument to `true` deletes the default subnet on `terraform destroy` ([#22253](https://github.com/hashicorp/terraform-provider-aws/issues/22253))
* resource/aws_default_vpc: If no default VPC exists in the current AWS Region one is now created. The `force_destroy` destroy argument has been added (defaults to `false`). Setting this argument to `true` deletes the default VPC on `terraform destroy` ([#22253](https://github.com/hashicorp/terraform-provider-aws/issues/22253))
* resource/aws_ec2_client_vpn_endpoint: The `status` attribute has been deprecated ([#22887](https://github.com/hashicorp/terraform-provider-aws/issues/22887))
* resource/aws_ec2_client_vpn_endpoint: The type of the `dns_servers` argument has changed from Set to List ([#22889](https://github.com/hashicorp/terraform-provider-aws/issues/22889))
* resource/aws_ec2_client_vpn_network_association: The `security_groups` argument has been deprecated. Use the `security_group_ids` argument of the `aws_ec2_client_vpn_endpoint` resource instead ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* resource/aws_ec2_client_vpn_network_association: The `status` attribute has been deprecated ([#22887](https://github.com/hashicorp/terraform-provider-aws/issues/22887))
* resource/aws_ec2_client_vpn_route: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))

FEATURES:

* **New Data Source:** `aws_cloudfront_realtime_log_config` ([#22620](https://github.com/hashicorp/terraform-provider-aws/issues/22620))
* **New Data Source:** `aws_ec2_client_vpn_endpoint` ([#14218](https://github.com/hashicorp/terraform-provider-aws/issues/14218))
* **New Data Source:** `aws_eips` ([#7537](https://github.com/hashicorp/terraform-provider-aws/issues/7537))
* **New Data Source:** `aws_s3_object` ([#22850](https://github.com/hashicorp/terraform-provider-aws/issues/22850))
* **New Data Source:** `aws_s3_objects` ([#22850](https://github.com/hashicorp/terraform-provider-aws/issues/22850))
* **New Resource:** `aws_cognito_user` ([#19919](https://github.com/hashicorp/terraform-provider-aws/issues/19919))
* **New Resource:** `aws_dataexchange_revision` ([#22933](https://github.com/hashicorp/terraform-provider-aws/issues/22933))
* **New Resource:** `aws_network_acl_association` ([#18807](https://github.com/hashicorp/terraform-provider-aws/issues/18807))
* **New Resource:** `aws_s3_bucket_accelerate_configuration` ([#22617](https://github.com/hashicorp/terraform-provider-aws/issues/22617))
* **New Resource:** `aws_s3_bucket_acl` ([#22853](https://github.com/hashicorp/terraform-provider-aws/issues/22853))
* **New Resource:** `aws_s3_bucket_cors_configuration` ([#12141](https://github.com/hashicorp/terraform-provider-aws/issues/12141))
* **New Resource:** `aws_s3_bucket_lifecycle_configuration` ([#22579](https://github.com/hashicorp/terraform-provider-aws/issues/22579))
* **New Resource:** `aws_s3_bucket_logging` ([#22608](https://github.com/hashicorp/terraform-provider-aws/issues/22608))
* **New Resource:** `aws_s3_bucket_object_lock_configuration` ([#22644](https://github.com/hashicorp/terraform-provider-aws/issues/22644))
* **New Resource:** `aws_s3_bucket_request_payment_configuration` ([#22649](https://github.com/hashicorp/terraform-provider-aws/issues/22649))
* **New Resource:** `aws_s3_bucket_server_side_encryption_configuration` ([#22609](https://github.com/hashicorp/terraform-provider-aws/issues/22609))
* **New Resource:** `aws_s3_bucket_versioning` ([#5132](https://github.com/hashicorp/terraform-provider-aws/issues/5132))
* **New Resource:** `aws_s3_bucket_website_configuration` ([#22648](https://github.com/hashicorp/terraform-provider-aws/issues/22648))
* **New Resource:** `aws_s3_object` ([#22850](https://github.com/hashicorp/terraform-provider-aws/issues/22850))

ENHANCEMENTS:

* data-source/aws_ec2_client_vpn_endpoint: Add `security_group_ids` and `vpc_id` attributes ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* provider: Add `duration` argument to the `assume_role` configuration block ([#23077](https://github.com/hashicorp/terraform-provider-aws/issues/23077))
* provider: Add `ec2_metadata_service_endpoint`, `ec2_metadata_service_endpoint_mode`, `use_dualstack_endpoint`, `use_fips_endpoint` arguments ([#22804](https://github.com/hashicorp/terraform-provider-aws/issues/22804))
* provider: Add environment variables `TF_AWS_DYNAMODB_ENDPOINT`, `TF_AWS_IAM_ENDPOINT`, `TF_AWS_S3_ENDPOINT`, and `TF_AWS_STS_ENDPOINT`. ([#23052](https://github.com/hashicorp/terraform-provider-aws/issues/23052))
* provider: Add support for `shared_config_file` parameter ([#20587](https://github.com/hashicorp/terraform-provider-aws/issues/20587))
* provider: Add support for `shared_credentials_files` parameter and deprecates `shared_credentials_file` ([#23080](https://github.com/hashicorp/terraform-provider-aws/issues/23080))
* provider: Adds `s3_use_path_style` parameter and deprecates `s3_force_path_style`. ([#23055](https://github.com/hashicorp/terraform-provider-aws/issues/23055))
* provider: Changes `shared_config_file` parameter to `shared_config_files` ([#23080](https://github.com/hashicorp/terraform-provider-aws/issues/23080))
* provider: Updates AWS authentication to use AWS SDK for Go v2 <https://aws.github.io/aws-sdk-go-v2/docs/> ([#20587](https://github.com/hashicorp/terraform-provider-aws/issues/20587))
* resource/aws_ec2_client_vpn_authorization_rule: Configurable Create and Delete timeouts ([#20688](https://github.com/hashicorp/terraform-provider-aws/issues/20688))
* resource/aws_ec2_client_vpn_endpoint: Add `client_connect_options` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `client_login_banner_options` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `security_group_ids` and `vpc_id` arguments ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* resource/aws_ec2_client_vpn_endpoint: Add `session_timeout_hours` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `vpn_port` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_network_association: Configurable Create and Delete timeouts ([#20689](https://github.com/hashicorp/terraform-provider-aws/issues/20689))
* resource/aws_fsx_lustre_file_system: Add `log_configuration` argument. ([#22935](https://github.com/hashicorp/terraform-provider-aws/issues/22935))
* resource/aws_fsx_ontap_file_system: Reduce the minimum valid value of the `throughput_capacity` argument to `128` (128 MB/s) ([#22898](https://github.com/hashicorp/terraform-provider-aws/issues/22898))
* resource/aws_glue_partition_index: Add support for custom timeouts. ([#22941](https://github.com/hashicorp/terraform-provider-aws/issues/22941))
* resource/aws_mq_broker: `auto_minor_version_upgrade` and `host_instance_type` can be changed without recreating broker ([#20661](https://github.com/hashicorp/terraform-provider-aws/issues/20661))
* resource/aws_vpn_connection: Add the ability to revert changes to unconfigured tunnel options made outside of Terraform to their [documented default values](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html) ([#17031](https://github.com/hashicorp/terraform-provider-aws/issues/17031))
* resource/aws_vpn_connection: Mark `customer_gateway_configuration` as [`Sensitive`](https://www.terraform.io/plugin/sdkv2/best-practices/sensitive-state#using-the-sensitive-flag) ([#15806](https://github.com/hashicorp/terraform-provider-aws/issues/15806))

BUG FIXES:

* resource/aws_ec2_client_vpn_authorization_rule: Don't raise an error when `InvalidClientVpnEndpointId.NotFound` is returned during refresh ([#20688](https://github.com/hashicorp/terraform-provider-aws/issues/20688))
* resource/aws_ec2_client_vpn_endpoint: `connection_log_options.cloudwatch_log_stream` argument is Computed, preventing spurious resource diffs ([#22891](https://github.com/hashicorp/terraform-provider-aws/issues/22891))

## 3.75.2 (May 20, 2022)

ENHANCEMENTS:


## 3.75.1 (March 24, 2022)

BUG FIXES:


## 3.75.0 (March 18, 2022)

NOTES:


FEATURES:

* **New Resource:** `aws_s3_bucket_accelerate_configuration` ([#23471](https://github.com/hashicorp/terraform-provider-aws/issues/23471))
* **New Resource:** `aws_s3_bucket_acl` ([#23419](https://github.com/hashicorp/terraform-provider-aws/issues/23419))
* **New Resource:** `aws_s3_bucket_cors_configuration` ([#23434](https://github.com/hashicorp/terraform-provider-aws/issues/23434))
* **New Resource:** `aws_s3_bucket_lifecycle_configuration` ([#23445](https://github.com/hashicorp/terraform-provider-aws/issues/23445))
* **New Resource:** `aws_s3_bucket_logging` ([#23430](https://github.com/hashicorp/terraform-provider-aws/issues/23430))
* **New Resource:** `aws_s3_bucket_object_lock_configuration` ([#23449](https://github.com/hashicorp/terraform-provider-aws/issues/23449))
* **New Resource:** `aws_s3_bucket_request_payment_configuration` ([#23473](https://github.com/hashicorp/terraform-provider-aws/issues/23473))
* **New Resource:** `aws_s3_bucket_server_side_encryption_configuration` ([#23476](https://github.com/hashicorp/terraform-provider-aws/issues/23476))
* **New Resource:** `aws_s3_bucket_versioning` ([#23432](https://github.com/hashicorp/terraform-provider-aws/issues/23432))
* **New Resource:** `aws_s3_bucket_website_configuration` ([#23435](https://github.com/hashicorp/terraform-provider-aws/issues/23435))

ENHANCEMENTS:

x-amz-bucket-object-lock-token for enabling replication on object lock enabled
buckets or enabling object lock on an existing bucket. ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))

BUG FIXES:


## 3.74.3 (February 17, 2022)

BUG FIXES:


## 3.74.2 (February 11, 2022)

BUG FIXES:

* resource/aws_vpn_connection: Add support for `ipsec.1-aes256` connection type ([#23127](https://github.com/hashicorp/terraform-provider-aws/issues/23127))

## 3.74.1 (February 7, 2022)

BUG FIXES:


## 3.74.0 (January 28, 2022)

FEATURES:

* **New Data Source:** `aws_api_gateway_export` ([#22731](https://github.com/hashicorp/terraform-provider-aws/issues/22731))
* **New Data Source:** `aws_api_gateway_sdk` ([#22731](https://github.com/hashicorp/terraform-provider-aws/issues/22731))
* **New Data Source:** `aws_apigatewayv2_export` ([#22732](https://github.com/hashicorp/terraform-provider-aws/issues/22732))
* **New Data Source:** `aws_connect_contact_flow_module` ([#22518](https://github.com/hashicorp/terraform-provider-aws/issues/22518))
* **New Data Source:** `aws_connect_prompt` ([#22636](https://github.com/hashicorp/terraform-provider-aws/issues/22636))
* **New Data Source:** `aws_connect_quick_connect` ([#22527](https://github.com/hashicorp/terraform-provider-aws/issues/22527))
* **New Data Source:** `aws_datapipeline_pipeline` ([#22597](https://github.com/hashicorp/terraform-provider-aws/issues/22597))
* **New Data Source:** `aws_datapipeline_pipeline_definition` ([#22597](https://github.com/hashicorp/terraform-provider-aws/issues/22597))
* **New Data Source:** `aws_imagebuilder_components` ([#21881](https://github.com/hashicorp/terraform-provider-aws/issues/21881))
* **New Data Source:** `aws_imagebuilder_distribution_configurations` ([#22733](https://github.com/hashicorp/terraform-provider-aws/issues/22733))
* **New Data Source:** `aws_imagebuilder_infrastructure_configurations` ([#22723](https://github.com/hashicorp/terraform-provider-aws/issues/22723))
* **New Resource:** `aws_connect_queue` ([#22566](https://github.com/hashicorp/terraform-provider-aws/issues/22566))
* **New Resource:** `aws_connect_security_profile` ([#22369](https://github.com/hashicorp/terraform-provider-aws/issues/22369))
* **New Resource:** `aws_dataexchange_data_set` ([#22697](https://github.com/hashicorp/terraform-provider-aws/issues/22697))
* **New Resource:** `aws_datapipeline_pipeline_definition` ([#22597](https://github.com/hashicorp/terraform-provider-aws/issues/22597))
* **New Resource:** `aws_devicefarm_test_grid_project` ([#22688](https://github.com/hashicorp/terraform-provider-aws/issues/22688))
* **New Resource:** `aws_ecs_cluster_capacity_providers` ([#22672](https://github.com/hashicorp/terraform-provider-aws/issues/22672))
* **New Resource:** `aws_sagemaker_project` ([#21534](https://github.com/hashicorp/terraform-provider-aws/issues/21534))

ENHANCEMENTS:


BUG FIXES:

* data-source/aws_lb: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* data-source/aws_lb: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* data-source/aws_lb_listener: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* data-source/aws_lb_target_group: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* data-source/aws_sqs_queue: Further refine tag error handling in ISO partitions ([#22780](https://github.com/hashicorp/terraform-provider-aws/issues/22780))
* resource/aws_instance: Prevent panic when reading the instance's block device mappings ([#22719](https://github.com/hashicorp/terraform-provider-aws/issues/22719))
* resource/aws_internet_gateway: No longer give up before the attachment timeout (4m) is exceeded (previously it was giving up after 20 not found checks). ([#22713](https://github.com/hashicorp/terraform-provider-aws/issues/22713))
* resource/aws_lb: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* resource/aws_lb_listener: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* resource/aws_lb_listener_rule: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* resource/aws_lb_target_group: Further refine tag error handling for ISO regions ([#22717](https://github.com/hashicorp/terraform-provider-aws/issues/22717))
* resource/aws_sns_topic: Further refine tag error handling in ISO partitions ([#22780](https://github.com/hashicorp/terraform-provider-aws/issues/22780))
* resource/aws_sqs_queue: Further refine tag error handling in ISO partitions ([#22780](https://github.com/hashicorp/terraform-provider-aws/issues/22780))

## 3.73.0 (January 21, 2022)

FEATURES:

* **New Data Source:** `aws_cloudfront_origin_access_identity` ([#22572](https://github.com/hashicorp/terraform-provider-aws/issues/22572))
* **New Data Source:** `aws_vpc_ipam_preview_next_cidr` ([#22643](https://github.com/hashicorp/terraform-provider-aws/issues/22643))
* **New Resource:** `aws_appsync_api_cache` ([#22578](https://github.com/hashicorp/terraform-provider-aws/issues/22578))
* **New Resource:** `aws_appsync_domain_name` ([#22487](https://github.com/hashicorp/terraform-provider-aws/issues/22487))
* **New Resource:** `aws_appsync_domain_name_api_association` ([#22487](https://github.com/hashicorp/terraform-provider-aws/issues/22487))
* **New Resource:** `aws_cloudsearch_domain` ([#17723](https://github.com/hashicorp/terraform-provider-aws/issues/17723))
* **New Resource:** `aws_cloudsearch_domain_service_access_policy` ([#17723](https://github.com/hashicorp/terraform-provider-aws/issues/17723))
* **New Resource:** `aws_detective_invitation_accepter` ([#22163](https://github.com/hashicorp/terraform-provider-aws/issues/22163))
* **New Resource:** `aws_detective_member` ([#22163](https://github.com/hashicorp/terraform-provider-aws/issues/22163))
* **New Resource:** `aws_fsx_data_repository_association` ([#22291](https://github.com/hashicorp/terraform-provider-aws/issues/22291))
* **New Resource:** `aws_lambda_invocation` ([#19488](https://github.com/hashicorp/terraform-provider-aws/issues/19488))


ENHANCEMENTS:

* resource/aws_fsx_lustre_file_system: Add `file_system_type_version` argument ([#22291](https://github.com/hashicorp/terraform-provider-aws/issues/22291))
* resource/aws_opsworks_custom_layer: Add plan time validation for `ebs_volume.type` and `custom_json`. ([#12433](https://github.com/hashicorp/terraform-provider-aws/issues/12433))
* resource/aws_opsworks_custom_layer: Add support for `cloudwatch_configuration` ([#12433](https://github.com/hashicorp/terraform-provider-aws/issues/12433))

BUG FIXES:

* resource/aws_load_balancer_policy: Suppress `policy_attribute` differences ([#21776](https://github.com/hashicorp/terraform-provider-aws/issues/21776))

## 3.72.0 (January 13, 2022)

FEATURES:

* **New Data Source:** `aws_cognito_user_pool_client` ([#22477](https://github.com/hashicorp/terraform-provider-aws/issues/22477))
* **New Resource:** `aws_cognito_identity_pool_provider_principal_tag` ([#22514](https://github.com/hashicorp/terraform-provider-aws/issues/22514))
* **New Resource:** `aws_connect_contact_flow_module` ([#22349](https://github.com/hashicorp/terraform-provider-aws/issues/22349))
* **New Resource:** `aws_connect_quick_connect` ([#22250](https://github.com/hashicorp/terraform-provider-aws/issues/22250))
* **New Resource:** `aws_devicefarm_instance_profile` ([#22458](https://github.com/hashicorp/terraform-provider-aws/issues/22458))
* **New Resource:** `aws_memorydb_snapshot` ([#22486](https://github.com/hashicorp/terraform-provider-aws/issues/22486))
* **New Resource:** `aws_shield_protection_health_check_association` ([#21993](https://github.com/hashicorp/terraform-provider-aws/issues/21993))

ENHANCEMENTS:

* data-source/aws_customer_gateway: Add `certificate_arn` attribute ([#22435](https://github.com/hashicorp/terraform-provider-aws/issues/22435))
* data-source/aws_elb_service_account: Add account ID for `ap-southeast-3` AWS Region ([#22453](https://github.com/hashicorp/terraform-provider-aws/issues/22453))
* data-source/aws_instance: Add the `instance_metadata_tags` attribute to the `metadata_options` configuration block ([#22463](https://github.com/hashicorp/terraform-provider-aws/issues/22463))
* data-source/aws_lb: Allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* data-source/aws_lb_listener: Allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* data-source/aws_lb_target_group: Allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* data-source/aws_sqs_queue: Allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22516](https://github.com/hashicorp/terraform-provider-aws/issues/22516))
* resource/aws_customer_gateway: Add `certificate_arn` argument ([#22435](https://github.com/hashicorp/terraform-provider-aws/issues/22435))
* resource/aws_glue_crawler: add `delta_target` argument. ([#22472](https://github.com/hashicorp/terraform-provider-aws/issues/22472))
* resource/aws_instance: Add the `instance_metadata_tags` argument to the `metadata_options` configuration block ([#22463](https://github.com/hashicorp/terraform-provider-aws/issues/22463))
* resource/aws_lb: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* resource/aws_lb_listener: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* resource/aws_lb_listener_rule: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* resource/aws_lb_target_group: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22551](https://github.com/hashicorp/terraform-provider-aws/issues/22551))
* resource/aws_sns_topic: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22511](https://github.com/hashicorp/terraform-provider-aws/issues/22511))
* resource/aws_sqs_queue: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#22516](https://github.com/hashicorp/terraform-provider-aws/issues/22516))
* resource/aws_vpn_connection: Add `vgw_telemetry.certificate_arn` attribute ([#19311](https://github.com/hashicorp/terraform-provider-aws/issues/19311))
* resource/aws_vpn_connection: `customer_gateway_id`, `transit_gateway_id` and `vpn_gateway_id` can be updated without recreating the resource ([#19311](https://github.com/hashicorp/terraform-provider-aws/issues/19311))
* resource/aws_vpn_connection: `tunnel1_preshared_key` and `tunnel2_preshared_key` can be updated without recreating the resource ([#19311](https://github.com/hashicorp/terraform-provider-aws/issues/19311))

BUG FIXES:

* resource/aws_fsx_lustre_file_system: Add missing values to `per_unit_storage_throughput` validation ([#22462](https://github.com/hashicorp/terraform-provider-aws/issues/22462))
* resource/aws_fsx_openzfs_file_system: Change `root_volume_configuration.copy_tags_to_snapshots` to ForceNew ([#22480](https://github.com/hashicorp/terraform-provider-aws/issues/22480))
* resource/aws_fsx_openzfs_file_system: Fix crash with nil `root_volume_configuration.nfs_exports` value ([#22480](https://github.com/hashicorp/terraform-provider-aws/issues/22480))
* resource/aws_memorydb_cluster: Correctly propagate configurable timeouts to waiters. ([#22489](https://github.com/hashicorp/terraform-provider-aws/issues/22489))
* resource/aws_sqs_queue: Don't timeout when a queue policy `Condition` value contains an empty array ([#22547](https://github.com/hashicorp/terraform-provider-aws/issues/22547))

## 3.71.0 (January 06, 2022)

FEATURES:

* **New Data Source:** `aws_batch_scheduling_policy` ([#22335](https://github.com/hashicorp/terraform-provider-aws/issues/22335))
* **New Data Source:** `aws_cognito_user_pool_clients` ([#22289](https://github.com/hashicorp/terraform-provider-aws/issues/22289))
* **New Data Source:** `aws_cognito_user_pool_signing_certificate` ([#22285](https://github.com/hashicorp/terraform-provider-aws/issues/22285))
* **New Data Source:** `aws_mskconnect_custom_plugin` ([#22333](https://github.com/hashicorp/terraform-provider-aws/issues/22333))
* **New Data Source:** `aws_mskconnect_worker_configuration` ([#22414](https://github.com/hashicorp/terraform-provider-aws/issues/22414))
* **New Data Source:** `aws_organizations_resource_tags` ([#22371](https://github.com/hashicorp/terraform-provider-aws/issues/22371))
* **New Data Source:** `aws_ses_active_receipt_rule_set` ([#22310](https://github.com/hashicorp/terraform-provider-aws/issues/22310))
* **New Data Source:** `aws_ses_domain_identity` ([#22321](https://github.com/hashicorp/terraform-provider-aws/issues/22321))
* **New Data Source:** `aws_ses_email_identity` ([#22321](https://github.com/hashicorp/terraform-provider-aws/issues/22321))
* **New Resource:** `aws_batch_scheduling_policy` ([#22262](https://github.com/hashicorp/terraform-provider-aws/issues/22262))
* **New Resource:** `aws_cloud9_environment_membership` ([#11857](https://github.com/hashicorp/terraform-provider-aws/issues/11857))
* **New Resource:** `aws_codebuild_resource_policy` ([#22196](https://github.com/hashicorp/terraform-provider-aws/issues/22196))
* **New Resource:** `aws_datasync_location_fsx_lustre_file_system` ([#22346](https://github.com/hashicorp/terraform-provider-aws/issues/22346))
* **New Resource:** `aws_datasync_location_hdfs` ([#22347](https://github.com/hashicorp/terraform-provider-aws/issues/22347))
* **New Resource:** `aws_devicefarm_device_pool` ([#21025](https://github.com/hashicorp/terraform-provider-aws/issues/21025))
* **New Resource:** `aws_devicefarm_network_profile` ([#22448](https://github.com/hashicorp/terraform-provider-aws/issues/22448))
* **New Resource:** `aws_devicefarm_upload` ([#22443](https://github.com/hashicorp/terraform-provider-aws/issues/22443))
* **New Resource:** `aws_fsx_openzfs_file_system` ([#22234](https://github.com/hashicorp/terraform-provider-aws/issues/22234))
* **New Resource:** `aws_fsx_openzfs_snapshot` ([#22234](https://github.com/hashicorp/terraform-provider-aws/issues/22234))
* **New Resource:** `aws_fsx_openzfs_volume` ([#22234](https://github.com/hashicorp/terraform-provider-aws/issues/22234))
* **New Resource:** `aws_memorydb_cluster` ([#22388](https://github.com/hashicorp/terraform-provider-aws/issues/22388))
* **New Resource:** `aws_memorydb_parameter_group` ([#22304](https://github.com/hashicorp/terraform-provider-aws/issues/22304))
* **New Resource:** `aws_memorydb_subnet_group` ([#22256](https://github.com/hashicorp/terraform-provider-aws/issues/22256))
* **New Resource:** `aws_memorydb_user` ([#22261](https://github.com/hashicorp/terraform-provider-aws/issues/22261))
* **New Resource:** `aws_mskconnect_custom_plugin` ([#22333](https://github.com/hashicorp/terraform-provider-aws/issues/22333))
* **New Resource:** `aws_mskconnect_worker_configuration` ([#22414](https://github.com/hashicorp/terraform-provider-aws/issues/22414))
* **New Resource:** `aws_sagemaker_device` ([#22427](https://github.com/hashicorp/terraform-provider-aws/issues/22427))
* **New Resource:** `aws_vpc_endpoint_connection_accepter` ([#19083](https://github.com/hashicorp/terraform-provider-aws/issues/19083))
* **New Resource:** `aws_vpc_ipam_organization_admin_account` ([#22394](https://github.com/hashicorp/terraform-provider-aws/issues/22394))

ENHANCEMENTS:

* data-source/aws_batch_job_queue: Add `scheduling_policy_arn` attribute ([#22348](https://github.com/hashicorp/terraform-provider-aws/issues/22348))
* data-source/aws_elb_hosted_zone_id: Add hosted zone ID for `ap-southeast-3` AWS Region ([#22295](https://github.com/hashicorp/terraform-provider-aws/issues/22295))
* provider: Add validation for the `duration`, `external_id` and `session_name` arguments in the `assume_role` configuration block ([#18085](https://github.com/hashicorp/terraform-provider-aws/issues/18085))
* resource/aws_batch_job_queue: Add `scheduling_policy_arn` attribute ([#22298](https://github.com/hashicorp/terraform-provider-aws/issues/22298))
* resource/aws_dax_cluster: Add `cluster_endpoint_encryption_type` argument ([#22396](https://github.com/hashicorp/terraform-provider-aws/issues/22396))
* resource/aws_dx_private_virtual_interface: Add `sitelink_enabled` argument ([#22350](https://github.com/hashicorp/terraform-provider-aws/issues/22350))
* resource/aws_dx_transit_virtual_interface: Add `sitelink_enabled` argument ([#22350](https://github.com/hashicorp/terraform-provider-aws/issues/22350))
* resource/aws_glue_trigger: Add `start_on_creation` argument ([#22439](https://github.com/hashicorp/terraform-provider-aws/issues/22439))
* resource/aws_timestreamwrite_table: Add `magnetic_store_write_properties` argument. ([#22363](https://github.com/hashicorp/terraform-provider-aws/issues/22363))

BUG FIXES:

* resource/aws_ec2_traffic_mirror_filter_rule: Prevent crash during resource read ([#22315](https://github.com/hashicorp/terraform-provider-aws/issues/22315))
* resource/aws_ssoadmin_managed_policy_attachment: Fix missing call to `ProvisionPermissionSet` after detaching the managed policy ([#21773](https://github.com/hashicorp/terraform-provider-aws/issues/21773))

## 3.70.0 (December 16, 2021)

NOTES:

* resource/aws_fsx_ontap_storage_virtual_machine: The `active_directory_configuration.self_managed_active_directory_configuration.organizational_unit_distinguidshed_name` attribute has been deprecated. All configurations using `active_directory_configuration.self_managed_active_directory_configuration.organizational_unit_distinguidshed_name` should be updated to use the new `active_directory_configuration.self_managed_active_directory_configuration.organizational_unit_distinguished_name` attribute instead ([#22246](https://github.com/hashicorp/terraform-provider-aws/issues/22246))

FEATURES:

* **New Data Source:** `aws_connect_bot_association` ([#21097](https://github.com/hashicorp/terraform-provider-aws/issues/21097))
* **New Data Source:** `aws_connect_hours_of_operation` ([#22207](https://github.com/hashicorp/terraform-provider-aws/issues/22207))
* **New Data Source:** `aws_connect_lambda_function_association` ([#21276](https://github.com/hashicorp/terraform-provider-aws/issues/21276))
* **New Resource:** `aws_connect_bot_association` ([#21097](https://github.com/hashicorp/terraform-provider-aws/issues/21097))
* **New Resource:** `aws_connect_hours_of_operation` ([#21934](https://github.com/hashicorp/terraform-provider-aws/issues/21934))
* **New Resource:** `aws_connect_lambda_function_association` ([#21276](https://github.com/hashicorp/terraform-provider-aws/issues/21276))
* **New Resource:** `aws_ecr_pull_through_cache_rule` ([#22172](https://github.com/hashicorp/terraform-provider-aws/issues/22172))
* **New Resource:** `aws_ecr_registry_scanning_configuration` ([#22179](https://github.com/hashicorp/terraform-provider-aws/issues/22179))
* **New Resource:** `aws_ecrpublic_repository_policy` ([#16901](https://github.com/hashicorp/terraform-provider-aws/issues/16901))

ENHANCEMENTS:

* resource/aws_fsx_backup: Add `volume_id` argument to support Amazon FSx for NetApp ONTAP backup ([#21960](https://github.com/hashicorp/terraform-provider-aws/issues/21960))
* resource/iam_service_linked_role: Add `tags` argument ([#22185](https://github.com/hashicorp/terraform-provider-aws/issues/22185))

BUG FIXES:

* resource/aws_glue_resource_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22167](https://github.com/hashicorp/terraform-provider-aws/issues/22167))
* resource/aws_media_store_container_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22170](https://github.com/hashicorp/terraform-provider-aws/issues/22170))
* resource/aws_secretsmanager_secret: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22217](https://github.com/hashicorp/terraform-provider-aws/issues/22217))
* resource/aws_secretsmanager_secret_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22217](https://github.com/hashicorp/terraform-provider-aws/issues/22217))
* resource/aws_sns_topic: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22213](https://github.com/hashicorp/terraform-provider-aws/issues/22213))
* resource/aws_sns_topic_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22213](https://github.com/hashicorp/terraform-provider-aws/issues/22213))
* resource/aws_sqs_queue: Fix "error reading, empty result" and various eventual consistency errors ([#22194](https://github.com/hashicorp/terraform-provider-aws/issues/22194))
* resource/aws_sqs_queue: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22194](https://github.com/hashicorp/terraform-provider-aws/issues/22194))
* resource/aws_sqs_queue_policy: Fix "error reading, empty result" and various eventual consistency errors ([#22194](https://github.com/hashicorp/terraform-provider-aws/issues/22194))
* resource/aws_sqs_queue_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22194](https://github.com/hashicorp/terraform-provider-aws/issues/22194))
* resource/aws_ssoadmin_permission_set_inline_policy: Fix erroneous diffs in `inline_policy` when no changes made or policies are equivalent ([#22192](https://github.com/hashicorp/terraform-provider-aws/issues/22192))

## 3.69.0 (December 09, 2021)

FEATURES:

* **New Resource:** `aws_codecommit_approval_rule_template_association` ([#13467](https://github.com/hashicorp/terraform-provider-aws/issues/13467))
* **New Resource:** `aws_detective_graph` ([#22042](https://github.com/hashicorp/terraform-provider-aws/issues/22042))
* **New Resource:** `aws_ec2_subnet_cidr_reservation` ([#22051](https://github.com/hashicorp/terraform-provider-aws/issues/22051))
* **New Resource:** `aws_ecs_task_set` ([#22096](https://github.com/hashicorp/terraform-provider-aws/issues/22096))
* **New Resource:** `aws_emr_studio` ([#21855](https://github.com/hashicorp/terraform-provider-aws/issues/21855))
* **New Resource:** `aws_emr_studio_session_mapping` ([#22140](https://github.com/hashicorp/terraform-provider-aws/issues/22140))

ENHANCEMENTS:

* resource/aws_sqs_queue: Add `sqs_managed_sse_enabled` argument ([#21954](https://github.com/hashicorp/terraform-provider-aws/issues/21954))

BUG FIXES:

* resource/aws_efs_file_system_policy: Fix erroneous diffs in `policy` when no changes made or policies are equivalent ([#22100](https://github.com/hashicorp/terraform-provider-aws/issues/22100))
* resource/aws_lb: Correctly configure `enable_waf_fail_open` during resource creation ([#22072](https://github.com/hashicorp/terraform-provider-aws/issues/22072))

## 3.68.0 (December 02, 2021)

FEATURES:

* **New Data Source:** `aws_codecommit_approval_rule_template` ([#11487](https://github.com/hashicorp/terraform-provider-aws/issues/11487))
* **New Data Source:** `aws_vpc_pool_data_source` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_codecommit_approval_rule_template` ([#11487](https://github.com/hashicorp/terraform-provider-aws/issues/11487))
* **New Resource:** `aws_vpc_ipam` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_vpc_ipam_pool` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_vpc_ipam_scope` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_vpc_ipam_pool_cidr` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_vpc_ipam_pool_cidr_allocation` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))
* **New Resource:** `aws_vpc_ipv6_cidr_block_association` ([#21998](https://github.com/hashicorp/terraform-provider-aws/issues/21998))

ENHANCEMENTS:

* resource/aws_account_alternate_contact: Add `account_id` argument ([#21888](https://github.com/hashicorp/terraform-provider-aws/issues/21888))
* resource/aws_lb_target_group: Add support for `connection_termination` argument for NLBs ([#21130](https://github.com/hashicorp/terraform-provider-aws/issues/21130))
* resource/aws_synthetics_canary: Add `artifact_config` argument. ([#21963](https://github.com/hashicorp/terraform-provider-aws/issues/21963))
* resource/aws_synthetics_canary: Make `artifact_s3_location` updateable. ([#21963](https://github.com/hashicorp/terraform-provider-aws/issues/21963))

BUG FIXES:


## 3.67.0 (November 25, 2021)

FEATURES:

* **New Data Source:** `aws_ec2_instance_types` ([#21850](https://github.com/hashicorp/terraform-provider-aws/issues/21850))
* **New Data Source:** `aws_imagebuilder_image_recipes` ([#21814](https://github.com/hashicorp/terraform-provider-aws/issues/21814))
* **New Resource:** `aws_account_alternate_contact` ([#21789](https://github.com/hashicorp/terraform-provider-aws/issues/21789))
* **New Resource:** `aws_appstream_stack_fleet_association` ([#21484](https://github.com/hashicorp/terraform-provider-aws/issues/21484))
* **New Resource:** `aws_appstream_stack_user_association` ([#21485](https://github.com/hashicorp/terraform-provider-aws/issues/21485))
* **New Resource:** `aws_appstream_user` ([#21485](https://github.com/hashicorp/terraform-provider-aws/issues/21485))
* **New Resource:** `aws_fsx_ontap_storage_virtual_machine` ([#21780](https://github.com/hashicorp/terraform-provider-aws/issues/21780))
* **New Resource:** `aws_fsx_ontap_volume` ([#21889](https://github.com/hashicorp/terraform-provider-aws/issues/21889))

ENHANCEMENTS:

* data source/aws_lambda_function: Add `image_uri` attribute ([#21015](https://github.com/hashicorp/terraform-provider-aws/issues/21015))
* data-source/aws_elb: Add `desync_mitigation_mode` attribute ([#14764](https://github.com/hashicorp/terraform-provider-aws/issues/14764))
* data-source/aws_lb: Add `desync_mitigation_mode` attribute ([#14764](https://github.com/hashicorp/terraform-provider-aws/issues/14764))
* data-source/aws_lb: Add `enable_waf_fail_open` attribute ([#16393](https://github.com/hashicorp/terraform-provider-aws/issues/16393))
* resource/aws_athena_workgroup: Add `engine_version` argument ([#17733](https://github.com/hashicorp/terraform-provider-aws/issues/17733))
* resource/aws_dlm_lifecycle_policy: Add `cross_region_copy_rule` argument in the `schedule` configuration block ([#12868](https://github.com/hashicorp/terraform-provider-aws/issues/12868))
* resource/aws_ec2_fleet: Support in-place update of Launch Template config ([#15387](https://github.com/hashicorp/terraform-provider-aws/issues/15387))
* resource/aws_elb: Add `desync_mitigation_mode` argument ([#14764](https://github.com/hashicorp/terraform-provider-aws/issues/14764))
* resource/aws_lb: Add `desync_mitigation_mode` argument ([#14764](https://github.com/hashicorp/terraform-provider-aws/issues/14764))
* resource/aws_lb: Add `enable_waf_fail_open` argument ([#16393](https://github.com/hashicorp/terraform-provider-aws/issues/16393))
* resource/aws_lb: Update `name` and `name_prefix` plan-time validation to exclude `"internal-"` ([#10693](https://github.com/hashicorp/terraform-provider-aws/issues/10693))

BUG FIXES:


## 3.66.0 (November 18, 2021)

FEATURES:

* **New Data Source:** `aws_emr_release_labels` ([#21767](https://github.com/hashicorp/terraform-provider-aws/issues/21767))
* **New Resource:** `aws_appstream_directory_config` ([#21505](https://github.com/hashicorp/terraform-provider-aws/issues/21505))
* **New Resource:** `aws_iot_thing_group` ([#21799](https://github.com/hashicorp/terraform-provider-aws/issues/21799))
* **New Resource:** `aws_iot_thing_group_membership` ([#21799](https://github.com/hashicorp/terraform-provider-aws/issues/21799))
* **New Resource:** `aws_lambda_layer_version_permission` ([#11941](https://github.com/hashicorp/terraform-provider-aws/issues/11941))
* **New Resource:** `aws_s3_bucket_replication_configuration` ([#20777](https://github.com/hashicorp/terraform-provider-aws/issues/20777))
* **New Resource:** `aws_s3control_access_point_policy` ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* **New Resource:** `aws_s3control_multi_region_access_point` ([#21060](https://github.com/hashicorp/terraform-provider-aws/issues/21060))
* **New Resource:** `aws_s3control_multi_region_access_point_policy` ([#21060](https://github.com/hashicorp/terraform-provider-aws/issues/21060))
* **New Resource:** `aws_s3control_object_lambda_access_point` ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* **New Resource:** `aws_s3control_object_lambda_access_point_policy` ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* **New Resource:** `aws_securityhub_finding_aggregator` ([#21560](https://github.com/hashicorp/terraform-provider-aws/issues/21560))

ENHANCEMENTS:

* aws_s3_access_point: Add `alias` attribute ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* aws_s3_access_point: Add `endpoints` attribute ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* data-source/aws_ec2_instance_type: Add `encryption_in_transit_supported` attribute ([#21837](https://github.com/hashicorp/terraform-provider-aws/issues/21837))
* resource/aws_neptune_cluster: Support in-place update of `engine_version` ([#21760](https://github.com/hashicorp/terraform-provider-aws/issues/21760))

BUG FIXES:

* aws_s3_access_point: `vpc_configuration.vpc_id` is _ForceNew_ ([#19294](https://github.com/hashicorp/terraform-provider-aws/issues/19294))
* resource/aws_internet_gateway: Allow `available` as a *pending* state during gateway detach ([#21794](https://github.com/hashicorp/terraform-provider-aws/issues/21794))

## 3.65.0 (November 11, 2021)

FEATURES:

* **New Data Source:** `aws_key_pair` ([#15829](https://github.com/hashicorp/terraform-provider-aws/issues/15829))
* **New Resource:** `aws_cloudfront_field_level_encryption_config` ([#15033](https://github.com/hashicorp/terraform-provider-aws/issues/15033))
* **New Resource:** `aws_cloudfront_field_level_encryption_profile` ([#12509](https://github.com/hashicorp/terraform-provider-aws/issues/12509))
* **New Resource:** `aws_docdb_global_cluster` ([#20978](https://github.com/hashicorp/terraform-provider-aws/issues/20978))
* **New Resource:** `aws_s3_bucket_intelligent_tiering_configuration` ([#20329](https://github.com/hashicorp/terraform-provider-aws/issues/20329))

ENHANCEMENTS:

* resource/aws_batch_job_queue: Remove limit of 3 items from the `compute_environments` argument ([#21737](https://github.com/hashicorp/terraform-provider-aws/issues/21737))
* resource/aws_docdb_cluster: Add `global_cluster_identifier` argument ([#20978](https://github.com/hashicorp/terraform-provider-aws/issues/20978))

BUG FIXES:

* resource/aws_default_route_table: Fix lack of pagination when describing route tables ([#21710](https://github.com/hashicorp/terraform-provider-aws/issues/21710))
* resource/aws_internet_gateway: Retry resource read after creation to deal with EC2 API eventual consistency ([#21542](https://github.com/hashicorp/terraform-provider-aws/issues/21542))
* resource/aws_servicecatalog_provisioned_product: Allow empty values in provisioning parameters ([#21669](https://github.com/hashicorp/terraform-provider-aws/issues/21669))

## 3.64.2 (November 05, 2021)

BUG FIXES:

* provider: Additional fixes to allow setting endpoints with non-standard, legacy keys. ([#21657](https://github.com/hashicorp/terraform-provider-aws/issues/21657))

## 3.64.1 (November 05, 2021)

BUG FIXES:

* provider: Fix bug preventing custom endpoints from being set ([#21639](https://github.com/hashicorp/terraform-provider-aws/issues/21639))
* provider: Fix bug preventing proper assignment of custom endpoints ([#21641](https://github.com/hashicorp/terraform-provider-aws/issues/21641))

## 3.64.0 (November 04, 2021)

FEATURES:

* **New Data Source:** `aws_cloudfront_response_headers_policy` ([#21620](https://github.com/hashicorp/terraform-provider-aws/issues/21620))
* **New Data Source:** `aws_iam_user_ssh_key` ([#21335](https://github.com/hashicorp/terraform-provider-aws/issues/21335))
* **New Resource:** `aws_backup_vault_lock_configuration` ([#21315](https://github.com/hashicorp/terraform-provider-aws/issues/21315))
* **New Resource:** `aws_cloudfront_response_headers_policy` ([#21620](https://github.com/hashicorp/terraform-provider-aws/issues/21620))
* **New Resource:** `aws_kms_replica_external_key` ([#20533](https://github.com/hashicorp/terraform-provider-aws/issues/20533))
* **New Resource:** `aws_kms_replica_key` ([#20533](https://github.com/hashicorp/terraform-provider-aws/issues/20533))
* **New Resource:** `aws_prometheus_alert_manager_definition` ([#21431](https://github.com/hashicorp/terraform-provider-aws/issues/21431))
* **New Resource:** `aws_prometheus_rule_group_namespace` ([#21470](https://github.com/hashicorp/terraform-provider-aws/issues/21470))

ENHANCEMENTS:

* resource/aws_batch_compute_environment: Add `ec2_configuration` argument to `compute_resources` configuration block ([#21565](https://github.com/hashicorp/terraform-provider-aws/issues/21565))
* resource/aws_glue_crawler: Add `dlq_event_queue_arn` and `event_queue_arn` arguments to the `s3_target` configuration block ([#21467](https://github.com/hashicorp/terraform-provider-aws/issues/21467))
* resource/aws_glue_data_catalog_encryption_settings: Disable encryption on resource deletion ([#21452](https://github.com/hashicorp/terraform-provider-aws/issues/21452))

BUG FIXES:

* aws/resource_aws_lex_slot_type: Correctly determine `version` attribute ([#21509](https://github.com/hashicorp/terraform-provider-aws/issues/21509))
* resource/aws_placement_group: `partition_count` argument is Computed, preventing spurious resource diffs ([#21555](https://github.com/hashicorp/terraform-provider-aws/issues/21555))

## 3.63.0 (October 14, 2021)

FEATURES:

* **New Resource:** `aws_chime_voice_connector_termination_credentials` ([#21162](https://github.com/hashicorp/terraform-provider-aws/issues/21162))
* **New Resource:** `aws_glue_partition_index` ([#21234](https://github.com/hashicorp/terraform-provider-aws/issues/21234))
* **New Resource:** `aws_sagemaker_model_package_group_policy` ([#21250](https://github.com/hashicorp/terraform-provider-aws/issues/21250))

ENHANCEMENTS:

* data-source/aws_instance: Add `placement_partition_number` attribute ([#7777](https://github.com/hashicorp/terraform-provider-aws/issues/7777))
* data-source/glue_connection: Add tagging support. ([#21226](https://github.com/hashicorp/terraform-provider-aws/issues/21226))
* resource/aws_glue_resource_policy: Add `enable_hybrid` argument. ([#21239](https://github.com/hashicorp/terraform-provider-aws/issues/21239))
* resource/aws_instance: Add `placement_partition_number` argument ([#7777](https://github.com/hashicorp/terraform-provider-aws/issues/7777))
* resource/aws_placement_group: Add `partition_count` argument ([#15360](https://github.com/hashicorp/terraform-provider-aws/issues/15360))
* resource/glue_connection: Add tagging support. ([#21226](https://github.com/hashicorp/terraform-provider-aws/issues/21226))
* resource/rds_cluster_instance: Add `performance_insights_retention_period` attribute ([#17111](https://github.com/hashicorp/terraform-provider-aws/issues/17111))

BUG FIXES:

* resource/aws_glue_catalog_table: change `partition_index.keys` to list instead of set ([#21234](https://github.com/hashicorp/terraform-provider-aws/issues/21234))
* resource/glue_catalog_table: Ignore not exists errors on delete ([#21227](https://github.com/hashicorp/terraform-provider-aws/issues/21227))

## 3.62.0 (October 08, 2021)

FEATURES:

* **New Resource:** `aws_dx_connection_confirmation` ([#16489](https://github.com/hashicorp/terraform-provider-aws/issues/16489))
* **New Resource:** `aws_dx_hosted_connection` ([#16489](https://github.com/hashicorp/terraform-provider-aws/issues/16489))

ENHANCEMENTS:

* resource/aws_default_route_table: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#21161](https://github.com/hashicorp/terraform-provider-aws/issues/21161))
* resource/aws_dx_lag: Add `connection_id` argument ([#16489](https://github.com/hashicorp/terraform-provider-aws/issues/16489))
* resource/aws_lex_bot: Added waiter support to account for `BUILDING` status ([#21122](https://github.com/hashicorp/terraform-provider-aws/issues/21122))
* resource/aws_volume_attachment: Add `stop_instance_before_detaching` argument ([#21144](https://github.com/hashicorp/terraform-provider-aws/issues/21144))
* resource/aws_vpn_gateway_route_propagation: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#21161](https://github.com/hashicorp/terraform-provider-aws/issues/21161))

BUG FIXES:

* aws/resource_aws_lex_bot: Correctly determine `version` attribute ([#20383](https://github.com/hashicorp/terraform-provider-aws/issues/20383))
* aws/resource_aws_lex_intent: Correctly determine `version` attribute ([#20383](https://github.com/hashicorp/terraform-provider-aws/issues/20383))
* resource/aws_lb_target_group: Handle attributes at creation: `deregistration_delay`, `load_balancing_algorithm_type`, `preserve_client_ip`, `proxy_protocol_v2`, `slow_start`, `stickiness`, and `lambda_multi_value_headers_enabled` ([#21187](https://github.com/hashicorp/terraform-provider-aws/issues/21187))

## 3.61.0 (October 01, 2021)

FEATURES:

* **New Data Source:** `aws_cloudcontrolapi_resource` ([#21110](https://github.com/hashicorp/terraform-provider-aws/issues/21110))
* **New Data Source:** `aws_db_proxy` ([#21053](https://github.com/hashicorp/terraform-provider-aws/issues/21053))
* **New Data Source:** `aws_ec2_host` ([#10817](https://github.com/hashicorp/terraform-provider-aws/issues/10817))
* **New Data Source:** `aws_kinesis_firehose_delivery_stream` ([#18445](https://github.com/hashicorp/terraform-provider-aws/issues/18445))
* **New Data Source:** `aws_ssm_parameters_by_path` ([#9615](https://github.com/hashicorp/terraform-provider-aws/issues/9615))
* **New Resource:** `aws_appstream_image_builder` ([#21036](https://github.com/hashicorp/terraform-provider-aws/issues/21036))
* **New Resource:** `aws_cloudcontrolapi_resource` ([#21110](https://github.com/hashicorp/terraform-provider-aws/issues/21110))
* **New Resource:** `aws_ec2_host` ([#10817](https://github.com/hashicorp/terraform-provider-aws/issues/10817))
* **New Resource:** `aws_iot_authorizer` ([#14671](https://github.com/hashicorp/terraform-provider-aws/issues/14671))
* **New Resource:** `aws_quicksight_data_source` ([#20710](https://github.com/hashicorp/terraform-provider-aws/issues/20710))
* **New Resource:** `aws_redshift_scheduled_action` ([#13474](https://github.com/hashicorp/terraform-provider-aws/issues/13474))
* **New Resource:** `aws_sagemaker_studio_lifecycle_config` ([#21041](https://github.com/hashicorp/terraform-provider-aws/issues/21041))

ENHANCEMENTS:

* provider: Add parameter `http_proxy` to provider configuration ([#21077](https://github.com/hashicorp/terraform-provider-aws/issues/21077))
* resource/aws_lb_target_group: Support `alb` value for `target_type` argument ([#21069](https://github.com/hashicorp/terraform-provider-aws/issues/21069))
* resource/aws_user_profile: Add `user_settings.jupyter_server_app_settings.lifecycle_config_arns` and `user_settings.kernel_gateway_app_settings.lifecycle_config_arns` arguments ([#21041](https://github.com/hashicorp/terraform-provider-aws/issues/21041))

BUG FIXES:

* resource/aws_dx_connection: Mark `provider_name` as Computed to avoid resource recreation with pre-v3.56.0 configurations ([#21085](https://github.com/hashicorp/terraform-provider-aws/issues/21085))
* resource/aws_dx_lag: Mark `provider_name` as Computed to avoid resource recreation with pre-v3.56.0 configurations ([#21085](https://github.com/hashicorp/terraform-provider-aws/issues/21085))

## 3.60.0 (September 23, 2021)

FEATURES:

* **New Data Source:** `aws_cloudfront_log_delivery_canonical_user_id` ([#15167](https://github.com/hashicorp/terraform-provider-aws/issues/15167))
* **New Data Source:** `aws_cloudwatch_log_groups` ([#17151](https://github.com/hashicorp/terraform-provider-aws/issues/17151))
* **New Data Source:** `aws_connect_contact_flow` ([#16854](https://github.com/hashicorp/terraform-provider-aws/issues/16854))
* **New Data Source:** `aws_connect_instance` ([#16709](https://github.com/hashicorp/terraform-provider-aws/issues/16709))
* **New Data Source:** `aws_iam_users` ([#20877](https://github.com/hashicorp/terraform-provider-aws/issues/20877))
* **New Data Source:** `aws_msk_broker_nodes` ([#20615](https://github.com/hashicorp/terraform-provider-aws/issues/20615))
* **New Data Source:** `aws_msk_kafka_version` ([#20638](https://github.com/hashicorp/terraform-provider-aws/issues/20638))
* **New Resource:** `aws_appstream_fleet` ([#20543](https://github.com/hashicorp/terraform-provider-aws/issues/20543))
* **New Resource:** `aws_chime_voice_connector_streaming` ([#20933](https://github.com/hashicorp/terraform-provider-aws/issues/20933))
* **New Resource:** `aws_connect_contact_flow` ([#16854](https://github.com/hashicorp/terraform-provider-aws/issues/16854))
* **New Resource:** `aws_connect_instance` ([#16709](https://github.com/hashicorp/terraform-provider-aws/issues/16709))
* **New Resource:** `aws_ec2_managed_prefix_list_entry` ([#19394](https://github.com/hashicorp/terraform-provider-aws/issues/19394))
* **New Resource:** `aws_fsx_ontap_filesystem` ([#20951](https://github.com/hashicorp/terraform-provider-aws/issues/20951))
* **New Resource:** `aws_sagemaker_flow_definition` ([#20825](https://github.com/hashicorp/terraform-provider-aws/issues/20825))

ENHANCEMENTS:

* data-source/efs_file_system: Add `transition_to_primary_storage_class` to `lifecycle_policy`. ([#20971](https://github.com/hashicorp/terraform-provider-aws/issues/20971))

BUG FIXES:

* resource/efs_file_system: Allow multiple lifecycle policies. ([#20971](https://github.com/hashicorp/terraform-provider-aws/issues/20971))

## 3.59.0 (September 16, 2021)

FEATURES:

* **New Data Source:** `aws_eks_clusters` ([#20315](https://github.com/hashicorp/terraform-provider-aws/issues/20315))
* **New Data Source:** `aws_eks_node_group` ([#13564](https://github.com/hashicorp/terraform-provider-aws/issues/13564))
* **New Data Source:** `aws_eks_node_groups` ([#13564](https://github.com/hashicorp/terraform-provider-aws/issues/13564))
* **New Resource:** `aws_chime_voice_connector_logging` ([#20863](https://github.com/hashicorp/terraform-provider-aws/issues/20863))
* **New Resource:** `aws_transfer_access` ([#20342](https://github.com/hashicorp/terraform-provider-aws/issues/20342))

ENHANCEMENTS:

* resource/aws_ec2_client_vpn_endpoint: Add `self_service_portal` and `authentication_options.self_service_saml_provider_arn` arguments to support self-service portal ([#17897](https://github.com/hashicorp/terraform-provider-aws/issues/17897))
* resource/aws_ec2_managed_prefix_list: allow updating `max_entries`. ([#20797](https://github.com/hashicorp/terraform-provider-aws/issues/20797))
* resource/aws_efs_file_system: Add `lifecycle_policy.transition_to_primary_storage_class` argument to support Intelligent-Tiering ([#20874](https://github.com/hashicorp/terraform-provider-aws/issues/20874))
* resource/aws_efs_file_system_policy: Add `bypass_policy_lockout_safety_check` argument ([#20838](https://github.com/hashicorp/terraform-provider-aws/issues/20838))
* resource/aws_mwaa_environment: Increase resource creation timeout to 2 hours ([#20861](https://github.com/hashicorp/terraform-provider-aws/issues/20861))

BUG FIXES:


## 3.58.0 (September 09, 2021)

FEATURES:

* **New Resource:** `aws_chime_voice_connector_origination` ([#20676](https://github.com/hashicorp/terraform-provider-aws/issues/20676))
* **New Resource:** `aws_chime_voice_connector_termination` ([#20667](https://github.com/hashicorp/terraform-provider-aws/issues/20667))
* **New Resource:** `aws_quicksight_group_membership` ([#20687](https://github.com/hashicorp/terraform-provider-aws/issues/20687))

## 3.57.0 (September 02, 2021)

FEATURES:

* **New Resource:** `aws_service_discovery_instance` ([#17498](https://github.com/hashicorp/terraform-provider-aws/issues/17498))

ENHANCEMENTS:

* data-source/aws_instance: Add `ipv6_addresses` attribute ([#17859](https://github.com/hashicorp/terraform-provider-aws/issues/17859))
* resource/aws_athena_database: Read the database name from the `AwsDataCatalog` ([#19765](https://github.com/hashicorp/terraform-provider-aws/issues/19765))
* resource/aws_service_discovery_service: Add `force_destroy` argument ([#3538](https://github.com/hashicorp/terraform-provider-aws/issues/3538))
* resource_aws_route53_health_check: Add `RECOVERY_CONTROL` health check type and `routing_control_arn` argument ([#20731](https://github.com/hashicorp/terraform-provider-aws/issues/20731))
* resource_vpn_connection: Handle paginated response when reading Transit Gateway Attachments ([#20775](https://github.com/hashicorp/terraform-provider-aws/issues/20775))

BUG FIXES:


## 3.56.0 (August 26, 2021)

FEATURES:

* **New Data Source:** `aws_dx_connection` ([#17852](https://github.com/hashicorp/terraform-provider-aws/issues/17852))
* **New Data Source:** `aws_dx_location` ([#9735](https://github.com/hashicorp/terraform-provider-aws/issues/9735))
* **New Data Source:** `aws_dx_locations` ([#9735](https://github.com/hashicorp/terraform-provider-aws/issues/9735))
* **New Resource:** `aws_appstream_stack` ([#20547](https://github.com/hashicorp/terraform-provider-aws/issues/20547))
* **New Resource:** `aws_autoscaling_group_tag` ([#20009](https://github.com/hashicorp/terraform-provider-aws/issues/20009))
* **New Resource:** `aws_dynamodb_tag` ([#13783](https://github.com/hashicorp/terraform-provider-aws/issues/13783))
* **New Resource:** `aws_ecs_tag` ([#13783](https://github.com/hashicorp/terraform-provider-aws/issues/13783))
* **New Resource:** `aws_route53recoverycontrolconfig_cluster` ([#20568](https://github.com/hashicorp/terraform-provider-aws/issues/20568))
* **New Resource:** `aws_route53recoverycontrolconfig_control_panel` ([#20568](https://github.com/hashicorp/terraform-provider-aws/issues/20568))
* **New Resource:** `aws_route53recoverycontrolconfig_routing_control` ([#20568](https://github.com/hashicorp/terraform-provider-aws/issues/20568))
* **New Resource:** `aws_route53recoverycontrolconfig_safety_rule` ([#20568](https://github.com/hashicorp/terraform-provider-aws/issues/20568))
* **New Resource:** `aws_route53recoveryreadiness_cell` ([#20526](https://github.com/hashicorp/terraform-provider-aws/issues/20526))
* **New Resource:** `aws_route53recoveryreadiness_readiness_check` ([#20526](https://github.com/hashicorp/terraform-provider-aws/issues/20526))
* **New Resource:** `aws_route53recoveryreadiness_recovery_group` ([#20526](https://github.com/hashicorp/terraform-provider-aws/issues/20526))
* **New Resource:** `aws_route53recoveryreadiness_resource_set` ([#20526](https://github.com/hashicorp/terraform-provider-aws/issues/20526))

ENHANCEMENTS:

* data-source/aws_efs_file_system: Add ability to filter results by `tags` ([#20399](https://github.com/hashicorp/terraform-provider-aws/issues/20399))
* resource/aws_dx_connection: Add `owner_account_id` attribute ([#17852](https://github.com/hashicorp/terraform-provider-aws/issues/17852))
* resource/aws_dx_connection: Add `provider_name` argument ([#17852](https://github.com/hashicorp/terraform-provider-aws/issues/17852))
* resource/aws_dx_lag: Add `owner_account_id` attribute ([#17852](https://github.com/hashicorp/terraform-provider-aws/issues/17852))
* resource/aws_dx_lag: Add `provider_name` argument ([#17852](https://github.com/hashicorp/terraform-provider-aws/issues/17852))
* resource/aws_fsx_lustre_filesystem: Allow creating filesystem from backup using `backup_id`. ([#20614](https://github.com/hashicorp/terraform-provider-aws/issues/20614))
* resource/aws_fsx_windows_filesystem: Allow creating filesystem from backup using `backup_id`. ([#20643](https://github.com/hashicorp/terraform-provider-aws/issues/20643))

BUG FIXES:


## 3.55.0 (August 19, 2021)

FEATURES:

* **New Data Source:** `aws_iam_roles` ([#18585](https://github.com/hashicorp/terraform-provider-aws/issues/18585))
* **New Data Source:** `aws_subnets` ([#18803](https://github.com/hashicorp/terraform-provider-aws/issues/18803))
* **New Resource:** `aws_chime_voice_connector_group` ([#20565](https://github.com/hashicorp/terraform-provider-aws/issues/20565))
* **New Resource:** `aws_fsx_backup` ([#20569](https://github.com/hashicorp/terraform-provider-aws/issues/20569))
* **New Resource:** `aws_sagemaker_device_fleet` ([#20058](https://github.com/hashicorp/terraform-provider-aws/issues/20058))
* **New Resource:** `aws_sagemaker_human_task_ui` ([#20570](https://github.com/hashicorp/terraform-provider-aws/issues/20570))

ENHANCEMENTS:

* aws/resource_aws_appconfig_deployment: Add `state` attribute ([#20288](https://github.com/hashicorp/terraform-provider-aws/issues/20288))

BUG FIXES:

* aws/resource_aws_appconfig_deployment: Remove internal waiter after start of deployment ([#20288](https://github.com/hashicorp/terraform-provider-aws/issues/20288))
* aws/resource_aws_cloudwatch_event_rule: Correctly handle ARN in `event_bus_name` argument ([#20312](https://github.com/hashicorp/terraform-provider-aws/issues/20312))
* aws/resource_aws_cloudwatch_event_target: Correctly handle ARN in `event_bus_name` argument ([#20312](https://github.com/hashicorp/terraform-provider-aws/issues/20312))

## 3.54.0 (August 12, 2021)

FEATURES:

* **New Resource:** `aws_chime_voice_connector` ([#19504](https://github.com/hashicorp/terraform-provider-aws/issues/19504))
* **New Resource:** `aws_shield_protection_group` ([#20491](https://github.com/hashicorp/terraform-provider-aws/issues/20491))

ENHANCEMENTS:

* data-source/aws_workspaces_directory: Add `workspace_access_properties.device_type_linux` attribute ([#20462](https://github.com/hashicorp/terraform-provider-aws/issues/20462))
* resource/aws_athena_workgroup: Add `requester_pays_enabled` argument ([#20457](https://github.com/hashicorp/terraform-provider-aws/issues/20457))
* resource/aws_workspaces_directory: Add `workspace_access_properties.device_type_linux` argument ([#20462](https://github.com/hashicorp/terraform-provider-aws/issues/20462))

BUG FIXES:

* aws/resource_aws_imagebuilder_infrastructure_configuration: Always set `terminate_instance_on_failure` on create and update ([#20464](https://github.com/hashicorp/terraform-provider-aws/issues/20464))
* resource/aws_synthetics_canary: Correctly report any resource creation errors ([#20463](https://github.com/hashicorp/terraform-provider-aws/issues/20463))

## 3.53.0 (August 05, 2021)

ENHANCEMENTS:

* data-source/aws_acm_certificate: Add status attribute ([#20232](https://github.com/hashicorp/terraform-provider-aws/issues/20232))
* data-source/aws_ec2_coip_pool: Add `arn` attribute ([#17046](https://github.com/hashicorp/terraform-provider-aws/issues/17046))

BUG FIXES:

* aws/resource_aws_amplify_branch: Correctly handle branch names that contain '/' ([#20426](https://github.com/hashicorp/terraform-provider-aws/issues/20426))
* aws/resource_aws_apigateway_vpc_link: Ensure deletion does not return an error when resource is not found ([#20441](https://github.com/hashicorp/terraform-provider-aws/issues/20441))
* aws/resource_aws_instance: Fix running `terraform plan` with with `skip_credentials_validation=true` ([#20357](https://github.com/hashicorp/terraform-provider-aws/issues/20357))
* aws/resource_aws_instance: Fix state refresh when launch template was deleted ([#20357](https://github.com/hashicorp/terraform-provider-aws/issues/20357))

## 3.52.0 (July 29, 2021)

FEATURES:

* **New Resource:** `aws_sagemaker_workforce` ([#20065](https://github.com/hashicorp/terraform-provider-aws/issues/20065))
* **New Resource:** `aws_sagemaker_workteam` ([#20122](https://github.com/hashicorp/terraform-provider-aws/issues/20122))
* **New Resource:** `aws_storagegateway_file_system_association` ([#20082](https://github.com/hashicorp/terraform-provider-aws/issues/20082))

ENHANCEMENTS:

* data-source/aws_ec2_instance_type_offerings: Add `locations` and `location_types` attributes ([#16704](https://github.com/hashicorp/terraform-provider-aws/issues/16704))
* data-source/aws_lb: Add ability to filter results by `tags` ([#6458](https://github.com/hashicorp/terraform-provider-aws/issues/6458))
* data-source/aws_qldb_ledger: Add `permissions_mode` attribute ([#20302](https://github.com/hashicorp/terraform-provider-aws/issues/20302))
* resource/aws_budgets_budget: Add the `cost_filter` argument which allows multiple `values` to be specified per filter. This new argument will eventually replace the `cost_filters` argument ([#9092](https://github.com/hashicorp/terraform-provider-aws/issues/9092))
* resource/aws_budgets_budget: Change `time_period_start` to an optional argument. If you don't specify a start date, AWS defaults to the start of your chosen time period ([#9092](https://github.com/hashicorp/terraform-provider-aws/issues/9092))
* resource/aws_dx_connection: Add support for `100Gbps` `bandwidth` [#20364](https://github.com/hashicorp/terraform-provider-aws/issues/20364))
* resource/aws_dx_lag: Add support for `100Gbps` `connections_bandwidth` [#20364](https://github.com/hashicorp/terraform-provider-aws/issues/20364))
* resource/aws_qldb_ledger: Add `permissions_mode` support ([#20302](https://github.com/hashicorp/terraform-provider-aws/issues/20302))
* resource/aws_secretsmanager_secret: Add replica support ([#20293](https://github.com/hashicorp/terraform-provider-aws/issues/20293))

BUG FIXES:

* resource/aws_budgets_budget: Change the service name in the `arn` attribute from `budgetservice` to `budgets` ([#9092](https://github.com/hashicorp/terraform-provider-aws/issues/9092))
* resource/aws_budgets_budget: Suppress plan differences with trailing zeroes for `limit_amount` ([#9092](https://github.com/hashicorp/terraform-provider-aws/issues/9092))
* resource/aws_budgets_budget_action: Change the service name in the `arn` attribute from `budgetservice` to `budgets` ([#9092](https://github.com/hashicorp/terraform-provider-aws/issues/9092))
* resource/aws_lex_bot: Fix computed `version` for dependent resources ([#20336](https://github.com/hashicorp/terraform-provider-aws/issues/20336))
* resource/aws_lex_intent: Fix computed `version` for dependent resources ([#20336](https://github.com/hashicorp/terraform-provider-aws/issues/20336))
* resource/aws_lex_slot_type: Fix computed `version` for dependent resources ([#20336](https://github.com/hashicorp/terraform-provider-aws/issues/20336))

## 3.51.0 (July 22, 2021)

FEATURES:

* **New Data Source:** `aws_elasticache_user` ([#16629](https://github.com/hashicorp/terraform-provider-aws/issues/16629))
* **New Resource:** `aws_appconfig_deployment` ([#20172](https://github.com/hashicorp/terraform-provider-aws/issues/20172))
* **New Resource:** `aws_elasticache_user` ([#16629](https://github.com/hashicorp/terraform-provider-aws/issues/16629))
* **New Resource:** `aws_elasticache_user_group` ([#16504](https://github.com/hashicorp/terraform-provider-aws/issues/16504))

ENHANCEMENTS:

* resource/aws_glue_crawler: Add `sample_size` argument in `s3_target` block. ([#20203](https://github.com/hashicorp/terraform-provider-aws/issues/20203))
* resource/aws_instance: Add support for configuration with Launch Template ([#10807](https://github.com/hashicorp/terraform-provider-aws/issues/10807))
* resource/aws_servicecatalog_provisioned_product: Increase timeouts to align with CloudFormation (30 min.) ([#20254](https://github.com/hashicorp/terraform-provider-aws/issues/20254))

BUG FIXES:

* aws/resource_aws_lambda_event_source_mapping: Ignore `InvalidParameterValueException` error caused by IAM propagation when creating Lambda event source mapping with Kinesis stream source ([#20229](https://github.com/hashicorp/terraform-provider-aws/issues/20229))
* aws/resource_aws_route_table_association: Correctly handle `associated` as a pending state when waiting for deletion of an association ([#20265](https://github.com/hashicorp/terraform-provider-aws/issues/20265))

## 3.50.0 (July 15, 2021)

NOTES:

* resource/aws_dx_gateway_association_proposal: If an accepted Proposal reaches end-of-life and is removed by AWS do not recreate the resource, instead refreshing Terraform state from the resource's Direct Connect Gateway ID and Associated Gateway ID. ([#19741](https://github.com/hashicorp/terraform-provider-aws/issues/19741))

FEATURES:

* **New Resource:** `aws_appconfig_application` ([#19307](https://github.com/hashicorp/terraform-provider-aws/issues/19307))
* **New Resource:** `aws_appconfig_configuration_profile` ([#19320](https://github.com/hashicorp/terraform-provider-aws/issues/19320))
* **New Resource:** `aws_appconfig_deployment_strategy` ([#19359](https://github.com/hashicorp/terraform-provider-aws/issues/19359))
* **New Resource:** `aws_appconfig_environment` ([#19307](https://github.com/hashicorp/terraform-provider-aws/issues/19307))
* **New Resource:** `aws_appconfig_hosted_configuration_version` ([#19324](https://github.com/hashicorp/terraform-provider-aws/issues/19324))
* **New Resource:** `aws_config_organization_conformance_pack` ([#17298](https://github.com/hashicorp/terraform-provider-aws/issues/17298))
* **New Resource:** `aws_securityhub_organization_configuration` ([#19108](https://github.com/hashicorp/terraform-provider-aws/issues/19108))
* **New Resource:** `aws_securityhub_standards_control` ([#14714](https://github.com/hashicorp/terraform-provider-aws/issues/14714))

ENHANCEMENTS:

* resource/aws_fsx_windows_file_system: Add `aliases` argument ([#20054](https://github.com/hashicorp/terraform-provider-aws/issues/20054))

BUG FIXES:


## 3.49.0 (July 08, 2021)

FEATURES:

* **New Resource:** `aws_eks_identity_provider_config` ([#17959](https://github.com/hashicorp/terraform-provider-aws/issues/17959))
* **New Resource:** `aws_rds_cluster_role_association` ([#12370](https://github.com/hashicorp/terraform-provider-aws/issues/12370))

ENHANCEMENTS:

* aws_rds_cluster: Set `iam_roles` as Computed to prevent drift when the `aws_rds_cluster_role_association` resource is used ([#12370](https://github.com/hashicorp/terraform-provider-aws/issues/12370))

BUG FIXES:

* data-source/aws_lakeformation_permissions: Fix various problems with permissions including select-only ([#20108](https://github.com/hashicorp/terraform-provider-aws/issues/20108))
* resource/aws_lakeformation_permissions: Fix various problems with permissions including select-only ([#20108](https://github.com/hashicorp/terraform-provider-aws/issues/20108))
* resource/aws_ram_resource_share_accepter: Allow destroy even where AWS API provides no way to disassociate ([#19718](https://github.com/hashicorp/terraform-provider-aws/issues/19718))

## 3.48.0 (July 02, 2021)

FEATURES:

* **New Data Source:** `aws_iam_session_context` ([#19957](https://github.com/hashicorp/terraform-provider-aws/issues/19957))
* **New Data Source:** `aws_servicecatalog_launch_paths` ([#19572](https://github.com/hashicorp/terraform-provider-aws/issues/19572))
* **New Data Source:** `aws_servicecatalog_portfolio_constraints` ([#19813](https://github.com/hashicorp/terraform-provider-aws/issues/19813))
* **New Resource:** `aws_cloudfront_monitoring_subscription` ([#18083](https://github.com/hashicorp/terraform-provider-aws/issues/18083))
* **New Resource:** `aws_servicecatalog_provisioned_product` ([#19459](https://github.com/hashicorp/terraform-provider-aws/issues/19459))

ENHANCEMENTS:

* resource/aws_fsx_windows_file_system: Add `audit_log_configuration` argument. ([#19970](https://github.com/hashicorp/terraform-provider-aws/issues/19970))

BUG FIXES:

* resource/aws_mwaa_environment: Changes to the `kms_key` argument force resource recreation ([#19994](https://github.com/hashicorp/terraform-provider-aws/issues/19994))

## 3.47.0 (June 24, 2021)

FEATURES:

* **New Resource:** `aws_cloudwatch_event_bus_policy` ([#16874](https://github.com/hashicorp/terraform-provider-aws/issues/16874))
* **New Resource:** `aws_efs_backup_policy` ([#18006](https://github.com/hashicorp/terraform-provider-aws/issues/18006))
* **New Resource:** `aws_elasticsearch_domain_saml_options` ([#19497](https://github.com/hashicorp/terraform-provider-aws/issues/19497))
* **New Resource:** `aws_neptune_cluster_endpoint` ([#19898](https://github.com/hashicorp/terraform-provider-aws/issues/19898))

ENHANCEMENTS:

* resource/aws_default_route_table: Add retries when creating, deleting and replacing routes ([#19426](https://github.com/hashicorp/terraform-provider-aws/issues/19426))
* resource/aws_default_route_table: Add retries when creating, deleting and replacing routes ([#19426](https://github.com/hashicorp/terraform-provider-aws/issues/19426))
* resource/aws_lb_target_group: Add support for `app_cookie` stickiness type and `cookie_name` argument ([#18102](https://github.com/hashicorp/terraform-provider-aws/issues/18102))
* resource/aws_main_route_table_association: Wait for association to reach the required state ([#19426](https://github.com/hashicorp/terraform-provider-aws/issues/19426))
* resource/aws_neptune_cluster: Add `copy_tags_to_snapshot` argument ([#19899](https://github.com/hashicorp/terraform-provider-aws/issues/19899))

BUG FIXES:

* resource/aws_glue_catalog_database: Set `location_uri` as compute to prevent drift when `target_table` has `location_uri` set. ([#19743](https://github.com/hashicorp/terraform-provider-aws/issues/19743))
* resource/aws_glue_catalog_table: Fix updating `schema_reference` when columns are present. ([#19742](https://github.com/hashicorp/terraform-provider-aws/issues/19742))

## 3.46.0 (June 17, 2021)

FEATURES:

* **New Data Source:** `aws_appmesh_virtual_service` ([#19774](https://github.com/hashicorp/terraform-provider-aws/issues/19774))
* **New Data Source:** `aws_servicecatalog_portfolio` ([#19500](https://github.com/hashicorp/terraform-provider-aws/issues/19500))
* **New Resource:** `aws_budgets_budget_action` ([#19554](https://github.com/hashicorp/terraform-provider-aws/issues/19554))
* **New Resource:** `aws_route53_resolver_firewall_config` ([#18733](https://github.com/hashicorp/terraform-provider-aws/issues/18733))

ENHANCEMENTS:


BUG FIXES:

* data-source/aws_directory_service_directory: Check VpcSettings and ConnectSettings for nil values ([#19820](https://github.com/hashicorp/terraform-provider-aws/issues/19820))
* data-source/aws_lakeformation_permissions: Fix diffs resulting from order of column names and exclude column names ([#19817](https://github.com/hashicorp/terraform-provider-aws/issues/19817))
* resource/aws_glue_trigger: Fix default timeouts for Create and Delete operations ([#19827](https://github.com/hashicorp/terraform-provider-aws/issues/19827))
* resource/aws_lakeformation_permissions: Fix bug preventing updates (inconsistent result) ([#19817](https://github.com/hashicorp/terraform-provider-aws/issues/19817))
* resource/aws_lakeformation_permissions: Fix bug where resource is not properly removed from state ([#19817](https://github.com/hashicorp/terraform-provider-aws/issues/19817))
* resource/aws_lakeformation_permissions: Fix diffs resulting only from order of column names and exclude column names ([#19817](https://github.com/hashicorp/terraform-provider-aws/issues/19817))
* resource/aws_sqs_queue: Correctly handle the default `kms_data_key_reuse_period_seconds` value of `300` for unencrypted queues ([#19834](https://github.com/hashicorp/terraform-provider-aws/issues/19834))

## 3.45.0 (June 10, 2021)

FEATURES:

* **New Data Source:** `aws_appmesh_mesh` ([#19577](https://github.com/hashicorp/terraform-provider-aws/issues/19577))
* **New Data Source:** `aws_globalaccelerator_accelerator` ([#19647](https://github.com/hashicorp/terraform-provider-aws/issues/19647))

ENHANCEMENTS:

* data-source/aws_nat_gateway: Add `connectivity_type` attribute ([#19758](https://github.com/hashicorp/terraform-provider-aws/issues/19758))
* resource/aws_cur_report_definition: Add `arn` attribute. ([#19705](https://github.com/hashicorp/terraform-provider-aws/issues/19705))
* resource/aws_cur_report_definition: Add plan time validation for `report_name`. ([#19705](https://github.com/hashicorp/terraform-provider-aws/issues/19705))
* resource/aws_cur_report_definition: Support updating definition. ([#19705](https://github.com/hashicorp/terraform-provider-aws/issues/19705))
* resource/aws_default_vpc_dhcp_options: Add `owner_id` argument. ([#19656](https://github.com/hashicorp/terraform-provider-aws/issues/19656))
* resource/aws_fsx_lustre_filesystem: Add `data_compression_type` argument. ([#19664](https://github.com/hashicorp/terraform-provider-aws/issues/19664))
* resource/aws_nat_gateway: Add `connectivity_type` argument ([#19758](https://github.com/hashicorp/terraform-provider-aws/issues/19758))
* resource/aws_sqs_queue: Add `deduplication_scope` and `fifo_throughput_limit` arguments ([#19639](https://github.com/hashicorp/terraform-provider-aws/issues/19639))
* resource/aws_sqs_queue: Add `url` attribute ([#19639](https://github.com/hashicorp/terraform-provider-aws/issues/19639))

BUG FIXES:

* data-source/aws_acmpca_certificate_authority: Fix `error setting tags` ([#19681](https://github.com/hashicorp/terraform-provider-aws/issues/19681))
* data-source/aws_servicequotas_service_quota: Correctly handle errors embedded in API struct ([#19722](https://github.com/hashicorp/terraform-provider-aws/issues/19722))
* resource/aws_batch_job_definition: Suppress differences for empty `linuxParameters.devices` and `linuxParameters.tmpfs` arrays in the `container_properties` argument ([#19666](https://github.com/hashicorp/terraform-provider-aws/issues/19666))
* resource/aws_servicequotas_service_quota: Correctly handle errors embedded in API struct ([#19722](https://github.com/hashicorp/terraform-provider-aws/issues/19722))
* resource/aws_sqs_queue: Allow `visibility_timeout_seconds` to be `0` when creating queue ([#19639](https://github.com/hashicorp/terraform-provider-aws/issues/19639))
* resource/aws_sqs_queue: Ensure that queue attributes propagate completely during Create and Update ([#19639](https://github.com/hashicorp/terraform-provider-aws/issues/19639))

## 3.44.0 (June 03, 2021)

FEATURES:

* **New Resource:** `aws_amplify_branch` ([#11937](https://github.com/hashicorp/terraform-provider-aws/issues/11937))
* **New Resource:** `aws_amplify_domain_association` ([#11938](https://github.com/hashicorp/terraform-provider-aws/issues/11938))
* **New Resource:** `aws_amplify_webhook` ([#11939](https://github.com/hashicorp/terraform-provider-aws/issues/11939))
* **New Resource:** `aws_servicecatalog_principal_portfolio_association` ([#19470](https://github.com/hashicorp/terraform-provider-aws/issues/19470))

ENHANCEMENTS:

* resource/aws_acmpca_certificate_authority: Add `s3_object_acl` argument to `revocation_configuration.crl_configuration` configuration block ([#19578](https://github.com/hashicorp/terraform-provider-aws/issues/19578))
* resource/aws_devicefarm_project: Add `default_job_timeout_minutes` and `tags` argument ([#19574](https://github.com/hashicorp/terraform-provider-aws/issues/19574))
* resource/aws_devicefarm_project: Add plan time validation for `name` ([#19574](https://github.com/hashicorp/terraform-provider-aws/issues/19574))
* resource/aws_fsx_lustre_filesystem: Allow updating `storage_capacity`. ([#19568](https://github.com/hashicorp/terraform-provider-aws/issues/19568))

BUG FIXES:

* resource/aws_amplify_app: Mark the `enable_performance_mode` argument in the `auto_branch_creation_config` configuration block as `ForceNew` ([#11937](https://github.com/hashicorp/terraform-provider-aws/issues/11937))

## 3.43.0 (June 01, 2021)

FEATURES:

* **New Data Source:** `aws_cloudwatch_event_connection` ([#18905](https://github.com/hashicorp/terraform-provider-aws/issues/18905))
* **New Resource:** `aws_amplify_app` ([#15966](https://github.com/hashicorp/terraform-provider-aws/issues/15966))
* **New Resource:** `aws_amplify_backend_environment` ([#11936](https://github.com/hashicorp/terraform-provider-aws/issues/11936))
* **New Resource:** `aws_cloudwatch_event_api_destination` ([#18905](https://github.com/hashicorp/terraform-provider-aws/issues/18905))
* **New Resource:** `aws_cloudwatch_event_connection` ([#18905](https://github.com/hashicorp/terraform-provider-aws/issues/18905))
* **New Resource:** `aws_schemas_discoverer` ([#19100](https://github.com/hashicorp/terraform-provider-aws/issues/19100))
* **New Resource:** `aws_schemas_registry` ([#19100](https://github.com/hashicorp/terraform-provider-aws/issues/19100))
* **New Resource:** `aws_schemas_schema` ([#19100](https://github.com/hashicorp/terraform-provider-aws/issues/19100))
* **New Resource:** `aws_servicecatalog_budget_resource_association` ([#19452](https://github.com/hashicorp/terraform-provider-aws/issues/19452))
* **New Resource:** `aws_servicecatalog_provisioning_artifact` ([#19316](https://github.com/hashicorp/terraform-provider-aws/issues/19316))
* **New Resource:** `aws_servicecatalog_tag_option_resource_association` ([#19448](https://github.com/hashicorp/terraform-provider-aws/issues/19448))

ENHANCEMENTS:

* resource/aws_ec2_capacity_reservation: Add `outpost_arn` argument ([#19535](https://github.com/hashicorp/terraform-provider-aws/issues/19535))
* resource/aws_sns_topic: Add `firehose_success_feedback_role_arn`, `firehose_success_feedback_sample_rate` and `firehose_failure_feedback_role_arn` arguments. ([#19528](https://github.com/hashicorp/terraform-provider-aws/issues/19528))
* resource/aws_sns_topic: Add `owner` attribute. ([#19528](https://github.com/hashicorp/terraform-provider-aws/issues/19528))
* resource/aws_sns_topic: Add plan time validation for `application_success_feedback_role_arn`, `application_failure_feedback_role_arn`, `http_success_feedback_role_arn`, `http_failure_feedback_role_arn`, `lambda_success_feedback_role_arn`, `lambda_failure_feedback_role_arn`, `sqs_success_feedback_role_arn`, `sqs_failure_feedback_role_arn`. ([#19528](https://github.com/hashicorp/terraform-provider-aws/issues/19528))

BUG FIXES:

* data-source/aws_mq_broker: Correct type for `logs.audit` attribute ([#19502](https://github.com/hashicorp/terraform-provider-aws/issues/19502))
* resource/aws_batch_job_definition: Don't crash when setting `timeout.attempt_duration_seconds` to `null` ([#19505](https://github.com/hashicorp/terraform-provider-aws/issues/19505))
* resource/aws_ec2_managed_prefix_list: Fix crash with multiple description-only updates ([#19517](https://github.com/hashicorp/terraform-provider-aws/issues/19517))
* resource/aws_glue_connection: `connection_properties` are optional ([#19375](https://github.com/hashicorp/terraform-provider-aws/issues/19375))
* resource/aws_lb_listener_rule: Allow blank string for `action.redirect.query` nested argument ([#19496](https://github.com/hashicorp/terraform-provider-aws/issues/19496))
* resource/aws_synthetics_canary: Change minimum `timeout_in_seconds` in `run_config` from `60` to `3` ([#19515](https://github.com/hashicorp/terraform-provider-aws/issues/19515))
* resource/aws_vpn_connection: Allow `local_ipv4_network_cidr`, `remote_ipv4_network_cidr`, `local_ipv6_network_cidr`, and `remote_ipv6_network_cidr` to be CIDRs of any size ([#17573](https://github.com/hashicorp/terraform-provider-aws/issues/17573))

## 3.42.0 (May 20, 2021)

FEATURES:

* **New Data Source:** `aws_service_discovery_dns_namespace` ([#6856](https://github.com/hashicorp/terraform-provider-aws/issues/6856))
* **New Resource:** `aws_cloudwatch_metric_stream` ([#18870](https://github.com/hashicorp/terraform-provider-aws/issues/18870))
* **New Resource:** `aws_servicecatalog_constraint` ([#19385](https://github.com/hashicorp/terraform-provider-aws/issues/19385))
* **New Resource:** `aws_servicecatalog_product_portfolio_association` ([#19385](https://github.com/hashicorp/terraform-provider-aws/issues/19385))
* **New Resource:** `aws_servicecatalog_service_action` ([#19369](https://github.com/hashicorp/terraform-provider-aws/issues/19369))

ENHANCEMENTS:


BUG FIXES:


## 3.41.0 (May 19, 2021)

FEATURES:

* **New Data Source:** `aws_cloudfront_function` ([#19315](https://github.com/hashicorp/terraform-provider-aws/issues/19315))
* **New Data Source:** `aws_glue_connection` ([#18802](https://github.com/hashicorp/terraform-provider-aws/issues/18802))
* **New Data Source:** `aws_glue_data_catalog_encryption_settings` ([#18802](https://github.com/hashicorp/terraform-provider-aws/issues/18802))
* **New Data Source:** `aws_organizations_delegated_administrators` ([#19389](https://github.com/hashicorp/terraform-provider-aws/issues/19389))
* **New Data Source:** `aws_organizations_delegated_services` ([#19389](https://github.com/hashicorp/terraform-provider-aws/issues/19389))
* **New Resource:** `aws_apprunner_auto_scaling_configuration_version` ([#19432](https://github.com/hashicorp/terraform-provider-aws/issues/19432))
* **New Resource:** `aws_apprunner_connection` ([#19432](https://github.com/hashicorp/terraform-provider-aws/issues/19432))
* **New Resource:** `aws_apprunner_custom_domain_association` ([#19432](https://github.com/hashicorp/terraform-provider-aws/issues/19432))
* **New Resource:** `aws_apprunner_service` ([#19432](https://github.com/hashicorp/terraform-provider-aws/issues/19432))
* **New Resource:** `aws_cloudfront_function` ([#19315](https://github.com/hashicorp/terraform-provider-aws/issues/19315))
* **New Resource:** `aws_macie2_invitation_accepter` ([#19304](https://github.com/hashicorp/terraform-provider-aws/issues/19304))
* **New Resource:** `aws_macie2_member` ([#19304](https://github.com/hashicorp/terraform-provider-aws/issues/19304))
* **New Resource:** `aws_macie2_organization_admin_account` ([#19303](https://github.com/hashicorp/terraform-provider-aws/issues/19303))
* **New Resource:** `aws_organizations_delegated_administrator` ([#19389](https://github.com/hashicorp/terraform-provider-aws/issues/19389))
* **New Resource:** `aws_servicecatalog_organizations_access` ([#19278](https://github.com/hashicorp/terraform-provider-aws/issues/19278))
* **New Resource:** `aws_servicecatalog_portfolio_share` ([#19278](https://github.com/hashicorp/terraform-provider-aws/issues/19278))

ENHANCEMENTS:

* data-source/aws_outposts_outpost: `owner_id` is now an optional argument ([#17585](https://github.com/hashicorp/terraform-provider-aws/issues/17585))
* data-source/aws_outposts_outposts: Add `owner_id` argument ([#17585](https://github.com/hashicorp/terraform-provider-aws/issues/17585))
* resource/aws_glue_catalog_database: Add `target_database` argument ([#19371](https://github.com/hashicorp/terraform-provider-aws/issues/19371))
* resource/aws_glue_catalog_table: Add `target_table` argument ([#19372](https://github.com/hashicorp/terraform-provider-aws/issues/19372))

BUG FIXES:


## 3.40.0 (May 13, 2021)

FEATURES:

* **New Resource:** `aws_macie2_custom_data_identifier` ([#19254](https://github.com/hashicorp/terraform-provider-aws/issues/19254))
* **New Resource:** `aws_macie2_findings_filter` ([#19283](https://github.com/hashicorp/terraform-provider-aws/issues/19283))
* **New Resource:** `aws_servicecatalog_tag_option` ([#19300](https://github.com/hashicorp/terraform-provider-aws/issues/19300))
* **New Resource:** `aws_timestreamwrite_database` ([#15463](https://github.com/hashicorp/terraform-provider-aws/issues/15463))
* **New Resource:** `aws_timestreamwrite_table` ([#19354](https://github.com/hashicorp/terraform-provider-aws/issues/19354))

ENHANCEMENTS:

* data-source/aws_lb_listener: Add `tags` attribute. ([#19286](https://github.com/hashicorp/terraform-provider-aws/issues/19286))
* resource/aws_lb_listener: Add `tags` argument & `tags_all` attribute. ([#19286](https://github.com/hashicorp/terraform-provider-aws/issues/19286))
* resource/aws_lb_listener_rule: Add plan time validation to `listener_arn`, `action.target_group_arn`, `action.forward.target_group.arn`, `action.redirect.host`, `action.redirect.path`, `action.redirect.query`, `action.redirect.status_code`, `action.fixed_response.message_body`, `action.authenticate_cognito.user_pool_arn`. ([#19285](https://github.com/hashicorp/terraform-provider-aws/issues/19285))
* resource/aws_lb_listener_rule: Add tagging support. ([#19285](https://github.com/hashicorp/terraform-provider-aws/issues/19285))

## 3.39.0 (May 06, 2021)

FEATURES:

* **New Data Source:** `aws_cloudwatch_event_source` ([#19219](https://github.com/hashicorp/terraform-provider-aws/issues/19219))
* **New Resource:** `aws_dynamodb_kinesis_streaming_destination` ([#16743](https://github.com/hashicorp/terraform-provider-aws/issues/16743))
* **New Resource:** `aws_macie2_classification_job` ([#19165](https://github.com/hashicorp/terraform-provider-aws/issues/19165))

ENHANCEMENTS:

* resource/aws_glue_connection: Add plan time validation for `connection_properties`, `description`, `match_criteria`, `name`, and `physical_connection_requirements.security_group_id_list` ([#19172](https://github.com/hashicorp/terraform-provider-aws/issues/19172))
* resource/aws_sfn_state_machine: Add `tracing_configuration` attribute ([#15434](https://github.com/hashicorp/terraform-provider-aws/issues/15434))
* resource/aws_shield_protection: Add `tags` argument ([#19168](https://github.com/hashicorp/terraform-provider-aws/issues/19168))

BUG FIXES:

* aws_batch_compute_environment: Allow update of just `service_role` for managed compute environments ([#19205](https://github.com/hashicorp/terraform-provider-aws/issues/19205))
* aws_batch_compute_environment: `service_role` argument is optional ([#19205](https://github.com/hashicorp/terraform-provider-aws/issues/19205))
* provider: Prevent `Provider produced inconsistent final plan` errors when lifecycle arguments apply to resource `tags` not known until apply ([#19251](https://github.com/hashicorp/terraform-provider-aws/issues/19251))
* resource/aws_batch_job_definition: Prevent diff with default value of `fargatePlatformConfiguration` ([#19207](https://github.com/hashicorp/terraform-provider-aws/issues/19207))
* resource/aws_lakeformation_permissions: Fix issues related to permissions not being revoked and attempts to revoke non-existent permissions ([#18505](https://github.com/hashicorp/terraform-provider-aws/issues/18505))
* resource/aws_mwaa_environment: Correctly apply `plugins_s3_object_version` change ([#19266](https://github.com/hashicorp/terraform-provider-aws/issues/19266))
* resource/aws_sfn_state_machine: Handle eventual consistency of state machine updates ([#15434](https://github.com/hashicorp/terraform-provider-aws/issues/15434))
* resource/aws_ssoadmin_managed_policy_attachment: Retry attachment/detachment when other permission-set attachment event was not yet propagated, to avoid ConflictException. ([#19216](https://github.com/hashicorp/terraform-provider-aws/issues/19216))

## 3.38.0 (April 30, 2021)

NOTES:

* provider: `default_tags` support generally available to all provider resources that support `tags` with the exception of `aws_autoscaling_group` ([#19084](https://github.com/hashicorp/terraform-provider-aws/issues/19084))

FEATURES:

* **New Data Source:** `aws_cloudformation_type` ([#18579](https://github.com/hashicorp/terraform-provider-aws/issues/18579))
* **New Data Source:** `aws_kms_public_key` ([#18873](https://github.com/hashicorp/terraform-provider-aws/issues/18873))
* **New Data Source:** `aws_resourcegroupstaggingapi_resources` ([#17804](https://github.com/hashicorp/terraform-provider-aws/issues/17804))
* **New Resource:** `aws_cloudformation_type` ([#18579](https://github.com/hashicorp/terraform-provider-aws/issues/18579))
* **New Resource:** `aws_codestarconnections_host` ([#16918](https://github.com/hashicorp/terraform-provider-aws/issues/16918))
* **New Resource:** `aws_macie2_account` ([#19069](https://github.com/hashicorp/terraform-provider-aws/issues/19069))
* **New Resource:** `aws_rds_proxy_endpoint` ([#18881](https://github.com/hashicorp/terraform-provider-aws/issues/18881))
* **New Resource:** `aws_route53_resolver_firewall_rule` ([#18712](https://github.com/hashicorp/terraform-provider-aws/issues/18712))
* **New Resource:** `aws_route53_resolver_firewall_rule_group_association` ([#19164](https://github.com/hashicorp/terraform-provider-aws/issues/19164))
* **New Resource:** `aws_servicecatalog_product` ([#19122](https://github.com/hashicorp/terraform-provider-aws/issues/19122))

ENHANCEMENTS:

* data-source/aws_efs_mount_target: Add `access_point_id`, `file_system_id` arguments ([#18918](https://github.com/hashicorp/terraform-provider-aws/issues/18918))
* data/source_aws_eks_addon: added validation for `cluster_name` ([#19078](https://github.com/hashicorp/terraform-provider-aws/issues/19078))
* data/source_aws_eks_cluster: added validation for `cluster_name` ([#19078](https://github.com/hashicorp/terraform-provider-aws/issues/19078))
* resource/aws_batch_compute_environment: Additional supported value `FARGATE` and `FARGATE_SPOT` for the `type` argument in the `compute_resources` configuration block ([#16819](https://github.com/hashicorp/terraform-provider-aws/issues/16819))
* resource/aws_batch_compute_environment: The `instance_role`, `instance_type` and `min_vcpus` arguments in the `compute_resources` configuration block are now optional ([#16819](https://github.com/hashicorp/terraform-provider-aws/issues/16819))
* resource/aws_batch_compute_environment: The `security_group_ids` and `subnets` arguments in the `compute_resources` configuration block can now be updated in-place for Fargate compute resources ([#16819](https://github.com/hashicorp/terraform-provider-aws/issues/16819))
* resource/aws_batch_job_definition: Add `propagate_tags` argument ([#18336](https://github.com/hashicorp/terraform-provider-aws/issues/18336))
* resource/aws_instance: Make `instance_initiated_shutdown_behavior` also computed, allowing value to be read ([#18880](https://github.com/hashicorp/terraform-provider-aws/issues/18880))
* resource/aws_sns_topic_policy: Add `owner` attribute ([#14123](https://github.com/hashicorp/terraform-provider-aws/issues/14123))
* resource/aws_sns_topic_policy: Add plan time validation to `arn` ([#14123](https://github.com/hashicorp/terraform-provider-aws/issues/14123))

BUG FIXES:

* provider: Prevent `Provider produced inconsistent final plan` errors when resource `tags` are not known until apply ([#18958](https://github.com/hashicorp/terraform-provider-aws/issues/18958))
* resource/aws_batch_job_definition: Treat empty `container_properties.logConfiguration.secretOptions` array as `null` to prevent continual diffs ([#16120](https://github.com/hashicorp/terraform-provider-aws/issues/16120))
* resource/aws_batch_job_queue: Recreate batch job queue if the `name` changes ([#19121](https://github.com/hashicorp/terraform-provider-aws/issues/19121))
* resource/aws_ec2_managed_prefix_list: Prevent `entry` `description` update errors ([#19095](https://github.com/hashicorp/terraform-provider-aws/issues/19095))
* resource/aws_glue_crawler: Allow '/' in `name` argument ([#19160](https://github.com/hashicorp/terraform-provider-aws/issues/19160))
* resource/aws_ram_principal_association: Improve handling of eventual consistency ([#17032](https://github.com/hashicorp/terraform-provider-aws/issues/17032))
* resource/aws_ram_resource_share: Improve handling of eventual consistency ([#17032](https://github.com/hashicorp/terraform-provider-aws/issues/17032))
* resource/aws_ram_resource_share_accepter: Improve handling of eventual consistency ([#17032](https://github.com/hashicorp/terraform-provider-aws/issues/17032))
* resource/aws_vpn_connection: Prevent flipped `tunnel1_*` and `tunnel2_*` ordering when `tunnel1_inside_cidr`, `tunnel1_inside_ipv6_cidr`, or `tunnel1_preshared_key` is configured ([#19077](https://github.com/hashicorp/terraform-provider-aws/issues/19077))

## 3.37.0 (April 16, 2021)

NOTES:

* provider: The HTTP User-Agent header has been reordered so the AWS SDK Go product is last, except when using the TF_APPEND_USER_AGENT environment variable. Environments dependent on the previous User-Agent header ordering may require updates. ([#18855](https://github.com/hashicorp/terraform-provider-aws/issues/18855))

FEATURES:

* **New Data Source:** `aws_eks_addon` ([#16972](https://github.com/hashicorp/terraform-provider-aws/issues/16972))
* **New Resource:** `aws_eks_addon` ([#16972](https://github.com/hashicorp/terraform-provider-aws/issues/16972))
* **New Resource:** `aws_route53_resolver_firewall_domain_list` ([#18558](https://github.com/hashicorp/terraform-provider-aws/issues/18558))
* **New Resource:** `aws_securityhub_insight` ([#18494](https://github.com/hashicorp/terraform-provider-aws/issues/18494))

ENHANCEMENTS:

* resource/aws_servicequotas_service_quota: Add plan time validation to `quota_code` and `service_code` ([#17992](https://github.com/hashicorp/terraform-provider-aws/issues/17992))
* resource/aws_sns_topic: Add `fifo_topic` and `content_based_deduplication` attributes ([#15828](https://github.com/hashicorp/terraform-provider-aws/issues/15828))

BUG FIXES:

* resource/aws_fms_policy: Use API model regular expression for `resource_type` and `resource_type_list` argument plan time validation ([#18600](https://github.com/hashicorp/terraform-provider-aws/issues/18600))
* resource/aws_sqs_queue: Append `.fifo` suffix for Terraform-assigned FIFO queue names ([#17164](https://github.com/hashicorp/terraform-provider-aws/issues/17164))

## 3.36.0 (April 09, 2021)

FEATURES:

* **New Resource:** `aws_cloudfront_key_group` ([#17041](https://github.com/hashicorp/terraform-provider-aws/issues/17041))
* **New Resource:** `aws_ecr_registry_policy` ([#16831](https://github.com/hashicorp/terraform-provider-aws/issues/16831))
* **New Resource:** `aws_ecr_replication_configuration` ([#16853](https://github.com/hashicorp/terraform-provider-aws/issues/16853))
* **New Resource:** `aws_kinesisanalyticsv2_application_snapshot` ([#18056](https://github.com/hashicorp/terraform-provider-aws/issues/18056))
* **New Resource:** `aws_mwaa_environment` ([#16616](https://github.com/hashicorp/terraform-provider-aws/issues/16616))

ENHANCEMENTS:

* data-source/aws_lb_listener: Add `alpn_policy` argument ([#14462](https://github.com/hashicorp/terraform-provider-aws/issues/14462))
* resource/aws_lb_listener: Add `alpn_policy` argument ([#14462](https://github.com/hashicorp/terraform-provider-aws/issues/14462))

BUG FIXES:

* resource/aws_docdb_cluster_parameter_group: Read all user parameters and parameters specified in the configuration. ([#18486](https://github.com/hashicorp/terraform-provider-aws/issues/18486))
* resource/aws_lb_target_group: Handle read-after-create eventual consistency ([#18634](https://github.com/hashicorp/terraform-provider-aws/issues/18634))
* resource/aws_secretsmanager_secret: Handle read-after-create eventual consistency ([#18462](https://github.com/hashicorp/terraform-provider-aws/issues/18462))
* resource/aws_secretsmanager_secret_policy: Handle read-after-create eventual consistency ([#18462](https://github.com/hashicorp/terraform-provider-aws/issues/18462))
* resource/aws_secretsmanager_secret_rotation: Handle read-after-create eventual consistency ([#18462](https://github.com/hashicorp/terraform-provider-aws/issues/18462))
* resource/aws_secretsmanager_secret_version: Handle read-after-create eventual consistency ([#18462](https://github.com/hashicorp/terraform-provider-aws/issues/18462))
* resource/aws_xray_sampling_rule: Change the maximum length of `rule_name` from 128 to 32 ([#18667](https://github.com/hashicorp/terraform-provider-aws/issues/18667))

## 3.35.0 (April 01, 2021)

FEATURES:

* **New Resource:** `aws_cloudwatch_query_definition` ([#17899](https://github.com/hashicorp/terraform-provider-aws/issues/17899))

ENHANCEMENTS:

* data-source/aws_efs_file_system: Add `availability_zone_id` and `availability_zone_name` attributes ([#18319](https://github.com/hashicorp/terraform-provider-aws/issues/18319))
* resource/aws_efs_file_system: Add `availability_zone_id` attribute and `availability_zone_name` argument ([#18319](https://github.com/hashicorp/terraform-provider-aws/issues/18319))
* resource/aws_efs_file_system: Add `number_of_mount_targets`, `size_in_bytes` and `owner_id` attributes ([#17969](https://github.com/hashicorp/terraform-provider-aws/issues/17969))
* resource/aws_lb_target_group: Add preserve_client_ip target attribute support ([#17731](https://github.com/hashicorp/terraform-provider-aws/issues/17731))

BUG FIXES:

* resource/aws_ec2_transit_gateway_route_table_propagation: Wait for enable and disable operations to complete ([#18470](https://github.com/hashicorp/terraform-provider-aws/issues/18470))
* resource/aws_sns_topic_subscription: Enforce lowercase `protocol` argument validation to match API and prevent resource errors ([#18475](https://github.com/hashicorp/terraform-provider-aws/issues/18475))
* resource/aws_sns_topic_subscription: Handle read-after-create eventual consistency ([#18475](https://github.com/hashicorp/terraform-provider-aws/issues/18475))
* resource/aws_synthetics_canary: Handle asynchronous IAM eventual consistency error on creation ([#18404](https://github.com/hashicorp/terraform-provider-aws/issues/18404))
* resource/aws_vpn_gateway_route_propagation: Improve eventual consistency handling and handling of out-of-band resource removal ([#17319](https://github.com/hashicorp/terraform-provider-aws/issues/17319))

## 3.34.0 (March 26, 2021)

NOTES:


FEATURES:

* **New Data Source:** `aws_codestarconnections_connection` ([#18129](https://github.com/hashicorp/terraform-provider-aws/issues/18129))
* **New Resource:** `aws_lightsail_instance_public_ports` ([#8611](https://github.com/hashicorp/terraform-provider-aws/issues/8611))

ENHANCEMENTS:

* resource/aws_fms_admin_account: Extend creation timeout to 10 minutes ([#17596](https://github.com/hashicorp/terraform-provider-aws/issues/17596))
* resource/aws_sns_topic_subscription: Add plan time validation for `subscription_role_arn` and `topic_arn` ([#14101](https://github.com/hashicorp/terraform-provider-aws/issues/14101))

BUG FIXES:

* resource/aws_mq_configuration: Add `ldap` as an `authentication_strategy` and `RabbitMQ` as an `engine_type` ([#18070](https://github.com/hashicorp/terraform-provider-aws/issues/18070))
* resource/aws_sns_topic_subscription: recreate subscription if topic is deleted ([#14101](https://github.com/hashicorp/terraform-provider-aws/issues/14101))

## 3.33.0 (March 18, 2021)

NOTES:

* provider: New `default_tags` argument as a public preview for applying tags across all resources under a provider. Support for the functionality must be added to individual resources in the codebase and is only implemented for the `aws_subnet` and `aws_vpc` resources at this time. Until a general availability announcement, no compatibility promises are made with these provider arguments and their functionality. ([#17974](https://github.com/hashicorp/terraform-provider-aws/issues/17974))

FEATURES:

* **New Data Source:** `aws_ec2_transit_gateway_route_tables` ([#17589](https://github.com/hashicorp/terraform-provider-aws/issues/17589))
* **New Data Source:** `aws_kinesis_stream_consumer` ([#17149](https://github.com/hashicorp/terraform-provider-aws/issues/17149))
* **New Resource:** `aws_kinesis_stream_consumer` ([#17149](https://github.com/hashicorp/terraform-provider-aws/issues/17149))

ENHANCEMENTS:

* provider: Add `default_tags` argument (in public preview, see note above) ([#17974](https://github.com/hashicorp/terraform-provider-aws/issues/17974))

BUG FIXES:

* resource/aws_globalaccelerator_accelerator: Correct length for `name` attribute validation ([#17985](https://github.com/hashicorp/terraform-provider-aws/issues/17985))
* resource/aws_lakeformation_permissions: Properly serialize SELECT permission for `permissions` and `permissions_with_grant_option` fields ([#18203](https://github.com/hashicorp/terraform-provider-aws/issues/18203))

## 3.32.0 (March 12, 2021)

FEATURES:

* **New Data Source:** `aws_acmpca_certificate` ([#10213](https://github.com/hashicorp/terraform-provider-aws/issues/10213))
* **New Resource:** `aws_acmpca_certificate` ([#10213](https://github.com/hashicorp/terraform-provider-aws/issues/10213))
* **New Resource:** `aws_acmpca_certificate_authority_certificate` ([#17850](https://github.com/hashicorp/terraform-provider-aws/issues/17850))

ENHANCEMENTS:

* resource/aws_mq_broker: Add RabbitMQ as option for `engine_type`, and new arguments `authentication_strategy`, `ldap_server_metadata`, and `storage_type`. Improve handling of eventual consistency. ([#16108](https://github.com/hashicorp/terraform-provider-aws/issues/16108))
* resource/aws_mq_broker: Support updating broker engine version without recreating broker ([#12758](https://github.com/hashicorp/terraform-provider-aws/issues/12758))

BUG FIXES:


## 3.31.0 (March 04, 2021)

FEATURES:

* **New Resource:** `aws_route53_hosted_zone_dnssec` ([#17474](https://github.com/hashicorp/terraform-provider-aws/issues/17474))

ENHANCEMENTS:

* provider: Support automatic region validation for `ap-northeast-3` ([#17934](https://github.com/hashicorp/terraform-provider-aws/issues/17934))
* resource/aws_globalaccelerator_accelerator: Add plan time validation to `name`, `flow_logs_s3_bucket` and `flow_logs_s3_prefix` attributes ([#17739](https://github.com/hashicorp/terraform-provider-aws/issues/17739))

BUG FIXES:

* resource/aws_acm_certificate: Trigger resource recreation with `VALIDATION_TIMED_OUT` status ([#17869](https://github.com/hashicorp/terraform-provider-aws/issues/17869))
* resource/aws_globalaccelerator_accelerator: Allow update of flow log attribute for active flow logs ([#17739](https://github.com/hashicorp/terraform-provider-aws/issues/17739))
* resource/aws_neptune_cluster_instance: Add "storage-optimization" to Neptune cluster instance create/update pending states ([#17901](https://github.com/hashicorp/terraform-provider-aws/issues/17901))
* resource/aws_neptune_cluster_parameter_group: Correctly update resource by `id` ([#17872](https://github.com/hashicorp/terraform-provider-aws/issues/17872))

## 3.30.0 (February 26, 2021)

FEATURES:

* **New Data Source:** `aws_apigatewayv2_api` ([#13883](https://github.com/hashicorp/terraform-provider-aws/issues/13883))
* **New Data Source:** `aws_apigatewayv2_apis` ([#13883](https://github.com/hashicorp/terraform-provider-aws/issues/13883))
* **New Resource:** `aws_cognito_user_pool_ui_customization` ([#8114](https://github.com/hashicorp/terraform-provider-aws/issues/8114))
* **New Resource:** `aws_ecrpublic_repository` ([#16865](https://github.com/hashicorp/terraform-provider-aws/issues/16865))
* **New Resource:** `aws_sagemaker_app` ([#17251](https://github.com/hashicorp/terraform-provider-aws/issues/17251))

ENHANCEMENTS:

* provider: Add validation for `role_arn`, `policy_arns`, and `policy` ([#12642](https://github.com/hashicorp/terraform-provider-aws/issues/12642))
* resource/aws_gamelift_build: Support all valid operating system values ([#17764](https://github.com/hashicorp/terraform-provider-aws/issues/17764))
* resource/aws_sns_topic_subscription: Add `email`, `email-json`, and `firehose` to protocol values. Add `subscription_role_arn` argument for Firehose support. Add `confirmation_was_authenticated`, `owner_id`, and `pending_confirmation` attributes. ([#14923](https://github.com/hashicorp/terraform-provider-aws/issues/14923))

BUG FIXES:

* provider: Underlying Terraform Plugin SDK update to ensure data source errors include configuration source (file and line) ([#17801](https://github.com/hashicorp/terraform-provider-aws/issues/17801))
* resource/aws_glue_trigger: Support starting ON_DEMAND triggers via `enabled` flag. ([#17488](https://github.com/hashicorp/terraform-provider-aws/issues/17488))
* resource/aws_sns_topic_subscription: Fix to avoid `delivery_policy` always showing diff. ([#14255](https://github.com/hashicorp/terraform-provider-aws/issues/14255))

## 3.29.1 (February 23, 2021)

ENHANCEMENTS:


BUG FIXES:

* resource/aws_instance: Prevent error with `iam_instance_profile` containing additional forward slashes from path ([#17734](https://github.com/hashicorp/terraform-provider-aws/issues/17734))
* resource/aws_lb_target_group_attachment: Retry InvalidTarget errors when creating ([#8538](https://github.com/hashicorp/terraform-provider-aws/issues/8538))
* resource/aws_synthetics_canary: Fix Canary Update when in running state ([#17704](https://github.com/hashicorp/terraform-provider-aws/issues/17704))

## 3.29.0 (February 19, 2021)

FEATURES:

* **New Resource:** `aws_cloudwatch_event_archive` ([#17270](https://github.com/hashicorp/terraform-provider-aws/issues/17270))
* **New Resource:** `aws_elasticache_global_replication_group` ([#15885](https://github.com/hashicorp/terraform-provider-aws/issues/15885))
* **New Resource:** `aws_s3_object_copy` ([#15461](https://github.com/hashicorp/terraform-provider-aws/issues/15461))
* **New Resource:** `aws_securityhub_invite_accepter` ([#12684](https://github.com/hashicorp/terraform-provider-aws/issues/12684))

ENHANCEMENTS:

* resource/aws_dx_gateway_association: Changes to `proposal_id` do not force resource recreation ([#12482](https://github.com/hashicorp/terraform-provider-aws/issues/12482))
* resource/ses_receipt_rule_set: Add `arn` attribute ([#17611](https://github.com/hashicorp/terraform-provider-aws/issues/17611))
* resource/ses_receipt_rule_set: Add plan time validation to `name` ([#17611](https://github.com/hashicorp/terraform-provider-aws/issues/17611))

BUG FIXES:

* resource/aws_fms_policy: Update `resource_type_list` plan-time validation to include `AWS::EC2::VPC`. ([#17595](https://github.com/hashicorp/terraform-provider-aws/issues/17595))
* resource/aws_lb_cookie_stickiness_policy: Allow zero value for `cookie_expiration_period` ([#17204](https://github.com/hashicorp/terraform-provider-aws/issues/17204))
* resource/aws_lb_listener_certificate: Prevent resource ID parsing error with IAM Server Certificate names containing underscores ([#17645](https://github.com/hashicorp/terraform-provider-aws/issues/17645))
* resource/aws_lb_target_group: Use gRPC matcher when using gRPC protocol ([#17534](https://github.com/hashicorp/terraform-provider-aws/issues/17534))

## 3.28.0 (February 12, 2021)

FEATURES:

* **New Data Source:** `aws_cloudfront_cache_policy` ([#17336](https://github.com/hashicorp/terraform-provider-aws/issues/17336))
* **New Resource:** `aws_cloudfront_cache_policy` ([#17336](https://github.com/hashicorp/terraform-provider-aws/issues/17336))
* **New Resource:** `aws_cloudfront_realtime_log_config` ([#14974](https://github.com/hashicorp/terraform-provider-aws/issues/14974))
* **New Resource:** `aws_config_conformance_pack` ([#17313](https://github.com/hashicorp/terraform-provider-aws/issues/17313))
* **New Resource:** `aws_sagemaker_model_package_group` ([#17366](https://github.com/hashicorp/terraform-provider-aws/issues/17366))
* **New Resource:** `aws_securityhub_organization_admin_account` ([#17501](https://github.com/hashicorp/terraform-provider-aws/issues/17501))
* **New Resource:** `aws_synthetics_canary` ([#13140](https://github.com/hashicorp/terraform-provider-aws/issues/13140))

ENHANCEMENTS:

* data-source/aws_customer_gateway: Add `device_name` attribute ([#14786](https://github.com/hashicorp/terraform-provider-aws/issues/14786))
* provider: Add terraform-provider-aws version to HTTP User-Agent header ([#17486](https://github.com/hashicorp/terraform-provider-aws/issues/17486))
* resource/aws_budgets_budget: Add `arn` attribute ([#13139](https://github.com/hashicorp/terraform-provider-aws/issues/13139))
* resource/aws_budgets_budget: Add plan time validation for `budget_type`, `time_unit`, and `subscriber_sns_topic_arns` arguments ([#13139](https://github.com/hashicorp/terraform-provider-aws/issues/13139))
* resource/aws_customer_gateway: Add `device_name` argument ([#14786](https://github.com/hashicorp/terraform-provider-aws/issues/14786))
* resource/aws_ec2_capacity_reservation: Add `owner_id` attribute ([#17129](https://github.com/hashicorp/terraform-provider-aws/issues/17129))
* resource/aws_ec2_traffic_mirror_filter: Add `arn` attribute. ([#13948](https://github.com/hashicorp/terraform-provider-aws/issues/13948))
* resource/aws_ec2_traffic_mirror_filter_rule: Add arn attribute. ([#13949](https://github.com/hashicorp/terraform-provider-aws/issues/13949))
* resource/aws_ec2_traffic_mirror_filter_rule: Add plan time validation to `destination_port_range.from_port`,
`destination_port_range.to_port`, `source_port_range.from_port`, and `source_port_range.to_port`. ([#13949](https://github.com/hashicorp/terraform-provider-aws/issues/13949))
`notifications.completed`, `notifications.error`, `notifications.progressing`, `notifications.warning`,
`thumbnail_config.storage_class`, `thumbnail_config_permissions.access`, `thumbnail_config_permissions.grantee_type` ([#13973](https://github.com/hashicorp/terraform-provider-aws/issues/13973))
* resource/aws_fms_policy: Allow use of `resource_type` or `resource_type_list` attributes ([#17418](https://github.com/hashicorp/terraform-provider-aws/issues/17418))
* resource/aws_lb_listener_certificate: Add import support ([#16474](https://github.com/hashicorp/terraform-provider-aws/issues/16474))
* resource/aws_licensemanager_license_configuration: Add `arn` and `owner_account_id` attributes ([#17160](https://github.com/hashicorp/terraform-provider-aws/issues/17160))
* resource/aws_sns_topic_subscription: Add `redrive_policy` argument ([#11770](https://github.com/hashicorp/terraform-provider-aws/issues/11770))
`approved_patches`, `rejected_patches`, `approval_rule.approve_after_days`, `approval_rule.patch_filter.key`, and `approval_rule.patch_filter.values`. ([#11772](https://github.com/hashicorp/terraform-provider-aws/issues/11772))

BUG FIXES:

* resource/aws_glue_catalog_database: Use Catalog Id when deleting Databases. ([#17489](https://github.com/hashicorp/terraform-provider-aws/issues/17489))
* resource/aws_instance: Fix use of `throughput` and `iops` for `gp3` volumes at the same time ([#17380](https://github.com/hashicorp/terraform-provider-aws/issues/17380))

## 3.27.0 (February 05, 2021)

FEATURES:

* **New Resource:** `aws_ec2_transit_gateway_prefix_list_reference` ([#16823](https://github.com/hashicorp/terraform-provider-aws/issues/16823))
* **New Resource:** `aws_route53_key_signing_key` ([#16840](https://github.com/hashicorp/terraform-provider-aws/issues/16840))
* **New Resource:** `aws_cloudfront_origin_request_policy` ([#17342](https://github.com/hashicorp/terraform-provider-aws/issues/17342))
* **New Data Source:** `aws_cloudfront_origin_request_policy` ([#17342](https://github.com/hashicorp/terraform-provider-aws/issues/17342))

ENHANCEMENTS:

* resource/aws_glue_catalog_table: Adds support for specifying schema from schema registry. ([#17335](https://github.com/hashicorp/terraform-provider-aws/issues/17335))
* resource/aws_lb: Add `ipv6_address` attribute ([#17229](https://github.com/hashicorp/terraform-provider-aws/issues/17229))
* resource/aws_sfn_state_machine: Add support for `EXPRESS` state machine `type` ([#12249](https://github.com/hashicorp/terraform-provider-aws/issues/12249))
* resource/aws_lb_target_group: Add `protocol_version` attribute ([#17260](https://github.com/hashicorp/terraform-provider-aws/issues/17260))

BUG FIXES:

* data-source/aws_partition: Correct `reverse_dns_prefix` value in AWS China, C2S, and SC2S ([#17142](https://github.com/hashicorp/terraform-provider-aws/issues/17142))
* provider: Only validate AWS shared configuration profile SSO configuration when attempting to use SSO cached credentials ([#17469](https://github.com/hashicorp/terraform-provider-aws/issues/17469))
* resource/aws_glue_crawler: Use standard retry timeout for IAM eventual consistency and retry on LakeFormation permissions errors ([#17256](https://github.com/hashicorp/terraform-provider-aws/issues/17256))
* resource/aws_glue_partition: Fix `partition_values` to preserve order. ([#17344](https://github.com/hashicorp/terraform-provider-aws/issues/17344))
* resource/aws_ssoadmin_permission_set: Properly update resource with `relay_state` argument ([#17423](https://github.com/hashicorp/terraform-provider-aws/issues/17423))

## 3.26.0 (January 28, 2021)

NOTES:


FEATURES:

* **New Data Source:** `aws_imagebuilder_image` ([#16710](https://github.com/hashicorp/terraform-provider-aws/issues/16710))
* **New Resource:** `aws_imagebuilder_image` ([#16710](https://github.com/hashicorp/terraform-provider-aws/issues/16710))
* **New Resource:** `aws_prometheus_workspace` ([#16882](https://github.com/hashicorp/terraform-provider-aws/issues/16882))
* **New Resource:** `aws_sagemaker_app_image_config` ([#17221](https://github.com/hashicorp/terraform-provider-aws/issues/17221))

ENHANCEMENTS:

* provider: Support AWS Single-Sign On (SSO) cached credentials ([#17340](https://github.com/hashicorp/terraform-provider-aws/issues/17340))

BUG FIXES:

* data-source/aws_elb_hosted_zone_id: Correct values for `cn-north-1` and `cn-northwest-1` regions ([#17226](https://github.com/hashicorp/terraform-provider-aws/issues/17226))
* data-source/aws_lb_listener: Prevent error when retrieving a listener whose default action contains weighted target groups ([#17238](https://github.com/hashicorp/terraform-provider-aws/issues/17238))

## 3.25.0 (January 22, 2021)

NOTES


FEATURES

* **New Resource:** `aws_backup_global_settings` ([#16475](https://github.com/hashicorp/terraform-provider-aws/issues/16475))
* **New Resource:** `aws_sagemaker_feature_group` ([#16728](https://github.com/hashicorp/terraform-provider-aws/issues/16728))
* **New Resource:** `aws_sagemaker_image_version` ([#17141](https://github.com/hashicorp/terraform-provider-aws/issues/17141))
* **New Resource:** `aws_sagemaker_user_profile` ([#17123](https://github.com/hashicorp/terraform-provider-aws/issues/17123))

ENHANCEMENTS


BUG FIXES

* provider: Fix error messages for missing required blocks not including the block name ([#17211](https://github.com/hashicorp/terraform-provider-aws/issues/17211))
* provider: Prevent panic when sending Ctrl-C (SIGINT) to Terraform ([#17211](https://github.com/hashicorp/terraform-provider-aws/issues/17211))
* resource/aws_lakeformation_data_lake_settings: Avoid unnecessary resource cycling ([#17189](https://github.com/hashicorp/terraform-provider-aws/issues/17189))
* resource/aws_lakeformation_permissions: Handle resources with multiple permissions ([#17189](https://github.com/hashicorp/terraform-provider-aws/issues/17189))


## 3.24.1 (January 15, 2021)

BUG FIXES

* data-source/instance: Fix EBS and root block device tags issue with "Invalid address to set" ([#17136](https://github.com/hashicorp/terraform-provider-aws/issues/17136))

## 3.24.0 (January 14, 2021)

FEATURES

* **New Data Source:** `aws_api_gateway_domain_name` ([#12489](https://github.com/hashicorp/terraform-provider-aws/issues/12489))
* **New Data Source:** `aws_identitystore_group` ([#15322](https://github.com/hashicorp/terraform-provider-aws/issues/15322))
* **New Data Source:** `aws_identitystore_user` ([#15322](https://github.com/hashicorp/terraform-provider-aws/issues/15322))
* **New Resource:** `aws_cloudwatch_composite_alarm` ([#15023](https://github.com/hashicorp/terraform-provider-aws/issues/15023))
* **New Resource:** `aws_fms_policy` ([#9594](https://github.com/hashicorp/terraform-provider-aws/issues/9594))
* **New Resource:** `aws_route53_resolver_dnssec_config` ([#17012](https://github.com/hashicorp/terraform-provider-aws/issues/17012))
* **New Resource:** `aws_sagemaker_domain` ([#16077](https://github.com/hashicorp/terraform-provider-aws/issues/16077))
* **New Resource:** `aws_ssoadmin_account_assignment` ([#15322](https://github.com/hashicorp/terraform-provider-aws/issues/15322))

ENHANCEMENTS

* data-source/aws_workspaces_directory: Add access properties ([#16688](https://github.com/hashicorp/terraform-provider-aws/issues/16688))
* resource/aws_globalaccelerator_accelerator: Add custom timeouts ([#17112](https://github.com/hashicorp/terraform-provider-aws/issues/17112))
* resource/aws_globalaccelerator_endpoint_group: Add custom timeouts ([#17112](https://github.com/hashicorp/terraform-provider-aws/issues/17112))
* resource/aws_globalaccelerator_endpoint_listener: Add custom timeouts ([#17112](https://github.com/hashicorp/terraform-provider-aws/issues/17112))
* resource/aws_instance: Add `tags` parameter to `root_block_device`, `ebs_block_device` blocks.([#15474](https://github.com/hashicorp/terraform-provider-aws/issues/15474))
* resource/aws_workspaces_directory: Add access properties ([#16688](https://github.com/hashicorp/terraform-provider-aws/issues/16688))

BUG FIXES

* resource/aws_instance: Prevent `volume_tags` from improperly interfering with `tags` in `aws_ebs_volume` ([#15474](https://github.com/hashicorp/terraform-provider-aws/issues/15474))
* service/apigateway: All operations will now automatically retry on `ConflictException: Unable to complete operation due to concurrent modification. Please try again later.` errors.

## 3.23.0 (January 08, 2021)

FEATURES

* **New Data Source:** `aws_ssoadmin_instances` ([#15808](https://github.com/hashicorp/terraform-provider-aws/issues/15808))
* **New Data Source:** `aws_ssoadmin_permission_set` ([#15808](https://github.com/hashicorp/terraform-provider-aws/issues/15808))
* **New Resource:** `aws_sagemaker_image` ([#16082](https://github.com/hashicorp/terraform-provider-aws/issues/16082))
* **New Resource:** `aws_ssoadmin_managed_policy_attachment` ([#15808](https://github.com/hashicorp/terraform-provider-aws/issues/15808))
* **New Resource:** `aws_ssoadmin_permission_set` ([#15808](https://github.com/hashicorp/terraform-provider-aws/issues/15808))
* **New Resource:** `aws_ssoadmin_permission_set_inline_policy` ([#15808](https://github.com/hashicorp/terraform-provider-aws/issues/15808))

ENHANCEMENTS

* resource/aws_glue_crawler: add support for `lineage_configuration` and `recrawl_policy` ([#16714](https://github.com/hashicorp/terraform-provider-aws/issues/16714))
* resource/aws_glue_crawler: add plan time validations to `name`, `description` and `table_prefix` ([#16714](https://github.com/hashicorp/terraform-provider-aws/issues/16714))

BUG FIXES

* resource/aws_glue_catalog_table: Glue table partition keys should be set to empty list instead of being unset ([#16727](https://github.com/hashicorp/terraform-provider-aws/issues/16727))

## 3.22.0 (December 18, 2020)

FEATURES

* **New Data Source:** `aws_ec2_managed_prefix_list` ([#16738](https://github.com/hashicorp/terraform-provider-aws/issues/16738))
* **New Data Source:** `aws_lakeformation_data_lake_settings` ([#13250](https://github.com/hashicorp/terraform-provider-aws/issues/13250))
* **New Data Source:** `aws_lakeformation_permissions` ([#13396](https://github.com/hashicorp/terraform-provider-aws/issues/13396))
* **New Data Source:** `aws_lakeformation_resource` ([#13396](https://github.com/hashicorp/terraform-provider-aws/issues/13396))
* **New Resource:** `aws_codestarconnections_connection` ([#15990](https://github.com/hashicorp/terraform-provider-aws/issues/15990))
* **New Resource:** `aws_ec2_managed_prefix_list` ([#14068](https://github.com/hashicorp/terraform-provider-aws/issues/14068))
* **New Resource:** `aws_lakeformation_data_lake_settings` ([#13250](https://github.com/hashicorp/terraform-provider-aws/issues/13250))
* **New Resource:** `aws_lakeformation_permissions` ([#13396](https://github.com/hashicorp/terraform-provider-aws/issues/13396))
* **New Resource:** `aws_lakeformation_resource` ([#13267](https://github.com/hashicorp/terraform-provider-aws/issues/13267))

ENHANCEMENTS

* data-source/aws_availability_zone: Add `parent_zone_id`, `parent_zone_name`, and `zone_type` attributes (additional support for Local and Wavelength Zones) ([#16770](https://github.com/hashicorp/terraform-provider-aws/issues/16770))
* data-source/aws_instance: Add `enclave_options` attribute (Nitro Enclaves) ([#16361](https://github.com/hashicorp/terraform-provider-aws/issues/16361))
* data-source/aws_instance: Add `ebs_block_device` and `root_block_device` configuration block `throughput` attribute ([#16620](https://github.com/hashicorp/terraform-provider-aws/issues/16620))
* resource/aws_instance: Add `enclave_options` configuration block (Nitro Enclaves) ([#16361](https://github.com/hashicorp/terraform-provider-aws/issues/16361))
* resource/aws_instance: Add `ebs_block_device` and `root_block_device` configuration block `throughput` attribute ([#16620](https://github.com/hashicorp/terraform-provider-aws/issues/16620))
* resource/aws_vpn_connection: Add support for VPN tunnel options and enable acceleration, DPDTimeoutAction, StartupAction, local/remote IPv4/IPv6 network CIDR and tunnel inside IP version. ([#14740](https://github.com/hashicorp/terraform-provider-aws/issues/14740))

BUG FIXES

* data-source/aws_ec2_coip_pools: Ensure all results from large environments are returned ([#16669](https://github.com/hashicorp/terraform-provider-aws/issues/16669))
* data-source/aws_ec2_local_gateways: Ensure all results from large environments are returned ([#16669](https://github.com/hashicorp/terraform-provider-aws/issues/16669))
* data-source/aws_ec2_local_gateway_route_tables: Ensure all results from large environments are returned ([#16669](https://github.com/hashicorp/terraform-provider-aws/issues/16669))
* data-source/aws_ec2_local_gateway_virtual_interface_groups: Ensure all results from large environments are returned ([#16669](https://github.com/hashicorp/terraform-provider-aws/issues/16669))
* data-source/aws_prefix_list: Using `name` argument no longer overrides other arguments ([#16739](https://github.com/hashicorp/terraform-provider-aws/issues/16739))
* resource/aws_main_route_table_association: Prevent crash on creation when VPC main route table association is not found ([#16680](https://github.com/hashicorp/terraform-provider-aws/issues/16680))
* resource/aws_workspaces_workspace: Prevent panic from terminated WorkSpace ([#16692](https://github.com/hashicorp/terraform-provider-aws/issues/16692))

## 3.21.0 (December 11, 2020)

NOTES


FEATURES

* **New Resource:** `aws_ec2_carrier_gateway` ([#16252](https://github.com/hashicorp/terraform-provider-aws/issues/16252))
* **New Resource:** `aws_glue_schema` ([#16612](https://github.com/hashicorp/terraform-provider-aws/issues/16612))

ENHANCEMENTS

* resource/aws_workspaces_workspace: Add failed request error code along with message ([#16459](https://github.com/hashicorp/terraform-provider-aws/issues/16459))

BUG FIXES

* data-source/aws_customer_gateway: Prevent missing `id` attribute when not configured as argument ([#16667](https://github.com/hashicorp/terraform-provider-aws/issues/16667))
* data-source/aws_ec2_transit_gateway: Prevent missing `id` attribute when not configured as argument ([#16667](https://github.com/hashicorp/terraform-provider-aws/issues/16667))
* data-source/aws_ec2_transit_gateway_peering_attachment: Prevent missing `id` attribute when not configured as argument ([#16667](https://github.com/hashicorp/terraform-provider-aws/issues/16667))
* data-source/aws_ec2_transit_gateway_route_table: Prevent missing `id` attribute when not configured as argument ([#16667](https://github.com/hashicorp/terraform-provider-aws/issues/16667))
* data-source/aws_ec2_transit_gateway_vpc_attachment: Prevent missing `id` attribute when not configured as argument ([#16667](https://github.com/hashicorp/terraform-provider-aws/issues/16667))
* resource/aws_workspaces_directory: Fix empty custom_security_group_id & default_ou ([#16589](https://github.com/hashicorp/terraform-provider-aws/issues/16589))

## 3.20.0 (December 03, 2020)

ENHANCEMENTS


BUG FIXES

* resource/aws_ec2_client_vpn_network_association: Increase associate and disassociate timeouts from 10min to 30min ([#16522](https://github.com/hashicorp/terraform-provider-aws/issues/16522))
* resource/aws_instance: Automatically retry instance restart on eventual consistency error during `instance_type` in-place update ([#16443](https://github.com/hashicorp/terraform-provider-aws/issues/16443))

## 3.19.0 (December 01, 2020)

FEATURES

* **New Resource:** `aws_glue_registry` ([#16418](https://github.com/hashicorp/terraform-provider-aws/issues/16418))

ENHANCEMENTS


BUG FIXES

* resource/aws_fsx_windows_file_system: Prevent potential panics, unexpected errors, and use correct operation timeout on update ([#16488](https://github.com/hashicorp/terraform-provider-aws/issues/16488))

## 3.18.0 (November 25, 2020)

FEATURES

* **New Data Source:** `aws_imagebuilder_image_pipeline` ([#16299](https://github.com/hashicorp/terraform-provider-aws/issues/16299))
* **New Data Source:** `aws_imagebuilder_image_recipe` ([#16218](https://github.com/hashicorp/terraform-provider-aws/issues/16218))
* **New Data Source:** `aws_serverlessrepository_application` ([#15874](https://github.com/hashicorp/terraform-provider-aws/issues/15874))
* **New Resource:** `aws_backup_region_settings` ([#16114](https://github.com/hashicorp/terraform-provider-aws/issues/16114))
* **New Resource:** `aws_imagebuilder_image_pipeline` ([#16299](https://github.com/hashicorp/terraform-provider-aws/issues/16299))
* **New Resource:** `aws_imagebuilder_image_recipe` ([#16218](https://github.com/hashicorp/terraform-provider-aws/issues/16218))
* **New Resource:** `aws_msk_scram_secret_association` ([#15302](https://github.com/hashicorp/terraform-provider-aws/issues/15302))
* **New Resource:** `aws_networkfirewall_resource_policy` ([#16279](https://github.com/hashicorp/terraform-provider-aws/issues/16279))
* **New Resource:** `aws_serverlessrepository_stack` ([#15874](https://github.com/hashicorp/terraform-provider-aws/issues/15874))

ENHANCEMENTS


BUG FIXES

* resource/aws_lb_listener: Mark `port` argument as optional and only default `protocol` argument to `HTTP` for Application Load Balancers (Support Gateway Load Balancer) ([#16306](https://github.com/hashicorp/terraform-provider-aws/issues/16306))

## 3.17.0 (November 24, 2020)

FEATURES

* **New Data Source:** `aws_lambda_code_signing_config` ([#16384](https://github.com/hashicorp/terraform-provider-aws/issues/16384))
* **New Data Source:** `aws_signer_signing_job` ([#16383](https://github.com/hashicorp/terraform-provider-aws/issues/16383))
* **New Data Source:** `aws_signer_signing_profile` ([#16383](https://github.com/hashicorp/terraform-provider-aws/issues/16383))
* **New Resource:** `aws_lambda_code_signing_config` ([#16384](https://github.com/hashicorp/terraform-provider-aws/issues/16384))
* **New Resource:** `aws_signer_signing_job` ([#16383](https://github.com/hashicorp/terraform-provider-aws/issues/16383))
* **New Resource:** `aws_signer_signing_profile` ([#16383](https://github.com/hashicorp/terraform-provider-aws/issues/16383))
* **New Resource:** `aws_signer_signing_profile_permission` ([#16383](https://github.com/hashicorp/terraform-provider-aws/issues/16383))

ENHANCEMENTS

* resource/aws_accessanalyzer_analyzer: Adds plan time validation to `analyzer_name` ([#16265](https://github.com/hashicorp/terraform-provider-aws/issues/16265))
* resource/aws_accessanalyzer_analyzer: Adds plan time validation to `analyzer_name` ([#16265](https://github.com/hashicorp/terraform-provider-aws/issues/16265))
* resource/aws_fsx_windows_file_system: Support updating `throughput_capacity` and `storage_capacity` ([#15582](https://github.com/hashicorp/terraform-provider-aws/issues/15582))
* resource/aws_glue_catalog_table: Add partition index support ([#16194](https://github.com/hashicorp/terraform-provider-aws/issues/16194))
* resource/aws_workspaces_directory: Allows assigning IP group ([#14451](https://github.com/hashicorp/terraform-provider-aws/issues/14451))

BUG FIXES

* resource/aws_fsx_windows_file_system: Update the default creation timeout from 30 to 45 minutes ([#16363](https://github.com/hashicorp/terraform-provider-aws/issues/16363))
* resource/aws_lb: Fix `enable_cross_zone_load_balancing` argument handling with Gateway Load Balancers ([#16314](https://github.com/hashicorp/terraform-provider-aws/issues/16314))

## 3.16.0 (November 18, 2020)

* **New Data Source:** `aws_imagebuilder_component` ([#16159](https://github.com/hashicorp/terraform-provider-aws/issues/16159))
* **New Data Source:** `aws_imagebuilder_distribution_configuration` ([#16180](https://github.com/hashicorp/terraform-provider-aws/issues/16180))
* **New Data Source:** `aws_imagebuilder_infrastructure_configuration` ([#16186](https://github.com/hashicorp/terraform-provider-aws/issues/16186))
* **New Resource:** `aws_api_gateway_rest_api_policy` ([#13619](https://github.com/hashicorp/terraform-provider-aws/issues/13619))
* **New Resource:** `aws_backup_vault_policy` ([#16112](https://github.com/hashicorp/terraform-provider-aws/issues/16112))
* **New Resource:** `aws_glue_dev_endpoint` ([#7895](https://github.com/hashicorp/terraform-provider-aws/pull/7895))
* **New Resource:** `aws_imagebuilder_component` ([#16159](https://github.com/hashicorp/terraform-provider-aws/issues/16159))
* **New Resource:** `aws_imagebuilder_distribution_configuration` ([#16180](https://github.com/hashicorp/terraform-provider-aws/issues/16180))
* **New Resource:** `aws_imagebuilder_infrastructure_configuration` ([#16186](https://github.com/hashicorp/terraform-provider-aws/issues/16186))
* **New Resource:** `aws_networkfirewall_firewall` ([#16277](https://github.com/hashicorp/terraform-provider-aws/issues/16277))
* **New Resource:** `aws_networkfirewall_firewall_policy` ([#16277](https://github.com/hashicorp/terraform-provider-aws/issues/16277))
* **New Resource:** `aws_networkfirewall_logging_configuration` ([#16277](https://github.com/hashicorp/terraform-provider-aws/issues/16277))
* **New Resource:** `aws_networkfirewall_rule_group` ([#16277](https://github.com/hashicorp/terraform-provider-aws/issues/16277))

ENHANCEMENTS

* resource/aws_globalaccelerator_endpoint_group: Add `arn` and `port_override` attributes ([#16121](https://github.com/hashicorp/terraform-provider-aws/issues/16121))
* resource/aws_glue_catalog_table: Add support for `parameters` argument to `storage_descriptor.columns` block ([#16052](https://github.com/hashicorp/terraform-provider-aws/issues/16052))
* resource/aws_glue_catalog_table: Add plan time validation for `description`, `name`, `partition_keys.name`, `partition_keys.comment`, `partition_keys.type`, `retention`, `view_original_text`, `view_expanded_text`, `storage_descriptor.name`, `storage_descriptor.comment`, `storage_descriptor.type`, `storage_descriptor.bucket_columns`, `storage_descriptor.ser_de_info.name`, `storage_descriptor.skewed_info.skewed_column_names`, `storage_descriptor.sort_columns.column`, `storage_descriptor.sort_columns.sort_order` ([#16052](https://github.com/hashicorp/terraform-provider-aws/issues/16052))
* resource_aws_storagegateway_nfs_file_share: Add `file_share_name` argument ([#16072](https://github.com/hashicorp/terraform-provider-aws/issues/16072))

BUG FIXES


## 3.15.0 (November 12, 2020)

ENHANCEMENTS

* data-source/aws_ec2_transit_gateway_route_table: Add `arn` attribute ([#13921](https://github.com/hashicorp/terraform-provider-aws/issues/13921))
* data-source/aws_ec2_transit_gateway_vpc_attachment: Add `appliance_mode_support` attribute ([#16159](https://github.com/hashicorp/terraform-provider-aws/issues/16159))
* resource/aws_default_route_table: Add `route` configuration block `vpc_endpoint_id` argument ([#16131](https://github.com/hashicorp/terraform-provider-aws/issues/16131))
* resource/aws_ec2_transit_gateway: Support in-place updates for most arguments ([#15556](https://github.com/hashicorp/terraform-provider-aws/issues/15556))
* resource/aws_ec2_transit_gateway_route_table: Add `arn` attribute ([#13921](https://github.com/hashicorp/terraform-provider-aws/issues/13921))
* resource/aws_ec2_transit_gateway_vpc_attachment: Add `appliance_mode_support` argument ([#16159](https://github.com/hashicorp/terraform-provider-aws/issues/16159))
* resource/aws_ec2_transit_gateway_vpc_attachment_accepter: Add `appliance_mode_support` attribute ([#16159](https://github.com/hashicorp/terraform-provider-aws/issues/16159))
* resource/aws_lb: Support `load_balancer_type` argument value of `gateway` ([#16131](https://github.com/hashicorp/terraform-provider-aws/issues/16131))
* resource/aws_lb_target_group: Support `protocol` argument value of `GENEVE` ([#16131](https://github.com/hashicorp/terraform-provider-aws/issues/16131))
* resource/aws_workspaces_workspace: Add configurable timeouts ([#15479](https://github.com/hashicorp/terraform-provider-aws/issues/15479))

BUG FIXES

* resource/aws_lb_target_group: Allow invalid configurations that were allowed prior to 3.10. ([#15613](https://github.com/hashicorp/terraform-provider-aws/issues/15613))
* resource/aws_workspaces_bundle: Fix empty (private) owner ([#14535](https://github.com/hashicorp/terraform-provider-aws/issues/14535))

## 3.14.1 (November 06, 2020)

BUG FIXES


## 3.14.0 (November 06, 2020)

FEATURES

* **New Data Source:** `aws_route53_resolver_endpoint` ([#8628](https://github.com/hashicorp/terraform-provider-aws/issues/8628))
* **New Data Source:** `aws_sagemaker_prebuilt_ecr_image` ([#15924](https://github.com/hashicorp/terraform-provider-aws/pull/15924))
* **New Data Source:** `aws_workspaces_workspace` ([#14135](https://github.com/hashicorp/terraform-provider-aws/issues/14135))
* **New Resource:** `aws_secretsmanager_secret_policy` ([#14468](https://github.com/hashicorp/terraform-provider-aws/pull/14468))

ENHANCEMENTS

* resource/aws_glue_catalog_database: add plan time validations for `description` and `name`. ([#15956](https://github.com/hashicorp/terraform-provider-aws/issues/15956))
* resource/aws_glue_crawler: Support MongoDB target ([#15934](https://github.com/hashicorp/terraform-provider-aws/issues/15934))
* resource/aws_glue_trigger: Add plan time validation to `name` ([#15793](https://github.com/hashicorp/terraform-provider-aws/issues/15793))
* resource/aws_glue_trigger: Add `security_configuration` and `notification_property` arguments to `actions` block ([#15793](https://github.com/hashicorp/terraform-provider-aws/issues/15793))

BUG FIXES

* resource/aws_glue_classifier: Fix `quote_symbol` being optional ([#15948](https://github.com/hashicorp/terraform-provider-aws/issues/15948))

## 3.13.0 (October 29, 2020)

NOTES

* data-source/aws_caller_identity: The `id` attribute has changed to the ID of the AWS Account. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_coip_pools: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_instance_type_offerings: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateway_route_tables: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateway_virtual_interface_groups: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateways: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_spot_price: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_efs_access_points: The `id` attribute has changed to the EFS File System identifier. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_glue_script: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_inspector_rules_packages: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_instances: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_outposts_outposts: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_outposts_sites: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))

FEATURES

* **New Resource:** `aws_glue_resource_policy` ([#10361](https://github.com/hashicorp/terraform-provider-aws/issues/10361))
* **New Resource:** `aws_s3control_bucket` ([#15510](https://github.com/hashicorp/terraform-provider-aws/issues/15510))
* **New Resource:** `aws_s3control_bucket_lifecycle_configuration` ([#15604](https://github.com/hashicorp/terraform-provider-aws/issues/15604))
* **New Resource:** `aws_s3control_bucket_policy` ([#15575](https://github.com/hashicorp/terraform-provider-aws/issues/15575))
* **New Resource:** `aws_s3outposts_endpoint` ([#15585](https://github.com/hashicorp/terraform-provider-aws/issues/15585))
* **New Resource:** `aws_sagemaker_code_repository` ([#15809](https://github.com/hashicorp/terraform-provider-aws/issues/15809))
* **New Resource:** `aws_storagegateway_tape_pool` ([#15370](https://github.com/hashicorp/terraform-provider-aws/issues/15370))

ENHANCEMENTS


BUG FIXES

* data-source/aws_caller_identity: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_coip_pools: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_instance_type_offerings: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateway_route_tables: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateway_virtual_interface_groups: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_local_gateways: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_ec2_spot_price: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_efs_access_points: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_glue_script: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_inspector_rules_packages: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_instances: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_outposts_outposts: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))
* data-source/aws_outposts_sites: Prevent plan differences with the `id` attribute ([#15896](https://github.com/hashicorp/terraform-provider-aws/issues/15896))

## 3.12.0 (October 22, 2020)

FEATURES

* **New Data Source:** `aws_rds_certificate` ([#15789](https://github.com/hashicorp/terraform-provider-aws/issues/15789))
* **New Resource:** `aws_autoscalingplans_scaling_plan` ([#8965](https://github.com/hashicorp/terraform-provider-aws/issues/8965))
* **New Resource:** `aws_cloudwatch_event_bus` ([#10256](https://github.com/hashicorp/terraform-provider-aws/issues/10256))
* **New Resource:** `aws_kinesisanalyticsv2_application` ([#11652](https://github.com/hashicorp/terraform-provider-aws/issues/11652))
* **New Resource:** `aws_storagegateway_stored_iscsi_volume` ([#12027](https://github.com/hashicorp/terraform-provider-aws/issues/12027))

ENHANCEMENTS

* resource/aws_fsx_lustre_file_system: Add `copy_tags_to_backups` support ([#15687](https://github.com/hashicorp/terraform-provider-aws/issues/15687))
* resource/aws_fsx_lustre_file_system: Increased maximum `automatic_backup_retention_days` from 35 to 90 ([#15641](https://github.com/hashicorp/terraform-provider-aws/issues/15641))
* resource/aws_fsx_windows_file_system: Increased maximum `automatic_backup_retention_days` from 35 to 90 ([#15641](https://github.com/hashicorp/terraform-provider-aws/issues/15641))
* resource/aws_glue_catalog_table: add validation checks for resource properties ([#12523](https://github.com/hashicorp/terraform-provider-aws/issues/12523))

BUGFIXES

* resource/aws_workspaces_workspace: Fix terminated state resolution ([#15705](https://github.com/hashicorp/terraform-provider-aws/issues/15705))
* resource/aws_glue_table_catalog_table: Prevent errors on `unset` of `ser_de_info.name` ([#15127](https://github.com/hashicorp/terraform-provider-aws/issues/15127))
* resource/aws_glue_security_configuration: Don't send empty `kms_arn` if mode is `DISABLED` ([#13618](https://github.com/hashicorp/terraform-provider-aws/issues/13618))

## 3.11.0 (October 15, 2020)

FEATURES

* **New Data Source:** `aws_codeartifact_repository_endpoint` ([#15566](https://github.com/hashicorp/terraform-provider-aws/issues/15566))
* **New Resource:** `aws_appmesh_gateway_route` ([#15638](https://github.com/hashicorp/terraform-provider-aws/issues/15638))
* **New Resource:** `aws_appmesh_virtual_gateway` ([#15611](https://github.com/hashicorp/terraform-provider-aws/issues/15611))

BUG FIXES

* resource/aws_ec2_transit_gateway_route: Prevent plan errors with compressed IPv6 addresses ([#14846](https://github.com/hashicorp/terraform-provider-aws/issues/14846))

ENHANCEMENTS

* data-source/aws_workspaces_directory: Add workspaces creation properties ([#14577](https://github.com/hashicorp/terraform-provider-aws/issues/14577))
* resource/aws_workspaces_directory: Add workspaces creation properties ([#14577](https://github.com/hashicorp/terraform-provider-aws/issues/14577))

## 3.10.0 (October 09, 2020)

FEATURES

* **New Data Source:** `aws_codeartifact_authorization_token` ([#15425](https://github.com/hashicorp/terraform-provider-aws/issues/15425))
* **New Data Source:** `aws_ec2_instance_type` ([#13124](https://github.com/hashicorp/terraform-provider-aws/issues/13124))
* **New Data Source:** `aws_lex_bot_alias` ([#8919](https://github.com/hashicorp/terraform-provider-aws/issues/8919))
* **New Data Source:** `aws_redshift_orderable_cluster` ([#15438](https://github.com/hashicorp/terraform-provider-aws/issues/15438))
* **New Resource:** `aws_codeartifact_repository_permissions_policy` ([#15562](https://github.com/hashicorp/terraform-provider-aws/issues/15562))
* **New Resource:** `aws_lex_bot_alias` ([#8919](https://github.com/hashicorp/terraform-provider-aws/issues/8919))
* **New Resource:** `aws_s3_bucket_ownership_controls` ([#15482](https://github.com/hashicorp/terraform-provider-aws/issues/15482))

NOTES

* data-source/aws_acm_certificate: The `id` attribute has changed to the ARN of the ACM Certificate. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_availability_zones: The `id` attribute has changed to the name of the AWS Region. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_ec2_instance_type_offering: The `id` attribute has changed to the EC2 Instance Type. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_partition: The `id` attribute has changed to the identifier of the AWS Partition. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_regions: The `id` attribute has changed to the identifier of the AWS Partition. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_sns_topic: The `id` attribute has changed to the ARN of the SNS Topic. The first apply of this updated data source may show this difference. ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))

ENHANCEMENTS

* data-source/aws_batch_compute_environment: Add `tags` attribute ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* data-source/aws_batch_job_queue: Add `tags` attribute ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* resource/aws_batch_compute_environment: Add `tags` argument ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* resource/aws_batch_job_definition: Add `tags` argument ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* resource/aws_batch_job_queue: Add `tags` argument ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* resource/aws_lb_target_group: Add `source_ip` as an option for the `stickiness.type` argument. ([#15295](https://github.com/hashicorp/terraform-provider-aws/issues/15295))
* resource/aws_sns_topic_subscription: Create subscriptions with attributes (delivery policy, filter policy, etc.) instead of separate API calls ([#10496](https://github.com/hashicorp/terraform-provider-aws/issues/10496))

BUG FIXES

* data-source/aws_acm_certificate: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_availability_zones: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_ec2_instance_type_offering: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_partition: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_regions: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* data-source/aws_sns_topic: Prevent plan differences with the `id` attribute ([#15399](https://github.com/hashicorp/terraform-provider-aws/issues/15399))
* resource/aws_acm_certificate: Prevent unexpected timeout error on deletion due to API retries ([#15522](https://github.com/hashicorp/terraform-provider-aws/issues/15522))
* resource/aws_batch_job_definition: Prevent unexpected plan difference for `container_properties` argument value with new secrets support ([#15470](https://github.com/hashicorp/terraform-provider-aws/issues/15470))
* resource/aws_gamelift_fleet: Prevent unexpected timeout error on creation due to API retries ([#15526](https://github.com/hashicorp/terraform-provider-aws/issues/15526))
* resource/aws_glue_workflow: Ensure `max_concurrent_runs` attribute is properly saved during import ([#15538](https://github.com/hashicorp/terraform-provider-aws/issues/15538))
* resource/aws_lex_bot: Prevent unexpected timeout error on creation due to API retries ([#15527](https://github.com/hashicorp/terraform-provider-aws/issues/15527))
* resource/aws_lex_bot_alias: Prevent unexpected timeout error on creation due to API retries ([#15527](https://github.com/hashicorp/terraform-provider-aws/issues/15527))
* resource/aws_lex_intent: Prevent unexpected timeout error on creation due to API retries ([#15527](https://github.com/hashicorp/terraform-provider-aws/issues/15527))
* resource/aws_lex_slot_type: Prevent unexpected timeout error on creation due to API retries ([#15527](https://github.com/hashicorp/terraform-provider-aws/issues/15527))

## 3.9.0 (October 02, 2020)

FEATURES

* **New Resource:** `aws_backup_vault_notifications` ([#12501](https://github.com/hashicorp/terraform-provider-aws/issues/12501))
* **New Resource:** `aws_codeartifact_domain` ([#13743](https://github.com/hashicorp/terraform-provider-aws/issues/13743))
* **New Resource:** `aws_codeartifact_domain_permissions` ([#13753](https://github.com/hashicorp/terraform-provider-aws/issues/13753))
* **New Resource:** `aws_codeartifact_repository` ([#14429](https://github.com/hashicorp/terraform-provider-aws/issues/14429))
* **New Resource:** `aws_db_proxy_target` ([#12784](https://github.com/hashicorp/terraform-provider-aws/issues/12784))
* **New Resource:** `aws_glue_data_catalog_encryption_settings` ([#14916](https://github.com/hashicorp/terraform-provider-aws/issues/14916))
* **New Resource:** `aws_glue_ml_transform` ([#14909](https://github.com/hashicorp/terraform-provider-aws/issues/14909))
* **New Resource:** `aws_glue_partition` ([#12547](https://github.com/hashicorp/terraform-provider-aws/issues/12547))
* **New Resource:** `aws_lex_bot` ([#8918](https://github.com/hashicorp/terraform-provider-aws/issues/8918))
* **New Resource:** `aws_lex_intent` ([#8917](https://github.com/hashicorp/terraform-provider-aws/issues/8917))
* **New Data Source:** `aws_lex_bot` ([#8918](https://github.com/hashicorp/terraform-provider-aws/issues/8918))
* **New Data Source:** `aws_lex_intent` ([#8917](https://github.com/hashicorp/terraform-provider-aws/issues/8917))

ENHANCEMENTS

* resource/aws_fsx_lustre_file_system: Add `auto_import_policy`  argument ([#15231](https://github.com/hashicorp/terraform-provider-aws/issues/15231))
* resource/aws_fsx_lustre_file_system: Support `daily_automatic_backup_start_time` ([#15299](https://github.com/hashicorp/terraform-provider-aws/issues/15299))
* resource/aws_fsx_lustre_file_system: Add `storage_type` and `drive_cache_type` ([#14727](https://github.com/hashicorp/terraform-provider-aws/issues/14727))
* resource/aws_glue_crawler: Add `connection_name` field to `s3_target` block ([#15350](https://github.com/hashicorp/terraform-provider-aws/issues/15350))

BUG FIXES

* resource/aws_ec2_client_vpn_authorization_rule: Increased active and revoked timeouts from 5 to 10 minutes ([#15367](https://github.com/hashicorp/terraform-provider-aws/issues/15367))

## 3.8.0 (September 24, 2020)

FEATURES

* **New Resource:** `aws_datasync_location_fsx_windows` ([#12686](https://github.com/hashicorp/terraform-provider-aws/pull/12686))
* **New Resource:** `aws_route53_resolver_query_log_config`. ([#14897](https://github.com/hashicorp/terraform-provider-aws/pull/14897))
* **New Resource:** `aws_route53_resolver_query_log_config_association`. ([#14901](https://github.com/hashicorp/terraform-provider-aws/pull/14901))
* **New Data Source:** `aws_rds_engine_version` ([#15228](https://github.com/hashicorp/terraform-provider-aws/pull/15228))
* **New Data Source:** `aws_docdb_engine_version` ([#15253](https://github.com/hashicorp/terraform-provider-aws/pull/15253))
* **New Data Source:** `aws_neptune_engine_version` ([#15259](https://github.com/hashicorp/terraform-provider-aws/pull/15259))
* **New Data Source:** `aws_workspaces_image` ([#11428](https://github.com/hashicorp/terraform-provider-aws/issues/11428))

ENHANCEMENTS

* data-source/aws_lb: Add `customer_owned_ipv4_pool` and `subnet_mapping` `outpost_id` attributes ([#15170](https://github.com/hashicorp/terraform-provider-aws/issues/15170))
* resource/aws_glue_connection: Support `NETWORK` connection type ([#14818](https://github.com/hashicorp/terraform-provider-aws/issues/14818))
* resource/aws_glue_crawler: Add support for `scan_all` and `scan_rate` arguments for ddb targets ([#14819](https://github.com/hashicorp/terraform-provider-aws/issues/14819))
* resource/aws_glue_crawler: Allow removing `table_prefix` ([#15268](https://github.com/hashicorp/terraform-provider-aws/issues/15268))
* resource/aws_glue_job: Add `non_overridable_arguments` argument ([#14793](https://github.com/hashicorp/terraform-provider-aws/issues/14793))
* resource/aws_glue_workflow: Add `tags` argument ([#14910](https://github.com/hashicorp/terraform-provider-aws/issues/14910))
* resource/aws_glue_workflow: Add `arn` attribute ([#14910](https://github.com/hashicorp/terraform-provider-aws/issues/14910))
* resource/aws_glue_workflow: Add `max_concurrent_runs` argument ([#14910](https://github.com/hashicorp/terraform-provider-aws/issues/14910))
* resource/aws_glue_workflow: Plan time validation for `name` ([#14910](https://github.com/hashicorp/terraform-provider-aws/issues/14910))
* resource/aws_fsx_lustre_file_system: Add support for backup retention ([#14446](https://github.com/hashicorp/terraform-provider-aws/issues/14446))
* resource/aws_fsx_lustre_file_system: Add `kms_key_id` argument ([#15057](https://github.com/hashicorp/terraform-provider-aws/issues/15057))
* resource/aws_fsx_lustre_file_system: Add `mount_name` argument ([#14313](https://github.com/hashicorp/terraform-provider-aws/issues/14313))
* resource/aws_lb: Add `customer_owned_ipv4_pool` argument and `subnet_mapping` `outpost_id` attribute ([#15170](https://github.com/hashicorp/terraform-provider-aws/issues/15170))

BUG FIXES

* resource/aws_efs_mount_target: Increase create timeout to 30 minutes ([#15293](https://github.com/hashicorp/terraform-provider-aws/issues/15293))
* resource/aws_fsx_lustre_file_system: Change `aws_fsx_lustre_file_system's`'s `network_interface_ids` to `TypeList` to preserve ordering. ([#14314](https://github.com/hashicorp/terraform-provider-aws/issues/14314))
* resource/aws_neptune_cluster_instance: Add `configuring-enhanced-monitoring` to expected states when creating and updating ([#15284](https://github.com/hashicorp/terraform-provider-aws/issues/15284))
* resource/aws_vpn_gateway: Increase VPC detachment timeout to 30 minutes ([#15201](https://github.com/hashicorp/terraform-provider-aws/issues/15201))
* resource/aws_vpn_gateway_attachment: Increase VPC detachment timeout to 30 minutes ([#15201](https://github.com/hashicorp/terraform-provider-aws/issues/15201))

## 3.7.0 (September 17, 2020)

FEATURES

* **New Resource:** `aws_config_remediation_configuration` ([#13884](https://github.com/hashicorp/terraform-provider-aws/issues/13884))

ENHANCEMENTS


BUG FIXES

* resource/aws_acm_certificate: Prevent tagging is not permitted on re-import error ([#15060](https://github.com/hashicorp/terraform-provider-aws/issues/15060))

## 3.6.0 (September 11, 2020)

FEATURES

* **New Resource:** `aws_db_proxy_default_target_group` ([#12743](https://github.com/hashicorp/terraform-provider-aws/issues/12743))

BUG FIXES

* resource/aws_ec2_client_vpn_authorization_rule: Increase active and revoked timeouts from 1 to 5 minutes ([#15037](https://github.com/hashicorp/terraform-provider-aws/issues/15037))

## 3.5.0 (September 03, 2020)

FEATURES

* **New Data Source:** `aws_docdb_orderable_db_instance` ([#14931](https://github.com/hashicorp/terraform-provider-aws/issues/14931))
* **New Data Source:** `aws_lex_slot_type` ([#8916](https://github.com/hashicorp/terraform-provider-aws/issues/8916))
* **New Data Source:** `aws_neptune_orderable_db_instance` ([#14953](https://github.com/hashicorp/terraform-provider-aws/issues/14953))
* **New Data Source:** `aws_rds_orderable_db_instance` ([#14834](https://github.com/hashicorp/terraform-provider-aws/issues/14834))
* **New Data Source:** `aws_vpc_peering_connections` ([#9491](https://github.com/hashicorp/terraform-provider-aws/issues/9491))
* **New Resource:** `aws_codebuild_report_group` ([#12573](https://github.com/hashicorp/terraform-provider-aws/issues/12573))
* **New Resource:** `aws_db_proxy` ([#12704](https://github.com/hashicorp/terraform-provider-aws/issues/12704))
* **New Resource:** `aws_emr_instance_fleet` ([#14813](https://github.com/hashicorp/terraform-provider-aws/issues/14813))
* **New Resource:** `aws_glue_user_defined_function` ([#12537](https://github.com/hashicorp/terraform-provider-aws/issues/12537))
* **New Resource:** `aws_guardduty_filter` ([#14876](https://github.com/hashicorp/terraform-provider-aws/issues/14876))
* **New Resource:** `aws_lex_slot_type` ([#8916](https://github.com/hashicorp/terraform-provider-aws/issues/8916))

ENHANCEMENTS

* data-source/aws_cur_report_definition: Add `refresh_closed_reports` and `report_versioning` attributes ([#12428](https://github.com/hashicorp/terraform-provider-aws/issues/12428))
* data-source/aws_outposts_outpost: Add `arn` argument ([#14967](https://github.com/hashicorp/terraform-provider-aws/issues/14967))
* resource/aws_acm_certificate: Provide additional plan-time validation for `subject_alternative_names` argument values ([#14782](https://github.com/hashicorp/terraform-provider-aws/issues/14782))
* resource/aws_cur_report_definition: Add `refresh_closed_reports` and `report_versioning` arguments ([#12428](https://github.com/hashicorp/terraform-provider-aws/issues/12428))
* resource/aws_cur_report_definition: Support `ATHENA` value in `additional_artifacts` argument plan-time validation ([#12428](https://github.com/hashicorp/terraform-provider-aws/issues/12428))
* resource/aws_cur_report_definition: Support `Parquet` value in `compression` and `format` argument plan-time validations ([#12428](https://github.com/hashicorp/terraform-provider-aws/issues/12428))
* resource/aws_cur_report_definition: Support `MONTHLY` value in `time_unit` argument plan-time validation ([#12428](https://github.com/hashicorp/terraform-provider-aws/issues/12428))
* resource/aws_ec2_client_vpn_endpoint: Support `authentication_options` `type` argument `federated-authentication` value and new `saml_provider_arn` argument ([#14171](https://github.com/hashicorp/terraform-provider-aws/issues/14171))
* resource/aws_instance: Support `io2` value for `volume_type` argument plan-time validation ([#14906](https://github.com/hashicorp/terraform-provider-aws/issues/14906))

BUG FIXES

* resource/aws_mq_configuration: Prevent additional revision creation with `tags` only updates ([#14850](https://github.com/hashicorp/terraform-provider-aws/issues/14850))
* resource/aws_opsworks_stack: Suppress equivalent `custom_json` differences ([#14886](https://github.com/hashicorp/terraform-provider-aws/issues/14886))
* resource/aws_workspaces_workspace: Prevent error when `workspace_properties` `running_mode` is set to `ALWAYS_ON` ([#13976](https://github.com/hashicorp/terraform-provider-aws/issues/13976))

## 3.4.0 (August 27, 2020)

FEATURES

* **New Data Source:** `aws_db_subnet_group` ([#9525](https://github.com/hashicorp/terraform-provider-aws/issues/9525))
* **New Resource:** `aws_emr_managed_scaling_policy` ([#13965](https://github.com/hashicorp/terraform-provider-aws/issues/13965))
* **New Resource:** `aws_guardduty_publishing_destination` ([#13894](https://github.com/hashicorp/terraform-provider-aws/issues/13894))
* **New Resource:** `aws_securityhub_action_target` ([#10493](https://github.com/hashicorp/terraform-provider-aws/issues/10493))
* **New Resource:** `aws_xray_encryption_config` ([#13600](https://github.com/hashicorp/terraform-provider-aws/issues/13600))
* **New Resource:** `aws_xray_group` ([#13597](https://github.com/hashicorp/terraform-provider-aws/issues/13597))

ENHANCEMENTS

* resource/aws_globalaccelerator_endpoint_group: Add `client_ip_preservation_enabled` argument to the `endpoint_configuration` configuration block ([#14486](https://github.com/hashicorp/terraform-provider-aws/issues/14486))
* resource/aws_xray_sampling_rule: Add `tags` argument ([#14831](https://github.com/hashicorp/terraform-provider-aws/issues/14831))

BUG FIXES

* resource/aws_acmpca_certificate_authority: Ensure `DELETED` status triggers state removal ([#13684](https://github.com/hashicorp/terraform-provider-aws/issues/13684))
* resource/aws_globalaccelerator_accelerator: Increase creation timeout to 10 minutes ([#14486](https://github.com/hashicorp/terraform-provider-aws/issues/14486))
* resource/aws_globalaccelerator_endpoint_group: Prevent differences with `health_check_path` defaults ([#14486](https://github.com/hashicorp/terraform-provider-aws/issues/14486))
* resource/aws_glue_crawler: Properly update `schedule` value ([#14792](https://github.com/hashicorp/terraform-provider-aws/issues/14792))

## 3.3.0 (August 20, 2020)

ENHANCEMENTS

* provider: Support for appending information to User-Agent request headers with the `TF_APPEND_USER_AGENT` environment variable ([#14555](https://github.com/hashicorp/terraform-provider-aws/issues/14555))
* resource/aws_customer_gateway: Support tag on create ([#14501](https://github.com/hashicorp/terraform-provider-aws/issues/14501))
* resource/aws_ec2_client_vpn_network_association: Allow specifying custom security groups ([#14146](https://github.com/hashicorp/terraform-provider-aws/issues/14146))
* resource/aws_ec2_client_vpn_network_association: Support resource import ([#14146](https://github.com/hashicorp/terraform-provider-aws/issues/14146))
* resource/aws_egress_only_intrenet_gateway:-Ssupport tag on create  ([#14501](https://github.com/hashicorp/terraform-provider-aws/issues/14501))
* resource/aws_internet_gateway: Support tag on create  ([#14501](https://github.com/hashicorp/terraform-provider-aws/issues/14501))
* resource/aws_vpn_connection: Support tag on create  ([#14501](https://github.com/hashicorp/terraform-provider-aws/issues/14501))
* resource/aws_vpn_gateway: Support tag on create  ([#14501](https://github.com/hashicorp/terraform-provider-aws/issues/14501))

BUG FIXES

* resource/aws_user_pool_domain: Ensure state removal when deleted outside Terraform ([#14732](https://github.com/hashicorp/terraform-provider-aws/issues/14732))

## 3.2.0 (August 14, 2020)

ENHANCEMENTS

* data-source/aws_lb: Add `subnet_mapping` `private_ipv4_address` attribute ([#14545](https://github.com/hashicorp/terraform-provider-aws/issues/14545))
* provider: Upgrade to Terraform Plugin SDK V2. There should be no breaking changes from a practitioner's perspective. Some validation errors should now feature enhanced messaging. ([#14432](https://github.com/hashicorp/terraform-provider-aws/issues/14432))
* resource/aws_accessanalyzer_analyzer: Support `ORGANIZATION` value in `type` argument ([#14493](https://github.com/hashicorp/terraform-provider-aws/issues/14493))

BUG FIXES

* resource/aws_acm_certificate_validation: Prevent panic with missing `DomainValidationOptions` `ResourceRecord` attribute in API response [[#14590](https://github.com/hashicorp/terraform-provider-aws/issues/14590)]

## 3.1.0 (August 07, 2020)

NOTES:


FEATURES

* **New Data Source:** `aws_ec2_spot_price` ([#12504](https://github.com/hashicorp/terraform-provider-aws/issues/12504))
* **New Resource**: `aws_route53_vpc_association_authorization` ([#14215](https://github.com/hashicorp/terraform-provider-aws/issues/14215))

ENHANCEMENTS

* resource/aws_lb: Add `subnet_mapping` configuration block `private_ipv4_address` argument ([#11404](https://github.com/hashicorp/terraform-provider-aws/issues/11404))

BUG FIXES

* data-source/aws_availability_zones: Prevent unexpected plan output every apply with `group_names` attribute ([#14412](https://github.com/hashicorp/terraform-provider-aws/issues/14412))
* resource/aws_secretsmanager_secret: Allow retries for IAM eventual consistency errors ([#14459](https://github.com/hashicorp/terraform-provider-aws/issues/14459))

## 3.0.0 (July 31, 2020)

NOTES:
* provider: This version is built using Go 1.14.5, including security fixes to the crypto/x509 and net/http packages.

BREAKING CHANGES

* provider: New versions of the provider can only be automatically installed on Terraform 0.12 and later ([#14143](https://github.com/hashicorp/terraform-provider-aws/issues/14143))
* provider: All "removed" attributes are cut, using them would result in a Terraform Core level error ([#14001](https://github.com/hashicorp/terraform-provider-aws/issues/14001))
* provider: Credential ordering has changed from static, environment, shared credentials, EC2 metadata, default AWS Go SDK (shared configuration, web identity, ECS, EC2 Metadata) to static, environment, shared credentials, default AWS Go SDK (shared configuration, web identity, ECS, EC2 Metadata) ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* provider: The `AWS_METADATA_TIMEOUT` environment variable no longer has any effect as we now depend on the default AWS Go SDK EC2 Metadata client timeout of one second with two retries ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* provider: Remove deprecated `kinesis_analytics` and `r53` custom service endpoint arguments ([#14238](https://github.com/hashicorp/terraform-provider-aws/issues/14238))
* data-source/aws_availability_zones: Remove deprecated `blacklisted_names` and `blacklisted_zone_ids` arguments ([#14134](https://github.com/hashicorp/terraform-provider-aws/issues/14134))
* data-source/aws_directory_service_directory: Return an error when a single result is not found ([#14006](https://github.com/hashicorp/terraform-provider-aws/issues/14006))
* data-source/aws_efs_file_system: Return an error when a single result is not found ([#14005](https://github.com/hashicorp/terraform-provider-aws/issues/14005))
* resource/aws_acm_certificate: `certificate_body`, `certificate_chain`, and `private_key` attributes are no longer stored in the Terraform state with hash values ([#9685](https://github.com/hashicorp/terraform-provider-aws/issues/9685))
* resource/aws_acm_certificate: `domain_validation_options` attribute changed from list to set ([#14199](https://github.com/hashicorp/terraform-provider-aws/issues/14199))
* resource/aws_acm_certificate: Plan-time validation added to `domain_name` and `subject_alternative_names` arguments to prevent usage of strings with trailing periods ([#14220](https://github.com/hashicorp/terraform-provider-aws/issues/14220))
* resource/aws_dx_gateway: Remove automatic `aws_dx_gateway_association` resource import ([#14124](https://github.com/hashicorp/terraform-provider-aws/issues/14124))
* resource/aws_dx_gateway_association: Remove deprecated `vpn_gateway_id` argument ([#14144](https://github.com/hashicorp/terraform-provider-aws/issues/14144))
* resource/aws_dx_gateway_association_proposal: Remove deprecated `vpn_gateway_id` argument ([#14144](https://github.com/hashicorp/terraform-provider-aws/issues/14144))
* resource/aws_glue_job: Remove deprecated `allocated_capacity` argument ([#14296](https://github.com/hashicorp/terraform-provider-aws/issues/14296))
* resource/aws_instance: Return an error when `ebs_block_device` `iops` or `root_block_device` `iops` argument set to a value greater than `0` for volume types other than `io1` ([#14310](https://github.com/hashicorp/terraform-provider-aws/issues/14310))
* resource/aws_lb_listener_rule: Remove deprecated `condition` configuration block `field` and `values` arguments ([#14309](https://github.com/hashicorp/terraform-provider-aws/issues/14309))
* resource/aws_sns_platform_application: `platform_credential` and `platform_principal` attributes are no longer stored in the Terraform state with hash values ([#3894](https://github.com/hashicorp/terraform-provider-aws/issues/3894))

FEATURES

* **New Data Source:** aws_workspaces_directory ([#13529](https://github.com/hashicorp/terraform-provider-aws/issues/13529))

ENHANCEMENTS

* provider: Always enable shared configuration file support (no longer require `AWS_SDK_LOAD_CONFIG` environment variable) ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* provider: Add `assume_role` configuration block `duration_seconds`, `policy_arns`, `tags`, and `transitive_tag_keys` arguments ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* data-source/aws_instance: Add `secondary_private_ips` attribute ([#14079](https://github.com/hashicorp/terraform-provider-aws/issues/14079))
* resource/aws_acm_certificate: Enable `domain_validation_options` usage in downstream resource `count` and `for_each` references ([#14199](https://github.com/hashicorp/terraform-provider-aws/issues/14199))
* resource/aws_fsx_lustre_file_system: Add `deployment_type` and `per_unit_storage_throughput` attributes ([#13639](https://github.com/hashicorp/terraform-provider-aws/issues/13639))
* resource_aws_fsx_windows_file_system - add `storage_type` argument. ([#14316](https://github.com/hashicorp/terraform-provider-aws/issues/14316))
* resource_aws_fsx_windows_file_system: add support for multi-az ([#12676](https://github.com/hashicorp/terraform-provider-aws/issues/12676))
* resource_aws_fsx_windows_file_system: add `SINGLE_AZ_2` deployment type ([#12676](https://github.com/hashicorp/terraform-provider-aws/issues/12676))
* resource_aws_fsx_windows_file_system: adds `preferred_file_server_ip`, `remote_administration_endpoint` attributes ([#12676](https://github.com/hashicorp/terraform-provider-aws/issues/12676))
* resource/aws_instance: Add `secondary_private_ips` argument (conflicts with `network_interface` configuration block) ([#14079](https://github.com/hashicorp/terraform-provider-aws/issues/14079))

BUG FIXES

* provider: Ensure nil is not passed to RetryError helpers, may result in some bug fixes ([#14104](https://github.com/hashicorp/terraform-provider-aws/issues/14104))
* provider: Ensure configured STS endpoint is used during `AssumeRole` API calls ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* provider: Prefer AWS shared configuration over EC2 metadata credentials by default ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* provider: Prefer CodeBuild, ECS, EKS credentials over EC2 metadata credentials by default ([#14077](https://github.com/hashicorp/terraform-provider-aws/issues/14077))
* data-source/aws_lb: `enable_http2` now properly set ([#14167](https://github.com/hashicorp/terraform-provider-aws/issues/14167))
* resource/aws_acm_certificate: Prevent unexpected ordering differences with `domain_validation_options` attribute ([#14199](https://github.com/hashicorp/terraform-provider-aws/issues/14199))
* resource/aws_dx_gateway_association: Increase default create/update/delete timeouts to 30 minutes ([#14144](https://github.com/hashicorp/terraform-provider-aws/issues/14144))
* resource/aws_lb_cookie_stickiness_policy: `lb_port` now properly set ([#14167](https://github.com/hashicorp/terraform-provider-aws/issues/14167))
* resource/aws_sns_topic_subscription: Immediately return `ListSubscriptionsByTopic` errors ([#14262](https://github.com/hashicorp/terraform-provider-aws/issues/14262))

## Previous Releases

For information on prior major releases, see their changelogs:

* [3.74 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/3.x/CHANGELOG.md)
* [2.70 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/2.x/CHANGELOG.md)
