<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core;

interface RegistryInterface
{
    /**
     * @param string $name
     * @param string $help
     * @param array $labels
     */
    public function registerCounter(string $name, string $help, array $labels): void;

    /**
     * @param string $name
     * @param string $help
     * @param array $labels
     */
    public function registerGauge(string $name, string $help, array $labels): void;

    /**
     * @param string $name
     * @param string $help
     * @param array $labels
     * @param array $buckets
     */
    public function registerHistogram(string $name, string $help, array $labels, array $buckets): void;

    /**
     * @param string $name
     * @return CounterInterface
     */
    public function getCounter(string $name): CounterInterface;

    /**
     * @param string $name
     * @return GaugeInterface
     */
    public function getGauge(string $name): GaugeInterface;

    /**
     * @param string $name
     * @return HistogramInterface
     */
    public function getHistogram(string $name): HistogramInterface;
}
