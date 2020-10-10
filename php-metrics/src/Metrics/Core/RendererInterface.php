<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\Core;

interface RendererInterface
{
    /**
     * @return string
     */
    public function renderMetrics(): string;
}
