<?php
declare(strict_types=1);

namespace Dimitriin\Probes;

interface HealthBitSenderInterface
{
    public function send(?int $ttl = null): void;
}