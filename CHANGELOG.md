## Regex Remove Stuff I don't care about
```
^\* (resource|data-source)/aws_(ami|api|app|auto|back|ce|cloud|code|cog|conf|con|data|db|dms|dyna|ebs|ecr|ecs|eip|eks|elas|emr|flow|glac|game|guar|iam|image|iot|kine|kms|lam|launch|light|msk|net|ops|org|pin|rds|red|route|sage|s3|secur|ses|spot|ssm|stor|sub|trans|vpc|waf).*\n
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

## Previous Releases

For information on prior major releases, see their changelogs:

* [3.74 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/3.x/CHANGELOG.md)
* [2.70 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/2.x/CHANGELOG.md)
