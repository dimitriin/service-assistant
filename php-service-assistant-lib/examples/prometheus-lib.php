<?php

require_once __DIR__ . '/../../vendor/autoload.php';

$promRegistry = new \Enalean\Prometheus\Registry\CollectorRegistry(new \Enalean\Prometheus\Storage\InMemoryStore());
$registry = new \Dimitriin\Metrics\PrometheusLib\Registry($promRegistry, new \Enalean\Prometheus\Renderer\RenderTextFormat());

$registry->registerCounter("app_counter", "Application counter", ["app_name"]);
$registry->registerHistogram("app_histogram", "Application histogram", ["app_name"], [0.1, 1, 10]);
$registry->registerGauge("app_gauge", "Application gauge", ["app_name"]);

$registry->getCounter("app_counter")->inc(["app_name" => "myapp"]);
$registry->getCounter("app_counter")->incBy(2, ["app_name" => "myapp"]);

$registry->getGauge("app_gauge")->set(100, ["app_name" => "myapp"]);
$registry->getGauge("app_gauge")->inc(["app_name" => "myapp"]);
$registry->getGauge("app_gauge")->incBy(2, ["app_name" => "myapp"]);
$registry->getGauge("app_gauge")->dec(["app_name" => "myapp"]);
$registry->getGauge("app_gauge")->decBy(2,["app_name" => "myapp"]);

$registry->getHistogram("app_histogram")->observe(0.05, ["app_name" => "myapp"]);
$registry->getHistogram("app_histogram")->observe(0.5, ["app_name" => "myapp"]);
$registry->getHistogram("app_histogram")->observe(5, ["app_name" => "myapp"]);
$registry->getHistogram("app_histogram")->observe(50, ["app_name" => "myapp"]);


echo $registry->renderMetrics();