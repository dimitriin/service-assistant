<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core\Config;

interface MetricsConfigProviderInterface
{
    /**
     * @return MetricConfig[]
     */
    public function getMetricsConfig(): array;
}