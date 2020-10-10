<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core;

interface GaugeInterface
{
    /**
     * @param float $val
     * @param array $labels
     * @return mixed
     */
    public function set(float $val, array $labels): void;

    /**
     * @param array $labels
     * @return mixed
     */
    public function inc(array $labels): void;

    /**
     * @param float $val
     * @param array $labels
     * @return mixed
     */
    public function incBy(float $val, array $labels): void;

    /**
     * @param array $labels
     * @return mixed
     */
    public function dec(array $labels): void;

    /**
     * @param float $val
     * @param array $labels
     * @return mixed
     */
    public function decBy(float $val, array $labels): void;

}
