<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core\Config;

final class MetricConfig
{
    public const COUNTER_METRIC_TYPE = 'counter';

    public const GAUGE_METRIC_TYPE = 'gauge';

    public const HISTOGRAM_METRIC_TYPE = 'histogram';

    /**
     * @var string
     */
    private $type;

    /**
     * @var string
     */
    private $name;

    /**
     * @var array
     */
    private $labels;

    /**
     * @var string
     */
    private $help;

    /**
     * @var array|null
     */
    private $buckets;

    /**
     * MetricConfig constructor.
     *
     * @param string     $type
     * @param string     $name
     * @param string     $help
     * @param array      $labels
     * @param array|null $buckets
     */
    public function __construct(string $type, string $name, string $help, array $labels, ?array $buckets)
    {
        $this->type = $type;
        $this->name = $name;
        $this->labels = $labels;
        $this->help = $help;
        $this->buckets = $buckets;
    }

    /**
     * @return string
     */
    public function getType(): string
    {
        return $this->type;
    }

    /**
     * @return string
     */
    public function getName(): string
    {
        return $this->name;
    }

    /**
     * @return array
     */
    public function getLabels(): array
    {
        return $this->labels;
    }

    /**
     * @return string
     */
    public function getHelp(): string
    {
        return $this->help;
    }

    /**
     * @return array|null
     */
    public function getBuckets(): ?array
    {
        return $this->buckets;
    }
}