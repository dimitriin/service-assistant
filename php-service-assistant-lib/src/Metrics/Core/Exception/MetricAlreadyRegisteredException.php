<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core\Exception;

use LogicException;
use Throwable;

final class MetricAlreadyRegisteredException extends LogicException implements ExceptionInterface
{
    /**
     * @var string
     */
    private $metricName;

    /**
     * MetricAlreadyRegisteredException constructor.
     *
     * @param string          $metricName
     * @param Throwable|null $p
     */
    public function __construct(string $metricName, Throwable $p = null)
    {
        $this->metricName = $metricName;

        parent::__construct("Metric {$metricName} already registered", 0, $p);
    }

}