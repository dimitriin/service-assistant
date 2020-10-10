<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib;

use Dimitriin\Metrics\Core\HistogramInterface;
use Enalean\Prometheus\Histogram as PrometheusHistogram;

final class Histogram implements HistogramInterface
{
    /**
     * @var PrometheusHistogram
     */
    private $histogram;

    /**
     * @var LabelsResolverInterface
     */
    private $labelsResolver;

    /**
     * @param PrometheusHistogram     $histogram
     * @param LabelsResolverInterface $labelsResolver
     */
    public function __construct(PrometheusHistogram $histogram, LabelsResolverInterface $labelsResolver)
    {
        $this->histogram = $histogram;
        $this->labelsResolver = $labelsResolver;
    }

    public function observe(float $val, array $labels): void
    {
        $this->histogram->observe($val, ...$this->labelsResolver->getLabelValues($labels));
    }
}