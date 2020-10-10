<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core;

interface HistogramInterface
{
    /**
     * @param float $val
     * @param array $labels
     */
    public function observe(float $val, array $labels): void;
}
