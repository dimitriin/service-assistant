<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib;

use Dimitriin\Metrics\Core\CounterInterface;
use Dimitriin\Metrics\Core\Exception\MetricAlreadyRegisteredException;
use Dimitriin\Metrics\Core\Exception\MetricNotRegisteredException;
use Dimitriin\Metrics\Core\GaugeInterface;
use Dimitriin\Metrics\Core\HistogramInterface;
use Dimitriin\Metrics\Core\RegistryInterface;
use Dimitriin\Metrics\Core\RendererInterface;
use Enalean\Prometheus\Registry\Registry as PrometheusRegistry;
use Enalean\Prometheus\Renderer\RenderTextFormat;
use Enalean\Prometheus\Value\HistogramLabelNames;
use Enalean\Prometheus\Value\MetricLabelNames;
use Enalean\Prometheus\Value\MetricName;

final class Registry implements RegistryInterface, RendererInterface
{
    /**
     * @var PrometheusRegistry
     */
    private $registry;

    /**
     * @var RenderTextFormat
     */
    private $renderer;

    private $metrics;

    /**
     * Registry constructor.
     *
     * @param PrometheusRegistry $registry
     * @param RenderTextFormat   $renderer
     */
    public function __construct(PrometheusRegistry $registry, RenderTextFormat $renderer)
    {
        $this->registry = $registry;
        $this->renderer = $renderer;
    }

    public function registerCounter(string $name, string $help, array $labels): void
    {
        $this->assertMetricNotRegistered($name);

        $this->metrics[$name] = new Counter(
            $this->registry->registerCounter(
                MetricName::fromName($name),
                $help,
                MetricLabelNames::fromNames(...$labels)
            ),
            new LabelsResolver($name, $labels)
        );
    }

    public function registerGauge(string $name, string $help, array $labels): void
    {
        $this->assertMetricNotRegistered($name);

        $this->metrics[$name] = new Gauge(
            $this->registry->registerGauge(
                MetricName::fromName($name),
                $help,
                MetricLabelNames::fromNames(...$labels)
            ),
            new LabelsResolver($name, $labels)
        );
    }

    public function registerHistogram(string $name, string $help, array $labels, array $buckets): void
    {
        $this->assertMetricNotRegistered($name);

        $this->metrics[$name] = new Histogram(
            $this->registry->registerHistogram(
                MetricName::fromName($name),
                $help,
                HistogramLabelNames::fromNames(...$labels),
                $buckets
            ),
            new LabelsResolver($name, $labels)
        );
    }

    public function getCounter(string $name): CounterInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    public function getGauge(string $name): GaugeInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    public function getHistogram(string $name): HistogramInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    public function renderMetrics(): string
    {
        return $this->renderer->render($this->registry->getMetricFamilySamples());
    }

    private function assertMetricNotRegistered(string $name): void {
        if(isset($this->metrics[$name])) {
            throw new MetricAlreadyRegisteredException($name);
        }
    }

    private function assertMetricRegistered(string $name): void {
        if(!isset($this->metrics[$name])) {
            throw new MetricNotRegisteredException($name);
        }
    }
}