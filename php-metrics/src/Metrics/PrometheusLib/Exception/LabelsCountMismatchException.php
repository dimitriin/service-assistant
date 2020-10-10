<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\PrometheusLib\Exception;

use Dimitriin\Metrics\Core\Exception\ExceptionInterface;
use InvalidArgumentException;
use Throwable;

final class LabelsCountMismatchException extends InvalidArgumentException implements ExceptionInterface
{
    /**
     * @var string
     */
    private $metricName;

    public function __construct(string $metricName, Throwable $previous = null)
    {
        $this->metricName = $metricName;

        parent::__construct("Labels count for metric {$metricName} mismatch", 0, $previous);
    }

    /**
     * @return string
     */
    public function getMetricName(): string
    {
        return $this->metricName;
    }
}