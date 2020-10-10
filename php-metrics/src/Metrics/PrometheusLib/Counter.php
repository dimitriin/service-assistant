<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib;

use Dimitriin\Metrics\Core\CounterInterface;
use Enalean\Prometheus\Counter as PrometheusCounter;

final class Counter implements CounterInterface
{
    /**
     * @var PrometheusCounter
     */
    private $counter;

    /**
     * @var LabelsResolverInterface
     */
    private $labelsResolver;

    /**
     * Counter constructor.
     *
     * @param PrometheusCounter       $counter
     * @param LabelsResolverInterface $labelsResolver
     */
    public function __construct(PrometheusCounter $counter, LabelsResolverInterface $labelsResolver)
    {
        $this->counter = $counter;
        $this->labelsResolver = $labelsResolver;
    }

    public function inc(array $labels): void
    {
        $this->counter->inc(...$this->labelsResolver->getLabelValues($labels));
    }

    public function incBy(int $val, array $labels): void
    {
        $this->counter->incBy(floatval($val), ...$this->labelsResolver->getLabelValues($labels));
    }
}