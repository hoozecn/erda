UPDATE `sp_metric_expression` SET `enable` = 0 WHERE `id` = (SELECT `id` FROM (SELECT `id` FROM `sp_metric_expression` WHERE `enable` = 1 AND `expression` LIKE "%\"alias\":\"trace\"%") AS `ids`);
INSERT `sp_metric_expression`(`attributes`,`expression`,`version`) VALUES("{}","{\"alias\":\"trace\",\"filter\":{},\"functions\":[{\"aggregator\":\"min\",\"alias\":\"start_time\",\"field\":\"start_time\"},{\"aggregator\":\"max\",\"alias\":\"end_time\",\"field\":\"end_time\"},{\"aggregator\":\"distinct\",\"alias\":\"components\",\"field\":\"components\",\"field_script\":\"function invoke(field, tag){ if(tag.span_layer=='http' || tag.span_layer=='rpc') return tag.service_name; return tag.component; }\"},{\"aggregator\":\"sum\",\"field\":\"errors\",\"field_script\":\"function invoke(field, tag){ return tag.error=='true'?1:0; }\"},{\"aggregator\":\"distinct\",\"alias\":\"application_names\",\"field\":\"applications\",\"field_script\":\"function invoke(field, tag){ return tag.application_name; }\"},{\"aggregator\":\"distinct\",\"alias\":\"applications_ids\",\"field\":\"applications_ids\",\"field_script\":\"function invoke(field, tag){ return tag.application_id; }\"},{\"aggregator\":\"distinct\",\"alias\":\"service_names\",\"field\":\"service_names\",\"field_script\":\"function invoke(field, tag){ return tag.service_name; }\"},{\"aggregator\":\"distinct\",\"alias\":\"service_ids\",\"field\":\"service_ids\",\"field_script\":\"function invoke(field, tag){ return tag.service_id; }\"},{\"aggregator\":\"distinct\",\"alias\":\"dubbo_methods\",\"field\":\"dubbo_methods\",\"field_script\":\"function invoke(field, tag){ return tag.dubbo_method; }\"},{\"aggregator\":\"distinct\",\"alias\":\"http_paths\",\"field\":\"http_paths\",\"field_script\":\"function invoke(field, tag){ return tag.http_path; }\"},{\"aggregator\":\"distinct\",\"alias\":\"terminus_keys\",\"field\":\"terminus_keys\",\"field_script\":\"function invoke(field, tag){ return tag.terminus_key; }\"},{\"aggregator\":\"value\",\"alias\":\"duration\",\"field\":\"duration\",\"field_script\":\"function invoke(field, tag){ return field.end_time - field.start_time; }\",\"trigger\":\"aggregated\"},{\"aggregator\":\"count\",\"alias\":\"span_count\",\"field\":\"span_count\",\"field_script\":\"function invoke(field, tag){ return 1; }\"}],\"group\":[\"trace_id\"],\"metric\":\"span\",\"outputs\":[\"metric\"],\"select\":{\"cluster_name\":\"#cluster_name\",\"org_id\":\"#org_id\",\"org_name\":\"#org_name\",\"project_id\":\"#project_id\",\"project_name\":\"#project_name\",\"trace_id\":\"#trace_id\",\"workspace\":\"#workspace\"},\"window\":3}","3.0");