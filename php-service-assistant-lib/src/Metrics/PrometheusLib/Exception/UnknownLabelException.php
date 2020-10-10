<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib\Exception;

use Dimitriin\Metrics\Core\Exception\ExceptionInterface;
use InvalidArgumentException;
use Throwable;

final class UnknownLabelException extends InvalidArgumentException implements ExceptionInterface
{
    /**
     * @var string
     */
    private $metricName;

    /**
     * @var string
     */
    private $label;

    public function __construct(string $metricName, string $label, Throwable $previous = null)
    {
        $this->metricName = $metricName;
        $this->label = $label;

        parent::__construct("Unknown label {$label} for metric {$metricName}", 0, $previous);
    }

    /**
     * @return string
     */
    public function getMetricName(): string
    {
        return $this->metricName;
    }

    /**
     * @return string
     */
    public function getLabel(): string
    {
        return $this->label;
    }
}