<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib;

use Dimitriin\Metrics\Core\GaugeInterface;
use Enalean\Prometheus\Gauge as PrometheusGauge;

final class Gauge implements GaugeInterface
{
    /**
     * @var PrometheusGauge
     */
    private $gauge;

    /**
     * @var LabelsResolverInterface
     */
    private $labelsResolver;

    /**
     * @param PrometheusGauge         $gauge
     * @param LabelsResolverInterface $labelsResolver
     */
    public function __construct(PrometheusGauge $gauge, LabelsResolverInterface $labelsResolver)
    {
        $this->gauge = $gauge;
        $this->labelsResolver = $labelsResolver;
    }

    public function set(float $val, array $labels): void
    {
        $this->gauge->set($val, ...$this->labelsResolver->getLabelValues($labels));
    }

    public function inc(array $labels): void
    {
        $this->gauge->inc(...$this->labelsResolver->getLabelValues($labels));
    }

    public function incBy(float $val, array $labels): void
    {
        $this->gauge->incBy($val, ...$this->labelsResolver->getLabelValues($labels));
    }

    public function dec(array $labels): void
    {
        $this->gauge->dec(...$this->labelsResolver->getLabelValues($labels));
    }

    public function decBy(float $val, array $labels): void
    {
        $this->gauge->decBy($val, ...$this->labelsResolver->getLabelValues($labels));
    }

}