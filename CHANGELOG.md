## 4.24.0 (Unreleased)

ENHANCEMENTS:

* resource/aws_ses_configuration_set: Add `tracking_options.0.custom_redirect_domain` argument (NOTE: This enhancement is provided as best effort due to testing limitations, i.e., the requirement of a verified domain) ([#26032](https://github.com/hashicorp/terraform-provider-aws/issues/26032))

## 4.24.0 (Unreleased)

FEATURES:

* **New Resource:** `aws_acmpca_permission` ([#12485](https://github.com/hashicorp/terraform-provider-aws/issues/12485))
* **New Resource:** `aws_ssm_service_setting` ([#13018](https://github.com/hashicorp/terraform-provider-aws/issues/13018))

ENHANCEMENTS:

* data-source/aws_ecs_service: Add `tags` attribute ([#25961](https://github.com/hashicorp/terraform-provider-aws/issues/25961))
* resource/aws_datasync_task: Add `includes` argument ([#25929](https://github.com/hashicorp/terraform-provider-aws/issues/25929))
* resource/aws_guardduty_detector: Add `malware_protection` attribute to the `datasources` configuration block ([#25994](https://github.com/hashicorp/terraform-provider-aws/issues/25994))
* resource/aws_guardduty_organization_configuration: Add `malware_protection` attribute to the `datasources` configuration block ([#25992](https://github.com/hashicorp/terraform-provider-aws/issues/25992))
* resource/aws_security_group: Additional plan-time validation for `name` and `name_prefix` ([#15011](https://github.com/hashicorp/terraform-provider-aws/issues/15011))
* resource/aws_security_group_rule: Add configurable Create timeout ([#24340](https://github.com/hashicorp/terraform-provider-aws/issues/24340))

BUG FIXES:

* data-source/aws_networkmanager_core_network_policy_document: Fix bug where bool values for `attachment-policy.action.require-acceptance` can only be `true` or omitted ([#26010](https://github.com/hashicorp/terraform-provider-aws/issues/26010))
* resource/aws_appmesh_gateway_route: Fix crash when only one of hostname rewrite or path rewrite is configured ([#26012](https://github.com/hashicorp/terraform-provider-aws/issues/26012))
* resource/aws_ce_anomaly_subscription:Fix crash upon adding or removing monitor ARNs to `monitor_arn_list`. ([#25941](https://github.com/hashicorp/terraform-provider-aws/issues/25941))
* resource/aws_cognito_identity_pool_provider_principal_tag: Fix read operation when using an OIDC provider ([#25964](https://github.com/hashicorp/terraform-provider-aws/issues/25964))
* resource/aws_route53_record: Don't ignore `dualstack` prefix in Route 53 Record alias names ([#10672](https://github.com/hashicorp/terraform-provider-aws/issues/10672))
* resource/aws_s3_bucket: Prevents unexpected import of existing bucket in `us-east-1`. ([#26011](https://github.com/hashicorp/terraform-provider-aws/issues/26011))
* resource/aws_s3_bucket: Refactored `object_lock_enabled` parameter's default assignment behavior to protect partitions without Object Lock available. ([#25098](https://github.com/hashicorp/terraform-provider-aws/issues/25098))

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

* data-source/aws_imagebuilder_container_recipe: Add `throughput` attribute to the `block_device_mapping` configuration block ([#25790](https://github.com/hashicorp/terraform-provider-aws/issues/25790))
* data-source/aws_imagebuilder_image_recipe: Add `throughput` attribute to the `block_device_mapping` configuration block ([#25790](https://github.com/hashicorp/terraform-provider-aws/issues/25790))
* data/aws_outposts_asset: Add `rack_elevation` attribute ([#25822](https://github.com/hashicorp/terraform-provider-aws/issues/25822))
* resource/aws_appmesh_gateway_route: Add `http2_route.action.rewrite`, `http2_route.match.hostname`, `http_route.action.rewrite` and `http_route.match.hostname` arguments ([#25819](https://github.com/hashicorp/terraform-provider-aws/issues/25819))
* resource/aws_ce_cost_category: Add `tags` argument and `tags_all` attribute to support resource tagging ([#25432](https://github.com/hashicorp/terraform-provider-aws/issues/25432))
* resource/aws_db_instance_automated_backups_replication: Add support for custom timeouts (create and delete) ([#25796](https://github.com/hashicorp/terraform-provider-aws/issues/25796))
* resource/aws_dynamodb_table: Add `replica.*.propagate_tags` argument to allow propagating tags to replicas ([#25866](https://github.com/hashicorp/terraform-provider-aws/issues/25866))
* resource/aws_flow_log: Add `transit_gateway_id` and `transit_gateway_attachment_id` arguments ([#25913](https://github.com/hashicorp/terraform-provider-aws/issues/25913))
* resource/aws_fsx_openzfs_file_system: Allow in-place update of `storage_capacity`, `throughput_capacity`, and `disk_iops_configuration`. ([#25841](https://github.com/hashicorp/terraform-provider-aws/issues/25841))
* resource/aws_guardduty_organization_configuration: Add `kubernetes` attribute to the `datasources` configuration block ([#25131](https://github.com/hashicorp/terraform-provider-aws/issues/25131))
* resource/aws_imagebuilder_container_recipe: Add `throughput` argument to the `block_device_mapping` configuration block ([#25790](https://github.com/hashicorp/terraform-provider-aws/issues/25790))
* resource/aws_imagebuilder_image_recipe: Add `throughput` argument to the `block_device_mapping` configuration block ([#25790](https://github.com/hashicorp/terraform-provider-aws/issues/25790))
* resource/aws_rds_cluster_instance: Allow `performance_insights_retention_period` values that are multiples of `31` ([#25729](https://github.com/hashicorp/terraform-provider-aws/issues/25729))

BUG FIXES:

* data-source/aws_networkmanager_core_network_policy_document: Fix bug where bool values in `segments` blocks weren't being included in json payloads ([#25789](https://github.com/hashicorp/terraform-provider-aws/issues/25789))
* resource/aws_connect_hours_of_operation: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_connect_queue: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_connect_quick_connect: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_connect_routing_profile: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_connect_security_profile: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_connect_user_hierarchy_group: Fix tags not being updated ([#24864](https://github.com/hashicorp/terraform-provider-aws/issues/24864))
* resource/aws_iam_role: Fix diffs in `assume_role_policy` when there are no semantic changes ([#23060](https://github.com/hashicorp/terraform-provider-aws/issues/23060))
* resource/aws_iam_role: Fix problem with exclusive management of inline and managed policies when empty (i.e., remove out-of-band policies) ([#23060](https://github.com/hashicorp/terraform-provider-aws/issues/23060))
* resource/aws_rds_cluster: Prevent failure of AWS RDS Cluster creation when it is in `rebooting` state. ([#25718](https://github.com/hashicorp/terraform-provider-aws/issues/25718))
* resource/aws_route_table: Retry resource Create for EC2 eventual consistency ([#25793](https://github.com/hashicorp/terraform-provider-aws/issues/25793))
* resource/aws_storagegateway_gateway: Only manage `average_download_rate_limit_in_bits_per_sec` and `average_upload_rate_limit_in_bits_per_sec` when gateway type supports rate limits ([#25922](https://github.com/hashicorp/terraform-provider-aws/issues/25922))

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

* data-source/aws_imagebuilder_distribution_configuration: Add `fast_launch_configuration` attribute to the `distribution` configuration block ([#25671](https://github.com/hashicorp/terraform-provider-aws/issues/25671))
* resource/aws_acmpca_certificate_authority: Add `revocation_configuration.ocsp_configuration` argument ([#25720](https://github.com/hashicorp/terraform-provider-aws/issues/25720))
* resource/aws_apprunner_service: Add `observability_configuration` argument configuration block ([#25697](https://github.com/hashicorp/terraform-provider-aws/issues/25697))
* resource/aws_autoscaling_group: Add `default_instance_warmup` attribute ([#25722](https://github.com/hashicorp/terraform-provider-aws/issues/25722))
* resource/aws_config_remediation_configuration: Add `parameter.*.static_values` attribute for a list of values ([#25738](https://github.com/hashicorp/terraform-provider-aws/issues/25738))
* resource/aws_dynamodb_table: Add `replica.*.point_in_time_recovery` argument ([#25659](https://github.com/hashicorp/terraform-provider-aws/issues/25659))
* resource/aws_ecr_repository: Add `force_delete` parameter. ([#9913](https://github.com/hashicorp/terraform-provider-aws/issues/9913))
* resource/aws_ecs_service: Add configurable timeouts for Create and Delete. ([#25641](https://github.com/hashicorp/terraform-provider-aws/issues/25641))
* resource/aws_emr_cluster: Add `core_instance_group.ebs_config.throughput` and `master_instance_group.ebs_config.throughput` arguments ([#25668](https://github.com/hashicorp/terraform-provider-aws/issues/25668))
* resource/aws_emr_cluster: Add `gp3` EBS volume support ([#25668](https://github.com/hashicorp/terraform-provider-aws/issues/25668))
* resource/aws_emr_cluster: Add `sc1` EBS volume support ([#25255](https://github.com/hashicorp/terraform-provider-aws/issues/25255))
* resource/aws_gamelift_game_session_queue: Add `notification_target` argument ([#25544](https://github.com/hashicorp/terraform-provider-aws/issues/25544))
* resource/aws_imagebuilder_distribution_configuration: Add `fast_launch_configuration` argument to the `distribution` configuration block ([#25671](https://github.com/hashicorp/terraform-provider-aws/issues/25671))
* resource/aws_placement_group: Add `spread_level` argument ([#25615](https://github.com/hashicorp/terraform-provider-aws/issues/25615))
* resource/aws_sagemaker_notebook_instance: Add `accelerator_types` argument ([#10210](https://github.com/hashicorp/terraform-provider-aws/issues/10210))
* resource/aws_sagemaker_project: Increase SageMaker Project create and delete timeout to 15 minutes ([#25638](https://github.com/hashicorp/terraform-provider-aws/issues/25638))
* resource/aws_ssm_parameter: Add `insecure_value` argument to enable dynamic use of SSM parameter values ([#25721](https://github.com/hashicorp/terraform-provider-aws/issues/25721))
* resource/aws_vpc_ipam_pool_cidr: Better error reporting ([#25287](https://github.com/hashicorp/terraform-provider-aws/issues/25287))

BUG FIXES:

* provider: Ensure that the configured `assume_role_with_web_identity` value is used ([#25681](https://github.com/hashicorp/terraform-provider-aws/issues/25681))
* resource/aws_acmpca_certificate_authority: Fix crash when `revocation_configuration` block is empty ([#25695](https://github.com/hashicorp/terraform-provider-aws/issues/25695))
* resource/aws_cognito_risk_configuration: Increase maximum allowed length of `account_takeover_risk_configuration.notify_configuration.block_email.html_body`, `account_takeover_risk_configuration.notify_configuration.block_email.text_body`, `account_takeover_risk_configuration.notify_configuration.mfa_email.html_body`, `account_takeover_risk_configuration.notify_configuration.mfa_email.text_body`, `account_takeover_risk_configuration.notify_configuration.no_action_email.html_body` and `account_takeover_risk_configuration.notify_configuration.no_action_email.text_body` arguments from `2000` to `20000` ([#25645](https://github.com/hashicorp/terraform-provider-aws/issues/25645))
* resource/aws_dynamodb_table: Prevent `restore_source_name` from forcing replacement when removed to enable restoring from a PITR backup ([#25659](https://github.com/hashicorp/terraform-provider-aws/issues/25659))
* resource/aws_dynamodb_table: Respect custom timeouts including when working with replicas ([#25659](https://github.com/hashicorp/terraform-provider-aws/issues/25659))
* resource/aws_ec2_transit_gateway: Fix MaxItems and subnet size validation in `transit_gateway_cidr_blocks` ([#25673](https://github.com/hashicorp/terraform-provider-aws/issues/25673))
* resource/aws_ecs_service: Fix "unexpected new value" errors on creation. ([#25641](https://github.com/hashicorp/terraform-provider-aws/issues/25641))
* resource/aws_ecs_service: Fix error where tags are sometimes not retrieved. ([#25641](https://github.com/hashicorp/terraform-provider-aws/issues/25641))
* resource/aws_emr_managed_scaling_policy: Support `maximum_ondemand_capacity_units` value of `0` ([#17134](https://github.com/hashicorp/terraform-provider-aws/issues/17134))

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

* data-source/aws_ami: Add `include_deprecated` argument ([#25566](https://github.com/hashicorp/terraform-provider-aws/issues/25566))
* data-source/aws_ami: Make `owners` optional ([#25566](https://github.com/hashicorp/terraform-provider-aws/issues/25566))
* data-source/aws_service_discovery_dns_namespace: Add `tags` attribute ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))
* data/aws_key_pair: New attribute `public_key` populated by setting the new `include_public_key` argument ([#25371](https://github.com/hashicorp/terraform-provider-aws/issues/25371))
* resource/aws_connect_instance: Configurable Create and Delete timeouts ([#24861](https://github.com/hashicorp/terraform-provider-aws/issues/24861))
* resource/aws_key_pair: Added 2 new attributes - `key_type` and `create_time` ([#25371](https://github.com/hashicorp/terraform-provider-aws/issues/25371))
* resource/aws_sagemaker_model: Add `repository_auth_config` arguments in support of [Private Docker Registry](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html) ([#25557](https://github.com/hashicorp/terraform-provider-aws/issues/25557))
* resource/aws_service_discovery_http_namespace: Add `http_name` attribute ([#25162](https://github.com/hashicorp/terraform-provider-aws/issues/25162))
* resource/aws_wafv2_web_acl: Add `rule.action.captcha` argument ([#21766](https://github.com/hashicorp/terraform-provider-aws/issues/21766))

BUG FIXES:

* resource/aws_api_gateway_model: Remove length validation from `schema` argument ([#25623](https://github.com/hashicorp/terraform-provider-aws/issues/25623))
* resource/aws_appstream_fleet_stack_association: Fix association not being found after creation ([#25370](https://github.com/hashicorp/terraform-provider-aws/issues/25370))
* resource/aws_appstream_stack: Fix crash when setting `embed_host_domains` ([#25372](https://github.com/hashicorp/terraform-provider-aws/issues/25372))
* resource/aws_route53_record: Successfully allow renaming of `set_identifier` (specified with multiple routing policies) ([#25620](https://github.com/hashicorp/terraform-provider-aws/issues/25620))

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
* data-source/aws_launch_template: Add `disable_api_stop` attribute ([#25185](https://github.com/hashicorp/terraform-provider-aws/issues/25185))
* data-source/aws_launch_template: Correctly set `credit_specification` for T4g instances ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* data-source/aws_vpc_endpoint: Add `dns_options` and `ip_address_type` attributes ([#25190](https://github.com/hashicorp/terraform-provider-aws/issues/25190))
* data-source/aws_vpc_endpoint_service: Add `supported_ip_address_types` attribute ([#25189](https://github.com/hashicorp/terraform-provider-aws/issues/25189))
* resource/aws_cloudwatch_event_api_destination: Remove validation of a maximum value for the `invocation_rate_limit_per_second` argument ([#25277](https://github.com/hashicorp/terraform-provider-aws/issues/25277))
* resource/aws_datasync_location_efs: Add `access_point_arn`, `file_system_access_role_arn`, and `in_transit_encryption` arguments ([#25182](https://github.com/hashicorp/terraform-provider-aws/issues/25182))
* resource/aws_datasync_location_efs: Add plan time validations for `ec2_config.security_group_arns` ([#25182](https://github.com/hashicorp/terraform-provider-aws/issues/25182))
* resource/aws_ec2_host: Add `outpost_arn` argument ([#25464](https://github.com/hashicorp/terraform-provider-aws/issues/25464))
* resource/aws_instance: Add `disable_api_stop` argument ([#25185](https://github.com/hashicorp/terraform-provider-aws/issues/25185))
* resource/aws_instance: Add `private_dns_name_options` argument ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_instance: Correctly handle `credit_specification` for T4g instances ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_launch_template: Add `disable_api_stop` argument ([#25185](https://github.com/hashicorp/terraform-provider-aws/issues/25185))
* resource/aws_launch_template: Correctly handle `credit_specification` for T4g instances ([#25161](https://github.com/hashicorp/terraform-provider-aws/issues/25161))
* resource/aws_s3_bucket_metric: Add validation to ensure name is <= 64 characters. ([#25260](https://github.com/hashicorp/terraform-provider-aws/issues/25260))
* resource/aws_sagemaker_endpoint_configuration: Add `serverless_config` argument ([#25218](https://github.com/hashicorp/terraform-provider-aws/issues/25218))
* resource/aws_sagemaker_endpoint_configuration: Make `production_variants.initial_instance_count` and `production_variants.instance_type` arguments optional ([#25218](https://github.com/hashicorp/terraform-provider-aws/issues/25218))
* resource/aws_sagemaker_notebook_instance: Add `instance_metadata_service_configuration` argument ([#25236](https://github.com/hashicorp/terraform-provider-aws/issues/25236))
* resource/aws_sagemaker_notebook_instance: Support `notebook-al2-v2` value for `platform_identifier` ([#25236](https://github.com/hashicorp/terraform-provider-aws/issues/25236))
* resource/aws_synthetics_canary: Add `delete_lambda` argument ([#25284](https://github.com/hashicorp/terraform-provider-aws/issues/25284))
* resource/aws_vpc_endpoint: Add `dns_options` and `ip_address_type` arguments ([#25190](https://github.com/hashicorp/terraform-provider-aws/issues/25190))
* resource/aws_vpc_endpoint_service: Add `supported_ip_address_types` argument ([#25189](https://github.com/hashicorp/terraform-provider-aws/issues/25189))
* resource/aws_vpn_connection: Add `outside_ip_address_type` and `transport_transit_gateway_attachment_id` arguments in support of [Private IP VPNs](https://docs.aws.amazon.com/vpn/latest/s2svpn/private-ip-dx.html) ([#25529](https://github.com/hashicorp/terraform-provider-aws/issues/25529))

BUG FIXES:

* data-source/aws_ecr_repository: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_elasticache_cluster: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_iam_policy: Add validation to prevent setting incompatible parameters. ([#25538](https://github.com/hashicorp/terraform-provider-aws/issues/25538))
* data-source/aws_iam_policy: Now loads tags. ([#25538](https://github.com/hashicorp/terraform-provider-aws/issues/25538))
* data-source/aws_lb: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_lb_listener: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_lb_target_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* data-source/aws_sqs_queue: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_api_gateway_model: Suppress whitespace differences between model schemas ([#25245](https://github.com/hashicorp/terraform-provider-aws/issues/25245))
* resource/aws_ce_cost_category: Allow duplicate values in `split_charge_rule.parameter.values` argument ([#25488](https://github.com/hashicorp/terraform-provider-aws/issues/25488))
* resource/aws_ce_cost_category: Fix error passing `split_charge_rule.parameter` to the AWS API ([#25488](https://github.com/hashicorp/terraform-provider-aws/issues/25488))
* resource/aws_cloudwatch_composite_alarm: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_cloudwatch_event_bus: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_cloudwatch_event_rule: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_cloudwatch_metric_alarm: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_cloudwatch_metric_stream: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_cognito_user_pool: Correctly handle missing or empty `account_recovery_setting` attribute ([#25184](https://github.com/hashicorp/terraform-provider-aws/issues/25184))
* resource/aws_ecr_repository: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_ecs_capacity_provider: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_ecs_cluster: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_ecs_service: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_ecs_task_definition: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_ecs_task_set: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_cluster: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_parameter_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_replication_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_subnet_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_user: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_elasticache_user_group: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_instance_profile: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_openid_connect_provider: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_policy: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_role: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_saml_provider: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_server_certificate: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_service_linked_role: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_user: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_iam_virtual_mfa_device: Prevent ISO-partition tagging precautions from eating legit errors ([#25549](https://github.com/hashicorp/terraform-provider-aws/issues/25549))
* resource/aws_keyspaces_table: Relax validation of the `schema_definition.column.type` argument to allow collection types ([#25230](https://github.com/hashicorp/terraform-provider-aws/issues/25230))
* resource/aws_launch_configuration: Remove default value for `associate_public_ip_address` argument and mark as Computed. This fixes a regression introduced in [v4.17.0](https://github.com/hashicorp/terraform-provider-aws/blob/main/CHANGELOG.md#4170-june--3-2022) via [#17695](https://github.com/hashicorp/terraform-provider-aws/issues/17695) when no value is configured, whilst honoring any configured value ([#25450](https://github.com/hashicorp/terraform-provider-aws/issues/25450))
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

* resource/aws_dynamodb_table_item: Fix to remove attribute from table item on update ([#25326](https://github.com/hashicorp/terraform-provider-aws/issues/25326))
* resource/aws_ec2_managed_prefix_list_entry: Fix error when attempting to create or delete multiple list entries ([#25046](https://github.com/hashicorp/terraform-provider-aws/issues/25046))

## 4.18.0 (June 10, 2022)

FEATURES:

* **New Resource:** `aws_ce_anomaly_monitor` ([#25177](https://github.com/hashicorp/terraform-provider-aws/issues/25177))
* **New Resource:** `aws_emrserverless_application` ([#25144](https://github.com/hashicorp/terraform-provider-aws/issues/25144))

ENHANCEMENTS:

* data-source/aws_cloudwatch_logs_groups: Make `log_group_name_prefix` optional ([#25187](https://github.com/hashicorp/terraform-provider-aws/issues/25187))
* data-source/aws_cognito_user_pool_client: Add `enable_propagate_additional_user_context_data` argument ([#25181](https://github.com/hashicorp/terraform-provider-aws/issues/25181))
* data-source/aws_ram_resource_share: Add `resource_share_status` argument. ([#25159](https://github.com/hashicorp/terraform-provider-aws/issues/25159))
* resource/aws_cognito_user_pool_client: Add `enable_propagate_additional_user_context_data` argument ([#25181](https://github.com/hashicorp/terraform-provider-aws/issues/25181))
* resource/aws_ebs_snapshot_copy: Add support for `timeouts` configuration block. ([#20912](https://github.com/hashicorp/terraform-provider-aws/issues/20912))
* resource/aws_ebs_volume: Add `final_snapshot` argument ([#21916](https://github.com/hashicorp/terraform-provider-aws/issues/21916))
* resource/aws_s3_bucket: Add error handling for `ErrCodeNotImplemented` and `ErrCodeXNotImplemented` errors when ready bucket information. ([#24764](https://github.com/hashicorp/terraform-provider-aws/issues/24764))
* resource/aws_vpc_ipam_pool_cidr_allocation: improve internal search mechanism ([#25257](https://github.com/hashicorp/terraform-provider-aws/issues/25257))

BUG FIXES:

* resource/aws_snapshot_create_volume_permission: Error if `account_id` is the snapshot's owner ([#12103](https://github.com/hashicorp/terraform-provider-aws/issues/12103))
* resource/aws_ssm_parameter: Allow `Intelligent-Tiering` to upgrade to `Advanced` tier as needed. ([#25174](https://github.com/hashicorp/terraform-provider-aws/issues/25174))

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

* resource/aws_dms_endpoint: Add `redshift_settings` configuration block ([#21846](https://github.com/hashicorp/terraform-provider-aws/issues/21846))
* resource/aws_dms_endpoint: Add ability to use AWS Secrets Manager with the `aurora-postgresql` and `mongodb` engines ([#23691](https://github.com/hashicorp/terraform-provider-aws/issues/23691))
* resource/aws_dms_endpoint: Add ability to use AWS Secrets Manager with the `aurora`, `mariadb` and `mysql` engines ([#24846](https://github.com/hashicorp/terraform-provider-aws/issues/24846))
* resource/aws_dms_endpoint: Add ability to use AWS Secrets Manager with the `redshift` engine ([#25080](https://github.com/hashicorp/terraform-provider-aws/issues/25080))
* resource/aws_dms_endpoint: Add ability to use AWS Secrets Manager with the `sqlserver` engine ([#22646](https://github.com/hashicorp/terraform-provider-aws/issues/22646))
* resource/aws_guardduty_detector: Add `kubernetes` attribute to the `datasources` configuration block ([#22859](https://github.com/hashicorp/terraform-provider-aws/issues/22859))
* resource/aws_ram_resource_share: Add `permission_arns` argument. ([#25113](https://github.com/hashicorp/terraform-provider-aws/issues/25113))
* resource/aws_redshift_cluster: The `default_iam_role_arn` argument is now Computed ([#25096](https://github.com/hashicorp/terraform-provider-aws/issues/25096))

BUG FIXES:

* data-source/aws_launch_configuration: Correct data type for `ebs_block_device.throughput` and `root_block_device.throughput` attributes ([#25097](https://github.com/hashicorp/terraform-provider-aws/issues/25097))
* resource/aws_db_instance_role_association: Extend timeout to 10 minutes ([#25145](https://github.com/hashicorp/terraform-provider-aws/issues/25145))
* resource/aws_ebs_volume: Fix to preserve `iops` when changing EBS volume type (`io1`, `io2`, `gp3`) ([#23280](https://github.com/hashicorp/terraform-provider-aws/issues/23280))
* resource/aws_launch_configuration: Honor associate_public_ip_address = false ([#17695](https://github.com/hashicorp/terraform-provider-aws/issues/17695))
* resource/aws_rds_cluster_role_association: Extend timeout to 10 minutes ([#25145](https://github.com/hashicorp/terraform-provider-aws/issues/25145))
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

* data-source/aws_ami: Add `tpm_support` attribute ([#25045](https://github.com/hashicorp/terraform-provider-aws/issues/25045))
* data-source/aws_redshift_cluster: Add `aqua_configuration_status` attribute. ([#24856](https://github.com/hashicorp/terraform-provider-aws/issues/24856))
* data-source/aws_redshift_cluster: Add `arn`, `cluster_nodes`, `cluster_nodes`, `maintenance_track_name`, `manual_snapshot_retention_period`, `log_destination_type`, and `log_exports` attributes. ([#24982](https://github.com/hashicorp/terraform-provider-aws/issues/24982))
* data-source/aws_cloudfront_response_headers_policy: Add `server_timing_headers_config` attribute ([#24913](https://github.com/hashicorp/terraform-provider-aws/issues/24913))
* resource/aws_ami: Add `tpm_support` argument ([#25045](https://github.com/hashicorp/terraform-provider-aws/issues/25045))
* resource/aws_ami_copy: Add `tpm_support` argument ([#25045](https://github.com/hashicorp/terraform-provider-aws/issues/25045))
* resource/aws_ami_from_instance: Add `tpm_support` argument ([#25045](https://github.com/hashicorp/terraform-provider-aws/issues/25045))
* resource/aws_autoscaling_group: Add `context` argument ([#24951](https://github.com/hashicorp/terraform-provider-aws/issues/24951))
* resource/aws_autoscaling_group: Add `mixed_instances_policy.launch_template.override.instance_requirements` argument ([#24795](https://github.com/hashicorp/terraform-provider-aws/issues/24795))
* resource/aws_cloudfront_response_headers_policy: Add `server_timing_headers_config` argument ([#24913](https://github.com/hashicorp/terraform-provider-aws/issues/24913))
* resource/aws_cloudsearch_domain: Add `index_field.source_fields` argument ([#24915](https://github.com/hashicorp/terraform-provider-aws/issues/24915))
* resource/aws_cloudwatch_metric_stream: Add `statistics_configuration` argument ([#24882](https://github.com/hashicorp/terraform-provider-aws/issues/24882))
* resource/aws_elasticache_global_replication_group: Add support for upgrading `engine_version`. ([#25077](https://github.com/hashicorp/terraform-provider-aws/issues/25077))
* resource/aws_msk_cluster: Support multiple attribute updates by refreshing `current_version` after each update ([#25062](https://github.com/hashicorp/terraform-provider-aws/issues/25062))
* resource/aws_redshift_cluster: Add `aqua_configuration_status` and `apply_immediately` arguments. ([#24856](https://github.com/hashicorp/terraform-provider-aws/issues/24856))
* resource/aws_redshift_cluster: Add `default_iam_role_arn`, `maintenance_track_name`, and `manual_snapshot_retention_period` arguments. ([#24982](https://github.com/hashicorp/terraform-provider-aws/issues/24982))
* resource/aws_redshift_cluster: Add `logging.log_destination_type` and `logging.log_exports` arguments. ([#24886](https://github.com/hashicorp/terraform-provider-aws/issues/24886))
* resource/aws_redshift_cluster: Add plan-time validation for `iam_roles`, `owner_account`, and `port`. ([#24856](https://github.com/hashicorp/terraform-provider-aws/issues/24856))
* resource/aws_redshift_event_subscription: Add plan time validations for `event_categories`, `source_type`, and `severity`. ([#24909](https://github.com/hashicorp/terraform-provider-aws/issues/24909))
* resource/aws_transfer_server: Add support for `TransferSecurityPolicy-2022-03` `security_policy_name` value ([#25060](https://github.com/hashicorp/terraform-provider-aws/issues/25060))

BUG FIXES:

* resource/aws_appflow_flow: Amend `task_properties` validation to avoid conflicting type assumption ([#24889](https://github.com/hashicorp/terraform-provider-aws/issues/24889))
* resource/aws_db_proxy_target: Fix `InvalidDBInstanceState: DB Instance is in an unsupported state - CREATING, needs to be in [AVAILABLE, MODIFYING, BACKING_UP]` error on resource Create ([#24875](https://github.com/hashicorp/terraform-provider-aws/issues/24875))
* resource/aws_instance: Correctly delete instance on destroy when `disable_api_termination` is `true` ([#19277](https://github.com/hashicorp/terraform-provider-aws/issues/19277))
* resource/aws_instance: Prevent error `InvalidParameterCombination: The parameter GroupName within placement information cannot be specified when instanceInterruptionBehavior is set to 'STOP'` when using a launch template that sets `instance_interruption_behavior` to `stop` ([#24695](https://github.com/hashicorp/terraform-provider-aws/issues/24695))
* resource/aws_msk_cluster: Prevent crash on apply when `client_authentication.tls` is empty ([#25072](https://github.com/hashicorp/terraform-provider-aws/issues/25072))
* resource/aws_servicecatalog_provisioned_product: Add possible `TAINTED` target state for resource update and remove one of the internal waiters during read ([#24804](https://github.com/hashicorp/terraform-provider-aws/issues/24804))

## 4.15.1 (May 20, 2022)

BUG FIXES:

* resource/aws_organizations_account: Fix reading account state for existing accounts ([#24899](https://github.com/hashicorp/terraform-provider-aws/issues/24899))

## 4.15.0 (May 20, 2022)

BREAKING CHANGES:

* resource/aws_msk_cluster: The `ebs_volume_size` argument is deprecated in favor of the `storage_info` block. The `storage_info` block can set `volume_size` and `provisioned_throughput` ([#24767](https://github.com/hashicorp/terraform-provider-aws/issues/24767))

FEATURES:

* **New Data Source:** `aws_lb_hosted_zone_id` ([#24749](https://github.com/hashicorp/terraform-provider-aws/issues/24749))
* **New Data Source:** `aws_networkmanager_core_network_policy_document` ([#24368](https://github.com/hashicorp/terraform-provider-aws/issues/24368))
* **New Resource:** `aws_db_snapshot_copy` ([#9886](https://github.com/hashicorp/terraform-provider-aws/issues/9886))
* **New Resource:** `aws_keyspaces_table` ([#24351](https://github.com/hashicorp/terraform-provider-aws/issues/24351))

ENHANCEMENTS:

* data-source/aws_route53_resolver_rules: add `name_regex` argument ([#24582](https://github.com/hashicorp/terraform-provider-aws/issues/24582))
* resource/aws_autoscaling_group: Add `instance_refresh.preferences.skip_matching` argument ([#23059](https://github.com/hashicorp/terraform-provider-aws/issues/23059))
* resource/aws_autoscaling_policy: Add `enabled` argument ([#12625](https://github.com/hashicorp/terraform-provider-aws/issues/12625))
* resource/aws_ec2_fleet: Add `arn` attribute ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))
* resource/aws_ec2_fleet: Add `launch_template_config.override.instance_requirements` argument ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))
* resource/aws_ec2_fleet: Add support for `capacity-optimized` and `capacity-optimized-prioritized` values for `spot_options.allocation_strategy` ([#24732](https://github.com/hashicorp/terraform-provider-aws/issues/24732))
* resource/aws_lambda_function: Add support for `nodejs16.x` `runtime` value ([#24768](https://github.com/hashicorp/terraform-provider-aws/issues/24768))
* resource/aws_lambda_layer_version: Add support for `nodejs16.x` `compatible_runtimes` value ([#24768](https://github.com/hashicorp/terraform-provider-aws/issues/24768))
* resource/aws_organizations_account: Add `create_govcloud` argument and `govcloud_id` attribute ([#24447](https://github.com/hashicorp/terraform-provider-aws/issues/24447))
* resource/aws_s3_bucket_website_configuration: Add `routing_rules` parameter to be used instead of `routing_rule` to support configurations with empty String values ([#24198](https://github.com/hashicorp/terraform-provider-aws/issues/24198))

BUG FIXES:

* resource/aws_autoscaling_group: Wait for correct number of ELBs when `wait_for_elb_capacity` is configured ([#20806](https://github.com/hashicorp/terraform-provider-aws/issues/20806))
* resource/aws_elasticache_replication_group: Fix perpetual diff on `auto_minor_version_upgrade` ([#24688](https://github.com/hashicorp/terraform-provider-aws/issues/24688))

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
* data-source/aws_autoscaling_group: Add `enabled_metrics` attribute ([#24691](https://github.com/hashicorp/terraform-provider-aws/issues/24691))
* data-source/aws_codestarconnections_connection: Support lookup by `name` ([#19262](https://github.com/hashicorp/terraform-provider-aws/issues/19262))
* data-source/aws_launch_template: Add `instance_requirements` attribute ([#24543](https://github.com/hashicorp/terraform-provider-aws/issues/24543))
* resource/aws_ebs_volume: Add support for `multi_attach_enabled` with `io2` volumes ([#19060](https://github.com/hashicorp/terraform-provider-aws/issues/19060))
* resource/aws_launch_template: Add `instance_requirements` argument ([#24543](https://github.com/hashicorp/terraform-provider-aws/issues/24543))
* resource/aws_servicecatalog_provisioned_product: Wait for provisioning to finish ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))
* resource/aws_servicecatalog_provisioned_product: Wait for update to finish ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))
* resource/aws_spot_fleet_request: Add `overrides.instance_requirements` argument ([#24448](https://github.com/hashicorp/terraform-provider-aws/issues/24448))

BUG FIXES:

* resource/aws_alb_listener_rule: Don't force recreate listener rule on priority change. ([#23768](https://github.com/hashicorp/terraform-provider-aws/issues/23768))
* resource/aws_default_subnet: Fix `InvalidSubnet.Conflict` errors when associating IPv6 CIDR blocks ([#24685](https://github.com/hashicorp/terraform-provider-aws/issues/24685))
* resource/aws_ebs_volume: Add configurable timeouts ([#24745](https://github.com/hashicorp/terraform-provider-aws/issues/24745))
* resource/aws_imagebuilder_image_recipe: Fix `ResourceDependencyException` errors when a dependency is modified ([#24708](https://github.com/hashicorp/terraform-provider-aws/issues/24708))
* resource/aws_kms_key: Retry on `MalformedPolicyDocumentException` errors when updating key policy ([#24697](https://github.com/hashicorp/terraform-provider-aws/issues/24697))
* resource/aws_servicecatalog_provisioned_product: Prevent error when retrieving a provisioned product in a non-available state ([#24758](https://github.com/hashicorp/terraform-provider-aws/issues/24758))
* resource/aws_subnet: Fix `InvalidSubnet.Conflict` errors when associating IPv6 CIDR blocks ([#24685](https://github.com/hashicorp/terraform-provider-aws/issues/24685))

## 4.13.0 (May  5, 2022)

FEATURES:

* **New Data Source:** `aws_emrcontainers_virtual_cluster` ([#20003](https://github.com/hashicorp/terraform-provider-aws/issues/20003))
* **New Data Source:** `aws_iam_instance_profiles` ([#24423](https://github.com/hashicorp/terraform-provider-aws/issues/24423))
* **New Data Source:** `aws_secretsmanager_secrets` ([#24514](https://github.com/hashicorp/terraform-provider-aws/issues/24514))
* **New Resource:** `aws_emrcontainers_virtual_cluster` ([#20003](https://github.com/hashicorp/terraform-provider-aws/issues/20003))
* **New Resource:** `aws_iot_topic_rule_destination` ([#24395](https://github.com/hashicorp/terraform-provider-aws/issues/24395))

ENHANCEMENTS:

* data-source/aws_ami: Add `deprecation_time` attribute ([#24489](https://github.com/hashicorp/terraform-provider-aws/issues/24489))
* data-source/aws_msk_cluster: Add `bootstrap_brokers_public_sasl_iam`, `bootstrap_brokers_public_sasl_scram` and `bootstrap_brokers_public_tls` attributes ([#21005](https://github.com/hashicorp/terraform-provider-aws/issues/21005))
* data-source/aws_ssm_patch_baseline: Add the following attributes: `approved_patches`, `approved_patches_compliance_level`, `approval_rule`, `global_filter`, `rejected_patches`, `rejected_patches_action`, `source` ([#24401](https://github.com/hashicorp/terraform-provider-aws/issues/24401))
* resource/aws_ami: Add `deprecation_time` argument ([#24489](https://github.com/hashicorp/terraform-provider-aws/issues/24489))
* resource/aws_ami_copy: Add `deprecation_time` argument ([#24489](https://github.com/hashicorp/terraform-provider-aws/issues/24489))
* resource/aws_ami_from_instance: Add `deprecation_time` argument ([#24489](https://github.com/hashicorp/terraform-provider-aws/issues/24489))
* resource/aws_iot_topic_rule: Add `http` and `error_action.http` arguments ([#16087](https://github.com/hashicorp/terraform-provider-aws/issues/16087))
* resource/aws_iot_topic_rule: Add `kafka` and `error_action.kafka` arguments ([#24395](https://github.com/hashicorp/terraform-provider-aws/issues/24395))
* resource/aws_iot_topic_rule: Add `s3.canned_acl` and `error_action.s3.canned_acl` arguments ([#19175](https://github.com/hashicorp/terraform-provider-aws/issues/19175))
* resource/aws_iot_topic_rule: Add `timestream` and `error_action.timestream` arguments ([#22337](https://github.com/hashicorp/terraform-provider-aws/issues/22337))
* resource/aws_lambda_permission: Add `function_url_auth_type` argument ([#24510](https://github.com/hashicorp/terraform-provider-aws/issues/24510))
* resource/aws_msk_cluster: Add `bootstrap_brokers_public_sasl_iam`, `bootstrap_brokers_public_sasl_scram` and `bootstrap_brokers_public_tls` attributes ([#21005](https://github.com/hashicorp/terraform-provider-aws/issues/21005))
* resource/aws_msk_cluster: Add `broker_node_group_info.connectivity_info` argument to support [public access](https://docs.aws.amazon.com/msk/latest/developerguide/public-access.html) ([#21005](https://github.com/hashicorp/terraform-provider-aws/issues/21005))
* resource/aws_msk_cluster: Add `client_authentication.unauthenticated` argument ([#21005](https://github.com/hashicorp/terraform-provider-aws/issues/21005))
* resource/aws_msk_cluster: Allow in-place update of `client_authentication` and `encryption_info.encryption_in_transit.client_broker` ([#21005](https://github.com/hashicorp/terraform-provider-aws/issues/21005))

BUG FIXES:

* resource/aws_cloudfront_distribution: Fix PreconditionFailed errors when other CloudFront resources are changed before the distribution ([#24537](https://github.com/hashicorp/terraform-provider-aws/issues/24537))
* resource/aws_ecs_service: Fix retry when using the `wait_for_steady_state` parameter ([#24541](https://github.com/hashicorp/terraform-provider-aws/issues/24541))
* resource/aws_launch_template: Fix crash when reading `license_specification` ([#24579](https://github.com/hashicorp/terraform-provider-aws/issues/24579))
* resource/aws_ssm_document: Always include `attachment_sources` when updating SSM documents ([#24530](https://github.com/hashicorp/terraform-provider-aws/issues/24530))

## 4.12.1 (April 29, 2022)

ENHANCEMENTS:

* resource/aws_kms_key: Add support for HMAC_256 customer master key spec ([#24450](https://github.com/hashicorp/terraform-provider-aws/issues/24450))

BUG FIXES:

* resource/aws_acm_certificate_validation: Restore certificate issuance timestamp as the resource `id` value, fixing error on existing resource Read ([#24453](https://github.com/hashicorp/terraform-provider-aws/issues/24453))
* resource/aws_kms_alias: Fix reserved prefix used in `name` and `name_prefix` plan time validation ([#24469](https://github.com/hashicorp/terraform-provider-aws/issues/24469))

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
* data-source/aws_launch_template: Add `maintenance_options` attribute ([#24377](https://github.com/hashicorp/terraform-provider-aws/issues/24377))
* provider: Add support for Assume Role with Web Identity. ([#24441](https://github.com/hashicorp/terraform-provider-aws/issues/24441))
* resource/aws_acm_certificate: Add `validation_option` argument ([#3853](https://github.com/hashicorp/terraform-provider-aws/issues/3853))
* resource/aws_acm_certificate_validation: Increase default resource Create (certificate issuance) timeout to 75 minutes ([#20073](https://github.com/hashicorp/terraform-provider-aws/issues/20073))
* resource/aws_emr_cluster: Add `list_steps_states` argument ([#20871](https://github.com/hashicorp/terraform-provider-aws/issues/20871))
* resource/aws_grafana_workspace: Add `tags` argument ([#24358](https://github.com/hashicorp/terraform-provider-aws/issues/24358))
* resource/aws_instance: Add `maintenance_options` argument ([#24377](https://github.com/hashicorp/terraform-provider-aws/issues/24377))
* resource/aws_launch_template: Add `maintenance_options` argument ([#24377](https://github.com/hashicorp/terraform-provider-aws/issues/24377))
* resource/aws_mq_broker: Make `maintenance_window_start_time` updateable without recreation. ([#24385](https://github.com/hashicorp/terraform-provider-aws/issues/24385))
* resource/aws_rds_cluster: Add `serverlessv2_scaling_configuration` argument to support [Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html) ([#24363](https://github.com/hashicorp/terraform-provider-aws/issues/24363))
* resource/aws_spot_fleet_request: Add `terminate_instances_on_delete` argument ([#17268](https://github.com/hashicorp/terraform-provider-aws/issues/17268))

BUG FIXES:

* data-source/aws_kms_alias: Fix `name` plan time validation ([#13000](https://github.com/hashicorp/terraform-provider-aws/issues/13000))
* provider: Setting `skip_metadata_api_check = false` now overrides `AWS_EC2_METADATA_DISABLED` environment variable. ([#24441](https://github.com/hashicorp/terraform-provider-aws/issues/24441))
* resource/aws_acm_certificate: Correctly handle SAN entries that match `domain_name` ([#20073](https://github.com/hashicorp/terraform-provider-aws/issues/20073))
* resource/aws_dms_replication_task: Fix to stop the task before updating, if required ([#24047](https://github.com/hashicorp/terraform-provider-aws/issues/24047))
* resource/aws_ec2_availability_zone_group: Don't crash if `group_name` is not found ([#24422](https://github.com/hashicorp/terraform-provider-aws/issues/24422))
* resource/aws_elasticache_cluster: Update regex pattern to target specific Redis V6 versions through the `engine_version` attribute ([#23734](https://github.com/hashicorp/terraform-provider-aws/issues/23734))
* resource/aws_elasticache_replication_group: Update regex pattern to target specific Redis V6 versions through the `engine_version` attribute ([#23734](https://github.com/hashicorp/terraform-provider-aws/issues/23734))
* resource/aws_kms_alias: Fix `name` and `name_prefix` plan time validation ([#13000](https://github.com/hashicorp/terraform-provider-aws/issues/13000))
* resource/aws_lb: Fix bug causing an error on update if tags unsupported in ISO region ([#24334](https://github.com/hashicorp/terraform-provider-aws/issues/24334))
* resource/aws_s3_bucket_policy: Let resource be removed from tfstate if bucket deleted outside Terraform ([#23510](https://github.com/hashicorp/terraform-provider-aws/issues/23510))
* resource/aws_s3_bucket_versioning: Let resource be removed from tfstate if bucket deleted outside Terraform ([#23510](https://github.com/hashicorp/terraform-provider-aws/issues/23510))
* resource/aws_ses_receipt_filter: Allow period character (`.`) in `name` argument ([#24383](https://github.com/hashicorp/terraform-provider-aws/issues/24383))

## 4.11.0 (April 22, 2022)

FEATURES:

* **New Data Source:** `aws_s3_bucket_policy` ([#17738](https://github.com/hashicorp/terraform-provider-aws/issues/17738))
* **New Resource:** `aws_transfer_workflow` ([#24248](https://github.com/hashicorp/terraform-provider-aws/issues/24248))

ENHANCEMENTS:

* data-source/aws_imagebuilder_infrastructure_configuration: Add `instance_metadata_options` attribute ([#24285](https://github.com/hashicorp/terraform-provider-aws/issues/24285))
* data-source/aws_opensearch_domain: Add `cold_storage_options` attribute to the `cluster_config` configuration block ([#24284](https://github.com/hashicorp/terraform-provider-aws/issues/24284))
* resource/aws_db_proxy: Add `auth.username` argument ([#24264](https://github.com/hashicorp/terraform-provider-aws/issues/24264))
* resource/aws_elasticache_user: Add plan-time validation of password argumnet length ([#24274](https://github.com/hashicorp/terraform-provider-aws/issues/24274))
* resource/aws_elasticsearch_domain: For Elasticsearch versions 6.7+, allow in-place update of `node_to_node_encryption.0.enabled` and `encrypt_at_rest.0.enabled`. ([#24222](https://github.com/hashicorp/terraform-provider-aws/issues/24222))
* resource/aws_fsx_ontap_file_system: Add support for `SINGLE_AZ_1` `deployment_type`. ([#24280](https://github.com/hashicorp/terraform-provider-aws/issues/24280))
* resource/aws_imagebuilder_infrastructure_configuration: Add `instance_metadata_options` argument ([#24285](https://github.com/hashicorp/terraform-provider-aws/issues/24285))
* resource/aws_instance: Add `capacity_reservation_specification.capacity_reservation_target.capacity_reservation_resource_group_arn` argument ([#24283](https://github.com/hashicorp/terraform-provider-aws/issues/24283))
* resource/aws_instance: Add `network_interface.network_card_index` argument ([#24283](https://github.com/hashicorp/terraform-provider-aws/issues/24283))
* resource/aws_opensearch_domain: Add `cold_storage_options` argument to the `cluster_config` configuration block ([#24284](https://github.com/hashicorp/terraform-provider-aws/issues/24284))
* resource/aws_opensearch_domain: For Elasticsearch versions 6.7+, allow in-place update of `node_to_node_encryption.0.enabled` and `encrypt_at_rest.0.enabled`. ([#24222](https://github.com/hashicorp/terraform-provider-aws/issues/24222))
* resource/aws_transfer_server: Add `workflow_details` argument ([#24248](https://github.com/hashicorp/terraform-provider-aws/issues/24248))
* resource/aws_waf_byte_match_set: Additional supported values for `byte_match_tuples.field_to_match.type` argument ([#24286](https://github.com/hashicorp/terraform-provider-aws/issues/24286))
* resource/aws_wafregional_web_acl: Additional supported values for `logging_configuration.redacted_fields.field_to_match.type` argument ([#24286](https://github.com/hashicorp/terraform-provider-aws/issues/24286))
* resource/aws_workspaces_workspace: Additional supported values for `workspace_properties.compute_type_name` argument ([#24286](https://github.com/hashicorp/terraform-provider-aws/issues/24286))

BUG FIXES:

* data-source/aws_db_instance: Prevent panic when setting instance connection endpoint values ([#24299](https://github.com/hashicorp/terraform-provider-aws/issues/24299))
* data-source/aws_efs_file_system: Prevent panic when searching by tag returns 0 or multiple results ([#24298](https://github.com/hashicorp/terraform-provider-aws/issues/24298))
* data-source/aws_elasticache_cluster: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_appstream_user_stack_association: Prevent panic during resource read ([#24303](https://github.com/hashicorp/terraform-provider-aws/issues/24303))
* resource/aws_cloudformation_stack_set: Prevent `Validation` errors when `operation_preferences.failure_tolerance_count` is zero ([#24250](https://github.com/hashicorp/terraform-provider-aws/issues/24250))
* resource/aws_elastic_beanstalk_environment: Correctly set `cname_prefix` attribute ([#24278](https://github.com/hashicorp/terraform-provider-aws/issues/24278))
* resource/aws_elasticache_cluster: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_elasticache_parameter_group: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_elasticache_replication_group: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_elasticache_subnet_group: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_elasticache_user: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_elasticache_user_group: Gracefully handle additional tagging error type in non-standard AWS partitions (i.e., ISO) ([#24275](https://github.com/hashicorp/terraform-provider-aws/issues/24275))
* resource/aws_instance: Fix issue with assuming Placement and disableApiTermination instance attributes exist when managing a Snowball Edge device ([#19256](https://github.com/hashicorp/terraform-provider-aws/issues/19256))
* resource/aws_kinesis_firehose_delivery_stream: Increase the maximum length of the `processing_configuration.processors.parameters.parameter_value` argument's value to `5120` ([#24312](https://github.com/hashicorp/terraform-provider-aws/issues/24312))
* resource/aws_macie2_member: Correct type for `invitation_disable_email_notification` parameter ([#24304](https://github.com/hashicorp/terraform-provider-aws/issues/24304))
* resource/aws_s3_bucket_server_side_encryption_configuration: Retry on `ServerSideEncryptionConfigurationNotFoundError` errors due to eventual consistency ([#24266](https://github.com/hashicorp/terraform-provider-aws/issues/24266))
* resource/aws_sfn_state_machine: Prevent panic during resource update ([#24302](https://github.com/hashicorp/terraform-provider-aws/issues/24302))
* resource/aws_shield_protection_group: When updating resource tags, use the `protection_group_arn` parameter instead of `arn`. ([#24296](https://github.com/hashicorp/terraform-provider-aws/issues/24296))
* resource/aws_ssm_association: Prevent panic when `wait_for_success_timeout_seconds` is configured ([#24300](https://github.com/hashicorp/terraform-provider-aws/issues/24300))

## 4.10.0 (April 14, 2022)

FEATURES:

* **New Data Source:** `aws_iam_saml_provider` ([#10498](https://github.com/hashicorp/terraform-provider-aws/issues/10498))
* **New Data Source:** `aws_nat_gateways` ([#24190](https://github.com/hashicorp/terraform-provider-aws/issues/24190))
* **New Resource:** `aws_datasync_location_fsx_openzfs_file_system` ([#24200](https://github.com/hashicorp/terraform-provider-aws/issues/24200))
* **New Resource:** `aws_elasticache_user_group_association` ([#24204](https://github.com/hashicorp/terraform-provider-aws/issues/24204))
* **New Resource:** `aws_qldb_stream` ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))

ENHANCEMENTS:

* data-source/aws_qldb_ledger: Add `kms_key` and `tags` attributes ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))
* resource/aws_ami_launch_permission: Add `group` argument ([#20677](https://github.com/hashicorp/terraform-provider-aws/issues/20677))
* resource/aws_ami_launch_permission: Add `organization_arn` and `organizational_unit_arn` arguments ([#21694](https://github.com/hashicorp/terraform-provider-aws/issues/21694))
* resource/aws_athena_database: Add `properties` argument. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))
* resource/aws_athena_database: Add import support. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))
* resource/aws_config_config_rule: Add `source.custom_policy_details` argument. ([#24057](https://github.com/hashicorp/terraform-provider-aws/issues/24057))
* resource/aws_config_config_rule: Add plan time validation for `source.source_detail.event_source` and `source.source_detail.message_type`. ([#24057](https://github.com/hashicorp/terraform-provider-aws/issues/24057))
* resource/aws_config_config_rule: Make `source.source_identifier` optional. ([#24057](https://github.com/hashicorp/terraform-provider-aws/issues/24057))
* resource/aws_eks_addon: Add `preserve` argument ([#24218](https://github.com/hashicorp/terraform-provider-aws/issues/24218))
* resource/aws_grafana_workspace: Add plan time validations for `authentication_providers`, `authentication_providers`, `authentication_providers`. ([#24170](https://github.com/hashicorp/terraform-provider-aws/issues/24170))
* resource/aws_qldb_ledger: Add `kms_key` argument ([#19297](https://github.com/hashicorp/terraform-provider-aws/issues/19297))
* resource/aws_vpc_ipam_scope: Add pagination when describing IPAM Scopes ([#24188](https://github.com/hashicorp/terraform-provider-aws/issues/24188))

BUG FIXES:

* resource/aws_athena_database: Add drift detection for `comment`. ([#24172](https://github.com/hashicorp/terraform-provider-aws/issues/24172))
* resource/aws_cloudformation_stack_set: Prevent `InvalidParameter` errors when updating `operation_preferences` ([#24202](https://github.com/hashicorp/terraform-provider-aws/issues/24202))
* resource/aws_cloudwatch_event_connection: Add validation to `auth_parameters.api_key.key`, `auth_parameters.api_key.value`, `auth_parameters.basic.username`, `auth_parameters.basic.password`, `auth_parameters.oauth.authorization_endpoint`, `auth_parameters.oauth.client_parameters.client_id` and `auth_parameters.oauth.client_parameters.client_secret` arguments ([#24154](https://github.com/hashicorp/terraform-provider-aws/issues/24154))
* resource/aws_cloudwatch_log_subscription_filter: Retry resource create and update when a conflicting operation error is returned ([#24148](https://github.com/hashicorp/terraform-provider-aws/issues/24148))
* resource/aws_ecs_service: Retry when using the `wait_for_steady_state` parameter and `ResourceNotReady` errors are returned from the AWS API ([#24223](https://github.com/hashicorp/terraform-provider-aws/issues/24223))
* resource/aws_ecs_service: Wait for service to reach an active state after create and update operations ([#24223](https://github.com/hashicorp/terraform-provider-aws/issues/24223))
* resource/aws_emr_cluster: Ignore `UnknownOperationException` errors when reading a cluster's auto-termination policy ([#24237](https://github.com/hashicorp/terraform-provider-aws/issues/24237))
* resource/aws_lambda_function_url: Ignore `ResourceConflictException` errors caused by existing `FunctionURLAllowPublicAccess` permission statements ([#24220](https://github.com/hashicorp/terraform-provider-aws/issues/24220))
* resource/aws_vpc_ipam_scope: Prevent panic when describing IPAM Scopes by ID ([#24188](https://github.com/hashicorp/terraform-provider-aws/issues/24188))

## 4.9.0 (April 07, 2022)

NOTES:

* resource/aws_s3_bucket: The `acceleration_status`, `acl`, `cors_rule`, `grant`, `lifecycle_rule`, `logging`, `object_lock_configuration.rule`, `policy`, `replication_configuration`, `request_payer`, `server_side_encryption_configuration`, `versioning`, and `website` parameters are now Optional. Please refer to the documentation for details on drift detection and potential conflicts when configuring these parameters with the standalone `aws_s3_bucket_*` resources. ([#23985](https://github.com/hashicorp/terraform-provider-aws/issues/23985))

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

* data-source/aws_imagebuilder_distribution_configuration: Add `account_id` attribute to the `launch_template_configuration` attribute of the `distribution` configuration block ([#23924](https://github.com/hashicorp/terraform-provider-aws/issues/23924))
* data-source/aws_route: Add `core_network_arn` argument ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* data-source/aws_route_table: Add 'routes.core_network_arn' attribute' ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* provider: Add support for reading custom CA bundle setting from shared config files ([#24064](https://github.com/hashicorp/terraform-provider-aws/issues/24064))
* resource/aws_cloudformation_stack_set: Add `operation_preferences` argument ([#23908](https://github.com/hashicorp/terraform-provider-aws/issues/23908))
* resource/aws_default_route_table: Add `core_network_arn` argument to the `route` configuration block ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_dlm_lifecycle_policy: Add `policy_details.schedule.create_rule.cron_expression`, `policy_details.schedule.retain_rule.interval`, `policy_details.schedule.retain_rule.interval_unit`, `policy_details.policy_type`, `policy_details.schedule.deprecate_rule`, `policy_details.parameters`, `policy_details.schedule.variable_tags`, `policy_details.schedule.fast_restore_rule`, `policy_details.schedule.share_rule`, `policy_details.resource_locations`, `policy_details.schedule.create_rule.location`, `policy_details.action` and `policy_details.event_source` arguments ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_dlm_lifecycle_policy: Add plan time validations for `policy_details.resource_types` and `description` arguments ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_dlm_lifecycle_policy: Make `policy_details.resource_types`, `policy_details.schedule`, `policy_details.target_tags`, `policy_details.schedule.retain_rule` and `policy_details.schedule.create_rule.interval` arguments optional ([#23880](https://github.com/hashicorp/terraform-provider-aws/issues/23880))
* resource/aws_elasticache_cluster: Add `auto_minor_version_upgrade` argument ([#23996](https://github.com/hashicorp/terraform-provider-aws/issues/23996))
* resource/aws_fms_policy: Retry when `InternalErrorException` errors are returned from the AWS API ([#23952](https://github.com/hashicorp/terraform-provider-aws/issues/23952))
* resource/aws_fsx_ontap_file_system: Support updating `storage_capacity`, `throughput_capacity`, and `disk_iops_configuration`. ([#24002](https://github.com/hashicorp/terraform-provider-aws/issues/24002))
* resource/aws_imagebuilder_distribution_configuration: Add `account_id` argument to the `launch_template_configuration` attribute of the `distribution` configuration block ([#23924](https://github.com/hashicorp/terraform-provider-aws/issues/23924))
* resource/aws_iot_authorizer: Add `enable_caching_for_http` argument ([#23993](https://github.com/hashicorp/terraform-provider-aws/issues/23993))
* resource/aws_lambda_permission: Add `principal_org_id` argument. ([#24001](https://github.com/hashicorp/terraform-provider-aws/issues/24001))
* resource/aws_mq_broker: Add validation to `broker_name` and `security_groups` arguments ([#18088](https://github.com/hashicorp/terraform-provider-aws/issues/18088))
* resource/aws_organizations_account: Add `close_on_deletion` argument to close account on deletion ([#23930](https://github.com/hashicorp/terraform-provider-aws/issues/23930))
* resource/aws_route: Add `core_network_arn` argument ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_route_table: Add `core_network_arn` argument to the `route` configuration block ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_s3_bucket: Speed up resource deletion, especially when the S3 buckets contains a large number of objects and `force_destroy` is `true` ([#24020](https://github.com/hashicorp/terraform-provider-aws/issues/24020))
* resource/aws_s3_bucket: Update `acceleration_status` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_accelerate_configuration` resource. ([#23816](https://github.com/hashicorp/terraform-provider-aws/issues/23816))
* resource/aws_s3_bucket: Update `acl` and `grant` parameters to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring these parameters with the standalone `aws_s3_bucket_acl` resource. ([#23798](https://github.com/hashicorp/terraform-provider-aws/issues/23798))
* resource/aws_s3_bucket: Update `cors_rule` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_cors_configuration` resource. ([#23817](https://github.com/hashicorp/terraform-provider-aws/issues/23817))
* resource/aws_s3_bucket: Update `lifecycle_rule` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_lifecycle_configuration` resource. ([#23818](https://github.com/hashicorp/terraform-provider-aws/issues/23818))
* resource/aws_s3_bucket: Update `logging` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_logging` resource. ([#23819](https://github.com/hashicorp/terraform-provider-aws/issues/23819))
* resource/aws_s3_bucket: Update `object_lock_configuration.rule` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_object_lock_configuration` resource. ([#23984](https://github.com/hashicorp/terraform-provider-aws/issues/23984))
* resource/aws_s3_bucket: Update `policy` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_policy` resource. ([#23843](https://github.com/hashicorp/terraform-provider-aws/issues/23843))
* resource/aws_s3_bucket: Update `replication_configuration` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_replication_configuration` resource. ([#23842](https://github.com/hashicorp/terraform-provider-aws/issues/23842))
* resource/aws_s3_bucket: Update `request_payer` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_request_payment_configuration` resource. ([#23844](https://github.com/hashicorp/terraform-provider-aws/issues/23844))
* resource/aws_s3_bucket: Update `server_side_encryption_configuration` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_server_side_encryption_configuration` resource. ([#23822](https://github.com/hashicorp/terraform-provider-aws/issues/23822))
* resource/aws_s3_bucket: Update `versioning` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_versioning` resource. ([#23820](https://github.com/hashicorp/terraform-provider-aws/issues/23820))
* resource/aws_s3_bucket: Update `website` parameter to be configurable. Please refer to the documentation for details on drift detection and potential conflicts when configuring this parameter with the standalone `aws_s3_bucket_website_configuration` resource. ([#23821](https://github.com/hashicorp/terraform-provider-aws/issues/23821))
* resource/aws_storagegateway_gateway: Add `maintenance_start_time` argument ([#15355](https://github.com/hashicorp/terraform-provider-aws/issues/15355))
* resource/aws_storagegateway_nfs_file_share: Add `bucket_region` and `vpc_endpoint_dns_name` arguments to support PrivateLink endpoints ([#24038](https://github.com/hashicorp/terraform-provider-aws/issues/24038))
* resource/aws_vpc_ipam: add `cascade` argument ([#23973](https://github.com/hashicorp/terraform-provider-aws/issues/23973))
* resource/aws_vpn_connection: Add `core_network_arn` and `core_network_attachment_arn` attributes ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_xray_group: Add `insights_configuration` argument ([#24028](https://github.com/hashicorp/terraform-provider-aws/issues/24028))

BUG FIXES:

* data-source/aws_elasticache_cluster: Allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_backup_report_plan: Wait for asynchronous lifecycle operations to complete ([#23967](https://github.com/hashicorp/terraform-provider-aws/issues/23967))
* resource/aws_cloudformation_stack_set: Consider `QUEUED` a valid pending state for resource creation ([#22160](https://github.com/hashicorp/terraform-provider-aws/issues/22160))
* resource/aws_dynamodb_table_item: Allow `item` names to still succeed if they include non-letters ([#14075](https://github.com/hashicorp/terraform-provider-aws/issues/14075))
* resource/aws_elasticache_cluster: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticache_parameter_group: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticache_replication_group: Allow disabling `auto_minor_version_upgrade` ([#23996](https://github.com/hashicorp/terraform-provider-aws/issues/23996))
* resource/aws_elasticache_replication_group: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticache_replication_group: Waits for available state before updating tags ([#24021](https://github.com/hashicorp/terraform-provider-aws/issues/24021))
* resource/aws_elasticache_subnet_group: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticache_user: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticache_user_group: Attempt `tags`-on-create, fallback to tag after create, and allow some `tags` errors to be non-fatal to support non-standard AWS partitions (i.e., ISO) ([#23979](https://github.com/hashicorp/terraform-provider-aws/issues/23979))
* resource/aws_elasticsearch_domain_saml_option: Fix difference caused by `subject_key` default not matching AWS default; old and new defaults are equivalent ([#20892](https://github.com/hashicorp/terraform-provider-aws/issues/20892))
* resource/aws_lb: Fix attribute key not recognized issue preventing creation in ISO-B regions ([#23972](https://github.com/hashicorp/terraform-provider-aws/issues/23972))
* resource/aws_redshift_cluster: Correctly use `number_of_nodes` argument value when restoring from snapshot ([#13203](https://github.com/hashicorp/terraform-provider-aws/issues/13203))
* resource/aws_route: Ensure that resource ID is set in case of wait-for-creation time out ([#24024](https://github.com/hashicorp/terraform-provider-aws/issues/24024))
* resource/aws_s3_bucket_lifecycle_configuration: Prevent `MalformedXML` errors when handling diffs in `rule.filter` ([#23893](https://github.com/hashicorp/terraform-provider-aws/issues/23893))

## 4.8.0 (March 25, 2022)

FEATURES:

* **New Data Source:** `aws_mskconnect_connector` ([#23792](https://github.com/hashicorp/terraform-provider-aws/issues/23544))
* **New Resource:** `aws_mskconnect_connector` ([#23765](https://github.com/hashicorp/terraform-provider-aws/issues/23544))

ENHANCEMENTS:

* data-source/aws_eips: Set `public_ips` for VPC as well as EC2 Classic ([#23859](https://github.com/hashicorp/terraform-provider-aws/issues/23859))
* data-source/aws_elasticache_cluster: Add `log_delivery_configuration` attribute ([#20068](https://github.com/hashicorp/terraform-provider-aws/issues/20068))
* data-source/aws_elasticache_replication_group: Add `log_delivery_configuration` attribute ([#20068](https://github.com/hashicorp/terraform-provider-aws/issues/20068))
* data-source/aws_elasticsearch_domain: Add `cold_storage_options` attribute to the `cluster_config` configuration block ([#19713](https://github.com/hashicorp/terraform-provider-aws/issues/19713))
* data-source/aws_lambda_function: Add `ephemeral_storage` attribute ([#23873](https://github.com/hashicorp/terraform-provider-aws/issues/23873))
* resource/aws_elasticache_cluster: Add `log_delivery_configuration` argument ([#20068](https://github.com/hashicorp/terraform-provider-aws/issues/20068))
* resource/aws_elasticache_replication_group: Add `log_delivery_configuration` argument ([#20068](https://github.com/hashicorp/terraform-provider-aws/issues/20068))
* resource/aws_elasticsearch_domain: Add `cold_storage_options` argument to the `cluster_config` configuration block ([#19713](https://github.com/hashicorp/terraform-provider-aws/issues/19713))
* resource/aws_elasticsearch_domain: Add configurable Create and Delete timeouts ([#19713](https://github.com/hashicorp/terraform-provider-aws/issues/19713))
* resource/aws_lambda_function: Add `ephemeral_storage` argument ([#23873](https://github.com/hashicorp/terraform-provider-aws/issues/23873))
* resource/aws_lambda_function: Add error handling for `ResourceConflictException` errors on create and update ([#23879](https://github.com/hashicorp/terraform-provider-aws/issues/23879))
* resource/aws_mskconnect_custom_plugin: Implement resource Delete ([#23544](https://github.com/hashicorp/terraform-provider-aws/issues/23544))
* resource/aws_mwaa_environment: Add `schedulers` argument ([#21941](https://github.com/hashicorp/terraform-provider-aws/issues/21941))
* resource/aws_network_firewall_policy: Allow use of managed rule group arns for network firewall managed rule groups. ([#22355](https://github.com/hashicorp/terraform-provider-aws/issues/22355))

BUG FIXES:

* resource/aws_autoscaling_group: Fix issue where group was not recreated if `initial_lifecycle_hook` changed ([#20708](https://github.com/hashicorp/terraform-provider-aws/issues/20708))
* resource/aws_cloudfront_distribution: Fix default value of `origin_path` in `origin` block ([#20709](https://github.com/hashicorp/terraform-provider-aws/issues/20709))
* resource/aws_cloudwatch_event_target: Fix setting `path_parameter_values`. ([#23862](https://github.com/hashicorp/terraform-provider-aws/issues/23862))

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

* data-source/aws_imagebuilder_distribution_configuration: Add `organization_arns` and `organizational_unit_arns` attributes to the `distribution.ami_distribution_configuration.launch_permission` configuration block ([#22104](https://github.com/hashicorp/terraform-provider-aws/issues/22104))
* data-source/aws_msk_cluster: Add `zookeeper_connect_string_tls` attribute ([#23804](https://github.com/hashicorp/terraform-provider-aws/issues/23804))
* data-source/aws_ssm_document: Support `TEXT` `document_format` ([#23757](https://github.com/hashicorp/terraform-provider-aws/issues/23757))
* resource/aws_api_gateway_stage: Add `canary_settings` argument. ([#23754](https://github.com/hashicorp/terraform-provider-aws/issues/23754))
* resource/aws_athena_workgroup: Add `acl_configuration` and `expected_bucket_owner` arguments to the `configuration.result_configuration` block ([#23748](https://github.com/hashicorp/terraform-provider-aws/issues/23748))
* resource/aws_autoscaling_group: Add `instance_reuse_policy` argument to support [Warm Pool scale-in](https://aws.amazon.com/about-aws/whats-new/2022/02/amazon-ec2-auto-scaling-warm-pools-supports-hibernating-returning-instances-warm-pools-scale-in/) ([#23769](https://github.com/hashicorp/terraform-provider-aws/issues/23769))
* resource/aws_autoscaling_group: Update documentation to include [Warm Pool hibernation](https://aws.amazon.com/about-aws/whats-new/2022/02/amazon-ec2-auto-scaling-warm-pools-supports-hibernating-returning-instances-warm-pools-scale-in/) ([#23772](https://github.com/hashicorp/terraform-provider-aws/issues/23772))
* resource/aws_cloudformation_stack_set_instance: Add `operation_preferences` argument ([#23666](https://github.com/hashicorp/terraform-provider-aws/issues/23666))
* resource/aws_cloudwatch_log_subscription_filter: Add plan time validations for `name`, `destination_arn`, `filter_pattern`, `role_arn`, `distribution`. ([#23760](https://github.com/hashicorp/terraform-provider-aws/issues/23760))
* resource/aws_glue_schema: Update documentation to include [Protobuf data format support](https://aws.amazon.com/about-aws/whats-new/2022/02/aws-glue-schema-registry-protocol-buffers) ([#23815](https://github.com/hashicorp/terraform-provider-aws/issues/23815))
* resource/aws_imagebuilder_distribution_configuration: Add `organization_arns` and `organizational_unit_arns` arguments to the `distribution.ami_distribution_configuration.launch_permission` configuration block ([#22104](https://github.com/hashicorp/terraform-provider-aws/issues/22104))
* resource/aws_instance: Add `user_data_replace_on_change` attribute ([#23604](https://github.com/hashicorp/terraform-provider-aws/issues/23604))
* resource/aws_ssm_maintenance_window_task: Add `arn` and `window_task_id` attributes. ([#23756](https://github.com/hashicorp/terraform-provider-aws/issues/23756))
* resource/aws_ssm_maintenance_window_task: Add `cutoff_behavior` argument. ([#23756](https://github.com/hashicorp/terraform-provider-aws/issues/23756))

BUG FIXES:

* data-source/aws_ssm_document: Dont generate `arn` for AWS managed docs. ([#23757](https://github.com/hashicorp/terraform-provider-aws/issues/23757))
* resource/aws_ecs_service: Ensure that `load_balancer` and `service_registries` can be updated in-place ([#23786](https://github.com/hashicorp/terraform-provider-aws/issues/23786))
* resource/aws_launch_template: Fix `network_interfaces.device_index` and `network_interfaces.network_card_index` of `0` not being set ([#23767](https://github.com/hashicorp/terraform-provider-aws/issues/23767))
* resource/aws_ssm_maintenance_window_task: Allow creating a window taks without targets. ([#23756](https://github.com/hashicorp/terraform-provider-aws/issues/23756))

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
* data-source/aws_imagebuilder_image: Add `container_recipe_arn` attribute ([#23647](https://github.com/hashicorp/terraform-provider-aws/issues/23647))
* data-source/aws_launch_template: Add `capacity_reservation_resource_group_arn` attribute to the `capacity_reservation_specification.capacity_reservation_target` configuration block ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* data-source/aws_launch_template: Add `capacity_reservation_specification`, `cpu_options`, `elastic_inference_accelerator` and `license_specification` attributes ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* data-source/aws_launch_template: Add `ipv4_prefixes`, `ipv4_prefix_count`, `ipv6_prefixes` and `ipv6_prefix_count` attributes to the `network_interfaces` configuration block ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* data-source/aws_launch_template: Add `private_dns_name_options` attribute ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* data-source/aws_redshift_cluster: Add `availability_zone_relocation_enabled` attribute. ([#20812](https://github.com/hashicorp/terraform-provider-aws/issues/20812))
* resource/aws_appconfig_configuration_profile: Add `type` argument to support [AWS AppConfig Feature Flags](https://aws.amazon.com/blogs/mt/using-aws-appconfig-feature-flags/) ([#23719](https://github.com/hashicorp/terraform-provider-aws/issues/23719))
* resource/aws_athena_database: Add `acl_configuration` and `expected_bucket_owner` arguments ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_athena_database: Add `comment` argument to support database descriptions ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_athena_database: Do not recreate the resource if `bucket` changes ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_cloud9_environment_ec2: Add `connection_type` and `image_id` arguments ([#19195](https://github.com/hashicorp/terraform-provider-aws/issues/19195))
* resource/aws_cloudformation_stack_set:_instance: Add `call_as` argument ([#23339](https://github.com/hashicorp/terraform-provider-aws/issues/23339))
* resource/aws_dms_replication_task: Add optional `start_replication_task` and `status` argument ([#23692](https://github.com/hashicorp/terraform-provider-aws/issues/23692))
* resource/aws_ec2_transit_gateway_connect_peer: Add `arn` attribute ([#13251](https://github.com/hashicorp/terraform-provider-aws/issues/13251))
* resource/aws_ecs_service: `enable_ecs_managed_tags`, `load_balancer`, `propagate_tags` and `service_registries` can now be updated in-place ([#23600](https://github.com/hashicorp/terraform-provider-aws/issues/23600))
* resource/aws_imagebuilder_image: Add `container_recipe_arn` argument ([#23647](https://github.com/hashicorp/terraform-provider-aws/issues/23647))
* resource/aws_iot_certificate: Add `ca_pem` argument, enabling the use of existing IoT certificates ([#23126](https://github.com/hashicorp/terraform-provider-aws/issues/23126))
* resource/aws_iot_topic_rule: Add `cloudwatch_logs` and `error_action.cloudwatch_logs` arguments ([#23440](https://github.com/hashicorp/terraform-provider-aws/issues/23440))
* resource/aws_launch_configuration: Add `ephemeral_block_device.no_device` argument ([#23152](https://github.com/hashicorp/terraform-provider-aws/issues/23152))
* resource/aws_launch_template: Add `capacity_reservation_resource_group_arn` argument to the `capacity_reservation_specification.capacity_reservation_target` configuration block ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* resource/aws_launch_template: Add `ipv4_prefixes`, `ipv4_prefix_count`, `ipv6_prefixes` and `ipv6_prefix_count` arguments to the `network_interfaces` configuration block ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* resource/aws_launch_template: Add `private_dns_name_options` argument ([#23365](https://github.com/hashicorp/terraform-provider-aws/issues/23365))
* resource/aws_msk_configuration: Correctly set `latest_revision` as Computed when `server_properties` changes ([#23662](https://github.com/hashicorp/terraform-provider-aws/issues/23662))
* resource/aws_quicksight_user: Allow custom values for `namespace` ([#23607](https://github.com/hashicorp/terraform-provider-aws/issues/23607))
* resource/aws_rds_cluster: Add `db_cluster_instance_class`, `allocated_storage`, `storage_type`, and `iops` arguments to support [Multi-AZ deployments for MySQL & PostgreSQL](https://aws.amazon.com/blogs/aws/amazon-rds-multi-az-db-cluster/) ([#23684](https://github.com/hashicorp/terraform-provider-aws/issues/23684))
* resource/aws_rds_global_cluster: Add configurable timeouts ([#23560](https://github.com/hashicorp/terraform-provider-aws/issues/23560))
* resource/aws_rds_instance: Add `source_db_instance_automated_backup_arn` option within `restore_to_point_in_time` attribute ([#23086](https://github.com/hashicorp/terraform-provider-aws/issues/23086))
* resource/aws_redshift_cluster: Add `availability_zone_relocation_enabled` attribute and allow `availability_zone` to be changed in-place. ([#20812](https://github.com/hashicorp/terraform-provider-aws/issues/20812))
* resource/aws_transfer_server: Add `pre_authentication_login_banner` and `post_authentication_login_banner` arguments ([#23631](https://github.com/hashicorp/terraform-provider-aws/issues/23631))
* resource/aws_vpc_endpoint: The `security_group_ids` attribute can now be empty when the resource is created. In this case the VPC's default security is associated with the VPC endpoint ([#13737](https://github.com/hashicorp/terraform-provider-aws/issues/13737))

BUG FIXES:

* resource/aws_amplify_app: Allow `repository` to be updated in-place ([#23517](https://github.com/hashicorp/terraform-provider-aws/issues/23517))
* resource/aws_api_gateway_stage: Fixed issue with providing `cache_cluster_size` without `cache_cluster_enabled` resulted in waiter error ([#23091](https://github.com/hashicorp/terraform-provider-aws/issues/23091))
* resource/aws_athena_database: Remove from state on resource Read if deleted outside of Terraform ([#23745](https://github.com/hashicorp/terraform-provider-aws/issues/23745))
* resource/aws_cloudformation_stack_set: Use `call_as` attribute when reading stack sets, fixing an error raised when using a delegated admistrator account ([#23339](https://github.com/hashicorp/terraform-provider-aws/issues/23339))
* resource/aws_cloudsearch_domain: Set correct defaults for `index_field.facet`, `index_field.highlight`, `index_field.return`, `index_field.search` and `index_field.sort`, preventing spurious resource diffs ([#23687](https://github.com/hashicorp/terraform-provider-aws/issues/23687))
* resource/aws_db_instance: Fix issues where configured update timeout was not respected, and update would fail if instance were in the process of being configured. ([#23560](https://github.com/hashicorp/terraform-provider-aws/issues/23560))
* resource/aws_rds_event_subscription: Fix issue where `enabled` was sometimes not updated ([#23560](https://github.com/hashicorp/terraform-provider-aws/issues/23560))
* resource/aws_rds_global_cluster: Fix ability to perform cluster version upgrades, including of clusters in distinct regions, such as previously got error: "Invalid database cluster identifier" ([#23560](https://github.com/hashicorp/terraform-provider-aws/issues/23560))
* resource/aws_route53domains_registered_domain: Redirect all Route 53 Domains AWS API calls to the `us-east-1` Region ([#23672](https://github.com/hashicorp/terraform-provider-aws/issues/23672))
* resource/aws_s3_bucket_acl: Fix resource import for S3 bucket names consisting of uppercase letters, underscores, and a maximum of 255 characters ([#23678](https://github.com/hashicorp/terraform-provider-aws/issues/23678))
* resource/aws_s3_bucket_lifecycle_configuration: Support empty string filtering (default behavior of the `aws_s3_bucket.lifecycle_rule` parameter in provider versions prior to v4.0) ([#23746](https://github.com/hashicorp/terraform-provider-aws/issues/23746))
* resource/aws_s3_bucket_replication_configuration: Change `rule` configuration block to list instead of set ([#23703](https://github.com/hashicorp/terraform-provider-aws/issues/23703))
* resource/aws_s3_bucket_replication_configuration: Set `rule.id` as Computed to prevent drift when the value is not configured ([#23703](https://github.com/hashicorp/terraform-provider-aws/issues/23703))
* resource/aws_s3_bucket_versioning: Add missing support for `Disabled` bucket versioning ([#23723](https://github.com/hashicorp/terraform-provider-aws/issues/23723))

## 4.5.0 (March 11, 2022)

ENHANCEMENTS:

* resource/aws_account_alternate_contact: Add configurable timeouts ([#23516](https://github.com/hashicorp/terraform-provider-aws/issues/23516))
* resource/aws_s3_bucket: Add error handling for `NotImplemented` errors when reading `object_lock_enabled` and `object_lock_configuration` into terraform state. ([#13366](https://github.com/hashicorp/terraform-provider-aws/issues/13366))
* resource/aws_s3_bucket: Add top-level `object_lock_enabled` parameter ([#23556](https://github.com/hashicorp/terraform-provider-aws/issues/23556))
* resource/aws_s3_bucket_replication_configuration: Add `token` field to specify
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
* resource/aws_image_builder_image_recipe: Fix regression in 4.3.0 whereby Windows-based images wouldn't build because of the newly introduced `systems_manager_agent.uninstall_after_build` argument. ([#23580](https://github.com/hashicorp/terraform-provider-aws/issues/23580))
* resource/aws_kms_external_key: Increase `tags` eventual consistency timeout from 5 minutes to 10 minutes ([#23593](https://github.com/hashicorp/terraform-provider-aws/issues/23593))
* resource/aws_kms_key: Increase `description` and `tags` eventual consistency timeouts from 5 minutes to 10 minutes ([#23593](https://github.com/hashicorp/terraform-provider-aws/issues/23593))
* resource/aws_kms_replica_external_key: Increase `tags` eventual consistency timeout from 5 minutes to 10 minutes ([#23593](https://github.com/hashicorp/terraform-provider-aws/issues/23593))
* resource/aws_kms_replica_key: Increase `tags` eventual consistency timeout from 5 minutes to 10 minutes ([#23593](https://github.com/hashicorp/terraform-provider-aws/issues/23593))
* resource/aws_s3_bucket_lifecycle_configuration: Correctly configure `rule.filter.object_size_greater_than` and `rule.filter.object_size_less_than` in API requests and terraform state ([#23441](https://github.com/hashicorp/terraform-provider-aws/issues/23441))
* resource/aws_s3_bucket_lifecycle_configuration: Prevent drift when `rule.noncurrent_version_expiration.newer_noncurrent_versions` or `rule.noncurrent_version_transition.newer_noncurrent_versions` is not specified ([#23441](https://github.com/hashicorp/terraform-provider-aws/issues/23441))
* resource/aws_s3_bucket_replication_configuration: Correctly configure empty `rule.filter` configuration block in API requests ([#23586](https://github.com/hashicorp/terraform-provider-aws/issues/23586))
* resource/aws_s3_bucket_replication_configuration: Ensure both `key` and `value` arguments of the `rule.filter.tag` configuration block are correctly populated in the outgoing API request and terraform state. ([#23579](https://github.com/hashicorp/terraform-provider-aws/issues/23579))
* resource/aws_s3_bucket_replication_configuration: Prevent inconsistent final plan when `rule.filter.prefix` is an empty string ([#23586](https://github.com/hashicorp/terraform-provider-aws/issues/23586))

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
* data-source/aws_eks_node_group: Add `taints` attribute ([#23452](https://github.com/hashicorp/terraform-provider-aws/issues/23452))
* resource/aws_apprunner_service: Add `network_configuration` argument ([#23173](https://github.com/hashicorp/terraform-provider-aws/issues/23173))
* resource/aws_cloudwatch_metric_alarm: Additional allowed values for `extended_statistic` and `metric_query.metric.stat` arguments ([#22942](https://github.com/hashicorp/terraform-provider-aws/issues/22942))
* resource/aws_ec2_transit_gateway: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* resource/aws_ec2_transit_gateway: Add `transit_gateway_cidr_blocks` argument ([#22181](https://github.com/hashicorp/terraform-provider-aws/issues/22181))
* resource/aws_eks_cluster: Retry when `ResourceInUseException` errors are returned from the AWS API during resource deletion ([#23366](https://github.com/hashicorp/terraform-provider-aws/issues/23366))
* resource/aws_glue_job: Add support for [streaming jobs](https://docs.aws.amazon.com/glue/latest/dg/add-job-streaming.html) by removing the default value for the `timeout` argument and marking it as Computed ([#23275](https://github.com/hashicorp/terraform-provider-aws/issues/23275))
* resource/aws_lambda_function: Add support for `dotnet6` `runtime` value ([#23426](https://github.com/hashicorp/terraform-provider-aws/issues/23426))
* resource/aws_lambda_layer_version: Add support for `dotnet6` `compatible_runtimes` value ([#23426](https://github.com/hashicorp/terraform-provider-aws/issues/23426))
* resource/aws_route: `nat_gateway_id` target no longer conflicts with `destination_ipv6_cidr_block` ([#23427](https://github.com/hashicorp/terraform-provider-aws/issues/23427))

BUG FIXES:

* resource/aws_dms_endpoint: Fix bug where KMS key was ignored for DynamoDB, OpenSearch, Kafka, Kinesis, Oracle, PostgreSQL, and S3 engines. ([#23444](https://github.com/hashicorp/terraform-provider-aws/issues/23444))
* resource/aws_networkfirewall_rule_group: Allow any character in `source` and `destination` `rule_group.rules_source.stateful_rule.header` arguments as per the AWS API docs ([#22727](https://github.com/hashicorp/terraform-provider-aws/issues/22727))
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
* resource/aws_route53_resolver_firewall_domain_list: Remove limit for number of `domains`. ([#23485](https://github.com/hashicorp/terraform-provider-aws/issues/23485))
* resource/aws_synthetics_canary: Retry canary creation if it fails because of IAM propagation. ([#23394](https://github.com/hashicorp/terraform-provider-aws/issues/23394))

## 4.3.0 (February 28, 2022)

NOTES:

* resource/aws_internet_gateway: Set `vpc_id` as Computed to prevent drift when the `aws_internet_gateway_attachment` resource is used ([#16386](https://github.com/hashicorp/terraform-provider-aws/issues/16386))
* resource/aws_s3_bucket_lifecycle_configuration: The `prefix` argument of the `rule` configuration block has been deprecated. Use the `filter` configuration block instead. ([#23325](https://github.com/hashicorp/terraform-provider-aws/issues/23325))

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
* resource/aws_connect_queue: The `quick_connect_ids` argument can now be updated in-place ([#22821](https://github.com/hashicorp/terraform-provider-aws/issues/22821))
* resource/aws_connect_security_profile: add `permissions` attribute to read ([#22761](https://github.com/hashicorp/terraform-provider-aws/issues/22761))
* resource/aws_ec2_fleet: Add `context` argument ([#23304](https://github.com/hashicorp/terraform-provider-aws/issues/23304))
* resource/aws_ec2_transit_gateway: Add `multicast_support` argument ([#22756](https://github.com/hashicorp/terraform-provider-aws/issues/22756))
* resource/aws_imagebuilder_image_pipeline: Add `schedule.timezone` argument ([#23322](https://github.com/hashicorp/terraform-provider-aws/issues/23322))
* resource/aws_imagebuilder_image_recipe: Add `systems_manager_agent.uninstall_after_build` argument ([#23293](https://github.com/hashicorp/terraform-provider-aws/issues/23293))
* resource/aws_instance: Prevent double base64 encoding of `user_data` and `user_data_base64` on update ([#23362](https://github.com/hashicorp/terraform-provider-aws/issues/23362))
* resource/aws_s3_bucket: Add error handling for `NotImplemented` error when reading `logging` into terraform state ([#23398](https://github.com/hashicorp/terraform-provider-aws/issues/23398))
* resource/aws_s3_bucket_object_lock_configuration: Mark `token` argument as sensitive ([#23368](https://github.com/hashicorp/terraform-provider-aws/issues/23368))
* resource/aws_servicecatalog_provisioned_product: Add `outputs` attribute ([#23270](https://github.com/hashicorp/terraform-provider-aws/issues/23270))

BUG FIXES:

* provider: Validates names of named profiles before use. ([#23388](https://github.com/hashicorp/terraform-provider-aws/issues/23388))
* resource/aws_dms_replication_task: Allow `cdc_start_position` to be computed ([#23328](https://github.com/hashicorp/terraform-provider-aws/issues/23328))
* resource/aws_ecs_cluster: Fix bug preventing describing clusters in ISO regions ([#23341](https://github.com/hashicorp/terraform-provider-aws/issues/23341))

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
* resource/aws_dms_endpoint: Add `s3_settings.add_column_name`, `s3_settings.canned_acl_for_objects`, `s3_settings.cdc_inserts_and_updates`, `s3_settings.cdc_inserts_only`, `s3_settings.cdc_max_batch_interval`, `s3_settings.cdc_min_file_size`, `s3_settings.cdc_path`, `s3_settings.csv_no_sup_value`, `s3_settings.csv_null_value`, `s3_settings.data_page_size`, `s3_settings.date_partition_delimiter`, `s3_settings.date_partition_sequence`, `s3_settings.dict_page_size_limit`, `s3_settings.enable_statistics`, `s3_settings.encoding_type`, `s3_settings.ignore_headers_row`, `s3_settings.include_op_for_full_load`, `s3_settings.max_file_size`, `s3_settings.preserve_transactions`, `s3_settings.rfc_4180`, `s3_settings.row_group_length`, `s3_settings.timestamp_column_name`, `s3_settings.use_csv_no_sup_value` arguments ([#20913](https://github.com/hashicorp/terraform-provider-aws/issues/20913))
* resource/aws_elasticache_replication_group: Add plan-time validation to `description` and `replication_group_description` to ensure non-empty strings ([#23254](https://github.com/hashicorp/terraform-provider-aws/issues/23254))
* resource/aws_fms_policy: Add `delete_unused_fm_managed_resources` argument ([#21295](https://github.com/hashicorp/terraform-provider-aws/issues/21295))
* resource/aws_fms_policy: Add `tags` argument and `tags_all` attribute to support resource tagging ([#21299](https://github.com/hashicorp/terraform-provider-aws/issues/21299))
* resource/aws_imagebuilder_image_recipe: Update plan time validation of `block_device_mapping.ebs.kms_key_id`, `block_device_mapping.ebs.snapshot_id`, `block_device_mapping.ebs.volume_type`, `name`, `parent_image`. ([#23235](https://github.com/hashicorp/terraform-provider-aws/issues/23235))
* resource/aws_instance: Allow updates to `user_data` and `user_data_base64` without forcing resource replacement ([#18043](https://github.com/hashicorp/terraform-provider-aws/issues/18043))
* resource/aws_s3_bucket: Add error handling for `MethodNotAllowed` and `XNotImplemented` errors when reading `website` into terraform state. ([#23278](https://github.com/hashicorp/terraform-provider-aws/issues/23278))
* resource/aws_s3_bucket: Add error handling for `NotImplemented` errors when reading `acceleration_status`, `policy`, or `request_payer` into terraform state. ([#23278](https://github.com/hashicorp/terraform-provider-aws/issues/23278))

BUG FIXES:

* provider: Credentials with expiry, such as assuming a role, would not renew. ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))
* provider: Setting a custom CA bundle caused the provider to fail. ([#23282](https://github.com/hashicorp/terraform-provider-aws/issues/23282))
* resource/aws_iam_instance_profile: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_openid_connect_provider: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_policy: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_saml_provider: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_server_certificate: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_service_linked_role: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_iam_virtual_mfa_device: Improve tag handling in ISO regions ([#23283](https://github.com/hashicorp/terraform-provider-aws/issues/23283))
* resource/aws_s3_bucket_lifecycle_configuration: Ensure both `key` and `value` arguments of the `filter` `tag` configuration block are correctly populated in the outgoing API request and terraform state. ([#23252](https://github.com/hashicorp/terraform-provider-aws/issues/23252))
* resource/aws_s3_bucket_lifecycle_configuration: Prevent non-empty plans when `filter` is an empty configuration block ([#23232](https://github.com/hashicorp/terraform-provider-aws/issues/23232))

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

* data-source/aws_imagebuilder_image_pipeline: Add `container_recipe_arn` attribute ([#23111](https://github.com/hashicorp/terraform-provider-aws/issues/23111))
* data-source/aws_kms_public_key: Add `public_key_pem` attribute ([#23130](https://github.com/hashicorp/terraform-provider-aws/issues/23130))
* resource/aws_api_gateway_authorizer: Add `arn` attribute. ([#23151](https://github.com/hashicorp/terraform-provider-aws/issues/23151))
* resource/aws_autoscaling_group: Disable scale-in protection before draining instances ([#23187](https://github.com/hashicorp/terraform-provider-aws/issues/23187))
* resource/aws_cloudformation_stack_set: Add `call_as` argument ([#22440](https://github.com/hashicorp/terraform-provider-aws/issues/22440))
* resource/aws_elastic_transcoder_preset: Add plan time validations to `audio.audio_packing_mode`,  `audio.channels`,
`audio.codec`,`audio.sample_rate`, `audio_codec_options.bit_depth`, `audio_codec_options.bit_order`,
`audio_codec_options.profile`, `audio_codec_options.signed`, `audio_codec_options.signed`,
`container`, `thumbnails.aspect_ratio`, `thumbnails.format`, `thumbnails.padding_policy`, `thumbnails.sizing_policy`,
`type`, `video.aspect_ratio`, `video.codec`, `video.display_aspect_ratio`, `video.fixed_gop`, `video.frame_rate`,   `video.max_frame_rate`,  `video.padding_policy`, `video.sizing_policy`, `video_watermarks.horizontal_align`,
`video_watermarks.id`, `video_watermarks.sizing_policy`, `video_watermarks.target`, `video_watermarks.vertical_align` ([#13974](https://github.com/hashicorp/terraform-provider-aws/issues/13974))
* resource/aws_elastic_transcoder_preset: Allow `audio.bit_rate` to be computed. ([#13974](https://github.com/hashicorp/terraform-provider-aws/issues/13974))
* resource/aws_gamelift_build: Add `object_version` argument to `storage_location` block. ([#22966](https://github.com/hashicorp/terraform-provider-aws/issues/22966))
* resource/aws_gamelift_build: Add import support ([#22966](https://github.com/hashicorp/terraform-provider-aws/issues/22966))
* resource/aws_gamelift_fleet: Add `certificate_configuration` argument ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Add import support ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Add plan time validation to `ec2_instance_type` ([#22967](https://github.com/hashicorp/terraform-provider-aws/issues/22967))
* resource/aws_gamelift_fleet: Adds `script_arn` attribute. ([#11560](https://github.com/hashicorp/terraform-provider-aws/issues/11560))
* resource/aws_gamelift_fleet: Adds `script_id` argument. ([#11560](https://github.com/hashicorp/terraform-provider-aws/issues/11560))
* resource/aws_glue_catalog_database: Add support `create_table_default_permission` argument ([#22964](https://github.com/hashicorp/terraform-provider-aws/issues/22964))
* resource/aws_glue_trigger: Add `event_batching_condition` argument. ([#22963](https://github.com/hashicorp/terraform-provider-aws/issues/22963))
* resource/aws_iam_user_login_profile: Make `pgp_key` optional ([#12384](https://github.com/hashicorp/terraform-provider-aws/issues/12384))
* resource/aws_imagebuilder_image_pipeline: Add `container_recipe_arn` argument ([#23111](https://github.com/hashicorp/terraform-provider-aws/issues/23111))
* resource/aws_prometheus_workspace: Add `tags` argument and `tags_all` attribute to support resource tagging ([#23202](https://github.com/hashicorp/terraform-provider-aws/issues/23202))
* resource/aws_ssm_association: Add `arn` attribute ([#17732](https://github.com/hashicorp/terraform-provider-aws/issues/17732))
* resource/aws_ssm_association: Add `wait_for_success_timeout_seconds` argument ([#17732](https://github.com/hashicorp/terraform-provider-aws/issues/17732))
* resource/aws_ssm_association: Add plan time validation to `association_name`, `document_version`, `schedule_expression`, `output_location.s3_bucket_name`, `output_location.s3_key_prefix`, `targets.key`, `targets.values`, `automation_target_parameter_name` ([#17732](https://github.com/hashicorp/terraform-provider-aws/issues/17732))

BUG FIXES:

* data-source/aws_vpc_ipam_pool: error if no pool found ([#23195](https://github.com/hashicorp/terraform-provider-aws/issues/23195))
* provider: Support `ap-northeast-3`, `ap-southeast-3` and `us-iso-west-1` as valid AWS Regions ([#23191](https://github.com/hashicorp/terraform-provider-aws/issues/23191))
* provider: Use AWS HTTP client which allows IMDS authentication in container environments and custom RootCAs in ISO regions ([#23191](https://github.com/hashicorp/terraform-provider-aws/issues/23191))
* resource/aws_appmesh_route: Handle zero `max_retries` ([#23035](https://github.com/hashicorp/terraform-provider-aws/issues/23035))
* resource/aws_elastic_transcoder_preset: Allow `video_codec_options` to be empty. ([#13974](https://github.com/hashicorp/terraform-provider-aws/issues/13974))
* resource/aws_rds_cluster: Fix crash when configured `engine_version` string is shorter than the `EngineVersion` string returned from the AWS API ([#23039](https://github.com/hashicorp/terraform-provider-aws/issues/23039))
* resource/aws_s3_bucket_lifecycle_configuration: Correctly handle the `days` value of the `rule` `transition` configuration block when set to `0` ([#23120](https://github.com/hashicorp/terraform-provider-aws/issues/23120))
* resource/aws_s3_bucket_lifecycle_configuration: Fix extraneous diffs especially after import ([#23144](https://github.com/hashicorp/terraform-provider-aws/issues/23144))
* resource/aws_sagemaker_endpoint_configuration: Emptiness check for arguments, Allow not passing `async_inference_config.kms_key_id`. ([#22960](https://github.com/hashicorp/terraform-provider-aws/issues/22960))
* resource/aws_vpn_connection: Add support for `ipsec.1-aes256` connection type ([#23127](https://github.com/hashicorp/terraform-provider-aws/issues/23127))

## 4.0.0 (February 10, 2022)

BREAKING CHANGES:

* data-source/aws_connect_hours_of_operation: The hours_of_operation_arn attribute is renamed to arn ([#22375](https://github.com/hashicorp/terraform-provider-aws/issues/22375))
* resource/aws_batch_compute_environment: No `compute_resources` configuration block can be specified when `type` is `UNMANAGED` ([#22805](https://github.com/hashicorp/terraform-provider-aws/issues/22805))
* resource/aws_cloudwatch_event_target: The `ecs_target` `launch_type` argument no longer has a default value (previously was `EC2`) ([#22803](https://github.com/hashicorp/terraform-provider-aws/issues/22803))
* resource/aws_cloudwatch_event_target: `ecs_target.0.launch_type` can no longer be set to `""`; instead, remove or set to `null` ([#22954](https://github.com/hashicorp/terraform-provider-aws/issues/22954))
* resource/aws_connect_hours_of_operation: The hours_of_operation_arn attribute is renamed to arn ([#22375](https://github.com/hashicorp/terraform-provider-aws/issues/22375))
* resource/aws_default_network_acl: These arguments can no longer be set to `""`: `egress.*.cidr_block`, `egress.*.ipv6_cidr_block`, `ingress.*.cidr_block`, or `ingress.*.ipv6_cidr_block` ([#22928](https://github.com/hashicorp/terraform-provider-aws/issues/22928))
* resource/aws_default_route_table: These arguments can no longer be set to `""`: `route.*.cidr_block`, `route.*.ipv6_cidr_block` ([#22931](https://github.com/hashicorp/terraform-provider-aws/issues/22931))
* resource/aws_default_vpc: `ipv6_cidr_block` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))
* resource/aws_efs_mount_target: `ip_address` can no longer be set to `""`; instead, remove or set to `null` ([#22954](https://github.com/hashicorp/terraform-provider-aws/issues/22954))
* resource/aws_elasticache_cluster: Either `engine` or `replication_group_id` must be specified ([#20482](https://github.com/hashicorp/terraform-provider-aws/issues/20482))
* resource/aws_elasticsearch_domain: `ebs_options.0.volume_type` can no longer be set to `""`; instead, remove or set to `null` ([#22954](https://github.com/hashicorp/terraform-provider-aws/issues/22954))
* resource/aws_fsx_ontap_storage_virtual_machine: Remove deprecated `active_directory_configuration.0.self_managed_active_directory_configuration.0.organizational_unit_distinguidshed_name`, migrating value to `active_directory_configuration.0.self_managed_active_directory_configuration.0.organizational_unit_distinguished_name` ([#22915](https://github.com/hashicorp/terraform-provider-aws/issues/22915))
* resource/aws_instance: `private_ip` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))
* resource/aws_lb_target_group: For `protocol = "TCP"`, `stickiness` can no longer be type set to `lb_cookie` even when `enabled = false`; instead use type `source_ip` ([#22996](https://github.com/hashicorp/terraform-provider-aws/issues/22996))
* resource/aws_network_acl: These arguments can no longer be set to `""`: `egress.*.cidr_block`, `egress.*.ipv6_cidr_block`, `ingress.*.cidr_block`, or `ingress.*.ipv6_cidr_block` ([#22928](https://github.com/hashicorp/terraform-provider-aws/issues/22928))
* resource/aws_route: Exactly one of these can be set: `destination_cidr_block`, `destination_ipv6_cidr_block`, `destination_prefix_list_id`. These arguments can no longer be set to `""`: `destination_cidr_block`, `destination_ipv6_cidr_block`. ([#22931](https://github.com/hashicorp/terraform-provider-aws/issues/22931))
* resource/aws_route_table: These arguments can no longer be set to `""`: `route.*.cidr_block`, `route.*.ipv6_cidr_block` ([#22931](https://github.com/hashicorp/terraform-provider-aws/issues/22931))
* resource/aws_s3_bucket: The `acceleration_status` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_accelerate_configuration` resource instead. ([#22610](https://github.com/hashicorp/terraform-provider-aws/issues/22610))
* resource/aws_s3_bucket: The `acl` and `grant` arguments have been deprecated and are now read-only. Use the `aws_s3_bucket_acl` resource instead. ([#22537](https://github.com/hashicorp/terraform-provider-aws/issues/22537))
* resource/aws_s3_bucket: The `cors_rule` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_cors_configuration` resource instead. ([#22611](https://github.com/hashicorp/terraform-provider-aws/issues/22611))
* resource/aws_s3_bucket: The `lifecycle_rule` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_lifecycle_configuration` resource instead. ([#22581](https://github.com/hashicorp/terraform-provider-aws/issues/22581))
* resource/aws_s3_bucket: The `logging` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_logging` resource instead. ([#22599](https://github.com/hashicorp/terraform-provider-aws/issues/22599))
* resource/aws_s3_bucket: The `object_lock_configuration` `rule` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_object_lock_configuration` resource instead. ([#22612](https://github.com/hashicorp/terraform-provider-aws/issues/22612))
* resource/aws_s3_bucket: The `policy` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_policy` resource instead. ([#22538](https://github.com/hashicorp/terraform-provider-aws/issues/22538))
* resource/aws_s3_bucket: The `replication_configuration` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_replication_configuration` resource instead. ([#22604](https://github.com/hashicorp/terraform-provider-aws/issues/22604))
* resource/aws_s3_bucket: The `request_payer` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_request_payment_configuration` resource instead. ([#22613](https://github.com/hashicorp/terraform-provider-aws/issues/22613))
* resource/aws_s3_bucket: The `server_side_encryption_configuration` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_server_side_encryption_configuration` resource instead. ([#22605](https://github.com/hashicorp/terraform-provider-aws/issues/22605))
* resource/aws_s3_bucket: The `versioning` argument has been deprecated and is now read-only. Use the `aws_s3_bucket_versioning` resource instead. ([#22606](https://github.com/hashicorp/terraform-provider-aws/issues/22606))
* resource/aws_s3_bucket: The `website`, `website_domain`, and `website_endpoint` arguments have been deprecated and are now read-only. Use the `aws_s3_bucket_website_configuration` resource instead. ([#22614](https://github.com/hashicorp/terraform-provider-aws/issues/22614))
* resource/aws_vpc: `ipv6_cidr_block` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))
* resource/aws_vpc_ipv6_cidr_block_association: `ipv6_cidr_block` can no longer be set to `""`; remove or set to `null` ([#22948](https://github.com/hashicorp/terraform-provider-aws/issues/22948))

NOTES:

* data-source/aws_cognito_user_pools: The type of the `ids` and `arns` attributes has changed from Set to List. If no volumes match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_db_event_categories: The type of the `ids` attribute has changed from Set to List. If no event categories match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ebs_volumes: The type of the `ids` attribute has changed from Set to List. If no volumes match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_coip_pools: The type of the `pool_ids` attribute has changed from Set to List. If no COIP pools match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateway_route_tables: The type of the `ids` attribute has changed from Set to List. If no local gateway route tables match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateway_virtual_interface_groups: The type of the `ids` and `local_gateway_virtual_interface_ids` attributes has changed from Set to List. If no local gateway virtual interface groups match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_local_gateways: The type of the `ids` attribute has changed from Set to List. If no local gateways match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ec2_transit_gateway_route_tables: The type of the `ids` attribute has changed from Set to List. If no transit gateway route tables match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_efs_access_points: The type of the `ids` and `arns` attributes has changed from Set to List. If no access points match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_elasticache_replication_group: The `number_cache_clusters` attribute has been deprecated. All configurations using `number_cache_clusters` should be updated to use the `num_cache_clusters` attribute instead ([#22667](https://github.com/hashicorp/terraform-provider-aws/issues/22667))
* data-source/aws_elasticache_replication_group: The `replication_group_description` attribute has been deprecated. All configurations using `replication_group_description` should be updated to use the `description` attribute instead ([#22667](https://github.com/hashicorp/terraform-provider-aws/issues/22667))
* data-source/aws_emr_release_labels: The type of the `ids` attribute has changed from Set to List. If no release labels match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_iam_policy_document: The `source_json` and `override_json` attributes have been deprecated. Use the `source_policy_documents` and `override_policy_documents` attributes respectively instead. ([#22890](https://github.com/hashicorp/terraform-provider-aws/issues/22890))
* data-source/aws_inspector_rules_packages: If no rules packages match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_instances: If no instances match the specified criteria an empty list is returned (previously an error was raised) ([#5055](https://github.com/hashicorp/terraform-provider-aws/issues/5055))
* data-source/aws_ip_ranges: If no ranges match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_network_acls: The type of the `ids` attribute has changed from Set to List. If no NACLs match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_network_interfaces: The type of the `ids` attribute has changed from Set to List. If no network interfaces match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_route_tables: The type of the `ids` attribute has changed from Set to List. If no route tables match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_s3_bucket_object: The data source is deprecated; use `aws_s3_object` instead ([#22877](https://github.com/hashicorp/terraform-provider-aws/issues/22877))
* data-source/aws_s3_bucket_objects: The data source is deprecated; use `aws_s3_objects` instead ([#22877](https://github.com/hashicorp/terraform-provider-aws/issues/22877))
* data-source/aws_security_groups: If no security groups match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_ssoadmin_instances: The type of the `identity_store_ids` and `arns` attributes has changed from Set to List. If no instances match the specified criteria an empty list is returned (previously an error was raised) ([#21219](https://github.com/hashicorp/terraform-provider-aws/issues/21219))
* data-source/aws_subnet_ids: The `aws_subnet_ids` data source has been deprecated and will be removed in a future version. Use the `aws_subnets` data source instead ([#22743](https://github.com/hashicorp/terraform-provider-aws/issues/22743))
* data-source/aws_vpcs: The type of the `ids` attributes has changed from Set to List. If no VPCs match the specified criteria an empty list is returned (previously an error was raised) ([#22253](https://github.com/hashicorp/terraform-provider-aws/issues/22253))
* provider: The `assume_role.duration_seconds` argument has been deprecated. All configurations using `assume_role.duration_seconds` should be updated to use the new `assume_role.duration` argument instead. ([#23077](https://github.com/hashicorp/terraform-provider-aws/issues/23077))
* resource/aws_acmpca_certificate_authority: The `status` attribute has been deprecated. Use the `enabled` attribute instead. ([#22878](https://github.com/hashicorp/terraform-provider-aws/issues/22878))
* resource/aws_autoscaling_attachment: The `alb_target_group_arn` argument has been deprecated. All configurations using `alb_target_group_arn` should be updated to use the new `lb_target_group_arn` argument instead ([#22662](https://github.com/hashicorp/terraform-provider-aws/issues/22662))
* resource/aws_autoscaling_group: The `tags` argument has been deprecated. All configurations using `tags` should be updated to use the `tag` argument instead ([#22663](https://github.com/hashicorp/terraform-provider-aws/issues/22663))
* resource/aws_budgets_budget: The `cost_filters` attribute has been deprecated. Use the `cost_filter` attribute instead. ([#22888](https://github.com/hashicorp/terraform-provider-aws/issues/22888))
* resource/aws_connect_hours_of_operation: Timeout support has been removed as it is not needed for this resource ([#22375](https://github.com/hashicorp/terraform-provider-aws/issues/22375))
* resource/aws_customer_gateway: `ip_address` can no longer be set to `""` ([#22926](https://github.com/hashicorp/terraform-provider-aws/issues/22926))
* resource/aws_db_instance: The `name` argument has been deprecated. All configurations using `name` should be updated to use the `db_name` argument instead ([#22668](https://github.com/hashicorp/terraform-provider-aws/issues/22668))
* resource/aws_default_subnet: If no default subnet exists in the specified Availability Zone one is now created. The `force_destroy` destroy argument has been added (defaults to `false`). Setting this argument to `true` deletes the default subnet on `terraform destroy` ([#22253](https://github.com/hashicorp/terraform-provider-aws/issues/22253))
* resource/aws_default_vpc: If no default VPC exists in the current AWS Region one is now created. The `force_destroy` destroy argument has been added (defaults to `false`). Setting this argument to `true` deletes the default VPC on `terraform destroy` ([#22253](https://github.com/hashicorp/terraform-provider-aws/issues/22253))
* resource/aws_ec2_client_vpn_endpoint: The `status` attribute has been deprecated ([#22887](https://github.com/hashicorp/terraform-provider-aws/issues/22887))
* resource/aws_ec2_client_vpn_endpoint: The type of the `dns_servers` argument has changed from Set to List ([#22889](https://github.com/hashicorp/terraform-provider-aws/issues/22889))
* resource/aws_ec2_client_vpn_network_association: The `security_groups` argument has been deprecated. Use the `security_group_ids` argument of the `aws_ec2_client_vpn_endpoint` resource instead ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* resource/aws_ec2_client_vpn_network_association: The `status` attribute has been deprecated ([#22887](https://github.com/hashicorp/terraform-provider-aws/issues/22887))
* resource/aws_ec2_client_vpn_route: Add [custom `timeouts`](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) block ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* resource/aws_ecs_cluster: The `capacity_providers` and `default_capacity_provider_strategy` arguments have been deprecated. Use the `aws_ecs_cluster_capacity_providers` resource instead. ([#22783](https://github.com/hashicorp/terraform-provider-aws/issues/22783))
* resource/aws_elasticache_replication_group: The `cluster_mode` argument has been deprecated. All configurations using `cluster_mode` should be updated to use the root-level `num_node_groups` and `replicas_per_node_group` arguments instead ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_elasticache_replication_group: The `number_cache_clusters` argument has been deprecated. All configurations using `number_cache_clusters` should be updated to use the `num_cache_clusters` argument instead ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_elasticache_replication_group: The `replication_group_description` argument has been deprecated. All configurations using `replication_group_description` should be updated to use the `description` argument instead ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_route: The `instance_id` argument has been deprecated. All configurations using `instance_id` should be updated to use the `network_interface_id` argument instead ([#22664](https://github.com/hashicorp/terraform-provider-aws/issues/22664))
* resource/aws_route_table: The `instance_id` argument of the `route` configuration block has been deprecated. All configurations using `route` `instance_id` should be updated to use the `route` `network_interface_id` argument instead ([#22664](https://github.com/hashicorp/terraform-provider-aws/issues/22664))
* resource/aws_s3_bucket_object: The resource is deprecated; use `aws_s3_object` instead ([#22877](https://github.com/hashicorp/terraform-provider-aws/issues/22877))

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

* data-source/aws_ami: Add `boot_mode` attribute. ([#22939](https://github.com/hashicorp/terraform-provider-aws/issues/22939))
* data-source/aws_cloudwatch_log_group: Automatically trim `:*` suffix from `arn` attribute ([#22043](https://github.com/hashicorp/terraform-provider-aws/issues/22043))
* data-source/aws_ec2_client_vpn_endpoint: Add `security_group_ids` and `vpc_id` attributes ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* data-source/aws_elasticache_replication_group: Add `description`, `num_cache_clusters`, `num_node_groups`, and `replicas_per_node_group` attributes ([#22667](https://github.com/hashicorp/terraform-provider-aws/issues/22667))
* data-source/aws_imagebuilder_distribution_configuration: Add `container_distribution_configuration` attribute to the `distribution` configuration block ([#22838](https://github.com/hashicorp/terraform-provider-aws/issues/22838))
* data-source/aws_imagebuilder_distribution_configuration: Add `launch_template_configuration` attribute to the `distribution` configuration block ([#22884](https://github.com/hashicorp/terraform-provider-aws/issues/22884))
* data-source/aws_imagebuilder_image_recipe: Add `parameter` attribute to the `component` configuration block ([#22856](https://github.com/hashicorp/terraform-provider-aws/issues/22856))
* provider: Add `duration` argument to the `assume_role` configuration block ([#23077](https://github.com/hashicorp/terraform-provider-aws/issues/23077))
* provider: Add `ec2_metadata_service_endpoint`, `ec2_metadata_service_endpoint_mode`, `use_dualstack_endpoint`, `use_fips_endpoint` arguments ([#22804](https://github.com/hashicorp/terraform-provider-aws/issues/22804))
* provider: Add environment variables `TF_AWS_DYNAMODB_ENDPOINT`, `TF_AWS_IAM_ENDPOINT`, `TF_AWS_S3_ENDPOINT`, and `TF_AWS_STS_ENDPOINT`. ([#23052](https://github.com/hashicorp/terraform-provider-aws/issues/23052))
* provider: Add support for `shared_config_file` parameter ([#20587](https://github.com/hashicorp/terraform-provider-aws/issues/20587))
* provider: Add support for `shared_credentials_files` parameter and deprecates `shared_credentials_file` ([#23080](https://github.com/hashicorp/terraform-provider-aws/issues/23080))
* provider: Adds `s3_use_path_style` parameter and deprecates `s3_force_path_style`. ([#23055](https://github.com/hashicorp/terraform-provider-aws/issues/23055))
* provider: Changes `shared_config_file` parameter to `shared_config_files` ([#23080](https://github.com/hashicorp/terraform-provider-aws/issues/23080))
* provider: Updates AWS authentication to use AWS SDK for Go v2 <https://aws.github.io/aws-sdk-go-v2/docs/> ([#20587](https://github.com/hashicorp/terraform-provider-aws/issues/20587))
* resource/aws_ami: Add `boot_mode` and `ebs_block_device.outpost_arn` arguments. ([#22939](https://github.com/hashicorp/terraform-provider-aws/issues/22939))
* resource/aws_ami_copy: Add `boot_mode` and `ebs_block_device.outpost_arn` attributes ([#22972](https://github.com/hashicorp/terraform-provider-aws/issues/22972))
* resource/aws_ami_from_instance: Add `boot_mode` and `ebs_block_device.outpost_arn` attributes ([#22972](https://github.com/hashicorp/terraform-provider-aws/issues/22972))
* resource/aws_api_gateway_domain_name: Add `ownership_verification_certificate_arn` argument. ([#21076](https://github.com/hashicorp/terraform-provider-aws/issues/21076))
* resource/aws_apigatewayv2_domain_name: Add `domain_name_configuration.ownership_verification_certificate_arn` argument. ([#21076](https://github.com/hashicorp/terraform-provider-aws/issues/21076))
* resource/aws_autoscaling_attachment: Add `lb_target_group_arn` argument ([#22662](https://github.com/hashicorp/terraform-provider-aws/issues/22662))
* resource/aws_cloudwatch_event_target: Add plan time validation for `input`, `input_path`, `run_command_targets.values`, `http_target.header_parameters`, `http_target.query_string_parameters`, `redshift_target.database`, `redshift_target.db_user`, `redshift_target.secrets_manager_arn`, `redshift_target.sql`, `redshift_target.statement_name`, `retry_policy.maximum_event_age_in_seconds`, `retry_policy.maximum_retry_attempts`. ([#22946](https://github.com/hashicorp/terraform-provider-aws/issues/22946))
* resource/aws_db_instance: Add `db_name` argument ([#22668](https://github.com/hashicorp/terraform-provider-aws/issues/22668))
* resource/aws_ec2_client_vpn_authorization_rule: Configurable Create and Delete timeouts ([#20688](https://github.com/hashicorp/terraform-provider-aws/issues/20688))
* resource/aws_ec2_client_vpn_endpoint: Add `client_connect_options` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `client_login_banner_options` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `security_group_ids` and `vpc_id` arguments ([#22911](https://github.com/hashicorp/terraform-provider-aws/issues/22911))
* resource/aws_ec2_client_vpn_endpoint: Add `session_timeout_hours` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_endpoint: Add `vpn_port` argument ([#22793](https://github.com/hashicorp/terraform-provider-aws/issues/22793))
* resource/aws_ec2_client_vpn_network_association: Configurable Create and Delete timeouts ([#20689](https://github.com/hashicorp/terraform-provider-aws/issues/20689))
* resource/aws_elasticache_replication_group: Add `description` argument ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_elasticache_replication_group: Add `num_cache_clusters` argument ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_elasticache_replication_group: Add `num_node_groups` and `replicas_per_node_group` arguments ([#22666](https://github.com/hashicorp/terraform-provider-aws/issues/22666))
* resource/aws_fsx_lustre_file_system: Add `log_configuration` argument. ([#22935](https://github.com/hashicorp/terraform-provider-aws/issues/22935))
* resource/aws_fsx_ontap_file_system: Reduce the minimum valid value of the `throughput_capacity` argument to `128` (128 MB/s) ([#22898](https://github.com/hashicorp/terraform-provider-aws/issues/22898))
* resource/aws_glue_partition_index: Add support for custom timeouts. ([#22941](https://github.com/hashicorp/terraform-provider-aws/issues/22941))
* resource/aws_imagebuilder_distribution_configuration: Add `launch_template_configuration` argument to the `distribution` configuration block ([#22842](https://github.com/hashicorp/terraform-provider-aws/issues/22842))
* resource/aws_imagebuilder_image_recipe: Add `parameter` argument to the `component` configuration block ([#22837](https://github.com/hashicorp/terraform-provider-aws/issues/22837))
* resource/aws_mq_broker: `auto_minor_version_upgrade` and `host_instance_type` can be changed without recreating broker ([#20661](https://github.com/hashicorp/terraform-provider-aws/issues/20661))
* resource/aws_s3_bucket_cors_configuration: Retry when `NoSuchCORSConfiguration` errors are returned from the AWS API ([#22977](https://github.com/hashicorp/terraform-provider-aws/issues/22977))
* resource/aws_s3_bucket_versioning: Add eventual consistency handling to help ensure bucket versioning is stabilized. ([#21076](https://github.com/hashicorp/terraform-provider-aws/issues/21076))
* resource/aws_vpn_connection: Add the ability to revert changes to unconfigured tunnel options made outside of Terraform to their [documented default values](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html) ([#17031](https://github.com/hashicorp/terraform-provider-aws/issues/17031))
* resource/aws_vpn_connection: Mark `customer_gateway_configuration` as [`Sensitive`](https://www.terraform.io/plugin/sdkv2/best-practices/sensitive-state#using-the-sensitive-flag) ([#15806](https://github.com/hashicorp/terraform-provider-aws/issues/15806))
* resource/aws_wafv2_web_acl: Support `version` on `managed_rule_group_statement` ([#21732](https://github.com/hashicorp/terraform-provider-aws/issues/21732))

BUG FIXES:

* data-source/aws_vpc_peering_connections: Return empty array instead of error when no connections found. ([#17382](https://github.com/hashicorp/terraform-provider-aws/issues/17382))
* resource/aws_cloudformation_stack: Retry resource Create and Update for IAM eventual consistency ([#22840](https://github.com/hashicorp/terraform-provider-aws/issues/22840))
* resource/aws_cloudwatch_event_target: Preserve order of `http_target.path_parameter_values`. ([#22946](https://github.com/hashicorp/terraform-provider-aws/issues/22946))
* resource/aws_db_instance: Fix error with reboot of replica ([#22178](https://github.com/hashicorp/terraform-provider-aws/issues/22178))
* resource/aws_ec2_client_vpn_authorization_rule: Don't raise an error when `InvalidClientVpnEndpointId.NotFound` is returned during refresh ([#20688](https://github.com/hashicorp/terraform-provider-aws/issues/20688))
* resource/aws_ec2_client_vpn_endpoint: `connection_log_options.cloudwatch_log_stream` argument is Computed, preventing spurious resource diffs ([#22891](https://github.com/hashicorp/terraform-provider-aws/issues/22891))
* resource/aws_ecs_capacity_provider: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_cluster: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_service: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_task_definition: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_task_set: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_route_table_association: Handle nil 'AssociationState' in ISO regions ([#22806](https://github.com/hashicorp/terraform-provider-aws/issues/22806))
* resource/aws_route_table_association: Retry resource Read for EC2 eventual consistency ([#22927](https://github.com/hashicorp/terraform-provider-aws/issues/22927))
* resource/aws_vpc_ipam: Correct update of `description` ([#22863](https://github.com/hashicorp/terraform-provider-aws/issues/22863))
* resource/aws_waf_rule_group: Prevent panic when expanding the rule group's set of `activated_rule` ([#22978](https://github.com/hashicorp/terraform-provider-aws/issues/22978))
* resource/aws_wafregional_rule_group: Prevent panic when expanding the rule group's set of `activated_rule` ([#22978](https://github.com/hashicorp/terraform-provider-aws/issues/22978))

## 3.75.2 (May 20, 2022)

ENHANCEMENTS:

* resource/aws_lambda_function: Add support for `nodejs16.x` `runtime` value ([#24874](https://github.com/hashicorp/terraform-provider-aws/issues/24874))
* resource/aws_lambda_layer_version: Add support for `nodejs16.x` `compatible_runtimes` value ([#24874](https://github.com/hashicorp/terraform-provider-aws/issues/24874))
* resource/aws_s3_bucket_website_configuration: Add `routing_rules` parameter to be used instead of `routing_rule` to support configurations with empty String values ([#24199](https://github.com/hashicorp/terraform-provider-aws/issues/24199))

## 3.75.1 (March 24, 2022)

BUG FIXES:

* resource/aws_route_table_association: Retry resource Read for EC2 eventual consistency ([#23806](https://github.com/hashicorp/terraform-provider-aws/issues/23806))

## 3.75.0 (March 18, 2022)

NOTES:

* resource/aws_s3_bucket: The `acceleration_status` argument has been deprecated. Use the `aws_s3_bucket_accelerate_configuration` resource instead. ([#23471](https://github.com/hashicorp/terraform-provider-aws/issues/23471))
* resource/aws_s3_bucket: The `acl` and `grant` arguments have been deprecated. Use the `aws_s3_bucket_acl` resource instead. ([#23419](https://github.com/hashicorp/terraform-provider-aws/issues/23419))
* resource/aws_s3_bucket: The `cors_rule` argument has been deprecated. Use the `aws_s3_bucket_cors_configuration` resource instead. ([#23434](https://github.com/hashicorp/terraform-provider-aws/issues/23434))
* resource/aws_s3_bucket: The `lifecycle_rule` argument has been deprecated. Use the `aws_s3_bucket_lifecycle_configuration` resource instead. ([#23445](https://github.com/hashicorp/terraform-provider-aws/issues/23445))
* resource/aws_s3_bucket: The `logging` argument has been deprecated. Use the `aws_s3_bucket_logging` resource instead. ([#23430](https://github.com/hashicorp/terraform-provider-aws/issues/23430))
* resource/aws_s3_bucket: The `object_lock_configuration.object_lock_enabled` argument has been deprecated. Use the top-level argument `object_lock_enabled` instead. ([#23449](https://github.com/hashicorp/terraform-provider-aws/issues/23449))
* resource/aws_s3_bucket: The `object_lock_configuration.rule` argument has been deprecated. Use the `aws_s3_bucket_object_lock_configuration` resource instead. ([#23449](https://github.com/hashicorp/terraform-provider-aws/issues/23449))
* resource/aws_s3_bucket: The `replication_configuration` argument has been deprecated. Use the `aws_s3_bucket_replication_configuration` resource instead. ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))
* resource/aws_s3_bucket: The `request_payer` argument has been deprecated. Use the `aws_s3_bucket_request_payment_configuration` resource instead. ([#23473](https://github.com/hashicorp/terraform-provider-aws/issues/23473))
* resource/aws_s3_bucket: The `server_side_encryption_configuration` argument has been deprecated. Use the `aws_s3_bucket_server_side_encryption_configuration` resource instead. ([#23476](https://github.com/hashicorp/terraform-provider-aws/issues/23476))
* resource/aws_s3_bucket: The `versioning` argument has been deprecated. Use the `aws_s3_bucket_versioning` resource instead. ([#23432](https://github.com/hashicorp/terraform-provider-aws/issues/23432))
* resource/aws_s3_bucket: The `website`, `website_domain`, and `website_endpoint` arguments have been deprecated. Use the `aws_s3_bucket_website_configuration` resource instead. ([#23435](https://github.com/hashicorp/terraform-provider-aws/issues/23435))

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

* resource/aws_lambda_function: Add support for `dotnet6` `runtime` value ([#23670](https://github.com/hashicorp/terraform-provider-aws/issues/23670))
* resource/aws_lambda_layer_version: Add support for `dotnet6` `compatible_runtimes` value ([#23670](https://github.com/hashicorp/terraform-provider-aws/issues/23670))
* resource/aws_s3_bucket: Add top-level `object_lock_enabled` parameter ([#23449](https://github.com/hashicorp/terraform-provider-aws/issues/23449))
* resource/aws_s3_bucket_acl: Support resource import for S3 bucket names consisting of uppercase letters, underscores, and a maximum of 255 characters ([#23679](https://github.com/hashicorp/terraform-provider-aws/issues/23679))
* resource/aws_s3_bucket_lifecycle_configuration: Support empty string filtering (default behavior of the `aws_s3_bucket.lifecycle_rule` parameter in provider versions prior to v4.0) ([#23750](https://github.com/hashicorp/terraform-provider-aws/issues/23750))
* resource/aws_s3_bucket_replication_configuration: Add `token` field to specify
x-amz-bucket-object-lock-token for enabling replication on object lock enabled
buckets or enabling object lock on an existing bucket. ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))
* resource/aws_s3_bucket_versioning: Add missing support for `Disabled` bucket versioning ([#23731](https://github.com/hashicorp/terraform-provider-aws/issues/23731))

BUG FIXES:

* resource/aws_s3_bucket: Prevent panic when expanding the bucket's list of `cors_rule` ([#7547](https://github.com/hashicorp/terraform-provider-aws/issues/7547))
* resource/aws_s3_bucket_replication_configuration: Change `rule` configuration block to list instead of set ([#23737](https://github.com/hashicorp/terraform-provider-aws/issues/23737))
* resource/aws_s3_bucket_replication_configuration: Correctly configure empty `rule.filter` configuration block in API requests ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))
* resource/aws_s3_bucket_replication_configuration: Ensure both `key` and `value` arguments of the `rule.filter.tag` configuration block are correctly populated in the outgoing API request and terraform state. ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))
* resource/aws_s3_bucket_replication_configuration: Prevent inconsistent final plan when `rule.filter.prefix` is an empty string ([#23716](https://github.com/hashicorp/terraform-provider-aws/issues/23716))
* resource/aws_s3_bucket_replication_configuration: Set `rule.id` as Computed to prevent drift when the value is not configured ([#23737](https://github.com/hashicorp/terraform-provider-aws/issues/23737))

## 3.74.3 (February 17, 2022)

BUG FIXES:

* resource/aws_ecs_capacity_provider: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_cluster: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_service: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_task_definition: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_ecs_task_set: Fix tagging error preventing use in ISO partitions ([#23030](https://github.com/hashicorp/terraform-provider-aws/issues/23030))
* resource/aws_waf_rule_group: Prevent panic when expanding the rule group's set of `activated_rule` ([#22978](https://github.com/hashicorp/terraform-provider-aws/issues/22978))
* resource/aws_wafregional_rule_group: Prevent panic when expanding the rule group's set of `activated_rule` ([#22978](https://github.com/hashicorp/terraform-provider-aws/issues/22978))

## 3.74.2 (February 11, 2022)

BUG FIXES:

* resource/aws_rds_cluster: Fix crash when configured `engine_version` string is shorter than the `EngineVersion` string returned from the AWS API ([#23039](https://github.com/hashicorp/terraform-provider-aws/issues/23039))
* resource/aws_vpn_connection: Add support for `ipsec.1-aes256` connection type ([#23127](https://github.com/hashicorp/terraform-provider-aws/issues/23127))

## 3.74.1 (February 7, 2022)

BUG FIXES:

* resource/aws_backup_selection: Fix permanent diffs for `condition` and `not_resources` arguments causing resource recreation ([#22882](https://github.com/hashicorp/terraform-provider-aws/issues/22882))

## Previous Releases

For information on prior major releases, see their changelogs:

* [3.74 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/3.x/CHANGELOG.md)
* [2.70 and earlier](https://github.com/hashicorp/terraform-provider-aws/blob/release/2.x/CHANGELOG.md)
