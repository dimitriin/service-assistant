<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib;

use Dimitriin\Metrics\PrometheusLib\Exception\LabelsCountMismatchException;
use Dimitriin\Metrics\PrometheusLib\Exception\UnknownLabelException;
use SplFixedArray;

final class LabelsResolver implements LabelsResolverInterface
{
    /**
     * @var string
     */
    private $metricName;

    /**
     * @var array
     */
    private $labelsIndex;

    /**
     * LabelsResolver constructor.
     *
     * @param string   $metricName
     * @param string[] $labelNames
     */
    public function __construct(string $metricName, array $labelNames) {
        $this->metricName = $metricName;

        for ($i = 0; $i < count($labelNames); $i++) {
            $this->labelsIndex[$labelNames[$i]] = $i;
        }
    }

    /**
     * @param array $labelsMap
     * @return string[]
     */
    public function getLabelValues(array $labelsMap): array
    {
        if (count($this->labelsIndex) != count($labelsMap)) {
            throw new LabelsCountMismatchException($this->metricName);
        }

       $labelValues = new SplFixedArray(count($this->labelsIndex));

        foreach ($labelsMap as $name => $value) {
            if (!array_key_exists($name, $this->labelsIndex)) {
                throw new UnknownLabelException($this->metricName, $name);
            }

            $labelValues[$this->labelsIndex[$name]] = $value;
        }

        return $labelValues->toArray();
    }
}
