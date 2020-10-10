<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core;

interface CounterInterface
{
    /**
     * @param array $labels
     */
    public function inc(array $labels): void;

    /**
     * @param int $val
     * @param array $labels
     */
    public function incBy(int $val, array $labels): void;
}
